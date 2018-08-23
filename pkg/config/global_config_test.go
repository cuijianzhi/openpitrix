// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package config_test

import (
	"testing"

	"openpitrix.io/openpitrix/pkg/config"
	"openpitrix.io/openpitrix/pkg/logger"
)

func Example_unmarshalInitConfig() {
	globalConfig := config.DecodeInitConfig()
	logger.Info(nil, "Got global config: \n%+v\n", globalConfig)
	logger.Info(nil, "Get global config string: \n%s\n", config.EncodeGlobalConfig(globalConfig))
}

func TestGlobalConfig_GetRuntimeImageIdAndUrl(t *testing.T) {
	globalConfig := config.DecodeInitConfig()
	logger.Info(nil, "Got global config: \n%+v\n", globalConfig)
	logger.Info(nil, "Get global config string: \n%s\n", config.EncodeGlobalConfig(globalConfig))

	imaggConf, err := globalConfig.GetRuntimeImageIdAndUrl("ec2.us-east-1.amazonaws.com", "us-east-1")
	if err != nil {
		t.Fatal(err)
	}
	logger.Info(nil, "%+v", imaggConf)
}
