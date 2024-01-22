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

// checks if the SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode{}

// SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode struct for SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode
type SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode struct {
	Data SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData `json:"data"`
}

// NewSubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode instantiates a new SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode(data SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData) *SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode {
	this := SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode{}
	this.Data = data
	return &this
}

// NewSubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCodeWithDefaults instantiates a new SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCodeWithDefaults() *SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode {
	this := SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode) GetData() SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData {
	if o == nil {
		var ret SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode) GetDataOk() (*SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode) SetData(v SubscriptionOfferCodeCustomCodeRelationshipsOfferCodeData) {
	o.Data = v
}

func (o SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableSubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode struct {
	value *SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode
	isSet bool
}

func (v NullableSubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode) Get() *SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode {
	return v.value
}

func (v *NullableSubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode) Set(val *SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode(val *SubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode) *NullableSubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode {
	return &NullableSubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeCustomCodeCreateRequestDataRelationshipsOfferCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

