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

// checks if the ChallengeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChallengeResponse{}

// ChallengeResponse struct for ChallengeResponse
type ChallengeResponse struct {
	RawValue *string `json:"rawValue,omitempty"`
	Parameters []Parameter `json:"parameters,omitempty"`
	Scheme *ChallengeScheme `json:"scheme,omitempty"`
	ServerNonce *string `json:"serverNonce,omitempty"`
	Realm *string `json:"realm,omitempty"`
	Opaque *string `json:"opaque,omitempty"`
	DigestAlgorithm *string `json:"digestAlgorithm,omitempty"`
	ClientNonce *string `json:"clientNonce,omitempty"`
	DigestRef *Reference `json:"digestRef,omitempty"`
	Identifier *string `json:"identifier,omitempty"`
	Quality *string `json:"quality,omitempty"`
	Secret []string `json:"secret,omitempty"`
	SecretAlgorithm *string `json:"secretAlgorithm,omitempty"`
	ServerNounceCount *int32 `json:"serverNounceCount,omitempty"`
	TimeIssued *int64 `json:"timeIssued,omitempty"`
	Principal *Principal `json:"principal,omitempty"`
	ServerNounceCountAsHex *string `json:"serverNounceCountAsHex,omitempty"`
}

// NewChallengeResponse instantiates a new ChallengeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChallengeResponse() *ChallengeResponse {
	this := ChallengeResponse{}
	return &this
}

// NewChallengeResponseWithDefaults instantiates a new ChallengeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChallengeResponseWithDefaults() *ChallengeResponse {
	this := ChallengeResponse{}
	return &this
}

// GetRawValue returns the RawValue field value if set, zero value otherwise.
func (o *ChallengeResponse) GetRawValue() string {
	if o == nil || IsNil(o.RawValue) {
		var ret string
		return ret
	}
	return *o.RawValue
}

// GetRawValueOk returns a tuple with the RawValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeResponse) GetRawValueOk() (*string, bool) {
	if o == nil || IsNil(o.RawValue) {
		return nil, false
	}
	return o.RawValue, true
}

// HasRawValue returns a boolean if a field has been set.
func (o *ChallengeResponse) HasRawValue() bool {
	if o != nil && !IsNil(o.RawValue) {
		return true
	}

	return false
}

// SetRawValue gets a reference to the given string and assigns it to the RawValue field.
func (o *ChallengeResponse) SetRawValue(v string) {
	o.RawValue = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ChallengeResponse) GetParameters() []Parameter {
	if o == nil || IsNil(o.Parameters) {
		var ret []Parameter
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeResponse) GetParametersOk() ([]Parameter, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ChallengeResponse) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []Parameter and assigns it to the Parameters field.
func (o *ChallengeResponse) SetParameters(v []Parameter) {
	o.Parameters = v
}

// GetScheme returns the Scheme field value if set, zero value otherwise.
func (o *ChallengeResponse) GetScheme() ChallengeScheme {
	if o == nil || IsNil(o.Scheme) {
		var ret ChallengeScheme
		return ret
	}
	return *o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeResponse) GetSchemeOk() (*ChallengeScheme, bool) {
	if o == nil || IsNil(o.Scheme) {
		return nil, false
	}
	return o.Scheme, true
}

// HasScheme returns a boolean if a field has been set.
func (o *ChallengeResponse) HasScheme() bool {
	if o != nil && !IsNil(o.Scheme) {
		return true
	}

	return false
}

// SetScheme gets a reference to the given ChallengeScheme and assigns it to the Scheme field.
func (o *ChallengeResponse) SetScheme(v ChallengeScheme) {
	o.Scheme = &v
}

// GetServerNonce returns the ServerNonce field value if set, zero value otherwise.
func (o *ChallengeResponse) GetServerNonce() string {
	if o == nil || IsNil(o.ServerNonce) {
		var ret string
		return ret
	}
	return *o.ServerNonce
}

// GetServerNonceOk returns a tuple with the ServerNonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeResponse) GetServerNonceOk() (*string, bool) {
	if o == nil || IsNil(o.ServerNonce) {
		return nil, false
	}
	return o.ServerNonce, true
}

