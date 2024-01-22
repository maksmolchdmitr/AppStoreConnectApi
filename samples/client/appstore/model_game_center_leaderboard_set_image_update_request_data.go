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

// checks if the GameCenterLeaderboardSetImageUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetImageUpdateRequestData{}

// GameCenterLeaderboardSetImageUpdateRequestData struct for GameCenterLeaderboardSetImageUpdateRequestData
type GameCenterLeaderboardSetImageUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppEventScreenshotUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

// NewGameCenterLeaderboardSetImageUpdateRequestData instantiates a new GameCenterLeaderboardSetImageUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetImageUpdateRequestData(type_ string, id string) *GameCenterLeaderboardSetImageUpdateRequestData {
	this := GameCenterLeaderboardSetImageUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterLeaderboardSetImageUpdateRequestDataWithDefaults instantiates a new GameCenterLeaderboardSetImageUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetImageUpdateRequestDataWithDefaults() *GameCenterLeaderboardSetImageUpdateRequestData {
	this := GameCenterLeaderboardSetImageUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterLeaderboardSetImageUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetImageUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterLeaderboardSetImageUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterLeaderboardSetImageUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetImageUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterLeaderboardSetImageUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetImageUpdateRequestData) GetAttributes() AppEventScreenshotUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppEventScreenshotUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetImageUpdateRequestData) GetAttributesOk() (*AppEventScreenshotUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetImageUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppEventScreenshotUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *GameCenterLeaderboardSetImageUpdateRequestData) SetAttributes(v AppEventScreenshotUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o GameCenterLeaderboardSetImageUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetImageUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableGameCenterLeaderboardSetImageUpdateRequestData struct {
	value *GameCenterLeaderboardSetImageUpdateRequestData
	isSet bool
}

func (v NullableGameCenterLeaderboardSetImageUpdateRequestData) Get() *GameCenterLeaderboardSetImageUpdateRequestData {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetImageUpdateRequestData) Set(val *GameCenterLeaderboardSetImageUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetImageUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetImageUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetImageUpdateRequestData(val *GameCenterLeaderboardSetImageUpdateRequestData) *NullableGameCenterLeaderboardSetImageUpdateRequestData {
	return &NullableGameCenterLeaderboardSetImageUpdateRequestData{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetImageUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetImageUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


