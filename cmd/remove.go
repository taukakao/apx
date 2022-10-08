package cmd

/*	License: GPLv3
	Authors:
		Mirko Brombin <send@mirko.pm>
		Pietro di Caprio <pietro@fabricators.ltd>
	Copyright: 2022
	Description: Apx is a wrapper around apt to make it works inside a container
	from outside, directly on the host.
*/

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vanilla-os/apx/core"
)

func removeUsage(*cobra.Command) error {
	fmt.Print(`Description: 
Remove packages.

Usage:
  apx remove <packages>

Examples:
  apx remove htop
`)
	return nil
}

func NewRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove packages.",
		RunE:  remove,
	}
	cmd.SetUsageFunc(removeUsage)
	return cmd
}

func remove(cmd *cobra.Command, args []string) error {
	sys := cmd.Flag("sys").Value.String() == "true"
	command := append([]string{}, core.GetPkgManager(sys)...)
	command = append(command, "remove")
	command = append(command, args...)

	if sys {
		core.AlmostRun(command...)
		return nil
	}

	core.RunContainer(command...)
	return nil
}