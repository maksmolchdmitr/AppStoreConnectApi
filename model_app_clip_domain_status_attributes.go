/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the AppClipDomainStatusAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipDomainStatusAttributes{}

// AppClipDomainStatusAttributes struct for AppClipDomainStatusAttributes
type AppClipDomainStatusAttributes struct {
	Domains         []AppClipDomainStatusAttributesDomainsInner `json:"domains,omitempty"`
	LastUpdatedDate *time.Time                                  `json:"lastUpdatedDate,omitempty"`
}

// NewAppClipDomainStatusAttributes instantiates a new AppClipDomainStatusAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipDomainStatusAttributes() *AppClipDomainStatusAttributes {
	this := AppClipDomainStatusAttributes{}
	return &this
}

// NewAppClipDomainStatusAttributesWithDefaults instantiates a new AppClipDomainStatusAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipDomainStatusAttributesWithDefaults() *AppClipDomainStatusAttributes {
	this := AppClipDomainStatusAttributes{}
	return &this
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *AppClipDomainStatusAttributes) GetDomains() []AppClipDomainStatusAttributesDomainsInner {
	if o == nil || IsNil(o.Domains) {
		var ret []AppClipDomainStatusAttributesDomainsInner
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDomainStatusAttributes) GetDomainsOk() ([]AppClipDomainStatusAttributesDomainsInner, bool) {
	if o == nil || IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *AppClipDomainStatusAttributes) HasDomains() bool {
	if o != nil && !IsNil(o.Domains) {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []AppClipDomainStatusAttributesDomainsInner and assigns it to the Domains field.
func (o *AppClipDomainStatusAttributes) SetDomains(v []AppClipDomainStatusAttributesDomainsInner) {
	o.Domains = v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value if set, zero value otherwise.
func (o *AppClipDomainStatusAttributes) GetLastUpdatedDate() time.Time {
	if o == nil || IsNil(o.LastUpdatedDate) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipDomainStatusAttributes) GetLastUpdatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedDate) {
		return nil, false
	}
	return o.LastUpdatedDate, true
}

// HasLastUpdatedDate returns a boolean if a field has been set.
func (o *AppClipDomainStatusAttributes) HasLastUpdatedDate() bool {
	if o != nil && !IsNil(o.LastUpdatedDate) {
		return true
	}

	return false
}

// SetLastUpdatedDate gets a reference to the given time.Time and assigns it to the LastUpdatedDate field.
func (o *AppClipDomainStatusAttributes) SetLastUpdatedDate(v time.Time) {
	o.LastUpdatedDate = &v
}

func (o AppClipDomainStatusAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipDomainStatusAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domains) {
		toSerialize["domains"] = o.Domains
	}
	if !IsNil(o.LastUpdatedDate) {
		toSerialize["lastUpdatedDate"] = o.LastUpdatedDate
	}
	return toSerialize, nil
}

type NullableAppClipDomainStatusAttributes struct {
	value *AppClipDomainStatusAttributes
	isSet bool
}

func (v NullableAppClipDomainStatusAttributes) Get() *AppClipDomainStatusAttributes {
	return v.value
}

func (v *NullableAppClipDomainStatusAttributes) Set(val *AppClipDomainStatusAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDomainStatusAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDomainStatusAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDomainStatusAttributes(val *AppClipDomainStatusAttributes) *NullableAppClipDomainStatusAttributes {
	return &NullableAppClipDomainStatusAttributes{value: val, isSet: true}
}

func (v NullableAppClipDomainStatusAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDomainStatusAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
