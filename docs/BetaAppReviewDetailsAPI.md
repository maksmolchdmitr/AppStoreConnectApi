# \BetaAppReviewDetailsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetaAppReviewDetailsAppGetToOneRelated**](BetaAppReviewDetailsAPI.md#BetaAppReviewDetailsAppGetToOneRelated) | **Get** /v1/betaAppReviewDetails/{id}/app | 
[**BetaAppReviewDetailsGetCollection**](BetaAppReviewDetailsAPI.md#BetaAppReviewDetailsGetCollection) | **Get** /v1/betaAppReviewDetails | 
[**BetaAppReviewDetailsGetInstance**](BetaAppReviewDetailsAPI.md#BetaAppReviewDetailsGetInstance) | **Get** /v1/betaAppReviewDetails/{id} | 
[**BetaAppReviewDetailsUpdateInstance**](BetaAppReviewDetailsAPI.md#BetaAppReviewDetailsUpdateInstance) | **Patch** /v1/betaAppReviewDetails/{id} | 



## BetaAppReviewDetailsAppGetToOneRelated

> AppWithoutIncludesResponse BetaAppReviewDetailsAppGetToOneRelated(ctx, id).FieldsApps(fieldsApps).Execute()



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
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BetaAppReviewDetailsAPI.BetaAppReviewDetailsAppGetToOneRelated(context.Background(), id).FieldsApps(fieldsApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppReviewDetailsAPI.BetaAppReviewDetailsAppGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaAppReviewDetailsAppGetToOneRelated`: AppWithoutIncludesResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaAppReviewDetailsAPI.BetaAppReviewDetailsAppGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaAppReviewDetailsAppGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 

### Return type

[**AppWithoutIncludesResponse**](AppWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaAppReviewDetailsGetCollection

> BetaAppReviewDetailsResponse BetaAppReviewDetailsGetCollection(ctx).FilterApp(filterApp).FieldsBetaAppReviewDetails(fieldsBetaAppReviewDetails).Limit(limit).Include(include).FieldsApps(fieldsApps).Execute()



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
    filterApp := []string{"Inner_example"} // []string | filter by id(s) of related 'app'
    fieldsBetaAppReviewDetails := []string{"FieldsBetaAppReviewDetails_example"} // []string | the fields to include for returned resources of type betaAppReviewDetails (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BetaAppReviewDetailsAPI.BetaAppReviewDetailsGetCollection(context.Background()).FilterApp(filterApp).FieldsBetaAppReviewDetails(fieldsBetaAppReviewDetails).Limit(limit).Include(include).FieldsApps(fieldsApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppReviewDetailsAPI.BetaAppReviewDetailsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaAppReviewDetailsGetCollection`: BetaAppReviewDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaAppReviewDetailsAPI.BetaAppReviewDetailsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetaAppReviewDetailsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterApp** | **[]string** | filter by id(s) of related &#39;app&#39; | 
 **fieldsBetaAppReviewDetails** | **[]string** | the fields to include for returned resources of type betaAppReviewDetails | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 

### Return type

[**BetaAppReviewDetailsResponse**](BetaAppReviewDetailsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaAppReviewDetailsGetInstance

> BetaAppReviewDetailResponse BetaAppReviewDetailsGetInstance(ctx, id).FieldsBetaAppReviewDetails(fieldsBetaAppReviewDetails).Include(include).FieldsApps(fieldsApps).Execute()



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
    fieldsBetaAppReviewDetails := []string{"FieldsBetaAppReviewDetails_example"} // []string | the fields to include for returned resources of type betaAppReviewDetails (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BetaAppReviewDetailsAPI.BetaAppReviewDetailsGetInstance(context.Background(), id).FieldsBetaAppReviewDetails(fieldsBetaAppReviewDetails).Include(include).FieldsApps(fieldsApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppReviewDetailsAPI.BetaAppReviewDetailsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaAppReviewDetailsGetInstance`: BetaAppReviewDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaAppReviewDetailsAPI.BetaAppReviewDetailsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaAppReviewDetailsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaAppReviewDetails** | **[]string** | the fields to include for returned resources of type betaAppReviewDetails | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 

### Return type

[**BetaAppReviewDetailResponse**](BetaAppReviewDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaAppReviewDetailsUpdateInstance

> BetaAppReviewDetailResponse BetaAppReviewDetailsUpdateInstance(ctx, id).BetaAppReviewDetailUpdateRequest(betaAppReviewDetailUpdateRequest).Execute()



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
    betaAppReviewDetailUpdateRequest := *openapiclient.NewBetaAppReviewDetailUpdateRequest(*openapiclient.NewBetaAppReviewDetailUpdateRequestData("Type_example", "Id_example")) // BetaAppReviewDetailUpdateRequest | BetaAppReviewDetail representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BetaAppReviewDetailsAPI.BetaAppReviewDetailsUpdateInstance(context.Background(), id).BetaAppReviewDetailUpdateRequest(betaAppReviewDetailUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppReviewDetailsAPI.BetaAppReviewDetailsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaAppReviewDetailsUpdateInstance`: BetaAppReviewDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaAppReviewDetailsAPI.BetaAppReviewDetailsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaAppReviewDetailsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaAppReviewDetailUpdateRequest** | [**BetaAppReviewDetailUpdateRequest**](BetaAppReviewDetailUpdateRequest.md) | BetaAppReviewDetail representation | 

### Return type

[**BetaAppReviewDetailResponse**](BetaAppReviewDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

