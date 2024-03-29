# RoutingAppCoverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AppStoreReviewAttachmentAttributes**](AppStoreReviewAttachmentAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**AppStoreVersionSubmissionRelationships**](AppStoreVersionSubmissionRelationships.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewRoutingAppCoverage

`func NewRoutingAppCoverage(type_ string, id string, ) *RoutingAppCoverage`

NewRoutingAppCoverage instantiates a new RoutingAppCoverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingAppCoverageWithDefaults

`func NewRoutingAppCoverageWithDefaults() *RoutingAppCoverage`

NewRoutingAppCoverageWithDefaults instantiates a new RoutingAppCoverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RoutingAppCoverage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutingAppCoverage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutingAppCoverage) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *RoutingAppCoverage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoutingAppCoverage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoutingAppCoverage) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *RoutingAppCoverage) GetAttributes() AppStoreReviewAttachmentAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *RoutingAppCoverage) GetAttributesOk() (*AppStoreReviewAttachmentAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *RoutingAppCoverage) SetAttributes(v AppStoreReviewAttachmentAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *RoutingAppCoverage) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *RoutingAppCoverage) GetRelationships() AppStoreVersionSubmissionRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *RoutingAppCoverage) GetRelationshipsOk() (*AppStoreVersionSubmissionRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *RoutingAppCoverage) SetRelationships(v AppStoreVersionSubmissionRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *RoutingAppCoverage) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *RoutingAppCoverage) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RoutingAppCoverage) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RoutingAppCoverage) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RoutingAppCoverage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


