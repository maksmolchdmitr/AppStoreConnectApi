# GameCenterLeaderboardSetImageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GameCenterLeaderboardSetImage**](GameCenterLeaderboardSetImage.md) |  | 
**Included** | Pointer to [**[]GameCenterLeaderboardSetLocalization**](GameCenterLeaderboardSetLocalization.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewGameCenterLeaderboardSetImageResponse

`func NewGameCenterLeaderboardSetImageResponse(data GameCenterLeaderboardSetImage, links DocumentLinks, ) *GameCenterLeaderboardSetImageResponse`

NewGameCenterLeaderboardSetImageResponse instantiates a new GameCenterLeaderboardSetImageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterLeaderboardSetImageResponseWithDefaults

`func NewGameCenterLeaderboardSetImageResponseWithDefaults() *GameCenterLeaderboardSetImageResponse`

NewGameCenterLeaderboardSetImageResponseWithDefaults instantiates a new GameCenterLeaderboardSetImageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterLeaderboardSetImageResponse) GetData() GameCenterLeaderboardSetImage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterLeaderboardSetImageResponse) GetDataOk() (*GameCenterLeaderboardSetImage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterLeaderboardSetImageResponse) SetData(v GameCenterLeaderboardSetImage)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterLeaderboardSetImageResponse) GetIncluded() []GameCenterLeaderboardSetLocalization`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterLeaderboardSetImageResponse) GetIncludedOk() (*[]GameCenterLeaderboardSetLocalization, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterLeaderboardSetImageResponse) SetIncluded(v []GameCenterLeaderboardSetLocalization)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterLeaderboardSetImageResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterLeaderboardSetImageResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterLeaderboardSetImageResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterLeaderboardSetImageResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


