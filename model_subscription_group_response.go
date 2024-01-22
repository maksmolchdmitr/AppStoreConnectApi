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

// checks if the SubscriptionGroupResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionGroupResponse{}

// SubscriptionGroupResponse struct for SubscriptionGroupResponse
type SubscriptionGroupResponse struct {
	Data     SubscriptionGroup                         `json:"data"`
	Included []SubscriptionGroupsResponseIncludedInner `json:"included,omitempty"`
	Links    DocumentLinks                             `json:"links"`
}

// NewSubscriptionGroupResponse instantiates a new SubscriptionGroupResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionGroupResponse(data SubscriptionGroup, links DocumentLinks) *SubscriptionGroupResponse {
	this := SubscriptionGroupResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewSubscriptionGroupResponseWithDefaults instantiates a new SubscriptionGroupResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionGroupResponseWithDefaults() *SubscriptionGroupResponse {
	this := SubscriptionGroupResponse{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionGroupResponse) GetData() SubscriptionGroup {
	if o == nil {
		var ret SubscriptionGroup
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupResponse) GetDataOk() (*SubscriptionGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionGroupResponse) SetData(v SubscriptionGroup) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *SubscriptionGroupResponse) GetIncluded() []SubscriptionGroupsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []SubscriptionGroupsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupResponse) GetIncludedOk() ([]SubscriptionGroupsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *SubscriptionGroupResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []SubscriptionGroupsResponseIncludedInner and assigns it to the Included field.
func (o *SubscriptionGroupResponse) SetIncluded(v []SubscriptionGroupsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *SubscriptionGroupResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SubscriptionGroupResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o SubscriptionGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionGroupResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableSubscriptionGroupResponse struct {
	value *SubscriptionGroupResponse
	isSet bool
}

func (v NullableSubscriptionGroupResponse) Get() *SubscriptionGroupResponse {
	return v.value
}

func (v *NullableSubscriptionGroupResponse) Set(val *SubscriptionGroupResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionGroupResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionGroupResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionGroupResponse(val *SubscriptionGroupResponse) *NullableSubscriptionGroupResponse {
	return &NullableSubscriptionGroupResponse{value: val, isSet: true}
}

func (v NullableSubscriptionGroupResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionGroupResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
