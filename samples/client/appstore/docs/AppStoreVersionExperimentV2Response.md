# AppStoreVersionExperimentV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AppStoreVersionExperimentV2**](AppStoreVersionExperimentV2.md) |  | 
**Included** | Pointer to [**[]AppStoreVersionExperimentsV2ResponseIncludedInner**](AppStoreVersionExperimentsV2ResponseIncludedInner.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewAppStoreVersionExperimentV2Response

`func NewAppStoreVersionExperimentV2Response(data AppStoreVersionExperimentV2, links DocumentLinks, ) *AppStoreVersionExperimentV2Response`

NewAppStoreVersionExperimentV2Response instantiates a new AppStoreVersionExperimentV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStoreVersionExperimentV2ResponseWithDefaults

`func NewAppStoreVersionExperimentV2ResponseWithDefaults() *AppStoreVersionExperimentV2Response`

NewAppStoreVersionExperimentV2ResponseWithDefaults instantiates a new AppStoreVersionExperimentV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppStoreVersionExperimentV2Response) GetData() AppStoreVersionExperimentV2`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppStoreVersionExperimentV2Response) GetDataOk() (*AppStoreVersionExperimentV2, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppStoreVersionExperimentV2Response) SetData(v AppStoreVersionExperimentV2)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppStoreVersionExperimentV2Response) GetIncluded() []AppStoreVersionExperimentsV2ResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppStoreVersionExperimentV2Response) GetIncludedOk() (*[]AppStoreVersionExperimentsV2ResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppStoreVersionExperimentV2Response) SetIncluded(v []AppStoreVersionExperimentsV2ResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppStoreVersionExperimentV2Response) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AppStoreVersionExperimentV2Response) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppStoreVersionExperimentV2Response) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppStoreVersionExperimentV2Response) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


