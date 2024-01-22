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

// checks if the ReviewSubmission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewSubmission{}

// ReviewSubmission struct for ReviewSubmission
type ReviewSubmission struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *ReviewSubmissionAttributes `json:"attributes,omitempty"`
	Relationships *ReviewSubmissionRelationships `json:"relationships,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

// NewReviewSubmission instantiates a new ReviewSubmission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewSubmission(type_ string, id string) *ReviewSubmission {
	this := ReviewSubmission{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewReviewSubmissionWithDefaults instantiates a new ReviewSubmission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewSubmissionWithDefaults() *ReviewSubmission {
	this := ReviewSubmission{}
	return &this
}

// GetType returns the Type field value
func (o *ReviewSubmission) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReviewSubmission) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReviewSubmission) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ReviewSubmission) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReviewSubmission) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReviewSubmission) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ReviewSubmission) GetAttributes() ReviewSubmissionAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret ReviewSubmissionAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmission) GetAttributesOk() (*ReviewSubmissionAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ReviewSubmission) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ReviewSubmissionAttributes and assigns it to the Attributes field.
func (o *ReviewSubmission) SetAttributes(v ReviewSubmissionAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *ReviewSubmission) GetRelationships() ReviewSubmissionRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret ReviewSubmissionRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmission) GetRelationshipsOk() (*ReviewSubmissionRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *ReviewSubmission) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given ReviewSubmissionRelationships and assigns it to the Relationships field.
func (o *ReviewSubmission) SetRelationships(v ReviewSubmissionRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ReviewSubmission) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmission) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ReviewSubmission) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *ReviewSubmission) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o ReviewSubmission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewSubmission) ToMap() (map[string]interface{}, error) {
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

type NullableReviewSubmission struct {
	value *ReviewSubmission
	isSet bool
}

func (v NullableReviewSubmission) Get() *ReviewSubmission {
	return v.value
}

func (v *NullableReviewSubmission) Set(val *ReviewSubmission) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewSubmission) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewSubmission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewSubmission(val *ReviewSubmission) *NullableReviewSubmission {
	return &NullableReviewSubmission{value: val, isSet: true}
}

func (v NullableReviewSubmission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewSubmission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


