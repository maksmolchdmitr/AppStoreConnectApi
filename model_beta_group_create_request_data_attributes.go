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

// checks if the BetaGroupCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaGroupCreateRequestDataAttributes{}

// BetaGroupCreateRequestDataAttributes struct for BetaGroupCreateRequestDataAttributes
type BetaGroupCreateRequestDataAttributes struct {
	Name                   string `json:"name"`
	IsInternalGroup        *bool  `json:"isInternalGroup,omitempty"`
	HasAccessToAllBuilds   *bool  `json:"hasAccessToAllBuilds,omitempty"`
	PublicLinkEnabled      *bool  `json:"publicLinkEnabled,omitempty"`
	PublicLinkLimitEnabled *bool  `json:"publicLinkLimitEnabled,omitempty"`
	PublicLinkLimit        *int32 `json:"publicLinkLimit,omitempty"`
	FeedbackEnabled        *bool  `json:"feedbackEnabled,omitempty"`
}

// NewBetaGroupCreateRequestDataAttributes instantiates a new BetaGroupCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaGroupCreateRequestDataAttributes(name string) *BetaGroupCreateRequestDataAttributes {
	this := BetaGroupCreateRequestDataAttributes{}
	this.Name = name
	return &this
}

// NewBetaGroupCreateRequestDataAttributesWithDefaults instantiates a new BetaGroupCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaGroupCreateRequestDataAttributesWithDefaults() *BetaGroupCreateRequestDataAttributes {
	this := BetaGroupCreateRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *BetaGroupCreateRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BetaGroupCreateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BetaGroupCreateRequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetIsInternalGroup returns the IsInternalGroup field value if set, zero value otherwise.
func (o *BetaGroupCreateRequestDataAttributes) GetIsInternalGroup() bool {
	if o == nil || IsNil(o.IsInternalGroup) {
		var ret bool
		return ret
	}
	return *o.IsInternalGroup
}

// GetIsInternalGroupOk returns a tuple with the IsInternalGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupCreateRequestDataAttributes) GetIsInternalGroupOk() (*bool, bool) {
	if o == nil || IsNil(o.IsInternalGroup) {
		return nil, false
	}
	return o.IsInternalGroup, true
}

// HasIsInternalGroup returns a boolean if a field has been set.
func (o *BetaGroupCreateRequestDataAttributes) HasIsInternalGroup() bool {
	if o != nil && !IsNil(o.IsInternalGroup) {
		return true
	}

	return false
}

// SetIsInternalGroup gets a reference to the given bool and assigns it to the IsInternalGroup field.
func (o *BetaGroupCreateRequestDataAttributes) SetIsInternalGroup(v bool) {
	o.IsInternalGroup = &v
}

// GetHasAccessToAllBuilds returns the HasAccessToAllBuilds field value if set, zero value otherwise.
func (o *BetaGroupCreateRequestDataAttributes) GetHasAccessToAllBuilds() bool {
	if o == nil || IsNil(o.HasAccessToAllBuilds) {
		var ret bool
		return ret
	}
	return *o.HasAccessToAllBuilds
}

// GetHasAccessToAllBuildsOk returns a tuple with the HasAccessToAllBuilds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupCreateRequestDataAttributes) GetHasAccessToAllBuildsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasAccessToAllBuilds) {
		return nil, false
	}
	return o.HasAccessToAllBuilds, true
}

// HasHasAccessToAllBuilds returns a boolean if a field has been set.
func (o *BetaGroupCreateRequestDataAttributes) HasHasAccessToAllBuilds() bool {
	if o != nil && !IsNil(o.HasAccessToAllBuilds) {
		return true
	}

	return false
}

// SetHasAccessToAllBuilds gets a reference to the given bool and assigns it to the HasAccessToAllBuilds field.
func (o *BetaGroupCreateRequestDataAttributes) SetHasAccessToAllBuilds(v bool) {
	o.HasAccessToAllBuilds = &v
}

// GetPublicLinkEnabled returns the PublicLinkEnabled field value if set, zero value otherwise.
func (o *BetaGroupCreateRequestDataAttributes) GetPublicLinkEnabled() bool {
	if o == nil || IsNil(o.PublicLinkEnabled) {
		var ret bool
		return ret
	}
	return *o.PublicLinkEnabled
}

// GetPublicLinkEnabledOk returns a tuple with the PublicLinkEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupCreateRequestDataAttributes) GetPublicLinkEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PublicLinkEnabled) {
		return nil, false
	}
	return o.PublicLinkEnabled, true
}

// HasPublicLinkEnabled returns a boolean if a field has been set.
func (o *BetaGroupCreateRequestDataAttributes) HasPublicLinkEnabled() bool {
	if o != nil && !IsNil(o.PublicLinkEnabled) {
		return true
	}

	return false
}

// SetPublicLinkEnabled gets a reference to the given bool and assigns it to the PublicLinkEnabled field.
func (o *BetaGroupCreateRequestDataAttributes) SetPublicLinkEnabled(v bool) {
	o.PublicLinkEnabled = &v
}

