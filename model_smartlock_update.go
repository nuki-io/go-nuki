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

// checks if the SmartlockUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmartlockUpdate{}

// SmartlockUpdate struct for SmartlockUpdate
type SmartlockUpdate struct {
	// The new name of the smartlock
	Name *string `json:"name,omitempty"`
	// True if the smartlock is favorite
	Favorite *bool `json:"favorite,omitempty"`
}

// NewSmartlockUpdate instantiates a new SmartlockUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartlockUpdate() *SmartlockUpdate {
	this := SmartlockUpdate{}
	return &this
}

// NewSmartlockUpdateWithDefaults instantiates a new SmartlockUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartlockUpdateWithDefaults() *SmartlockUpdate {
	this := SmartlockUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SmartlockUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SmartlockUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SmartlockUpdate) SetName(v string) {
	o.Name = &v
}

// GetFavorite returns the Favorite field value if set, zero value otherwise.
func (o *SmartlockUpdate) GetFavorite() bool {
	if o == nil || IsNil(o.Favorite) {
		var ret bool
		return ret
	}
	return *o.Favorite
}

// GetFavoriteOk returns a tuple with the Favorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockUpdate) GetFavoriteOk() (*bool, bool) {
	if o == nil || IsNil(o.Favorite) {
		return nil, false
	}
	return o.Favorite, true
}

// HasFavorite returns a boolean if a field has been set.
func (o *SmartlockUpdate) HasFavorite() bool {
	if o != nil && !IsNil(o.Favorite) {
		return true
	}

	return false
}

// SetFavorite gets a reference to the given bool and assigns it to the Favorite field.
func (o *SmartlockUpdate) SetFavorite(v bool) {
	o.Favorite = &v
}

func (o SmartlockUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmartlockUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Favorite) {
		toSerialize["favorite"] = o.Favorite
	}
	return toSerialize, nil
}

type NullableSmartlockUpdate struct {
	value *SmartlockUpdate
	isSet bool
}

func (v NullableSmartlockUpdate) Get() *SmartlockUpdate {
	return v.value
}

func (v *NullableSmartlockUpdate) Set(val *SmartlockUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartlockUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartlockUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartlockUpdate(val *SmartlockUpdate) *NullableSmartlockUpdate {
	return &NullableSmartlockUpdate{value: val, isSet: true}
}

func (v NullableSmartlockUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartlockUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


