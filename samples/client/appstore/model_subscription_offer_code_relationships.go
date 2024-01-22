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

// checks if the SubscriptionOfferCodeRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeRelationships{}

// SubscriptionOfferCodeRelationships struct for SubscriptionOfferCodeRelationships
type SubscriptionOfferCodeRelationships struct {
	Subscription *PromotedPurchaseRelationshipsSubscription `json:"subscription,omitempty"`
	OneTimeUseCodes *SubscriptionOfferCodeRelationshipsOneTimeUseCodes `json:"oneTimeUseCodes,omitempty"`
	CustomCodes *SubscriptionOfferCodeRelationshipsCustomCodes `json:"customCodes,omitempty"`
	Prices *SubscriptionOfferCodeRelationshipsPrices `json:"prices,omitempty"`
}

// NewSubscriptionOfferCodeRelationships instantiates a new SubscriptionOfferCodeRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeRelationships() *SubscriptionOfferCodeRelationships {
	this := SubscriptionOfferCodeRelationships{}
	return &this
}

// NewSubscriptionOfferCodeRelationshipsWithDefaults instantiates a new SubscriptionOfferCodeRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeRelationshipsWithDefaults() *SubscriptionOfferCodeRelationships {
	this := SubscriptionOfferCodeRelationships{}
	return &this
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeRelationships) GetSubscription() PromotedPurchaseRelationshipsSubscription {
	if o == nil || IsNil(o.Subscription) {
		var ret PromotedPurchaseRelationshipsSubscription
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeRelationships) GetSubscriptionOk() (*PromotedPurchaseRelationshipsSubscription, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeRelationships) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given PromotedPurchaseRelationshipsSubscription and assigns it to the Subscription field.
func (o *SubscriptionOfferCodeRelationships) SetSubscription(v PromotedPurchaseRelationshipsSubscription) {
	o.Subscription = &v
}

// GetOneTimeUseCodes returns the OneTimeUseCodes field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeRelationships) GetOneTimeUseCodes() SubscriptionOfferCodeRelationshipsOneTimeUseCodes {
	if o == nil || IsNil(o.OneTimeUseCodes) {
		var ret SubscriptionOfferCodeRelationshipsOneTimeUseCodes
		return ret
	}
	return *o.OneTimeUseCodes
}

// GetOneTimeUseCodesOk returns a tuple with the OneTimeUseCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeRelationships) GetOneTimeUseCodesOk() (*SubscriptionOfferCodeRelationshipsOneTimeUseCodes, bool) {
	if o == nil || IsNil(o.OneTimeUseCodes) {
		return nil, false
	}
	return o.OneTimeUseCodes, true
}

// HasOneTimeUseCodes returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeRelationships) HasOneTimeUseCodes() bool {
	if o != nil && !IsNil(o.OneTimeUseCodes) {
		return true
	}

	return false
}

// SetOneTimeUseCodes gets a reference to the given SubscriptionOfferCodeRelationshipsOneTimeUseCodes and assigns it to the OneTimeUseCodes field.
func (o *SubscriptionOfferCodeRelationships) SetOneTimeUseCodes(v SubscriptionOfferCodeRelationshipsOneTimeUseCodes) {
	o.OneTimeUseCodes = &v
}

// GetCustomCodes returns the CustomCodes field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeRelationships) GetCustomCodes() SubscriptionOfferCodeRelationshipsCustomCodes {
	if o == nil || IsNil(o.CustomCodes) {
		var ret SubscriptionOfferCodeRelationshipsCustomCodes
		return ret
	}
	return *o.CustomCodes
}

// GetCustomCodesOk returns a tuple with the CustomCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeRelationships) GetCustomCodesOk() (*SubscriptionOfferCodeRelationshipsCustomCodes, bool) {
	if o == nil || IsNil(o.CustomCodes) {
		return nil, false
	}
	return o.CustomCodes, true
}

// HasCustomCodes returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeRelationships) HasCustomCodes() bool {
	if o != nil && !IsNil(o.CustomCodes) {
		return true
	}

	return false
}

// SetCustomCodes gets a reference to the given SubscriptionOfferCodeRelationshipsCustomCodes and assigns it to the CustomCodes field.
func (o *SubscriptionOfferCodeRelationships) SetCustomCodes(v SubscriptionOfferCodeRelationshipsCustomCodes) {
	o.CustomCodes = &v
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeRelationships) GetPrices() SubscriptionOfferCodeRelationshipsPrices {
	if o == nil || IsNil(o.Prices) {
		var ret SubscriptionOfferCodeRelationshipsPrices
		return ret
	}
	return *o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeRelationships) GetPricesOk() (*SubscriptionOfferCodeRelationshipsPrices, bool) {
	if o == nil || IsNil(o.Prices) {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeRelationships) HasPrices() bool {
	if o != nil && !IsNil(o.Prices) {
		return true
	}

	return false
}

// SetPrices gets a reference to the given SubscriptionOfferCodeRelationshipsPrices and assigns it to the Prices field.
func (o *SubscriptionOfferCodeRelationships) SetPrices(v SubscriptionOfferCodeRelationshipsPrices) {
	o.Prices = &v
}

func (o SubscriptionOfferCodeRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	if !IsNil(o.OneTimeUseCodes) {
		toSerialize["oneTimeUseCodes"] = o.OneTimeUseCodes
	}
	if !IsNil(o.CustomCodes) {
		toSerialize["customCodes"] = o.CustomCodes
	}
	if !IsNil(o.Prices) {
		toSerialize["prices"] = o.Prices
	}
	return toSerialize, nil
}

type NullableSubscriptionOfferCodeRelationships struct {
	value *SubscriptionOfferCodeRelationships
	isSet bool
}

func (v NullableSubscriptionOfferCodeRelationships) Get() *SubscriptionOfferCodeRelationships {
	return v.value
}

func (v *NullableSubscriptionOfferCodeRelationships) Set(val *SubscriptionOfferCodeRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeRelationships(val *SubscriptionOfferCodeRelationships) *NullableSubscriptionOfferCodeRelationships {
	return &NullableSubscriptionOfferCodeRelationships{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


