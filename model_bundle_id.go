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

// checks if the BundleId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleId{}

// BundleId struct for BundleId
type BundleId struct {
	Type          string                 `json:"type"`
	Id            string                 `json:"id"`
	Attributes    *BundleIdAttributes    `json:"attributes,omitempty"`
	Relationships *BundleIdRelationships `json:"relationships,omitempty"`
	Links         *ResourceLinks         `json:"links,omitempty"`
}

// NewBundleId instantiates a new BundleId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleId(type_ string, id string) *BundleId {
	this := BundleId{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewBundleIdWithDefaults instantiates a new BundleId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleIdWithDefaults() *BundleId {
	this := BundleId{}
	return &this
}

// GetType returns the Type field value
func (o *BundleId) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BundleId) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BundleId) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BundleId) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BundleId) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BundleId) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *BundleId) GetAttributes() BundleIdAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret BundleIdAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleId) GetAttributesOk() (*BundleIdAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BundleId) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given BundleIdAttributes and assigns it to the Attributes field.
func (o *BundleId) SetAttributes(v BundleIdAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BundleId) GetRelationships() BundleIdRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret BundleIdRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleId) GetRelationshipsOk() (*BundleIdRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BundleId) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given BundleIdRelationships and assigns it to the Relationships field.
func (o *BundleId) SetRelationships(v BundleIdRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BundleId) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleId) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BundleId) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *BundleId) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o BundleId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleId) ToMap() (map[string]interface{}, error) {
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

type NullableBundleId struct {
	value *BundleId
	isSet bool
}

func (v NullableBundleId) Get() *BundleId {
	return v.value
}

func (v *NullableBundleId) Set(val *BundleId) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleId) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleId(val *BundleId) *NullableBundleId {
	return &NullableBundleId{value: val, isSet: true}
}

func (v NullableBundleId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}