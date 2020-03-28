/*
Copyright 2019 The Jetstack cert-manager contributors.

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

package meta

// ConditionStatus represents a condition's status.
type ConditionStatus string

// These are valid condition statuses. "ConditionTrue" means a resource is in
// the condition; "ConditionFalse" means a resource is not in the condition;
// "ConditionUnknown" means kubernetes can't decide if a resource is in the
// condition or not. In the future, we could add other intermediate
// conditions, e.g. ConditionDegraded.
const (
	// ConditionTrue represents the fact that a given condition is true
	ConditionTrue ConditionStatus = "True"

	// ConditionFalse represents the fact that a given condition is false
	ConditionFalse ConditionStatus = "False"

	// ConditionUnknown represents the fact that a given condition is unknown
	ConditionUnknown ConditionStatus = "Unknown"
)

type LocalObjectReference struct {
	// Name of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// TODO: Add other useful fields. apiVersion, kind, uid?
	Name string
}

// ObjectReference is a reference to an object with a given name, kind and group.
type ObjectReference struct {
	Name  string
	Kind  string
	Group string
}

type SecretKeySelector struct {
	// The name of the secret in the pod's namespace to select from.
	LocalObjectReference
	// The key of the secret to select from. Must be a valid secret key.
	Key string
}

const (
	TLSCAKey = "ca.crt"
)

// HTTPSelfCheckStrategy describes how the self check should be performed for HTTP01 solvers
// +kubebuilder:validation:Enum=Standard;Disabled
type HTTPSelfCheckStrategy string

const (
	// HTTPSelfCheckStrategyStandard means that a regular self check will be performed by querying the challenge located at the DNS name.
	HTTPSelfCheckStrategyStandard HTTPSelfCheckStrategy = "Standard"

	// HTTPSelfCheckStrategyDisabled means that no self check will be performed.
	HTTPSelfCheckStrategyDisabled HTTPSelfCheckStrategy = "Disabled"
)
