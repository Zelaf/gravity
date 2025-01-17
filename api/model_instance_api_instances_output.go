/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.6.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the InstanceAPIInstancesOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceAPIInstancesOutput{}

// InstanceAPIInstancesOutput struct for InstanceAPIInstancesOutput
type InstanceAPIInstancesOutput struct {
	Instances []InstanceInstanceInfo `json:"instances"`
}

// NewInstanceAPIInstancesOutput instantiates a new InstanceAPIInstancesOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceAPIInstancesOutput(instances []InstanceInstanceInfo) *InstanceAPIInstancesOutput {
	this := InstanceAPIInstancesOutput{}
	this.Instances = instances
	return &this
}

// NewInstanceAPIInstancesOutputWithDefaults instantiates a new InstanceAPIInstancesOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceAPIInstancesOutputWithDefaults() *InstanceAPIInstancesOutput {
	this := InstanceAPIInstancesOutput{}
	return &this
}

// GetInstances returns the Instances field value
// If the value is explicit nil, the zero value for []InstanceInstanceInfo will be returned
func (o *InstanceAPIInstancesOutput) GetInstances() []InstanceInstanceInfo {
	if o == nil {
		var ret []InstanceInstanceInfo
		return ret
	}

	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstanceAPIInstancesOutput) GetInstancesOk() ([]InstanceInstanceInfo, bool) {
	if o == nil || IsNil(o.Instances) {
		return nil, false
	}
	return o.Instances, true
}

// SetInstances sets field value
func (o *InstanceAPIInstancesOutput) SetInstances(v []InstanceInstanceInfo) {
	o.Instances = v
}

func (o InstanceAPIInstancesOutput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceAPIInstancesOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Instances != nil {
		toSerialize["instances"] = o.Instances
	}
	return toSerialize, nil
}

type NullableInstanceAPIInstancesOutput struct {
	value *InstanceAPIInstancesOutput
	isSet bool
}

func (v NullableInstanceAPIInstancesOutput) Get() *InstanceAPIInstancesOutput {
	return v.value
}

func (v *NullableInstanceAPIInstancesOutput) Set(val *InstanceAPIInstancesOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceAPIInstancesOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceAPIInstancesOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceAPIInstancesOutput(val *InstanceAPIInstancesOutput) *NullableInstanceAPIInstancesOutput {
	return &NullableInstanceAPIInstancesOutput{value: val, isSet: true}
}

func (v NullableInstanceAPIInstancesOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceAPIInstancesOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
