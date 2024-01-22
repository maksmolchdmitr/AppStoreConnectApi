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

// checks if the SubscriptionOfferCodeCustomCodeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeCustomCodeResponse{}

// SubscriptionOfferCodeCustomCodeResponse struct for SubscriptionOfferCodeCustomCodeResponse
type SubscriptionOfferCodeCustomCodeResponse struct {
	Data     SubscriptionOfferCodeCustomCode `json:"data"`
	Included []SubscriptionOfferCode         `json:"included,omitempty"`
	Links    DocumentLinks                   `json:"links"`
}

// NewSubscriptionOfferCodeCustomCodeResponse instantiates a new SubscriptionOfferCodeCustomCodeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeCustomCodeResponse(data SubscriptionOfferCodeCustomCode, links DocumentLinks) *SubscriptionOfferCodeCustomCodeResponse {
	this := SubscriptionOfferCodeCustomCodeResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewSubscriptionOfferCodeCustomCodeResponseWithDefaults instantiates a new SubscriptionOfferCodeCustomCodeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeCustomCodeResponseWithDefaults() *SubscriptionOfferCodeCustomCodeResponse {
	this := SubscriptionOfferCodeCustomCodeResponse{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionOfferCodeCustomCodeResponse) GetData() SubscriptionOfferCodeCustomCode {
	if o == nil {
		var ret SubscriptionOfferCodeCustomCode
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCustomCodeResponse) GetDataOk() (*SubscriptionOfferCodeCustomCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionOfferCodeCustomCodeResponse) SetData(v SubscriptionOfferCodeCustomCode) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeCustomCodeResponse) GetIncluded() []SubscriptionOfferCode {
	if o == nil || IsNil(o.Included) {
		var ret []SubscriptionOfferCode
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCustomCodeResponse) GetIncludedOk() ([]SubscriptionOfferCode, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeCustomCodeResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []SubscriptionOfferCode and assigns it to the Included field.
func (o *SubscriptionOfferCodeCustomCodeResponse) SetIncluded(v []SubscriptionOfferCode) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *SubscriptionOfferCodeCustomCodeResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCustomCodeResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SubscriptionOfferCodeCustomCodeResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o SubscriptionOfferCodeCustomCodeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeCustomCodeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableSubscriptionOfferCodeCustomCodeResponse struct {
	value *SubscriptionOfferCodeCustomCodeResponse
	isSet bool
}

func (v NullableSubscriptionOfferCodeCustomCodeResponse) Get() *SubscriptionOfferCodeCustomCodeResponse {
	return v.value
}

func (v *NullableSubscriptionOfferCodeCustomCodeResponse) Set(val *SubscriptionOfferCodeCustomCodeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeCustomCodeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeCustomCodeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeCustomCodeResponse(val *SubscriptionOfferCodeCustomCodeResponse) *NullableSubscriptionOfferCodeCustomCodeResponse {
	return &NullableSubscriptionOfferCodeCustomCodeResponse{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeCustomCodeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeCustomCodeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
