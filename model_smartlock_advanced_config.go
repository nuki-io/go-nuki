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

// checks if the SmartlockAdvancedConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmartlockAdvancedConfig{}

// SmartlockAdvancedConfig struct for SmartlockAdvancedConfig
type SmartlockAdvancedConfig struct {
	// Timeout in seconds for lock ‘n’ go
	LngTimeout *int32 `json:"lngTimeout,omitempty"`
	// The desired action, if the button is pressed once: 0 .. no action, 1 .. intelligent, 2 .. unlock, 3 .. lock, 4 .. unlatch, 5 .. lock 'n' go, 6 .. show status
	SingleButtonPressAction *int32 `json:"singleButtonPressAction,omitempty"`
	// The desired action, if the button is pressed twice: 0 .. no action, 1 .. intelligent, 2 .. unlock, 3 .. lock, 4 .. unlatch, 5 .. lock 'n' go, 6 .. show status
	DoubleButtonPressAction *int32 `json:"doubleButtonPressAction,omitempty"`
	// Flag that indicates if the automatic detection of the battery type is enabled
	AutomaticBatteryTypeDetection *bool `json:"automaticBatteryTypeDetection,omitempty"`
	// Duration in seconds for holding the latch in unlatched position
	UnlatchDuration *int32 `json:"unlatchDuration,omitempty"`
	// The operation id - if set it's locked for another operation
	OperationId *string `json:"operationId,omitempty"`
	// The absolute total position in degrees that has been reached during calibration
	TotalDegrees int32 `json:"totalDegrees"`
	// Offset that alters the single locked position
	SingleLockedPositionOffsetDegrees int32 `json:"singleLockedPositionOffsetDegrees"`
	// Offset that alters the position where transition from unlocked to locked happens
	UnlockedToLockedTransitionOffsetDegrees *int32 `json:"unlockedToLockedTransitionOffsetDegrees,omitempty"`
	// Offset that alters the unlocked position
	UnlockedPositionOffsetDegrees int32 `json:"unlockedPositionOffsetDegrees"`
	// Offset that alters the locked position
	LockedPositionOffsetDegrees int32 `json:"lockedPositionOffsetDegrees"`
	// Flag that indicates that the inner side of the used cylinder is detached from the outer side
	DetachedCylinder *bool `json:"detachedCylinder,omitempty"`
	// The type of the batteries present in the smart lock: 0 .. alkali, 1 .. accumulator, 2 .. lithium
	BatteryType int32 `json:"batteryType"`
	// New separate flag with FW >= 2.7.8/1.9.1: The Auto Lock feature automatically locks your door when it has been unlocked for a certain period of time
	AutoLock *bool `json:"autoLock,omitempty"`
	// Seconds until the smart lock relocks itself after it has been unlocked. FW < 2.7.8/1.9.1: No auto relock if value is 0, FW >= 2.7.8/1.9.1: has to be >=2 (defaults to 2 for values <2 if autoLock is set to true)
	AutoLockTimeout *int32 `json:"autoLockTimeout,omitempty"`
	// Flag that indicates if available firmware updates for the deviceshould be installed automatically
	AutoUpdateEnabled *bool `json:"autoUpdateEnabled,omitempty"`
	// Field used for setting the motor speed. 0x00 ... standard, 0x01 ... fast, 0x02 ... slow
	MotorSpeed *int32 `json:"motorSpeed,omitempty"`
	// Flag indicating if the slow speed shall be applied during NightMode
	EnableSlowSpeedDuringNightmode *bool `json:"enableSlowSpeedDuringNightmode,omitempty"`
}

type _SmartlockAdvancedConfig SmartlockAdvancedConfig

// NewSmartlockAdvancedConfig instantiates a new SmartlockAdvancedConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmartlockAdvancedConfig(totalDegrees int32, singleLockedPositionOffsetDegrees int32, unlockedPositionOffsetDegrees int32, lockedPositionOffsetDegrees int32, batteryType int32) *SmartlockAdvancedConfig {
	this := SmartlockAdvancedConfig{}
	this.TotalDegrees = totalDegrees
	this.SingleLockedPositionOffsetDegrees = singleLockedPositionOffsetDegrees
	this.UnlockedPositionOffsetDegrees = unlockedPositionOffsetDegrees
	this.LockedPositionOffsetDegrees = lockedPositionOffsetDegrees
	this.BatteryType = batteryType
	return &this
}

