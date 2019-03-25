package commands

import (
	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
)

func NewRootCmd(dockerCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Short: "Build with BuildKit",
		Use:   "buildx",
	}
	addCommands(cmd, dockerCli)
	return cmd
}

func addCommands(cmd *cobra.Command, dockerCli command.Cli) {
	cmd.AddCommand(
		buildCmd(dockerCli),
	)
}
