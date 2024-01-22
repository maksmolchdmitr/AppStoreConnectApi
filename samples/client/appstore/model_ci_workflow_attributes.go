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

// checks if the CiWorkflowAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiWorkflowAttributes{}

// CiWorkflowAttributes struct for CiWorkflowAttributes
type CiWorkflowAttributes struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	BranchStartCondition *CiBranchStartCondition `json:"branchStartCondition,omitempty"`
	TagStartCondition *CiTagStartCondition `json:"tagStartCondition,omitempty"`
	PullRequestStartCondition *CiPullRequestStartCondition `json:"pullRequestStartCondition,omitempty"`
	ScheduledStartCondition *CiScheduledStartCondition `json:"scheduledStartCondition,omitempty"`
	Actions []CiAction `json:"actions,omitempty"`
	IsEnabled *bool `json:"isEnabled,omitempty"`
	IsLockedForEditing *bool `json:"isLockedForEditing,omitempty"`
	Clean *bool `json:"clean,omitempty"`
	ContainerFilePath *string `json:"containerFilePath,omitempty"`
	LastModifiedDate *time.Time `json:"lastModifiedDate,omitempty"`
}

// NewCiWorkflowAttributes instantiates a new CiWorkflowAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiWorkflowAttributes() *CiWorkflowAttributes {
	this := CiWorkflowAttributes{}
	return &this
}

// NewCiWorkflowAttributesWithDefaults instantiates a new CiWorkflowAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiWorkflowAttributesWithDefaults() *CiWorkflowAttributes {
	this := CiWorkflowAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CiWorkflowAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CiWorkflowAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CiWorkflowAttributes) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CiWorkflowAttributes) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CiWorkflowAttributes) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CiWorkflowAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetBranchStartCondition returns the BranchStartCondition field value if set, zero value otherwise.
func (o *CiWorkflowAttributes) GetBranchStartCondition() CiBranchStartCondition {
	if o == nil || IsNil(o.BranchStartCondition) {
		var ret CiBranchStartCondition
		return ret
	}
	return *o.BranchStartCondition
}

// GetBranchStartConditionOk returns a tuple with the BranchStartCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowAttributes) GetBranchStartConditionOk() (*CiBranchStartCondition, bool) {
	if o == nil || IsNil(o.BranchStartCondition) {
		return nil, false
	}
	return o.BranchStartCondition, true
}

// HasBranchStartCondition returns a boolean if a field has been set.
func (o *CiWorkflowAttributes) HasBranchStartCondition() bool {
	if o != nil && !IsNil(o.BranchStartCondition) {
		return true
	}

	return false
}

// SetBranchStartCondition gets a reference to the given CiBranchStartCondition and assigns it to the BranchStartCondition field.
func (o *CiWorkflowAttributes) SetBranchStartCondition(v CiBranchStartCondition) {
	o.BranchStartCondition = &v
}

// GetTagStartCondition returns the TagStartCondition field value if set, zero value otherwise.
func (o *CiWorkflowAttributes) GetTagStartCondition() CiTagStartCondition {
	if o == nil || IsNil(o.TagStartCondition) {
		var ret CiTagStartCondition
		return ret
	}
	return *o.TagStartCondition
}

// GetTagStartConditionOk returns a tuple with the TagStartCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowAttributes) GetTagStartConditionOk() (*CiTagStartCondition, bool) {
	if o == nil || IsNil(o.TagStartCondition) {
		return nil, false
	}
	return o.TagStartCondition, true
}

// HasTagStartCondition returns a boolean if a field has been set.
func (o *CiWorkflowAttributes) HasTagStartCondition() bool {
	if o != nil && !IsNil(o.TagStartCondition) {
		return true
	}

	return false
}

// SetTagStartCondition gets a reference to the given CiTagStartCondition and assigns it to the TagStartCondition field.
func (o *CiWorkflowAttributes) SetTagStartCondition(v CiTagStartCondition) {
	o.TagStartCondition = &v
}

// GetPullRequestStartCondition returns the PullRequestStartCondition field value if set, zero value otherwise.
func (o *CiWorkflowAttributes) GetPullRequestStartCondition() CiPullRequestStartCondition {
	if o == nil || IsNil(o.PullRequestStartCondition) {
		var ret CiPullRequestStartCondition
		return ret
	}
	return *o.PullRequestStartCondition
}

