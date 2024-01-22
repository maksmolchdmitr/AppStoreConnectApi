# CertificatesWithoutIncludesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Profile**](Profile.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewCertificatesWithoutIncludesResponse

`func NewCertificatesWithoutIncludesResponse(data []Profile, links PagedDocumentLinks, ) *CertificatesWithoutIncludesResponse`

NewCertificatesWithoutIncludesResponse instantiates a new CertificatesWithoutIncludesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatesWithoutIncludesResponseWithDefaults

`func NewCertificatesWithoutIncludesResponseWithDefaults() *CertificatesWithoutIncludesResponse`

NewCertificatesWithoutIncludesResponseWithDefaults instantiates a new CertificatesWithoutIncludesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CertificatesWithoutIncludesResponse) GetData() []Profile`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CertificatesWithoutIncludesResponse) GetDataOk() (*[]Profile, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CertificatesWithoutIncludesResponse) SetData(v []Profile)`

SetData sets Data field to given value.


### GetLinks

`func (o *CertificatesWithoutIncludesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CertificatesWithoutIncludesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CertificatesWithoutIncludesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *CertificatesWithoutIncludesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CertificatesWithoutIncludesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CertificatesWithoutIncludesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CertificatesWithoutIncludesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


