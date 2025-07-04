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

// checks if the AccountUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountUser{}

// AccountUser struct for AccountUser
type AccountUser struct {
	// The account user id
	AccountUserId int32 `json:"accountUserId"`
	// The account id
	AccountId int32 `json:"accountId"`
	// The optional type: 0 .. user, 1 .. company
	Type *int32 `json:"type,omitempty"`
	// The email address
	Email string `json:"email"`
	// The name
	Name string `json:"name"`
	// The language code
	Language *string `json:"language,omitempty"`
	// The operation id - if set it's locked for another operation
	OperationId *string `json:"operationId,omitempty"`
	// The creation date
	CreationDate time.Time `json:"creationDate"`
	// The update date
	UpdateDate time.Time `json:"updateDate"`
}

type _AccountUser AccountUser

// NewAccountUser instantiates a new AccountUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUser(accountUserId int32, accountId int32, email string, name string, creationDate time.Time, updateDate time.Time) *AccountUser {
	this := AccountUser{}
	this.AccountUserId = accountUserId
	this.AccountId = accountId
	this.Email = email
	this.Name = name
	this.CreationDate = creationDate
	this.UpdateDate = updateDate
	return &this
}

// NewAccountUserWithDefaults instantiates a new AccountUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUserWithDefaults() *AccountUser {
	this := AccountUser{}
	return &this
}

// GetAccountUserId returns the AccountUserId field value
func (o *AccountUser) GetAccountUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountUserId
}

// GetAccountUserIdOk returns a tuple with the AccountUserId field value
// and a boolean to check if the value has been set.
func (o *AccountUser) GetAccountUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountUserId, true
}

// SetAccountUserId sets field value
func (o *AccountUser) SetAccountUserId(v int32) {
	o.AccountUserId = v
}

// GetAccountId returns the AccountId field value
func (o *AccountUser) GetAccountId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AccountUser) GetAccountIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AccountUser) SetAccountId(v int32) {
	o.AccountId = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccountUser) GetType() int32 {
	if o == nil || IsNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUser) GetTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccountUser) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *AccountUser) SetType(v int32) {
	o.Type = &v
}

// GetEmail returns the Email field value
func (o *AccountUser) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *AccountUser) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *AccountUser) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value
func (o *AccountUser) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountUser) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountUser) SetName(v string) {
	o.Name = v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *AccountUser) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUser) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *AccountUser) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *AccountUser) SetLanguage(v string) {
	o.Language = &v
}

// GetOperationId returns the OperationId field value if set, zero value otherwise.
func (o *AccountUser) GetOperationId() string {
	if o == nil || IsNil(o.OperationId) {
		var ret string
		return ret
	}
	return *o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUser) GetOperationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OperationId) {
		return nil, false
	}
	return o.OperationId, true
}

// HasOperationId returns a boolean if a field has been set.
func (o *AccountUser) HasOperationId() bool {
	if o != nil && !IsNil(o.OperationId) {
		return true
	}

	return false
}

// SetOperationId gets a reference to the given string and assigns it to the OperationId field.
func (o *AccountUser) SetOperationId(v string) {
	o.OperationId = &v
}

// GetCreationDate returns the CreationDate field value
func (o *AccountUser) GetCreationDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value
// and a boolean to check if the value has been set.
func (o *AccountUser) GetCreationDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreationDate, true
}

// SetCreationDate sets field value
func (o *AccountUser) SetCreationDate(v time.Time) {
	o.CreationDate = v
}

// GetUpdateDate returns the UpdateDate field value
func (o *AccountUser) GetUpdateDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdateDate
}

// GetUpdateDateOk returns a tuple with the UpdateDate field value
// and a boolean to check if the value has been set.
func (o *AccountUser) GetUpdateDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateDate, true
}

// SetUpdateDate sets field value
func (o *AccountUser) SetUpdateDate(v time.Time) {
	o.UpdateDate = v
}

func (o AccountUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountUserId"] = o.AccountUserId
	toSerialize["accountId"] = o.AccountId
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["email"] = o.Email
	toSerialize["name"] = o.Name
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.OperationId) {
		toSerialize["operationId"] = o.OperationId
	}
	toSerialize["creationDate"] = o.CreationDate
	toSerialize["updateDate"] = o.UpdateDate
	return toSerialize, nil
}

func (o *AccountUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accountUserId",
		"accountId",
		"email",
		"name",
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

	varAccountUser := _AccountUser{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountUser)

	if err != nil {
		return err
	}

	*o = AccountUser(varAccountUser)

	return err
}

type NullableAccountUser struct {
	value *AccountUser
	isSet bool
}

func (v NullableAccountUser) Get() *AccountUser {
	return v.value
}

func (v *NullableAccountUser) Set(val *AccountUser) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUser) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUser(val *AccountUser) *NullableAccountUser {
	return &NullableAccountUser{value: val, isSet: true}
}

func (v NullableAccountUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


