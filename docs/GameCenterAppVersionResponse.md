# GameCenterAppVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GameCenterAppVersion**](GameCenterAppVersion.md) |  | 
**Included** | Pointer to [**[]GameCenterAppVersionsResponseIncludedInner**](GameCenterAppVersionsResponseIncludedInner.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewGameCenterAppVersionResponse

`func NewGameCenterAppVersionResponse(data GameCenterAppVersion, links DocumentLinks, ) *GameCenterAppVersionResponse`

NewGameCenterAppVersionResponse instantiates a new GameCenterAppVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterAppVersionResponseWithDefaults

`func NewGameCenterAppVersionResponseWithDefaults() *GameCenterAppVersionResponse`

NewGameCenterAppVersionResponseWithDefaults instantiates a new GameCenterAppVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterAppVersionResponse) GetData() GameCenterAppVersion`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterAppVersionResponse) GetDataOk() (*GameCenterAppVersion, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterAppVersionResponse) SetData(v GameCenterAppVersion)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterAppVersionResponse) GetIncluded() []GameCenterAppVersionsResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterAppVersionResponse) GetIncludedOk() (*[]GameCenterAppVersionsResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterAppVersionResponse) SetIncluded(v []GameCenterAppVersionsResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterAppVersionResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterAppVersionResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterAppVersionResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterAppVersionResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


