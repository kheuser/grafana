// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v0alpha1

// ResourceCountApplyConfiguration represents a declarative configuration of the ResourceCount type for use
// with apply.
type ResourceCountApplyConfiguration struct {
	Group    *string `json:"group,omitempty"`
	Resource *string `json:"resource,omitempty"`
	Count    *int64  `json:"count,omitempty"`
}

// ResourceCountApplyConfiguration constructs a declarative configuration of the ResourceCount type for use with
// apply.
func ResourceCount() *ResourceCountApplyConfiguration {
	return &ResourceCountApplyConfiguration{}
}

// WithGroup sets the Group field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Group field is set to the value of the last call.
func (b *ResourceCountApplyConfiguration) WithGroup(value string) *ResourceCountApplyConfiguration {
	b.Group = &value
	return b
}

// WithResource sets the Resource field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resource field is set to the value of the last call.
func (b *ResourceCountApplyConfiguration) WithResource(value string) *ResourceCountApplyConfiguration {
	b.Resource = &value
	return b
}

// WithCount sets the Count field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Count field is set to the value of the last call.
func (b *ResourceCountApplyConfiguration) WithCount(value int64) *ResourceCountApplyConfiguration {
	b.Count = &value
	return b
}
