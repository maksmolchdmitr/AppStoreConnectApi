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

// checks if the ScmGitReferenceRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScmGitReferenceRelationships{}

// ScmGitReferenceRelationships struct for ScmGitReferenceRelationships
type ScmGitReferenceRelationships struct {
	Repository *CiWorkflowRelationshipsRepository `json:"repository,omitempty"`
}

// NewScmGitReferenceRelationships instantiates a new ScmGitReferenceRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScmGitReferenceRelationships() *ScmGitReferenceRelationships {
	this := ScmGitReferenceRelationships{}
	return &this
}

// NewScmGitReferenceRelationshipsWithDefaults instantiates a new ScmGitReferenceRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScmGitReferenceRelationshipsWithDefaults() *ScmGitReferenceRelationships {
	this := ScmGitReferenceRelationships{}
	return &this
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *ScmGitReferenceRelationships) GetRepository() CiWorkflowRelationshipsRepository {
	if o == nil || IsNil(o.Repository) {
		var ret CiWorkflowRelationshipsRepository
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmGitReferenceRelationships) GetRepositoryOk() (*CiWorkflowRelationshipsRepository, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *ScmGitReferenceRelationships) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given CiWorkflowRelationshipsRepository and assigns it to the Repository field.
func (o *ScmGitReferenceRelationships) SetRepository(v CiWorkflowRelationshipsRepository) {
	o.Repository = &v
}

func (o ScmGitReferenceRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScmGitReferenceRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	return toSerialize, nil
}

type NullableScmGitReferenceRelationships struct {
	value *ScmGitReferenceRelationships
	isSet bool
}

func (v NullableScmGitReferenceRelationships) Get() *ScmGitReferenceRelationships {
	return v.value
}

func (v *NullableScmGitReferenceRelationships) Set(val *ScmGitReferenceRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableScmGitReferenceRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableScmGitReferenceRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScmGitReferenceRelationships(val *ScmGitReferenceRelationships) *NullableScmGitReferenceRelationships {
	return &NullableScmGitReferenceRelationships{value: val, isSet: true}
}

func (v NullableScmGitReferenceRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScmGitReferenceRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


