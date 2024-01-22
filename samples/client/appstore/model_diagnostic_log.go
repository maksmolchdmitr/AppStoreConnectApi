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

// checks if the DiagnosticLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiagnosticLog{}

// DiagnosticLog struct for DiagnosticLog
type DiagnosticLog struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Links *ResourceLinks `json:"links,omitempty"`
}

// NewDiagnosticLog instantiates a new DiagnosticLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiagnosticLog(type_ string, id string) *DiagnosticLog {
	this := DiagnosticLog{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewDiagnosticLogWithDefaults instantiates a new DiagnosticLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiagnosticLogWithDefaults() *DiagnosticLog {
	this := DiagnosticLog{}
	return &this
}

// GetType returns the Type field value
func (o *DiagnosticLog) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DiagnosticLog) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DiagnosticLog) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *DiagnosticLog) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DiagnosticLog) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DiagnosticLog) SetId(v string) {
	o.Id = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DiagnosticLog) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticLog) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DiagnosticLog) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *DiagnosticLog) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o DiagnosticLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiagnosticLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableDiagnosticLog struct {
	value *DiagnosticLog
	isSet bool
}

func (v NullableDiagnosticLog) Get() *DiagnosticLog {
	return v.value
}

func (v *NullableDiagnosticLog) Set(val *DiagnosticLog) {
	v.value = val
	v.isSet = true
}

func (v NullableDiagnosticLog) IsSet() bool {
	return v.isSet
}

func (v *NullableDiagnosticLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiagnosticLog(val *DiagnosticLog) *NullableDiagnosticLog {
	return &NullableDiagnosticLog{value: val, isSet: true}
}

func (v NullableDiagnosticLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiagnosticLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


