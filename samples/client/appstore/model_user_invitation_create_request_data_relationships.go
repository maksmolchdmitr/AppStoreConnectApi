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

// checks if the UserInvitationCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserInvitationCreateRequestDataRelationships{}

// UserInvitationCreateRequestDataRelationships struct for UserInvitationCreateRequestDataRelationships
type UserInvitationCreateRequestDataRelationships struct {
	VisibleApps *UserInvitationCreateRequestDataRelationshipsVisibleApps `json:"visibleApps,omitempty"`
}

// NewUserInvitationCreateRequestDataRelationships instantiates a new UserInvitationCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInvitationCreateRequestDataRelationships() *UserInvitationCreateRequestDataRelationships {
	this := UserInvitationCreateRequestDataRelationships{}
	return &this
}

// NewUserInvitationCreateRequestDataRelationshipsWithDefaults instantiates a new UserInvitationCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInvitationCreateRequestDataRelationshipsWithDefaults() *UserInvitationCreateRequestDataRelationships {
	this := UserInvitationCreateRequestDataRelationships{}
	return &this
}

// GetVisibleApps returns the VisibleApps field value if set, zero value otherwise.
func (o *UserInvitationCreateRequestDataRelationships) GetVisibleApps() UserInvitationCreateRequestDataRelationshipsVisibleApps {
	if o == nil || IsNil(o.VisibleApps) {
		var ret UserInvitationCreateRequestDataRelationshipsVisibleApps
		return ret
	}
	return *o.VisibleApps
}

// GetVisibleAppsOk returns a tuple with the VisibleApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInvitationCreateRequestDataRelationships) GetVisibleAppsOk() (*UserInvitationCreateRequestDataRelationshipsVisibleApps, bool) {
	if o == nil || IsNil(o.VisibleApps) {
		return nil, false
	}
	return o.VisibleApps, true
}

// HasVisibleApps returns a boolean if a field has been set.
func (o *UserInvitationCreateRequestDataRelationships) HasVisibleApps() bool {
	if o != nil && !IsNil(o.VisibleApps) {
		return true
	}

	return false
}

// SetVisibleApps gets a reference to the given UserInvitationCreateRequestDataRelationshipsVisibleApps and assigns it to the VisibleApps field.
func (o *UserInvitationCreateRequestDataRelationships) SetVisibleApps(v UserInvitationCreateRequestDataRelationshipsVisibleApps) {
	o.VisibleApps = &v
}

func (o UserInvitationCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserInvitationCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VisibleApps) {
		toSerialize["visibleApps"] = o.VisibleApps
	}
	return toSerialize, nil
}

type NullableUserInvitationCreateRequestDataRelationships struct {
	value *UserInvitationCreateRequestDataRelationships
	isSet bool
}

func (v NullableUserInvitationCreateRequestDataRelationships) Get() *UserInvitationCreateRequestDataRelationships {
	return v.value
}

func (v *NullableUserInvitationCreateRequestDataRelationships) Set(val *UserInvitationCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInvitationCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInvitationCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInvitationCreateRequestDataRelationships(val *UserInvitationCreateRequestDataRelationships) *NullableUserInvitationCreateRequestDataRelationships {
	return &NullableUserInvitationCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableUserInvitationCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInvitationCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


