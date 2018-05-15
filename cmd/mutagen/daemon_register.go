package main

import (
	"github.com/pkg/errors"

	"github.com/spf13/cobra"

	"github.com/havoc-io/mutagen/cmd"
	"github.com/havoc-io/mutagen/pkg/daemon"
)

func daemonRegisterMain(command *cobra.Command, arguments []string) {
	// Validate arguments.
	if len(arguments) != 0 {
		cmd.Fatal(errors.New("unexpected arguments provided"))
	}

	// Attempt registration.
	if err := daemon.Register(); err != nil {
		cmd.Fatal(err)
	}
}

var daemonRegisterCommand = &cobra.Command{
	Use:   "register",
	Short: "Registers Mutagen to start as a per-user daemon on login",
	Run:   daemonRegisterMain,
}

var daemonRegisterConfiguration struct {
	help bool
}

func init() {
	// Bind flags to configuration. We manually add help to override the default
	// message, but Cobra still implements it automatically.
	flags := daemonRegisterCommand.Flags()
	flags.BoolVarP(&daemonRegisterConfiguration.help, "help", "h", false, "Show help information")
}
