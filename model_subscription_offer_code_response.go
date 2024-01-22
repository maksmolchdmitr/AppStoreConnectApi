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

// checks if the SubscriptionOfferCodeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeResponse{}

// SubscriptionOfferCodeResponse struct for SubscriptionOfferCodeResponse
type SubscriptionOfferCodeResponse struct {
	Data     SubscriptionOfferCode                         `json:"data"`
	Included []SubscriptionOfferCodesResponseIncludedInner `json:"included,omitempty"`
	Links    DocumentLinks                                 `json:"links"`
}

// NewSubscriptionOfferCodeResponse instantiates a new SubscriptionOfferCodeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeResponse(data SubscriptionOfferCode, links DocumentLinks) *SubscriptionOfferCodeResponse {
	this := SubscriptionOfferCodeResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewSubscriptionOfferCodeResponseWithDefaults instantiates a new SubscriptionOfferCodeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeResponseWithDefaults() *SubscriptionOfferCodeResponse {
	this := SubscriptionOfferCodeResponse{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionOfferCodeResponse) GetData() SubscriptionOfferCode {
	if o == nil {
		var ret SubscriptionOfferCode
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeResponse) GetDataOk() (*SubscriptionOfferCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionOfferCodeResponse) SetData(v SubscriptionOfferCode) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeResponse) GetIncluded() []SubscriptionOfferCodesResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []SubscriptionOfferCodesResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeResponse) GetIncludedOk() ([]SubscriptionOfferCodesResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []SubscriptionOfferCodesResponseIncludedInner and assigns it to the Included field.
func (o *SubscriptionOfferCodeResponse) SetIncluded(v []SubscriptionOfferCodesResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *SubscriptionOfferCodeResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SubscriptionOfferCodeResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o SubscriptionOfferCodeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableSubscriptionOfferCodeResponse struct {
	value *SubscriptionOfferCodeResponse
	isSet bool
}

func (v NullableSubscriptionOfferCodeResponse) Get() *SubscriptionOfferCodeResponse {
	return v.value
}

func (v *NullableSubscriptionOfferCodeResponse) Set(val *SubscriptionOfferCodeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeResponse(val *SubscriptionOfferCodeResponse) *NullableSubscriptionOfferCodeResponse {
	return &NullableSubscriptionOfferCodeResponse{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