// HasServerNonce returns a boolean if a field has been set.
func (o *ChallengeResponse) HasServerNonce() bool {
	if o != nil && !IsNil(o.ServerNonce) {
		return true
	}

	return false
}

// SetServerNonce gets a reference to the given string and assigns it to the ServerNonce field.
func (o *ChallengeResponse) SetServerNonce(v string) {
	o.ServerNonce = &v
}

// GetRealm returns the Realm field value if set, zero value otherwise.
func (o *ChallengeResponse) GetRealm() string {
	if o == nil || IsNil(o.Realm) {
		var ret string
		return ret
	}
	return *o.Realm
}

// GetRealmOk returns a tuple with the Realm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeResponse) GetRealmOk() (*string, bool) {
	if o == nil || IsNil(o.Realm) {
		return nil, false
	}
	return o.Realm, true
}

// HasRealm returns a boolean if a field has been set.
func (o *ChallengeResponse) HasRealm() bool {
	if o != nil && !IsNil(o.Realm) {
		return true
	}

	return false
}

// SetRealm gets a reference to the given string and assigns it to the Realm field.
func (o *ChallengeResponse) SetRealm(v string) {
	o.Realm = &v
}

// GetOpaque returns the Opaque field value if set, zero value otherwise.
func (o *ChallengeResponse) GetOpaque() string {
	if o == nil || IsNil(o.Opaque) {
		var ret string
		return ret
	}
	return *o.Opaque
}

// GetOpaqueOk returns a tuple with the Opaque field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeResponse) GetOpaqueOk() (*string, bool) {
	if o == nil || IsNil(o.Opaque) {
		return nil, false
	}
	return o.Opaque, true
}

// HasOpaque returns a boolean if a field has been set.
func (o *ChallengeResponse) HasOpaque() bool {
	if o != nil && !IsNil(o.Opaque) {
		return true
	}

	return false
}

// SetOpaque gets a reference to the given string and assigns it to the Opaque field.
func (o *ChallengeResponse) SetOpaque(v string) {
	o.Opaque = &v
}

// GetDigestAlgorithm returns the DigestAlgorithm field value if set, zero value otherwise.
func (o *ChallengeResponse) GetDigestAlgorithm() string {
	if o == nil || IsNil(o.DigestAlgorithm) {
		var ret string
		return ret
	}
	return *o.DigestAlgorithm
}

// GetDigestAlgorithmOk returns a tuple with the DigestAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeResponse) GetDigestAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.DigestAlgorithm) {
		return nil, false
	}
	return o.DigestAlgorithm, true
}

// HasDigestAlgorithm returns a boolean if a field has been set.
func (o *ChallengeResponse) HasDigestAlgorithm() bool {
	if o != nil && !IsNil(o.DigestAlgorithm) {
		return true
	}

	return false
}

// SetDigestAlgorithm gets a reference to the given string and assigns it to the DigestAlgorithm field.
func (o *ChallengeResponse) SetDigestAlgorithm(v string) {
	o.DigestAlgorithm = &v
}

// GetClientNonce returns the ClientNonce field value if set, zero value otherwise.
func (o *ChallengeResponse) GetClientNonce() string {
	if o == nil || IsNil(o.ClientNonce) {
		var ret string
		return ret
	}
	return *o.ClientNonce
}

// GetClientNonceOk returns a tuple with the ClientNonce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeResponse) GetClientNonceOk() (*string, bool) {
	if o == nil || IsNil(o.ClientNonce) {
		return nil, false
	}
	return o.ClientNonce, true
}

// HasClientNonce returns a boolean if a field has been set.
func (o *ChallengeResponse) HasClientNonce() bool {
	if o != nil && !IsNil(o.ClientNonce) {
		return true
	}

	return false
}

// SetClientNonce gets a reference to the given string and assigns it to the ClientNonce field.
func (o *ChallengeResponse) SetClientNonce(v string) {
	o.ClientNonce = &v
}

// GetDigestRef returns the DigestRef field value if set, zero value otherwise.
func (o *ChallengeResponse) GetDigestRef() Reference {
	if o == nil || IsNil(o.DigestRef) {
		var ret Reference
		return ret
	}
	return *o.DigestRef
}

