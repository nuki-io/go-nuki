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
	"time"
	"bytes"
	"fmt"
)

// checks if the SmartlockAuthCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmartlockAuthCreate{}

// SmartlockAuthCreate struct for SmartlockAuthCreate
type SmartlockAuthCreate struct {
	// The name of the authorization (max 32 chars)
	Name string `json:"name"`
	// The allowed from date
	AllowedFromDate *time.Time `json:"allowedFromDate,omitempty"`
	// The allowed until date
	AllowedUntilDate *time.Time `json:"allowedUntilDate,omitempty"`
	// The allowed weekdays bitmask: 64 .. monday, 32 .. tuesday, 16 .. wednesday, 8 .. thursday, 4 .. friday, 2 .. saturday, 1 .. sunday
	AllowedWeekDays *int32 `json:"allowedWeekDays,omitempty"`
	// The allowed from time (in minutes from midnight)
	AllowedFromTime *int32 `json:"allowedFromTime,omitempty"`
	// The allowed until time (in minutes from midnight)
	AllowedUntilTime *int32 `json:"allowedUntilTime,omitempty"`
	// The id of the linked account user (required if type is NOT 13 .. keypad)
	AccountUserId *int32 `json:"accountUserId,omitempty"`
	// True if the auth has remote access
	RemoteAllowed bool `json:"remoteAllowed"`
	// The smart actions enabled flag
	SmartActionsEnabled *bool `json:"smartActionsEnabled,omitempty"`
	// The optional type of the auth 0 .. app (default), 2 .. fob, 13 .. keypad
	Type *int32 `json:"type,omitempty"`
	// The code of the keypad authorization (only for keypad)
	Code *int32 `json:"code,omitempty"`
}

type _SmartlockAuthCreate SmartlockAuthCreate

// NewSmartlockAuthCreate instantiates a new SmartlockAuthCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartlockAuthCreate(name string, remoteAllowed bool) *SmartlockAuthCreate {
	this := SmartlockAuthCreate{}
	this.Name = name
	this.RemoteAllowed = remoteAllowed
	return &this
}

// NewSmartlockAuthCreateWithDefaults instantiates a new SmartlockAuthCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartlockAuthCreateWithDefaults() *SmartlockAuthCreate {
	this := SmartlockAuthCreate{}
	return &this
}

// GetName returns the Name field value
func (o *SmartlockAuthCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SmartlockAuthCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SmartlockAuthCreate) SetName(v string) {
	o.Name = v
}

// GetAllowedFromDate returns the AllowedFromDate field value if set, zero value otherwise.
func (o *SmartlockAuthCreate) GetAllowedFromDate() time.Time {
	if o == nil || IsNil(o.AllowedFromDate) {
		var ret time.Time
		return ret
	}
	return *o.AllowedFromDate
}

// GetAllowedFromDateOk returns a tuple with the AllowedFromDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockAuthCreate) GetAllowedFromDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AllowedFromDate) {
		return nil, false
	}
	return o.AllowedFromDate, true
}

// HasAllowedFromDate returns a boolean if a field has been set.
func (o *SmartlockAuthCreate) HasAllowedFromDate() bool {
	if o != nil && !IsNil(o.AllowedFromDate) {
		return true
	}

	return false
}

// SetAllowedFromDate gets a reference to the given time.Time and assigns it to the AllowedFromDate field.
func (o *SmartlockAuthCreate) SetAllowedFromDate(v time.Time) {
	o.AllowedFromDate = &v
}

// GetAllowedUntilDate returns the AllowedUntilDate field value if set, zero value otherwise.
func (o *SmartlockAuthCreate) GetAllowedUntilDate() time.Time {
	if o == nil || IsNil(o.AllowedUntilDate) {
		var ret time.Time
		return ret
	}
	return *o.AllowedUntilDate
}

// GetAllowedUntilDateOk returns a tuple with the AllowedUntilDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockAuthCreate) GetAllowedUntilDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AllowedUntilDate) {
		return nil, false
	}
	return o.AllowedUntilDate, true
}

