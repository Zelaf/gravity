/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.3.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// DhcpAPIRoleConfigOutput struct for DhcpAPIRoleConfigOutput
type DhcpAPIRoleConfigOutput struct {
	Config DhcpRoleConfig `json:"config"`
}

// NewDhcpAPIRoleConfigOutput instantiates a new DhcpAPIRoleConfigOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpAPIRoleConfigOutput(config DhcpRoleConfig) *DhcpAPIRoleConfigOutput {
	this := DhcpAPIRoleConfigOutput{}
	this.Config = config
	return &this
}

// NewDhcpAPIRoleConfigOutputWithDefaults instantiates a new DhcpAPIRoleConfigOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpAPIRoleConfigOutputWithDefaults() *DhcpAPIRoleConfigOutput {
	this := DhcpAPIRoleConfigOutput{}
	return &this
}

// GetConfig returns the Config field value
func (o *DhcpAPIRoleConfigOutput) GetConfig() DhcpRoleConfig {
	if o == nil {
		var ret DhcpRoleConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *DhcpAPIRoleConfigOutput) GetConfigOk() (*DhcpRoleConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *DhcpAPIRoleConfigOutput) SetConfig(v DhcpRoleConfig) {
	o.Config = v
}

func (o DhcpAPIRoleConfigOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullableDhcpAPIRoleConfigOutput struct {
	value *DhcpAPIRoleConfigOutput
	isSet bool
}

func (v NullableDhcpAPIRoleConfigOutput) Get() *DhcpAPIRoleConfigOutput {
	return v.value
}

func (v *NullableDhcpAPIRoleConfigOutput) Set(val *DhcpAPIRoleConfigOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpAPIRoleConfigOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpAPIRoleConfigOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpAPIRoleConfigOutput(val *DhcpAPIRoleConfigOutput) *NullableDhcpAPIRoleConfigOutput {
	return &NullableDhcpAPIRoleConfigOutput{value: val, isSet: true}
}

func (v NullableDhcpAPIRoleConfigOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpAPIRoleConfigOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
