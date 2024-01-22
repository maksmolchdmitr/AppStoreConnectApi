# \BetaAppLocalizationsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetaAppLocalizationsAppGetToOneRelated**](BetaAppLocalizationsAPI.md#BetaAppLocalizationsAppGetToOneRelated) | **Get** /v1/betaAppLocalizations/{id}/app | 
[**BetaAppLocalizationsCreateInstance**](BetaAppLocalizationsAPI.md#BetaAppLocalizationsCreateInstance) | **Post** /v1/betaAppLocalizations | 
[**BetaAppLocalizationsDeleteInstance**](BetaAppLocalizationsAPI.md#BetaAppLocalizationsDeleteInstance) | **Delete** /v1/betaAppLocalizations/{id} | 
[**BetaAppLocalizationsGetCollection**](BetaAppLocalizationsAPI.md#BetaAppLocalizationsGetCollection) | **Get** /v1/betaAppLocalizations | 
[**BetaAppLocalizationsGetInstance**](BetaAppLocalizationsAPI.md#BetaAppLocalizationsGetInstance) | **Get** /v1/betaAppLocalizations/{id} | 
[**BetaAppLocalizationsUpdateInstance**](BetaAppLocalizationsAPI.md#BetaAppLocalizationsUpdateInstance) | **Patch** /v1/betaAppLocalizations/{id} | 



## BetaAppLocalizationsAppGetToOneRelated

> AppWithoutIncludesResponse BetaAppLocalizationsAppGetToOneRelated(ctx, id).FieldsApps(fieldsApps).Execute()



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
    resp, r, err := apiClient.BetaAppLocalizationsAPI.BetaAppLocalizationsAppGetToOneRelated(context.Background(), id).FieldsApps(fieldsApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppLocalizationsAPI.BetaAppLocalizationsAppGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaAppLocalizationsAppGetToOneRelated`: AppWithoutIncludesResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaAppLocalizationsAPI.BetaAppLocalizationsAppGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaAppLocalizationsAppGetToOneRelatedRequest struct via the builder pattern


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


## BetaAppLocalizationsCreateInstance

> BetaAppLocalizationResponse BetaAppLocalizationsCreateInstance(ctx).BetaAppLocalizationCreateRequest(betaAppLocalizationCreateRequest).Execute()



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
    betaAppLocalizationCreateRequest := *openapiclient.NewBetaAppLocalizationCreateRequest(*openapiclient.NewBetaAppLocalizationCreateRequestData("Type_example", *openapiclient.NewBetaAppLocalizationCreateRequestDataAttributes("Locale_example"), *openapiclient.NewAppEventCreateRequestDataRelationships(*openapiclient.NewAppAvailabilityV2CreateRequestDataRelationshipsApp(*openapiclient.NewAppAvailabilityV2CreateRequestDataRelationshipsAppData("Type_example", "Id_example"))))) // BetaAppLocalizationCreateRequest | BetaAppLocalization representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BetaAppLocalizationsAPI.BetaAppLocalizationsCreateInstance(context.Background()).BetaAppLocalizationCreateRequest(betaAppLocalizationCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppLocalizationsAPI.BetaAppLocalizationsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaAppLocalizationsCreateInstance`: BetaAppLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaAppLocalizationsAPI.BetaAppLocalizationsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetaAppLocalizationsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **betaAppLocalizationCreateRequest** | [**BetaAppLocalizationCreateRequest**](BetaAppLocalizationCreateRequest.md) | BetaAppLocalization representation | 

### Return type

[**BetaAppLocalizationResponse**](BetaAppLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaAppLocalizationsDeleteInstance

> BetaAppLocalizationsDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.BetaAppLocalizationsAPI.BetaAppLocalizationsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppLocalizationsAPI.BetaAppLocalizationsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaAppLocalizationsDeleteInstanceRequest struct via the builder pattern


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


## BetaAppLocalizationsGetCollection

> BetaAppLocalizationsResponse BetaAppLocalizationsGetCollection(ctx).FilterLocale(filterLocale).FilterApp(filterApp).FieldsBetaAppLocalizations(fieldsBetaAppLocalizations).Limit(limit).Include(include).FieldsApps(fieldsApps).Execute()



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
    filterLocale := []string{"Inner_example"} // []string | filter by attribute 'locale' (optional)
    filterApp := []string{"Inner_example"} // []string | filter by id(s) of related 'app' (optional)
    fieldsBetaAppLocalizations := []string{"FieldsBetaAppLocalizations_example"} // []string | the fields to include for returned resources of type betaAppLocalizations (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BetaAppLocalizationsAPI.BetaAppLocalizationsGetCollection(context.Background()).FilterLocale(filterLocale).FilterApp(filterApp).FieldsBetaAppLocalizations(fieldsBetaAppLocalizations).Limit(limit).Include(include).FieldsApps(fieldsApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppLocalizationsAPI.BetaAppLocalizationsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaAppLocalizationsGetCollection`: BetaAppLocalizationsResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaAppLocalizationsAPI.BetaAppLocalizationsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetaAppLocalizationsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterLocale** | **[]string** | filter by attribute &#39;locale&#39; | 
 **filterApp** | **[]string** | filter by id(s) of related &#39;app&#39; | 
 **fieldsBetaAppLocalizations** | **[]string** | the fields to include for returned resources of type betaAppLocalizations | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 

### Return type

[**BetaAppLocalizationsResponse**](BetaAppLocalizationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaAppLocalizationsGetInstance

> BetaAppLocalizationResponse BetaAppLocalizationsGetInstance(ctx, id).FieldsBetaAppLocalizations(fieldsBetaAppLocalizations).Include(include).FieldsApps(fieldsApps).Execute()



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
    fieldsBetaAppLocalizations := []string{"FieldsBetaAppLocalizations_example"} // []string | the fields to include for returned resources of type betaAppLocalizations (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BetaAppLocalizationsAPI.BetaAppLocalizationsGetInstance(context.Background(), id).FieldsBetaAppLocalizations(fieldsBetaAppLocalizations).Include(include).FieldsApps(fieldsApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppLocalizationsAPI.BetaAppLocalizationsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaAppLocalizationsGetInstance`: BetaAppLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaAppLocalizationsAPI.BetaAppLocalizationsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaAppLocalizationsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaAppLocalizations** | **[]string** | the fields to include for returned resources of type betaAppLocalizations | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 

### Return type

[**BetaAppLocalizationResponse**](BetaAppLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaAppLocalizationsUpdateInstance

> BetaAppLocalizationResponse BetaAppLocalizationsUpdateInstance(ctx, id).BetaAppLocalizationUpdateRequest(betaAppLocalizationUpdateRequest).Execute()



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
    betaAppLocalizationUpdateRequest := *openapiclient.NewBetaAppLocalizationUpdateRequest(*openapiclient.NewBetaAppLocalizationUpdateRequestData("Type_example", "Id_example")) // BetaAppLocalizationUpdateRequest | BetaAppLocalization representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BetaAppLocalizationsAPI.BetaAppLocalizationsUpdateInstance(context.Background(), id).BetaAppLocalizationUpdateRequest(betaAppLocalizationUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppLocalizationsAPI.BetaAppLocalizationsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaAppLocalizationsUpdateInstance`: BetaAppLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaAppLocalizationsAPI.BetaAppLocalizationsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaAppLocalizationsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaAppLocalizationUpdateRequest** | [**BetaAppLocalizationUpdateRequest**](BetaAppLocalizationUpdateRequest.md) | BetaAppLocalization representation | 

### Return type

[**BetaAppLocalizationResponse**](BetaAppLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

