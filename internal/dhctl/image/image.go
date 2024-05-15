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

package image

import (
	"fmt"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"

	"github.com/deckhouse/deckhouse-cli/internal/mirror/contexts"
	"github.com/deckhouse/deckhouse-cli/internal/mirror/util/auth"
	"github.com/deckhouse/deckhouse-cli/internal/mirror/util/errorutil"
	"github.com/deckhouse/deckhouse-cli/internal/mirror/util/log"
)

// InstallerContext holds data related to pending mirroring-from-registry operation.
type InstallerContext struct {
	contexts.BaseContext
	Args     []string
	ImageTag string
}

func PullInstallerImage(ctx *InstallerContext) error {
	nameOpts, remoteOpts := auth.MakeRemoteRegistryRequestOptions(ctx.RegistryAuth, ctx.Insecure, ctx.SkipTLSVerification)
	imageRef := fmt.Sprintf("%s/%s:%s", ctx.DeckhouseRegistryRepo, "install", ctx.ImageTag)
	ref, err := name.ParseReference(imageRef, nameOpts...)
	if err != nil {
		return err
	}

	log.InfoF("Pulling %s...", imageRef)
	_, err = remote.Image(ref, remoteOpts...)
	if err != nil {
		if errorutil.IsImageNotFoundError(err) {
			return fmt.Errorf("⚠️ %s Not found in registry", imageRef)
		}

		return fmt.Errorf("pull image %s metadata: %w", imageRef, err)
	}
	return nil
}