// NewSmartlockAdvancedConfigWithDefaults instantiates a new SmartlockAdvancedConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmartlockAdvancedConfigWithDefaults() *SmartlockAdvancedConfig {
	this := SmartlockAdvancedConfig{}
	return &this
}

// GetLngTimeout returns the LngTimeout field value if set, zero value otherwise.
func (o *SmartlockAdvancedConfig) GetLngTimeout() int32 {
	if o == nil || IsNil(o.LngTimeout) {
		var ret int32
		return ret
	}
	return *o.LngTimeout
}

// GetLngTimeoutOk returns a tuple with the LngTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockAdvancedConfig) GetLngTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.LngTimeout) {
		return nil, false
	}
	return o.LngTimeout, true
}

// HasLngTimeout returns a boolean if a field has been set.
func (o *SmartlockAdvancedConfig) HasLngTimeout() bool {
	if o != nil && !IsNil(o.LngTimeout) {
		return true
	}

	return false
}

// SetLngTimeout gets a reference to the given int32 and assigns it to the LngTimeout field.
func (o *SmartlockAdvancedConfig) SetLngTimeout(v int32) {
	o.LngTimeout = &v
}

// GetSingleButtonPressAction returns the SingleButtonPressAction field value if set, zero value otherwise.
func (o *SmartlockAdvancedConfig) GetSingleButtonPressAction() int32 {
	if o == nil || IsNil(o.SingleButtonPressAction) {
		var ret int32
		return ret
	}
	return *o.SingleButtonPressAction
}

// GetSingleButtonPressActionOk returns a tuple with the SingleButtonPressAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockAdvancedConfig) GetSingleButtonPressActionOk() (*int32, bool) {
	if o == nil || IsNil(o.SingleButtonPressAction) {
		return nil, false
	}
	return o.SingleButtonPressAction, true
}

// HasSingleButtonPressAction returns a boolean if a field has been set.
func (o *SmartlockAdvancedConfig) HasSingleButtonPressAction() bool {
	if o != nil && !IsNil(o.SingleButtonPressAction) {
		return true
	}

	return false
}

// SetSingleButtonPressAction gets a reference to the given int32 and assigns it to the SingleButtonPressAction field.
func (o *SmartlockAdvancedConfig) SetSingleButtonPressAction(v int32) {
	o.SingleButtonPressAction = &v
}

// GetDoubleButtonPressAction returns the DoubleButtonPressAction field value if set, zero value otherwise.
func (o *SmartlockAdvancedConfig) GetDoubleButtonPressAction() int32 {
	if o == nil || IsNil(o.DoubleButtonPressAction) {
		var ret int32
		return ret
	}
	return *o.DoubleButtonPressAction
}

// GetDoubleButtonPressActionOk returns a tuple with the DoubleButtonPressAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockAdvancedConfig) GetDoubleButtonPressActionOk() (*int32, bool) {
	if o == nil || IsNil(o.DoubleButtonPressAction) {
		return nil, false
	}
	return o.DoubleButtonPressAction, true
}

// HasDoubleButtonPressAction returns a boolean if a field has been set.
func (o *SmartlockAdvancedConfig) HasDoubleButtonPressAction() bool {
	if o != nil && !IsNil(o.DoubleButtonPressAction) {
		return true
	}

	return false
}

// SetDoubleButtonPressAction gets a reference to the given int32 and assigns it to the DoubleButtonPressAction field.
func (o *SmartlockAdvancedConfig) SetDoubleButtonPressAction(v int32) {
	o.DoubleButtonPressAction = &v
}

// GetAutomaticBatteryTypeDetection returns the AutomaticBatteryTypeDetection field value if set, zero value otherwise.
func (o *SmartlockAdvancedConfig) GetAutomaticBatteryTypeDetection() bool {
	if o == nil || IsNil(o.AutomaticBatteryTypeDetection) {
		var ret bool
		return ret
	}
	return *o.AutomaticBatteryTypeDetection
}

// GetAutomaticBatteryTypeDetectionOk returns a tuple with the AutomaticBatteryTypeDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockAdvancedConfig) GetAutomaticBatteryTypeDetectionOk() (*bool, bool) {
	if o == nil || IsNil(o.AutomaticBatteryTypeDetection) {
		return nil, false
	}
	return o.AutomaticBatteryTypeDetection, true
}

