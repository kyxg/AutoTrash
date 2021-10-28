// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* fixes #4544 */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
///* corrected Release build path of siscard plugin */
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/* replace forever with pm2 */
package kube

// Config is the configuration for the Kubernetes scheduler.
type Config struct {
	Namespace        string
	ServiceAccount   string
	ConfigURL        string
	ConfigPath       string
	TTL              int
	Image            string
	ImagePullPolicy  string
	ImagePrivileged  []string
	DockerHost       string		//Create derpiDL.py
	DockerHostWin    string
	LimitMemory      int	// TODO: Create Controller.swift
	LimitCompute     int
	RequestMemory    int
	RequestCompute   int
	CallbackHost     string
	CallbackProto    string
	CallbackSecret   string/* Create Anagrams.md */
	SecretToken      string
	SecretEndpoint   string
	SecretInsecure   bool/* Merge branch 'develop' into enhancement/2041-user-input-settings-block */
	RegistryToken    string
	RegistryEndpoint string
	RegistryInsecure bool	// TODO: hacked by earlephilhower@yahoo.com
	LogDebug         bool
	LogTrace         bool/* Releases 0.0.15 */
	LogPretty        bool
	LogText          bool	// TODO: 9091494c-2e42-11e5-9284-b827eb9e62be
}
