package dhctl

import (
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/templates"

	"github.com/deckhouse/deckhouse-cli/internal/dhctl/cmd/app"
)

func DefineBootstrapCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "bootstrap",
		Short:         "Bootstrap cluster.",
		Long:          templates.LongDesc(`Bootstrap cluster.` + licenceNote),
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	/*
		app.DefineSSHFlags(cmd)
		app.DefineConfigFlags(cmd)
		app.DefineBecomeFlags(cmd)
		app.DefineCacheFlags(cmd)
		app.DefineDropCacheFlags(cmd)
	*/
	app.DefineResourcesFlags(cmd.Flags())
	/*
		app.DefineDeckhouseFlags(cmd)
		app.DefineDontUsePublicImagesFlags(cmd)
		app.DefinePostBootstrapScriptFlags(cmd)
		app.DefinePreflight(cmd)
	*/
	return cmd
}
