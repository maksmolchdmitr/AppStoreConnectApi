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

// checks if the SubscriptionPromotionalOfferUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPromotionalOfferUpdateRequest{}

// SubscriptionPromotionalOfferUpdateRequest struct for SubscriptionPromotionalOfferUpdateRequest
type SubscriptionPromotionalOfferUpdateRequest struct {
	Data SubscriptionPromotionalOfferUpdateRequestData `json:"data"`
	Included []SubscriptionPromotionalOfferPriceInlineCreate `json:"included,omitempty"`
}

// NewSubscriptionPromotionalOfferUpdateRequest instantiates a new SubscriptionPromotionalOfferUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPromotionalOfferUpdateRequest(data SubscriptionPromotionalOfferUpdateRequestData) *SubscriptionPromotionalOfferUpdateRequest {
	this := SubscriptionPromotionalOfferUpdateRequest{}
	this.Data = data
	return &this
}

// NewSubscriptionPromotionalOfferUpdateRequestWithDefaults instantiates a new SubscriptionPromotionalOfferUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPromotionalOfferUpdateRequestWithDefaults() *SubscriptionPromotionalOfferUpdateRequest {
	this := SubscriptionPromotionalOfferUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionPromotionalOfferUpdateRequest) GetData() SubscriptionPromotionalOfferUpdateRequestData {
	if o == nil {
		var ret SubscriptionPromotionalOfferUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferUpdateRequest) GetDataOk() (*SubscriptionPromotionalOfferUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionPromotionalOfferUpdateRequest) SetData(v SubscriptionPromotionalOfferUpdateRequestData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *SubscriptionPromotionalOfferUpdateRequest) GetIncluded() []SubscriptionPromotionalOfferPriceInlineCreate {
	if o == nil || IsNil(o.Included) {
		var ret []SubscriptionPromotionalOfferPriceInlineCreate
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferUpdateRequest) GetIncludedOk() ([]SubscriptionPromotionalOfferPriceInlineCreate, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *SubscriptionPromotionalOfferUpdateRequest) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []SubscriptionPromotionalOfferPriceInlineCreate and assigns it to the Included field.
func (o *SubscriptionPromotionalOfferUpdateRequest) SetIncluded(v []SubscriptionPromotionalOfferPriceInlineCreate) {
	o.Included = v
}

func (o SubscriptionPromotionalOfferUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPromotionalOfferUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	return toSerialize, nil
}

type NullableSubscriptionPromotionalOfferUpdateRequest struct {
	value *SubscriptionPromotionalOfferUpdateRequest
	isSet bool
}

func (v NullableSubscriptionPromotionalOfferUpdateRequest) Get() *SubscriptionPromotionalOfferUpdateRequest {
	return v.value
}

func (v *NullableSubscriptionPromotionalOfferUpdateRequest) Set(val *SubscriptionPromotionalOfferUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPromotionalOfferUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPromotionalOfferUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPromotionalOfferUpdateRequest(val *SubscriptionPromotionalOfferUpdateRequest) *NullableSubscriptionPromotionalOfferUpdateRequest {
	return &NullableSubscriptionPromotionalOfferUpdateRequest{value: val, isSet: true}
}

func (v NullableSubscriptionPromotionalOfferUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPromotionalOfferUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


