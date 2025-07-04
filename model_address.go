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

// checks if the Address type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Address{}

// Address struct for Address
type Address struct {
	// The address id
	AddressId int32 `json:"addressId"`
	// The account id
	AccountId int32 `json:"accountId"`
	// The name of the address
	Name string `json:"name"`
	// The smartlocks for this address
	SmartlockIds []int64 `json:"smartlockIds"`
	// The optional service id if the address is from an partner service
	ServiceId *string `json:"serviceId,omitempty"`
	// The timezone
	TimeZone *string `json:"timeZone,omitempty"`
	// The optional check in time (minutes of the day)
	CheckInTime *int32 `json:"checkInTime,omitempty"`
	// The optional check out time (minutes of the day)
	CheckOutTime *int32 `json:"checkOutTime,omitempty"`
	// The optional settings object
	Settings map[string]map[string]interface{} `json:"settings,omitempty"`
	// The creation date
	CreationDate time.Time `json:"creationDate"`
	// The update date
	UpdateDate time.Time `json:"updateDate"`
}

type _Address Address

// NewAddress instantiates a new Address object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress(addressId int32, accountId int32, name string, smartlockIds []int64, creationDate time.Time, updateDate time.Time) *Address {
	this := Address{}
	this.AddressId = addressId
	this.AccountId = accountId
	this.Name = name
	this.SmartlockIds = smartlockIds
	this.CreationDate = creationDate
	this.UpdateDate = updateDate
	return &this
}

// NewAddressWithDefaults instantiates a new Address object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressWithDefaults() *Address {
	this := Address{}
	return &this
}

// GetAddressId returns the AddressId field value
func (o *Address) GetAddressId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AddressId
}

// GetAddressIdOk returns a tuple with the AddressId field value
// and a boolean to check if the value has been set.
func (o *Address) GetAddressIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressId, true
}

// SetAddressId sets field value
func (o *Address) SetAddressId(v int32) {
	o.AddressId = v
}

// GetAccountId returns the AccountId field value
func (o *Address) GetAccountId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *Address) GetAccountIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *Address) SetAccountId(v int32) {
	o.AccountId = v
}

// GetName returns the Name field value
func (o *Address) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Address) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Address) SetName(v string) {
	o.Name = v
}

// GetSmartlockIds returns the SmartlockIds field value
func (o *Address) GetSmartlockIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.SmartlockIds
}

// GetSmartlockIdsOk returns a tuple with the SmartlockIds field value
// and a boolean to check if the value has been set.
func (o *Address) GetSmartlockIdsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SmartlockIds, true
}

// SetSmartlockIds sets field value
func (o *Address) SetSmartlockIds(v []int64) {
	o.SmartlockIds = v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *Address) GetServiceId() string {
	if o == nil || IsNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *Address) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *Address) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *Address) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *Address) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *Address) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetCheckInTime returns the CheckInTime field value if set, zero value otherwise.
func (o *Address) GetCheckInTime() int32 {
	if o == nil || IsNil(o.CheckInTime) {
		var ret int32
		return ret
	}
	return *o.CheckInTime
}

// GetCheckInTimeOk returns a tuple with the CheckInTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetCheckInTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.CheckInTime) {
		return nil, false
	}
	return o.CheckInTime, true
}

// HasCheckInTime returns a boolean if a field has been set.
func (o *Address) HasCheckInTime() bool {
	if o != nil && !IsNil(o.CheckInTime) {
		return true
	}

	return false
}

// SetCheckInTime gets a reference to the given int32 and assigns it to the CheckInTime field.
func (o *Address) SetCheckInTime(v int32) {
	o.CheckInTime = &v
}

// GetCheckOutTime returns the CheckOutTime field value if set, zero value otherwise.
func (o *Address) GetCheckOutTime() int32 {
	if o == nil || IsNil(o.CheckOutTime) {
		var ret int32
		return ret
	}
	return *o.CheckOutTime
}

// GetCheckOutTimeOk returns a tuple with the CheckOutTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetCheckOutTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.CheckOutTime) {
		return nil, false
	}
	return o.CheckOutTime, true
}

// HasCheckOutTime returns a boolean if a field has been set.
func (o *Address) HasCheckOutTime() bool {
	if o != nil && !IsNil(o.CheckOutTime) {
		return true
	}

	return false
}

// SetCheckOutTime gets a reference to the given int32 and assigns it to the CheckOutTime field.
func (o *Address) SetCheckOutTime(v int32) {
	o.CheckOutTime = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *Address) GetSettings() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Settings) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetSettingsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Settings) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *Address) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given map[string]map[string]interface{} and assigns it to the Settings field.
func (o *Address) SetSettings(v map[string]map[string]interface{}) {
	o.Settings = v
}

// GetCreationDate returns the CreationDate field value
func (o *Address) GetCreationDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value
// and a boolean to check if the value has been set.
func (o *Address) GetCreationDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationDate, true
}

// SetCreationDate sets field value
func (o *Address) SetCreationDate(v time.Time) {
	o.CreationDate = v
}

// GetUpdateDate returns the UpdateDate field value
func (o *Address) GetUpdateDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdateDate
}

// GetUpdateDateOk returns a tuple with the UpdateDate field value
// and a boolean to check if the value has been set.
func (o *Address) GetUpdateDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateDate, true
}

// SetUpdateDate sets field value
func (o *Address) SetUpdateDate(v time.Time) {
	o.UpdateDate = v
}

func (o Address) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Address) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addressId"] = o.AddressId
	toSerialize["accountId"] = o.AccountId
	toSerialize["name"] = o.Name
	toSerialize["smartlockIds"] = o.SmartlockIds
	if !IsNil(o.ServiceId) {
		toSerialize["serviceId"] = o.ServiceId
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !IsNil(o.CheckInTime) {
		toSerialize["checkInTime"] = o.CheckInTime
	}
	if !IsNil(o.CheckOutTime) {
		toSerialize["checkOutTime"] = o.CheckOutTime
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	toSerialize["creationDate"] = o.CreationDate
	toSerialize["updateDate"] = o.UpdateDate
	return toSerialize, nil
}

func (o *Address) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"addressId",
		"accountId",
		"name",
		"smartlockIds",
		"creationDate",
		"updateDate",
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

	varAddress := _Address{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddress)

	if err != nil {
		return err
	}

	*o = Address(varAddress)

	return err
}

type NullableAddress struct {
	value *Address
	isSet bool
}

func (v NullableAddress) Get() *Address {
	return v.value
}

func (v *NullableAddress) Set(val *Address) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress(val *Address) *NullableAddress {
	return &NullableAddress{value: val, isSet: true}
}

func (v NullableAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


