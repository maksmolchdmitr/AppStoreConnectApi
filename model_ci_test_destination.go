/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CiTestDestination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiTestDestination{}

// CiTestDestination struct for CiTestDestination
type CiTestDestination struct {
	DeviceTypeName       *string                `json:"deviceTypeName,omitempty"`
	DeviceTypeIdentifier *string                `json:"deviceTypeIdentifier,omitempty"`
	RuntimeName          *string                `json:"runtimeName,omitempty"`
	RuntimeIdentifier    *string                `json:"runtimeIdentifier,omitempty"`
	Kind                 *CiTestDestinationKind `json:"kind,omitempty"`
}

// NewCiTestDestination instantiates a new CiTestDestination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiTestDestination() *CiTestDestination {
	this := CiTestDestination{}
	return &this
}

// NewCiTestDestinationWithDefaults instantiates a new CiTestDestination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiTestDestinationWithDefaults() *CiTestDestination {
	this := CiTestDestination{}
	return &this
}

// GetDeviceTypeName returns the DeviceTypeName field value if set, zero value otherwise.
func (o *CiTestDestination) GetDeviceTypeName() string {
	if o == nil || IsNil(o.DeviceTypeName) {
		var ret string
		return ret
	}
	return *o.DeviceTypeName
}

// GetDeviceTypeNameOk returns a tuple with the DeviceTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiTestDestination) GetDeviceTypeNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceTypeName) {
		return nil, false
	}
	return o.DeviceTypeName, true
}

// HasDeviceTypeName returns a boolean if a field has been set.
func (o *CiTestDestination) HasDeviceTypeName() bool {
	if o != nil && !IsNil(o.DeviceTypeName) {
		return true
	}

	return false
}

// SetDeviceTypeName gets a reference to the given string and assigns it to the DeviceTypeName field.
func (o *CiTestDestination) SetDeviceTypeName(v string) {
	o.DeviceTypeName = &v
}

// GetDeviceTypeIdentifier returns the DeviceTypeIdentifier field value if set, zero value otherwise.
func (o *CiTestDestination) GetDeviceTypeIdentifier() string {
	if o == nil || IsNil(o.DeviceTypeIdentifier) {
		var ret string
		return ret
	}
	return *o.DeviceTypeIdentifier
}

// GetDeviceTypeIdentifierOk returns a tuple with the DeviceTypeIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiTestDestination) GetDeviceTypeIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceTypeIdentifier) {
		return nil, false
	}
	return o.DeviceTypeIdentifier, true
}

// HasDeviceTypeIdentifier returns a boolean if a field has been set.
func (o *CiTestDestination) HasDeviceTypeIdentifier() bool {
	if o != nil && !IsNil(o.DeviceTypeIdentifier) {
		return true
	}

	return false
}

// SetDeviceTypeIdentifier gets a reference to the given string and assigns it to the DeviceTypeIdentifier field.
func (o *CiTestDestination) SetDeviceTypeIdentifier(v string) {
	o.DeviceTypeIdentifier = &v
}

// GetRuntimeName returns the RuntimeName field value if set, zero value otherwise.
func (o *CiTestDestination) GetRuntimeName() string {
	if o == nil || IsNil(o.RuntimeName) {
		var ret string
		return ret
	}
	return *o.RuntimeName
}

// GetRuntimeNameOk returns a tuple with the RuntimeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiTestDestination) GetRuntimeNameOk() (*string, bool) {
	if o == nil || IsNil(o.RuntimeName) {
		return nil, false
	}
	return o.RuntimeName, true
}

// HasRuntimeName returns a boolean if a field has been set.
func (o *CiTestDestination) HasRuntimeName() bool {
	if o != nil && !IsNil(o.RuntimeName) {
		return true
	}

	return false
}

// SetRuntimeName gets a reference to the given string and assigns it to the RuntimeName field.
func (o *CiTestDestination) SetRuntimeName(v string) {
	o.RuntimeName = &v
}

// GetRuntimeIdentifier returns the RuntimeIdentifier field value if set, zero value otherwise.
func (o *CiTestDestination) GetRuntimeIdentifier() string {
	if o == nil || IsNil(o.RuntimeIdentifier) {
		var ret string
		return ret
	}
	return *o.RuntimeIdentifier
}

// GetRuntimeIdentifierOk returns a tuple with the RuntimeIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiTestDestination) GetRuntimeIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.RuntimeIdentifier) {
		return nil, false
	}
	return o.RuntimeIdentifier, true
}

// HasRuntimeIdentifier returns a boolean if a field has been set.
func (o *CiTestDestination) HasRuntimeIdentifier() bool {
	if o != nil && !IsNil(o.RuntimeIdentifier) {
		return true
	}

	return false
}

// SetRuntimeIdentifier gets a reference to the given string and assigns it to the RuntimeIdentifier field.
func (o *CiTestDestination) SetRuntimeIdentifier(v string) {
	o.RuntimeIdentifier = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *CiTestDestination) GetKind() CiTestDestinationKind {
	if o == nil || IsNil(o.Kind) {
		var ret CiTestDestinationKind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiTestDestination) GetKindOk() (*CiTestDestinationKind, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *CiTestDestination) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given CiTestDestinationKind and assigns it to the Kind field.
func (o *CiTestDestination) SetKind(v CiTestDestinationKind) {
	o.Kind = &v
}

func (o CiTestDestination) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiTestDestination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceTypeName) {
		toSerialize["deviceTypeName"] = o.DeviceTypeName
	}
	if !IsNil(o.DeviceTypeIdentifier) {
		toSerialize["deviceTypeIdentifier"] = o.DeviceTypeIdentifier
	}
	if !IsNil(o.RuntimeName) {
		toSerialize["runtimeName"] = o.RuntimeName
	}
	if !IsNil(o.RuntimeIdentifier) {
		toSerialize["runtimeIdentifier"] = o.RuntimeIdentifier
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	return toSerialize, nil
}

type NullableCiTestDestination struct {
	value *CiTestDestination
	isSet bool
}

func (v NullableCiTestDestination) Get() *CiTestDestination {
	return v.value
}

func (v *NullableCiTestDestination) Set(val *CiTestDestination) {
	v.value = val
	v.isSet = true
}

func (v NullableCiTestDestination) IsSet() bool {
	return v.isSet
}

func (v *NullableCiTestDestination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiTestDestination(val *CiTestDestination) *NullableCiTestDestination {
	return &NullableCiTestDestination{value: val, isSet: true}
}

func (v NullableCiTestDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiTestDestination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}