# AppStoreVersionExperimentV2CreateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**AppStoreVersionExperimentV2CreateRequestDataAttributes**](AppStoreVersionExperimentV2CreateRequestDataAttributes.md) |  | 
**Relationships** | [**AppEventCreateRequestDataRelationships**](AppEventCreateRequestDataRelationships.md) |  | 

## Methods

### NewAppStoreVersionExperimentV2CreateRequestData

`func NewAppStoreVersionExperimentV2CreateRequestData(type_ string, attributes AppStoreVersionExperimentV2CreateRequestDataAttributes, relationships AppEventCreateRequestDataRelationships, ) *AppStoreVersionExperimentV2CreateRequestData`

NewAppStoreVersionExperimentV2CreateRequestData instantiates a new AppStoreVersionExperimentV2CreateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStoreVersionExperimentV2CreateRequestDataWithDefaults

`func NewAppStoreVersionExperimentV2CreateRequestDataWithDefaults() *AppStoreVersionExperimentV2CreateRequestData`

NewAppStoreVersionExperimentV2CreateRequestDataWithDefaults instantiates a new AppStoreVersionExperimentV2CreateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppStoreVersionExperimentV2CreateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppStoreVersionExperimentV2CreateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppStoreVersionExperimentV2CreateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *AppStoreVersionExperimentV2CreateRequestData) GetAttributes() AppStoreVersionExperimentV2CreateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppStoreVersionExperimentV2CreateRequestData) GetAttributesOk() (*AppStoreVersionExperimentV2CreateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppStoreVersionExperimentV2CreateRequestData) SetAttributes(v AppStoreVersionExperimentV2CreateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *AppStoreVersionExperimentV2CreateRequestData) GetRelationships() AppEventCreateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AppStoreVersionExperimentV2CreateRequestData) GetRelationshipsOk() (*AppEventCreateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AppStoreVersionExperimentV2CreateRequestData) SetRelationships(v AppEventCreateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


