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

// checks if the CiBuildAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiBuildAction{}

// CiBuildAction struct for CiBuildAction
type CiBuildAction struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *CiBuildActionAttributes `json:"attributes,omitempty"`
	Relationships *CiBuildActionRelationships `json:"relationships,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

// NewCiBuildAction instantiates a new CiBuildAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiBuildAction(type_ string, id string) *CiBuildAction {
	this := CiBuildAction{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewCiBuildActionWithDefaults instantiates a new CiBuildAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiBuildActionWithDefaults() *CiBuildAction {
	this := CiBuildAction{}
	return &this
}

// GetType returns the Type field value
func (o *CiBuildAction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CiBuildAction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CiBuildAction) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CiBuildAction) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CiBuildAction) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CiBuildAction) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CiBuildAction) GetAttributes() CiBuildActionAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret CiBuildActionAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildAction) GetAttributesOk() (*CiBuildActionAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CiBuildAction) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given CiBuildActionAttributes and assigns it to the Attributes field.
func (o *CiBuildAction) SetAttributes(v CiBuildActionAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *CiBuildAction) GetRelationships() CiBuildActionRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret CiBuildActionRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildAction) GetRelationshipsOk() (*CiBuildActionRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *CiBuildAction) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given CiBuildActionRelationships and assigns it to the Relationships field.
func (o *CiBuildAction) SetRelationships(v CiBuildActionRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CiBuildAction) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildAction) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CiBuildAction) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *CiBuildAction) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o CiBuildAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiBuildAction) ToMap() (map[string]interface{}, error) {
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

type NullableCiBuildAction struct {
	value *CiBuildAction
	isSet bool
}

func (v NullableCiBuildAction) Get() *CiBuildAction {
	return v.value
}

func (v *NullableCiBuildAction) Set(val *CiBuildAction) {
	v.value = val
	v.isSet = true
}

func (v NullableCiBuildAction) IsSet() bool {
	return v.isSet
}

func (v *NullableCiBuildAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiBuildAction(val *CiBuildAction) *NullableCiBuildAction {
	return &NullableCiBuildAction{value: val, isSet: true}
}

func (v NullableCiBuildAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiBuildAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


