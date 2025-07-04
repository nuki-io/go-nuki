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

// checks if the MetadataService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataService{}

// MetadataService struct for MetadataService
type MetadataService struct {
	Context *Context `json:"context,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Started *bool `json:"started,omitempty"`
	DefaultCharacterSet *CharacterSet `json:"defaultCharacterSet,omitempty"`
	DefaultEncoding *Encoding `json:"defaultEncoding,omitempty"`
	DefaultLanguage *Language `json:"defaultLanguage,omitempty"`
	DefaultMediaType *MediaType `json:"defaultMediaType,omitempty"`
	AllMediaTypeExtensionNames []string `json:"allMediaTypeExtensionNames,omitempty"`
	AllCharacterSetExtensionNames []string `json:"allCharacterSetExtensionNames,omitempty"`
	AllEncodingExtensionNames []string `json:"allEncodingExtensionNames,omitempty"`
	AllExtensionNames []string `json:"allExtensionNames,omitempty"`
	AllLanguageExtensionNames []string `json:"allLanguageExtensionNames,omitempty"`
	Stopped *bool `json:"stopped,omitempty"`
}

// NewMetadataService instantiates a new MetadataService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataService() *MetadataService {
	this := MetadataService{}
	return &this
}

// NewMetadataServiceWithDefaults instantiates a new MetadataService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataServiceWithDefaults() *MetadataService {
	this := MetadataService{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *MetadataService) GetContext() Context {
	if o == nil || IsNil(o.Context) {
		var ret Context
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataService) GetContextOk() (*Context, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *MetadataService) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given Context and assigns it to the Context field.
func (o *MetadataService) SetContext(v Context) {
	o.Context = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *MetadataService) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataService) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *MetadataService) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *MetadataService) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetStarted returns the Started field value if set, zero value otherwise.
func (o *MetadataService) GetStarted() bool {
	if o == nil || IsNil(o.Started) {
		var ret bool
		return ret
	}
	return *o.Started
}

// GetStartedOk returns a tuple with the Started field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataService) GetStartedOk() (*bool, bool) {
	if o == nil || IsNil(o.Started) {
		return nil, false
	}
	return o.Started, true
}

// HasStarted returns a boolean if a field has been set.
func (o *MetadataService) HasStarted() bool {
	if o != nil && !IsNil(o.Started) {
		return true
	}

	return false
}

// SetStarted gets a reference to the given bool and assigns it to the Started field.
func (o *MetadataService) SetStarted(v bool) {
	o.Started = &v
}

// GetDefaultCharacterSet returns the DefaultCharacterSet field value if set, zero value otherwise.
func (o *MetadataService) GetDefaultCharacterSet() CharacterSet {
	if o == nil || IsNil(o.DefaultCharacterSet) {
		var ret CharacterSet
		return ret
	}
	return *o.DefaultCharacterSet
}

// GetDefaultCharacterSetOk returns a tuple with the DefaultCharacterSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataService) GetDefaultCharacterSetOk() (*CharacterSet, bool) {
	if o == nil || IsNil(o.DefaultCharacterSet) {
		return nil, false
	}
	return o.DefaultCharacterSet, true
}

// HasDefaultCharacterSet returns a boolean if a field has been set.
func (o *MetadataService) HasDefaultCharacterSet() bool {
	if o != nil && !IsNil(o.DefaultCharacterSet) {
		return true
	}

	return false
}

// SetDefaultCharacterSet gets a reference to the given CharacterSet and assigns it to the DefaultCharacterSet field.
func (o *MetadataService) SetDefaultCharacterSet(v CharacterSet) {
	o.DefaultCharacterSet = &v
}

// GetDefaultEncoding returns the DefaultEncoding field value if set, zero value otherwise.
func (o *MetadataService) GetDefaultEncoding() Encoding {
	if o == nil || IsNil(o.DefaultEncoding) {
		var ret Encoding
		return ret
	}
	return *o.DefaultEncoding
}

// GetDefaultEncodingOk returns a tuple with the DefaultEncoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataService) GetDefaultEncodingOk() (*Encoding, bool) {
	if o == nil || IsNil(o.DefaultEncoding) {
		return nil, false
	}
	return o.DefaultEncoding, true
}

// HasDefaultEncoding returns a boolean if a field has been set.
func (o *MetadataService) HasDefaultEncoding() bool {
	if o != nil && !IsNil(o.DefaultEncoding) {
		return true
	}

	return false
}

// SetDefaultEncoding gets a reference to the given Encoding and assigns it to the DefaultEncoding field.
func (o *MetadataService) SetDefaultEncoding(v Encoding) {
	o.DefaultEncoding = &v
}

// GetDefaultLanguage returns the DefaultLanguage field value if set, zero value otherwise.
func (o *MetadataService) GetDefaultLanguage() Language {
	if o == nil || IsNil(o.DefaultLanguage) {
		var ret Language
		return ret
	}
	return *o.DefaultLanguage
}

// GetDefaultLanguageOk returns a tuple with the DefaultLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataService) GetDefaultLanguageOk() (*Language, bool) {
	if o == nil || IsNil(o.DefaultLanguage) {
		return nil, false
	}
	return o.DefaultLanguage, true
}

// HasDefaultLanguage returns a boolean if a field has been set.
func (o *MetadataService) HasDefaultLanguage() bool {
	if o != nil && !IsNil(o.DefaultLanguage) {
		return true
	}

	return false
}

// SetDefaultLanguage gets a reference to the given Language and assigns it to the DefaultLanguage field.
func (o *MetadataService) SetDefaultLanguage(v Language) {
	o.DefaultLanguage = &v
}

// GetDefaultMediaType returns the DefaultMediaType field value if set, zero value otherwise.
func (o *MetadataService) GetDefaultMediaType() MediaType {
	if o == nil || IsNil(o.DefaultMediaType) {
		var ret MediaType
		return ret
	}
	return *o.DefaultMediaType
}

// GetDefaultMediaTypeOk returns a tuple with the DefaultMediaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataService) GetDefaultMediaTypeOk() (*MediaType, bool) {
	if o == nil || IsNil(o.DefaultMediaType) {
		return nil, false
	}
	return o.DefaultMediaType, true
}

// HasDefaultMediaType returns a boolean if a field has been set.
func (o *MetadataService) HasDefaultMediaType() bool {
	if o != nil && !IsNil(o.DefaultMediaType) {
		return true
	}

	return false
}

// SetDefaultMediaType gets a reference to the given MediaType and assigns it to the DefaultMediaType field.
func (o *MetadataService) SetDefaultMediaType(v MediaType) {
	o.DefaultMediaType = &v
}

// GetAllMediaTypeExtensionNames returns the AllMediaTypeExtensionNames field value if set, zero value otherwise.
func (o *MetadataService) GetAllMediaTypeExtensionNames() []string {
	if o == nil || IsNil(o.AllMediaTypeExtensionNames) {
		var ret []string
		return ret
	}
	return o.AllMediaTypeExtensionNames
}

// GetAllMediaTypeExtensionNamesOk returns a tuple with the AllMediaTypeExtensionNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataService) GetAllMediaTypeExtensionNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllMediaTypeExtensionNames) {
		return nil, false
	}
	return o.AllMediaTypeExtensionNames, true
}

// HasAllMediaTypeExtensionNames returns a boolean if a field has been set.
func (o *MetadataService) HasAllMediaTypeExtensionNames() bool {
	if o != nil && !IsNil(o.AllMediaTypeExtensionNames) {
		return true
	}

	return false
}

// SetAllMediaTypeExtensionNames gets a reference to the given []string and assigns it to the AllMediaTypeExtensionNames field.
func (o *MetadataService) SetAllMediaTypeExtensionNames(v []string) {
	o.AllMediaTypeExtensionNames = v
}

// GetAllCharacterSetExtensionNames returns the AllCharacterSetExtensionNames field value if set, zero value otherwise.
func (o *MetadataService) GetAllCharacterSetExtensionNames() []string {
	if o == nil || IsNil(o.AllCharacterSetExtensionNames) {
		var ret []string
		return ret
	}
	return o.AllCharacterSetExtensionNames
}

// GetAllCharacterSetExtensionNamesOk returns a tuple with the AllCharacterSetExtensionNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataService) GetAllCharacterSetExtensionNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllCharacterSetExtensionNames) {
		return nil, false
	}
	return o.AllCharacterSetExtensionNames, true
}

// HasAllCharacterSetExtensionNames returns a boolean if a field has been set.
func (o *MetadataService) HasAllCharacterSetExtensionNames() bool {
	if o != nil && !IsNil(o.AllCharacterSetExtensionNames) {
		return true
	}

	return false
}

// SetAllCharacterSetExtensionNames gets a reference to the given []string and assigns it to the AllCharacterSetExtensionNames field.
func (o *MetadataService) SetAllCharacterSetExtensionNames(v []string) {
	o.AllCharacterSetExtensionNames = v
}

// GetAllEncodingExtensionNames returns the AllEncodingExtensionNames field value if set, zero value otherwise.
func (o *MetadataService) GetAllEncodingExtensionNames() []string {
	if o == nil || IsNil(o.AllEncodingExtensionNames) {
		var ret []string
		return ret
	}
	return o.AllEncodingExtensionNames
}

// GetAllEncodingExtensionNamesOk returns a tuple with the AllEncodingExtensionNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataService) GetAllEncodingExtensionNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllEncodingExtensionNames) {
		return nil, false
	}
	return o.AllEncodingExtensionNames, true
}

// HasAllEncodingExtensionNames returns a boolean if a field has been set.
func (o *MetadataService) HasAllEncodingExtensionNames() bool {
	if o != nil && !IsNil(o.AllEncodingExtensionNames) {
		return true
	}

	return false
}

// SetAllEncodingExtensionNames gets a reference to the given []string and assigns it to the AllEncodingExtensionNames field.
func (o *MetadataService) SetAllEncodingExtensionNames(v []string) {
	o.AllEncodingExtensionNames = v
}

// GetAllExtensionNames returns the AllExtensionNames field value if set, zero value otherwise.
func (o *MetadataService) GetAllExtensionNames() []string {
	if o == nil || IsNil(o.AllExtensionNames) {
		var ret []string
		return ret
	}
	return o.AllExtensionNames
}

// GetAllExtensionNamesOk returns a tuple with the AllExtensionNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataService) GetAllExtensionNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllExtensionNames) {
		return nil, false
	}
	return o.AllExtensionNames, true
}

// HasAllExtensionNames returns a boolean if a field has been set.
func (o *MetadataService) HasAllExtensionNames() bool {
	if o != nil && !IsNil(o.AllExtensionNames) {
		return true
	}

	return false
}

// SetAllExtensionNames gets a reference to the given []string and assigns it to the AllExtensionNames field.
func (o *MetadataService) SetAllExtensionNames(v []string) {
	o.AllExtensionNames = v
}

// GetAllLanguageExtensionNames returns the AllLanguageExtensionNames field value if set, zero value otherwise.
func (o *MetadataService) GetAllLanguageExtensionNames() []string {
	if o == nil || IsNil(o.AllLanguageExtensionNames) {
		var ret []string
		return ret
	}
	return o.AllLanguageExtensionNames
}

// GetAllLanguageExtensionNamesOk returns a tuple with the AllLanguageExtensionNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataService) GetAllLanguageExtensionNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.AllLanguageExtensionNames) {
		return nil, false
	}
	return o.AllLanguageExtensionNames, true
}

// HasAllLanguageExtensionNames returns a boolean if a field has been set.
func (o *MetadataService) HasAllLanguageExtensionNames() bool {
	if o != nil && !IsNil(o.AllLanguageExtensionNames) {
		return true
	}

	return false
}

// SetAllLanguageExtensionNames gets a reference to the given []string and assigns it to the AllLanguageExtensionNames field.
func (o *MetadataService) SetAllLanguageExtensionNames(v []string) {
	o.AllLanguageExtensionNames = v
}

// GetStopped returns the Stopped field value if set, zero value otherwise.
func (o *MetadataService) GetStopped() bool {
	if o == nil || IsNil(o.Stopped) {
		var ret bool
		return ret
	}
	return *o.Stopped
}

// GetStoppedOk returns a tuple with the Stopped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataService) GetStoppedOk() (*bool, bool) {
	if o == nil || IsNil(o.Stopped) {
		return nil, false
	}
	return o.Stopped, true
}

// HasStopped returns a boolean if a field has been set.
func (o *MetadataService) HasStopped() bool {
	if o != nil && !IsNil(o.Stopped) {
		return true
	}

	return false
}

// SetStopped gets a reference to the given bool and assigns it to the Stopped field.
func (o *MetadataService) SetStopped(v bool) {
	o.Stopped = &v
}

func (o MetadataService) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Started) {
		toSerialize["started"] = o.Started
	}
	if !IsNil(o.DefaultCharacterSet) {
		toSerialize["defaultCharacterSet"] = o.DefaultCharacterSet
	}
	if !IsNil(o.DefaultEncoding) {
		toSerialize["defaultEncoding"] = o.DefaultEncoding
	}
	if !IsNil(o.DefaultLanguage) {
		toSerialize["defaultLanguage"] = o.DefaultLanguage
	}
	if !IsNil(o.DefaultMediaType) {
		toSerialize["defaultMediaType"] = o.DefaultMediaType
	}
	if !IsNil(o.AllMediaTypeExtensionNames) {
		toSerialize["allMediaTypeExtensionNames"] = o.AllMediaTypeExtensionNames
	}
	if !IsNil(o.AllCharacterSetExtensionNames) {
		toSerialize["allCharacterSetExtensionNames"] = o.AllCharacterSetExtensionNames
	}
	if !IsNil(o.AllEncodingExtensionNames) {
		toSerialize["allEncodingExtensionNames"] = o.AllEncodingExtensionNames
	}
	if !IsNil(o.AllExtensionNames) {
		toSerialize["allExtensionNames"] = o.AllExtensionNames
	}
	if !IsNil(o.AllLanguageExtensionNames) {
		toSerialize["allLanguageExtensionNames"] = o.AllLanguageExtensionNames
	}
	if !IsNil(o.Stopped) {
		toSerialize["stopped"] = o.Stopped
	}
	return toSerialize, nil
}

type NullableMetadataService struct {
	value *MetadataService
	isSet bool
}

func (v NullableMetadataService) Get() *MetadataService {
	return v.value
}

func (v *NullableMetadataService) Set(val *MetadataService) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataService) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataService(val *MetadataService) *NullableMetadataService {
	return &NullableMetadataService{value: val, isSet: true}
}

func (v NullableMetadataService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


