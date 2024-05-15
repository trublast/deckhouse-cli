package dhctl

import (
	"github.com/spf13/cobra"
	"k8s.io/component-base/logs"
)

const (
	licenceNote = `LICENSE NOTE: The d8 dhctl functionality is exclusively available to users holding a
valid license for any commercial version of the Deckhouse Kubernetes Platform.

© Flant JSC 2024`
)

func NewCommand() *cobra.Command {
	dhctlCmd := &cobra.Command{
		Use:   "i",
		Short: "run Deckhouse installer tool",
		Long:  "Run Deckhouse installer tool.\n" + licenceNote,
	}

	dhctlCmd.AddCommand(
		DefineBootstrapCommand(),
	)

	logs.AddFlags(dhctlCmd.PersistentFlags())

	return dhctlCmd
}
