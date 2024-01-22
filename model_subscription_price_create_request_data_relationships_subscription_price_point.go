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

// checks if the SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint{}

// SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint struct for SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint
type SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint struct {
	Data SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePointData `json:"data"`
}

// NewSubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint instantiates a new SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint(data SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePointData) *SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint {
	this := SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint{}
	this.Data = data
	return &this
}

// NewSubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePointWithDefaults instantiates a new SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePointWithDefaults() *SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint {
	this := SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint) GetData() SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePointData {
	if o == nil {
		var ret SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePointData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint) GetDataOk() (*SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePointData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint) SetData(v SubscriptionIntroductoryOfferRelationshipsSubscriptionPricePointData) {
	o.Data = v
}

func (o SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableSubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint struct {
	value *SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint
	isSet bool
}

func (v NullableSubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint) Get() *SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint {
	return v.value
}

func (v *NullableSubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint) Set(val *SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint(val *SubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint) *NullableSubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint {
	return &NullableSubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint{value: val, isSet: true}
}

func (v NullableSubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPriceCreateRequestDataRelationshipsSubscriptionPricePoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
