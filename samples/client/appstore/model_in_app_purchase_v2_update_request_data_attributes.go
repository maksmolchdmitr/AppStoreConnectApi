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

// checks if the InAppPurchaseV2UpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseV2UpdateRequestDataAttributes{}

// InAppPurchaseV2UpdateRequestDataAttributes struct for InAppPurchaseV2UpdateRequestDataAttributes
type InAppPurchaseV2UpdateRequestDataAttributes struct {
	Name *string `json:"name,omitempty"`
	ReviewNote *string `json:"reviewNote,omitempty"`
	FamilySharable *bool `json:"familySharable,omitempty"`
	// Deprecated
	AvailableInAllTerritories *bool `json:"availableInAllTerritories,omitempty"`
}

// NewInAppPurchaseV2UpdateRequestDataAttributes instantiates a new InAppPurchaseV2UpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseV2UpdateRequestDataAttributes() *InAppPurchaseV2UpdateRequestDataAttributes {
	this := InAppPurchaseV2UpdateRequestDataAttributes{}
	return &this
}

// NewInAppPurchaseV2UpdateRequestDataAttributesWithDefaults instantiates a new InAppPurchaseV2UpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseV2UpdateRequestDataAttributesWithDefaults() *InAppPurchaseV2UpdateRequestDataAttributes {
	this := InAppPurchaseV2UpdateRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InAppPurchaseV2UpdateRequestDataAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2UpdateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InAppPurchaseV2UpdateRequestDataAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InAppPurchaseV2UpdateRequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetReviewNote returns the ReviewNote field value if set, zero value otherwise.
func (o *InAppPurchaseV2UpdateRequestDataAttributes) GetReviewNote() string {
	if o == nil || IsNil(o.ReviewNote) {
		var ret string
		return ret
	}
	return *o.ReviewNote
}

// GetReviewNoteOk returns a tuple with the ReviewNote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2UpdateRequestDataAttributes) GetReviewNoteOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewNote) {
		return nil, false
	}
	return o.ReviewNote, true
}

// HasReviewNote returns a boolean if a field has been set.
func (o *InAppPurchaseV2UpdateRequestDataAttributes) HasReviewNote() bool {
	if o != nil && !IsNil(o.ReviewNote) {
		return true
	}

	return false
}

// SetReviewNote gets a reference to the given string and assigns it to the ReviewNote field.
func (o *InAppPurchaseV2UpdateRequestDataAttributes) SetReviewNote(v string) {
	o.ReviewNote = &v
}

// GetFamilySharable returns the FamilySharable field value if set, zero value otherwise.
func (o *InAppPurchaseV2UpdateRequestDataAttributes) GetFamilySharable() bool {
	if o == nil || IsNil(o.FamilySharable) {
		var ret bool
		return ret
	}
	return *o.FamilySharable
}

// GetFamilySharableOk returns a tuple with the FamilySharable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2UpdateRequestDataAttributes) GetFamilySharableOk() (*bool, bool) {
	if o == nil || IsNil(o.FamilySharable) {
		return nil, false
	}
	return o.FamilySharable, true
}

// HasFamilySharable returns a boolean if a field has been set.
func (o *InAppPurchaseV2UpdateRequestDataAttributes) HasFamilySharable() bool {
	if o != nil && !IsNil(o.FamilySharable) {
		return true
	}

	return false
}

// SetFamilySharable gets a reference to the given bool and assigns it to the FamilySharable field.
func (o *InAppPurchaseV2UpdateRequestDataAttributes) SetFamilySharable(v bool) {
	o.FamilySharable = &v
}

// GetAvailableInAllTerritories returns the AvailableInAllTerritories field value if set, zero value otherwise.
// Deprecated
func (o *InAppPurchaseV2UpdateRequestDataAttributes) GetAvailableInAllTerritories() bool {
	if o == nil || IsNil(o.AvailableInAllTerritories) {
		var ret bool
		return ret
	}
	return *o.AvailableInAllTerritories
}

// GetAvailableInAllTerritoriesOk returns a tuple with the AvailableInAllTerritories field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *InAppPurchaseV2UpdateRequestDataAttributes) GetAvailableInAllTerritoriesOk() (*bool, bool) {
	if o == nil || IsNil(o.AvailableInAllTerritories) {
		return nil, false
	}
	return o.AvailableInAllTerritories, true
}

// HasAvailableInAllTerritories returns a boolean if a field has been set.
func (o *InAppPurchaseV2UpdateRequestDataAttributes) HasAvailableInAllTerritories() bool {
	if o != nil && !IsNil(o.AvailableInAllTerritories) {
		return true
	}

	return false
}

// SetAvailableInAllTerritories gets a reference to the given bool and assigns it to the AvailableInAllTerritories field.
// Deprecated
func (o *InAppPurchaseV2UpdateRequestDataAttributes) SetAvailableInAllTerritories(v bool) {
	o.AvailableInAllTerritories = &v
}

func (o InAppPurchaseV2UpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseV2UpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ReviewNote) {
		toSerialize["reviewNote"] = o.ReviewNote
	}
	if !IsNil(o.FamilySharable) {
		toSerialize["familySharable"] = o.FamilySharable
	}
	if !IsNil(o.AvailableInAllTerritories) {
		toSerialize["availableInAllTerritories"] = o.AvailableInAllTerritories
	}
	return toSerialize, nil
}

type NullableInAppPurchaseV2UpdateRequestDataAttributes struct {
	value *InAppPurchaseV2UpdateRequestDataAttributes
	isSet bool
}

func (v NullableInAppPurchaseV2UpdateRequestDataAttributes) Get() *InAppPurchaseV2UpdateRequestDataAttributes {
	return v.value
}

func (v *NullableInAppPurchaseV2UpdateRequestDataAttributes) Set(val *InAppPurchaseV2UpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseV2UpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseV2UpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseV2UpdateRequestDataAttributes(val *InAppPurchaseV2UpdateRequestDataAttributes) *NullableInAppPurchaseV2UpdateRequestDataAttributes {
	return &NullableInAppPurchaseV2UpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableInAppPurchaseV2UpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseV2UpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


