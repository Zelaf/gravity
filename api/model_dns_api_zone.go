/*
gravity

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.6.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DnsAPIZone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsAPIZone{}

// DnsAPIZone struct for DnsAPIZone
type DnsAPIZone struct {
	Authoritative  bool                `json:"authoritative"`
	DefaultTTL     int32               `json:"defaultTTL"`
	HandlerConfigs []map[string]string `json:"handlerConfigs"`
	Name           string              `json:"name"`
}

// NewDnsAPIZone instantiates a new DnsAPIZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsAPIZone(authoritative bool, defaultTTL int32, handlerConfigs []map[string]string, name string) *DnsAPIZone {
	this := DnsAPIZone{}
	this.Authoritative = authoritative
	this.DefaultTTL = defaultTTL
	this.HandlerConfigs = handlerConfigs
	this.Name = name
	return &this
}

// NewDnsAPIZoneWithDefaults instantiates a new DnsAPIZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsAPIZoneWithDefaults() *DnsAPIZone {
	this := DnsAPIZone{}
	return &this
}

// GetAuthoritative returns the Authoritative field value
func (o *DnsAPIZone) GetAuthoritative() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Authoritative
}

// GetAuthoritativeOk returns a tuple with the Authoritative field value
// and a boolean to check if the value has been set.
func (o *DnsAPIZone) GetAuthoritativeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authoritative, true
}

// SetAuthoritative sets field value
func (o *DnsAPIZone) SetAuthoritative(v bool) {
	o.Authoritative = v
}

// GetDefaultTTL returns the DefaultTTL field value
func (o *DnsAPIZone) GetDefaultTTL() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DefaultTTL
}

// GetDefaultTTLOk returns a tuple with the DefaultTTL field value
// and a boolean to check if the value has been set.
func (o *DnsAPIZone) GetDefaultTTLOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultTTL, true
}

// SetDefaultTTL sets field value
func (o *DnsAPIZone) SetDefaultTTL(v int32) {
	o.DefaultTTL = v
}

// GetHandlerConfigs returns the HandlerConfigs field value
// If the value is explicit nil, the zero value for []map[string]string will be returned
func (o *DnsAPIZone) GetHandlerConfigs() []map[string]string {
	if o == nil {
		var ret []map[string]string
		return ret
	}

	return o.HandlerConfigs
}

// GetHandlerConfigsOk returns a tuple with the HandlerConfigs field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DnsAPIZone) GetHandlerConfigsOk() ([]map[string]string, bool) {
	if o == nil || IsNil(o.HandlerConfigs) {
		return nil, false
	}
	return o.HandlerConfigs, true
}

// SetHandlerConfigs sets field value
func (o *DnsAPIZone) SetHandlerConfigs(v []map[string]string) {
	o.HandlerConfigs = v
}

// GetName returns the Name field value
func (o *DnsAPIZone) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DnsAPIZone) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DnsAPIZone) SetName(v string) {
	o.Name = v
}

func (o DnsAPIZone) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsAPIZone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authoritative"] = o.Authoritative
	toSerialize["defaultTTL"] = o.DefaultTTL
	if o.HandlerConfigs != nil {
		toSerialize["handlerConfigs"] = o.HandlerConfigs
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableDnsAPIZone struct {
	value *DnsAPIZone
	isSet bool
}

func (v NullableDnsAPIZone) Get() *DnsAPIZone {
	return v.value
}

func (v *NullableDnsAPIZone) Set(val *DnsAPIZone) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsAPIZone) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsAPIZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsAPIZone(val *DnsAPIZone) *NullableDnsAPIZone {
	return &NullableDnsAPIZone{value: val, isSet: true}
}

func (v NullableDnsAPIZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsAPIZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
