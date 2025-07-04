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

// checks if the AccountUserUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountUserUpdate{}

// AccountUserUpdate struct for AccountUserUpdate
type AccountUserUpdate struct {
	// The new email address
	Email *string `json:"email,omitempty"`
	// The new name of the sub account
	Name *string `json:"name,omitempty"`
	// The new language code
	Language *string `json:"language,omitempty"`
}

// NewAccountUserUpdate instantiates a new AccountUserUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUserUpdate() *AccountUserUpdate {
	this := AccountUserUpdate{}
	return &this
}

// NewAccountUserUpdateWithDefaults instantiates a new AccountUserUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUserUpdateWithDefaults() *AccountUserUpdate {
	this := AccountUserUpdate{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AccountUserUpdate) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUserUpdate) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AccountUserUpdate) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AccountUserUpdate) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccountUserUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUserUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccountUserUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccountUserUpdate) SetName(v string) {
	o.Name = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *AccountUserUpdate) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUserUpdate) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *AccountUserUpdate) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *AccountUserUpdate) SetLanguage(v string) {
	o.Language = &v
}

func (o AccountUserUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountUserUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	return toSerialize, nil
}

type NullableAccountUserUpdate struct {
	value *AccountUserUpdate
	isSet bool
}

func (v NullableAccountUserUpdate) Get() *AccountUserUpdate {
	return v.value
}

func (v *NullableAccountUserUpdate) Set(val *AccountUserUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUserUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUserUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUserUpdate(val *AccountUserUpdate) *NullableAccountUserUpdate {
	return &NullableAccountUserUpdate{value: val, isSet: true}
}

func (v NullableAccountUserUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUserUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


