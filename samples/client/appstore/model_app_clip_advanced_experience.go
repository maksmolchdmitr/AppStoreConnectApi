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

// checks if the AppClipAdvancedExperience type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAdvancedExperience{}

// AppClipAdvancedExperience struct for AppClipAdvancedExperience
type AppClipAdvancedExperience struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppClipAdvancedExperienceAttributes `json:"attributes,omitempty"`
	Relationships *AppClipAdvancedExperienceRelationships `json:"relationships,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

// NewAppClipAdvancedExperience instantiates a new AppClipAdvancedExperience object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAdvancedExperience(type_ string, id string) *AppClipAdvancedExperience {
	this := AppClipAdvancedExperience{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppClipAdvancedExperienceWithDefaults instantiates a new AppClipAdvancedExperience object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAdvancedExperienceWithDefaults() *AppClipAdvancedExperience {
	this := AppClipAdvancedExperience{}
	return &this
}

// GetType returns the Type field value
func (o *AppClipAdvancedExperience) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperience) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppClipAdvancedExperience) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppClipAdvancedExperience) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperience) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppClipAdvancedExperience) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppClipAdvancedExperience) GetAttributes() AppClipAdvancedExperienceAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppClipAdvancedExperienceAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperience) GetAttributesOk() (*AppClipAdvancedExperienceAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppClipAdvancedExperience) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppClipAdvancedExperienceAttributes and assigns it to the Attributes field.
func (o *AppClipAdvancedExperience) SetAttributes(v AppClipAdvancedExperienceAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppClipAdvancedExperience) GetRelationships() AppClipAdvancedExperienceRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AppClipAdvancedExperienceRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperience) GetRelationshipsOk() (*AppClipAdvancedExperienceRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppClipAdvancedExperience) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppClipAdvancedExperienceRelationships and assigns it to the Relationships field.
func (o *AppClipAdvancedExperience) SetRelationships(v AppClipAdvancedExperienceRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppClipAdvancedExperience) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperience) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppClipAdvancedExperience) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *AppClipAdvancedExperience) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o AppClipAdvancedExperience) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAdvancedExperience) ToMap() (map[string]interface{}, error) {
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

type NullableAppClipAdvancedExperience struct {
	value *AppClipAdvancedExperience
	isSet bool
}

func (v NullableAppClipAdvancedExperience) Get() *AppClipAdvancedExperience {
	return v.value
}

func (v *NullableAppClipAdvancedExperience) Set(val *AppClipAdvancedExperience) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperience) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperience) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperience(val *AppClipAdvancedExperience) *NullableAppClipAdvancedExperience {
	return &NullableAppClipAdvancedExperience{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperience) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperience) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


