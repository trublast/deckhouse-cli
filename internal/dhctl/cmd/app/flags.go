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
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"time"

	"github.com/spf13/pflag"
)

var TmpDirName = filepath.Join(os.TempDir(), "dhctl")

var (
	ConfigPaths = make([]string, 0)
	SanityCheck = false
	LoggerType  = "pretty"
)

func GlobalFlags(flagSet *pflag.FlagSet) {
	LoggerType = SetStringVarFromEnv("LOGGER_TYPE", LoggerType)
	flagSet.StringVar(
		&LoggerType,
		"logger-type",
		LoggerType,
		"Format logs output of a dhctl in different ways.",
	)

	TmpDirName = SetStringVarFromEnv("TMP_DIR", TmpDirName)
	flagSet.StringVar(
		&TmpDirName,
		"tmp-dir",
		TmpDirName,
		"Set temporary directory for debug purposes.",
	)
}

func DefineConfigFlags(flagSet *pflag.FlagSet) {
	ConfigPaths = SetStringSliceVarFromEnv("CONFIG", ConfigPaths)
	flagSet.StringSliceVar(
		&ConfigPaths,
		"config",
		ConfigPaths,
		`Path to a file with bootstrap configuration and declared Kubernetes resources in YAML format.
It can be go-template file (for only string keys!). Passed data contains next keys:
cloudDiscovery - the data discovered by applying Terraform and getting its output. It depends on the cloud provider.
`,
	)
}

func DefineSanityFlags(flagSet *pflag.FlagSet) {
	flagSet.BoolVar(
		&SanityCheck,
		"yes-i-am-sane-and-i-understand-what-i-am-doing",
		false,
		"You should double check what you are doing here.",
	)
}

func ConfigEnvName(name string) string {
	return "DHCTL_CLI_" + name
}

func CheckConfigParameters() error {
	var availableLoggerTypes = []string{"pretty", "simple", "json"}
	if slices.Contains(availableLoggerTypes, LoggerType) {
		return nil
	}
	return fmt.Errorf("invalid logger type: %s", LoggerType)
}

func SetStringVarFromEnv(envName, defaultValue string) string {
	if v := os.Getenv(ConfigEnvName(envName)); v != "" {
		return v
	}
	return defaultValue
}

func SetBoolVarFromEnv(envName string, defaultValue bool) bool {
	if v := os.Getenv(ConfigEnvName(envName)); v == "true" {
		return true
	}
	return defaultValue
}

func SetDurationVarFromEnv(envName string, defaultValue time.Duration) time.Duration {
	dt := os.Getenv(ConfigEnvName(envName))
	if dt == "" {
		return defaultValue
	}
	dur, err := time.ParseDuration(dt)
	if err != nil {
		panic(err)
	}
	return dur
}

var (
	envVarValuesSeparator = "\r?\n"
	envVarValuesTrimmer   = regexp.MustCompile(envVarValuesSeparator + "$")
	envVarValuesSplitter  = regexp.MustCompile(envVarValuesSeparator)
)

func SetStringSliceVarFromEnv(envName string, defaultValue []string) []string {
	envarValue := os.Getenv(envName)
	if envarValue == "" {
		return defaultValue
	}
	// Split by new line to extract multiple values, if any.
	trimmed := envVarValuesTrimmer.ReplaceAllString(envarValue, "")
	return envVarValuesSplitter.Split(trimmed, -1)
}
