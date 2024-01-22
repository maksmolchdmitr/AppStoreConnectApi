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

// checks if the CiBuildActionAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiBuildActionAttributes{}

// CiBuildActionAttributes struct for CiBuildActionAttributes
type CiBuildActionAttributes struct {
	Name *string `json:"name,omitempty"`
	ActionType *CiActionType `json:"actionType,omitempty"`
	StartedDate *time.Time `json:"startedDate,omitempty"`
	FinishedDate *time.Time `json:"finishedDate,omitempty"`
	IssueCounts *CiIssueCounts `json:"issueCounts,omitempty"`
	ExecutionProgress *CiExecutionProgress `json:"executionProgress,omitempty"`
	CompletionStatus *CiCompletionStatus `json:"completionStatus,omitempty"`
	IsRequiredToPass *bool `json:"isRequiredToPass,omitempty"`
}

// NewCiBuildActionAttributes instantiates a new CiBuildActionAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiBuildActionAttributes() *CiBuildActionAttributes {
	this := CiBuildActionAttributes{}
	return &this
}

// NewCiBuildActionAttributesWithDefaults instantiates a new CiBuildActionAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiBuildActionAttributesWithDefaults() *CiBuildActionAttributes {
	this := CiBuildActionAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CiBuildActionAttributes) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildActionAttributes) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CiBuildActionAttributes) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CiBuildActionAttributes) SetName(v string) {
	o.Name = &v
}

// GetActionType returns the ActionType field value if set, zero value otherwise.
func (o *CiBuildActionAttributes) GetActionType() CiActionType {
	if o == nil || IsNil(o.ActionType) {
		var ret CiActionType
		return ret
	}
	return *o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildActionAttributes) GetActionTypeOk() (*CiActionType, bool) {
	if o == nil || IsNil(o.ActionType) {
		return nil, false
	}
	return o.ActionType, true
}

// HasActionType returns a boolean if a field has been set.
func (o *CiBuildActionAttributes) HasActionType() bool {
	if o != nil && !IsNil(o.ActionType) {
		return true
	}

	return false
}

// SetActionType gets a reference to the given CiActionType and assigns it to the ActionType field.
func (o *CiBuildActionAttributes) SetActionType(v CiActionType) {
	o.ActionType = &v
}

// GetStartedDate returns the StartedDate field value if set, zero value otherwise.
func (o *CiBuildActionAttributes) GetStartedDate() time.Time {
	if o == nil || IsNil(o.StartedDate) {
		var ret time.Time
		return ret
	}
	return *o.StartedDate
}

// GetStartedDateOk returns a tuple with the StartedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildActionAttributes) GetStartedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartedDate) {
		return nil, false
	}
	return o.StartedDate, true
}

// HasStartedDate returns a boolean if a field has been set.
func (o *CiBuildActionAttributes) HasStartedDate() bool {
	if o != nil && !IsNil(o.StartedDate) {
		return true
	}

	return false
}

// SetStartedDate gets a reference to the given time.Time and assigns it to the StartedDate field.
func (o *CiBuildActionAttributes) SetStartedDate(v time.Time) {
	o.StartedDate = &v
}

// GetFinishedDate returns the FinishedDate field value if set, zero value otherwise.
func (o *CiBuildActionAttributes) GetFinishedDate() time.Time {
	if o == nil || IsNil(o.FinishedDate) {
		var ret time.Time
		return ret
	}
	return *o.FinishedDate
}

// GetFinishedDateOk returns a tuple with the FinishedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildActionAttributes) GetFinishedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FinishedDate) {
		return nil, false
	}
	return o.FinishedDate, true
}

// HasFinishedDate returns a boolean if a field has been set.
func (o *CiBuildActionAttributes) HasFinishedDate() bool {
	if o != nil && !IsNil(o.FinishedDate) {
		return true
	}

	return false
}

// SetFinishedDate gets a reference to the given time.Time and assigns it to the FinishedDate field.
func (o *CiBuildActionAttributes) SetFinishedDate(v time.Time) {
	o.FinishedDate = &v
}

// GetIssueCounts returns the IssueCounts field value if set, zero value otherwise.
func (o *CiBuildActionAttributes) GetIssueCounts() CiIssueCounts {
	if o == nil || IsNil(o.IssueCounts) {
		var ret CiIssueCounts
		return ret
	}
	return *o.IssueCounts
}

// GetIssueCountsOk returns a tuple with the IssueCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildActionAttributes) GetIssueCountsOk() (*CiIssueCounts, bool) {
	if o == nil || IsNil(o.IssueCounts) {
		return nil, false
	}
	return o.IssueCounts, true
}

// HasIssueCounts returns a boolean if a field has been set.
func (o *CiBuildActionAttributes) HasIssueCounts() bool {
	if o != nil && !IsNil(o.IssueCounts) {
		return true
	}

	return false
}

