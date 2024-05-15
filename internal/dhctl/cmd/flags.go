// Copyright 2024 Flant JSC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dhctl

import (
	"os"
	"path/filepath"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/spf13/pflag"

	"github.com/deckhouse/deckhouse-cli/internal/dhctl/cmd/app"
	"github.com/deckhouse/deckhouse-cli/internal/dhctl/image"
	"github.com/deckhouse/deckhouse-cli/internal/mirror/contexts"
)

const (
	deckhouseRegistryHost     = "registry.deckhouse.io"
	enterpriseEditionRepoPath = "/deckhouse/ee"
	enterpriseEditionRepo     = deckhouseRegistryHost + enterpriseEditionRepoPath
)

var (
	TempDir       = filepath.Join(os.TempDir(), "dhctl")
	Insecure      bool
	TLSSkipVerify bool

	RegistryRepo     = enterpriseEditionRepo // Fallback to EE if nothing was given as source.
	RegistryLogin    string
	RegistryPassword string
	LicenseToken     string
	ImageTag         string
)

func addRegistryFlags(flagSet *pflag.FlagSet) {
	RegistryRepo = os.Getenv(app.ConfigEnvName("REGISTRY"))
	flagSet.StringVar(
		&RegistryRepo,
		"registry",
		enterpriseEditionRepo,
		"Pull dhctl from registry address.",
	)

	RegistryLogin = os.Getenv(app.ConfigEnvName("REGISTRY_LOGIN"))
	flagSet.StringVar(
		&RegistryLogin,
		"login",
		"",
		"Registry login.",
	)

	RegistryPassword = os.Getenv(app.ConfigEnvName("REGISTRY_PASSWORD"))
	flagSet.StringVar(
		&RegistryPassword,
		"password",
		"",
		"Registry password.",
	)

	LicenseToken = os.Getenv(app.ConfigEnvName("LICENSE_TOKEN"))
	flagSet.StringVar(
		&LicenseToken,
		"license",
		"",
		"Pull dhctl to local machine using license key. Shortcut for --login=license-token --password=<>.",
	)

	ImageTag = os.Getenv(app.ConfigEnvName("IMAGE_TAG"))
	flagSet.StringVar(
		&ImageTag,
		"tag",
		"stable",
		"Pull dhctl with specified tag.",
	)

	TLSSkipVerify = app.SetBoolVarFromEnv("TLS-SKIP-VERIFY", TLSSkipVerify)
	flagSet.BoolVar(
		&TLSSkipVerify,
		"tls-skip-verify",
		false,
		"Disable TLS certificate validation.",
	)

	Insecure = app.SetBoolVarFromEnv("INSECURE", Insecure)
	flagSet.BoolVar(
		&Insecure,
		"insecure",
		false,
		"Interact with registries over HTTP.",
	)
}

func getSourceRegistryAuthProvider() authn.Authenticator {
	if RegistryLogin != "" {
		return authn.FromConfig(authn.AuthConfig{
			Username: RegistryLogin,
			Password: RegistryPassword,
		})
	}

	if LicenseToken != "" {
		return authn.FromConfig(authn.AuthConfig{
			Username: "license-token",
			Password: LicenseToken,
		})
	}

	return authn.Anonymous
}

func buildInstallerContext() *image.InstallerContext {
	ctx := &image.InstallerContext{
		BaseContext: contexts.BaseContext{
			Insecure:              Insecure,
			SkipTLSVerification:   TLSSkipVerify,
			DeckhouseRegistryRepo: RegistryRepo,
			RegistryAuth:          getSourceRegistryAuthProvider(),
			UnpackedImagesPath:    TempDir,
		},
		Args:     os.Args[1:],
		ImageTag: ImageTag,
	}
	return ctx
}
