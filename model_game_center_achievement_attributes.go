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

// checks if the GameCenterAchievementAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementAttributes{}

// GameCenterAchievementAttributes struct for GameCenterAchievementAttributes
type GameCenterAchievementAttributes struct {
	ReferenceName    *string `json:"referenceName,omitempty"`
	VendorIdentifier *string `json:"vendorIdentifier,omitempty"`
	Points           *int32  `json:"points,omitempty"`
	ShowBeforeEarned *bool   `json:"showBeforeEarned,omitempty"`
	Repeatable       *bool   `json:"repeatable,omitempty"`
	Archived         *bool   `json:"archived,omitempty"`
}

// NewGameCenterAchievementAttributes instantiates a new GameCenterAchievementAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementAttributes() *GameCenterAchievementAttributes {
	this := GameCenterAchievementAttributes{}
	return &this
}

// NewGameCenterAchievementAttributesWithDefaults instantiates a new GameCenterAchievementAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementAttributesWithDefaults() *GameCenterAchievementAttributes {
	this := GameCenterAchievementAttributes{}
	return &this
}

// GetReferenceName returns the ReferenceName field value if set, zero value otherwise.
func (o *GameCenterAchievementAttributes) GetReferenceName() string {
	if o == nil || IsNil(o.ReferenceName) {
		var ret string
		return ret
	}
	return *o.ReferenceName
}

// GetReferenceNameOk returns a tuple with the ReferenceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementAttributes) GetReferenceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceName) {
		return nil, false
	}
	return o.ReferenceName, true
}

// HasReferenceName returns a boolean if a field has been set.
func (o *GameCenterAchievementAttributes) HasReferenceName() bool {
	if o != nil && !IsNil(o.ReferenceName) {
		return true
	}

	return false
}

// SetReferenceName gets a reference to the given string and assigns it to the ReferenceName field.
func (o *GameCenterAchievementAttributes) SetReferenceName(v string) {
	o.ReferenceName = &v
}

// GetVendorIdentifier returns the VendorIdentifier field value if set, zero value otherwise.
func (o *GameCenterAchievementAttributes) GetVendorIdentifier() string {
	if o == nil || IsNil(o.VendorIdentifier) {
		var ret string
		return ret
	}
	return *o.VendorIdentifier
}

// GetVendorIdentifierOk returns a tuple with the VendorIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementAttributes) GetVendorIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.VendorIdentifier) {
		return nil, false
	}
	return o.VendorIdentifier, true
}

// HasVendorIdentifier returns a boolean if a field has been set.
func (o *GameCenterAchievementAttributes) HasVendorIdentifier() bool {
	if o != nil && !IsNil(o.VendorIdentifier) {
		return true
	}

	return false
}

// SetVendorIdentifier gets a reference to the given string and assigns it to the VendorIdentifier field.
func (o *GameCenterAchievementAttributes) SetVendorIdentifier(v string) {
	o.VendorIdentifier = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *GameCenterAchievementAttributes) GetPoints() int32 {
	if o == nil || IsNil(o.Points) {
		var ret int32
		return ret
	}
	return *o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementAttributes) GetPointsOk() (*int32, bool) {
	if o == nil || IsNil(o.Points) {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *GameCenterAchievementAttributes) HasPoints() bool {
	if o != nil && !IsNil(o.Points) {
		return true
	}

	return false
}

// SetPoints gets a reference to the given int32 and assigns it to the Points field.
func (o *GameCenterAchievementAttributes) SetPoints(v int32) {
	o.Points = &v
}

// GetShowBeforeEarned returns the ShowBeforeEarned field value if set, zero value otherwise.
func (o *GameCenterAchievementAttributes) GetShowBeforeEarned() bool {
	if o == nil || IsNil(o.ShowBeforeEarned) {
		var ret bool
		return ret
	}
	return *o.ShowBeforeEarned
}

// GetShowBeforeEarnedOk returns a tuple with the ShowBeforeEarned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementAttributes) GetShowBeforeEarnedOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowBeforeEarned) {
		return nil, false
	}
	return o.ShowBeforeEarned, true
}

// HasShowBeforeEarned returns a boolean if a field has been set.
func (o *GameCenterAchievementAttributes) HasShowBeforeEarned() bool {
	if o != nil && !IsNil(o.ShowBeforeEarned) {
		return true
	}

	return false
}

// SetShowBeforeEarned gets a reference to the given bool and assigns it to the ShowBeforeEarned field.
func (o *GameCenterAchievementAttributes) SetShowBeforeEarned(v bool) {
	o.ShowBeforeEarned = &v
}

// GetRepeatable returns the Repeatable field value if set, zero value otherwise.
func (o *GameCenterAchievementAttributes) GetRepeatable() bool {
	if o == nil || IsNil(o.Repeatable) {
		var ret bool
		return ret
	}
	return *o.Repeatable
}

// GetRepeatableOk returns a tuple with the Repeatable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementAttributes) GetRepeatableOk() (*bool, bool) {
	if o == nil || IsNil(o.Repeatable) {
		return nil, false
	}
	return o.Repeatable, true
}

// HasRepeatable returns a boolean if a field has been set.
func (o *GameCenterAchievementAttributes) HasRepeatable() bool {
	if o != nil && !IsNil(o.Repeatable) {
		return true
	}

	return false
}

// SetRepeatable gets a reference to the given bool and assigns it to the Repeatable field.
func (o *GameCenterAchievementAttributes) SetRepeatable(v bool) {
	o.Repeatable = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *GameCenterAchievementAttributes) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementAttributes) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *GameCenterAchievementAttributes) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *GameCenterAchievementAttributes) SetArchived(v bool) {
	o.Archived = &v
}

func (o GameCenterAchievementAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReferenceName) {
		toSerialize["referenceName"] = o.ReferenceName
	}
	if !IsNil(o.VendorIdentifier) {
		toSerialize["vendorIdentifier"] = o.VendorIdentifier
	}
	if !IsNil(o.Points) {
		toSerialize["points"] = o.Points
	}
	if !IsNil(o.ShowBeforeEarned) {
		toSerialize["showBeforeEarned"] = o.ShowBeforeEarned
	}
	if !IsNil(o.Repeatable) {
		toSerialize["repeatable"] = o.Repeatable
	}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	return toSerialize, nil
}

type NullableGameCenterAchievementAttributes struct {
	value *GameCenterAchievementAttributes
	isSet bool
}

func (v NullableGameCenterAchievementAttributes) Get() *GameCenterAchievementAttributes {
	return v.value
}

func (v *NullableGameCenterAchievementAttributes) Set(val *GameCenterAchievementAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementAttributes(val *GameCenterAchievementAttributes) *NullableGameCenterAchievementAttributes {
	return &NullableGameCenterAchievementAttributes{value: val, isSet: true}
}

func (v NullableGameCenterAchievementAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
