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

// checks if the CiBuildRunAttributesSourceCommit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiBuildRunAttributesSourceCommit{}

// CiBuildRunAttributesSourceCommit struct for CiBuildRunAttributesSourceCommit
type CiBuildRunAttributesSourceCommit struct {
	CommitSha *string    `json:"commitSha,omitempty"`
	Message   *string    `json:"message,omitempty"`
	Author    *CiGitUser `json:"author,omitempty"`
	Committer *CiGitUser `json:"committer,omitempty"`
	WebUrl    *string    `json:"webUrl,omitempty"`
}

// NewCiBuildRunAttributesSourceCommit instantiates a new CiBuildRunAttributesSourceCommit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiBuildRunAttributesSourceCommit() *CiBuildRunAttributesSourceCommit {
	this := CiBuildRunAttributesSourceCommit{}
	return &this
}

// NewCiBuildRunAttributesSourceCommitWithDefaults instantiates a new CiBuildRunAttributesSourceCommit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiBuildRunAttributesSourceCommitWithDefaults() *CiBuildRunAttributesSourceCommit {
	this := CiBuildRunAttributesSourceCommit{}
	return &this
}

// GetCommitSha returns the CommitSha field value if set, zero value otherwise.
func (o *CiBuildRunAttributesSourceCommit) GetCommitSha() string {
	if o == nil || IsNil(o.CommitSha) {
		var ret string
		return ret
	}
	return *o.CommitSha
}

// GetCommitShaOk returns a tuple with the CommitSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildRunAttributesSourceCommit) GetCommitShaOk() (*string, bool) {
	if o == nil || IsNil(o.CommitSha) {
		return nil, false
	}
	return o.CommitSha, true
}

// HasCommitSha returns a boolean if a field has been set.
func (o *CiBuildRunAttributesSourceCommit) HasCommitSha() bool {
	if o != nil && !IsNil(o.CommitSha) {
		return true
	}

	return false
}

// SetCommitSha gets a reference to the given string and assigns it to the CommitSha field.
func (o *CiBuildRunAttributesSourceCommit) SetCommitSha(v string) {
	o.CommitSha = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CiBuildRunAttributesSourceCommit) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildRunAttributesSourceCommit) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CiBuildRunAttributesSourceCommit) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CiBuildRunAttributesSourceCommit) SetMessage(v string) {
	o.Message = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *CiBuildRunAttributesSourceCommit) GetAuthor() CiGitUser {
	if o == nil || IsNil(o.Author) {
		var ret CiGitUser
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildRunAttributesSourceCommit) GetAuthorOk() (*CiGitUser, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *CiBuildRunAttributesSourceCommit) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given CiGitUser and assigns it to the Author field.
func (o *CiBuildRunAttributesSourceCommit) SetAuthor(v CiGitUser) {
	o.Author = &v
}

// GetCommitter returns the Committer field value if set, zero value otherwise.
func (o *CiBuildRunAttributesSourceCommit) GetCommitter() CiGitUser {
	if o == nil || IsNil(o.Committer) {
		var ret CiGitUser
		return ret
	}
	return *o.Committer
}

// GetCommitterOk returns a tuple with the Committer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildRunAttributesSourceCommit) GetCommitterOk() (*CiGitUser, bool) {
	if o == nil || IsNil(o.Committer) {
		return nil, false
	}
	return o.Committer, true
}

// HasCommitter returns a boolean if a field has been set.
func (o *CiBuildRunAttributesSourceCommit) HasCommitter() bool {
	if o != nil && !IsNil(o.Committer) {
		return true
	}

	return false
}

// SetCommitter gets a reference to the given CiGitUser and assigns it to the Committer field.
func (o *CiBuildRunAttributesSourceCommit) SetCommitter(v CiGitUser) {
	o.Committer = &v
}

// GetWebUrl returns the WebUrl field value if set, zero value otherwise.
func (o *CiBuildRunAttributesSourceCommit) GetWebUrl() string {
	if o == nil || IsNil(o.WebUrl) {
		var ret string
		return ret
	}
	return *o.WebUrl
}

// GetWebUrlOk returns a tuple with the WebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildRunAttributesSourceCommit) GetWebUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebUrl) {
		return nil, false
	}
	return o.WebUrl, true
}

// HasWebUrl returns a boolean if a field has been set.
func (o *CiBuildRunAttributesSourceCommit) HasWebUrl() bool {
	if o != nil && !IsNil(o.WebUrl) {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.
func (o *CiBuildRunAttributesSourceCommit) SetWebUrl(v string) {
	o.WebUrl = &v
}

func (o CiBuildRunAttributesSourceCommit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiBuildRunAttributesSourceCommit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommitSha) {
		toSerialize["commitSha"] = o.CommitSha
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Committer) {
		toSerialize["committer"] = o.Committer
	}
	if !IsNil(o.WebUrl) {
		toSerialize["webUrl"] = o.WebUrl
	}
	return toSerialize, nil
}

type NullableCiBuildRunAttributesSourceCommit struct {
	value *CiBuildRunAttributesSourceCommit
	isSet bool
}

func (v NullableCiBuildRunAttributesSourceCommit) Get() *CiBuildRunAttributesSourceCommit {
	return v.value
}

func (v *NullableCiBuildRunAttributesSourceCommit) Set(val *CiBuildRunAttributesSourceCommit) {
	v.value = val
	v.isSet = true
}

func (v NullableCiBuildRunAttributesSourceCommit) IsSet() bool {
	return v.isSet
}

func (v *NullableCiBuildRunAttributesSourceCommit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiBuildRunAttributesSourceCommit(val *CiBuildRunAttributesSourceCommit) *NullableCiBuildRunAttributesSourceCommit {
	return &NullableCiBuildRunAttributesSourceCommit{value: val, isSet: true}
}

func (v NullableCiBuildRunAttributesSourceCommit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiBuildRunAttributesSourceCommit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
