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

// checks if the BuildBundleRelationshipsAppClipDomainCacheStatusData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildBundleRelationshipsAppClipDomainCacheStatusData{}

// BuildBundleRelationshipsAppClipDomainCacheStatusData struct for BuildBundleRelationshipsAppClipDomainCacheStatusData
type BuildBundleRelationshipsAppClipDomainCacheStatusData struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewBuildBundleRelationshipsAppClipDomainCacheStatusData instantiates a new BuildBundleRelationshipsAppClipDomainCacheStatusData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildBundleRelationshipsAppClipDomainCacheStatusData(type_ string, id string) *BuildBundleRelationshipsAppClipDomainCacheStatusData {
	this := BuildBundleRelationshipsAppClipDomainCacheStatusData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewBuildBundleRelationshipsAppClipDomainCacheStatusDataWithDefaults instantiates a new BuildBundleRelationshipsAppClipDomainCacheStatusData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildBundleRelationshipsAppClipDomainCacheStatusDataWithDefaults() *BuildBundleRelationshipsAppClipDomainCacheStatusData {
	this := BuildBundleRelationshipsAppClipDomainCacheStatusData{}
	return &this
}

// GetType returns the Type field value
func (o *BuildBundleRelationshipsAppClipDomainCacheStatusData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BuildBundleRelationshipsAppClipDomainCacheStatusData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BuildBundleRelationshipsAppClipDomainCacheStatusData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BuildBundleRelationshipsAppClipDomainCacheStatusData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BuildBundleRelationshipsAppClipDomainCacheStatusData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BuildBundleRelationshipsAppClipDomainCacheStatusData) SetId(v string) {
	o.Id = v
}

func (o BuildBundleRelationshipsAppClipDomainCacheStatusData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildBundleRelationshipsAppClipDomainCacheStatusData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableBuildBundleRelationshipsAppClipDomainCacheStatusData struct {
	value *BuildBundleRelationshipsAppClipDomainCacheStatusData
	isSet bool
}

func (v NullableBuildBundleRelationshipsAppClipDomainCacheStatusData) Get() *BuildBundleRelationshipsAppClipDomainCacheStatusData {
	return v.value
}

func (v *NullableBuildBundleRelationshipsAppClipDomainCacheStatusData) Set(val *BuildBundleRelationshipsAppClipDomainCacheStatusData) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildBundleRelationshipsAppClipDomainCacheStatusData) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildBundleRelationshipsAppClipDomainCacheStatusData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildBundleRelationshipsAppClipDomainCacheStatusData(val *BuildBundleRelationshipsAppClipDomainCacheStatusData) *NullableBuildBundleRelationshipsAppClipDomainCacheStatusData {
	return &NullableBuildBundleRelationshipsAppClipDomainCacheStatusData{value: val, isSet: true}
}

func (v NullableBuildBundleRelationshipsAppClipDomainCacheStatusData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildBundleRelationshipsAppClipDomainCacheStatusData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


