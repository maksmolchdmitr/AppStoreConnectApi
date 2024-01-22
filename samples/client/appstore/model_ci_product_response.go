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

// checks if the CiProductResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiProductResponse{}

// CiProductResponse struct for CiProductResponse
type CiProductResponse struct {
	Data CiProduct `json:"data"`
	Included []CiProductsResponseIncludedInner `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// NewCiProductResponse instantiates a new CiProductResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiProductResponse(data CiProduct, links DocumentLinks) *CiProductResponse {
	this := CiProductResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewCiProductResponseWithDefaults instantiates a new CiProductResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiProductResponseWithDefaults() *CiProductResponse {
	this := CiProductResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CiProductResponse) GetData() CiProduct {
	if o == nil {
		var ret CiProduct
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CiProductResponse) GetDataOk() (*CiProduct, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CiProductResponse) SetData(v CiProduct) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *CiProductResponse) GetIncluded() []CiProductsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []CiProductsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiProductResponse) GetIncludedOk() ([]CiProductsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *CiProductResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []CiProductsResponseIncludedInner and assigns it to the Included field.
func (o *CiProductResponse) SetIncluded(v []CiProductsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *CiProductResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CiProductResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CiProductResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o CiProductResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiProductResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableCiProductResponse struct {
	value *CiProductResponse
	isSet bool
}

func (v NullableCiProductResponse) Get() *CiProductResponse {
	return v.value
}

func (v *NullableCiProductResponse) Set(val *CiProductResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCiProductResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCiProductResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiProductResponse(val *CiProductResponse) *NullableCiProductResponse {
	return &NullableCiProductResponse{value: val, isSet: true}
}

func (v NullableCiProductResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiProductResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


