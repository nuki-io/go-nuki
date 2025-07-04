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
	"bytes"
	"fmt"
)

// checks if the SmartlockLogOpenerLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmartlockLogOpenerLog{}

// SmartlockLogOpenerLog struct for SmartlockLogOpenerLog
type SmartlockLogOpenerLog struct {
	// Flag indicating if continuous mode was active
	ActiveCm bool `json:"activeCm"`
	// Flag indicating if ring to open was active
	ActiveRto bool `json:"activeRto"`
	// The cause of the activation of ring to open or continuous mode: 0 .. doorbell, 1 .. timecontrol, 2 .. app, 3 .. button, 4 .. fob, 5 .. bridge, 6 .. keypad
	Source int32 `json:"source"`
	// Flag indicating a geo fence induced action
	FlagGeoFence bool `json:"flagGeoFence"`
	// Flag indicating a force induced action
	FlagForce bool `json:"flagForce"`
	// Flag indicating if doorbell suppression was active
	FlagDoorbellSuppression bool `json:"flagDoorbellSuppression"`
}

type _SmartlockLogOpenerLog SmartlockLogOpenerLog

// NewSmartlockLogOpenerLog instantiates a new SmartlockLogOpenerLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartlockLogOpenerLog(activeCm bool, activeRto bool, source int32, flagGeoFence bool, flagForce bool, flagDoorbellSuppression bool) *SmartlockLogOpenerLog {
	this := SmartlockLogOpenerLog{}
	this.ActiveCm = activeCm
	this.ActiveRto = activeRto
	this.Source = source
	this.FlagGeoFence = flagGeoFence
	this.FlagForce = flagForce
	this.FlagDoorbellSuppression = flagDoorbellSuppression
	return &this
}

// NewSmartlockLogOpenerLogWithDefaults instantiates a new SmartlockLogOpenerLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartlockLogOpenerLogWithDefaults() *SmartlockLogOpenerLog {
	this := SmartlockLogOpenerLog{}
	return &this
}

// GetActiveCm returns the ActiveCm field value
func (o *SmartlockLogOpenerLog) GetActiveCm() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ActiveCm
}

// GetActiveCmOk returns a tuple with the ActiveCm field value
// and a boolean to check if the value has been set.
func (o *SmartlockLogOpenerLog) GetActiveCmOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveCm, true
}

// SetActiveCm sets field value
func (o *SmartlockLogOpenerLog) SetActiveCm(v bool) {
	o.ActiveCm = v
}

// GetActiveRto returns the ActiveRto field value
func (o *SmartlockLogOpenerLog) GetActiveRto() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ActiveRto
}

// GetActiveRtoOk returns a tuple with the ActiveRto field value
// and a boolean to check if the value has been set.
func (o *SmartlockLogOpenerLog) GetActiveRtoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveRto, true
}

// SetActiveRto sets field value
func (o *SmartlockLogOpenerLog) SetActiveRto(v bool) {
	o.ActiveRto = v
}

// GetSource returns the Source field value
func (o *SmartlockLogOpenerLog) GetSource() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *SmartlockLogOpenerLog) GetSourceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *SmartlockLogOpenerLog) SetSource(v int32) {
	o.Source = v
}

// GetFlagGeoFence returns the FlagGeoFence field value
func (o *SmartlockLogOpenerLog) GetFlagGeoFence() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FlagGeoFence
}

// GetFlagGeoFenceOk returns a tuple with the FlagGeoFence field value
// and a boolean to check if the value has been set.
func (o *SmartlockLogOpenerLog) GetFlagGeoFenceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlagGeoFence, true
}

// SetFlagGeoFence sets field value
func (o *SmartlockLogOpenerLog) SetFlagGeoFence(v bool) {
	o.FlagGeoFence = v
}

// GetFlagForce returns the FlagForce field value
func (o *SmartlockLogOpenerLog) GetFlagForce() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FlagForce
}

// GetFlagForceOk returns a tuple with the FlagForce field value
// and a boolean to check if the value has been set.
func (o *SmartlockLogOpenerLog) GetFlagForceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlagForce, true
}

// SetFlagForce sets field value
func (o *SmartlockLogOpenerLog) SetFlagForce(v bool) {
	o.FlagForce = v
}

// GetFlagDoorbellSuppression returns the FlagDoorbellSuppression field value
func (o *SmartlockLogOpenerLog) GetFlagDoorbellSuppression() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FlagDoorbellSuppression
}

// GetFlagDoorbellSuppressionOk returns a tuple with the FlagDoorbellSuppression field value
// and a boolean to check if the value has been set.
func (o *SmartlockLogOpenerLog) GetFlagDoorbellSuppressionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FlagDoorbellSuppression, true
}

// SetFlagDoorbellSuppression sets field value
func (o *SmartlockLogOpenerLog) SetFlagDoorbellSuppression(v bool) {
	o.FlagDoorbellSuppression = v
}

func (o SmartlockLogOpenerLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmartlockLogOpenerLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["activeCm"] = o.ActiveCm
	toSerialize["activeRto"] = o.ActiveRto
	toSerialize["source"] = o.Source
	toSerialize["flagGeoFence"] = o.FlagGeoFence
	toSerialize["flagForce"] = o.FlagForce
	toSerialize["flagDoorbellSuppression"] = o.FlagDoorbellSuppression
	return toSerialize, nil
}

func (o *SmartlockLogOpenerLog) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"activeCm",
		"activeRto",
		"source",
		"flagGeoFence",
		"flagForce",
		"flagDoorbellSuppression",
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

	varSmartlockLogOpenerLog := _SmartlockLogOpenerLog{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSmartlockLogOpenerLog)

	if err != nil {
		return err
	}

	*o = SmartlockLogOpenerLog(varSmartlockLogOpenerLog)

	return err
}

type NullableSmartlockLogOpenerLog struct {
	value *SmartlockLogOpenerLog
	isSet bool
}

func (v NullableSmartlockLogOpenerLog) Get() *SmartlockLogOpenerLog {
	return v.value
}

func (v *NullableSmartlockLogOpenerLog) Set(val *SmartlockLogOpenerLog) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartlockLogOpenerLog) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartlockLogOpenerLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartlockLogOpenerLog(val *SmartlockLogOpenerLog) *NullableSmartlockLogOpenerLog {
	return &NullableSmartlockLogOpenerLog{value: val, isSet: true}
}

func (v NullableSmartlockLogOpenerLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartlockLogOpenerLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


