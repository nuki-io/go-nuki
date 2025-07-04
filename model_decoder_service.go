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

// checks if the DecoderService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DecoderService{}

// DecoderService struct for DecoderService
type DecoderService struct {
	Context *Context `json:"context,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Started *bool `json:"started,omitempty"`
	Stopped *bool `json:"stopped,omitempty"`
}

// NewDecoderService instantiates a new DecoderService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecoderService() *DecoderService {
	this := DecoderService{}
	return &this
}

// NewDecoderServiceWithDefaults instantiates a new DecoderService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecoderServiceWithDefaults() *DecoderService {
	this := DecoderService{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *DecoderService) GetContext() Context {
	if o == nil || IsNil(o.Context) {
		var ret Context
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecoderService) GetContextOk() (*Context, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *DecoderService) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given Context and assigns it to the Context field.
func (o *DecoderService) SetContext(v Context) {
	o.Context = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DecoderService) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecoderService) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DecoderService) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DecoderService) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetStarted returns the Started field value if set, zero value otherwise.
func (o *DecoderService) GetStarted() bool {
	if o == nil || IsNil(o.Started) {
		var ret bool
		return ret
	}
	return *o.Started
}

// GetStartedOk returns a tuple with the Started field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecoderService) GetStartedOk() (*bool, bool) {
	if o == nil || IsNil(o.Started) {
		return nil, false
	}
	return o.Started, true
}

// HasStarted returns a boolean if a field has been set.
func (o *DecoderService) HasStarted() bool {
	if o != nil && !IsNil(o.Started) {
		return true
	}

	return false
}

// SetStarted gets a reference to the given bool and assigns it to the Started field.
func (o *DecoderService) SetStarted(v bool) {
	o.Started = &v
}

// GetStopped returns the Stopped field value if set, zero value otherwise.
func (o *DecoderService) GetStopped() bool {
	if o == nil || IsNil(o.Stopped) {
		var ret bool
		return ret
	}
	return *o.Stopped
}

// GetStoppedOk returns a tuple with the Stopped field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecoderService) GetStoppedOk() (*bool, bool) {
	if o == nil || IsNil(o.Stopped) {
		return nil, false
	}
	return o.Stopped, true
}

// HasStopped returns a boolean if a field has been set.
func (o *DecoderService) HasStopped() bool {
	if o != nil && !IsNil(o.Stopped) {
		return true
	}

	return false
}

// SetStopped gets a reference to the given bool and assigns it to the Stopped field.
func (o *DecoderService) SetStopped(v bool) {
	o.Stopped = &v
}

func (o DecoderService) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DecoderService) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Stopped) {
		toSerialize["stopped"] = o.Stopped
	}
	return toSerialize, nil
}

type NullableDecoderService struct {
	value *DecoderService
	isSet bool
}

func (v NullableDecoderService) Get() *DecoderService {
	return v.value
}

func (v *NullableDecoderService) Set(val *DecoderService) {
	v.value = val
	v.isSet = true
}

func (v NullableDecoderService) IsSet() bool {
	return v.isSet
}

func (v *NullableDecoderService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecoderService(val *DecoderService) *NullableDecoderService {
	return &NullableDecoderService{value: val, isSet: true}
}

func (v NullableDecoderService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecoderService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


