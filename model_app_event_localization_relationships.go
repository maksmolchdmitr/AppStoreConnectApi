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

// checks if the AppEventLocalizationRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventLocalizationRelationships{}

// AppEventLocalizationRelationships struct for AppEventLocalizationRelationships
type AppEventLocalizationRelationships struct {
	AppEvent            *AppEventLocalizationRelationshipsAppEvent            `json:"appEvent,omitempty"`
	AppEventScreenshots *AppEventLocalizationRelationshipsAppEventScreenshots `json:"appEventScreenshots,omitempty"`
	AppEventVideoClips  *AppEventLocalizationRelationshipsAppEventVideoClips  `json:"appEventVideoClips,omitempty"`
}

// NewAppEventLocalizationRelationships instantiates a new AppEventLocalizationRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventLocalizationRelationships() *AppEventLocalizationRelationships {
	this := AppEventLocalizationRelationships{}
	return &this
}

// NewAppEventLocalizationRelationshipsWithDefaults instantiates a new AppEventLocalizationRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventLocalizationRelationshipsWithDefaults() *AppEventLocalizationRelationships {
	this := AppEventLocalizationRelationships{}
	return &this
}

// GetAppEvent returns the AppEvent field value if set, zero value otherwise.
func (o *AppEventLocalizationRelationships) GetAppEvent() AppEventLocalizationRelationshipsAppEvent {
	if o == nil || IsNil(o.AppEvent) {
		var ret AppEventLocalizationRelationshipsAppEvent
		return ret
	}
	return *o.AppEvent
}

// GetAppEventOk returns a tuple with the AppEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationRelationships) GetAppEventOk() (*AppEventLocalizationRelationshipsAppEvent, bool) {
	if o == nil || IsNil(o.AppEvent) {
		return nil, false
	}
	return o.AppEvent, true
}

// HasAppEvent returns a boolean if a field has been set.
func (o *AppEventLocalizationRelationships) HasAppEvent() bool {
	if o != nil && !IsNil(o.AppEvent) {
		return true
	}

	return false
}

// SetAppEvent gets a reference to the given AppEventLocalizationRelationshipsAppEvent and assigns it to the AppEvent field.
func (o *AppEventLocalizationRelationships) SetAppEvent(v AppEventLocalizationRelationshipsAppEvent) {
	o.AppEvent = &v
}

// GetAppEventScreenshots returns the AppEventScreenshots field value if set, zero value otherwise.
func (o *AppEventLocalizationRelationships) GetAppEventScreenshots() AppEventLocalizationRelationshipsAppEventScreenshots {
	if o == nil || IsNil(o.AppEventScreenshots) {
		var ret AppEventLocalizationRelationshipsAppEventScreenshots
		return ret
	}
	return *o.AppEventScreenshots
}

// GetAppEventScreenshotsOk returns a tuple with the AppEventScreenshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationRelationships) GetAppEventScreenshotsOk() (*AppEventLocalizationRelationshipsAppEventScreenshots, bool) {
	if o == nil || IsNil(o.AppEventScreenshots) {
		return nil, false
	}
	return o.AppEventScreenshots, true
}

// HasAppEventScreenshots returns a boolean if a field has been set.
func (o *AppEventLocalizationRelationships) HasAppEventScreenshots() bool {
	if o != nil && !IsNil(o.AppEventScreenshots) {
		return true
	}

	return false
}

// SetAppEventScreenshots gets a reference to the given AppEventLocalizationRelationshipsAppEventScreenshots and assigns it to the AppEventScreenshots field.
func (o *AppEventLocalizationRelationships) SetAppEventScreenshots(v AppEventLocalizationRelationshipsAppEventScreenshots) {
	o.AppEventScreenshots = &v
}

// GetAppEventVideoClips returns the AppEventVideoClips field value if set, zero value otherwise.
func (o *AppEventLocalizationRelationships) GetAppEventVideoClips() AppEventLocalizationRelationshipsAppEventVideoClips {
	if o == nil || IsNil(o.AppEventVideoClips) {
		var ret AppEventLocalizationRelationshipsAppEventVideoClips
		return ret
	}
	return *o.AppEventVideoClips
}

// GetAppEventVideoClipsOk returns a tuple with the AppEventVideoClips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventLocalizationRelationships) GetAppEventVideoClipsOk() (*AppEventLocalizationRelationshipsAppEventVideoClips, bool) {
	if o == nil || IsNil(o.AppEventVideoClips) {
		return nil, false
	}
	return o.AppEventVideoClips, true
}

// HasAppEventVideoClips returns a boolean if a field has been set.
func (o *AppEventLocalizationRelationships) HasAppEventVideoClips() bool {
	if o != nil && !IsNil(o.AppEventVideoClips) {
		return true
	}

	return false
}

// SetAppEventVideoClips gets a reference to the given AppEventLocalizationRelationshipsAppEventVideoClips and assigns it to the AppEventVideoClips field.
func (o *AppEventLocalizationRelationships) SetAppEventVideoClips(v AppEventLocalizationRelationshipsAppEventVideoClips) {
	o.AppEventVideoClips = &v
}

func (o AppEventLocalizationRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventLocalizationRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppEvent) {
		toSerialize["appEvent"] = o.AppEvent
	}
	if !IsNil(o.AppEventScreenshots) {
		toSerialize["appEventScreenshots"] = o.AppEventScreenshots
	}
	if !IsNil(o.AppEventVideoClips) {
		toSerialize["appEventVideoClips"] = o.AppEventVideoClips
	}
	return toSerialize, nil
}

type NullableAppEventLocalizationRelationships struct {
	value *AppEventLocalizationRelationships
	isSet bool
}

func (v NullableAppEventLocalizationRelationships) Get() *AppEventLocalizationRelationships {
	return v.value
}

func (v *NullableAppEventLocalizationRelationships) Set(val *AppEventLocalizationRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventLocalizationRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventLocalizationRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventLocalizationRelationships(val *AppEventLocalizationRelationships) *NullableAppEventLocalizationRelationships {
	return &NullableAppEventLocalizationRelationships{value: val, isSet: true}
}

func (v NullableAppEventLocalizationRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventLocalizationRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
