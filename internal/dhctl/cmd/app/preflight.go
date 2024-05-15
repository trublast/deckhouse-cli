// Copyright 2023 Flant JSC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"github.com/spf13/pflag"
)

var (
	PreflightSkipAll                   = false
	PreflightSkipSSHForword            = false
	PreflightSkipAvailabilityPorts     = false
	PreflightSkipResolvingLocalhost    = false
	PreflightSkipDeckhouseVersionCheck = false
	PreflightSkipRegistryThroughProxy  = false
)

const (
	SSHForwardArgName                = "preflight-skip-ssh-forward-check"
	PortsAvailabilityArgName         = "preflight-skip-availability-ports-check"
	ResolvingLocalhostArgName        = "preflight-skip-resolving-localhost-check"
	DeckhouseVersionCheckArgName     = "preflight-skip-deckhouse-version-check"
	RegistryThroughProxyCheckArgName = "preflight-skip-registry-through-proxy"
)

func DefinePreflight(flagSet *pflag.FlagSet) {
	PreflightSkipAll = SetBoolVarFromEnv("PREFLIGHT_SKIP_ALL_CHECKS", PreflightSkipAll)
	flagSet.BoolVar(
		&PreflightSkipAll,
		"preflight-skip-all-checks",
		PreflightSkipAll,
		"Skip all preflight checks.",
	)

	PreflightSkipSSHForword = SetBoolVarFromEnv("PREFLIGHT_SKIP_SSH_FORWARD_CHECK", PreflightSkipSSHForword)
	flagSet.BoolVar(
		&PreflightSkipSSHForword,
		SSHForwardArgName,
		PreflightSkipSSHForword,
		"Skip SSH forward preflight check.",
	)

	PreflightSkipAvailabilityPorts = SetBoolVarFromEnv("PREFLIGHT_SKIP_AVAILABILITY_PORTS_CHECK", PreflightSkipAvailabilityPorts)
	flagSet.BoolVar(
		&PreflightSkipAvailabilityPorts,
		PortsAvailabilityArgName,
		PreflightSkipAvailabilityPorts,
		"Skip availability ports preflight check.",
	)

	PreflightSkipResolvingLocalhost = SetBoolVarFromEnv("PREFLIGHT_SKIP_RESOLVING_LOCALHOST_CHECK", PreflightSkipResolvingLocalhost)
	flagSet.BoolVar(
		&PreflightSkipResolvingLocalhost,
		ResolvingLocalhostArgName,
		PreflightSkipResolvingLocalhost,
		"Skip resolving the localhost domain.",
	)

	PreflightSkipDeckhouseVersionCheck = SetBoolVarFromEnv("PREFLIGHT_SKIP_INCOMPATIBLE_VERSION_CHECK", PreflightSkipDeckhouseVersionCheck)
	flagSet.BoolVar(
		&PreflightSkipDeckhouseVersionCheck,
		DeckhouseVersionCheckArgName,
		PreflightSkipDeckhouseVersionCheck,
		"Skip verifying deckhouse version.",
	)

	PreflightSkipRegistryThroughProxy = SetBoolVarFromEnv("PREFLIGHT_SKIP_REGISTRY_THROUGH_PROXY", PreflightSkipRegistryThroughProxy)
	flagSet.BoolVar(
		&PreflightSkipRegistryThroughProxy,
		RegistryThroughProxyCheckArgName,
		PreflightSkipRegistryThroughProxy,
		"Skip checking registry through proxy.",
	)
}
