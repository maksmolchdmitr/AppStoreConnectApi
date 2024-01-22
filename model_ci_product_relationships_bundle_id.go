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

// checks if the CiProductRelationshipsBundleId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiProductRelationshipsBundleId{}

// CiProductRelationshipsBundleId struct for CiProductRelationshipsBundleId
type CiProductRelationshipsBundleId struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks   `json:"links,omitempty"`
	Data  *BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData `json:"data,omitempty"`
}

// NewCiProductRelationshipsBundleId instantiates a new CiProductRelationshipsBundleId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiProductRelationshipsBundleId() *CiProductRelationshipsBundleId {
	this := CiProductRelationshipsBundleId{}
	return &this
}

// NewCiProductRelationshipsBundleIdWithDefaults instantiates a new CiProductRelationshipsBundleId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiProductRelationshipsBundleIdWithDefaults() *CiProductRelationshipsBundleId {
	this := CiProductRelationshipsBundleId{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CiProductRelationshipsBundleId) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiProductRelationshipsBundleId) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CiProductRelationshipsBundleId) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *CiProductRelationshipsBundleId) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CiProductRelationshipsBundleId) GetData() BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData {
	if o == nil || IsNil(o.Data) {
		var ret BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiProductRelationshipsBundleId) GetDataOk() (*BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CiProductRelationshipsBundleId) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData and assigns it to the Data field.
func (o *CiProductRelationshipsBundleId) SetData(v BundleIdCapabilityCreateRequestDataRelationshipsBundleIdData) {
	o.Data = &v
}

func (o CiProductRelationshipsBundleId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiProductRelationshipsBundleId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCiProductRelationshipsBundleId struct {
	value *CiProductRelationshipsBundleId
	isSet bool
}

func (v NullableCiProductRelationshipsBundleId) Get() *CiProductRelationshipsBundleId {
	return v.value
}

func (v *NullableCiProductRelationshipsBundleId) Set(val *CiProductRelationshipsBundleId) {
	v.value = val
	v.isSet = true
}

func (v NullableCiProductRelationshipsBundleId) IsSet() bool {
	return v.isSet
}

func (v *NullableCiProductRelationshipsBundleId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiProductRelationshipsBundleId(val *CiProductRelationshipsBundleId) *NullableCiProductRelationshipsBundleId {
	return &NullableCiProductRelationshipsBundleId{value: val, isSet: true}
}

func (v NullableCiProductRelationshipsBundleId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiProductRelationshipsBundleId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
