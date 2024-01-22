# GameCenterDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GameCenterDetail**](GameCenterDetail.md) |  | 
**Included** | Pointer to [**[]GameCenterDetailsResponseIncludedInner**](GameCenterDetailsResponseIncludedInner.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewGameCenterDetailResponse

`func NewGameCenterDetailResponse(data GameCenterDetail, links DocumentLinks, ) *GameCenterDetailResponse`

NewGameCenterDetailResponse instantiates a new GameCenterDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterDetailResponseWithDefaults

`func NewGameCenterDetailResponseWithDefaults() *GameCenterDetailResponse`

NewGameCenterDetailResponseWithDefaults instantiates a new GameCenterDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterDetailResponse) GetData() GameCenterDetail`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterDetailResponse) GetDataOk() (*GameCenterDetail, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterDetailResponse) SetData(v GameCenterDetail)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterDetailResponse) GetIncluded() []GameCenterDetailsResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterDetailResponse) GetIncludedOk() (*[]GameCenterDetailsResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterDetailResponse) SetIncluded(v []GameCenterDetailsResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterDetailResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterDetailResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterDetailResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterDetailResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


