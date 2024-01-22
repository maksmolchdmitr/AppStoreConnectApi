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

// checks if the AppCustomProductPageVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageVersion{}

// AppCustomProductPageVersion struct for AppCustomProductPageVersion
type AppCustomProductPageVersion struct {
	Type          string                                    `json:"type"`
	Id            string                                    `json:"id"`
	Attributes    *AppCustomProductPageVersionAttributes    `json:"attributes,omitempty"`
	Relationships *AppCustomProductPageVersionRelationships `json:"relationships,omitempty"`
	Links         *ResourceLinks                            `json:"links,omitempty"`
}

// NewAppCustomProductPageVersion instantiates a new AppCustomProductPageVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageVersion(type_ string, id string) *AppCustomProductPageVersion {
	this := AppCustomProductPageVersion{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppCustomProductPageVersionWithDefaults instantiates a new AppCustomProductPageVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageVersionWithDefaults() *AppCustomProductPageVersion {
	this := AppCustomProductPageVersion{}
	return &this
}

// GetType returns the Type field value
func (o *AppCustomProductPageVersion) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageVersion) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppCustomProductPageVersion) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppCustomProductPageVersion) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageVersion) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppCustomProductPageVersion) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppCustomProductPageVersion) GetAttributes() AppCustomProductPageVersionAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppCustomProductPageVersionAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageVersion) GetAttributesOk() (*AppCustomProductPageVersionAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppCustomProductPageVersion) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppCustomProductPageVersionAttributes and assigns it to the Attributes field.
func (o *AppCustomProductPageVersion) SetAttributes(v AppCustomProductPageVersionAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppCustomProductPageVersion) GetRelationships() AppCustomProductPageVersionRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AppCustomProductPageVersionRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageVersion) GetRelationshipsOk() (*AppCustomProductPageVersionRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppCustomProductPageVersion) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppCustomProductPageVersionRelationships and assigns it to the Relationships field.
func (o *AppCustomProductPageVersion) SetRelationships(v AppCustomProductPageVersionRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppCustomProductPageVersion) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageVersion) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppCustomProductPageVersion) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *AppCustomProductPageVersion) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o AppCustomProductPageVersion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageVersion) ToMap() (map[string]interface{}, error) {
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

type NullableAppCustomProductPageVersion struct {
	value *AppCustomProductPageVersion
	isSet bool
}

func (v NullableAppCustomProductPageVersion) Get() *AppCustomProductPageVersion {
	return v.value
}

func (v *NullableAppCustomProductPageVersion) Set(val *AppCustomProductPageVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageVersion(val *AppCustomProductPageVersion) *NullableAppCustomProductPageVersion {
	return &NullableAppCustomProductPageVersion{value: val, isSet: true}
}

func (v NullableAppCustomProductPageVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}