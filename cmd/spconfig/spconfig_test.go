// Copyright (c) 2021, SailPoint Technologies, Inc. All rights reserved.

package spconfig

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sailpoint-oss/sailpoint-cli/internal/config"
)

func TestNewSPConfigCommand(t *testing.T) {
	err := config.InitConfig()
	if err != nil {
		t.Fatalf("Error initializing config: %v", err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cmd := NewSPConfigCmd()

	b := new(bytes.Buffer)
	cmd.SetOut(b)

	err = cmd.Execute()
	if err != nil {
		t.Fatalf("TestNewCreateCmd: Unable to execute the command successfully: %v", err)
	}
}
