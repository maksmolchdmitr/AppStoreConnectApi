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

// checks if the AppPricePointV3Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPricePointV3Response{}

// AppPricePointV3Response struct for AppPricePointV3Response
type AppPricePointV3Response struct {
	Data     AppPricePointV3                        `json:"data"`
	Included []AppAvailabilityResponseIncludedInner `json:"included,omitempty"`
	Links    DocumentLinks                          `json:"links"`
}

// NewAppPricePointV3Response instantiates a new AppPricePointV3Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPricePointV3Response(data AppPricePointV3, links DocumentLinks) *AppPricePointV3Response {
	this := AppPricePointV3Response{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppPricePointV3ResponseWithDefaults instantiates a new AppPricePointV3Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPricePointV3ResponseWithDefaults() *AppPricePointV3Response {
	this := AppPricePointV3Response{}
	return &this
}

// GetData returns the Data field value
func (o *AppPricePointV3Response) GetData() AppPricePointV3 {
	if o == nil {
		var ret AppPricePointV3
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppPricePointV3Response) GetDataOk() (*AppPricePointV3, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppPricePointV3Response) SetData(v AppPricePointV3) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppPricePointV3Response) GetIncluded() []AppAvailabilityResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []AppAvailabilityResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricePointV3Response) GetIncludedOk() ([]AppAvailabilityResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppPricePointV3Response) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppAvailabilityResponseIncludedInner and assigns it to the Included field.
func (o *AppPricePointV3Response) SetIncluded(v []AppAvailabilityResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppPricePointV3Response) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppPricePointV3Response) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppPricePointV3Response) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppPricePointV3Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPricePointV3Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableAppPricePointV3Response struct {
	value *AppPricePointV3Response
	isSet bool
}

func (v NullableAppPricePointV3Response) Get() *AppPricePointV3Response {
	return v.value
}

func (v *NullableAppPricePointV3Response) Set(val *AppPricePointV3Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPricePointV3Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPricePointV3Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPricePointV3Response(val *AppPricePointV3Response) *NullableAppPricePointV3Response {
	return &NullableAppPricePointV3Response{value: val, isSet: true}
}

func (v NullableAppPricePointV3Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPricePointV3Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}