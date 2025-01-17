package va

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/sailpoint-oss/sailpoint-cli/internal/terminal"
	"github.com/sailpoint-oss/sailpoint-cli/internal/va"
	"github.com/spf13/cobra"
)

func newUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "update",
		Short:   "update a va",
		Long:    "update a Virtual Appliance.",
		Example: "sail va update 10.10.10.10 10.10.10.11",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var credentials []string
			for credential := 0; credential < len(args); credential++ {
				password, _ := terminal.PromptPassword(fmt.Sprintf("Enter Password for %v:", args[credential]))
				credentials = append(credentials, password)
			}
			for i := 0; i < len(args); i++ {
				endpoint := args[i]
				fmt.Printf("Starting update for %v\n", endpoint)
				password := credentials[i]
				_, updateErr := va.RunVACmd(endpoint, password, UpdateCommand)
				if updateErr != nil {
					return updateErr
				} else {
					color.Green("Initiating update check and install (%v)", endpoint)
				}
				reboot, rebootErr := va.RunVACmd(endpoint, password, RebootCommand)
				if rebootErr != nil {
					color.Green("Rebooting Virtual Appliance (%v)", endpoint)
				} else {
					color.Red(reboot)
				}
			}
			return nil
		},
	}
	return cmd
}
