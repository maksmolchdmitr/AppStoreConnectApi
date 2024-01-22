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

// checks if the AppCustomProductPageLocalizationRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageLocalizationRelationships{}

// AppCustomProductPageLocalizationRelationships struct for AppCustomProductPageLocalizationRelationships
type AppCustomProductPageLocalizationRelationships struct {
	AppCustomProductPageVersion *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion `json:"appCustomProductPageVersion,omitempty"`
	AppScreenshotSets           *AppCustomProductPageLocalizationRelationshipsAppScreenshotSets           `json:"appScreenshotSets,omitempty"`
	AppPreviewSets              *AppCustomProductPageLocalizationRelationshipsAppPreviewSets              `json:"appPreviewSets,omitempty"`
}

// NewAppCustomProductPageLocalizationRelationships instantiates a new AppCustomProductPageLocalizationRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageLocalizationRelationships() *AppCustomProductPageLocalizationRelationships {
	this := AppCustomProductPageLocalizationRelationships{}
	return &this
}

// NewAppCustomProductPageLocalizationRelationshipsWithDefaults instantiates a new AppCustomProductPageLocalizationRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageLocalizationRelationshipsWithDefaults() *AppCustomProductPageLocalizationRelationships {
	this := AppCustomProductPageLocalizationRelationships{}
	return &this
}

// GetAppCustomProductPageVersion returns the AppCustomProductPageVersion field value if set, zero value otherwise.
func (o *AppCustomProductPageLocalizationRelationships) GetAppCustomProductPageVersion() AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion {
	if o == nil || IsNil(o.AppCustomProductPageVersion) {
		var ret AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion
		return ret
	}
	return *o.AppCustomProductPageVersion
}

// GetAppCustomProductPageVersionOk returns a tuple with the AppCustomProductPageVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageLocalizationRelationships) GetAppCustomProductPageVersionOk() (*AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion, bool) {
	if o == nil || IsNil(o.AppCustomProductPageVersion) {
		return nil, false
	}
	return o.AppCustomProductPageVersion, true
}

// HasAppCustomProductPageVersion returns a boolean if a field has been set.
func (o *AppCustomProductPageLocalizationRelationships) HasAppCustomProductPageVersion() bool {
	if o != nil && !IsNil(o.AppCustomProductPageVersion) {
		return true
	}

	return false
}

// SetAppCustomProductPageVersion gets a reference to the given AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion and assigns it to the AppCustomProductPageVersion field.
func (o *AppCustomProductPageLocalizationRelationships) SetAppCustomProductPageVersion(v AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersion) {
	o.AppCustomProductPageVersion = &v
}

// GetAppScreenshotSets returns the AppScreenshotSets field value if set, zero value otherwise.
func (o *AppCustomProductPageLocalizationRelationships) GetAppScreenshotSets() AppCustomProductPageLocalizationRelationshipsAppScreenshotSets {
	if o == nil || IsNil(o.AppScreenshotSets) {
		var ret AppCustomProductPageLocalizationRelationshipsAppScreenshotSets
		return ret
	}
	return *o.AppScreenshotSets
}

// GetAppScreenshotSetsOk returns a tuple with the AppScreenshotSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageLocalizationRelationships) GetAppScreenshotSetsOk() (*AppCustomProductPageLocalizationRelationshipsAppScreenshotSets, bool) {
	if o == nil || IsNil(o.AppScreenshotSets) {
		return nil, false
	}
	return o.AppScreenshotSets, true
}

// HasAppScreenshotSets returns a boolean if a field has been set.
func (o *AppCustomProductPageLocalizationRelationships) HasAppScreenshotSets() bool {
	if o != nil && !IsNil(o.AppScreenshotSets) {
		return true
	}

	return false
}

// SetAppScreenshotSets gets a reference to the given AppCustomProductPageLocalizationRelationshipsAppScreenshotSets and assigns it to the AppScreenshotSets field.
func (o *AppCustomProductPageLocalizationRelationships) SetAppScreenshotSets(v AppCustomProductPageLocalizationRelationshipsAppScreenshotSets) {
	o.AppScreenshotSets = &v
}

// GetAppPreviewSets returns the AppPreviewSets field value if set, zero value otherwise.
func (o *AppCustomProductPageLocalizationRelationships) GetAppPreviewSets() AppCustomProductPageLocalizationRelationshipsAppPreviewSets {
	if o == nil || IsNil(o.AppPreviewSets) {
		var ret AppCustomProductPageLocalizationRelationshipsAppPreviewSets
		return ret
	}
	return *o.AppPreviewSets
}

// GetAppPreviewSetsOk returns a tuple with the AppPreviewSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageLocalizationRelationships) GetAppPreviewSetsOk() (*AppCustomProductPageLocalizationRelationshipsAppPreviewSets, bool) {
	if o == nil || IsNil(o.AppPreviewSets) {
		return nil, false
	}
	return o.AppPreviewSets, true
}

// HasAppPreviewSets returns a boolean if a field has been set.
func (o *AppCustomProductPageLocalizationRelationships) HasAppPreviewSets() bool {
	if o != nil && !IsNil(o.AppPreviewSets) {
		return true
	}

	return false
}

// SetAppPreviewSets gets a reference to the given AppCustomProductPageLocalizationRelationshipsAppPreviewSets and assigns it to the AppPreviewSets field.
func (o *AppCustomProductPageLocalizationRelationships) SetAppPreviewSets(v AppCustomProductPageLocalizationRelationshipsAppPreviewSets) {
	o.AppPreviewSets = &v
}

func (o AppCustomProductPageLocalizationRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageLocalizationRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppCustomProductPageVersion) {
		toSerialize["appCustomProductPageVersion"] = o.AppCustomProductPageVersion
	}
	if !IsNil(o.AppScreenshotSets) {
		toSerialize["appScreenshotSets"] = o.AppScreenshotSets
	}
	if !IsNil(o.AppPreviewSets) {
		toSerialize["appPreviewSets"] = o.AppPreviewSets
	}
	return toSerialize, nil
}

type NullableAppCustomProductPageLocalizationRelationships struct {
	value *AppCustomProductPageLocalizationRelationships
	isSet bool
}

func (v NullableAppCustomProductPageLocalizationRelationships) Get() *AppCustomProductPageLocalizationRelationships {
	return v.value
}

func (v *NullableAppCustomProductPageLocalizationRelationships) Set(val *AppCustomProductPageLocalizationRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageLocalizationRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageLocalizationRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageLocalizationRelationships(val *AppCustomProductPageLocalizationRelationships) *NullableAppCustomProductPageLocalizationRelationships {
	return &NullableAppCustomProductPageLocalizationRelationships{value: val, isSet: true}
}

func (v NullableAppCustomProductPageLocalizationRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageLocalizationRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
