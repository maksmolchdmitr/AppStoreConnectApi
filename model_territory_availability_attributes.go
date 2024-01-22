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

// checks if the TerritoryAvailabilityAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TerritoryAvailabilityAttributes{}

// TerritoryAvailabilityAttributes struct for TerritoryAvailabilityAttributes
type TerritoryAvailabilityAttributes struct {
	Available           *bool    `json:"available,omitempty"`
	ReleaseDate         *string  `json:"releaseDate,omitempty"`
	PreOrderEnabled     *bool    `json:"preOrderEnabled,omitempty"`
	PreOrderPublishDate *string  `json:"preOrderPublishDate,omitempty"`
	ContentStatuses     []string `json:"contentStatuses,omitempty"`
}

// NewTerritoryAvailabilityAttributes instantiates a new TerritoryAvailabilityAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTerritoryAvailabilityAttributes() *TerritoryAvailabilityAttributes {
	this := TerritoryAvailabilityAttributes{}
	return &this
}

// NewTerritoryAvailabilityAttributesWithDefaults instantiates a new TerritoryAvailabilityAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTerritoryAvailabilityAttributesWithDefaults() *TerritoryAvailabilityAttributes {
	this := TerritoryAvailabilityAttributes{}
	return &this
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *TerritoryAvailabilityAttributes) GetAvailable() bool {
	if o == nil || IsNil(o.Available) {
		var ret bool
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerritoryAvailabilityAttributes) GetAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *TerritoryAvailabilityAttributes) HasAvailable() bool {
	if o != nil && !IsNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given bool and assigns it to the Available field.
func (o *TerritoryAvailabilityAttributes) SetAvailable(v bool) {
	o.Available = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *TerritoryAvailabilityAttributes) GetReleaseDate() string {
	if o == nil || IsNil(o.ReleaseDate) {
		var ret string
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerritoryAvailabilityAttributes) GetReleaseDateOk() (*string, bool) {
	if o == nil || IsNil(o.ReleaseDate) {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *TerritoryAvailabilityAttributes) HasReleaseDate() bool {
	if o != nil && !IsNil(o.ReleaseDate) {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given string and assigns it to the ReleaseDate field.
func (o *TerritoryAvailabilityAttributes) SetReleaseDate(v string) {
	o.ReleaseDate = &v
}

// GetPreOrderEnabled returns the PreOrderEnabled field value if set, zero value otherwise.
func (o *TerritoryAvailabilityAttributes) GetPreOrderEnabled() bool {
	if o == nil || IsNil(o.PreOrderEnabled) {
		var ret bool
		return ret
	}
	return *o.PreOrderEnabled
}

// GetPreOrderEnabledOk returns a tuple with the PreOrderEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerritoryAvailabilityAttributes) GetPreOrderEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PreOrderEnabled) {
		return nil, false
	}
	return o.PreOrderEnabled, true
}

// HasPreOrderEnabled returns a boolean if a field has been set.
func (o *TerritoryAvailabilityAttributes) HasPreOrderEnabled() bool {
	if o != nil && !IsNil(o.PreOrderEnabled) {
		return true
	}

	return false
}

// SetPreOrderEnabled gets a reference to the given bool and assigns it to the PreOrderEnabled field.
func (o *TerritoryAvailabilityAttributes) SetPreOrderEnabled(v bool) {
	o.PreOrderEnabled = &v
}

// GetPreOrderPublishDate returns the PreOrderPublishDate field value if set, zero value otherwise.
func (o *TerritoryAvailabilityAttributes) GetPreOrderPublishDate() string {
	if o == nil || IsNil(o.PreOrderPublishDate) {
		var ret string
		return ret
	}
	return *o.PreOrderPublishDate
}

// GetPreOrderPublishDateOk returns a tuple with the PreOrderPublishDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerritoryAvailabilityAttributes) GetPreOrderPublishDateOk() (*string, bool) {
	if o == nil || IsNil(o.PreOrderPublishDate) {
		return nil, false
	}
	return o.PreOrderPublishDate, true
}

// HasPreOrderPublishDate returns a boolean if a field has been set.
func (o *TerritoryAvailabilityAttributes) HasPreOrderPublishDate() bool {
	if o != nil && !IsNil(o.PreOrderPublishDate) {
		return true
	}

	return false
}

// SetPreOrderPublishDate gets a reference to the given string and assigns it to the PreOrderPublishDate field.
func (o *TerritoryAvailabilityAttributes) SetPreOrderPublishDate(v string) {
	o.PreOrderPublishDate = &v
}

// GetContentStatuses returns the ContentStatuses field value if set, zero value otherwise.
func (o *TerritoryAvailabilityAttributes) GetContentStatuses() []string {
	if o == nil || IsNil(o.ContentStatuses) {
		var ret []string
		return ret
	}
	return o.ContentStatuses
}

// GetContentStatusesOk returns a tuple with the ContentStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TerritoryAvailabilityAttributes) GetContentStatusesOk() ([]string, bool) {
	if o == nil || IsNil(o.ContentStatuses) {
		return nil, false
	}
	return o.ContentStatuses, true
}

// HasContentStatuses returns a boolean if a field has been set.
func (o *TerritoryAvailabilityAttributes) HasContentStatuses() bool {
	if o != nil && !IsNil(o.ContentStatuses) {
		return true
	}

	return false
}

// SetContentStatuses gets a reference to the given []string and assigns it to the ContentStatuses field.
func (o *TerritoryAvailabilityAttributes) SetContentStatuses(v []string) {
	o.ContentStatuses = v
}

func (o TerritoryAvailabilityAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TerritoryAvailabilityAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !IsNil(o.ReleaseDate) {
		toSerialize["releaseDate"] = o.ReleaseDate
	}
	if !IsNil(o.PreOrderEnabled) {
		toSerialize["preOrderEnabled"] = o.PreOrderEnabled
	}
	if !IsNil(o.PreOrderPublishDate) {
		toSerialize["preOrderPublishDate"] = o.PreOrderPublishDate
	}
	if !IsNil(o.ContentStatuses) {
		toSerialize["contentStatuses"] = o.ContentStatuses
	}
	return toSerialize, nil
}

type NullableTerritoryAvailabilityAttributes struct {
	value *TerritoryAvailabilityAttributes
	isSet bool
}

func (v NullableTerritoryAvailabilityAttributes) Get() *TerritoryAvailabilityAttributes {
	return v.value
}

func (v *NullableTerritoryAvailabilityAttributes) Set(val *TerritoryAvailabilityAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableTerritoryAvailabilityAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableTerritoryAvailabilityAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTerritoryAvailabilityAttributes(val *TerritoryAvailabilityAttributes) *NullableTerritoryAvailabilityAttributes {
	return &NullableTerritoryAvailabilityAttributes{value: val, isSet: true}
}

func (v NullableTerritoryAvailabilityAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTerritoryAvailabilityAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