// HasAutomaticBatteryTypeDetection returns a boolean if a field has been set.
func (o *SmartlockAdvancedConfig) HasAutomaticBatteryTypeDetection() bool {
	if o != nil && !IsNil(o.AutomaticBatteryTypeDetection) {
		return true
	}

	return false
}

// SetAutomaticBatteryTypeDetection gets a reference to the given bool and assigns it to the AutomaticBatteryTypeDetection field.
func (o *SmartlockAdvancedConfig) SetAutomaticBatteryTypeDetection(v bool) {
	o.AutomaticBatteryTypeDetection = &v
}

// GetUnlatchDuration returns the UnlatchDuration field value if set, zero value otherwise.
func (o *SmartlockAdvancedConfig) GetUnlatchDuration() int32 {
	if o == nil || IsNil(o.UnlatchDuration) {
		var ret int32
		return ret
	}
	return *o.UnlatchDuration
}

// GetUnlatchDurationOk returns a tuple with the UnlatchDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockAdvancedConfig) GetUnlatchDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.UnlatchDuration) {
		return nil, false
	}
	return o.UnlatchDuration, true
}

// HasUnlatchDuration returns a boolean if a field has been set.
func (o *SmartlockAdvancedConfig) HasUnlatchDuration() bool {
	if o != nil && !IsNil(o.UnlatchDuration) {
		return true
	}

	return false
}

// SetUnlatchDuration gets a reference to the given int32 and assigns it to the UnlatchDuration field.
func (o *SmartlockAdvancedConfig) SetUnlatchDuration(v int32) {
	o.UnlatchDuration = &v
}

// GetOperationId returns the OperationId field value if set, zero value otherwise.
func (o *SmartlockAdvancedConfig) GetOperationId() string {
	if o == nil || IsNil(o.OperationId) {
		var ret string
		return ret
	}
	return *o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockAdvancedConfig) GetOperationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OperationId) {
		return nil, false
	}
	return o.OperationId, true
}

// HasOperationId returns a boolean if a field has been set.
func (o *SmartlockAdvancedConfig) HasOperationId() bool {
	if o != nil && !IsNil(o.OperationId) {
		return true
	}

	return false
}

// SetOperationId gets a reference to the given string and assigns it to the OperationId field.
func (o *SmartlockAdvancedConfig) SetOperationId(v string) {
	o.OperationId = &v
}

// GetTotalDegrees returns the TotalDegrees field value
func (o *SmartlockAdvancedConfig) GetTotalDegrees() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalDegrees
}

// GetTotalDegreesOk returns a tuple with the TotalDegrees field value
// and a boolean to check if the value has been set.
func (o *SmartlockAdvancedConfig) GetTotalDegreesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalDegrees, true
}

// SetTotalDegrees sets field value
func (o *SmartlockAdvancedConfig) SetTotalDegrees(v int32) {
	o.TotalDegrees = v
}

// GetSingleLockedPositionOffsetDegrees returns the SingleLockedPositionOffsetDegrees field value
func (o *SmartlockAdvancedConfig) GetSingleLockedPositionOffsetDegrees() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SingleLockedPositionOffsetDegrees
}

// GetSingleLockedPositionOffsetDegreesOk returns a tuple with the SingleLockedPositionOffsetDegrees field value
// and a boolean to check if the value has been set.
func (o *SmartlockAdvancedConfig) GetSingleLockedPositionOffsetDegreesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SingleLockedPositionOffsetDegrees, true
}

// SetSingleLockedPositionOffsetDegrees sets field value
func (o *SmartlockAdvancedConfig) SetSingleLockedPositionOffsetDegrees(v int32) {
	o.SingleLockedPositionOffsetDegrees = v
}

// GetUnlockedToLockedTransitionOffsetDegrees returns the UnlockedToLockedTransitionOffsetDegrees field value if set, zero value otherwise.
func (o *SmartlockAdvancedConfig) GetUnlockedToLockedTransitionOffsetDegrees() int32 {
	if o == nil || IsNil(o.UnlockedToLockedTransitionOffsetDegrees) {
		var ret int32
		return ret
	}
	return *o.UnlockedToLockedTransitionOffsetDegrees
}

