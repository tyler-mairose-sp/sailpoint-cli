// Copyright (c) 2021, SailPoint Technologies, Inc. All rights reserved.
package connector

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/sailpoint-oss/sailpoint-cli/internal/client"
	"github.com/sailpoint-oss/sailpoint-cli/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

const (
	connectorsEndpoint = "/beta/platform-connectors"
)

func NewConnCmd() *cobra.Command {
	conn := &cobra.Command{
		Use:     "connectors",
		Short:   "manage connectors",
		Aliases: []string{"conn"},
		Run: func(command *cobra.Command, args []string) {
			_, _ = fmt.Fprintf(command.OutOrStdout(), command.UsageString())
		},
	}

	Config, err := config.GetConfig()
	if err != nil {
		cobra.CheckErr(err)
	}

	Client := client.NewSpClient(Config)

	conn.PersistentFlags().StringP("conn-endpoint", "e", connectorsEndpoint, "Override connectors endpoint")

	conn.AddCommand(
		newConnInitCmd(),
		newConnListCmd(Client),
		newConnGetCmd(Client),
		newConnUpdateCmd(Client),
		newConnCreateCmd(Client),
		newConnCreateVersionCmd(Client),
		newConnVersionsCmd(Client),
		newConnInvokeCmd(Client),
		newConnValidateCmd(Client),
		newConnTagCmd(Client),
		newConnValidateSourcesCmd(Client),
		newConnLogsCmd(Client),
		newConnStatsCmd(Client),
		newConnDeleteCmd(Client),
	)

	return conn
}

type devConfig struct {
	ID     string                 `yaml:"id"`
	Config map[string]interface{} `yaml:"config"`
}

func bindDevConfig(flags *pflag.FlagSet) {
	cfg := &devConfig{}
	raw, err := os.ReadFile(".dev.yaml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(raw, cfg)
	if err != nil {
		log.Printf("Failed to unmarshal '.dev.yaml': %s", err)
		return
	}

	if cfg.ID != "" {
		f := flags.Lookup("id")
		if f != nil && !f.Changed {
			flags.Set("id", cfg.ID)
		}
	}

	if len(cfg.Config) > 0 {
		f := flags.Lookup("config-json")
		if f != nil && !f.Changed {
			raw, err := json.Marshal(cfg.Config)
			if err != nil {
				panic(fmt.Sprintf("Failed to encode config as json: %s", err))
			}
			flags.Set("config-json", string(raw))
		}
	}
}