// GetPublicLinkLimitEnabled returns the PublicLinkLimitEnabled field value if set, zero value otherwise.
func (o *BetaGroupCreateRequestDataAttributes) GetPublicLinkLimitEnabled() bool {
	if o == nil || IsNil(o.PublicLinkLimitEnabled) {
		var ret bool
		return ret
	}
	return *o.PublicLinkLimitEnabled
}

// GetPublicLinkLimitEnabledOk returns a tuple with the PublicLinkLimitEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupCreateRequestDataAttributes) GetPublicLinkLimitEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PublicLinkLimitEnabled) {
		return nil, false
	}
	return o.PublicLinkLimitEnabled, true
}

// HasPublicLinkLimitEnabled returns a boolean if a field has been set.
func (o *BetaGroupCreateRequestDataAttributes) HasPublicLinkLimitEnabled() bool {
	if o != nil && !IsNil(o.PublicLinkLimitEnabled) {
		return true
	}

	return false
}

// SetPublicLinkLimitEnabled gets a reference to the given bool and assigns it to the PublicLinkLimitEnabled field.
func (o *BetaGroupCreateRequestDataAttributes) SetPublicLinkLimitEnabled(v bool) {
	o.PublicLinkLimitEnabled = &v
}

// GetPublicLinkLimit returns the PublicLinkLimit field value if set, zero value otherwise.
func (o *BetaGroupCreateRequestDataAttributes) GetPublicLinkLimit() int32 {
	if o == nil || IsNil(o.PublicLinkLimit) {
		var ret int32
		return ret
	}
	return *o.PublicLinkLimit
}

// GetPublicLinkLimitOk returns a tuple with the PublicLinkLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupCreateRequestDataAttributes) GetPublicLinkLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.PublicLinkLimit) {
		return nil, false
	}
	return o.PublicLinkLimit, true
}

// HasPublicLinkLimit returns a boolean if a field has been set.
func (o *BetaGroupCreateRequestDataAttributes) HasPublicLinkLimit() bool {
	if o != nil && !IsNil(o.PublicLinkLimit) {
		return true
	}

	return false
}

// SetPublicLinkLimit gets a reference to the given int32 and assigns it to the PublicLinkLimit field.
func (o *BetaGroupCreateRequestDataAttributes) SetPublicLinkLimit(v int32) {
	o.PublicLinkLimit = &v
}

// GetFeedbackEnabled returns the FeedbackEnabled field value if set, zero value otherwise.
func (o *BetaGroupCreateRequestDataAttributes) GetFeedbackEnabled() bool {
	if o == nil || IsNil(o.FeedbackEnabled) {
		var ret bool
		return ret
	}
	return *o.FeedbackEnabled
}

// GetFeedbackEnabledOk returns a tuple with the FeedbackEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupCreateRequestDataAttributes) GetFeedbackEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FeedbackEnabled) {
		return nil, false
	}
	return o.FeedbackEnabled, true
}

// HasFeedbackEnabled returns a boolean if a field has been set.
func (o *BetaGroupCreateRequestDataAttributes) HasFeedbackEnabled() bool {
	if o != nil && !IsNil(o.FeedbackEnabled) {
		return true
	}

	return false
}

// SetFeedbackEnabled gets a reference to the given bool and assigns it to the FeedbackEnabled field.
func (o *BetaGroupCreateRequestDataAttributes) SetFeedbackEnabled(v bool) {
	o.FeedbackEnabled = &v
}

func (o BetaGroupCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaGroupCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.IsInternalGroup) {
		toSerialize["isInternalGroup"] = o.IsInternalGroup
	}
	if !IsNil(o.HasAccessToAllBuilds) {
		toSerialize["hasAccessToAllBuilds"] = o.HasAccessToAllBuilds
	}
	if !IsNil(o.PublicLinkEnabled) {
		toSerialize["publicLinkEnabled"] = o.PublicLinkEnabled
	}
	if !IsNil(o.PublicLinkLimitEnabled) {
		toSerialize["publicLinkLimitEnabled"] = o.PublicLinkLimitEnabled
	}
	if !IsNil(o.PublicLinkLimit) {
		toSerialize["publicLinkLimit"] = o.PublicLinkLimit
	}
	if !IsNil(o.FeedbackEnabled) {
		toSerialize["feedbackEnabled"] = o.FeedbackEnabled
	}
	return toSerialize, nil
}

type NullableBetaGroupCreateRequestDataAttributes struct {
	value *BetaGroupCreateRequestDataAttributes
	isSet bool
}

func (v NullableBetaGroupCreateRequestDataAttributes) Get() *BetaGroupCreateRequestDataAttributes {
	return v.value
}

func (v *NullableBetaGroupCreateRequestDataAttributes) Set(val *BetaGroupCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaGroupCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaGroupCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaGroupCreateRequestDataAttributes(val *BetaGroupCreateRequestDataAttributes) *NullableBetaGroupCreateRequestDataAttributes {
	return &NullableBetaGroupCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableBetaGroupCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaGroupCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
