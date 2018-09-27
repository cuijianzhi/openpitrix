// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package app

import (
	"context"
	"fmt"
	"sort"
	"time"

	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/db"
	"openpitrix.io/openpitrix/pkg/logger"
	"openpitrix.io/openpitrix/pkg/models"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/pi"
	"openpitrix.io/openpitrix/pkg/repoiface"
	"openpitrix.io/openpitrix/pkg/repoiface/wrapper"
	"openpitrix.io/openpitrix/pkg/service/category/categoryutil"
	"openpitrix.io/openpitrix/pkg/util/stringutil"
)

type repoProxy struct {
	repo *pb.Repo
}

func newRepoProxy(repo *pb.Repo) *repoProxy {
	vp := &repoProxy{}
	vp.repo = repo
	return vp
}

func (rp *repoProxy) deleteAppVersions(ctx context.Context) error {
	repoId := rp.repo.RepoId
	_, err := pi.Global().DB(ctx).
		Update(constants.TableApp).
		Set(constants.ColumnStatus, constants.StatusDeleted).
		Where(db.Eq(constants.ColumnRepoId, repoId)).
		Exec()
	if err != nil {
		return err
	}
	_, err = pi.Global().DB(ctx).
		Update(constants.TableAppVersion).
		Set(constants.ColumnStatus, constants.StatusDeleted).
		Where(db.Eq(constants.ColumnRepoId, repoId)).
		Exec()
	return err
}

func (rp *repoProxy) SyncRepo(ctx context.Context) error {
	repo := rp.repo
	if repo.Status.GetValue() == constants.StatusDeleted {
		return rp.deleteAppVersions(ctx)
	}
	reader, err := repoiface.NewReader(ctx, repo)
	if err != nil {
		return err
	}
	indexFile, err := reader.GetIndex(ctx)
	if err != nil {
		return err
	}
	var appIds []string
	for appName, appVersions := range indexFile.GetEntries() {
		var appId string
		logger.Debug(ctx, "Start index app [%s]", appName)
		logger.Debug(ctx, "App [%s] has [%d] versions", appName, appVersions.Len())
		if len(appVersions) == 0 {
			return fmt.Errorf("failed to sync app [%s], no versions", appName)
		}
		sort.Sort(appVersions)
		appId, err := rp.syncAppInfo(ctx, appVersions[0])
		if err != nil {
			logger.Error(ctx, "Failed to sync app [%s] to app info", appName)
			return err
		}
		logger.Info(ctx, "Sync [%s] to app [%s] success", appName, appId)
		var versionIds []string
		for _, appVersion := range appVersions {
			var versionId string
			versionId, err = rp.syncAppVersionInfo(ctx, appId, appVersion)
			if err != nil {
				logger.Error(ctx, "Failed to sync app version [%s] to app version", appVersion.GetAppVersion())
				return err
			}
			logger.Debug(ctx, "App version [%s] sync to app version [%s]", appVersion.GetVersion(), versionId)
			versionIds = append(versionIds, versionId)
		}
		err = clearAppVersions(ctx, appId, versionIds)
		if err != nil {
			return err
		}
		err = resortAppVersions(ctx, appId)
		if err != nil {
			return err
		}
		err = syncAppStatus(ctx, appId)
		if err != nil {
			return err
		}
		appIds = append(appIds, appId)
	}
	err = clearApps(ctx, rp.repo.RepoId.GetValue(), appIds)
	if err != nil {
		return err
	}
	return nil
}

func (rp *repoProxy) syncAppInfo(ctx context.Context, appIface wrapper.VersionInterface) (string, error) {
	chartName := appIface.GetName()
	repoId := rp.repo.GetRepoId().GetValue()

	var enabledCategoryIds []string
	var disabledCategoryIds []string

	for _, c := range rp.repo.GetCategorySet() {
		switch c.Status.GetValue() {
		case constants.StatusEnabled:
			enabledCategoryIds = append(enabledCategoryIds, c.CategoryId.GetValue())
		case constants.StatusDisabled:
			disabledCategoryIds = append(disabledCategoryIds, c.CategoryId.GetValue())
		}
	}
	if len(enabledCategoryIds) == 0 {
		enabledCategoryIds = append(enabledCategoryIds, models.UncategorizedId)
	}

	var appId string
	var app = &models.App{}
	err := pi.Global().DB(ctx).
		Select(models.AppColumns...).
		From(constants.TableApp).
		Where(db.Eq(constants.ColumnRepoId, repoId)).
		Where(db.Eq(constants.ColumnChartName, chartName)).
		LoadOne(&app)
	if err != nil {
		// insert new
		if err != db.ErrNotFound {
			return appId, err
		}
		app = models.NewApp(
			chartName,
			repoId,
			"",
			rp.repo.GetOwner().GetValue(),
			chartName,
		)
		_, err = pi.Global().DB(ctx).
			InsertInto(constants.TableApp).
			Record(app).
			Exec()
		if err != nil {
			return appId, err
		}

		err = categoryutil.SyncResourceCategories(
			ctx,
			pi.Global().DB(ctx),
			app.AppId,
			enabledCategoryIds,
		)
		if err != nil {
			return appId, err
		}

		return app.AppId, err
	}
	appId = app.AppId
	// update exists, only need sync categories
	appCategories, err := getAppCategories(ctx, appId)
	if err != nil {
		return appId, err
	}
	var categoryMap = make(map[string]bool)
	for _, c := range appCategories {
		categoryId := c.GetCategoryId().GetValue()
		// app follow repo's categories:
		// if repo *disable* some categories, app MUST *disable* it
		// if repo *enable*  some categories, app MUST *enable*  it
		if c.GetStatus().GetValue() == constants.StatusEnabled {
			if !stringutil.StringIn(categoryId, disabledCategoryIds) {
				categoryMap[categoryId] = true
			}
		}
	}
	for _, c := range enabledCategoryIds {
		categoryMap[c] = true
	}
	var categoryIds []string
	for c := range categoryMap {
		if c == models.UncategorizedId && len(categoryMap) > 1 {
			continue
		}
		categoryIds = append(categoryIds, c)
	}
	err = categoryutil.SyncResourceCategories(
		ctx,
		pi.Global().DB(ctx),
		appId,
		categoryIds,
	)
	return app.AppId, err
}

