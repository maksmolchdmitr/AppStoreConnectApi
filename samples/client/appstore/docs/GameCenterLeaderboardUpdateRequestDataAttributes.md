# GameCenterLeaderboardUpdateRequestDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultFormatter** | Pointer to [**GameCenterLeaderboardFormatter**](GameCenterLeaderboardFormatter.md) |  | [optional] 
**ReferenceName** | Pointer to **string** |  | [optional] 
**SubmissionType** | Pointer to **string** |  | [optional] 
**ScoreSortType** | Pointer to **string** |  | [optional] 
**ScoreRangeStart** | Pointer to **float64** |  | [optional] 
**ScoreRangeEnd** | Pointer to **float64** |  | [optional] 
**RecurrenceStartDate** | Pointer to **time.Time** |  | [optional] 
**RecurrenceDuration** | Pointer to **string** |  | [optional] 
**RecurrenceRule** | Pointer to **string** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 

## Methods

### NewGameCenterLeaderboardUpdateRequestDataAttributes

`func NewGameCenterLeaderboardUpdateRequestDataAttributes() *GameCenterLeaderboardUpdateRequestDataAttributes`

NewGameCenterLeaderboardUpdateRequestDataAttributes instantiates a new GameCenterLeaderboardUpdateRequestDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterLeaderboardUpdateRequestDataAttributesWithDefaults

`func NewGameCenterLeaderboardUpdateRequestDataAttributesWithDefaults() *GameCenterLeaderboardUpdateRequestDataAttributes`

NewGameCenterLeaderboardUpdateRequestDataAttributesWithDefaults instantiates a new GameCenterLeaderboardUpdateRequestDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultFormatter

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetDefaultFormatter() GameCenterLeaderboardFormatter`

GetDefaultFormatter returns the DefaultFormatter field if non-nil, zero value otherwise.

### GetDefaultFormatterOk

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetDefaultFormatterOk() (*GameCenterLeaderboardFormatter, bool)`

GetDefaultFormatterOk returns a tuple with the DefaultFormatter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFormatter

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) SetDefaultFormatter(v GameCenterLeaderboardFormatter)`

SetDefaultFormatter sets DefaultFormatter field to given value.

### HasDefaultFormatter

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) HasDefaultFormatter() bool`

HasDefaultFormatter returns a boolean if a field has been set.

### GetReferenceName

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetReferenceName() string`

GetReferenceName returns the ReferenceName field if non-nil, zero value otherwise.

### GetReferenceNameOk

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetReferenceNameOk() (*string, bool)`

GetReferenceNameOk returns a tuple with the ReferenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceName

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) SetReferenceName(v string)`

SetReferenceName sets ReferenceName field to given value.

### HasReferenceName

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) HasReferenceName() bool`

HasReferenceName returns a boolean if a field has been set.

### GetSubmissionType

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetSubmissionType() string`

GetSubmissionType returns the SubmissionType field if non-nil, zero value otherwise.

### GetSubmissionTypeOk

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetSubmissionTypeOk() (*string, bool)`

GetSubmissionTypeOk returns a tuple with the SubmissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionType

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) SetSubmissionType(v string)`

SetSubmissionType sets SubmissionType field to given value.

### HasSubmissionType

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) HasSubmissionType() bool`

HasSubmissionType returns a boolean if a field has been set.

### GetScoreSortType

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetScoreSortType() string`

GetScoreSortType returns the ScoreSortType field if non-nil, zero value otherwise.

### GetScoreSortTypeOk

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetScoreSortTypeOk() (*string, bool)`

GetScoreSortTypeOk returns a tuple with the ScoreSortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreSortType

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) SetScoreSortType(v string)`

SetScoreSortType sets ScoreSortType field to given value.

### HasScoreSortType

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) HasScoreSortType() bool`

HasScoreSortType returns a boolean if a field has been set.

### GetScoreRangeStart

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetScoreRangeStart() float64`

GetScoreRangeStart returns the ScoreRangeStart field if non-nil, zero value otherwise.

### GetScoreRangeStartOk

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetScoreRangeStartOk() (*float64, bool)`

GetScoreRangeStartOk returns a tuple with the ScoreRangeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreRangeStart

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) SetScoreRangeStart(v float64)`

SetScoreRangeStart sets ScoreRangeStart field to given value.

### HasScoreRangeStart

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) HasScoreRangeStart() bool`

HasScoreRangeStart returns a boolean if a field has been set.

### GetScoreRangeEnd

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetScoreRangeEnd() float64`

GetScoreRangeEnd returns the ScoreRangeEnd field if non-nil, zero value otherwise.

### GetScoreRangeEndOk

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetScoreRangeEndOk() (*float64, bool)`

GetScoreRangeEndOk returns a tuple with the ScoreRangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreRangeEnd

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) SetScoreRangeEnd(v float64)`

SetScoreRangeEnd sets ScoreRangeEnd field to given value.

### HasScoreRangeEnd

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) HasScoreRangeEnd() bool`

HasScoreRangeEnd returns a boolean if a field has been set.

### GetRecurrenceStartDate

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetRecurrenceStartDate() time.Time`

GetRecurrenceStartDate returns the RecurrenceStartDate field if non-nil, zero value otherwise.

### GetRecurrenceStartDateOk

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetRecurrenceStartDateOk() (*time.Time, bool)`

GetRecurrenceStartDateOk returns a tuple with the RecurrenceStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceStartDate

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) SetRecurrenceStartDate(v time.Time)`

SetRecurrenceStartDate sets RecurrenceStartDate field to given value.

### HasRecurrenceStartDate

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) HasRecurrenceStartDate() bool`

HasRecurrenceStartDate returns a boolean if a field has been set.

### GetRecurrenceDuration

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetRecurrenceDuration() string`

GetRecurrenceDuration returns the RecurrenceDuration field if non-nil, zero value otherwise.

### GetRecurrenceDurationOk

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetRecurrenceDurationOk() (*string, bool)`

GetRecurrenceDurationOk returns a tuple with the RecurrenceDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceDuration

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) SetRecurrenceDuration(v string)`

SetRecurrenceDuration sets RecurrenceDuration field to given value.

### HasRecurrenceDuration

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) HasRecurrenceDuration() bool`

HasRecurrenceDuration returns a boolean if a field has been set.

### GetRecurrenceRule

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetRecurrenceRule() string`

GetRecurrenceRule returns the RecurrenceRule field if non-nil, zero value otherwise.

### GetRecurrenceRuleOk

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetRecurrenceRuleOk() (*string, bool)`

GetRecurrenceRuleOk returns a tuple with the RecurrenceRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceRule

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) SetRecurrenceRule(v string)`

SetRecurrenceRule sets RecurrenceRule field to given value.

### HasRecurrenceRule

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) HasRecurrenceRule() bool`

HasRecurrenceRule returns a boolean if a field has been set.

### GetArchived

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *GameCenterLeaderboardUpdateRequestDataAttributes) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


