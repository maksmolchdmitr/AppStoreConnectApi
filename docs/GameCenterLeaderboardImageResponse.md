# GameCenterLeaderboardImageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GameCenterLeaderboardImage**](GameCenterLeaderboardImage.md) |  | 
**Included** | Pointer to [**[]GameCenterLeaderboardLocalization**](GameCenterLeaderboardLocalization.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewGameCenterLeaderboardImageResponse

`func NewGameCenterLeaderboardImageResponse(data GameCenterLeaderboardImage, links DocumentLinks, ) *GameCenterLeaderboardImageResponse`

NewGameCenterLeaderboardImageResponse instantiates a new GameCenterLeaderboardImageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterLeaderboardImageResponseWithDefaults

`func NewGameCenterLeaderboardImageResponseWithDefaults() *GameCenterLeaderboardImageResponse`

NewGameCenterLeaderboardImageResponseWithDefaults instantiates a new GameCenterLeaderboardImageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterLeaderboardImageResponse) GetData() GameCenterLeaderboardImage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterLeaderboardImageResponse) GetDataOk() (*GameCenterLeaderboardImage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterLeaderboardImageResponse) SetData(v GameCenterLeaderboardImage)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterLeaderboardImageResponse) GetIncluded() []GameCenterLeaderboardLocalization`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterLeaderboardImageResponse) GetIncludedOk() (*[]GameCenterLeaderboardLocalization, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterLeaderboardImageResponse) SetIncluded(v []GameCenterLeaderboardLocalization)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterLeaderboardImageResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterLeaderboardImageResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterLeaderboardImageResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterLeaderboardImageResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


