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

// checks if the AppPricePointsV3Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPricePointsV3Response{}

// AppPricePointsV3Response struct for AppPricePointsV3Response
type AppPricePointsV3Response struct {
	Data     []AppPricePointV3                      `json:"data"`
	Included []AppAvailabilityResponseIncludedInner `json:"included,omitempty"`
	Links    PagedDocumentLinks                     `json:"links"`
	Meta     *PagingInformation                     `json:"meta,omitempty"`
}

// NewAppPricePointsV3Response instantiates a new AppPricePointsV3Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPricePointsV3Response(data []AppPricePointV3, links PagedDocumentLinks) *AppPricePointsV3Response {
	this := AppPricePointsV3Response{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppPricePointsV3ResponseWithDefaults instantiates a new AppPricePointsV3Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPricePointsV3ResponseWithDefaults() *AppPricePointsV3Response {
	this := AppPricePointsV3Response{}
	return &this
}

// GetData returns the Data field value
func (o *AppPricePointsV3Response) GetData() []AppPricePointV3 {
	if o == nil {
		var ret []AppPricePointV3
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppPricePointsV3Response) GetDataOk() ([]AppPricePointV3, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AppPricePointsV3Response) SetData(v []AppPricePointV3) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppPricePointsV3Response) GetIncluded() []AppAvailabilityResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []AppAvailabilityResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricePointsV3Response) GetIncludedOk() ([]AppAvailabilityResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppPricePointsV3Response) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppAvailabilityResponseIncludedInner and assigns it to the Included field.
func (o *AppPricePointsV3Response) SetIncluded(v []AppAvailabilityResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppPricePointsV3Response) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppPricePointsV3Response) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppPricePointsV3Response) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppPricePointsV3Response) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPricePointsV3Response) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppPricePointsV3Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppPricePointsV3Response) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o AppPricePointsV3Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPricePointsV3Response) ToMap() (map[string]interface{}, error) {
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

type NullableAppPricePointsV3Response struct {
	value *AppPricePointsV3Response
	isSet bool
}

func (v NullableAppPricePointsV3Response) Get() *AppPricePointsV3Response {
	return v.value
}

func (v *NullableAppPricePointsV3Response) Set(val *AppPricePointsV3Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPricePointsV3Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPricePointsV3Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPricePointsV3Response(val *AppPricePointsV3Response) *NullableAppPricePointsV3Response {
	return &NullableAppPricePointsV3Response{value: val, isSet: true}
}

func (v NullableAppPricePointsV3Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPricePointsV3Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
