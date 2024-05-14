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

package app

import (
	"os"
	"time"

	"github.com/spf13/pflag"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	ResourcesPath                   = ""
	ResourcesTimeout                = "15m"
	DontUsePublicControlPlaneImages = false
	DeckhouseTimeout                = 15 * time.Minute
	PostBootstrapScriptTimeout      = 10 * time.Minute
	PostBootstrapScriptPath         = ""
)

func DefineResourcesFlags(flagSet *pflag.FlagSet) {
	ResourcesPath = os.Getenv(ConfigEnvName("RESOURCES"))
	flagSet.StringVar(
		&ResourcesPath,
		"resources",
		"",
		"Path to a file with declared Kubernetes resources in YAML format.\nDeprecated. Please use --config flag multiple repeatedly for logical resources separation.",
	)

	ResourcesTimeout = os.Getenv(ConfigEnvName("RESOURCES_TIMEOUT"))
	flagSet.StringVar(
		&ResourcesTimeout,
		"resources-timeout",
		"",
		"Timeout to create resources. Experimental. This feature may be deleted in the future.",
	)
}

func DefineDeckhouseFlags(cmd *kingpin.CmdClause) {
	cmd.Flag("deckhouse-timeout", "Timeout to install deckhouse. Experimental. This feature may be deleted in the future.").
		Envar(ConfigEnvName("DECKHOUSE_TIMEOUT")).
		Default(DeckhouseTimeout.String()).
		DurationVar(&DeckhouseTimeout)
}

func DefineDontUsePublicImagesFlags(cmd *kingpin.CmdClause) {
	const help = `DEPRECATED. Don't use public images for control-plane components.`
	cmd.Flag("dont-use-public-control-plane-images", help).
		Envar(ConfigEnvName("DONT_USE_PUBLIC_CONTROL_PLANE_IMAGES")).
		Default("false").
		BoolVar(&DontUsePublicControlPlaneImages)
}

func DefinePostBootstrapScriptFlags(cmd *kingpin.CmdClause) {
	cmd.Flag("post-bootstrap-script-path", `Path to bash (or another interpreted language which installed on master node) script which will execute after bootstrap resources.
All output of the script will be logged with Info level with prefix 'Post-bootstrap script result:'.
If you want save to state cache on key 'post-bootstrap-result' you need to out result with prefix 'Result of post-bootstrap script:' in one line.
Experimental. This feature may be deleted in the future.`).
		Envar(ConfigEnvName("POST_BOOTSTRAP_SCRIPT_PATH")).
		StringVar(&PostBootstrapScriptPath)

	cmd.Flag("post-bootstrap-script-timeout", "Timeout to execute after bootstrap resources script. Experimental. This feature may be deleted in the future.").
		Envar(ConfigEnvName("POST_BOOTSTRAP_SCRIPT_TIMEOUT")).
		Default(PostBootstrapScriptTimeout.String()).
		DurationVar(&PostBootstrapScriptTimeout)
}
