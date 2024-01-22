# \SubscriptionLocalizationsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionLocalizationsCreateInstance**](SubscriptionLocalizationsAPI.md#SubscriptionLocalizationsCreateInstance) | **Post** /v1/subscriptionLocalizations | 
[**SubscriptionLocalizationsDeleteInstance**](SubscriptionLocalizationsAPI.md#SubscriptionLocalizationsDeleteInstance) | **Delete** /v1/subscriptionLocalizations/{id} | 
[**SubscriptionLocalizationsGetInstance**](SubscriptionLocalizationsAPI.md#SubscriptionLocalizationsGetInstance) | **Get** /v1/subscriptionLocalizations/{id} | 
[**SubscriptionLocalizationsUpdateInstance**](SubscriptionLocalizationsAPI.md#SubscriptionLocalizationsUpdateInstance) | **Patch** /v1/subscriptionLocalizations/{id} | 



## SubscriptionLocalizationsCreateInstance

> SubscriptionLocalizationResponse SubscriptionLocalizationsCreateInstance(ctx).SubscriptionLocalizationCreateRequest(subscriptionLocalizationCreateRequest).Execute()



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
    subscriptionLocalizationCreateRequest := *openapiclient.NewSubscriptionLocalizationCreateRequest(*openapiclient.NewSubscriptionLocalizationCreateRequestData("Type_example", *openapiclient.NewInAppPurchaseLocalizationCreateRequestDataAttributes("Name_example", "Locale_example"), *openapiclient.NewSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationships(*openapiclient.NewSubscriptionAppStoreReviewScreenshotCreateRequestDataRelationshipsSubscription(*openapiclient.NewPromotedPurchaseRelationshipsSubscriptionData("Type_example", "Id_example"))))) // SubscriptionLocalizationCreateRequest | SubscriptionLocalization representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionLocalizationsAPI.SubscriptionLocalizationsCreateInstance(context.Background()).SubscriptionLocalizationCreateRequest(subscriptionLocalizationCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionLocalizationsAPI.SubscriptionLocalizationsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionLocalizationsCreateInstance`: SubscriptionLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionLocalizationsAPI.SubscriptionLocalizationsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionLocalizationsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionLocalizationCreateRequest** | [**SubscriptionLocalizationCreateRequest**](SubscriptionLocalizationCreateRequest.md) | SubscriptionLocalization representation | 

### Return type

[**SubscriptionLocalizationResponse**](SubscriptionLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionLocalizationsDeleteInstance

> SubscriptionLocalizationsDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.SubscriptionLocalizationsAPI.SubscriptionLocalizationsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionLocalizationsAPI.SubscriptionLocalizationsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSubscriptionLocalizationsDeleteInstanceRequest struct via the builder pattern


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


## SubscriptionLocalizationsGetInstance

> SubscriptionLocalizationResponse SubscriptionLocalizationsGetInstance(ctx, id).FieldsSubscriptionLocalizations(fieldsSubscriptionLocalizations).Include(include).Execute()



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
    fieldsSubscriptionLocalizations := []string{"FieldsSubscriptionLocalizations_example"} // []string | the fields to include for returned resources of type subscriptionLocalizations (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionLocalizationsAPI.SubscriptionLocalizationsGetInstance(context.Background(), id).FieldsSubscriptionLocalizations(fieldsSubscriptionLocalizations).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionLocalizationsAPI.SubscriptionLocalizationsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionLocalizationsGetInstance`: SubscriptionLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionLocalizationsAPI.SubscriptionLocalizationsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionLocalizationsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsSubscriptionLocalizations** | **[]string** | the fields to include for returned resources of type subscriptionLocalizations | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**SubscriptionLocalizationResponse**](SubscriptionLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionLocalizationsUpdateInstance

> SubscriptionLocalizationResponse SubscriptionLocalizationsUpdateInstance(ctx, id).SubscriptionLocalizationUpdateRequest(subscriptionLocalizationUpdateRequest).Execute()



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
    subscriptionLocalizationUpdateRequest := *openapiclient.NewSubscriptionLocalizationUpdateRequest(*openapiclient.NewSubscriptionLocalizationUpdateRequestData("Type_example", "Id_example")) // SubscriptionLocalizationUpdateRequest | SubscriptionLocalization representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionLocalizationsAPI.SubscriptionLocalizationsUpdateInstance(context.Background(), id).SubscriptionLocalizationUpdateRequest(subscriptionLocalizationUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionLocalizationsAPI.SubscriptionLocalizationsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionLocalizationsUpdateInstance`: SubscriptionLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionLocalizationsAPI.SubscriptionLocalizationsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionLocalizationsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionLocalizationUpdateRequest** | [**SubscriptionLocalizationUpdateRequest**](SubscriptionLocalizationUpdateRequest.md) | SubscriptionLocalization representation | 

### Return type

[**SubscriptionLocalizationResponse**](SubscriptionLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

