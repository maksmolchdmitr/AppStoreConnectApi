# GameCenterAppVersionCompatibilityVersionsLinkagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner**](GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewGameCenterAppVersionCompatibilityVersionsLinkagesResponse

`func NewGameCenterAppVersionCompatibilityVersionsLinkagesResponse(data []GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner, links PagedDocumentLinks, ) *GameCenterAppVersionCompatibilityVersionsLinkagesResponse`

NewGameCenterAppVersionCompatibilityVersionsLinkagesResponse instantiates a new GameCenterAppVersionCompatibilityVersionsLinkagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterAppVersionCompatibilityVersionsLinkagesResponseWithDefaults

`func NewGameCenterAppVersionCompatibilityVersionsLinkagesResponseWithDefaults() *GameCenterAppVersionCompatibilityVersionsLinkagesResponse`

NewGameCenterAppVersionCompatibilityVersionsLinkagesResponseWithDefaults instantiates a new GameCenterAppVersionCompatibilityVersionsLinkagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterAppVersionCompatibilityVersionsLinkagesResponse) GetData() []GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterAppVersionCompatibilityVersionsLinkagesResponse) GetDataOk() (*[]GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterAppVersionCompatibilityVersionsLinkagesResponse) SetData(v []GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner)`

SetData sets Data field to given value.


### GetLinks

`func (o *GameCenterAppVersionCompatibilityVersionsLinkagesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterAppVersionCompatibilityVersionsLinkagesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterAppVersionCompatibilityVersionsLinkagesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *GameCenterAppVersionCompatibilityVersionsLinkagesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GameCenterAppVersionCompatibilityVersionsLinkagesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GameCenterAppVersionCompatibilityVersionsLinkagesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GameCenterAppVersionCompatibilityVersionsLinkagesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