func (rp *repoProxy) syncAppVersionInfo(ctx context.Context, appId string, versionInterface wrapper.VersionInterface) (string, error) {
	versionName := versionInterface.GetVersionName()
	var appVersion = &models.AppVersion{}
	var versionId = ""
	err := pi.Global().DB(ctx).
		Select(models.AppVersionColumns...).
		From(constants.TableAppVersion).
		Where(db.Eq(constants.ColumnAppId, appId)).
		Where(db.Eq(constants.ColumnName, versionName)).
		LoadOne(&appVersion)
	if err != nil {
		if err != db.ErrNotFound {
			return versionId, err
		}
		// not found version, create new
		appVersion = models.NewAppVersion(
			appId,
			versionName,
			versionInterface.GetDescription(),
			rp.repo.GetOwner().GetValue(),
			versionInterface.GetPackageName(),
		)

		appVersion.PackageName = versionInterface.GetPackageName()
		appVersion.Description = versionInterface.GetDescription()
		appVersion.Home = versionInterface.GetHome()
		appVersion.Icon = versionInterface.GetIcon()
		appVersion.Screenshots = versionInterface.GetScreenshots()
		appVersion.Maintainers = versionInterface.GetMaintainers()
		appVersion.Keywords = versionInterface.GetKeywords()
		appVersion.Sources = versionInterface.GetSources()

		appVersion.Status = getAppDefaultStatus(rp.repo)

		_, err = pi.Global().DB(ctx).
			InsertInto(constants.TableAppVersion).
			Record(appVersion).
			Exec()
		if err != nil {
			return versionId, err
		}
	}
	// update exists
	versionId = appVersion.VersionId
	var updateAttr = make(map[string]interface{})

	if appVersion.Status != rp.repo.GetAppDefaultStatus().GetValue() {
		updateAttr[constants.ColumnStatus] = getAppVersionStatus(
			getAppDefaultStatus(rp.repo),
			appVersion.Status,
		)
	}

	if appVersion.PackageName != versionInterface.GetPackageName() {
		updateAttr[constants.ColumnPackageName] = versionInterface.GetPackageName()
	}
	if appVersion.Description != versionInterface.GetDescription() {
		updateAttr[constants.ColumnDescription] = versionInterface.GetDescription()
	}
	if appVersion.Home != versionInterface.GetHome() {
		updateAttr[constants.ColumnHome] = versionInterface.GetHome()
	}
	if appVersion.Icon != versionInterface.GetIcon() {
		updateAttr[constants.ColumnIcon] = versionInterface.GetIcon()
	}
	if appVersion.Screenshots != versionInterface.GetScreenshots() {
		updateAttr[constants.ColumnScreenshots] = versionInterface.GetScreenshots()
	}
	if appVersion.Maintainers != versionInterface.GetMaintainers() {
		updateAttr[constants.ColumnMaintainers] = versionInterface.GetMaintainers()
	}
	if appVersion.Keywords != versionInterface.GetKeywords() {
		updateAttr[constants.ColumnKeywords] = versionInterface.GetKeywords()
	}
	if appVersion.Sources != versionInterface.GetSources() {
		updateAttr[constants.ColumnSources] = versionInterface.GetSources()
	}
	if len(updateAttr) == 0 {
		return versionId, nil
	}
	_, err = pi.Global().DB(ctx).
		Update(constants.TableAppVersion).
		SetMap(updateAttr).
		Set(constants.ColumnUpdateTime, time.Now()).
		Where(db.Eq(constants.ColumnVersionId, versionId)).
		Exec()
	return versionId, err
}