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

// checks if the AppPreviewSetRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPreviewSetRelationships{}

// AppPreviewSetRelationships struct for AppPreviewSetRelationships
type AppPreviewSetRelationships struct {
	AppStoreVersionLocalization                    *AppPreviewSetRelationshipsAppStoreVersionLocalization                    `json:"appStoreVersionLocalization,omitempty"`
	AppCustomProductPageLocalization               *AppPreviewSetRelationshipsAppCustomProductPageLocalization               `json:"appCustomProductPageLocalization,omitempty"`
	AppStoreVersionExperimentTreatmentLocalization *AppPreviewSetRelationshipsAppStoreVersionExperimentTreatmentLocalization `json:"appStoreVersionExperimentTreatmentLocalization,omitempty"`
	AppPreviews                                    *AppPreviewSetRelationshipsAppPreviews                                    `json:"appPreviews,omitempty"`
}

// NewAppPreviewSetRelationships instantiates a new AppPreviewSetRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreviewSetRelationships() *AppPreviewSetRelationships {
	this := AppPreviewSetRelationships{}
	return &this
}

// NewAppPreviewSetRelationshipsWithDefaults instantiates a new AppPreviewSetRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreviewSetRelationshipsWithDefaults() *AppPreviewSetRelationships {
	this := AppPreviewSetRelationships{}
	return &this
}

// GetAppStoreVersionLocalization returns the AppStoreVersionLocalization field value if set, zero value otherwise.
func (o *AppPreviewSetRelationships) GetAppStoreVersionLocalization() AppPreviewSetRelationshipsAppStoreVersionLocalization {
	if o == nil || IsNil(o.AppStoreVersionLocalization) {
		var ret AppPreviewSetRelationshipsAppStoreVersionLocalization
		return ret
	}
	return *o.AppStoreVersionLocalization
}

// GetAppStoreVersionLocalizationOk returns a tuple with the AppStoreVersionLocalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewSetRelationships) GetAppStoreVersionLocalizationOk() (*AppPreviewSetRelationshipsAppStoreVersionLocalization, bool) {
	if o == nil || IsNil(o.AppStoreVersionLocalization) {
		return nil, false
	}
	return o.AppStoreVersionLocalization, true
}

// HasAppStoreVersionLocalization returns a boolean if a field has been set.
func (o *AppPreviewSetRelationships) HasAppStoreVersionLocalization() bool {
	if o != nil && !IsNil(o.AppStoreVersionLocalization) {
		return true
	}

	return false
}

// SetAppStoreVersionLocalization gets a reference to the given AppPreviewSetRelationshipsAppStoreVersionLocalization and assigns it to the AppStoreVersionLocalization field.
func (o *AppPreviewSetRelationships) SetAppStoreVersionLocalization(v AppPreviewSetRelationshipsAppStoreVersionLocalization) {
	o.AppStoreVersionLocalization = &v
}

// GetAppCustomProductPageLocalization returns the AppCustomProductPageLocalization field value if set, zero value otherwise.
func (o *AppPreviewSetRelationships) GetAppCustomProductPageLocalization() AppPreviewSetRelationshipsAppCustomProductPageLocalization {
	if o == nil || IsNil(o.AppCustomProductPageLocalization) {
		var ret AppPreviewSetRelationshipsAppCustomProductPageLocalization
		return ret
	}
	return *o.AppCustomProductPageLocalization
}

// GetAppCustomProductPageLocalizationOk returns a tuple with the AppCustomProductPageLocalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewSetRelationships) GetAppCustomProductPageLocalizationOk() (*AppPreviewSetRelationshipsAppCustomProductPageLocalization, bool) {
	if o == nil || IsNil(o.AppCustomProductPageLocalization) {
		return nil, false
	}
	return o.AppCustomProductPageLocalization, true
}

// HasAppCustomProductPageLocalization returns a boolean if a field has been set.
func (o *AppPreviewSetRelationships) HasAppCustomProductPageLocalization() bool {
	if o != nil && !IsNil(o.AppCustomProductPageLocalization) {
		return true
	}

	return false
}

// SetAppCustomProductPageLocalization gets a reference to the given AppPreviewSetRelationshipsAppCustomProductPageLocalization and assigns it to the AppCustomProductPageLocalization field.
func (o *AppPreviewSetRelationships) SetAppCustomProductPageLocalization(v AppPreviewSetRelationshipsAppCustomProductPageLocalization) {
	o.AppCustomProductPageLocalization = &v
}

