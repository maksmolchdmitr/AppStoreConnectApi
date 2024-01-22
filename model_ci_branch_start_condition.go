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

// checks if the CiBranchStartCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiBranchStartCondition{}

// CiBranchStartCondition struct for CiBranchStartCondition
type CiBranchStartCondition struct {
	Source              *CiBranchPatterns      `json:"source,omitempty"`
	FilesAndFoldersRule *CiFilesAndFoldersRule `json:"filesAndFoldersRule,omitempty"`
	AutoCancel          *bool                  `json:"autoCancel,omitempty"`
}

// NewCiBranchStartCondition instantiates a new CiBranchStartCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiBranchStartCondition() *CiBranchStartCondition {
	this := CiBranchStartCondition{}
	return &this
}

// NewCiBranchStartConditionWithDefaults instantiates a new CiBranchStartCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiBranchStartConditionWithDefaults() *CiBranchStartCondition {
	this := CiBranchStartCondition{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CiBranchStartCondition) GetSource() CiBranchPatterns {
	if o == nil || IsNil(o.Source) {
		var ret CiBranchPatterns
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBranchStartCondition) GetSourceOk() (*CiBranchPatterns, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CiBranchStartCondition) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given CiBranchPatterns and assigns it to the Source field.
func (o *CiBranchStartCondition) SetSource(v CiBranchPatterns) {
	o.Source = &v
}

// GetFilesAndFoldersRule returns the FilesAndFoldersRule field value if set, zero value otherwise.
func (o *CiBranchStartCondition) GetFilesAndFoldersRule() CiFilesAndFoldersRule {
	if o == nil || IsNil(o.FilesAndFoldersRule) {
		var ret CiFilesAndFoldersRule
		return ret
	}
	return *o.FilesAndFoldersRule
}

// GetFilesAndFoldersRuleOk returns a tuple with the FilesAndFoldersRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBranchStartCondition) GetFilesAndFoldersRuleOk() (*CiFilesAndFoldersRule, bool) {
	if o == nil || IsNil(o.FilesAndFoldersRule) {
		return nil, false
	}
	return o.FilesAndFoldersRule, true
}

// HasFilesAndFoldersRule returns a boolean if a field has been set.
func (o *CiBranchStartCondition) HasFilesAndFoldersRule() bool {
	if o != nil && !IsNil(o.FilesAndFoldersRule) {
		return true
	}

	return false
}

// SetFilesAndFoldersRule gets a reference to the given CiFilesAndFoldersRule and assigns it to the FilesAndFoldersRule field.
func (o *CiBranchStartCondition) SetFilesAndFoldersRule(v CiFilesAndFoldersRule) {
	o.FilesAndFoldersRule = &v
}

// GetAutoCancel returns the AutoCancel field value if set, zero value otherwise.
func (o *CiBranchStartCondition) GetAutoCancel() bool {
	if o == nil || IsNil(o.AutoCancel) {
		var ret bool
		return ret
	}
	return *o.AutoCancel
}

// GetAutoCancelOk returns a tuple with the AutoCancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBranchStartCondition) GetAutoCancelOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoCancel) {
		return nil, false
	}
	return o.AutoCancel, true
}

// HasAutoCancel returns a boolean if a field has been set.
func (o *CiBranchStartCondition) HasAutoCancel() bool {
	if o != nil && !IsNil(o.AutoCancel) {
		return true
	}

	return false
}

// SetAutoCancel gets a reference to the given bool and assigns it to the AutoCancel field.
func (o *CiBranchStartCondition) SetAutoCancel(v bool) {
	o.AutoCancel = &v
}

func (o CiBranchStartCondition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiBranchStartCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.FilesAndFoldersRule) {
		toSerialize["filesAndFoldersRule"] = o.FilesAndFoldersRule
	}
	if !IsNil(o.AutoCancel) {
		toSerialize["autoCancel"] = o.AutoCancel
	}
	return toSerialize, nil
}

type NullableCiBranchStartCondition struct {
	value *CiBranchStartCondition
	isSet bool
}

func (v NullableCiBranchStartCondition) Get() *CiBranchStartCondition {
	return v.value
}

func (v *NullableCiBranchStartCondition) Set(val *CiBranchStartCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableCiBranchStartCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableCiBranchStartCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiBranchStartCondition(val *CiBranchStartCondition) *NullableCiBranchStartCondition {
	return &NullableCiBranchStartCondition{value: val, isSet: true}
}

func (v NullableCiBranchStartCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiBranchStartCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
