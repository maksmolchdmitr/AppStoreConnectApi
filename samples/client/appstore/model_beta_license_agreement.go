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

// checks if the BetaLicenseAgreement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaLicenseAgreement{}

// BetaLicenseAgreement struct for BetaLicenseAgreement
type BetaLicenseAgreement struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *BetaLicenseAgreementAttributes `json:"attributes,omitempty"`
	Relationships *AppPreOrderRelationships `json:"relationships,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

// NewBetaLicenseAgreement instantiates a new BetaLicenseAgreement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaLicenseAgreement(type_ string, id string) *BetaLicenseAgreement {
	this := BetaLicenseAgreement{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewBetaLicenseAgreementWithDefaults instantiates a new BetaLicenseAgreement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaLicenseAgreementWithDefaults() *BetaLicenseAgreement {
	this := BetaLicenseAgreement{}
	return &this
}

// GetType returns the Type field value
func (o *BetaLicenseAgreement) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BetaLicenseAgreement) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BetaLicenseAgreement) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BetaLicenseAgreement) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BetaLicenseAgreement) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BetaLicenseAgreement) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *BetaLicenseAgreement) GetAttributes() BetaLicenseAgreementAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret BetaLicenseAgreementAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaLicenseAgreement) GetAttributesOk() (*BetaLicenseAgreementAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BetaLicenseAgreement) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given BetaLicenseAgreementAttributes and assigns it to the Attributes field.
func (o *BetaLicenseAgreement) SetAttributes(v BetaLicenseAgreementAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BetaLicenseAgreement) GetRelationships() AppPreOrderRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AppPreOrderRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaLicenseAgreement) GetRelationshipsOk() (*AppPreOrderRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BetaLicenseAgreement) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppPreOrderRelationships and assigns it to the Relationships field.
func (o *BetaLicenseAgreement) SetRelationships(v AppPreOrderRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BetaLicenseAgreement) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaLicenseAgreement) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BetaLicenseAgreement) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *BetaLicenseAgreement) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o BetaLicenseAgreement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaLicenseAgreement) ToMap() (map[string]interface{}, error) {
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

type NullableBetaLicenseAgreement struct {
	value *BetaLicenseAgreement
	isSet bool
}

func (v NullableBetaLicenseAgreement) Get() *BetaLicenseAgreement {
	return v.value
}

func (v *NullableBetaLicenseAgreement) Set(val *BetaLicenseAgreement) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaLicenseAgreement) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaLicenseAgreement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaLicenseAgreement(val *BetaLicenseAgreement) *NullableBetaLicenseAgreement {
	return &NullableBetaLicenseAgreement{value: val, isSet: true}
}

func (v NullableBetaLicenseAgreement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaLicenseAgreement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


