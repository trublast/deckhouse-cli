package dhctl

import (
	"github.com/spf13/cobra"
	"k8s.io/component-base/logs"
	"k8s.io/kubectl/pkg/util/templates"
)

const (
	AppName                   = "dhctl"
	deckhouseRegistryHost     = "registry.deckhouse.io"
	enterpriseEditionRepoPath = "/deckhouse/ee"
	enterpriseEditionRepo     = deckhouseRegistryHost + enterpriseEditionRepoPath
	licenceNote               = `

LICENSE NOTE: The d8 dhctl functionality is exclusively available to users holding a
valid license for any commercial version of the Deckhouse Kubernetes Platform.

© Flant JSC 2024`
)

var (
	Insecure      bool
	TLSSkipVerify bool

	RegistryRepo     = enterpriseEditionRepo // Fallback to EE if nothing was given as source.
	RegistryLogin    string
	RegistryPassword string
	LicenseToken     string
	ImageTag         string
)

func NewCommand() *cobra.Command {
	dhctlCmd := &cobra.Command{
		Use:           "i",
		Aliases:       []string{"dhctl"},
		Short:         "Run dhctl tool",
		Long:          templates.LongDesc(`Run dhctl tool.` + licenceNote),
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	logs.AddFlags(dhctlCmd.PersistentFlags())

	addRegistryFlags(dhctlCmd.Flags())

	dhctlCmd.AddCommand(
		DefineBootstrapCommand(),
	)

	return dhctlCmd
}

func parseAndValidateParameters(cmd *cobra.Command, args []string) error {
	return nil
}

func dhctl(cmd *cobra.Command, args []string) error {
	return nil
}

func cleanup(_ *cobra.Command, _ []string) error {
	return nil
}
