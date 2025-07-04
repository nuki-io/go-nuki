/*
Nuki API

The Nuki Web Api

API version: 3.10.1
Contact: contact@nuki.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the NamedValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NamedValue{}

// NamedValue struct for NamedValue
type NamedValue struct {
	Name *string `json:"name,omitempty"`
	Value map[string]interface{} `json:"value,omitempty"`
}

// NewNamedValue instantiates a new NamedValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNamedValue() *NamedValue {
	this := NamedValue{}
	return &this
}

// NewNamedValueWithDefaults instantiates a new NamedValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNamedValueWithDefaults() *NamedValue {
	this := NamedValue{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NamedValue) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamedValue) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NamedValue) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NamedValue) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *NamedValue) GetValue() map[string]interface{} {
	if o == nil || IsNil(o.Value) {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NamedValue) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return map[string]interface{}{}, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *NamedValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *NamedValue) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o NamedValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NamedValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableNamedValue struct {
	value *NamedValue
	isSet bool
}

func (v NullableNamedValue) Get() *NamedValue {
	return v.value
}

func (v *NullableNamedValue) Set(val *NamedValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNamedValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNamedValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNamedValue(val *NamedValue) *NullableNamedValue {
	return &NullableNamedValue{value: val, isSet: true}
}

func (v NullableNamedValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNamedValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


