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

// checks if the BetaTesterInvitationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaTesterInvitationResponse{}

// BetaTesterInvitationResponse struct for BetaTesterInvitationResponse
type BetaTesterInvitationResponse struct {
	Data  BetaTesterInvitation `json:"data"`
	Links DocumentLinks        `json:"links"`
}

// NewBetaTesterInvitationResponse instantiates a new BetaTesterInvitationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaTesterInvitationResponse(data BetaTesterInvitation, links DocumentLinks) *BetaTesterInvitationResponse {
	this := BetaTesterInvitationResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBetaTesterInvitationResponseWithDefaults instantiates a new BetaTesterInvitationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaTesterInvitationResponseWithDefaults() *BetaTesterInvitationResponse {
	this := BetaTesterInvitationResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BetaTesterInvitationResponse) GetData() BetaTesterInvitation {
	if o == nil {
		var ret BetaTesterInvitation
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaTesterInvitationResponse) GetDataOk() (*BetaTesterInvitation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetaTesterInvitationResponse) SetData(v BetaTesterInvitation) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *BetaTesterInvitationResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BetaTesterInvitationResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BetaTesterInvitationResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o BetaTesterInvitationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaTesterInvitationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableBetaTesterInvitationResponse struct {
	value *BetaTesterInvitationResponse
	isSet bool
}

func (v NullableBetaTesterInvitationResponse) Get() *BetaTesterInvitationResponse {
	return v.value
}

func (v *NullableBetaTesterInvitationResponse) Set(val *BetaTesterInvitationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaTesterInvitationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaTesterInvitationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaTesterInvitationResponse(val *BetaTesterInvitationResponse) *NullableBetaTesterInvitationResponse {
	return &NullableBetaTesterInvitationResponse{value: val, isSet: true}
}

func (v NullableBetaTesterInvitationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaTesterInvitationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
