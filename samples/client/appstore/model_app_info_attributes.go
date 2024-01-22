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

// checks if the AppInfoAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppInfoAttributes{}

// AppInfoAttributes struct for AppInfoAttributes
type AppInfoAttributes struct {
	AppStoreState *AppStoreVersionState `json:"appStoreState,omitempty"`
	AppStoreAgeRating *AppStoreAgeRating `json:"appStoreAgeRating,omitempty"`
	BrazilAgeRating *BrazilAgeRating `json:"brazilAgeRating,omitempty"`
	BrazilAgeRatingV2 *string `json:"brazilAgeRatingV2,omitempty"`
	KidsAgeBand *KidsAgeBand `json:"kidsAgeBand,omitempty"`
}

// NewAppInfoAttributes instantiates a new AppInfoAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInfoAttributes() *AppInfoAttributes {
	this := AppInfoAttributes{}
	return &this
}

// NewAppInfoAttributesWithDefaults instantiates a new AppInfoAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInfoAttributesWithDefaults() *AppInfoAttributes {
	this := AppInfoAttributes{}
	return &this
}

// GetAppStoreState returns the AppStoreState field value if set, zero value otherwise.
func (o *AppInfoAttributes) GetAppStoreState() AppStoreVersionState {
	if o == nil || IsNil(o.AppStoreState) {
		var ret AppStoreVersionState
		return ret
	}
	return *o.AppStoreState
}

// GetAppStoreStateOk returns a tuple with the AppStoreState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoAttributes) GetAppStoreStateOk() (*AppStoreVersionState, bool) {
	if o == nil || IsNil(o.AppStoreState) {
		return nil, false
	}
	return o.AppStoreState, true
}

// HasAppStoreState returns a boolean if a field has been set.
func (o *AppInfoAttributes) HasAppStoreState() bool {
	if o != nil && !IsNil(o.AppStoreState) {
		return true
	}

	return false
}

// SetAppStoreState gets a reference to the given AppStoreVersionState and assigns it to the AppStoreState field.
func (o *AppInfoAttributes) SetAppStoreState(v AppStoreVersionState) {
	o.AppStoreState = &v
}

// GetAppStoreAgeRating returns the AppStoreAgeRating field value if set, zero value otherwise.
func (o *AppInfoAttributes) GetAppStoreAgeRating() AppStoreAgeRating {
	if o == nil || IsNil(o.AppStoreAgeRating) {
		var ret AppStoreAgeRating
		return ret
	}
	return *o.AppStoreAgeRating
}

// GetAppStoreAgeRatingOk returns a tuple with the AppStoreAgeRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoAttributes) GetAppStoreAgeRatingOk() (*AppStoreAgeRating, bool) {
	if o == nil || IsNil(o.AppStoreAgeRating) {
		return nil, false
	}
	return o.AppStoreAgeRating, true
}

// HasAppStoreAgeRating returns a boolean if a field has been set.
func (o *AppInfoAttributes) HasAppStoreAgeRating() bool {
	if o != nil && !IsNil(o.AppStoreAgeRating) {
		return true
	}

	return false
}

// SetAppStoreAgeRating gets a reference to the given AppStoreAgeRating and assigns it to the AppStoreAgeRating field.
func (o *AppInfoAttributes) SetAppStoreAgeRating(v AppStoreAgeRating) {
	o.AppStoreAgeRating = &v
}

// GetBrazilAgeRating returns the BrazilAgeRating field value if set, zero value otherwise.
func (o *AppInfoAttributes) GetBrazilAgeRating() BrazilAgeRating {
	if o == nil || IsNil(o.BrazilAgeRating) {
		var ret BrazilAgeRating
		return ret
	}
	return *o.BrazilAgeRating
}

// GetBrazilAgeRatingOk returns a tuple with the BrazilAgeRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoAttributes) GetBrazilAgeRatingOk() (*BrazilAgeRating, bool) {
	if o == nil || IsNil(o.BrazilAgeRating) {
		return nil, false
	}
	return o.BrazilAgeRating, true
}

