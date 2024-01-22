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

// checks if the BetaAppLocalizationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppLocalizationResponse{}

// BetaAppLocalizationResponse struct for BetaAppLocalizationResponse
type BetaAppLocalizationResponse struct {
	Data     BetaAppLocalization `json:"data"`
	Included []App               `json:"included,omitempty"`
	Links    DocumentLinks       `json:"links"`
}

// NewBetaAppLocalizationResponse instantiates a new BetaAppLocalizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppLocalizationResponse(data BetaAppLocalization, links DocumentLinks) *BetaAppLocalizationResponse {
	this := BetaAppLocalizationResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBetaAppLocalizationResponseWithDefaults instantiates a new BetaAppLocalizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppLocalizationResponseWithDefaults() *BetaAppLocalizationResponse {
	this := BetaAppLocalizationResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BetaAppLocalizationResponse) GetData() BetaAppLocalization {
	if o == nil {
		var ret BetaAppLocalization
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaAppLocalizationResponse) GetDataOk() (*BetaAppLocalization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetaAppLocalizationResponse) SetData(v BetaAppLocalization) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *BetaAppLocalizationResponse) GetIncluded() []App {
	if o == nil || IsNil(o.Included) {
		var ret []App
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppLocalizationResponse) GetIncludedOk() ([]App, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *BetaAppLocalizationResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []App and assigns it to the Included field.
func (o *BetaAppLocalizationResponse) SetIncluded(v []App) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *BetaAppLocalizationResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BetaAppLocalizationResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BetaAppLocalizationResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o BetaAppLocalizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppLocalizationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableBetaAppLocalizationResponse struct {
	value *BetaAppLocalizationResponse
	isSet bool
}

func (v NullableBetaAppLocalizationResponse) Get() *BetaAppLocalizationResponse {
	return v.value
}

func (v *NullableBetaAppLocalizationResponse) Set(val *BetaAppLocalizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppLocalizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppLocalizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppLocalizationResponse(val *BetaAppLocalizationResponse) *NullableBetaAppLocalizationResponse {
	return &NullableBetaAppLocalizationResponse{value: val, isSet: true}
}

func (v NullableBetaAppLocalizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppLocalizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
