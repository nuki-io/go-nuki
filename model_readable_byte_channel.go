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

// checks if the ReadableByteChannel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadableByteChannel{}

// ReadableByteChannel struct for ReadableByteChannel
type ReadableByteChannel struct {
	Open *bool `json:"open,omitempty"`
}

// NewReadableByteChannel instantiates a new ReadableByteChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadableByteChannel() *ReadableByteChannel {
	this := ReadableByteChannel{}
	return &this
}

// NewReadableByteChannelWithDefaults instantiates a new ReadableByteChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadableByteChannelWithDefaults() *ReadableByteChannel {
	this := ReadableByteChannel{}
	return &this
}

// GetOpen returns the Open field value if set, zero value otherwise.
func (o *ReadableByteChannel) GetOpen() bool {
	if o == nil || IsNil(o.Open) {
		var ret bool
		return ret
	}
	return *o.Open
}

// GetOpenOk returns a tuple with the Open field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadableByteChannel) GetOpenOk() (*bool, bool) {
	if o == nil || IsNil(o.Open) {
		return nil, false
	}
	return o.Open, true
}

// HasOpen returns a boolean if a field has been set.
func (o *ReadableByteChannel) HasOpen() bool {
	if o != nil && !IsNil(o.Open) {
		return true
	}

	return false
}

// SetOpen gets a reference to the given bool and assigns it to the Open field.
func (o *ReadableByteChannel) SetOpen(v bool) {
	o.Open = &v
}

func (o ReadableByteChannel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadableByteChannel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Open) {
		toSerialize["open"] = o.Open
	}
	return toSerialize, nil
}

type NullableReadableByteChannel struct {
	value *ReadableByteChannel
	isSet bool
}

func (v NullableReadableByteChannel) Get() *ReadableByteChannel {
	return v.value
}

func (v *NullableReadableByteChannel) Set(val *ReadableByteChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableReadableByteChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableReadableByteChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadableByteChannel(val *ReadableByteChannel) *NullableReadableByteChannel {
	return &NullableReadableByteChannel{value: val, isSet: true}
}

func (v NullableReadableByteChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadableByteChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