// SetIssueCounts gets a reference to the given CiIssueCounts and assigns it to the IssueCounts field.
func (o *CiBuildActionAttributes) SetIssueCounts(v CiIssueCounts) {
	o.IssueCounts = &v
}

// GetExecutionProgress returns the ExecutionProgress field value if set, zero value otherwise.
func (o *CiBuildActionAttributes) GetExecutionProgress() CiExecutionProgress {
	if o == nil || IsNil(o.ExecutionProgress) {
		var ret CiExecutionProgress
		return ret
	}
	return *o.ExecutionProgress
}

// GetExecutionProgressOk returns a tuple with the ExecutionProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildActionAttributes) GetExecutionProgressOk() (*CiExecutionProgress, bool) {
	if o == nil || IsNil(o.ExecutionProgress) {
		return nil, false
	}
	return o.ExecutionProgress, true
}

// HasExecutionProgress returns a boolean if a field has been set.
func (o *CiBuildActionAttributes) HasExecutionProgress() bool {
	if o != nil && !IsNil(o.ExecutionProgress) {
		return true
	}

	return false
}

// SetExecutionProgress gets a reference to the given CiExecutionProgress and assigns it to the ExecutionProgress field.
func (o *CiBuildActionAttributes) SetExecutionProgress(v CiExecutionProgress) {
	o.ExecutionProgress = &v
}

// GetCompletionStatus returns the CompletionStatus field value if set, zero value otherwise.
func (o *CiBuildActionAttributes) GetCompletionStatus() CiCompletionStatus {
	if o == nil || IsNil(o.CompletionStatus) {
		var ret CiCompletionStatus
		return ret
	}
	return *o.CompletionStatus
}

// GetCompletionStatusOk returns a tuple with the CompletionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildActionAttributes) GetCompletionStatusOk() (*CiCompletionStatus, bool) {
	if o == nil || IsNil(o.CompletionStatus) {
		return nil, false
	}
	return o.CompletionStatus, true
}

// HasCompletionStatus returns a boolean if a field has been set.
func (o *CiBuildActionAttributes) HasCompletionStatus() bool {
	if o != nil && !IsNil(o.CompletionStatus) {
		return true
	}

	return false
}

// SetCompletionStatus gets a reference to the given CiCompletionStatus and assigns it to the CompletionStatus field.
func (o *CiBuildActionAttributes) SetCompletionStatus(v CiCompletionStatus) {
	o.CompletionStatus = &v
}

// GetIsRequiredToPass returns the IsRequiredToPass field value if set, zero value otherwise.
func (o *CiBuildActionAttributes) GetIsRequiredToPass() bool {
	if o == nil || IsNil(o.IsRequiredToPass) {
		var ret bool
		return ret
	}
	return *o.IsRequiredToPass
}

// GetIsRequiredToPassOk returns a tuple with the IsRequiredToPass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildActionAttributes) GetIsRequiredToPassOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRequiredToPass) {
		return nil, false
	}
	return o.IsRequiredToPass, true
}

// HasIsRequiredToPass returns a boolean if a field has been set.
func (o *CiBuildActionAttributes) HasIsRequiredToPass() bool {
	if o != nil && !IsNil(o.IsRequiredToPass) {
		return true
	}

	return false
}

// SetIsRequiredToPass gets a reference to the given bool and assigns it to the IsRequiredToPass field.
func (o *CiBuildActionAttributes) SetIsRequiredToPass(v bool) {
	o.IsRequiredToPass = &v
}

func (o CiBuildActionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiBuildActionAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ActionType) {
		toSerialize["actionType"] = o.ActionType
	}
	if !IsNil(o.StartedDate) {
		toSerialize["startedDate"] = o.StartedDate
	}
	if !IsNil(o.FinishedDate) {
		toSerialize["finishedDate"] = o.FinishedDate
	}
	if !IsNil(o.IssueCounts) {
		toSerialize["issueCounts"] = o.IssueCounts
	}
	if !IsNil(o.ExecutionProgress) {
		toSerialize["executionProgress"] = o.ExecutionProgress
	}
	if !IsNil(o.CompletionStatus) {
		toSerialize["completionStatus"] = o.CompletionStatus
	}
	if !IsNil(o.IsRequiredToPass) {
		toSerialize["isRequiredToPass"] = o.IsRequiredToPass
	}
	return toSerialize, nil
}

type NullableCiBuildActionAttributes struct {
	value *CiBuildActionAttributes
	isSet bool
}

func (v NullableCiBuildActionAttributes) Get() *CiBuildActionAttributes {
	return v.value
}

func (v *NullableCiBuildActionAttributes) Set(val *CiBuildActionAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCiBuildActionAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCiBuildActionAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiBuildActionAttributes(val *CiBuildActionAttributes) *NullableCiBuildActionAttributes {
	return &NullableCiBuildActionAttributes{value: val, isSet: true}
}

func (v NullableCiBuildActionAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiBuildActionAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