// GetPullRequestStartConditionOk returns a tuple with the PullRequestStartCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowAttributes) GetPullRequestStartConditionOk() (*CiPullRequestStartCondition, bool) {
	if o == nil || IsNil(o.PullRequestStartCondition) {
		return nil, false
	}
	return o.PullRequestStartCondition, true
}

// HasPullRequestStartCondition returns a boolean if a field has been set.
func (o *CiWorkflowAttributes) HasPullRequestStartCondition() bool {
	if o != nil && !IsNil(o.PullRequestStartCondition) {
		return true
	}

	return false
}

// SetPullRequestStartCondition gets a reference to the given CiPullRequestStartCondition and assigns it to the PullRequestStartCondition field.
func (o *CiWorkflowAttributes) SetPullRequestStartCondition(v CiPullRequestStartCondition) {
	o.PullRequestStartCondition = &v
}

// GetScheduledStartCondition returns the ScheduledStartCondition field value if set, zero value otherwise.
func (o *CiWorkflowAttributes) GetScheduledStartCondition() CiScheduledStartCondition {
	if o == nil || IsNil(o.ScheduledStartCondition) {
		var ret CiScheduledStartCondition
		return ret
	}
	return *o.ScheduledStartCondition
}

// GetScheduledStartConditionOk returns a tuple with the ScheduledStartCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowAttributes) GetScheduledStartConditionOk() (*CiScheduledStartCondition, bool) {
	if o == nil || IsNil(o.ScheduledStartCondition) {
		return nil, false
	}
	return o.ScheduledStartCondition, true
}

// HasScheduledStartCondition returns a boolean if a field has been set.
func (o *CiWorkflowAttributes) HasScheduledStartCondition() bool {
	if o != nil && !IsNil(o.ScheduledStartCondition) {
		return true
	}

	return false
}

// SetScheduledStartCondition gets a reference to the given CiScheduledStartCondition and assigns it to the ScheduledStartCondition field.
func (o *CiWorkflowAttributes) SetScheduledStartCondition(v CiScheduledStartCondition) {
	o.ScheduledStartCondition = &v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *CiWorkflowAttributes) GetActions() []CiAction {
	if o == nil || IsNil(o.Actions) {
		var ret []CiAction
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowAttributes) GetActionsOk() ([]CiAction, bool) {
	if o == nil || IsNil(o.Actions) {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *CiWorkflowAttributes) HasActions() bool {
	if o != nil && !IsNil(o.Actions) {
		return true
	}

	return false
}

// SetActions gets a reference to the given []CiAction and assigns it to the Actions field.
func (o *CiWorkflowAttributes) SetActions(v []CiAction) {
	o.Actions = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *CiWorkflowAttributes) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowAttributes) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *CiWorkflowAttributes) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *CiWorkflowAttributes) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetIsLockedForEditing returns the IsLockedForEditing field value if set, zero value otherwise.
func (o *CiWorkflowAttributes) GetIsLockedForEditing() bool {
	if o == nil || IsNil(o.IsLockedForEditing) {
		var ret bool
		return ret
	}
	return *o.IsLockedForEditing
}

// GetIsLockedForEditingOk returns a tuple with the IsLockedForEditing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowAttributes) GetIsLockedForEditingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLockedForEditing) {
		return nil, false
	}
	return o.IsLockedForEditing, true
}

// HasIsLockedForEditing returns a boolean if a field has been set.
func (o *CiWorkflowAttributes) HasIsLockedForEditing() bool {
	if o != nil && !IsNil(o.IsLockedForEditing) {
		return true
	}

	return false
}

// SetIsLockedForEditing gets a reference to the given bool and assigns it to the IsLockedForEditing field.
func (o *CiWorkflowAttributes) SetIsLockedForEditing(v bool) {
	o.IsLockedForEditing = &v
}

// GetClean returns the Clean field value if set, zero value otherwise.
func (o *CiWorkflowAttributes) GetClean() bool {
	if o == nil || IsNil(o.Clean) {
		var ret bool
		return ret
	}
	return *o.Clean
}

