# GameCenterLeaderboardCreateRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultFormatter** | [**GameCenterLeaderboardFormatter**](GameCenterLeaderboardFormatter.md) |  | 
**ReferenceName** | **string** |  | 
**VendorIdentifier** | **string** |  | 
**SubmissionType** | **string** |  | 
**ScoreSortType** | **string** |  | 
**ScoreRangeStart** | Pointer to **float64** |  | [optional] 
**ScoreRangeEnd** | Pointer to **float64** |  | [optional] 
**RecurrenceStartDate** | Pointer to **time.Time** |  | [optional] 
**RecurrenceDuration** | Pointer to **string** |  | [optional] 
**RecurrenceRule** | Pointer to **string** |  | [optional] 

## Methods

### NewGameCenterLeaderboardCreateRequestDataAttributes

`func NewGameCenterLeaderboardCreateRequestDataAttributes(defaultFormatter GameCenterLeaderboardFormatter, referenceName string, vendorIdentifier string, submissionType string, scoreSortType string, ) *GameCenterLeaderboardCreateRequestDataAttributes`

NewGameCenterLeaderboardCreateRequestDataAttributes instantiates a new GameCenterLeaderboardCreateRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterLeaderboardCreateRequestDataAttributesWithDefaults

`func NewGameCenterLeaderboardCreateRequestDataAttributesWithDefaults() *GameCenterLeaderboardCreateRequestDataAttributes`

NewGameCenterLeaderboardCreateRequestDataAttributesWithDefaults instantiates a new GameCenterLeaderboardCreateRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultFormatter

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetDefaultFormatter() GameCenterLeaderboardFormatter`

GetDefaultFormatter returns the DefaultFormatter field if non-nil, zero value otherwise.

### GetDefaultFormatterOk

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetDefaultFormatterOk() (*GameCenterLeaderboardFormatter, bool)`

GetDefaultFormatterOk returns a tuple with the DefaultFormatter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFormatter

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) SetDefaultFormatter(v GameCenterLeaderboardFormatter)`

SetDefaultFormatter sets DefaultFormatter field to given value.


### GetReferenceName

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetReferenceName() string`

GetReferenceName returns the ReferenceName field if non-nil, zero value otherwise.

### GetReferenceNameOk

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetReferenceNameOk() (*string, bool)`

GetReferenceNameOk returns a tuple with the ReferenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceName

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) SetReferenceName(v string)`

SetReferenceName sets ReferenceName field to given value.


### GetVendorIdentifier

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetVendorIdentifier() string`

GetVendorIdentifier returns the VendorIdentifier field if non-nil, zero value otherwise.

### GetVendorIdentifierOk

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetVendorIdentifierOk() (*string, bool)`

GetVendorIdentifierOk returns a tuple with the VendorIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorIdentifier

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) SetVendorIdentifier(v string)`

SetVendorIdentifier sets VendorIdentifier field to given value.


### GetSubmissionType

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetSubmissionType() string`

GetSubmissionType returns the SubmissionType field if non-nil, zero value otherwise.

### GetSubmissionTypeOk

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetSubmissionTypeOk() (*string, bool)`

GetSubmissionTypeOk returns a tuple with the SubmissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionType

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) SetSubmissionType(v string)`

SetSubmissionType sets SubmissionType field to given value.


### GetScoreSortType

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetScoreSortType() string`

GetScoreSortType returns the ScoreSortType field if non-nil, zero value otherwise.

### GetScoreSortTypeOk

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetScoreSortTypeOk() (*string, bool)`

GetScoreSortTypeOk returns a tuple with the ScoreSortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreSortType

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) SetScoreSortType(v string)`

SetScoreSortType sets ScoreSortType field to given value.


### GetScoreRangeStart

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetScoreRangeStart() float64`

GetScoreRangeStart returns the ScoreRangeStart field if non-nil, zero value otherwise.

### GetScoreRangeStartOk

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetScoreRangeStartOk() (*float64, bool)`

GetScoreRangeStartOk returns a tuple with the ScoreRangeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreRangeStart

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) SetScoreRangeStart(v float64)`

SetScoreRangeStart sets ScoreRangeStart field to given value.

### HasScoreRangeStart

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) HasScoreRangeStart() bool`

HasScoreRangeStart returns a boolean if a field has been set.

### GetScoreRangeEnd

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetScoreRangeEnd() float64`

GetScoreRangeEnd returns the ScoreRangeEnd field if non-nil, zero value otherwise.

### GetScoreRangeEndOk

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetScoreRangeEndOk() (*float64, bool)`

GetScoreRangeEndOk returns a tuple with the ScoreRangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreRangeEnd

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) SetScoreRangeEnd(v float64)`

SetScoreRangeEnd sets ScoreRangeEnd field to given value.

### HasScoreRangeEnd

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) HasScoreRangeEnd() bool`

HasScoreRangeEnd returns a boolean if a field has been set.

### GetRecurrenceStartDate

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetRecurrenceStartDate() time.Time`

GetRecurrenceStartDate returns the RecurrenceStartDate field if non-nil, zero value otherwise.

### GetRecurrenceStartDateOk

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetRecurrenceStartDateOk() (*time.Time, bool)`

GetRecurrenceStartDateOk returns a tuple with the RecurrenceStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceStartDate

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) SetRecurrenceStartDate(v time.Time)`

SetRecurrenceStartDate sets RecurrenceStartDate field to given value.

### HasRecurrenceStartDate

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) HasRecurrenceStartDate() bool`

HasRecurrenceStartDate returns a boolean if a field has been set.

### GetRecurrenceDuration

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetRecurrenceDuration() string`

GetRecurrenceDuration returns the RecurrenceDuration field if non-nil, zero value otherwise.

### GetRecurrenceDurationOk

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetRecurrenceDurationOk() (*string, bool)`

GetRecurrenceDurationOk returns a tuple with the RecurrenceDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceDuration

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) SetRecurrenceDuration(v string)`

SetRecurrenceDuration sets RecurrenceDuration field to given value.

### HasRecurrenceDuration

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) HasRecurrenceDuration() bool`

HasRecurrenceDuration returns a boolean if a field has been set.

### GetRecurrenceRule

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetRecurrenceRule() string`

GetRecurrenceRule returns the RecurrenceRule field if non-nil, zero value otherwise.

### GetRecurrenceRuleOk

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) GetRecurrenceRuleOk() (*string, bool)`

GetRecurrenceRuleOk returns a tuple with the RecurrenceRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceRule

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) SetRecurrenceRule(v string)`

SetRecurrenceRule sets RecurrenceRule field to given value.

### HasRecurrenceRule

`func (o *GameCenterLeaderboardCreateRequestDataAttributes) HasRecurrenceRule() bool`

HasRecurrenceRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


