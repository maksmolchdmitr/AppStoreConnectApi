/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the ProfileAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileAttributes{}

// ProfileAttributes struct for ProfileAttributes
type ProfileAttributes struct {
	Name           *string           `json:"name,omitempty"`
	Platform       *BundleIdPlatform `json:"platform,omitempty"`
	ProfileType    *string           `json:"profileType,omitempty"`
	ProfileState   *string           `json:"profileState,omitempty"`
	ProfileContent *string           `json:"profileContent,omitempty"`
	Uuid           *string           `json:"uuid,omitempty"`
	CreatedDate    *time.Time        `json:"createdDate,omitempty"`
	ExpirationDate *time.Time        `json:"expirationDate,omitempty"`
}

// NewProfileAttributes instantiates a new ProfileAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileAttributes() *ProfileAttributes {
	this := ProfileAttributes{}
	return &this
}

// NewProfileAttributesWithDefaults instantiates a new ProfileAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileAttributesWithDefaults() *ProfileAttributes {
	this := ProfileAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProfileAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProfileAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProfileAttributes) SetName(v string) {
	o.Name = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *ProfileAttributes) GetPlatform() BundleIdPlatform {
	if o == nil || IsNil(o.Platform) {
		var ret BundleIdPlatform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileAttributes) GetPlatformOk() (*BundleIdPlatform, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *ProfileAttributes) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given BundleIdPlatform and assigns it to the Platform field.
func (o *ProfileAttributes) SetPlatform(v BundleIdPlatform) {
	o.Platform = &v
}

// GetProfileType returns the ProfileType field value if set, zero value otherwise.
func (o *ProfileAttributes) GetProfileType() string {
	if o == nil || IsNil(o.ProfileType) {
		var ret string
		return ret
	}
	return *o.ProfileType
}

// GetProfileTypeOk returns a tuple with the ProfileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileAttributes) GetProfileTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileType) {
		return nil, false
	}
	return o.ProfileType, true
}

// HasProfileType returns a boolean if a field has been set.
func (o *ProfileAttributes) HasProfileType() bool {
	if o != nil && !IsNil(o.ProfileType) {
		return true
	}

	return false
}

// SetProfileType gets a reference to the given string and assigns it to the ProfileType field.
func (o *ProfileAttributes) SetProfileType(v string) {
	o.ProfileType = &v
}

// GetProfileState returns the ProfileState field value if set, zero value otherwise.
func (o *ProfileAttributes) GetProfileState() string {
	if o == nil || IsNil(o.ProfileState) {
		var ret string
		return ret
	}
	return *o.ProfileState
}

// GetProfileStateOk returns a tuple with the ProfileState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileAttributes) GetProfileStateOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileState) {
		return nil, false
	}
	return o.ProfileState, true
}

// HasProfileState returns a boolean if a field has been set.
func (o *ProfileAttributes) HasProfileState() bool {
	if o != nil && !IsNil(o.ProfileState) {
		return true
	}

	return false
}

// SetProfileState gets a reference to the given string and assigns it to the ProfileState field.
func (o *ProfileAttributes) SetProfileState(v string) {
	o.ProfileState = &v
}

// GetProfileContent returns the ProfileContent field value if set, zero value otherwise.
func (o *ProfileAttributes) GetProfileContent() string {
	if o == nil || IsNil(o.ProfileContent) {
		var ret string
		return ret
	}
	return *o.ProfileContent
}

// GetProfileContentOk returns a tuple with the ProfileContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileAttributes) GetProfileContentOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileContent) {
		return nil, false
	}
	return o.ProfileContent, true
}

// HasProfileContent returns a boolean if a field has been set.
func (o *ProfileAttributes) HasProfileContent() bool {
	if o != nil && !IsNil(o.ProfileContent) {
		return true
	}

	return false
}

// SetProfileContent gets a reference to the given string and assigns it to the ProfileContent field.
func (o *ProfileAttributes) SetProfileContent(v string) {
	o.ProfileContent = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ProfileAttributes) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileAttributes) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ProfileAttributes) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ProfileAttributes) SetUuid(v string) {
	o.Uuid = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *ProfileAttributes) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileAttributes) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *ProfileAttributes) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *ProfileAttributes) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *ProfileAttributes) GetExpirationDate() time.Time {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileAttributes) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *ProfileAttributes) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given time.Time and assigns it to the ExpirationDate field.
func (o *ProfileAttributes) SetExpirationDate(v time.Time) {
	o.ExpirationDate = &v
}

func (o ProfileAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.ProfileType) {
		toSerialize["profileType"] = o.ProfileType
	}
	if !IsNil(o.ProfileState) {
		toSerialize["profileState"] = o.ProfileState
	}
	if !IsNil(o.ProfileContent) {
		toSerialize["profileContent"] = o.ProfileContent
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	return toSerialize, nil
}

type NullableProfileAttributes struct {
	value *ProfileAttributes
	isSet bool
}

func (v NullableProfileAttributes) Get() *ProfileAttributes {
	return v.value
}

func (v *NullableProfileAttributes) Set(val *ProfileAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileAttributes(val *ProfileAttributes) *NullableProfileAttributes {
	return &NullableProfileAttributes{value: val, isSet: true}
}

func (v NullableProfileAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
