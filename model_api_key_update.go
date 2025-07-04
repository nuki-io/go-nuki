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

// checks if the ApiKeyUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiKeyUpdate{}

// ApiKeyUpdate struct for ApiKeyUpdate
type ApiKeyUpdate struct {
	// The description
	Description *string `json:"description,omitempty"`
	// The list of redirect uris
	RedirectUris []string `json:"redirectUris,omitempty"`
}

// NewApiKeyUpdate instantiates a new ApiKeyUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeyUpdate() *ApiKeyUpdate {
	this := ApiKeyUpdate{}
	return &this
}

// NewApiKeyUpdateWithDefaults instantiates a new ApiKeyUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyUpdateWithDefaults() *ApiKeyUpdate {
	this := ApiKeyUpdate{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiKeyUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiKeyUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiKeyUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *ApiKeyUpdate) GetRedirectUris() []string {
	if o == nil || IsNil(o.RedirectUris) {
		var ret []string
		return ret
	}
	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeyUpdate) GetRedirectUrisOk() ([]string, bool) {
	if o == nil || IsNil(o.RedirectUris) {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *ApiKeyUpdate) HasRedirectUris() bool {
	if o != nil && !IsNil(o.RedirectUris) {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given []string and assigns it to the RedirectUris field.
func (o *ApiKeyUpdate) SetRedirectUris(v []string) {
	o.RedirectUris = v
}

func (o ApiKeyUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiKeyUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.RedirectUris) {
		toSerialize["redirectUris"] = o.RedirectUris
	}
	return toSerialize, nil
}

type NullableApiKeyUpdate struct {
	value *ApiKeyUpdate
	isSet bool
}

func (v NullableApiKeyUpdate) Get() *ApiKeyUpdate {
	return v.value
}

func (v *NullableApiKeyUpdate) Set(val *ApiKeyUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeyUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeyUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeyUpdate(val *ApiKeyUpdate) *NullableApiKeyUpdate {
	return &NullableApiKeyUpdate{value: val, isSet: true}
}

func (v NullableApiKeyUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeyUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