// HasAllowedUntilDate returns a boolean if a field has been set.
func (o *SmartlockAuthCreate) HasAllowedUntilDate() bool {
	if o != nil && !IsNil(o.AllowedUntilDate) {
		return true
	}

	return false
}

// SetAllowedUntilDate gets a reference to the given time.Time and assigns it to the AllowedUntilDate field.
func (o *SmartlockAuthCreate) SetAllowedUntilDate(v time.Time) {
	o.AllowedUntilDate = &v
}

// GetAllowedWeekDays returns the AllowedWeekDays field value if set, zero value otherwise.
func (o *SmartlockAuthCreate) GetAllowedWeekDays() int32 {
	if o == nil || IsNil(o.AllowedWeekDays) {
		var ret int32
		return ret
	}
	return *o.AllowedWeekDays
}

// GetAllowedWeekDaysOk returns a tuple with the AllowedWeekDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockAuthCreate) GetAllowedWeekDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.AllowedWeekDays) {
		return nil, false
	}
	return o.AllowedWeekDays, true
}

// HasAllowedWeekDays returns a boolean if a field has been set.
func (o *SmartlockAuthCreate) HasAllowedWeekDays() bool {
	if o != nil && !IsNil(o.AllowedWeekDays) {
		return true
	}

	return false
}

// SetAllowedWeekDays gets a reference to the given int32 and assigns it to the AllowedWeekDays field.
func (o *SmartlockAuthCreate) SetAllowedWeekDays(v int32) {
	o.AllowedWeekDays = &v
}

// GetAllowedFromTime returns the AllowedFromTime field value if set, zero value otherwise.
func (o *SmartlockAuthCreate) GetAllowedFromTime() int32 {
	if o == nil || IsNil(o.AllowedFromTime) {
		var ret int32
		return ret
	}
	return *o.AllowedFromTime
}

// GetAllowedFromTimeOk returns a tuple with the AllowedFromTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockAuthCreate) GetAllowedFromTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.AllowedFromTime) {
		return nil, false
	}
	return o.AllowedFromTime, true
}

// HasAllowedFromTime returns a boolean if a field has been set.
func (o *SmartlockAuthCreate) HasAllowedFromTime() bool {
	if o != nil && !IsNil(o.AllowedFromTime) {
		return true
	}

	return false
}

// SetAllowedFromTime gets a reference to the given int32 and assigns it to the AllowedFromTime field.
func (o *SmartlockAuthCreate) SetAllowedFromTime(v int32) {
	o.AllowedFromTime = &v
}

// GetAllowedUntilTime returns the AllowedUntilTime field value if set, zero value otherwise.
func (o *SmartlockAuthCreate) GetAllowedUntilTime() int32 {
	if o == nil || IsNil(o.AllowedUntilTime) {
		var ret int32
		return ret
	}
	return *o.AllowedUntilTime
}

// GetAllowedUntilTimeOk returns a tuple with the AllowedUntilTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockAuthCreate) GetAllowedUntilTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.AllowedUntilTime) {
		return nil, false
	}
	return o.AllowedUntilTime, true
}

// HasAllowedUntilTime returns a boolean if a field has been set.
func (o *SmartlockAuthCreate) HasAllowedUntilTime() bool {
	if o != nil && !IsNil(o.AllowedUntilTime) {
		return true
	}

	return false
}

// SetAllowedUntilTime gets a reference to the given int32 and assigns it to the AllowedUntilTime field.
func (o *SmartlockAuthCreate) SetAllowedUntilTime(v int32) {
	o.AllowedUntilTime = &v
}

// GetAccountUserId returns the AccountUserId field value if set, zero value otherwise.
func (o *SmartlockAuthCreate) GetAccountUserId() int32 {
	if o == nil || IsNil(o.AccountUserId) {
		var ret int32
		return ret
	}
	return *o.AccountUserId
}

// GetAccountUserIdOk returns a tuple with the AccountUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockAuthCreate) GetAccountUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AccountUserId) {
		return nil, false
	}
	return o.AccountUserId, true
}

// HasAccountUserId returns a boolean if a field has been set.
func (o *SmartlockAuthCreate) HasAccountUserId() bool {
	if o != nil && !IsNil(o.AccountUserId) {
		return true
	}

	return false
}