// GetDigestRefOk returns a tuple with the DigestRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeResponse) GetDigestRefOk() (*Reference, bool) {
	if o == nil || IsNil(o.DigestRef) {
		return nil, false
	}
	return o.DigestRef, true
}

// HasDigestRef returns a boolean if a field has been set.
func (o *ChallengeResponse) HasDigestRef() bool {
	if o != nil && !IsNil(o.DigestRef) {
		return true
	}

	return false
}

// SetDigestRef gets a reference to the given Reference and assigns it to the DigestRef field.
func (o *ChallengeResponse) SetDigestRef(v Reference) {
	o.DigestRef = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *ChallengeResponse) GetIdentifier() string {
	if o == nil || IsNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeResponse) GetIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *ChallengeResponse) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *ChallengeResponse) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *ChallengeResponse) GetQuality() string {
	if o == nil || IsNil(o.Quality) {
		var ret string
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeResponse) GetQualityOk() (*string, bool) {
	if o == nil || IsNil(o.Quality) {
		return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *ChallengeResponse) HasQuality() bool {
	if o != nil && !IsNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given string and assigns it to the Quality field.
func (o *ChallengeResponse) SetQuality(v string) {
	o.Quality = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ChallengeResponse) GetSecret() []string {
	if o == nil || IsNil(o.Secret) {
		var ret []string
		return ret
	}
	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeResponse) GetSecretOk() ([]string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ChallengeResponse) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given []string and assigns it to the Secret field.
func (o *ChallengeResponse) SetSecret(v []string) {
	o.Secret = v
}

// GetSecretAlgorithm returns the SecretAlgorithm field value if set, zero value otherwise.
func (o *ChallengeResponse) GetSecretAlgorithm() string {
	if o == nil || IsNil(o.SecretAlgorithm) {
		var ret string
		return ret
	}
	return *o.SecretAlgorithm
}

// GetSecretAlgorithmOk returns a tuple with the SecretAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeResponse) GetSecretAlgorithmOk() (*string, bool) {
	if o == nil || IsNil(o.SecretAlgorithm) {
		return nil, false
	}
	return o.SecretAlgorithm, true
}

// HasSecretAlgorithm returns a boolean if a field has been set.
func (o *ChallengeResponse) HasSecretAlgorithm() bool {
	if o != nil && !IsNil(o.SecretAlgorithm) {
		return true
	}

	return false
}

// SetSecretAlgorithm gets a reference to the given string and assigns it to the SecretAlgorithm field.
func (o *ChallengeResponse) SetSecretAlgorithm(v string) {
	o.SecretAlgorithm = &v
}

// GetServerNounceCount returns the ServerNounceCount field value if set, zero value otherwise.
func (o *ChallengeResponse) GetServerNounceCount() int32 {
	if o == nil || IsNil(o.ServerNounceCount) {
		var ret int32
		return ret
	}
	return *o.ServerNounceCount
}

// GetServerNounceCountOk returns a tuple with the ServerNounceCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeResponse) GetServerNounceCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ServerNounceCount) {
		return nil, false
	}
	return o.ServerNounceCount, true
}

// HasServerNounceCount returns a boolean if a field has been set.
func (o *ChallengeResponse) HasServerNounceCount() bool {
	if o != nil && !IsNil(o.ServerNounceCount) {
		return true
	}

	return false
}

// SetServerNounceCount gets a reference to the given int32 and assigns it to the ServerNounceCount field.
func (o *ChallengeResponse) SetServerNounceCount(v int32) {
	o.ServerNounceCount = &v
}

// GetTimeIssued returns the TimeIssued field value if set, zero value otherwise.
func (o *ChallengeResponse) GetTimeIssued() int64 {
	if o == nil || IsNil(o.TimeIssued) {
		var ret int64
		return ret
	}
	return *o.TimeIssued
}

// GetTimeIssuedOk returns a tuple with the TimeIssued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeResponse) GetTimeIssuedOk() (*int64, bool) {
	if o == nil || IsNil(o.TimeIssued) {
		return nil, false
	}
	return o.TimeIssued, true
}

