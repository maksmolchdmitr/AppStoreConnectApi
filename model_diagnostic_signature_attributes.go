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

// checks if the DiagnosticSignatureAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiagnosticSignatureAttributes{}

// DiagnosticSignatureAttributes struct for DiagnosticSignatureAttributes
type DiagnosticSignatureAttributes struct {
	DiagnosticType *string  `json:"diagnosticType,omitempty"`
	Signature      *string  `json:"signature,omitempty"`
	Weight         *float32 `json:"weight,omitempty"`
}

// NewDiagnosticSignatureAttributes instantiates a new DiagnosticSignatureAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiagnosticSignatureAttributes() *DiagnosticSignatureAttributes {
	this := DiagnosticSignatureAttributes{}
	return &this
}

// NewDiagnosticSignatureAttributesWithDefaults instantiates a new DiagnosticSignatureAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiagnosticSignatureAttributesWithDefaults() *DiagnosticSignatureAttributes {
	this := DiagnosticSignatureAttributes{}
	return &this
}

// GetDiagnosticType returns the DiagnosticType field value if set, zero value otherwise.
func (o *DiagnosticSignatureAttributes) GetDiagnosticType() string {
	if o == nil || IsNil(o.DiagnosticType) {
		var ret string
		return ret
	}
	return *o.DiagnosticType
}

// GetDiagnosticTypeOk returns a tuple with the DiagnosticType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSignatureAttributes) GetDiagnosticTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DiagnosticType) {
		return nil, false
	}
	return o.DiagnosticType, true
}

// HasDiagnosticType returns a boolean if a field has been set.
func (o *DiagnosticSignatureAttributes) HasDiagnosticType() bool {
	if o != nil && !IsNil(o.DiagnosticType) {
		return true
	}

	return false
}

// SetDiagnosticType gets a reference to the given string and assigns it to the DiagnosticType field.
func (o *DiagnosticSignatureAttributes) SetDiagnosticType(v string) {
	o.DiagnosticType = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *DiagnosticSignatureAttributes) GetSignature() string {
	if o == nil || IsNil(o.Signature) {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSignatureAttributes) GetSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *DiagnosticSignatureAttributes) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *DiagnosticSignatureAttributes) SetSignature(v string) {
	o.Signature = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *DiagnosticSignatureAttributes) GetWeight() float32 {
	if o == nil || IsNil(o.Weight) {
		var ret float32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticSignatureAttributes) GetWeightOk() (*float32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *DiagnosticSignatureAttributes) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given float32 and assigns it to the Weight field.
func (o *DiagnosticSignatureAttributes) SetWeight(v float32) {
	o.Weight = &v
}

func (o DiagnosticSignatureAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiagnosticSignatureAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DiagnosticType) {
		toSerialize["diagnosticType"] = o.DiagnosticType
	}
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

type NullableDiagnosticSignatureAttributes struct {
	value *DiagnosticSignatureAttributes
	isSet bool
}

func (v NullableDiagnosticSignatureAttributes) Get() *DiagnosticSignatureAttributes {
	return v.value
}

func (v *NullableDiagnosticSignatureAttributes) Set(val *DiagnosticSignatureAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableDiagnosticSignatureAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableDiagnosticSignatureAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiagnosticSignatureAttributes(val *DiagnosticSignatureAttributes) *NullableDiagnosticSignatureAttributes {
	return &NullableDiagnosticSignatureAttributes{value: val, isSet: true}
}

func (v NullableDiagnosticSignatureAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiagnosticSignatureAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
