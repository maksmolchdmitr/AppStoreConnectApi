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

// checks if the AppStoreVersionLocalizationRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionLocalizationRelationships{}

// AppStoreVersionLocalizationRelationships struct for AppStoreVersionLocalizationRelationships
type AppStoreVersionLocalizationRelationships struct {
	AppStoreVersion   *AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion `json:"appStoreVersion,omitempty"`
	AppScreenshotSets *AppCustomProductPageLocalizationRelationshipsAppScreenshotSets  `json:"appScreenshotSets,omitempty"`
	AppPreviewSets    *AppCustomProductPageLocalizationRelationshipsAppPreviewSets     `json:"appPreviewSets,omitempty"`
}

// NewAppStoreVersionLocalizationRelationships instantiates a new AppStoreVersionLocalizationRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionLocalizationRelationships() *AppStoreVersionLocalizationRelationships {
	this := AppStoreVersionLocalizationRelationships{}
	return &this
}

// NewAppStoreVersionLocalizationRelationshipsWithDefaults instantiates a new AppStoreVersionLocalizationRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionLocalizationRelationshipsWithDefaults() *AppStoreVersionLocalizationRelationships {
	this := AppStoreVersionLocalizationRelationships{}
	return &this
}

// GetAppStoreVersion returns the AppStoreVersion field value if set, zero value otherwise.
func (o *AppStoreVersionLocalizationRelationships) GetAppStoreVersion() AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion {
	if o == nil || IsNil(o.AppStoreVersion) {
		var ret AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion
		return ret
	}
	return *o.AppStoreVersion
}

// GetAppStoreVersionOk returns a tuple with the AppStoreVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationRelationships) GetAppStoreVersionOk() (*AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion, bool) {
	if o == nil || IsNil(o.AppStoreVersion) {
		return nil, false
	}
	return o.AppStoreVersion, true
}

// HasAppStoreVersion returns a boolean if a field has been set.
func (o *AppStoreVersionLocalizationRelationships) HasAppStoreVersion() bool {
	if o != nil && !IsNil(o.AppStoreVersion) {
		return true
	}

	return false
}

// SetAppStoreVersion gets a reference to the given AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion and assigns it to the AppStoreVersion field.
func (o *AppStoreVersionLocalizationRelationships) SetAppStoreVersion(v AppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersion) {
	o.AppStoreVersion = &v
}

// GetAppScreenshotSets returns the AppScreenshotSets field value if set, zero value otherwise.
func (o *AppStoreVersionLocalizationRelationships) GetAppScreenshotSets() AppCustomProductPageLocalizationRelationshipsAppScreenshotSets {
	if o == nil || IsNil(o.AppScreenshotSets) {
		var ret AppCustomProductPageLocalizationRelationshipsAppScreenshotSets
		return ret
	}
	return *o.AppScreenshotSets
}

// GetAppScreenshotSetsOk returns a tuple with the AppScreenshotSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationRelationships) GetAppScreenshotSetsOk() (*AppCustomProductPageLocalizationRelationshipsAppScreenshotSets, bool) {
	if o == nil || IsNil(o.AppScreenshotSets) {
		return nil, false
	}
	return o.AppScreenshotSets, true
}

// HasAppScreenshotSets returns a boolean if a field has been set.
func (o *AppStoreVersionLocalizationRelationships) HasAppScreenshotSets() bool {
	if o != nil && !IsNil(o.AppScreenshotSets) {
		return true
	}

	return false
}

// SetAppScreenshotSets gets a reference to the given AppCustomProductPageLocalizationRelationshipsAppScreenshotSets and assigns it to the AppScreenshotSets field.
func (o *AppStoreVersionLocalizationRelationships) SetAppScreenshotSets(v AppCustomProductPageLocalizationRelationshipsAppScreenshotSets) {
	o.AppScreenshotSets = &v
}

// GetAppPreviewSets returns the AppPreviewSets field value if set, zero value otherwise.
func (o *AppStoreVersionLocalizationRelationships) GetAppPreviewSets() AppCustomProductPageLocalizationRelationshipsAppPreviewSets {
	if o == nil || IsNil(o.AppPreviewSets) {
		var ret AppCustomProductPageLocalizationRelationshipsAppPreviewSets
		return ret
	}
	return *o.AppPreviewSets
}

// GetAppPreviewSetsOk returns a tuple with the AppPreviewSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationRelationships) GetAppPreviewSetsOk() (*AppCustomProductPageLocalizationRelationshipsAppPreviewSets, bool) {
	if o == nil || IsNil(o.AppPreviewSets) {
		return nil, false
	}
	return o.AppPreviewSets, true
}

// HasAppPreviewSets returns a boolean if a field has been set.
func (o *AppStoreVersionLocalizationRelationships) HasAppPreviewSets() bool {
	if o != nil && !IsNil(o.AppPreviewSets) {
		return true
	}

	return false
}

// SetAppPreviewSets gets a reference to the given AppCustomProductPageLocalizationRelationshipsAppPreviewSets and assigns it to the AppPreviewSets field.
func (o *AppStoreVersionLocalizationRelationships) SetAppPreviewSets(v AppCustomProductPageLocalizationRelationshipsAppPreviewSets) {
	o.AppPreviewSets = &v
}

func (o AppStoreVersionLocalizationRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionLocalizationRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppStoreVersion) {
		toSerialize["appStoreVersion"] = o.AppStoreVersion
	}
	if !IsNil(o.AppScreenshotSets) {
		toSerialize["appScreenshotSets"] = o.AppScreenshotSets
	}
	if !IsNil(o.AppPreviewSets) {
		toSerialize["appPreviewSets"] = o.AppPreviewSets
	}
	return toSerialize, nil
}

type NullableAppStoreVersionLocalizationRelationships struct {
	value *AppStoreVersionLocalizationRelationships
	isSet bool
}

func (v NullableAppStoreVersionLocalizationRelationships) Get() *AppStoreVersionLocalizationRelationships {
	return v.value
}

func (v *NullableAppStoreVersionLocalizationRelationships) Set(val *AppStoreVersionLocalizationRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionLocalizationRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionLocalizationRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionLocalizationRelationships(val *AppStoreVersionLocalizationRelationships) *NullableAppStoreVersionLocalizationRelationships {
	return &NullableAppStoreVersionLocalizationRelationships{value: val, isSet: true}
}

func (v NullableAppStoreVersionLocalizationRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionLocalizationRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}