# \AppPreOrdersAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppPreOrdersCreateInstance**](AppPreOrdersAPI.md#AppPreOrdersCreateInstance) | **Post** /v1/appPreOrders | 
[**AppPreOrdersDeleteInstance**](AppPreOrdersAPI.md#AppPreOrdersDeleteInstance) | **Delete** /v1/appPreOrders/{id} | 
[**AppPreOrdersGetInstance**](AppPreOrdersAPI.md#AppPreOrdersGetInstance) | **Get** /v1/appPreOrders/{id} | 
[**AppPreOrdersUpdateInstance**](AppPreOrdersAPI.md#AppPreOrdersUpdateInstance) | **Patch** /v1/appPreOrders/{id} | 



## AppPreOrdersCreateInstance

> AppPreOrderResponse AppPreOrdersCreateInstance(ctx).AppPreOrderCreateRequest(appPreOrderCreateRequest).Execute()



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
    appPreOrderCreateRequest := *openapiclient.NewAppPreOrderCreateRequest(*openapiclient.NewAppPreOrderCreateRequestData("Type_example", *openapiclient.NewAppEventCreateRequestDataRelationships(*openapiclient.NewAppAvailabilityV2CreateRequestDataRelationshipsApp(*openapiclient.NewAppAvailabilityV2CreateRequestDataRelationshipsAppData("Type_example", "Id_example"))))) // AppPreOrderCreateRequest | AppPreOrder representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppPreOrdersAPI.AppPreOrdersCreateInstance(context.Background()).AppPreOrderCreateRequest(appPreOrderCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppPreOrdersAPI.AppPreOrdersCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPreOrdersCreateInstance`: AppPreOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `AppPreOrdersAPI.AppPreOrdersCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppPreOrdersCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appPreOrderCreateRequest** | [**AppPreOrderCreateRequest**](AppPreOrderCreateRequest.md) | AppPreOrder representation | 

### Return type

[**AppPreOrderResponse**](AppPreOrderResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPreOrdersDeleteInstance

> AppPreOrdersDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.AppPreOrdersAPI.AppPreOrdersDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppPreOrdersAPI.AppPreOrdersDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppPreOrdersDeleteInstanceRequest struct via the builder pattern


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


## AppPreOrdersGetInstance

> AppPreOrderResponse AppPreOrdersGetInstance(ctx, id).FieldsAppPreOrders(fieldsAppPreOrders).Include(include).Execute()



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
    fieldsAppPreOrders := []string{"FieldsAppPreOrders_example"} // []string | the fields to include for returned resources of type appPreOrders (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppPreOrdersAPI.AppPreOrdersGetInstance(context.Background(), id).FieldsAppPreOrders(fieldsAppPreOrders).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppPreOrdersAPI.AppPreOrdersGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPreOrdersGetInstance`: AppPreOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `AppPreOrdersAPI.AppPreOrdersGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppPreOrdersGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppPreOrders** | **[]string** | the fields to include for returned resources of type appPreOrders | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**AppPreOrderResponse**](AppPreOrderResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPreOrdersUpdateInstance

> AppPreOrderResponse AppPreOrdersUpdateInstance(ctx, id).AppPreOrderUpdateRequest(appPreOrderUpdateRequest).Execute()



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
    appPreOrderUpdateRequest := *openapiclient.NewAppPreOrderUpdateRequest(*openapiclient.NewAppPreOrderUpdateRequestData("Type_example", "Id_example")) // AppPreOrderUpdateRequest | AppPreOrder representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppPreOrdersAPI.AppPreOrdersUpdateInstance(context.Background(), id).AppPreOrderUpdateRequest(appPreOrderUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppPreOrdersAPI.AppPreOrdersUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPreOrdersUpdateInstance`: AppPreOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `AppPreOrdersAPI.AppPreOrdersUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppPreOrdersUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appPreOrderUpdateRequest** | [**AppPreOrderUpdateRequest**](AppPreOrderUpdateRequest.md) | AppPreOrder representation | 

### Return type

[**AppPreOrderResponse**](AppPreOrderResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

