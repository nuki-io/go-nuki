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

// checks if the Method type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Method{}

// Method struct for Method
type Method struct {
	Description *string `json:"description,omitempty"`
	Idempotent *bool `json:"idempotent,omitempty"`
	Name *string `json:"name,omitempty"`
	Replying *bool `json:"replying,omitempty"`
	Safe *bool `json:"safe,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewMethod instantiates a new Method object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMethod() *Method {
	this := Method{}
	return &this
}

// NewMethodWithDefaults instantiates a new Method object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMethodWithDefaults() *Method {
	this := Method{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Method) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Method) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Method) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Method) SetDescription(v string) {
	o.Description = &v
}

// GetIdempotent returns the Idempotent field value if set, zero value otherwise.
func (o *Method) GetIdempotent() bool {
	if o == nil || IsNil(o.Idempotent) {
		var ret bool
		return ret
	}
	return *o.Idempotent
}

// GetIdempotentOk returns a tuple with the Idempotent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Method) GetIdempotentOk() (*bool, bool) {
	if o == nil || IsNil(o.Idempotent) {
		return nil, false
	}
	return o.Idempotent, true
}

// HasIdempotent returns a boolean if a field has been set.
func (o *Method) HasIdempotent() bool {
	if o != nil && !IsNil(o.Idempotent) {
		return true
	}

	return false
}

// SetIdempotent gets a reference to the given bool and assigns it to the Idempotent field.
func (o *Method) SetIdempotent(v bool) {
	o.Idempotent = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Method) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Method) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Method) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Method) SetName(v string) {
	o.Name = &v
}

// GetReplying returns the Replying field value if set, zero value otherwise.
func (o *Method) GetReplying() bool {
	if o == nil || IsNil(o.Replying) {
		var ret bool
		return ret
	}
	return *o.Replying
}

// GetReplyingOk returns a tuple with the Replying field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Method) GetReplyingOk() (*bool, bool) {
	if o == nil || IsNil(o.Replying) {
		return nil, false
	}
	return o.Replying, true
}

// HasReplying returns a boolean if a field has been set.
func (o *Method) HasReplying() bool {
	if o != nil && !IsNil(o.Replying) {
		return true
	}

	return false
}

// SetReplying gets a reference to the given bool and assigns it to the Replying field.
func (o *Method) SetReplying(v bool) {
	o.Replying = &v
}

// GetSafe returns the Safe field value if set, zero value otherwise.
func (o *Method) GetSafe() bool {
	if o == nil || IsNil(o.Safe) {
		var ret bool
		return ret
	}
	return *o.Safe
}

// GetSafeOk returns a tuple with the Safe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Method) GetSafeOk() (*bool, bool) {
	if o == nil || IsNil(o.Safe) {
		return nil, false
	}
	return o.Safe, true
}

// HasSafe returns a boolean if a field has been set.
func (o *Method) HasSafe() bool {
	if o != nil && !IsNil(o.Safe) {
		return true
	}

	return false
}

// SetSafe gets a reference to the given bool and assigns it to the Safe field.
func (o *Method) SetSafe(v bool) {
	o.Safe = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *Method) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Method) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *Method) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *Method) SetUri(v string) {
	o.Uri = &v
}

func (o Method) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Method) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Idempotent) {
		toSerialize["idempotent"] = o.Idempotent
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Replying) {
		toSerialize["replying"] = o.Replying
	}
	if !IsNil(o.Safe) {
		toSerialize["safe"] = o.Safe
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	return toSerialize, nil
}

type NullableMethod struct {
	value *Method
	isSet bool
}

func (v NullableMethod) Get() *Method {
	return v.value
}

func (v *NullableMethod) Set(val *Method) {
	v.value = val
	v.isSet = true
}

func (v NullableMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMethod(val *Method) *NullableMethod {
	return &NullableMethod{value: val, isSet: true}
}

func (v NullableMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