// GetUnlockedToLockedTransitionOffsetDegreesOk returns a tuple with the UnlockedToLockedTransitionOffsetDegrees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockAdvancedConfig) GetUnlockedToLockedTransitionOffsetDegreesOk() (*int32, bool) {
	if o == nil || IsNil(o.UnlockedToLockedTransitionOffsetDegrees) {
		return nil, false
	}
	return o.UnlockedToLockedTransitionOffsetDegrees, true
}

// HasUnlockedToLockedTransitionOffsetDegrees returns a boolean if a field has been set.
func (o *SmartlockAdvancedConfig) HasUnlockedToLockedTransitionOffsetDegrees() bool {
	if o != nil && !IsNil(o.UnlockedToLockedTransitionOffsetDegrees) {
		return true
	}

	return false
}

// SetUnlockedToLockedTransitionOffsetDegrees gets a reference to the given int32 and assigns it to the UnlockedToLockedTransitionOffsetDegrees field.
func (o *SmartlockAdvancedConfig) SetUnlockedToLockedTransitionOffsetDegrees(v int32) {
	o.UnlockedToLockedTransitionOffsetDegrees = &v
}

// GetUnlockedPositionOffsetDegrees returns the UnlockedPositionOffsetDegrees field value
func (o *SmartlockAdvancedConfig) GetUnlockedPositionOffsetDegrees() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UnlockedPositionOffsetDegrees
}

// GetUnlockedPositionOffsetDegreesOk returns a tuple with the UnlockedPositionOffsetDegrees field value
// and a boolean to check if the value has been set.
func (o *SmartlockAdvancedConfig) GetUnlockedPositionOffsetDegreesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnlockedPositionOffsetDegrees, true
}

// SetUnlockedPositionOffsetDegrees sets field value
func (o *SmartlockAdvancedConfig) SetUnlockedPositionOffsetDegrees(v int32) {
	o.UnlockedPositionOffsetDegrees = v
}

// GetLockedPositionOffsetDegrees returns the LockedPositionOffsetDegrees field value
func (o *SmartlockAdvancedConfig) GetLockedPositionOffsetDegrees() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LockedPositionOffsetDegrees
}

// GetLockedPositionOffsetDegreesOk returns a tuple with the LockedPositionOffsetDegrees field value
// and a boolean to check if the value has been set.
func (o *SmartlockAdvancedConfig) GetLockedPositionOffsetDegreesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LockedPositionOffsetDegrees, true
}

// SetLockedPositionOffsetDegrees sets field value
func (o *SmartlockAdvancedConfig) SetLockedPositionOffsetDegrees(v int32) {
	o.LockedPositionOffsetDegrees = v
}

// GetDetachedCylinder returns the DetachedCylinder field value if set, zero value otherwise.
func (o *SmartlockAdvancedConfig) GetDetachedCylinder() bool {
	if o == nil || IsNil(o.DetachedCylinder) {
		var ret bool
		return ret
	}
	return *o.DetachedCylinder
}

// GetDetachedCylinderOk returns a tuple with the DetachedCylinder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockAdvancedConfig) GetDetachedCylinderOk() (*bool, bool) {
	if o == nil || IsNil(o.DetachedCylinder) {
		return nil, false
	}
	return o.DetachedCylinder, true
}

// HasDetachedCylinder returns a boolean if a field has been set.
func (o *SmartlockAdvancedConfig) HasDetachedCylinder() bool {
	if o != nil && !IsNil(o.DetachedCylinder) {
		return true
	}

	return false
}

// SetDetachedCylinder gets a reference to the given bool and assigns it to the DetachedCylinder field.
func (o *SmartlockAdvancedConfig) SetDetachedCylinder(v bool) {
	o.DetachedCylinder = &v
}

// GetBatteryType returns the BatteryType field value
func (o *SmartlockAdvancedConfig) GetBatteryType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BatteryType
}

// GetBatteryTypeOk returns a tuple with the BatteryType field value
// and a boolean to check if the value has been set.
func (o *SmartlockAdvancedConfig) GetBatteryTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BatteryType, true
}

// SetBatteryType sets field value
func (o *SmartlockAdvancedConfig) SetBatteryType(v int32) {
	o.BatteryType = v
}

// GetAutoLock returns the AutoLock field value if set, zero value otherwise.
func (o *SmartlockAdvancedConfig) GetAutoLock() bool {
	if o == nil || IsNil(o.AutoLock) {
		var ret bool
		return ret
	}
	return *o.AutoLock
}

