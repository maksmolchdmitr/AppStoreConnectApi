# AppStoreVersionExperimentV2CreateRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Platform** | [**Platform**](Platform.md) |  | 
**TrafficProportion** | **int32** |  | 

## Methods

### NewAppStoreVersionExperimentV2CreateRequestDataAttributes

`func NewAppStoreVersionExperimentV2CreateRequestDataAttributes(name string, platform Platform, trafficProportion int32, ) *AppStoreVersionExperimentV2CreateRequestDataAttributes`

NewAppStoreVersionExperimentV2CreateRequestDataAttributes instantiates a new AppStoreVersionExperimentV2CreateRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStoreVersionExperimentV2CreateRequestDataAttributesWithDefaults

`func NewAppStoreVersionExperimentV2CreateRequestDataAttributesWithDefaults() *AppStoreVersionExperimentV2CreateRequestDataAttributes`

NewAppStoreVersionExperimentV2CreateRequestDataAttributesWithDefaults instantiates a new AppStoreVersionExperimentV2CreateRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AppStoreVersionExperimentV2CreateRequestDataAttributes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppStoreVersionExperimentV2CreateRequestDataAttributes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppStoreVersionExperimentV2CreateRequestDataAttributes) SetName(v string)`

SetName sets Name field to given value.


### GetPlatform

`func (o *AppStoreVersionExperimentV2CreateRequestDataAttributes) GetPlatform() Platform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AppStoreVersionExperimentV2CreateRequestDataAttributes) GetPlatformOk() (*Platform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AppStoreVersionExperimentV2CreateRequestDataAttributes) SetPlatform(v Platform)`

SetPlatform sets Platform field to given value.


### GetTrafficProportion

`func (o *AppStoreVersionExperimentV2CreateRequestDataAttributes) GetTrafficProportion() int32`

GetTrafficProportion returns the TrafficProportion field if non-nil, zero value otherwise.

### GetTrafficProportionOk

`func (o *AppStoreVersionExperimentV2CreateRequestDataAttributes) GetTrafficProportionOk() (*int32, bool)`

GetTrafficProportionOk returns a tuple with the TrafficProportion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficProportion

`func (o *AppStoreVersionExperimentV2CreateRequestDataAttributes) SetTrafficProportion(v int32)`

SetTrafficProportion sets TrafficProportion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