// SetAccountUserId gets a reference to the given int32 and assigns it to the AccountUserId field.
func (o *SmartlockAuthCreate) SetAccountUserId(v int32) {
	o.AccountUserId = &v
}

// GetRemoteAllowed returns the RemoteAllowed field value
func (o *SmartlockAuthCreate) GetRemoteAllowed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RemoteAllowed
}

// GetRemoteAllowedOk returns a tuple with the RemoteAllowed field value
// and a boolean to check if the value has been set.
func (o *SmartlockAuthCreate) GetRemoteAllowedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteAllowed, true
}

// SetRemoteAllowed sets field value
func (o *SmartlockAuthCreate) SetRemoteAllowed(v bool) {
	o.RemoteAllowed = v
}

// GetSmartActionsEnabled returns the SmartActionsEnabled field value if set, zero value otherwise.
func (o *SmartlockAuthCreate) GetSmartActionsEnabled() bool {
	if o == nil || IsNil(o.SmartActionsEnabled) {
		var ret bool
		return ret
	}
	return *o.SmartActionsEnabled
}

// GetSmartActionsEnabledOk returns a tuple with the SmartActionsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockAuthCreate) GetSmartActionsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SmartActionsEnabled) {
		return nil, false
	}
	return o.SmartActionsEnabled, true
}

// HasSmartActionsEnabled returns a boolean if a field has been set.
func (o *SmartlockAuthCreate) HasSmartActionsEnabled() bool {
	if o != nil && !IsNil(o.SmartActionsEnabled) {
		return true
	}

	return false
}

// SetSmartActionsEnabled gets a reference to the given bool and assigns it to the SmartActionsEnabled field.
func (o *SmartlockAuthCreate) SetSmartActionsEnabled(v bool) {
	o.SmartActionsEnabled = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SmartlockAuthCreate) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockAuthCreate) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SmartlockAuthCreate) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *SmartlockAuthCreate) SetType(v int32) {
	o.Type = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SmartlockAuthCreate) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockAuthCreate) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SmartlockAuthCreate) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *SmartlockAuthCreate) SetCode(v int32) {
	o.Code = &v
}

func (o SmartlockAuthCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmartlockAuthCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.AllowedFromDate) {
		toSerialize["allowedFromDate"] = o.AllowedFromDate
	}
	if !IsNil(o.AllowedUntilDate) {
		toSerialize["allowedUntilDate"] = o.AllowedUntilDate
	}
	if !IsNil(o.AllowedWeekDays) {
		toSerialize["allowedWeekDays"] = o.AllowedWeekDays
	}
	if !IsNil(o.AllowedFromTime) {
		toSerialize["allowedFromTime"] = o.AllowedFromTime
	}
	if !IsNil(o.AllowedUntilTime) {
		toSerialize["allowedUntilTime"] = o.AllowedUntilTime
	}
	if !IsNil(o.AccountUserId) {
		toSerialize["accountUserId"] = o.AccountUserId
	}
	toSerialize["remoteAllowed"] = o.RemoteAllowed
	if !IsNil(o.SmartActionsEnabled) {
		toSerialize["smartActionsEnabled"] = o.SmartActionsEnabled
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return toSerialize, nil
}

func (o *SmartlockAuthCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"remoteAllowed",
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

	varSmartlockAuthCreate := _SmartlockAuthCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSmartlockAuthCreate)

	if err != nil {
		return err
	}

	*o = SmartlockAuthCreate(varSmartlockAuthCreate)

	return err
}

type NullableSmartlockAuthCreate struct {
	value *SmartlockAuthCreate
	isSet bool
}

func (v NullableSmartlockAuthCreate) Get() *SmartlockAuthCreate {
	return v.value
}

func (v *NullableSmartlockAuthCreate) Set(val *SmartlockAuthCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartlockAuthCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartlockAuthCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartlockAuthCreate(val *SmartlockAuthCreate) *NullableSmartlockAuthCreate {
	return &NullableSmartlockAuthCreate{value: val, isSet: true}
}

func (v NullableSmartlockAuthCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartlockAuthCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


