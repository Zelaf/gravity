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

// ApiAPIExportOutput struct for ApiAPIExportOutput
type ApiAPIExportOutput struct {
	Entries []ApiAPITransportEntry `json:"entries,omitempty"`
}

// NewApiAPIExportOutput instantiates a new ApiAPIExportOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAPIExportOutput() *ApiAPIExportOutput {
	this := ApiAPIExportOutput{}
	return &this
}

// NewApiAPIExportOutputWithDefaults instantiates a new ApiAPIExportOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAPIExportOutputWithDefaults() *ApiAPIExportOutput {
	this := ApiAPIExportOutput{}
	return &this
}

// GetEntries returns the Entries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiAPIExportOutput) GetEntries() []ApiAPITransportEntry {
	if o == nil {
		var ret []ApiAPITransportEntry
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiAPIExportOutput) GetEntriesOk() ([]ApiAPITransportEntry, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *ApiAPIExportOutput) HasEntries() bool {
	if o != nil && o.Entries != nil {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []ApiAPITransportEntry and assigns it to the Entries field.
func (o *ApiAPIExportOutput) SetEntries(v []ApiAPITransportEntry) {
	o.Entries = v
}

func (o ApiAPIExportOutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}
	return json.Marshal(toSerialize)
}

type NullableApiAPIExportOutput struct {
	value *ApiAPIExportOutput
	isSet bool
}

func (v NullableApiAPIExportOutput) Get() *ApiAPIExportOutput {
	return v.value
}

func (v *NullableApiAPIExportOutput) Set(val *ApiAPIExportOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAPIExportOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAPIExportOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAPIExportOutput(val *ApiAPIExportOutput) *NullableApiAPIExportOutput {
	return &NullableApiAPIExportOutput{value: val, isSet: true}
}

func (v NullableApiAPIExportOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAPIExportOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
