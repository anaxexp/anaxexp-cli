package ci

import (
	"github.com/anaxexp/anaxexp-cli/cmd/anaxexp/ci/build"
	"github.com/anaxexp/anaxexp-cli/cmd/anaxexp/ci/deploy"
	"github.com/anaxexp/anaxexp-cli/cmd/anaxexp/ci/initialize"
	"github.com/anaxexp/anaxexp-cli/cmd/anaxexp/ci/release"
	"github.com/anaxexp/anaxexp-cli/cmd/anaxexp/ci/run"
	"github.com/spf13/cobra"
)

// Cmd represents the ci command.
var Cmd = &cobra.Command{
	Use:   "ci [command]",
	Short: "ci commands",
}

func init() {
	Cmd.AddCommand(initialize.Cmd)
	Cmd.AddCommand(build.Cmd)
	Cmd.AddCommand(release.Cmd)
	Cmd.AddCommand(deploy.Cmd)
	Cmd.AddCommand(run.Cmd)
}
