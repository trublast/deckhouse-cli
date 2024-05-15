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

const (
	UseStateCacheAsk = "ask"
	UseStateCacheYes = "yes"
	UseStateCacheNo  = "no"
)

var (
	CacheDir   = filepath.Join(os.TempDir(), "dhctl")
	UseTfCache = "ask"

	DropCache = false

	CacheKubeConfig          = ""
	CacheKubeConfigContext   = ""
	CacheKubeConfigInCluster = false
	CacheKubeNamespace       = ""
	CacheKubeName            = ""
	CacheKubeLabels          = make(map[string]string)
)

func DefineCacheFlags(flagSet *pflag.FlagSet) {
	CacheDir = SetStringVarFromEnv("CACHE_DIR", CacheDir)
	flagSet.StringVar(
		&CacheDir,
		"cache-dir",
		CacheDir,
		"Directory to store the cache.",
	)

	UseTfCache = SetStringVarFromEnv("USE_CACHE", UseTfCache)
	flagSet.StringVar(
		&UseTfCache,
		"use-cache",
		UseStateCacheAsk,
		fmt.Sprintf(`Behaviour for using terraform state cache. May be:
	%s - ask user about it (Default)
   	%s - use cache
	%s  - don't use cache
	`, UseStateCacheAsk, UseStateCacheYes, UseStateCacheNo),
	)

	CacheKubeConfig = SetStringVarFromEnv("CACHE_STORE_KUBE_CONFIG", CacheKubeConfig)
	flagSet.StringVar(
		&CacheKubeConfig,
		"kube-cache-store-kubeconfig",
		CacheKubeConfig,
		"Path to kubernetes config file for storing cache in kubernetes secret.",
	)

	CacheKubeConfigContext = SetStringVarFromEnv("CACHE_STORE_KUBE_CONFIG_CONTEXT", CacheKubeConfigContext)
	flagSet.StringVar(
		&CacheKubeConfigContext,
		"kube-cachestore-kubeconfig-context",
		CacheKubeConfigContext,
		"Context from kubernetes config to connect to Kubernetes API. for storing cache in kubernetes secret.",
	)

	CacheKubeConfigInCluster = SetBoolVarFromEnv("CACHE_STORE_KUBE_CLIENT_FROM_CLUSTER", CacheKubeConfigInCluster)
	flagSet.BoolVar(
		&CacheKubeConfigInCluster,
		"kube-cachestore-kube-client-from-cluster",
		CacheKubeConfigInCluster,
		"Use in-cluster Kubernetes API access. for storing cache in kubernetes secret.",
	)

	CacheKubeNamespace = SetStringVarFromEnv("CACHE_STORE_KUBE_NAMESPACE", CacheKubeNamespace)
	flagSet.StringVar(
		&CacheKubeNamespace,
		"kube-cachestore-namespace",
		CacheKubeNamespace,
		"Use in-cluster Kubernetes API access. for storing cache in kubernetes secret.",
	)

	for _, s := range SetStringSliceVarFromEnv("CACHE_STORE_KUBE_LABELS", []string{}) {
		item := strings.Split(s, ":")
		CacheKubeLabels[item[0]] = item[1]
	}
	flagSet.StringToStringVar(
		&CacheKubeLabels,
		"kube-cachestore-labels",
		CacheKubeLabels,
		"List labels for cache secrets.",
	)

	CacheKubeName = SetStringVarFromEnv("CACHE_STORE_KUBE_NAME", CacheKubeName)
	flagSet.StringVar(
		&CacheKubeName,
		"kube-cachestore-name",
		CacheKubeName,
		"Name for cache secret.",
	)
}

func DefineDropCacheFlags(flagSet *pflag.FlagSet) {
	flagSet.BoolVar(
		&DropCache,
		"yes-i-want-to-drop-cache",
		DropCache,
		"All cached information will be deleted from your local cache.",
	)
}
