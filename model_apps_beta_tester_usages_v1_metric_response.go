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

// checks if the AppsBetaTesterUsagesV1MetricResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppsBetaTesterUsagesV1MetricResponse{}

// AppsBetaTesterUsagesV1MetricResponse struct for AppsBetaTesterUsagesV1MetricResponse
type AppsBetaTesterUsagesV1MetricResponse struct {
	Data     []AppsBetaTesterUsagesV1MetricResponseDataInner `json:"data"`
	Links    PagedDocumentLinks                              `json:"links"`
	Meta     *PagingInformation                              `json:"meta,omitempty"`
	Included []BetaTester                                    `json:"included,omitempty"`
}

// NewAppsBetaTesterUsagesV1MetricResponse instantiates a new AppsBetaTesterUsagesV1MetricResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppsBetaTesterUsagesV1MetricResponse(data []AppsBetaTesterUsagesV1MetricResponseDataInner, links PagedDocumentLinks) *AppsBetaTesterUsagesV1MetricResponse {
	this := AppsBetaTesterUsagesV1MetricResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppsBetaTesterUsagesV1MetricResponseWithDefaults instantiates a new AppsBetaTesterUsagesV1MetricResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppsBetaTesterUsagesV1MetricResponseWithDefaults() *AppsBetaTesterUsagesV1MetricResponse {
	this := AppsBetaTesterUsagesV1MetricResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppsBetaTesterUsagesV1MetricResponse) GetData() []AppsBetaTesterUsagesV1MetricResponseDataInner {
	if o == nil {
		var ret []AppsBetaTesterUsagesV1MetricResponseDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppsBetaTesterUsagesV1MetricResponse) GetDataOk() ([]AppsBetaTesterUsagesV1MetricResponseDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AppsBetaTesterUsagesV1MetricResponse) SetData(v []AppsBetaTesterUsagesV1MetricResponseDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *AppsBetaTesterUsagesV1MetricResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppsBetaTesterUsagesV1MetricResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppsBetaTesterUsagesV1MetricResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppsBetaTesterUsagesV1MetricResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppsBetaTesterUsagesV1MetricResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppsBetaTesterUsagesV1MetricResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppsBetaTesterUsagesV1MetricResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppsBetaTesterUsagesV1MetricResponse) GetIncluded() []BetaTester {
	if o == nil || IsNil(o.Included) {
		var ret []BetaTester
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppsBetaTesterUsagesV1MetricResponse) GetIncludedOk() ([]BetaTester, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppsBetaTesterUsagesV1MetricResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []BetaTester and assigns it to the Included field.
func (o *AppsBetaTesterUsagesV1MetricResponse) SetIncluded(v []BetaTester) {
	o.Included = v
}

func (o AppsBetaTesterUsagesV1MetricResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppsBetaTesterUsagesV1MetricResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	return toSerialize, nil
}

type NullableAppsBetaTesterUsagesV1MetricResponse struct {
	value *AppsBetaTesterUsagesV1MetricResponse
	isSet bool
}

func (v NullableAppsBetaTesterUsagesV1MetricResponse) Get() *AppsBetaTesterUsagesV1MetricResponse {
	return v.value
}

func (v *NullableAppsBetaTesterUsagesV1MetricResponse) Set(val *AppsBetaTesterUsagesV1MetricResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppsBetaTesterUsagesV1MetricResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppsBetaTesterUsagesV1MetricResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppsBetaTesterUsagesV1MetricResponse(val *AppsBetaTesterUsagesV1MetricResponse) *NullableAppsBetaTesterUsagesV1MetricResponse {
	return &NullableAppsBetaTesterUsagesV1MetricResponse{value: val, isSet: true}
}

func (v NullableAppsBetaTesterUsagesV1MetricResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppsBetaTesterUsagesV1MetricResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
