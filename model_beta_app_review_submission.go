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

// checks if the BetaAppReviewSubmission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppReviewSubmission{}

// BetaAppReviewSubmission struct for BetaAppReviewSubmission
type BetaAppReviewSubmission struct {
	Type          string                                `json:"type"`
	Id            string                                `json:"id"`
	Attributes    *BetaAppReviewSubmissionAttributes    `json:"attributes,omitempty"`
	Relationships *BetaAppReviewSubmissionRelationships `json:"relationships,omitempty"`
	Links         *ResourceLinks                        `json:"links,omitempty"`
}

// NewBetaAppReviewSubmission instantiates a new BetaAppReviewSubmission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppReviewSubmission(type_ string, id string) *BetaAppReviewSubmission {
	this := BetaAppReviewSubmission{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewBetaAppReviewSubmissionWithDefaults instantiates a new BetaAppReviewSubmission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppReviewSubmissionWithDefaults() *BetaAppReviewSubmission {
	this := BetaAppReviewSubmission{}
	return &this
}

// GetType returns the Type field value
func (o *BetaAppReviewSubmission) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BetaAppReviewSubmission) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BetaAppReviewSubmission) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BetaAppReviewSubmission) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BetaAppReviewSubmission) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BetaAppReviewSubmission) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *BetaAppReviewSubmission) GetAttributes() BetaAppReviewSubmissionAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret BetaAppReviewSubmissionAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppReviewSubmission) GetAttributesOk() (*BetaAppReviewSubmissionAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BetaAppReviewSubmission) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given BetaAppReviewSubmissionAttributes and assigns it to the Attributes field.
func (o *BetaAppReviewSubmission) SetAttributes(v BetaAppReviewSubmissionAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BetaAppReviewSubmission) GetRelationships() BetaAppReviewSubmissionRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret BetaAppReviewSubmissionRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppReviewSubmission) GetRelationshipsOk() (*BetaAppReviewSubmissionRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BetaAppReviewSubmission) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given BetaAppReviewSubmissionRelationships and assigns it to the Relationships field.
func (o *BetaAppReviewSubmission) SetRelationships(v BetaAppReviewSubmissionRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BetaAppReviewSubmission) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppReviewSubmission) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BetaAppReviewSubmission) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *BetaAppReviewSubmission) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o BetaAppReviewSubmission) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppReviewSubmission) ToMap() (map[string]interface{}, error) {
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

type NullableBetaAppReviewSubmission struct {
	value *BetaAppReviewSubmission
	isSet bool
}

func (v NullableBetaAppReviewSubmission) Get() *BetaAppReviewSubmission {
	return v.value
}

func (v *NullableBetaAppReviewSubmission) Set(val *BetaAppReviewSubmission) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppReviewSubmission) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppReviewSubmission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppReviewSubmission(val *BetaAppReviewSubmission) *NullableBetaAppReviewSubmission {
	return &NullableBetaAppReviewSubmission{value: val, isSet: true}
}

func (v NullableBetaAppReviewSubmission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppReviewSubmission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
