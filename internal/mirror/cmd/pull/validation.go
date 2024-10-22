/*
Copyright 2024 Flant JSC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pull

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io/fs"
	"os"
	"path/filepath"
)

func parseAndValidateParameters(_ *cobra.Command, args []string) error {
	var err error
	if err = validateImagesBundlePathArg(args); err != nil {
		return err
	}
	if err = validateChunkSizeFlag(); err != nil {
		return err
	}

	return nil
}

func validateImagesBundlePathArg(args []string) error {
	if len(args) != 1 {
		return errors.New("invalid number of arguments")
	}

	ImagesBundlePath = filepath.Clean(args[0])
	if filepath.Ext(ImagesBundlePath) != ".tar" {
		return errors.New("images-bundle-path argument should be a path to tar archive (.tar)")
	}

	stats, err := os.Stat(ImagesBundlePath)
	switch {
	case errors.Is(err, fs.ErrNotExist):
		// If only the file is not there it is fine, it will be created, but if directories on the path are also missing, this is bad.
		tarBundleDir := filepath.Dir(ImagesBundlePath)
		if _, err = os.Stat(tarBundleDir); err != nil {
			return err
		}
		break
	case err != nil && !errors.Is(err, fs.ErrNotExist):
		return err
	case stats.IsDir() || filepath.Ext(ImagesBundlePath) != ".tar":
		return fmt.Errorf("%s should be a tar archive", ImagesBundlePath)
	}
	return nil
}

func validateChunkSizeFlag() error {
	if ImagesBundleChunkSizeGB < 0 {
		return errors.New("Chunk size cannot be less than zero GB")
	}

	return nil
}