// HasBrazilAgeRating returns a boolean if a field has been set.
func (o *AppInfoAttributes) HasBrazilAgeRating() bool {
	if o != nil && !IsNil(o.BrazilAgeRating) {
		return true
	}

	return false
}

// SetBrazilAgeRating gets a reference to the given BrazilAgeRating and assigns it to the BrazilAgeRating field.
func (o *AppInfoAttributes) SetBrazilAgeRating(v BrazilAgeRating) {
	o.BrazilAgeRating = &v
}

// GetBrazilAgeRatingV2 returns the BrazilAgeRatingV2 field value if set, zero value otherwise.
func (o *AppInfoAttributes) GetBrazilAgeRatingV2() string {
	if o == nil || IsNil(o.BrazilAgeRatingV2) {
		var ret string
		return ret
	}
	return *o.BrazilAgeRatingV2
}

// GetBrazilAgeRatingV2Ok returns a tuple with the BrazilAgeRatingV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoAttributes) GetBrazilAgeRatingV2Ok() (*string, bool) {
	if o == nil || IsNil(o.BrazilAgeRatingV2) {
		return nil, false
	}
	return o.BrazilAgeRatingV2, true
}

// HasBrazilAgeRatingV2 returns a boolean if a field has been set.
func (o *AppInfoAttributes) HasBrazilAgeRatingV2() bool {
	if o != nil && !IsNil(o.BrazilAgeRatingV2) {
		return true
	}

	return false
}

// SetBrazilAgeRatingV2 gets a reference to the given string and assigns it to the BrazilAgeRatingV2 field.
func (o *AppInfoAttributes) SetBrazilAgeRatingV2(v string) {
	o.BrazilAgeRatingV2 = &v
}

// GetKidsAgeBand returns the KidsAgeBand field value if set, zero value otherwise.
func (o *AppInfoAttributes) GetKidsAgeBand() KidsAgeBand {
	if o == nil || IsNil(o.KidsAgeBand) {
		var ret KidsAgeBand
		return ret
	}
	return *o.KidsAgeBand
}

// GetKidsAgeBandOk returns a tuple with the KidsAgeBand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoAttributes) GetKidsAgeBandOk() (*KidsAgeBand, bool) {
	if o == nil || IsNil(o.KidsAgeBand) {
		return nil, false
	}
	return o.KidsAgeBand, true
}

// HasKidsAgeBand returns a boolean if a field has been set.
func (o *AppInfoAttributes) HasKidsAgeBand() bool {
	if o != nil && !IsNil(o.KidsAgeBand) {
		return true
	}

	return false
}

// SetKidsAgeBand gets a reference to the given KidsAgeBand and assigns it to the KidsAgeBand field.
func (o *AppInfoAttributes) SetKidsAgeBand(v KidsAgeBand) {
	o.KidsAgeBand = &v
}

func (o AppInfoAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppInfoAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppStoreState) {
		toSerialize["appStoreState"] = o.AppStoreState
	}
	if !IsNil(o.AppStoreAgeRating) {
		toSerialize["appStoreAgeRating"] = o.AppStoreAgeRating
	}
	if !IsNil(o.BrazilAgeRating) {
		toSerialize["brazilAgeRating"] = o.BrazilAgeRating
	}
	if !IsNil(o.BrazilAgeRatingV2) {
		toSerialize["brazilAgeRatingV2"] = o.BrazilAgeRatingV2
	}
	if !IsNil(o.KidsAgeBand) {
		toSerialize["kidsAgeBand"] = o.KidsAgeBand
	}
	return toSerialize, nil
}

type NullableAppInfoAttributes struct {
	value *AppInfoAttributes
	isSet bool
}

func (v NullableAppInfoAttributes) Get() *AppInfoAttributes {
	return v.value
}

func (v *NullableAppInfoAttributes) Set(val *AppInfoAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInfoAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInfoAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInfoAttributes(val *AppInfoAttributes) *NullableAppInfoAttributes {
	return &NullableAppInfoAttributes{value: val, isSet: true}
}

func (v NullableAppInfoAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInfoAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


