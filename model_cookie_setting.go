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

// checks if the CookieSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CookieSetting{}

// CookieSetting struct for CookieSetting
type CookieSetting struct {
	Domain *string `json:"domain,omitempty"`
	Name *string `json:"name,omitempty"`
	Path *string `json:"path,omitempty"`
	Value *string `json:"value,omitempty"`
	Version *int32 `json:"version,omitempty"`
	AccessRestricted *bool `json:"accessRestricted,omitempty"`
	Comment *string `json:"comment,omitempty"`
	MaxAge *int32 `json:"maxAge,omitempty"`
	Secure *bool `json:"secure,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewCookieSetting instantiates a new CookieSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCookieSetting() *CookieSetting {
	this := CookieSetting{}
	return &this
}

// NewCookieSettingWithDefaults instantiates a new CookieSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCookieSettingWithDefaults() *CookieSetting {
	this := CookieSetting{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *CookieSetting) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CookieSetting) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *CookieSetting) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *CookieSetting) SetDomain(v string) {
	o.Domain = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CookieSetting) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CookieSetting) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CookieSetting) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CookieSetting) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *CookieSetting) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CookieSetting) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *CookieSetting) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *CookieSetting) SetPath(v string) {
	o.Path = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CookieSetting) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CookieSetting) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CookieSetting) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CookieSetting) SetValue(v string) {
	o.Value = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CookieSetting) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CookieSetting) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CookieSetting) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *CookieSetting) SetVersion(v int32) {
	o.Version = &v
}

// GetAccessRestricted returns the AccessRestricted field value if set, zero value otherwise.
func (o *CookieSetting) GetAccessRestricted() bool {
	if o == nil || IsNil(o.AccessRestricted) {
		var ret bool
		return ret
	}
	return *o.AccessRestricted
}

// GetAccessRestrictedOk returns a tuple with the AccessRestricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CookieSetting) GetAccessRestrictedOk() (*bool, bool) {
	if o == nil || IsNil(o.AccessRestricted) {
		return nil, false
	}
	return o.AccessRestricted, true
}

// HasAccessRestricted returns a boolean if a field has been set.
func (o *CookieSetting) HasAccessRestricted() bool {
	if o != nil && !IsNil(o.AccessRestricted) {
		return true
	}

	return false
}

// SetAccessRestricted gets a reference to the given bool and assigns it to the AccessRestricted field.
func (o *CookieSetting) SetAccessRestricted(v bool) {
	o.AccessRestricted = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CookieSetting) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CookieSetting) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CookieSetting) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CookieSetting) SetComment(v string) {
	o.Comment = &v
}

// GetMaxAge returns the MaxAge field value if set, zero value otherwise.
func (o *CookieSetting) GetMaxAge() int32 {
	if o == nil || IsNil(o.MaxAge) {
		var ret int32
		return ret
	}
	return *o.MaxAge
}

// GetMaxAgeOk returns a tuple with the MaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CookieSetting) GetMaxAgeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAge) {
		return nil, false
	}
	return o.MaxAge, true
}

// HasMaxAge returns a boolean if a field has been set.
func (o *CookieSetting) HasMaxAge() bool {
	if o != nil && !IsNil(o.MaxAge) {
		return true
	}

	return false
}

// SetMaxAge gets a reference to the given int32 and assigns it to the MaxAge field.
func (o *CookieSetting) SetMaxAge(v int32) {
	o.MaxAge = &v
}

// GetSecure returns the Secure field value if set, zero value otherwise.
func (o *CookieSetting) GetSecure() bool {
	if o == nil || IsNil(o.Secure) {
		var ret bool
		return ret
	}
	return *o.Secure
}

// GetSecureOk returns a tuple with the Secure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CookieSetting) GetSecureOk() (*bool, bool) {
	if o == nil || IsNil(o.Secure) {
		return nil, false
	}
	return o.Secure, true
}

// HasSecure returns a boolean if a field has been set.
func (o *CookieSetting) HasSecure() bool {
	if o != nil && !IsNil(o.Secure) {
		return true
	}

	return false
}

// SetSecure gets a reference to the given bool and assigns it to the Secure field.
func (o *CookieSetting) SetSecure(v bool) {
	o.Secure = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CookieSetting) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CookieSetting) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CookieSetting) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CookieSetting) SetDescription(v string) {
	o.Description = &v
}

func (o CookieSetting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CookieSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.AccessRestricted) {
		toSerialize["accessRestricted"] = o.AccessRestricted
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.MaxAge) {
		toSerialize["maxAge"] = o.MaxAge
	}
	if !IsNil(o.Secure) {
		toSerialize["secure"] = o.Secure
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableCookieSetting struct {
	value *CookieSetting
	isSet bool
}

func (v NullableCookieSetting) Get() *CookieSetting {
	return v.value
}

func (v *NullableCookieSetting) Set(val *CookieSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableCookieSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableCookieSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCookieSetting(val *CookieSetting) *NullableCookieSetting {
	return &NullableCookieSetting{value: val, isSet: true}
}

func (v NullableCookieSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCookieSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


