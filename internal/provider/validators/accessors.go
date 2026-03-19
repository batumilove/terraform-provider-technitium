// Copyright (c) 2026 Alex Ackerman
// SPDX-License-Identifier: MPL-2.0

package validators

// ConfigAccessor provides read access to Terraform configuration values.
type ConfigAccessor interface {
	GetString(path string) (string, bool)
	GetBool(path string) (bool, bool)
	GetStringList(path string) ([]string, bool)
}

// PlanAccessor provides read access to Terraform plan values.
type PlanAccessor interface {
	ConfigAccessor
}

// StateAccessor provides read access to Terraform state values.
type StateAccessor interface {
	ConfigAccessor
}
