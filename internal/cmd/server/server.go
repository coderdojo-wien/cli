package server

import (
	"github.com/hetznercloud/cli/internal/state"
	"github.com/spf13/cobra"
)

func NewCommand(cli *state.State) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "server",
		Short:                 "Manage servers",
		Args:                  cobra.NoArgs,
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
	}
	cmd.AddCommand(
		newListCommand(cli),
		newDescribeCommand(cli),
		newCreateCommand(cli),
		newDeleteCommand(cli),
		newRebootCommand(cli),
		newPoweronCommand(cli),
		newPoweroffCommand(cli),
		newResetCommand(cli),
		newShutdownCommand(cli),
		newCreateImageCommand(cli),
		newResetPasswordCommand(cli),
		newEnableRescueCommand(cli),
		newDisableRescueCommand(cli),
		newAttachISOCommand(cli),
		newDetachISOCommand(cli),
		newUpdateCommand(cli),
		newChangeTypeCommand(cli),
		newRebuildCommand(cli),
		newEnableBackupCommand(cli),
		newDisableBackupCommand(cli),
		newEnableProtectionCommand(cli),
		newDisableProtectionCommand(cli),
		newSSHCommand(cli),
		newAddLabelCommand(cli),
		newRemoveLabelCommand(cli),
		newSetRDNSCommand(cli),
		newAttachToNetworkCommand(cli),
		newDetachFromNetworkCommand(cli),
		newChangeAliasIPsCommand(cli),
		newIPCommand(cli),
		newRequestConsoleCommand(cli),
		newMetricsCommand(cli),
		//newApplyFirewallCommand(cli),
		//newRemoveFirewallCommand(cli),
	)
	return cmd
}