// GetCleanOk returns a tuple with the Clean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowAttributes) GetCleanOk() (*bool, bool) {
	if o == nil || IsNil(o.Clean) {
		return nil, false
	}
	return o.Clean, true
}

// HasClean returns a boolean if a field has been set.
func (o *CiWorkflowAttributes) HasClean() bool {
	if o != nil && !IsNil(o.Clean) {
		return true
	}

	return false
}

// SetClean gets a reference to the given bool and assigns it to the Clean field.
func (o *CiWorkflowAttributes) SetClean(v bool) {
	o.Clean = &v
}

// GetContainerFilePath returns the ContainerFilePath field value if set, zero value otherwise.
func (o *CiWorkflowAttributes) GetContainerFilePath() string {
	if o == nil || IsNil(o.ContainerFilePath) {
		var ret string
		return ret
	}
	return *o.ContainerFilePath
}

// GetContainerFilePathOk returns a tuple with the ContainerFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowAttributes) GetContainerFilePathOk() (*string, bool) {
	if o == nil || IsNil(o.ContainerFilePath) {
		return nil, false
	}
	return o.ContainerFilePath, true
}

// HasContainerFilePath returns a boolean if a field has been set.
func (o *CiWorkflowAttributes) HasContainerFilePath() bool {
	if o != nil && !IsNil(o.ContainerFilePath) {
		return true
	}

	return false
}

// SetContainerFilePath gets a reference to the given string and assigns it to the ContainerFilePath field.
func (o *CiWorkflowAttributes) SetContainerFilePath(v string) {
	o.ContainerFilePath = &v
}

// GetLastModifiedDate returns the LastModifiedDate field value if set, zero value otherwise.
func (o *CiWorkflowAttributes) GetLastModifiedDate() time.Time {
	if o == nil || IsNil(o.LastModifiedDate) {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDate
}

// GetLastModifiedDateOk returns a tuple with the LastModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowAttributes) GetLastModifiedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedDate) {
		return nil, false
	}
	return o.LastModifiedDate, true
}

// HasLastModifiedDate returns a boolean if a field has been set.
func (o *CiWorkflowAttributes) HasLastModifiedDate() bool {
	if o != nil && !IsNil(o.LastModifiedDate) {
		return true
	}

	return false
}

// SetLastModifiedDate gets a reference to the given time.Time and assigns it to the LastModifiedDate field.
func (o *CiWorkflowAttributes) SetLastModifiedDate(v time.Time) {
	o.LastModifiedDate = &v
}

func (o CiWorkflowAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiWorkflowAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.BranchStartCondition) {
		toSerialize["branchStartCondition"] = o.BranchStartCondition
	}
	if !IsNil(o.TagStartCondition) {
		toSerialize["tagStartCondition"] = o.TagStartCondition
	}
	if !IsNil(o.PullRequestStartCondition) {
		toSerialize["pullRequestStartCondition"] = o.PullRequestStartCondition
	}
	if !IsNil(o.ScheduledStartCondition) {
		toSerialize["scheduledStartCondition"] = o.ScheduledStartCondition
	}
	if !IsNil(o.Actions) {
		toSerialize["actions"] = o.Actions
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !IsNil(o.IsLockedForEditing) {
		toSerialize["isLockedForEditing"] = o.IsLockedForEditing
	}
	if !IsNil(o.Clean) {
		toSerialize["clean"] = o.Clean
	}
	if !IsNil(o.ContainerFilePath) {
		toSerialize["containerFilePath"] = o.ContainerFilePath
	}
	if !IsNil(o.LastModifiedDate) {
		toSerialize["lastModifiedDate"] = o.LastModifiedDate
	}
	return toSerialize, nil
}

type NullableCiWorkflowAttributes struct {
	value *CiWorkflowAttributes
	isSet bool
}

func (v NullableCiWorkflowAttributes) Get() *CiWorkflowAttributes {
	return v.value
}

func (v *NullableCiWorkflowAttributes) Set(val *CiWorkflowAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCiWorkflowAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCiWorkflowAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiWorkflowAttributes(val *CiWorkflowAttributes) *NullableCiWorkflowAttributes {
	return &NullableCiWorkflowAttributes{value: val, isSet: true}
}

func (v NullableCiWorkflowAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiWorkflowAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