// GetAutoLockOk returns a tuple with the AutoLock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockAdvancedConfig) GetAutoLockOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoLock) {
		return nil, false
	}
	return o.AutoLock, true
}

// HasAutoLock returns a boolean if a field has been set.
func (o *SmartlockAdvancedConfig) HasAutoLock() bool {
	if o != nil && !IsNil(o.AutoLock) {
		return true
	}

	return false
}

// SetAutoLock gets a reference to the given bool and assigns it to the AutoLock field.
func (o *SmartlockAdvancedConfig) SetAutoLock(v bool) {
	o.AutoLock = &v
}

// GetAutoLockTimeout returns the AutoLockTimeout field value if set, zero value otherwise.
func (o *SmartlockAdvancedConfig) GetAutoLockTimeout() int32 {
	if o == nil || IsNil(o.AutoLockTimeout) {
		var ret int32
		return ret
	}
	return *o.AutoLockTimeout
}

// GetAutoLockTimeoutOk returns a tuple with the AutoLockTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockAdvancedConfig) GetAutoLockTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.AutoLockTimeout) {
		return nil, false
	}
	return o.AutoLockTimeout, true
}

// HasAutoLockTimeout returns a boolean if a field has been set.
func (o *SmartlockAdvancedConfig) HasAutoLockTimeout() bool {
	if o != nil && !IsNil(o.AutoLockTimeout) {
		return true
	}

	return false
}

// SetAutoLockTimeout gets a reference to the given int32 and assigns it to the AutoLockTimeout field.
func (o *SmartlockAdvancedConfig) SetAutoLockTimeout(v int32) {
	o.AutoLockTimeout = &v
}

// GetAutoUpdateEnabled returns the AutoUpdateEnabled field value if set, zero value otherwise.
func (o *SmartlockAdvancedConfig) GetAutoUpdateEnabled() bool {
	if o == nil || IsNil(o.AutoUpdateEnabled) {
		var ret bool
		return ret
	}
	return *o.AutoUpdateEnabled
}

// GetAutoUpdateEnabledOk returns a tuple with the AutoUpdateEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockAdvancedConfig) GetAutoUpdateEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoUpdateEnabled) {
		return nil, false
	}
	return o.AutoUpdateEnabled, true
}

// HasAutoUpdateEnabled returns a boolean if a field has been set.
func (o *SmartlockAdvancedConfig) HasAutoUpdateEnabled() bool {
	if o != nil && !IsNil(o.AutoUpdateEnabled) {
		return true
	}

	return false
}

// SetAutoUpdateEnabled gets a reference to the given bool and assigns it to the AutoUpdateEnabled field.
func (o *SmartlockAdvancedConfig) SetAutoUpdateEnabled(v bool) {
	o.AutoUpdateEnabled = &v
}

// GetMotorSpeed returns the MotorSpeed field value if set, zero value otherwise.
func (o *SmartlockAdvancedConfig) GetMotorSpeed() int32 {
	if o == nil || IsNil(o.MotorSpeed) {
		var ret int32
		return ret
	}
	return *o.MotorSpeed
}

// GetMotorSpeedOk returns a tuple with the MotorSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockAdvancedConfig) GetMotorSpeedOk() (*int32, bool) {
	if o == nil || IsNil(o.MotorSpeed) {
		return nil, false
	}
	return o.MotorSpeed, true
}

// HasMotorSpeed returns a boolean if a field has been set.
func (o *SmartlockAdvancedConfig) HasMotorSpeed() bool {
	if o != nil && !IsNil(o.MotorSpeed) {
		return true
	}

	return false
}

// SetMotorSpeed gets a reference to the given int32 and assigns it to the MotorSpeed field.
func (o *SmartlockAdvancedConfig) SetMotorSpeed(v int32) {
	o.MotorSpeed = &v
}

// GetEnableSlowSpeedDuringNightmode returns the EnableSlowSpeedDuringNightmode field value if set, zero value otherwise.
func (o *SmartlockAdvancedConfig) GetEnableSlowSpeedDuringNightmode() bool {
	if o == nil || IsNil(o.EnableSlowSpeedDuringNightmode) {
		var ret bool
		return ret
	}
	return *o.EnableSlowSpeedDuringNightmode
}

// GetEnableSlowSpeedDuringNightmodeOk returns a tuple with the EnableSlowSpeedDuringNightmode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmartlockAdvancedConfig) GetEnableSlowSpeedDuringNightmodeOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableSlowSpeedDuringNightmode) {
		return nil, false
	}
	return o.EnableSlowSpeedDuringNightmode, true
}

