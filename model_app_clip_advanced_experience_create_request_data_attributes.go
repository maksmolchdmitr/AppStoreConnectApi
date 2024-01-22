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

// checks if the AppClipAdvancedExperienceCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAdvancedExperienceCreateRequestDataAttributes{}

// AppClipAdvancedExperienceCreateRequestDataAttributes struct for AppClipAdvancedExperienceCreateRequestDataAttributes
type AppClipAdvancedExperienceCreateRequestDataAttributes struct {
	Link             string                                    `json:"link"`
	Action           *AppClipAction                            `json:"action,omitempty"`
	IsPoweredBy      bool                                      `json:"isPoweredBy"`
	Place            *AppClipAdvancedExperienceAttributesPlace `json:"place,omitempty"`
	BusinessCategory *string                                   `json:"businessCategory,omitempty"`
	DefaultLanguage  AppClipAdvancedExperienceLanguage         `json:"defaultLanguage"`
}

// NewAppClipAdvancedExperienceCreateRequestDataAttributes instantiates a new AppClipAdvancedExperienceCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAdvancedExperienceCreateRequestDataAttributes(link string, isPoweredBy bool, defaultLanguage AppClipAdvancedExperienceLanguage) *AppClipAdvancedExperienceCreateRequestDataAttributes {
	this := AppClipAdvancedExperienceCreateRequestDataAttributes{}
	this.Link = link
	this.IsPoweredBy = isPoweredBy
	this.DefaultLanguage = defaultLanguage
	return &this
}

// NewAppClipAdvancedExperienceCreateRequestDataAttributesWithDefaults instantiates a new AppClipAdvancedExperienceCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAdvancedExperienceCreateRequestDataAttributesWithDefaults() *AppClipAdvancedExperienceCreateRequestDataAttributes {
	this := AppClipAdvancedExperienceCreateRequestDataAttributes{}
	return &this
}

// GetLink returns the Link field value
func (o *AppClipAdvancedExperienceCreateRequestDataAttributes) GetLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceCreateRequestDataAttributes) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *AppClipAdvancedExperienceCreateRequestDataAttributes) SetLink(v string) {
	o.Link = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceCreateRequestDataAttributes) GetAction() AppClipAction {
	if o == nil || IsNil(o.Action) {
		var ret AppClipAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceCreateRequestDataAttributes) GetActionOk() (*AppClipAction, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceCreateRequestDataAttributes) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given AppClipAction and assigns it to the Action field.
func (o *AppClipAdvancedExperienceCreateRequestDataAttributes) SetAction(v AppClipAction) {
	o.Action = &v
}

// GetIsPoweredBy returns the IsPoweredBy field value
func (o *AppClipAdvancedExperienceCreateRequestDataAttributes) GetIsPoweredBy() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPoweredBy
}

// GetIsPoweredByOk returns a tuple with the IsPoweredBy field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceCreateRequestDataAttributes) GetIsPoweredByOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPoweredBy, true
}

// SetIsPoweredBy sets field value
func (o *AppClipAdvancedExperienceCreateRequestDataAttributes) SetIsPoweredBy(v bool) {
	o.IsPoweredBy = v
}

// GetPlace returns the Place field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceCreateRequestDataAttributes) GetPlace() AppClipAdvancedExperienceAttributesPlace {
	if o == nil || IsNil(o.Place) {
		var ret AppClipAdvancedExperienceAttributesPlace
		return ret
	}
	return *o.Place
}

// GetPlaceOk returns a tuple with the Place field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceCreateRequestDataAttributes) GetPlaceOk() (*AppClipAdvancedExperienceAttributesPlace, bool) {
	if o == nil || IsNil(o.Place) {
		return nil, false
	}
	return o.Place, true
}

// HasPlace returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceCreateRequestDataAttributes) HasPlace() bool {
	if o != nil && !IsNil(o.Place) {
		return true
	}

	return false
}

// SetPlace gets a reference to the given AppClipAdvancedExperienceAttributesPlace and assigns it to the Place field.
func (o *AppClipAdvancedExperienceCreateRequestDataAttributes) SetPlace(v AppClipAdvancedExperienceAttributesPlace) {
	o.Place = &v
}

// GetBusinessCategory returns the BusinessCategory field value if set, zero value otherwise.
func (o *AppClipAdvancedExperienceCreateRequestDataAttributes) GetBusinessCategory() string {
	if o == nil || IsNil(o.BusinessCategory) {
		var ret string
		return ret
	}
	return *o.BusinessCategory
}

// GetBusinessCategoryOk returns a tuple with the BusinessCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceCreateRequestDataAttributes) GetBusinessCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessCategory) {
		return nil, false
	}
	return o.BusinessCategory, true
}

// HasBusinessCategory returns a boolean if a field has been set.
func (o *AppClipAdvancedExperienceCreateRequestDataAttributes) HasBusinessCategory() bool {
	if o != nil && !IsNil(o.BusinessCategory) {
		return true
	}

	return false
}

// SetBusinessCategory gets a reference to the given string and assigns it to the BusinessCategory field.
func (o *AppClipAdvancedExperienceCreateRequestDataAttributes) SetBusinessCategory(v string) {
	o.BusinessCategory = &v
}

// GetDefaultLanguage returns the DefaultLanguage field value
func (o *AppClipAdvancedExperienceCreateRequestDataAttributes) GetDefaultLanguage() AppClipAdvancedExperienceLanguage {
	if o == nil {
		var ret AppClipAdvancedExperienceLanguage
		return ret
	}

	return o.DefaultLanguage
}

// GetDefaultLanguageOk returns a tuple with the DefaultLanguage field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceCreateRequestDataAttributes) GetDefaultLanguageOk() (*AppClipAdvancedExperienceLanguage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultLanguage, true
}

// SetDefaultLanguage sets field value
func (o *AppClipAdvancedExperienceCreateRequestDataAttributes) SetDefaultLanguage(v AppClipAdvancedExperienceLanguage) {
	o.DefaultLanguage = v
}

func (o AppClipAdvancedExperienceCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAdvancedExperienceCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["link"] = o.Link
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	toSerialize["isPoweredBy"] = o.IsPoweredBy
	if !IsNil(o.Place) {
		toSerialize["place"] = o.Place
	}
	if !IsNil(o.BusinessCategory) {
		toSerialize["businessCategory"] = o.BusinessCategory
	}
	toSerialize["defaultLanguage"] = o.DefaultLanguage
	return toSerialize, nil
}

type NullableAppClipAdvancedExperienceCreateRequestDataAttributes struct {
	value *AppClipAdvancedExperienceCreateRequestDataAttributes
	isSet bool
}

func (v NullableAppClipAdvancedExperienceCreateRequestDataAttributes) Get() *AppClipAdvancedExperienceCreateRequestDataAttributes {
	return v.value
}

func (v *NullableAppClipAdvancedExperienceCreateRequestDataAttributes) Set(val *AppClipAdvancedExperienceCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperienceCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperienceCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperienceCreateRequestDataAttributes(val *AppClipAdvancedExperienceCreateRequestDataAttributes) *NullableAppClipAdvancedExperienceCreateRequestDataAttributes {
	return &NullableAppClipAdvancedExperienceCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperienceCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperienceCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
