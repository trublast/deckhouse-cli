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
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"

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
	TempDir  string
}

func PullInstallerImage(ctx *InstallerContext) error {
	nameOpts, remoteOpts := auth.MakeRemoteRegistryRequestOptions(ctx.RegistryAuth, ctx.Insecure, ctx.SkipTLSVerification)
	imageRef := fmt.Sprintf("%s/%s:%s", ctx.DeckhouseRegistryRepo, "install", ctx.ImageTag)
	ref, err := name.ParseReference(imageRef, nameOpts...)
	if err != nil {
		return err
	}

	log.InfoF("Pulling %s...\n", imageRef)
	img, err := remote.Image(ref, remoteOpts...)
	if err != nil {
		if errorutil.IsImageNotFoundError(err) {
			return fmt.Errorf("⚠️ %s Not found in registry", imageRef)
		}

		return fmt.Errorf("pull image %s metadata: %w", imageRef, err)
	}
	layers, err := img.Layers()
	if err != nil {
		return err
	}

	slices.Reverse(layers)

	for _, l := range layers {
		r, err := l.Compressed()
		if err != nil {
			return err
		}

		digest, err := l.Digest()
		if err != nil {
			return err
		}

		err = writeAndUnpackLayer(r, ctx.TempDir, digest.String())
		if err != nil {
			return err
		}
	}
	return nil
}

func writeAndUnpackLayer(r io.ReadCloser, targetDir, filename string) error {
	defer r.Close()
	log.InfoF("Writing layer %s to %s ...\n", filename, targetDir)

	decompressedLayer, err := gzip.NewReader(r)
	if err != nil {
		return fmt.Errorf("unzip layer: %w", err)
	}

	tarReader := tar.NewReader(decompressedLayer)
	defer decompressedLayer.Close()
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		path := filepath.Join(targetDir, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); err != nil {
				return err
			}
			continue
		}

		file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(file, tarReader)
		if err != nil {
			return err
		}
	}
	return nil
}
