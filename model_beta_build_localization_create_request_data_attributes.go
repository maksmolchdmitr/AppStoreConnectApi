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

// checks if the BetaBuildLocalizationCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaBuildLocalizationCreateRequestDataAttributes{}

// BetaBuildLocalizationCreateRequestDataAttributes struct for BetaBuildLocalizationCreateRequestDataAttributes
type BetaBuildLocalizationCreateRequestDataAttributes struct {
	WhatsNew *string `json:"whatsNew,omitempty"`
	Locale   string  `json:"locale"`
}

// NewBetaBuildLocalizationCreateRequestDataAttributes instantiates a new BetaBuildLocalizationCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaBuildLocalizationCreateRequestDataAttributes(locale string) *BetaBuildLocalizationCreateRequestDataAttributes {
	this := BetaBuildLocalizationCreateRequestDataAttributes{}
	this.Locale = locale
	return &this
}

// NewBetaBuildLocalizationCreateRequestDataAttributesWithDefaults instantiates a new BetaBuildLocalizationCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaBuildLocalizationCreateRequestDataAttributesWithDefaults() *BetaBuildLocalizationCreateRequestDataAttributes {
	this := BetaBuildLocalizationCreateRequestDataAttributes{}
	return &this
}

// GetWhatsNew returns the WhatsNew field value if set, zero value otherwise.
func (o *BetaBuildLocalizationCreateRequestDataAttributes) GetWhatsNew() string {
	if o == nil || IsNil(o.WhatsNew) {
		var ret string
		return ret
	}
	return *o.WhatsNew
}

// GetWhatsNewOk returns a tuple with the WhatsNew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaBuildLocalizationCreateRequestDataAttributes) GetWhatsNewOk() (*string, bool) {
	if o == nil || IsNil(o.WhatsNew) {
		return nil, false
	}
	return o.WhatsNew, true
}

// HasWhatsNew returns a boolean if a field has been set.
func (o *BetaBuildLocalizationCreateRequestDataAttributes) HasWhatsNew() bool {
	if o != nil && !IsNil(o.WhatsNew) {
		return true
	}

	return false
}

// SetWhatsNew gets a reference to the given string and assigns it to the WhatsNew field.
func (o *BetaBuildLocalizationCreateRequestDataAttributes) SetWhatsNew(v string) {
	o.WhatsNew = &v
}

// GetLocale returns the Locale field value
func (o *BetaBuildLocalizationCreateRequestDataAttributes) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *BetaBuildLocalizationCreateRequestDataAttributes) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *BetaBuildLocalizationCreateRequestDataAttributes) SetLocale(v string) {
	o.Locale = v
}

func (o BetaBuildLocalizationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaBuildLocalizationCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WhatsNew) {
		toSerialize["whatsNew"] = o.WhatsNew
	}
	toSerialize["locale"] = o.Locale
	return toSerialize, nil
}

type NullableBetaBuildLocalizationCreateRequestDataAttributes struct {
	value *BetaBuildLocalizationCreateRequestDataAttributes
	isSet bool
}

func (v NullableBetaBuildLocalizationCreateRequestDataAttributes) Get() *BetaBuildLocalizationCreateRequestDataAttributes {
	return v.value
}

func (v *NullableBetaBuildLocalizationCreateRequestDataAttributes) Set(val *BetaBuildLocalizationCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaBuildLocalizationCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaBuildLocalizationCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaBuildLocalizationCreateRequestDataAttributes(val *BetaBuildLocalizationCreateRequestDataAttributes) *NullableBetaBuildLocalizationCreateRequestDataAttributes {
	return &NullableBetaBuildLocalizationCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableBetaBuildLocalizationCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaBuildLocalizationCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
