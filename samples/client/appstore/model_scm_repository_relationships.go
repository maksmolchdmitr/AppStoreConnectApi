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

// checks if the ScmRepositoryRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScmRepositoryRelationships{}

// ScmRepositoryRelationships struct for ScmRepositoryRelationships
type ScmRepositoryRelationships struct {
	ScmProvider *ScmRepositoryRelationshipsScmProvider `json:"scmProvider,omitempty"`
	DefaultBranch *CiBuildRunRelationshipsSourceBranchOrTag `json:"defaultBranch,omitempty"`
}

// NewScmRepositoryRelationships instantiates a new ScmRepositoryRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScmRepositoryRelationships() *ScmRepositoryRelationships {
	this := ScmRepositoryRelationships{}
	return &this
}

// NewScmRepositoryRelationshipsWithDefaults instantiates a new ScmRepositoryRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScmRepositoryRelationshipsWithDefaults() *ScmRepositoryRelationships {
	this := ScmRepositoryRelationships{}
	return &this
}

// GetScmProvider returns the ScmProvider field value if set, zero value otherwise.
func (o *ScmRepositoryRelationships) GetScmProvider() ScmRepositoryRelationshipsScmProvider {
	if o == nil || IsNil(o.ScmProvider) {
		var ret ScmRepositoryRelationshipsScmProvider
		return ret
	}
	return *o.ScmProvider
}

// GetScmProviderOk returns a tuple with the ScmProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmRepositoryRelationships) GetScmProviderOk() (*ScmRepositoryRelationshipsScmProvider, bool) {
	if o == nil || IsNil(o.ScmProvider) {
		return nil, false
	}
	return o.ScmProvider, true
}

// HasScmProvider returns a boolean if a field has been set.
func (o *ScmRepositoryRelationships) HasScmProvider() bool {
	if o != nil && !IsNil(o.ScmProvider) {
		return true
	}

	return false
}

// SetScmProvider gets a reference to the given ScmRepositoryRelationshipsScmProvider and assigns it to the ScmProvider field.
func (o *ScmRepositoryRelationships) SetScmProvider(v ScmRepositoryRelationshipsScmProvider) {
	o.ScmProvider = &v
}

// GetDefaultBranch returns the DefaultBranch field value if set, zero value otherwise.
func (o *ScmRepositoryRelationships) GetDefaultBranch() CiBuildRunRelationshipsSourceBranchOrTag {
	if o == nil || IsNil(o.DefaultBranch) {
		var ret CiBuildRunRelationshipsSourceBranchOrTag
		return ret
	}
	return *o.DefaultBranch
}

// GetDefaultBranchOk returns a tuple with the DefaultBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmRepositoryRelationships) GetDefaultBranchOk() (*CiBuildRunRelationshipsSourceBranchOrTag, bool) {
	if o == nil || IsNil(o.DefaultBranch) {
		return nil, false
	}
	return o.DefaultBranch, true
}

// HasDefaultBranch returns a boolean if a field has been set.
func (o *ScmRepositoryRelationships) HasDefaultBranch() bool {
	if o != nil && !IsNil(o.DefaultBranch) {
		return true
	}

	return false
}

// SetDefaultBranch gets a reference to the given CiBuildRunRelationshipsSourceBranchOrTag and assigns it to the DefaultBranch field.
func (o *ScmRepositoryRelationships) SetDefaultBranch(v CiBuildRunRelationshipsSourceBranchOrTag) {
	o.DefaultBranch = &v
}

func (o ScmRepositoryRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScmRepositoryRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScmProvider) {
		toSerialize["scmProvider"] = o.ScmProvider
	}
	if !IsNil(o.DefaultBranch) {
		toSerialize["defaultBranch"] = o.DefaultBranch
	}
	return toSerialize, nil
}

type NullableScmRepositoryRelationships struct {
	value *ScmRepositoryRelationships
	isSet bool
}

func (v NullableScmRepositoryRelationships) Get() *ScmRepositoryRelationships {
	return v.value
}

func (v *NullableScmRepositoryRelationships) Set(val *ScmRepositoryRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableScmRepositoryRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableScmRepositoryRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScmRepositoryRelationships(val *ScmRepositoryRelationships) *NullableScmRepositoryRelationships {
	return &NullableScmRepositoryRelationships{value: val, isSet: true}
}

func (v NullableScmRepositoryRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScmRepositoryRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


