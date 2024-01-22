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

// checks if the SubscriptionPromotionalOfferPricesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionPromotionalOfferPricesResponse{}

// SubscriptionPromotionalOfferPricesResponse struct for SubscriptionPromotionalOfferPricesResponse
type SubscriptionPromotionalOfferPricesResponse struct {
	Data []SubscriptionPromotionalOfferPrice `json:"data"`
	Included []SubscriptionOfferCodePricesResponseIncludedInner `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewSubscriptionPromotionalOfferPricesResponse instantiates a new SubscriptionPromotionalOfferPricesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionPromotionalOfferPricesResponse(data []SubscriptionPromotionalOfferPrice, links PagedDocumentLinks) *SubscriptionPromotionalOfferPricesResponse {
	this := SubscriptionPromotionalOfferPricesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewSubscriptionPromotionalOfferPricesResponseWithDefaults instantiates a new SubscriptionPromotionalOfferPricesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionPromotionalOfferPricesResponseWithDefaults() *SubscriptionPromotionalOfferPricesResponse {
	this := SubscriptionPromotionalOfferPricesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionPromotionalOfferPricesResponse) GetData() []SubscriptionPromotionalOfferPrice {
	if o == nil {
		var ret []SubscriptionPromotionalOfferPrice
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferPricesResponse) GetDataOk() ([]SubscriptionPromotionalOfferPrice, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *SubscriptionPromotionalOfferPricesResponse) SetData(v []SubscriptionPromotionalOfferPrice) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *SubscriptionPromotionalOfferPricesResponse) GetIncluded() []SubscriptionOfferCodePricesResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []SubscriptionOfferCodePricesResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferPricesResponse) GetIncludedOk() ([]SubscriptionOfferCodePricesResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *SubscriptionPromotionalOfferPricesResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []SubscriptionOfferCodePricesResponseIncludedInner and assigns it to the Included field.
func (o *SubscriptionPromotionalOfferPricesResponse) SetIncluded(v []SubscriptionOfferCodePricesResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *SubscriptionPromotionalOfferPricesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferPricesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *SubscriptionPromotionalOfferPricesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SubscriptionPromotionalOfferPricesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionPromotionalOfferPricesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SubscriptionPromotionalOfferPricesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *SubscriptionPromotionalOfferPricesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o SubscriptionPromotionalOfferPricesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionPromotionalOfferPricesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableSubscriptionPromotionalOfferPricesResponse struct {
	value *SubscriptionPromotionalOfferPricesResponse
	isSet bool
}

func (v NullableSubscriptionPromotionalOfferPricesResponse) Get() *SubscriptionPromotionalOfferPricesResponse {
	return v.value
}

func (v *NullableSubscriptionPromotionalOfferPricesResponse) Set(val *SubscriptionPromotionalOfferPricesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionPromotionalOfferPricesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionPromotionalOfferPricesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionPromotionalOfferPricesResponse(val *SubscriptionPromotionalOfferPricesResponse) *NullableSubscriptionPromotionalOfferPricesResponse {
	return &NullableSubscriptionPromotionalOfferPricesResponse{value: val, isSet: true}
}

func (v NullableSubscriptionPromotionalOfferPricesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionPromotionalOfferPricesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


