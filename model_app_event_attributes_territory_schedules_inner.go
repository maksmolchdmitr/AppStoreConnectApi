/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the AppEventAttributesTerritorySchedulesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventAttributesTerritorySchedulesInner{}

// AppEventAttributesTerritorySchedulesInner struct for AppEventAttributesTerritorySchedulesInner
type AppEventAttributesTerritorySchedulesInner struct {
	Territories  []string   `json:"territories,omitempty"`
	PublishStart *time.Time `json:"publishStart,omitempty"`
	EventStart   *time.Time `json:"eventStart,omitempty"`
	EventEnd     *time.Time `json:"eventEnd,omitempty"`
}

// NewAppEventAttributesTerritorySchedulesInner instantiates a new AppEventAttributesTerritorySchedulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventAttributesTerritorySchedulesInner() *AppEventAttributesTerritorySchedulesInner {
	this := AppEventAttributesTerritorySchedulesInner{}
	return &this
}

// NewAppEventAttributesTerritorySchedulesInnerWithDefaults instantiates a new AppEventAttributesTerritorySchedulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventAttributesTerritorySchedulesInnerWithDefaults() *AppEventAttributesTerritorySchedulesInner {
	this := AppEventAttributesTerritorySchedulesInner{}
	return &this
}

// GetTerritories returns the Territories field value if set, zero value otherwise.
func (o *AppEventAttributesTerritorySchedulesInner) GetTerritories() []string {
	if o == nil || IsNil(o.Territories) {
		var ret []string
		return ret
	}
	return o.Territories
}

// GetTerritoriesOk returns a tuple with the Territories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventAttributesTerritorySchedulesInner) GetTerritoriesOk() ([]string, bool) {
	if o == nil || IsNil(o.Territories) {
		return nil, false
	}
	return o.Territories, true
}

// HasTerritories returns a boolean if a field has been set.
func (o *AppEventAttributesTerritorySchedulesInner) HasTerritories() bool {
	if o != nil && !IsNil(o.Territories) {
		return true
	}

	return false
}

// SetTerritories gets a reference to the given []string and assigns it to the Territories field.
func (o *AppEventAttributesTerritorySchedulesInner) SetTerritories(v []string) {
	o.Territories = v
}

// GetPublishStart returns the PublishStart field value if set, zero value otherwise.
func (o *AppEventAttributesTerritorySchedulesInner) GetPublishStart() time.Time {
	if o == nil || IsNil(o.PublishStart) {
		var ret time.Time
		return ret
	}
	return *o.PublishStart
}

// GetPublishStartOk returns a tuple with the PublishStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventAttributesTerritorySchedulesInner) GetPublishStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PublishStart) {
		return nil, false
	}
	return o.PublishStart, true
}

// HasPublishStart returns a boolean if a field has been set.
func (o *AppEventAttributesTerritorySchedulesInner) HasPublishStart() bool {
	if o != nil && !IsNil(o.PublishStart) {
		return true
	}

	return false
}

// SetPublishStart gets a reference to the given time.Time and assigns it to the PublishStart field.
func (o *AppEventAttributesTerritorySchedulesInner) SetPublishStart(v time.Time) {
	o.PublishStart = &v
}

// GetEventStart returns the EventStart field value if set, zero value otherwise.
func (o *AppEventAttributesTerritorySchedulesInner) GetEventStart() time.Time {
	if o == nil || IsNil(o.EventStart) {
		var ret time.Time
		return ret
	}
	return *o.EventStart
}

// GetEventStartOk returns a tuple with the EventStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventAttributesTerritorySchedulesInner) GetEventStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EventStart) {
		return nil, false
	}
	return o.EventStart, true
}

// HasEventStart returns a boolean if a field has been set.
func (o *AppEventAttributesTerritorySchedulesInner) HasEventStart() bool {
	if o != nil && !IsNil(o.EventStart) {
		return true
	}

	return false
}

// SetEventStart gets a reference to the given time.Time and assigns it to the EventStart field.
func (o *AppEventAttributesTerritorySchedulesInner) SetEventStart(v time.Time) {
	o.EventStart = &v
}

// GetEventEnd returns the EventEnd field value if set, zero value otherwise.
func (o *AppEventAttributesTerritorySchedulesInner) GetEventEnd() time.Time {
	if o == nil || IsNil(o.EventEnd) {
		var ret time.Time
		return ret
	}
	return *o.EventEnd
}

// GetEventEndOk returns a tuple with the EventEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventAttributesTerritorySchedulesInner) GetEventEndOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EventEnd) {
		return nil, false
	}
	return o.EventEnd, true
}

// HasEventEnd returns a boolean if a field has been set.
func (o *AppEventAttributesTerritorySchedulesInner) HasEventEnd() bool {
	if o != nil && !IsNil(o.EventEnd) {
		return true
	}

	return false
}

// SetEventEnd gets a reference to the given time.Time and assigns it to the EventEnd field.
func (o *AppEventAttributesTerritorySchedulesInner) SetEventEnd(v time.Time) {
	o.EventEnd = &v
}

func (o AppEventAttributesTerritorySchedulesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventAttributesTerritorySchedulesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Territories) {
		toSerialize["territories"] = o.Territories
	}
	if !IsNil(o.PublishStart) {
		toSerialize["publishStart"] = o.PublishStart
	}
	if !IsNil(o.EventStart) {
		toSerialize["eventStart"] = o.EventStart
	}
	if !IsNil(o.EventEnd) {
		toSerialize["eventEnd"] = o.EventEnd
	}
	return toSerialize, nil
}

type NullableAppEventAttributesTerritorySchedulesInner struct {
	value *AppEventAttributesTerritorySchedulesInner
	isSet bool
}

func (v NullableAppEventAttributesTerritorySchedulesInner) Get() *AppEventAttributesTerritorySchedulesInner {
	return v.value
}

func (v *NullableAppEventAttributesTerritorySchedulesInner) Set(val *AppEventAttributesTerritorySchedulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventAttributesTerritorySchedulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventAttributesTerritorySchedulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventAttributesTerritorySchedulesInner(val *AppEventAttributesTerritorySchedulesInner) *NullableAppEventAttributesTerritorySchedulesInner {
	return &NullableAppEventAttributesTerritorySchedulesInner{value: val, isSet: true}
}

func (v NullableAppEventAttributesTerritorySchedulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventAttributesTerritorySchedulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}