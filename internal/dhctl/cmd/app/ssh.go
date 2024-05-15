// Copyright 2021 Flant JSC
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
	"strings"

	"github.com/spf13/pflag"
)

const DefaultSSHAgentPrivateKeys = "~/.ssh/id_rsa"

var (
	SSHAgentPrivateKeys = make([]string, 0)
	SSHPrivateKeys      = make([]string, 0)
	SSHBastionHost      = ""
	SSHBastionPort      = ""
	SSHBastionUser      = os.Getenv("USER")
	SSHUser             = os.Getenv("USER")
	SSHHosts            = make([]string, 0)
	SSHPort             = ""
	SSHExtraArgs        = ""

	AskBecomePass = false
	BecomePass    = ""
)

func DefineSSHFlags(flagSet *pflag.FlagSet) {
	SSHAgentPrivateKeys = SetStringSliceVarFromEnv("SSH_AGENT_PRIVATE_KEYS", SSHAgentPrivateKeys)
	flagSet.StringSliceVar(
		&SSHAgentPrivateKeys,
		"ssh-agent-private-keys",
		SSHAgentPrivateKeys,
		"Paths to private keys. Those keys will be used to connect to servers and to the bastion. Can be specified multiple times (default: '~/.ssh/id_rsa')",
	)

	SSHHosts = SetStringSliceVarFromEnv("SSH_HOSTS", SSHHosts)
	flagSet.StringSliceVar(
		&SSHHosts,
		"ssh-host",
		SSHHosts,
		"SSH destination hosts, can be specified multiple times",
	)

	SSHBastionHost = SetStringVarFromEnv("SSH_BASTION_HOST", SSHBastionHost)
	flagSet.StringVar(
		&SSHBastionHost,
		"ssh-bastion-host",
		SSHBastionHost,
		"Jumper (bastion) host to connect to servers (will be used both by terraform and ansible). Only IPs or hostnames are supported, name from ssh-config will not work.",
	)

	SSHBastionPort = SetStringVarFromEnv("SSH_BASTION_PORT", SSHBastionPort)
	flagSet.StringVar(
		&SSHBastionPort,
		"ssh-bastion-port",
		SSHBastionPort,
		"SSH destination port.",
	)

	flagSet.StringVar(
		&SSHBastionUser,
		"ssh-bastion-user",
		SSHBastionUser,
		"User to authenticate under when connecting to bastion (default: $USER)",
	)

	flagSet.StringVar(
		&SSHUser,
		"ssh-user",
		SSHUser,
		"User to authenticate under (default: $USER)",
	)

	SSHPort = SetStringVarFromEnv("SSH_PORT", SSHPort)
	flagSet.StringVar(
		&SSHPort,
		"ssh-port",
		SSHPort,
		"SSH destination port.",
	)

	SSHExtraArgs = SetStringVarFromEnv("SSH_EXTRA_ARGS", SSHExtraArgs)
	flagSet.StringVar(
		&SSHExtraArgs,
		"ssh-extra-args",
		SSHExtraArgs,
		"extra args for ssh commands (-vvv)",
	)
}

func CheckSSHParameters() error {
	if len(SSHAgentPrivateKeys) == 0 {
		SSHAgentPrivateKeys = append(SSHAgentPrivateKeys, DefaultSSHAgentPrivateKeys)
	}
	var err error
	SSHPrivateKeys, err = ParseSSHPrivateKeyPaths(SSHAgentPrivateKeys)
	if err != nil {
		return fmt.Errorf("ssh private keys: %v", err)
	}
	return nil
}

func ParseSSHPrivateKeyPaths(pathSets []string) ([]string, error) {
	res := make([]string, 0)
	if len(pathSets) == 0 || (len(pathSets) == 1 && pathSets[0] == "") {
		return res, nil
	}

	for _, pathSet := range pathSets {
		keys := strings.Split(pathSet, ",")
		for _, k := range keys {
			if strings.HasPrefix(k, "~") {
				home := os.Getenv("HOME")
				if home == "" {
					return nil, fmt.Errorf("HOME is not defined for key '%s'", k)
				}
				k = strings.Replace(k, "~", home, 1)
			}

			keyPath, err := filepath.Abs(k)
			if err != nil {
				return nil, fmt.Errorf("get absolute path for '%s': %v", k, err)
			}
			res = append(res, keyPath)
		}
	}
	return res, nil
}

func DefineBecomeFlags(flagSet *pflag.FlagSet) {
	// Ansible compatible
	AskBecomePass = SetBoolVarFromEnv("ASK_BECOME_PASS", AskBecomePass)
	flagSet.BoolVarP(
		&AskBecomePass,
		"ask-become-pass",
		"K",
		false,
		"Ask for sudo password before the installation process.",
	)
}