// GetAppStoreVersionExperimentTreatmentLocalization returns the AppStoreVersionExperimentTreatmentLocalization field value if set, zero value otherwise.
func (o *AppPreviewSetRelationships) GetAppStoreVersionExperimentTreatmentLocalization() AppPreviewSetRelationshipsAppStoreVersionExperimentTreatmentLocalization {
	if o == nil || IsNil(o.AppStoreVersionExperimentTreatmentLocalization) {
		var ret AppPreviewSetRelationshipsAppStoreVersionExperimentTreatmentLocalization
		return ret
	}
	return *o.AppStoreVersionExperimentTreatmentLocalization
}

// GetAppStoreVersionExperimentTreatmentLocalizationOk returns a tuple with the AppStoreVersionExperimentTreatmentLocalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewSetRelationships) GetAppStoreVersionExperimentTreatmentLocalizationOk() (*AppPreviewSetRelationshipsAppStoreVersionExperimentTreatmentLocalization, bool) {
	if o == nil || IsNil(o.AppStoreVersionExperimentTreatmentLocalization) {
		return nil, false
	}
	return o.AppStoreVersionExperimentTreatmentLocalization, true
}

// HasAppStoreVersionExperimentTreatmentLocalization returns a boolean if a field has been set.
func (o *AppPreviewSetRelationships) HasAppStoreVersionExperimentTreatmentLocalization() bool {
	if o != nil && !IsNil(o.AppStoreVersionExperimentTreatmentLocalization) {
		return true
	}

	return false
}

// SetAppStoreVersionExperimentTreatmentLocalization gets a reference to the given AppPreviewSetRelationshipsAppStoreVersionExperimentTreatmentLocalization and assigns it to the AppStoreVersionExperimentTreatmentLocalization field.
func (o *AppPreviewSetRelationships) SetAppStoreVersionExperimentTreatmentLocalization(v AppPreviewSetRelationshipsAppStoreVersionExperimentTreatmentLocalization) {
	o.AppStoreVersionExperimentTreatmentLocalization = &v
}

// GetAppPreviews returns the AppPreviews field value if set, zero value otherwise.
func (o *AppPreviewSetRelationships) GetAppPreviews() AppPreviewSetRelationshipsAppPreviews {
	if o == nil || IsNil(o.AppPreviews) {
		var ret AppPreviewSetRelationshipsAppPreviews
		return ret
	}
	return *o.AppPreviews
}

// GetAppPreviewsOk returns a tuple with the AppPreviews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewSetRelationships) GetAppPreviewsOk() (*AppPreviewSetRelationshipsAppPreviews, bool) {
	if o == nil || IsNil(o.AppPreviews) {
		return nil, false
	}
	return o.AppPreviews, true
}

// HasAppPreviews returns a boolean if a field has been set.
func (o *AppPreviewSetRelationships) HasAppPreviews() bool {
	if o != nil && !IsNil(o.AppPreviews) {
		return true
	}

	return false
}

// SetAppPreviews gets a reference to the given AppPreviewSetRelationshipsAppPreviews and assigns it to the AppPreviews field.
func (o *AppPreviewSetRelationships) SetAppPreviews(v AppPreviewSetRelationshipsAppPreviews) {
	o.AppPreviews = &v
}

func (o AppPreviewSetRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPreviewSetRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppStoreVersionLocalization) {
		toSerialize["appStoreVersionLocalization"] = o.AppStoreVersionLocalization
	}
	if !IsNil(o.AppCustomProductPageLocalization) {
		toSerialize["appCustomProductPageLocalization"] = o.AppCustomProductPageLocalization
	}
	if !IsNil(o.AppStoreVersionExperimentTreatmentLocalization) {
		toSerialize["appStoreVersionExperimentTreatmentLocalization"] = o.AppStoreVersionExperimentTreatmentLocalization
	}
	if !IsNil(o.AppPreviews) {
		toSerialize["appPreviews"] = o.AppPreviews
	}
	return toSerialize, nil
}

type NullableAppPreviewSetRelationships struct {
	value *AppPreviewSetRelationships
	isSet bool
}

func (v NullableAppPreviewSetRelationships) Get() *AppPreviewSetRelationships {
	return v.value
}

func (v *NullableAppPreviewSetRelationships) Set(val *AppPreviewSetRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewSetRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewSetRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewSetRelationships(val *AppPreviewSetRelationships) *NullableAppPreviewSetRelationships {
	return &NullableAppPreviewSetRelationships{value: val, isSet: true}
}

func (v NullableAppPreviewSetRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewSetRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
