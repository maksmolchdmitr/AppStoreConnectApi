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

// checks if the BetaGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaGroupResponse{}

// BetaGroupResponse struct for BetaGroupResponse
type BetaGroupResponse struct {
	Data     BetaGroup                         `json:"data"`
	Included []BetaGroupsResponseIncludedInner `json:"included,omitempty"`
	Links    DocumentLinks                     `json:"links"`
}

// NewBetaGroupResponse instantiates a new BetaGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaGroupResponse(data BetaGroup, links DocumentLinks) *BetaGroupResponse {
	this := BetaGroupResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBetaGroupResponseWithDefaults instantiates a new BetaGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaGroupResponseWithDefaults() *BetaGroupResponse {
	this := BetaGroupResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BetaGroupResponse) GetData() BetaGroup {
	if o == nil {
		var ret BetaGroup
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaGroupResponse) GetDataOk() (*BetaGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetaGroupResponse) SetData(v BetaGroup) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *BetaGroupResponse) GetIncluded() []BetaGroupsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []BetaGroupsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupResponse) GetIncludedOk() ([]BetaGroupsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *BetaGroupResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []BetaGroupsResponseIncludedInner and assigns it to the Included field.
func (o *BetaGroupResponse) SetIncluded(v []BetaGroupsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *BetaGroupResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BetaGroupResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BetaGroupResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o BetaGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableBetaGroupResponse struct {
	value *BetaGroupResponse
	isSet bool
}

func (v NullableBetaGroupResponse) Get() *BetaGroupResponse {
	return v.value
}

func (v *NullableBetaGroupResponse) Set(val *BetaGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaGroupResponse(val *BetaGroupResponse) *NullableBetaGroupResponse {
	return &NullableBetaGroupResponse{value: val, isSet: true}
}

func (v NullableBetaGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
