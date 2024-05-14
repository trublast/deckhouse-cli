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
	"path/filepath"

	"gopkg.in/alecthomas/kingpin.v2"
)

var TmpDirName = filepath.Join(os.TempDir(), "dhctl")

var (
	ConfigPaths = make([]string, 0)
	SanityCheck = false
	LoggerType  = "pretty"
)

func GlobalFlags(cmd *kingpin.Application) {
	cmd.Flag("logger-type", "Format logs output of a dhctl in different ways.").
		Envar(ConfigEnvName("LOGGER_TYPE")).
		Default("pretty").
		EnumVar(&LoggerType, "pretty", "simple", "json")
	cmd.Flag("tmp-dir", "Set temporary directory for debug purposes.").
		Envar(ConfigEnvName("TMP_DIR")).
		Default(TmpDirName).
		StringVar(&TmpDirName)
}

func DefineConfigFlags(cmd *kingpin.CmdClause) {
	cmd.Flag("config", `Path to a file with bootstrap configuration and declared Kubernetes resources in YAML format.
It can be go-template file (for only string keys!). Passed data contains next keys:
  cloudDiscovery - the data discovered by applying Terraform and getting its output. It depends on the cloud provider.
`).
		Required().
		Envar(ConfigEnvName("CONFIG")).
		StringsVar(&ConfigPaths)
}

func DefineSanityFlags(cmd *kingpin.CmdClause) {
	cmd.Flag("yes-i-am-sane-and-i-understand-what-i-am-doing", "You should double check what you are doing here.").
		Default("false").
		BoolVar(&SanityCheck)
}

func ConfigEnvName(name string) string {
	return "DHCTL_CLI_" + name
}