// HasEnableSlowSpeedDuringNightmode returns a boolean if a field has been set.
func (o *SmartlockAdvancedConfig) HasEnableSlowSpeedDuringNightmode() bool {
	if o != nil && !IsNil(o.EnableSlowSpeedDuringNightmode) {
		return true
	}

	return false
}

// SetEnableSlowSpeedDuringNightmode gets a reference to the given bool and assigns it to the EnableSlowSpeedDuringNightmode field.
func (o *SmartlockAdvancedConfig) SetEnableSlowSpeedDuringNightmode(v bool) {
	o.EnableSlowSpeedDuringNightmode = &v
}

func (o SmartlockAdvancedConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmartlockAdvancedConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LngTimeout) {
		toSerialize["lngTimeout"] = o.LngTimeout
	}
	if !IsNil(o.SingleButtonPressAction) {
		toSerialize["singleButtonPressAction"] = o.SingleButtonPressAction
	}
	if !IsNil(o.DoubleButtonPressAction) {
		toSerialize["doubleButtonPressAction"] = o.DoubleButtonPressAction
	}
	if !IsNil(o.AutomaticBatteryTypeDetection) {
		toSerialize["automaticBatteryTypeDetection"] = o.AutomaticBatteryTypeDetection
	}
	if !IsNil(o.UnlatchDuration) {
		toSerialize["unlatchDuration"] = o.UnlatchDuration
	}
	if !IsNil(o.OperationId) {
		toSerialize["operationId"] = o.OperationId
	}
	toSerialize["totalDegrees"] = o.TotalDegrees
	toSerialize["singleLockedPositionOffsetDegrees"] = o.SingleLockedPositionOffsetDegrees
	if !IsNil(o.UnlockedToLockedTransitionOffsetDegrees) {
		toSerialize["unlockedToLockedTransitionOffsetDegrees"] = o.UnlockedToLockedTransitionOffsetDegrees
	}
	toSerialize["unlockedPositionOffsetDegrees"] = o.UnlockedPositionOffsetDegrees
	toSerialize["lockedPositionOffsetDegrees"] = o.LockedPositionOffsetDegrees
	if !IsNil(o.DetachedCylinder) {
		toSerialize["detachedCylinder"] = o.DetachedCylinder
	}
	toSerialize["batteryType"] = o.BatteryType
	if !IsNil(o.AutoLock) {
		toSerialize["autoLock"] = o.AutoLock
	}
	if !IsNil(o.AutoLockTimeout) {
		toSerialize["autoLockTimeout"] = o.AutoLockTimeout
	}
	if !IsNil(o.AutoUpdateEnabled) {
		toSerialize["autoUpdateEnabled"] = o.AutoUpdateEnabled
	}
	if !IsNil(o.MotorSpeed) {
		toSerialize["motorSpeed"] = o.MotorSpeed
	}
	if !IsNil(o.EnableSlowSpeedDuringNightmode) {
		toSerialize["enableSlowSpeedDuringNightmode"] = o.EnableSlowSpeedDuringNightmode
	}
	return toSerialize, nil
}

func (o *SmartlockAdvancedConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"totalDegrees",
		"singleLockedPositionOffsetDegrees",
		"unlockedPositionOffsetDegrees",
		"lockedPositionOffsetDegrees",
		"batteryType",
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

	varSmartlockAdvancedConfig := _SmartlockAdvancedConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSmartlockAdvancedConfig)

	if err != nil {
		return err
	}

	*o = SmartlockAdvancedConfig(varSmartlockAdvancedConfig)

	return err
}

type NullableSmartlockAdvancedConfig struct {
	value *SmartlockAdvancedConfig
	isSet bool
}

func (v NullableSmartlockAdvancedConfig) Get() *SmartlockAdvancedConfig {
	return v.value
}

func (v *NullableSmartlockAdvancedConfig) Set(val *SmartlockAdvancedConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSmartlockAdvancedConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSmartlockAdvancedConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmartlockAdvancedConfig(val *SmartlockAdvancedConfig) *NullableSmartlockAdvancedConfig {
	return &NullableSmartlockAdvancedConfig{value: val, isSet: true}
}

func (v NullableSmartlockAdvancedConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmartlockAdvancedConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


