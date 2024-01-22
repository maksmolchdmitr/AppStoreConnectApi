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

// checks if the AppPriceTierResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPriceTierResponse{}

// AppPriceTierResponse struct for AppPriceTierResponse
type AppPriceTierResponse struct {
	// Deprecated
	Data     AppPriceTier    `json:"data"`
	Included []AppPricePoint `json:"included,omitempty"`
	Links    DocumentLinks   `json:"links"`
}

// NewAppPriceTierResponse instantiates a new AppPriceTierResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPriceTierResponse(data AppPriceTier, links DocumentLinks) *AppPriceTierResponse {
	this := AppPriceTierResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppPriceTierResponseWithDefaults instantiates a new AppPriceTierResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPriceTierResponseWithDefaults() *AppPriceTierResponse {
	this := AppPriceTierResponse{}
	return &this
}

// GetData returns the Data field value
// Deprecated
func (o *AppPriceTierResponse) GetData() AppPriceTier {
	if o == nil {
		var ret AppPriceTier
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *AppPriceTierResponse) GetDataOk() (*AppPriceTier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
// Deprecated
func (o *AppPriceTierResponse) SetData(v AppPriceTier) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppPriceTierResponse) GetIncluded() []AppPricePoint {
	if o == nil || IsNil(o.Included) {
		var ret []AppPricePoint
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPriceTierResponse) GetIncludedOk() ([]AppPricePoint, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppPriceTierResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppPricePoint and assigns it to the Included field.
func (o *AppPriceTierResponse) SetIncluded(v []AppPricePoint) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppPriceTierResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppPriceTierResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppPriceTierResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppPriceTierResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPriceTierResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableAppPriceTierResponse struct {
	value *AppPriceTierResponse
	isSet bool
}

func (v NullableAppPriceTierResponse) Get() *AppPriceTierResponse {
	return v.value
}

func (v *NullableAppPriceTierResponse) Set(val *AppPriceTierResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPriceTierResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPriceTierResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPriceTierResponse(val *AppPriceTierResponse) *NullableAppPriceTierResponse {
	return &NullableAppPriceTierResponse{value: val, isSet: true}
}

func (v NullableAppPriceTierResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPriceTierResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}