package sshkey

import (
	"errors"
	"fmt"

	"github.com/hetznercloud/cli/internal/cmd/cmpl"
	"github.com/hetznercloud/cli/internal/state"
	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/spf13/cobra"
)

func newUpdateCommand(cli *state.State) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "update [FLAGS] SSHKEY",
		Short:                 "Update a SSH key",
		Args:                  cobra.ExactArgs(1),
		ValidArgsFunction:     cmpl.SuggestArgs(cmpl.SuggestCandidatesF(cli.SSHKeyNames)),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		PreRunE:               cli.EnsureToken,
		RunE:                  cli.Wrap(runUpdate),
	}

	cmd.Flags().String("name", "", "SSH key name")

	return cmd
}

func runUpdate(cli *state.State, cmd *cobra.Command, args []string) error {
	idOrName := args[0]
	sshKey, _, err := cli.Client().SSHKey.Get(cli.Context, idOrName)
	if err != nil {
		return err
	}
	if sshKey == nil {
		return fmt.Errorf("SSH key not found: %s", idOrName)
	}

	name, _ := cmd.Flags().GetString("name")
	opts := hcloud.SSHKeyUpdateOpts{
		Name: name,
	}
	if opts.Name == "" {
		return errors.New("no updates")
	}

	_, _, err = cli.Client().SSHKey.Update(cli.Context, sshKey, opts)
	if err != nil {
		return err
	}
	fmt.Printf("SSH key %d updated\n", sshKey.ID)
	return nil
}
