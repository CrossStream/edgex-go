//
// Copyright (c) 2018
// Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package main

import (
	"testing"

	"github.com/edgexfoundry/edgex-go/internal/export/client"
	"github.com/edgexfoundry/edgex-go/internal/pkg/config"
)

func TestToml(t *testing.T) {
	configuration := &client.ConfigurationStruct{}
	if err := config.VerifyTomlFiles(configuration); err != nil {
		t.Fatalf("%v", err)
	}
	if configuration.Service.StartupMsg == "" {
		t.Errorf("configuration.StartupMsg is zero length.")
	}
}
