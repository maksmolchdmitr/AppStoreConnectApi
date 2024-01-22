# \PromotedPurchaseImagesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PromotedPurchaseImagesCreateInstance**](PromotedPurchaseImagesAPI.md#PromotedPurchaseImagesCreateInstance) | **Post** /v1/promotedPurchaseImages | 
[**PromotedPurchaseImagesDeleteInstance**](PromotedPurchaseImagesAPI.md#PromotedPurchaseImagesDeleteInstance) | **Delete** /v1/promotedPurchaseImages/{id} | 
[**PromotedPurchaseImagesGetInstance**](PromotedPurchaseImagesAPI.md#PromotedPurchaseImagesGetInstance) | **Get** /v1/promotedPurchaseImages/{id} | 
[**PromotedPurchaseImagesUpdateInstance**](PromotedPurchaseImagesAPI.md#PromotedPurchaseImagesUpdateInstance) | **Patch** /v1/promotedPurchaseImages/{id} | 



## PromotedPurchaseImagesCreateInstance

> PromotedPurchaseImageResponse PromotedPurchaseImagesCreateInstance(ctx).PromotedPurchaseImageCreateRequest(promotedPurchaseImageCreateRequest).Execute()



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
    promotedPurchaseImageCreateRequest := *openapiclient.NewPromotedPurchaseImageCreateRequest(*openapiclient.NewPromotedPurchaseImageCreateRequestData("Type_example", *openapiclient.NewAppClipAdvancedExperienceImageCreateRequestDataAttributes(int32(123), "FileName_example"), *openapiclient.NewPromotedPurchaseImageCreateRequestDataRelationships(*openapiclient.NewPromotedPurchaseImageCreateRequestDataRelationshipsPromotedPurchase(*openapiclient.NewAppRelationshipsPromotedPurchasesDataInner("Type_example", "Id_example"))))) // PromotedPurchaseImageCreateRequest | PromotedPurchaseImage representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromotedPurchaseImagesAPI.PromotedPurchaseImagesCreateInstance(context.Background()).PromotedPurchaseImageCreateRequest(promotedPurchaseImageCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotedPurchaseImagesAPI.PromotedPurchaseImagesCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PromotedPurchaseImagesCreateInstance`: PromotedPurchaseImageResponse
    fmt.Fprintf(os.Stdout, "Response from `PromotedPurchaseImagesAPI.PromotedPurchaseImagesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPromotedPurchaseImagesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **promotedPurchaseImageCreateRequest** | [**PromotedPurchaseImageCreateRequest**](PromotedPurchaseImageCreateRequest.md) | PromotedPurchaseImage representation | 

### Return type

[**PromotedPurchaseImageResponse**](PromotedPurchaseImageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromotedPurchaseImagesDeleteInstance

> PromotedPurchaseImagesDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.PromotedPurchaseImagesAPI.PromotedPurchaseImagesDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotedPurchaseImagesAPI.PromotedPurchaseImagesDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPromotedPurchaseImagesDeleteInstanceRequest struct via the builder pattern


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


## PromotedPurchaseImagesGetInstance

> PromotedPurchaseImageResponse PromotedPurchaseImagesGetInstance(ctx, id).FieldsPromotedPurchaseImages(fieldsPromotedPurchaseImages).Include(include).Execute()



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
    fieldsPromotedPurchaseImages := []string{"FieldsPromotedPurchaseImages_example"} // []string | the fields to include for returned resources of type promotedPurchaseImages (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromotedPurchaseImagesAPI.PromotedPurchaseImagesGetInstance(context.Background(), id).FieldsPromotedPurchaseImages(fieldsPromotedPurchaseImages).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotedPurchaseImagesAPI.PromotedPurchaseImagesGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PromotedPurchaseImagesGetInstance`: PromotedPurchaseImageResponse
    fmt.Fprintf(os.Stdout, "Response from `PromotedPurchaseImagesAPI.PromotedPurchaseImagesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiPromotedPurchaseImagesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsPromotedPurchaseImages** | **[]string** | the fields to include for returned resources of type promotedPurchaseImages | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**PromotedPurchaseImageResponse**](PromotedPurchaseImageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromotedPurchaseImagesUpdateInstance

> PromotedPurchaseImageResponse PromotedPurchaseImagesUpdateInstance(ctx, id).PromotedPurchaseImageUpdateRequest(promotedPurchaseImageUpdateRequest).Execute()



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
    promotedPurchaseImageUpdateRequest := *openapiclient.NewPromotedPurchaseImageUpdateRequest(*openapiclient.NewPromotedPurchaseImageUpdateRequestData("Type_example", "Id_example")) // PromotedPurchaseImageUpdateRequest | PromotedPurchaseImage representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PromotedPurchaseImagesAPI.PromotedPurchaseImagesUpdateInstance(context.Background(), id).PromotedPurchaseImageUpdateRequest(promotedPurchaseImageUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PromotedPurchaseImagesAPI.PromotedPurchaseImagesUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PromotedPurchaseImagesUpdateInstance`: PromotedPurchaseImageResponse
    fmt.Fprintf(os.Stdout, "Response from `PromotedPurchaseImagesAPI.PromotedPurchaseImagesUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiPromotedPurchaseImagesUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **promotedPurchaseImageUpdateRequest** | [**PromotedPurchaseImageUpdateRequest**](PromotedPurchaseImageUpdateRequest.md) | PromotedPurchaseImage representation | 

### Return type

[**PromotedPurchaseImageResponse**](PromotedPurchaseImageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