// HasTimeIssued returns a boolean if a field has been set.
func (o *ChallengeResponse) HasTimeIssued() bool {
	if o != nil && !IsNil(o.TimeIssued) {
		return true
	}

	return false
}

// SetTimeIssued gets a reference to the given int64 and assigns it to the TimeIssued field.
func (o *ChallengeResponse) SetTimeIssued(v int64) {
	o.TimeIssued = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *ChallengeResponse) GetPrincipal() Principal {
	if o == nil || IsNil(o.Principal) {
		var ret Principal
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeResponse) GetPrincipalOk() (*Principal, bool) {
	if o == nil || IsNil(o.Principal) {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *ChallengeResponse) HasPrincipal() bool {
	if o != nil && !IsNil(o.Principal) {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given Principal and assigns it to the Principal field.
func (o *ChallengeResponse) SetPrincipal(v Principal) {
	o.Principal = &v
}

// GetServerNounceCountAsHex returns the ServerNounceCountAsHex field value if set, zero value otherwise.
func (o *ChallengeResponse) GetServerNounceCountAsHex() string {
	if o == nil || IsNil(o.ServerNounceCountAsHex) {
		var ret string
		return ret
	}
	return *o.ServerNounceCountAsHex
}

// GetServerNounceCountAsHexOk returns a tuple with the ServerNounceCountAsHex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeResponse) GetServerNounceCountAsHexOk() (*string, bool) {
	if o == nil || IsNil(o.ServerNounceCountAsHex) {
		return nil, false
	}
	return o.ServerNounceCountAsHex, true
}

// HasServerNounceCountAsHex returns a boolean if a field has been set.
func (o *ChallengeResponse) HasServerNounceCountAsHex() bool {
	if o != nil && !IsNil(o.ServerNounceCountAsHex) {
		return true
	}

	return false
}

// SetServerNounceCountAsHex gets a reference to the given string and assigns it to the ServerNounceCountAsHex field.
func (o *ChallengeResponse) SetServerNounceCountAsHex(v string) {
	o.ServerNounceCountAsHex = &v
}

func (o ChallengeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChallengeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RawValue) {
		toSerialize["rawValue"] = o.RawValue
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.Scheme) {
		toSerialize["scheme"] = o.Scheme
	}
	if !IsNil(o.ServerNonce) {
		toSerialize["serverNonce"] = o.ServerNonce
	}
	if !IsNil(o.Realm) {
		toSerialize["realm"] = o.Realm
	}
	if !IsNil(o.Opaque) {
		toSerialize["opaque"] = o.Opaque
	}
	if !IsNil(o.DigestAlgorithm) {
		toSerialize["digestAlgorithm"] = o.DigestAlgorithm
	}
	if !IsNil(o.ClientNonce) {
		toSerialize["clientNonce"] = o.ClientNonce
	}
	if !IsNil(o.DigestRef) {
		toSerialize["digestRef"] = o.DigestRef
	}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !IsNil(o.SecretAlgorithm) {
		toSerialize["secretAlgorithm"] = o.SecretAlgorithm
	}
	if !IsNil(o.ServerNounceCount) {
		toSerialize["serverNounceCount"] = o.ServerNounceCount
	}
	if !IsNil(o.TimeIssued) {
		toSerialize["timeIssued"] = o.TimeIssued
	}
	if !IsNil(o.Principal) {
		toSerialize["principal"] = o.Principal
	}
	if !IsNil(o.ServerNounceCountAsHex) {
		toSerialize["serverNounceCountAsHex"] = o.ServerNounceCountAsHex
	}
	return toSerialize, nil
}

type NullableChallengeResponse struct {
	value *ChallengeResponse
	isSet bool
}

func (v NullableChallengeResponse) Get() *ChallengeResponse {
	return v.value
}

func (v *NullableChallengeResponse) Set(val *ChallengeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChallengeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChallengeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChallengeResponse(val *ChallengeResponse) *NullableChallengeResponse {
	return &NullableChallengeResponse{value: val, isSet: true}
}

func (v NullableChallengeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChallengeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


