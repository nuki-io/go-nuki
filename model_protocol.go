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

// checks if the Protocol type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Protocol{}

// Protocol struct for Protocol
type Protocol struct {
	Confidential *bool `json:"confidential,omitempty"`
	DefaultPort *int32 `json:"defaultPort,omitempty"`
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
	SchemeName *string `json:"schemeName,omitempty"`
	TechnicalName *string `json:"technicalName,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewProtocol instantiates a new Protocol object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProtocol() *Protocol {
	this := Protocol{}
	return &this
}

// NewProtocolWithDefaults instantiates a new Protocol object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProtocolWithDefaults() *Protocol {
	this := Protocol{}
	return &this
}

// GetConfidential returns the Confidential field value if set, zero value otherwise.
func (o *Protocol) GetConfidential() bool {
	if o == nil || IsNil(o.Confidential) {
		var ret bool
		return ret
	}
	return *o.Confidential
}

// GetConfidentialOk returns a tuple with the Confidential field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Protocol) GetConfidentialOk() (*bool, bool) {
	if o == nil || IsNil(o.Confidential) {
		return nil, false
	}
	return o.Confidential, true
}

// HasConfidential returns a boolean if a field has been set.
func (o *Protocol) HasConfidential() bool {
	if o != nil && !IsNil(o.Confidential) {
		return true
	}

	return false
}

// SetConfidential gets a reference to the given bool and assigns it to the Confidential field.
func (o *Protocol) SetConfidential(v bool) {
	o.Confidential = &v
}

// GetDefaultPort returns the DefaultPort field value if set, zero value otherwise.
func (o *Protocol) GetDefaultPort() int32 {
	if o == nil || IsNil(o.DefaultPort) {
		var ret int32
		return ret
	}
	return *o.DefaultPort
}

// GetDefaultPortOk returns a tuple with the DefaultPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Protocol) GetDefaultPortOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultPort) {
		return nil, false
	}
	return o.DefaultPort, true
}

// HasDefaultPort returns a boolean if a field has been set.
func (o *Protocol) HasDefaultPort() bool {
	if o != nil && !IsNil(o.DefaultPort) {
		return true
	}

	return false
}

// SetDefaultPort gets a reference to the given int32 and assigns it to the DefaultPort field.
func (o *Protocol) SetDefaultPort(v int32) {
	o.DefaultPort = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Protocol) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Protocol) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Protocol) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Protocol) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Protocol) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Protocol) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Protocol) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Protocol) SetName(v string) {
	o.Name = &v
}

// GetSchemeName returns the SchemeName field value if set, zero value otherwise.
func (o *Protocol) GetSchemeName() string {
	if o == nil || IsNil(o.SchemeName) {
		var ret string
		return ret
	}
	return *o.SchemeName
}

// GetSchemeNameOk returns a tuple with the SchemeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Protocol) GetSchemeNameOk() (*string, bool) {
	if o == nil || IsNil(o.SchemeName) {
		return nil, false
	}
	return o.SchemeName, true
}

// HasSchemeName returns a boolean if a field has been set.
func (o *Protocol) HasSchemeName() bool {
	if o != nil && !IsNil(o.SchemeName) {
		return true
	}

	return false
}

// SetSchemeName gets a reference to the given string and assigns it to the SchemeName field.
func (o *Protocol) SetSchemeName(v string) {
	o.SchemeName = &v
}

// GetTechnicalName returns the TechnicalName field value if set, zero value otherwise.
func (o *Protocol) GetTechnicalName() string {
	if o == nil || IsNil(o.TechnicalName) {
		var ret string
		return ret
	}
	return *o.TechnicalName
}

// GetTechnicalNameOk returns a tuple with the TechnicalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Protocol) GetTechnicalNameOk() (*string, bool) {
	if o == nil || IsNil(o.TechnicalName) {
		return nil, false
	}
	return o.TechnicalName, true
}

// HasTechnicalName returns a boolean if a field has been set.
func (o *Protocol) HasTechnicalName() bool {
	if o != nil && !IsNil(o.TechnicalName) {
		return true
	}

	return false
}

// SetTechnicalName gets a reference to the given string and assigns it to the TechnicalName field.
func (o *Protocol) SetTechnicalName(v string) {
	o.TechnicalName = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Protocol) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Protocol) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Protocol) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Protocol) SetVersion(v string) {
	o.Version = &v
}

func (o Protocol) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Protocol) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Confidential) {
		toSerialize["confidential"] = o.Confidential
	}
	if !IsNil(o.DefaultPort) {
		toSerialize["defaultPort"] = o.DefaultPort
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SchemeName) {
		toSerialize["schemeName"] = o.SchemeName
	}
	if !IsNil(o.TechnicalName) {
		toSerialize["technicalName"] = o.TechnicalName
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableProtocol struct {
	value *Protocol
	isSet bool
}

func (v NullableProtocol) Get() *Protocol {
	return v.value
}

func (v *NullableProtocol) Set(val *Protocol) {
	v.value = val
	v.isSet = true
}

func (v NullableProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtocol(val *Protocol) *NullableProtocol {
	return &NullableProtocol{value: val, isSet: true}
}

func (v NullableProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


