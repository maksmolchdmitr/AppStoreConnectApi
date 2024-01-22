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

// checks if the AppEventResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventResponse{}

// AppEventResponse struct for AppEventResponse
type AppEventResponse struct {
	Data     AppEvent               `json:"data"`
	Included []AppEventLocalization `json:"included,omitempty"`
	Links    DocumentLinks          `json:"links"`
}

// NewAppEventResponse instantiates a new AppEventResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventResponse(data AppEvent, links DocumentLinks) *AppEventResponse {
	this := AppEventResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppEventResponseWithDefaults instantiates a new AppEventResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventResponseWithDefaults() *AppEventResponse {
	this := AppEventResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppEventResponse) GetData() AppEvent {
	if o == nil {
		var ret AppEvent
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppEventResponse) GetDataOk() (*AppEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppEventResponse) SetData(v AppEvent) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppEventResponse) GetIncluded() []AppEventLocalization {
	if o == nil || IsNil(o.Included) {
		var ret []AppEventLocalization
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventResponse) GetIncludedOk() ([]AppEventLocalization, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppEventResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppEventLocalization and assigns it to the Included field.
func (o *AppEventResponse) SetIncluded(v []AppEventLocalization) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppEventResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppEventResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppEventResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppEventResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableAppEventResponse struct {
	value *AppEventResponse
	isSet bool
}

func (v NullableAppEventResponse) Get() *AppEventResponse {
	return v.value
}

func (v *NullableAppEventResponse) Set(val *AppEventResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventResponse(val *AppEventResponse) *NullableAppEventResponse {
	return &NullableAppEventResponse{value: val, isSet: true}
}

func (v NullableAppEventResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
