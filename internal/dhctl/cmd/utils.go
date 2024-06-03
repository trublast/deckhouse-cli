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
	"strings"

	"github.com/google/go-containerregistry/pkg/authn"

	"github.com/deckhouse/deckhouse-cli/internal/mirror/contexts"
)

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

func buildContext() *contexts.BaseContext {
	ctx := &contexts.BaseContext{
		Insecure:              Insecure,
		SkipTLSVerification:   TLSSkipVerify,
		DeckhouseRegistryRepo: RegistryRepo,
		RegistryAuth:          getSourceRegistryAuthProvider(),
		UnpackedImagesPath:    TempDir,
	}
	return ctx
}

func getDhctlEnvs() map[string]string {
	ret := make(map[string]string)
	for _, s := range os.Environ() {
		pair := strings.SplitN(s, "=", 2)
		if strings.HasPrefix(pair[0], "DHCTL_") {
			ret[pair[0]] = pair[1]
		}
	}
	return ret
}
