# GameCenterAppVersionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GameCenterAppVersion**](GameCenterAppVersion.md) |  | 
**Included** | Pointer to [**[]GameCenterAppVersionsResponseIncludedInner**](GameCenterAppVersionsResponseIncludedInner.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewGameCenterAppVersionsResponse

`func NewGameCenterAppVersionsResponse(data []GameCenterAppVersion, links PagedDocumentLinks, ) *GameCenterAppVersionsResponse`

NewGameCenterAppVersionsResponse instantiates a new GameCenterAppVersionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterAppVersionsResponseWithDefaults

`func NewGameCenterAppVersionsResponseWithDefaults() *GameCenterAppVersionsResponse`

NewGameCenterAppVersionsResponseWithDefaults instantiates a new GameCenterAppVersionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterAppVersionsResponse) GetData() []GameCenterAppVersion`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterAppVersionsResponse) GetDataOk() (*[]GameCenterAppVersion, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterAppVersionsResponse) SetData(v []GameCenterAppVersion)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterAppVersionsResponse) GetIncluded() []GameCenterAppVersionsResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterAppVersionsResponse) GetIncludedOk() (*[]GameCenterAppVersionsResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterAppVersionsResponse) SetIncluded(v []GameCenterAppVersionsResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterAppVersionsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterAppVersionsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterAppVersionsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterAppVersionsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *GameCenterAppVersionsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GameCenterAppVersionsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GameCenterAppVersionsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GameCenterAppVersionsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


