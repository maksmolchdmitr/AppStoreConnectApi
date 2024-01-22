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

// checks if the AppEventVideoClipCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventVideoClipCreateRequestDataAttributes{}

// AppEventVideoClipCreateRequestDataAttributes struct for AppEventVideoClipCreateRequestDataAttributes
type AppEventVideoClipCreateRequestDataAttributes struct {
	FileSize int32 `json:"fileSize"`
	FileName string `json:"fileName"`
	PreviewFrameTimeCode *string `json:"previewFrameTimeCode,omitempty"`
	AppEventAssetType AppEventAssetType `json:"appEventAssetType"`
}

// NewAppEventVideoClipCreateRequestDataAttributes instantiates a new AppEventVideoClipCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventVideoClipCreateRequestDataAttributes(fileSize int32, fileName string, appEventAssetType AppEventAssetType) *AppEventVideoClipCreateRequestDataAttributes {
	this := AppEventVideoClipCreateRequestDataAttributes{}
	this.FileSize = fileSize
	this.FileName = fileName
	this.AppEventAssetType = appEventAssetType
	return &this
}

// NewAppEventVideoClipCreateRequestDataAttributesWithDefaults instantiates a new AppEventVideoClipCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventVideoClipCreateRequestDataAttributesWithDefaults() *AppEventVideoClipCreateRequestDataAttributes {
	this := AppEventVideoClipCreateRequestDataAttributes{}
	return &this
}

// GetFileSize returns the FileSize field value
func (o *AppEventVideoClipCreateRequestDataAttributes) GetFileSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value
// and a boolean to check if the value has been set.
func (o *AppEventVideoClipCreateRequestDataAttributes) GetFileSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileSize, true
}

// SetFileSize sets field value
func (o *AppEventVideoClipCreateRequestDataAttributes) SetFileSize(v int32) {
	o.FileSize = v
}

// GetFileName returns the FileName field value
func (o *AppEventVideoClipCreateRequestDataAttributes) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *AppEventVideoClipCreateRequestDataAttributes) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *AppEventVideoClipCreateRequestDataAttributes) SetFileName(v string) {
	o.FileName = v
}

// GetPreviewFrameTimeCode returns the PreviewFrameTimeCode field value if set, zero value otherwise.
func (o *AppEventVideoClipCreateRequestDataAttributes) GetPreviewFrameTimeCode() string {
	if o == nil || IsNil(o.PreviewFrameTimeCode) {
		var ret string
		return ret
	}
	return *o.PreviewFrameTimeCode
}

// GetPreviewFrameTimeCodeOk returns a tuple with the PreviewFrameTimeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventVideoClipCreateRequestDataAttributes) GetPreviewFrameTimeCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PreviewFrameTimeCode) {
		return nil, false
	}
	return o.PreviewFrameTimeCode, true
}

// HasPreviewFrameTimeCode returns a boolean if a field has been set.
func (o *AppEventVideoClipCreateRequestDataAttributes) HasPreviewFrameTimeCode() bool {
	if o != nil && !IsNil(o.PreviewFrameTimeCode) {
		return true
	}

	return false
}

// SetPreviewFrameTimeCode gets a reference to the given string and assigns it to the PreviewFrameTimeCode field.
func (o *AppEventVideoClipCreateRequestDataAttributes) SetPreviewFrameTimeCode(v string) {
	o.PreviewFrameTimeCode = &v
}

// GetAppEventAssetType returns the AppEventAssetType field value
func (o *AppEventVideoClipCreateRequestDataAttributes) GetAppEventAssetType() AppEventAssetType {
	if o == nil {
		var ret AppEventAssetType
		return ret
	}

	return o.AppEventAssetType
}

// GetAppEventAssetTypeOk returns a tuple with the AppEventAssetType field value
// and a boolean to check if the value has been set.
func (o *AppEventVideoClipCreateRequestDataAttributes) GetAppEventAssetTypeOk() (*AppEventAssetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppEventAssetType, true
}

// SetAppEventAssetType sets field value
func (o *AppEventVideoClipCreateRequestDataAttributes) SetAppEventAssetType(v AppEventAssetType) {
	o.AppEventAssetType = v
}

func (o AppEventVideoClipCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventVideoClipCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fileSize"] = o.FileSize
	toSerialize["fileName"] = o.FileName
	if !IsNil(o.PreviewFrameTimeCode) {
		toSerialize["previewFrameTimeCode"] = o.PreviewFrameTimeCode
	}
	toSerialize["appEventAssetType"] = o.AppEventAssetType
	return toSerialize, nil
}

type NullableAppEventVideoClipCreateRequestDataAttributes struct {
	value *AppEventVideoClipCreateRequestDataAttributes
	isSet bool
}

func (v NullableAppEventVideoClipCreateRequestDataAttributes) Get() *AppEventVideoClipCreateRequestDataAttributes {
	return v.value
}

func (v *NullableAppEventVideoClipCreateRequestDataAttributes) Set(val *AppEventVideoClipCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventVideoClipCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventVideoClipCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventVideoClipCreateRequestDataAttributes(val *AppEventVideoClipCreateRequestDataAttributes) *NullableAppEventVideoClipCreateRequestDataAttributes {
	return &NullableAppEventVideoClipCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppEventVideoClipCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventVideoClipCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


