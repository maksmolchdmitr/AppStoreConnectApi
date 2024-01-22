# \EndUserLicenseAgreementsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EndUserLicenseAgreementsCreateInstance**](EndUserLicenseAgreementsAPI.md#EndUserLicenseAgreementsCreateInstance) | **Post** /v1/endUserLicenseAgreements | 
[**EndUserLicenseAgreementsDeleteInstance**](EndUserLicenseAgreementsAPI.md#EndUserLicenseAgreementsDeleteInstance) | **Delete** /v1/endUserLicenseAgreements/{id} | 
[**EndUserLicenseAgreementsGetInstance**](EndUserLicenseAgreementsAPI.md#EndUserLicenseAgreementsGetInstance) | **Get** /v1/endUserLicenseAgreements/{id} | 
[**EndUserLicenseAgreementsTerritoriesGetToManyRelated**](EndUserLicenseAgreementsAPI.md#EndUserLicenseAgreementsTerritoriesGetToManyRelated) | **Get** /v1/endUserLicenseAgreements/{id}/territories | 
[**EndUserLicenseAgreementsUpdateInstance**](EndUserLicenseAgreementsAPI.md#EndUserLicenseAgreementsUpdateInstance) | **Patch** /v1/endUserLicenseAgreements/{id} | 



## EndUserLicenseAgreementsCreateInstance

> EndUserLicenseAgreementResponse EndUserLicenseAgreementsCreateInstance(ctx).EndUserLicenseAgreementCreateRequest(endUserLicenseAgreementCreateRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    endUserLicenseAgreementCreateRequest := *openapiclient.NewEndUserLicenseAgreementCreateRequest(*openapiclient.NewEndUserLicenseAgreementCreateRequestData("Type_example", *openapiclient.NewEndUserLicenseAgreementCreateRequestDataAttributes("AgreementText_example"), *openapiclient.NewEndUserLicenseAgreementCreateRequestDataRelationships(*openapiclient.NewAppAvailabilityV2CreateRequestDataRelationshipsApp(*openapiclient.NewAppAvailabilityV2CreateRequestDataRelationshipsAppData("Type_example", "Id_example")), *openapiclient.NewAppAvailabilityCreateRequestDataRelationshipsAvailableTerritories([]openapiclient.AppAvailabilityRelationshipsAvailableTerritoriesDataInner{*openapiclient.NewAppAvailabilityRelationshipsAvailableTerritoriesDataInner("Type_example", "Id_example")})))) // EndUserLicenseAgreementCreateRequest | EndUserLicenseAgreement representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EndUserLicenseAgreementsAPI.EndUserLicenseAgreementsCreateInstance(context.Background()).EndUserLicenseAgreementCreateRequest(endUserLicenseAgreementCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndUserLicenseAgreementsAPI.EndUserLicenseAgreementsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EndUserLicenseAgreementsCreateInstance`: EndUserLicenseAgreementResponse
    fmt.Fprintf(os.Stdout, "Response from `EndUserLicenseAgreementsAPI.EndUserLicenseAgreementsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEndUserLicenseAgreementsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endUserLicenseAgreementCreateRequest** | [**EndUserLicenseAgreementCreateRequest**](EndUserLicenseAgreementCreateRequest.md) | EndUserLicenseAgreement representation | 

### Return type

[**EndUserLicenseAgreementResponse**](EndUserLicenseAgreementResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndUserLicenseAgreementsDeleteInstance

> EndUserLicenseAgreementsDeleteInstance(ctx, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | the id of the requested resource

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EndUserLicenseAgreementsAPI.EndUserLicenseAgreementsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndUserLicenseAgreementsAPI.EndUserLicenseAgreementsDeleteInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndUserLicenseAgreementsDeleteInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndUserLicenseAgreementsGetInstance

> EndUserLicenseAgreementResponse EndUserLicenseAgreementsGetInstance(ctx, id).FieldsEndUserLicenseAgreements(fieldsEndUserLicenseAgreements).Include(include).FieldsTerritories(fieldsTerritories).LimitTerritories(limitTerritories).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | the id of the requested resource
    fieldsEndUserLicenseAgreements := []string{"FieldsEndUserLicenseAgreements_example"} // []string | the fields to include for returned resources of type endUserLicenseAgreements (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
    limitTerritories := int32(56) // int32 | maximum number of related territories returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EndUserLicenseAgreementsAPI.EndUserLicenseAgreementsGetInstance(context.Background(), id).FieldsEndUserLicenseAgreements(fieldsEndUserLicenseAgreements).Include(include).FieldsTerritories(fieldsTerritories).LimitTerritories(limitTerritories).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndUserLicenseAgreementsAPI.EndUserLicenseAgreementsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EndUserLicenseAgreementsGetInstance`: EndUserLicenseAgreementResponse
    fmt.Fprintf(os.Stdout, "Response from `EndUserLicenseAgreementsAPI.EndUserLicenseAgreementsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndUserLicenseAgreementsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsEndUserLicenseAgreements** | **[]string** | the fields to include for returned resources of type endUserLicenseAgreements | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 
 **limitTerritories** | **int32** | maximum number of related territories returned (when they are included) | 

### Return type

[**EndUserLicenseAgreementResponse**](EndUserLicenseAgreementResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndUserLicenseAgreementsTerritoriesGetToManyRelated

> TerritoriesWithoutIncludesResponse EndUserLicenseAgreementsTerritoriesGetToManyRelated(ctx, id).FieldsTerritories(fieldsTerritories).Limit(limit).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | the id of the requested resource
    fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EndUserLicenseAgreementsAPI.EndUserLicenseAgreementsTerritoriesGetToManyRelated(context.Background(), id).FieldsTerritories(fieldsTerritories).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndUserLicenseAgreementsAPI.EndUserLicenseAgreementsTerritoriesGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EndUserLicenseAgreementsTerritoriesGetToManyRelated`: TerritoriesWithoutIncludesResponse
    fmt.Fprintf(os.Stdout, "Response from `EndUserLicenseAgreementsAPI.EndUserLicenseAgreementsTerritoriesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndUserLicenseAgreementsTerritoriesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**TerritoriesWithoutIncludesResponse**](TerritoriesWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EndUserLicenseAgreementsUpdateInstance

> EndUserLicenseAgreementResponse EndUserLicenseAgreementsUpdateInstance(ctx, id).EndUserLicenseAgreementUpdateRequest(endUserLicenseAgreementUpdateRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    id := "id_example" // string | the id of the requested resource
    endUserLicenseAgreementUpdateRequest := *openapiclient.NewEndUserLicenseAgreementUpdateRequest(*openapiclient.NewEndUserLicenseAgreementUpdateRequestData("Type_example", "Id_example")) // EndUserLicenseAgreementUpdateRequest | EndUserLicenseAgreement representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EndUserLicenseAgreementsAPI.EndUserLicenseAgreementsUpdateInstance(context.Background(), id).EndUserLicenseAgreementUpdateRequest(endUserLicenseAgreementUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndUserLicenseAgreementsAPI.EndUserLicenseAgreementsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EndUserLicenseAgreementsUpdateInstance`: EndUserLicenseAgreementResponse
    fmt.Fprintf(os.Stdout, "Response from `EndUserLicenseAgreementsAPI.EndUserLicenseAgreementsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndUserLicenseAgreementsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endUserLicenseAgreementUpdateRequest** | [**EndUserLicenseAgreementUpdateRequest**](EndUserLicenseAgreementUpdateRequest.md) | EndUserLicenseAgreement representation | 

### Return type

[**EndUserLicenseAgreementResponse**](EndUserLicenseAgreementResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

