package dhctl

import (
	"os"

	"github.com/spf13/cobra"
	"k8s.io/component-base/logs"

	"github.com/deckhouse/deckhouse-cli/internal/dhctl/cmd/app"
	"github.com/deckhouse/deckhouse-cli/internal/dhctl/image"
)

func DefineBootstrapCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "bootstrap",
		Short:         "Bootstrap cluster.",
		Long:          "Bootstrap cluster.\n" + licenceNote,
		SilenceErrors: true,
		SilenceUsage:  true,
		PreRunE:       parseAndValidateParameters,
		RunE:          bootstrap,
		PostRunE:      cleanup,
	}
	/*
	 */
	addRegistryFlags(cmd.Flags())

	app.DefineCacheFlags(cmd.Flags())
	app.DefineDropCacheFlags(cmd.Flags())
	app.GlobalFlags(cmd.Flags())
	app.DefineConfigFlags(cmd.Flags())
	app.DefineSSHFlags(cmd.Flags())
	app.DefineBecomeFlags(cmd.Flags())
	app.DefineResourcesFlags(cmd.Flags())
	app.DefinePreflight(cmd.Flags())
	app.DefineDeckhouseFlags(cmd.Flags())
	app.DefineDontUsePublicImagesFlags(cmd.Flags())
	app.DefinePostBootstrapScriptFlags(cmd.Flags())

	logs.AddFlags(cmd.Flags())
	return cmd
}

func parseAndValidateParameters(cmd *cobra.Command, args []string) error {
	err := app.CheckConfigParameters()
	if err != nil {
		return err
	}
	return app.CheckSSHParameters()
}

func bootstrap(cmd *cobra.Command, args []string) error {
	ctx := buildInstallerContext()
	return image.PullInstallerImage(ctx)

}

func cleanup(cmd *cobra.Command, args []string) error {
	return os.RemoveAll(TempDir)
}
