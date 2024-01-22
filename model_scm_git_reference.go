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

// checks if the ScmGitReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScmGitReference{}

// ScmGitReference struct for ScmGitReference
type ScmGitReference struct {
	Type          string                        `json:"type"`
	Id            string                        `json:"id"`
	Attributes    *ScmGitReferenceAttributes    `json:"attributes,omitempty"`
	Relationships *ScmGitReferenceRelationships `json:"relationships,omitempty"`
	Links         *ResourceLinks                `json:"links,omitempty"`
}

// NewScmGitReference instantiates a new ScmGitReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScmGitReference(type_ string, id string) *ScmGitReference {
	this := ScmGitReference{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewScmGitReferenceWithDefaults instantiates a new ScmGitReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScmGitReferenceWithDefaults() *ScmGitReference {
	this := ScmGitReference{}
	return &this
}

// GetType returns the Type field value
func (o *ScmGitReference) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScmGitReference) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ScmGitReference) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ScmGitReference) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScmGitReference) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ScmGitReference) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ScmGitReference) GetAttributes() ScmGitReferenceAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret ScmGitReferenceAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmGitReference) GetAttributesOk() (*ScmGitReferenceAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ScmGitReference) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ScmGitReferenceAttributes and assigns it to the Attributes field.
func (o *ScmGitReference) SetAttributes(v ScmGitReferenceAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ScmGitReference) GetRelationships() ScmGitReferenceRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret ScmGitReferenceRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmGitReference) GetRelationshipsOk() (*ScmGitReferenceRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ScmGitReference) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ScmGitReferenceRelationships and assigns it to the Relationships field.
func (o *ScmGitReference) SetRelationships(v ScmGitReferenceRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ScmGitReference) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmGitReference) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ScmGitReference) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *ScmGitReference) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o ScmGitReference) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScmGitReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableScmGitReference struct {
	value *ScmGitReference
	isSet bool
}

func (v NullableScmGitReference) Get() *ScmGitReference {
	return v.value
}

func (v *NullableScmGitReference) Set(val *ScmGitReference) {
	v.value = val
	v.isSet = true
}

func (v NullableScmGitReference) IsSet() bool {
	return v.isSet
}

func (v *NullableScmGitReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScmGitReference(val *ScmGitReference) *NullableScmGitReference {
	return &NullableScmGitReference{value: val, isSet: true}
}

func (v NullableScmGitReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScmGitReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
