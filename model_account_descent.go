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
	"bytes"
	"fmt"
)

// checks if the AccountDescent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountDescent{}

// AccountDescent struct for AccountDescent
type AccountDescent struct {
	// The account origin source
	Origin string `json:"origin"`
}

type _AccountDescent AccountDescent

// NewAccountDescent instantiates a new AccountDescent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountDescent(origin string) *AccountDescent {
	this := AccountDescent{}
	this.Origin = origin
	return &this
}

// NewAccountDescentWithDefaults instantiates a new AccountDescent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountDescentWithDefaults() *AccountDescent {
	this := AccountDescent{}
	return &this
}

// GetOrigin returns the Origin field value
func (o *AccountDescent) GetOrigin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Origin
}

// GetOriginOk returns a tuple with the Origin field value
// and a boolean to check if the value has been set.
func (o *AccountDescent) GetOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Origin, true
}

// SetOrigin sets field value
func (o *AccountDescent) SetOrigin(v string) {
	o.Origin = v
}

func (o AccountDescent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountDescent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["origin"] = o.Origin
	return toSerialize, nil
}

func (o *AccountDescent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"origin",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAccountDescent := _AccountDescent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountDescent)

	if err != nil {
		return err
	}

	*o = AccountDescent(varAccountDescent)

	return err
}

type NullableAccountDescent struct {
	value *AccountDescent
	isSet bool
}

func (v NullableAccountDescent) Get() *AccountDescent {
	return v.value
}

func (v *NullableAccountDescent) Set(val *AccountDescent) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountDescent) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountDescent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountDescent(val *AccountDescent) *NullableAccountDescent {
	return &NullableAccountDescent{value: val, isSet: true}
}

func (v NullableAccountDescent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountDescent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


