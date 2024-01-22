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

// checks if the AppAvailabilityV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAvailabilityV2{}

// AppAvailabilityV2 struct for AppAvailabilityV2
type AppAvailabilityV2 struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppAvailabilityV2Attributes `json:"attributes,omitempty"`
	Relationships *AppAvailabilityV2Relationships `json:"relationships,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

// NewAppAvailabilityV2 instantiates a new AppAvailabilityV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAvailabilityV2(type_ string, id string) *AppAvailabilityV2 {
	this := AppAvailabilityV2{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppAvailabilityV2WithDefaults instantiates a new AppAvailabilityV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAvailabilityV2WithDefaults() *AppAvailabilityV2 {
	this := AppAvailabilityV2{}
	return &this
}

// GetType returns the Type field value
func (o *AppAvailabilityV2) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppAvailabilityV2) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppAvailabilityV2) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppAvailabilityV2) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppAvailabilityV2) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppAvailabilityV2) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppAvailabilityV2) GetAttributes() AppAvailabilityV2Attributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppAvailabilityV2Attributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAvailabilityV2) GetAttributesOk() (*AppAvailabilityV2Attributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppAvailabilityV2) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppAvailabilityV2Attributes and assigns it to the Attributes field.
func (o *AppAvailabilityV2) SetAttributes(v AppAvailabilityV2Attributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppAvailabilityV2) GetRelationships() AppAvailabilityV2Relationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AppAvailabilityV2Relationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAvailabilityV2) GetRelationshipsOk() (*AppAvailabilityV2Relationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppAvailabilityV2) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppAvailabilityV2Relationships and assigns it to the Relationships field.
func (o *AppAvailabilityV2) SetRelationships(v AppAvailabilityV2Relationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppAvailabilityV2) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAvailabilityV2) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppAvailabilityV2) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *AppAvailabilityV2) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o AppAvailabilityV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAvailabilityV2) ToMap() (map[string]interface{}, error) {
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

type NullableAppAvailabilityV2 struct {
	value *AppAvailabilityV2
	isSet bool
}

func (v NullableAppAvailabilityV2) Get() *AppAvailabilityV2 {
	return v.value
}

func (v *NullableAppAvailabilityV2) Set(val *AppAvailabilityV2) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAvailabilityV2) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAvailabilityV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAvailabilityV2(val *AppAvailabilityV2) *NullableAppAvailabilityV2 {
	return &NullableAppAvailabilityV2{value: val, isSet: true}
}

func (v NullableAppAvailabilityV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAvailabilityV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


