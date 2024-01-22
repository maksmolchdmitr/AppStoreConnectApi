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

// checks if the Actor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Actor{}

// Actor struct for Actor
type Actor struct {
	Type       string           `json:"type"`
	Id         string           `json:"id"`
	Attributes *ActorAttributes `json:"attributes,omitempty"`
	Links      *ResourceLinks   `json:"links,omitempty"`
}

// NewActor instantiates a new Actor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActor(type_ string, id string) *Actor {
	this := Actor{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewActorWithDefaults instantiates a new Actor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActorWithDefaults() *Actor {
	this := Actor{}
	return &this
}

// GetType returns the Type field value
func (o *Actor) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Actor) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Actor) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *Actor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Actor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Actor) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Actor) GetAttributes() ActorAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret ActorAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actor) GetAttributesOk() (*ActorAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Actor) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ActorAttributes and assigns it to the Attributes field.
func (o *Actor) SetAttributes(v ActorAttributes) {
	o.Attributes = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Actor) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actor) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Actor) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *Actor) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o Actor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Actor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableActor struct {
	value *Actor
	isSet bool
}

func (v NullableActor) Get() *Actor {
	return v.value
}

func (v *NullableActor) Set(val *Actor) {
	v.value = val
	v.isSet = true
}

func (v NullableActor) IsSet() bool {
	return v.isSet
}

func (v *NullableActor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActor(val *Actor) *NullableActor {
	return &NullableActor{value: val, isSet: true}
}

func (v NullableActor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
