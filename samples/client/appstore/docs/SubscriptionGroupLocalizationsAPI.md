# \SubscriptionGroupLocalizationsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionGroupLocalizationsCreateInstance**](SubscriptionGroupLocalizationsAPI.md#SubscriptionGroupLocalizationsCreateInstance) | **Post** /v1/subscriptionGroupLocalizations | 
[**SubscriptionGroupLocalizationsDeleteInstance**](SubscriptionGroupLocalizationsAPI.md#SubscriptionGroupLocalizationsDeleteInstance) | **Delete** /v1/subscriptionGroupLocalizations/{id} | 
[**SubscriptionGroupLocalizationsGetInstance**](SubscriptionGroupLocalizationsAPI.md#SubscriptionGroupLocalizationsGetInstance) | **Get** /v1/subscriptionGroupLocalizations/{id} | 
[**SubscriptionGroupLocalizationsUpdateInstance**](SubscriptionGroupLocalizationsAPI.md#SubscriptionGroupLocalizationsUpdateInstance) | **Patch** /v1/subscriptionGroupLocalizations/{id} | 



## SubscriptionGroupLocalizationsCreateInstance

> SubscriptionGroupLocalizationResponse SubscriptionGroupLocalizationsCreateInstance(ctx).SubscriptionGroupLocalizationCreateRequest(subscriptionGroupLocalizationCreateRequest).Execute()



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
    subscriptionGroupLocalizationCreateRequest := *openapiclient.NewSubscriptionGroupLocalizationCreateRequest(*openapiclient.NewSubscriptionGroupLocalizationCreateRequestData("Type_example", *openapiclient.NewSubscriptionGroupLocalizationCreateRequestDataAttributes("Name_example", "Locale_example"), *openapiclient.NewSubscriptionGroupLocalizationCreateRequestDataRelationships(*openapiclient.NewSubscriptionGroupLocalizationCreateRequestDataRelationshipsSubscriptionGroup(*openapiclient.NewAppRelationshipsSubscriptionGroupsDataInner("Type_example", "Id_example"))))) // SubscriptionGroupLocalizationCreateRequest | SubscriptionGroupLocalization representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionGroupLocalizationsAPI.SubscriptionGroupLocalizationsCreateInstance(context.Background()).SubscriptionGroupLocalizationCreateRequest(subscriptionGroupLocalizationCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionGroupLocalizationsAPI.SubscriptionGroupLocalizationsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionGroupLocalizationsCreateInstance`: SubscriptionGroupLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionGroupLocalizationsAPI.SubscriptionGroupLocalizationsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionGroupLocalizationsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionGroupLocalizationCreateRequest** | [**SubscriptionGroupLocalizationCreateRequest**](SubscriptionGroupLocalizationCreateRequest.md) | SubscriptionGroupLocalization representation | 

### Return type

[**SubscriptionGroupLocalizationResponse**](SubscriptionGroupLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionGroupLocalizationsDeleteInstance

> SubscriptionGroupLocalizationsDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.SubscriptionGroupLocalizationsAPI.SubscriptionGroupLocalizationsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionGroupLocalizationsAPI.SubscriptionGroupLocalizationsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSubscriptionGroupLocalizationsDeleteInstanceRequest struct via the builder pattern


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


## SubscriptionGroupLocalizationsGetInstance

> SubscriptionGroupLocalizationResponse SubscriptionGroupLocalizationsGetInstance(ctx, id).FieldsSubscriptionGroupLocalizations(fieldsSubscriptionGroupLocalizations).Include(include).Execute()



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
    fieldsSubscriptionGroupLocalizations := []string{"FieldsSubscriptionGroupLocalizations_example"} // []string | the fields to include for returned resources of type subscriptionGroupLocalizations (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionGroupLocalizationsAPI.SubscriptionGroupLocalizationsGetInstance(context.Background(), id).FieldsSubscriptionGroupLocalizations(fieldsSubscriptionGroupLocalizations).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionGroupLocalizationsAPI.SubscriptionGroupLocalizationsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionGroupLocalizationsGetInstance`: SubscriptionGroupLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionGroupLocalizationsAPI.SubscriptionGroupLocalizationsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionGroupLocalizationsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsSubscriptionGroupLocalizations** | **[]string** | the fields to include for returned resources of type subscriptionGroupLocalizations | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**SubscriptionGroupLocalizationResponse**](SubscriptionGroupLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionGroupLocalizationsUpdateInstance

> SubscriptionGroupLocalizationResponse SubscriptionGroupLocalizationsUpdateInstance(ctx, id).SubscriptionGroupLocalizationUpdateRequest(subscriptionGroupLocalizationUpdateRequest).Execute()



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
    subscriptionGroupLocalizationUpdateRequest := *openapiclient.NewSubscriptionGroupLocalizationUpdateRequest(*openapiclient.NewSubscriptionGroupLocalizationUpdateRequestData("Type_example", "Id_example")) // SubscriptionGroupLocalizationUpdateRequest | SubscriptionGroupLocalization representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionGroupLocalizationsAPI.SubscriptionGroupLocalizationsUpdateInstance(context.Background(), id).SubscriptionGroupLocalizationUpdateRequest(subscriptionGroupLocalizationUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionGroupLocalizationsAPI.SubscriptionGroupLocalizationsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionGroupLocalizationsUpdateInstance`: SubscriptionGroupLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionGroupLocalizationsAPI.SubscriptionGroupLocalizationsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionGroupLocalizationsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionGroupLocalizationUpdateRequest** | [**SubscriptionGroupLocalizationUpdateRequest**](SubscriptionGroupLocalizationUpdateRequest.md) | SubscriptionGroupLocalization representation | 

### Return type

[**SubscriptionGroupLocalizationResponse**](SubscriptionGroupLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

