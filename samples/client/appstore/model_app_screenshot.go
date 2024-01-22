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

// checks if the AppScreenshot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppScreenshot{}

// AppScreenshot struct for AppScreenshot
type AppScreenshot struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppScreenshotAttributes `json:"attributes,omitempty"`
	Relationships *AppScreenshotRelationships `json:"relationships,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

// NewAppScreenshot instantiates a new AppScreenshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppScreenshot(type_ string, id string) *AppScreenshot {
	this := AppScreenshot{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppScreenshotWithDefaults instantiates a new AppScreenshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppScreenshotWithDefaults() *AppScreenshot {
	this := AppScreenshot{}
	return &this
}

// GetType returns the Type field value
func (o *AppScreenshot) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppScreenshot) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppScreenshot) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppScreenshot) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppScreenshot) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppScreenshot) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppScreenshot) GetAttributes() AppScreenshotAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppScreenshotAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppScreenshot) GetAttributesOk() (*AppScreenshotAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppScreenshot) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppScreenshotAttributes and assigns it to the Attributes field.
func (o *AppScreenshot) SetAttributes(v AppScreenshotAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *AppScreenshot) GetRelationships() AppScreenshotRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret AppScreenshotRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppScreenshot) GetRelationshipsOk() (*AppScreenshotRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *AppScreenshot) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given AppScreenshotRelationships and assigns it to the Relationships field.
func (o *AppScreenshot) SetRelationships(v AppScreenshotRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AppScreenshot) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppScreenshot) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AppScreenshot) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *AppScreenshot) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o AppScreenshot) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppScreenshot) ToMap() (map[string]interface{}, error) {
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

type NullableAppScreenshot struct {
	value *AppScreenshot
	isSet bool
}

func (v NullableAppScreenshot) Get() *AppScreenshot {
	return v.value
}

func (v *NullableAppScreenshot) Set(val *AppScreenshot) {
	v.value = val
	v.isSet = true
}

func (v NullableAppScreenshot) IsSet() bool {
	return v.isSet
}

func (v *NullableAppScreenshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppScreenshot(val *AppScreenshot) *NullableAppScreenshot {
	return &NullableAppScreenshot{value: val, isSet: true}
}

func (v NullableAppScreenshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppScreenshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


