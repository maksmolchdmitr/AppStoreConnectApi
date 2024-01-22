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

// checks if the AppStoreVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersion{}

// AppStoreVersion struct for AppStoreVersion
type AppStoreVersion struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppStoreVersionAttributes `json:"attributes,omitempty"`
	Relationships *AppStoreVersionRelationships `json:"relationships,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

// NewAppStoreVersion instantiates a new AppStoreVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersion(type_ string, id string) *AppStoreVersion {
	this := AppStoreVersion{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppStoreVersionWithDefaults instantiates a new AppStoreVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionWithDefaults() *AppStoreVersion {
	this := AppStoreVersion{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreVersion) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersion) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreVersion) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppStoreVersion) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersion) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppStoreVersion) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppStoreVersion) GetAttributes() AppStoreVersionAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppStoreVersionAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersion) GetAttributesOk() (*AppStoreVersionAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppStoreVersion) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppStoreVersionAttributes and assigns it to the Attributes field.
func (o *AppStoreVersion) SetAttributes(v AppStoreVersionAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppStoreVersion) GetRelationships() AppStoreVersionRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AppStoreVersionRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersion) GetRelationshipsOk() (*AppStoreVersionRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppStoreVersion) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppStoreVersionRelationships and assigns it to the Relationships field.
func (o *AppStoreVersion) SetRelationships(v AppStoreVersionRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppStoreVersion) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersion) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppStoreVersion) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *AppStoreVersion) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o AppStoreVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersion) ToMap() (map[string]interface{}, error) {
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

type NullableAppStoreVersion struct {
	value *AppStoreVersion
	isSet bool
}

func (v NullableAppStoreVersion) Get() *AppStoreVersion {
	return v.value
}

func (v *NullableAppStoreVersion) Set(val *AppStoreVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersion(val *AppStoreVersion) *NullableAppStoreVersion {
	return &NullableAppStoreVersion{value: val, isSet: true}
}

func (v NullableAppStoreVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


