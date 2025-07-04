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

// checks if the PublicKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicKey{}

// PublicKey struct for PublicKey
type PublicKey struct {
	Algorithm *string `json:"algorithm,omitempty"`
	Format *string `json:"format,omitempty"`
	Encoded []string `json:"encoded,omitempty"`
}

// NewPublicKey instantiates a new PublicKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicKey() *PublicKey {
	this := PublicKey{}
	return &this
}

// NewPublicKeyWithDefaults instantiates a new PublicKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicKeyWithDefaults() *PublicKey {
	this := PublicKey{}
	return &this
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *PublicKey) GetAlgorithm() string {
	if o == nil || IsNil(o.Algorithm) {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicKey) GetAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.Algorithm) {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *PublicKey) HasAlgorithm() bool {
	if o != nil && !IsNil(o.Algorithm) {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *PublicKey) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *PublicKey) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicKey) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *PublicKey) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *PublicKey) SetFormat(v string) {
	o.Format = &v
}

// GetEncoded returns the Encoded field value if set, zero value otherwise.
func (o *PublicKey) GetEncoded() []string {
	if o == nil || IsNil(o.Encoded) {
		var ret []string
		return ret
	}
	return o.Encoded
}

// GetEncodedOk returns a tuple with the Encoded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicKey) GetEncodedOk() ([]string, bool) {
	if o == nil || IsNil(o.Encoded) {
		return nil, false
	}
	return o.Encoded, true
}

// HasEncoded returns a boolean if a field has been set.
func (o *PublicKey) HasEncoded() bool {
	if o != nil && !IsNil(o.Encoded) {
		return true
	}

	return false
}

// SetEncoded gets a reference to the given []string and assigns it to the Encoded field.
func (o *PublicKey) SetEncoded(v []string) {
	o.Encoded = v
}

func (o PublicKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Algorithm) {
		toSerialize["algorithm"] = o.Algorithm
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.Encoded) {
		toSerialize["encoded"] = o.Encoded
	}
	return toSerialize, nil
}

type NullablePublicKey struct {
	value *PublicKey
	isSet bool
}

func (v NullablePublicKey) Get() *PublicKey {
	return v.value
}

func (v *NullablePublicKey) Set(val *PublicKey) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicKey) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicKey(val *PublicKey) *NullablePublicKey {
	return &NullablePublicKey{value: val, isSet: true}
}

func (v NullablePublicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


