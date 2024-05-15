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
	"time"

	"github.com/spf13/pflag"
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
	ResourcesPath = SetStringVarFromEnv("RESOURCES", ResourcesPath)
	flagSet.StringVar(
		&ResourcesPath,
		"resources",
		ResourcesPath,
		"Path to a file with declared Kubernetes resources in YAML format.\nDeprecated. Please use --config flag multiple repeatedly for logical resources separation.",
	)

	ResourcesTimeout = SetStringVarFromEnv("RESOURCES_TIMEOUT", ResourcesTimeout)
	flagSet.StringVar(
		&ResourcesTimeout,
		"resources-timeout",
		ResourcesTimeout,
		"Timeout to create resources. Experimental. This feature may be deleted in the future.",
	)
}

func DefineDeckhouseFlags(flagSet *pflag.FlagSet) {
	DeckhouseTimeout = SetDurationVarFromEnv("DECKHOUSE_TIMEOUT", DeckhouseTimeout)
	flagSet.DurationVar(
		&DeckhouseTimeout,
		"deckhouse-timeout",
		DeckhouseTimeout,
		"Timeout to install deckhouse. Experimental. This feature may be deleted in the future.",
	)
}

func DefineDontUsePublicImagesFlags(flagSet *pflag.FlagSet) {
	DontUsePublicControlPlaneImages = SetBoolVarFromEnv("DONT_USE_PUBLIC_CONTROL_PLANE_IMAGES", DontUsePublicControlPlaneImages)
	flagSet.BoolVar(
		&DontUsePublicControlPlaneImages,
		"dont-use-public-control-plane-images",
		DontUsePublicControlPlaneImages,
		"DEPRECATED. Don't use public images for control-plane components.",
	)
}

func DefinePostBootstrapScriptFlags(flagSet *pflag.FlagSet) {
	PostBootstrapScriptPath = SetStringVarFromEnv("POST_BOOTSTRAP_SCRIPT_PATH", PostBootstrapScriptPath)
	flagSet.StringVar(
		&PostBootstrapScriptPath,
		"post-bootstrap-script-path",
		PostBootstrapScriptPath,
		`Path to bash (or another interpreted language which installed on master node) script which will execute after bootstrap resources.
All output of the script will be logged with Info level with prefix 'Post-bootstrap script result:'.
If you want save to state cache on key 'post-bootstrap-result' you need to out result with prefix 'Result of post-bootstrap script:' in one line.
Experimental. This feature may be deleted in the future.`,
	)

	PostBootstrapScriptTimeout = SetDurationVarFromEnv("POST_BOOTSTRAP_SCRIPT_TIMEOUT", PostBootstrapScriptTimeout)
	flagSet.DurationVar(
		&PostBootstrapScriptTimeout,
		"post-bootstrap-script-timeout",
		PostBootstrapScriptTimeout,
		"Timeout to execute after bootstrap resources script. Experimental. This feature may be deleted in the future.",
	)
}
