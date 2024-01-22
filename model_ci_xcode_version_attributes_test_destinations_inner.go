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

// checks if the CiXcodeVersionAttributesTestDestinationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiXcodeVersionAttributesTestDestinationsInner{}

// CiXcodeVersionAttributesTestDestinationsInner struct for CiXcodeVersionAttributesTestDestinationsInner
type CiXcodeVersionAttributesTestDestinationsInner struct {
	DeviceTypeName       *string                                                               `json:"deviceTypeName,omitempty"`
	DeviceTypeIdentifier *string                                                               `json:"deviceTypeIdentifier,omitempty"`
	AvailableRuntimes    []CiXcodeVersionAttributesTestDestinationsInnerAvailableRuntimesInner `json:"availableRuntimes,omitempty"`
	Kind                 *CiTestDestinationKind                                                `json:"kind,omitempty"`
}

// NewCiXcodeVersionAttributesTestDestinationsInner instantiates a new CiXcodeVersionAttributesTestDestinationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiXcodeVersionAttributesTestDestinationsInner() *CiXcodeVersionAttributesTestDestinationsInner {
	this := CiXcodeVersionAttributesTestDestinationsInner{}
	return &this
}

// NewCiXcodeVersionAttributesTestDestinationsInnerWithDefaults instantiates a new CiXcodeVersionAttributesTestDestinationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiXcodeVersionAttributesTestDestinationsInnerWithDefaults() *CiXcodeVersionAttributesTestDestinationsInner {
	this := CiXcodeVersionAttributesTestDestinationsInner{}
	return &this
}

// GetDeviceTypeName returns the DeviceTypeName field value if set, zero value otherwise.
func (o *CiXcodeVersionAttributesTestDestinationsInner) GetDeviceTypeName() string {
	if o == nil || IsNil(o.DeviceTypeName) {
		var ret string
		return ret
	}
	return *o.DeviceTypeName
}

// GetDeviceTypeNameOk returns a tuple with the DeviceTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiXcodeVersionAttributesTestDestinationsInner) GetDeviceTypeNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceTypeName) {
		return nil, false
	}
	return o.DeviceTypeName, true
}

// HasDeviceTypeName returns a boolean if a field has been set.
func (o *CiXcodeVersionAttributesTestDestinationsInner) HasDeviceTypeName() bool {
	if o != nil && !IsNil(o.DeviceTypeName) {
		return true
	}

	return false
}

// SetDeviceTypeName gets a reference to the given string and assigns it to the DeviceTypeName field.
func (o *CiXcodeVersionAttributesTestDestinationsInner) SetDeviceTypeName(v string) {
	o.DeviceTypeName = &v
}

// GetDeviceTypeIdentifier returns the DeviceTypeIdentifier field value if set, zero value otherwise.
func (o *CiXcodeVersionAttributesTestDestinationsInner) GetDeviceTypeIdentifier() string {
	if o == nil || IsNil(o.DeviceTypeIdentifier) {
		var ret string
		return ret
	}
	return *o.DeviceTypeIdentifier
}

// GetDeviceTypeIdentifierOk returns a tuple with the DeviceTypeIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiXcodeVersionAttributesTestDestinationsInner) GetDeviceTypeIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceTypeIdentifier) {
		return nil, false
	}
	return o.DeviceTypeIdentifier, true
}

// HasDeviceTypeIdentifier returns a boolean if a field has been set.
func (o *CiXcodeVersionAttributesTestDestinationsInner) HasDeviceTypeIdentifier() bool {
	if o != nil && !IsNil(o.DeviceTypeIdentifier) {
		return true
	}

	return false
}

// SetDeviceTypeIdentifier gets a reference to the given string and assigns it to the DeviceTypeIdentifier field.
func (o *CiXcodeVersionAttributesTestDestinationsInner) SetDeviceTypeIdentifier(v string) {
	o.DeviceTypeIdentifier = &v
}

// GetAvailableRuntimes returns the AvailableRuntimes field value if set, zero value otherwise.
func (o *CiXcodeVersionAttributesTestDestinationsInner) GetAvailableRuntimes() []CiXcodeVersionAttributesTestDestinationsInnerAvailableRuntimesInner {
	if o == nil || IsNil(o.AvailableRuntimes) {
		var ret []CiXcodeVersionAttributesTestDestinationsInnerAvailableRuntimesInner
		return ret
	}
	return o.AvailableRuntimes
}

// GetAvailableRuntimesOk returns a tuple with the AvailableRuntimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiXcodeVersionAttributesTestDestinationsInner) GetAvailableRuntimesOk() ([]CiXcodeVersionAttributesTestDestinationsInnerAvailableRuntimesInner, bool) {
	if o == nil || IsNil(o.AvailableRuntimes) {
		return nil, false
	}
	return o.AvailableRuntimes, true
}

// HasAvailableRuntimes returns a boolean if a field has been set.
func (o *CiXcodeVersionAttributesTestDestinationsInner) HasAvailableRuntimes() bool {
	if o != nil && !IsNil(o.AvailableRuntimes) {
		return true
	}

	return false
}

// SetAvailableRuntimes gets a reference to the given []CiXcodeVersionAttributesTestDestinationsInnerAvailableRuntimesInner and assigns it to the AvailableRuntimes field.
func (o *CiXcodeVersionAttributesTestDestinationsInner) SetAvailableRuntimes(v []CiXcodeVersionAttributesTestDestinationsInnerAvailableRuntimesInner) {
	o.AvailableRuntimes = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *CiXcodeVersionAttributesTestDestinationsInner) GetKind() CiTestDestinationKind {
	if o == nil || IsNil(o.Kind) {
		var ret CiTestDestinationKind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiXcodeVersionAttributesTestDestinationsInner) GetKindOk() (*CiTestDestinationKind, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *CiXcodeVersionAttributesTestDestinationsInner) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given CiTestDestinationKind and assigns it to the Kind field.
func (o *CiXcodeVersionAttributesTestDestinationsInner) SetKind(v CiTestDestinationKind) {
	o.Kind = &v
}

func (o CiXcodeVersionAttributesTestDestinationsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiXcodeVersionAttributesTestDestinationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceTypeName) {
		toSerialize["deviceTypeName"] = o.DeviceTypeName
	}
	if !IsNil(o.DeviceTypeIdentifier) {
		toSerialize["deviceTypeIdentifier"] = o.DeviceTypeIdentifier
	}
	if !IsNil(o.AvailableRuntimes) {
		toSerialize["availableRuntimes"] = o.AvailableRuntimes
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	return toSerialize, nil
}

type NullableCiXcodeVersionAttributesTestDestinationsInner struct {
	value *CiXcodeVersionAttributesTestDestinationsInner
	isSet bool
}

func (v NullableCiXcodeVersionAttributesTestDestinationsInner) Get() *CiXcodeVersionAttributesTestDestinationsInner {
	return v.value
}

func (v *NullableCiXcodeVersionAttributesTestDestinationsInner) Set(val *CiXcodeVersionAttributesTestDestinationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCiXcodeVersionAttributesTestDestinationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCiXcodeVersionAttributesTestDestinationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiXcodeVersionAttributesTestDestinationsInner(val *CiXcodeVersionAttributesTestDestinationsInner) *NullableCiXcodeVersionAttributesTestDestinationsInner {
	return &NullableCiXcodeVersionAttributesTestDestinationsInner{value: val, isSet: true}
}

func (v NullableCiXcodeVersionAttributesTestDestinationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiXcodeVersionAttributesTestDestinationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
