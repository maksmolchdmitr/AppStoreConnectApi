# \AppClipAdvancedExperienceImagesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppClipAdvancedExperienceImagesCreateInstance**](AppClipAdvancedExperienceImagesAPI.md#AppClipAdvancedExperienceImagesCreateInstance) | **Post** /v1/appClipAdvancedExperienceImages | 
[**AppClipAdvancedExperienceImagesGetInstance**](AppClipAdvancedExperienceImagesAPI.md#AppClipAdvancedExperienceImagesGetInstance) | **Get** /v1/appClipAdvancedExperienceImages/{id} | 
[**AppClipAdvancedExperienceImagesUpdateInstance**](AppClipAdvancedExperienceImagesAPI.md#AppClipAdvancedExperienceImagesUpdateInstance) | **Patch** /v1/appClipAdvancedExperienceImages/{id} | 



## AppClipAdvancedExperienceImagesCreateInstance

> AppClipAdvancedExperienceImageResponse AppClipAdvancedExperienceImagesCreateInstance(ctx).AppClipAdvancedExperienceImageCreateRequest(appClipAdvancedExperienceImageCreateRequest).Execute()



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
    appClipAdvancedExperienceImageCreateRequest := *openapiclient.NewAppClipAdvancedExperienceImageCreateRequest(*openapiclient.NewAppClipAdvancedExperienceImageCreateRequestData("Type_example", *openapiclient.NewAppClipAdvancedExperienceImageCreateRequestDataAttributes(int32(123), "FileName_example"))) // AppClipAdvancedExperienceImageCreateRequest | AppClipAdvancedExperienceImage representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppClipAdvancedExperienceImagesAPI.AppClipAdvancedExperienceImagesCreateInstance(context.Background()).AppClipAdvancedExperienceImageCreateRequest(appClipAdvancedExperienceImageCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppClipAdvancedExperienceImagesAPI.AppClipAdvancedExperienceImagesCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppClipAdvancedExperienceImagesCreateInstance`: AppClipAdvancedExperienceImageResponse
    fmt.Fprintf(os.Stdout, "Response from `AppClipAdvancedExperienceImagesAPI.AppClipAdvancedExperienceImagesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppClipAdvancedExperienceImagesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appClipAdvancedExperienceImageCreateRequest** | [**AppClipAdvancedExperienceImageCreateRequest**](AppClipAdvancedExperienceImageCreateRequest.md) | AppClipAdvancedExperienceImage representation | 

### Return type

[**AppClipAdvancedExperienceImageResponse**](AppClipAdvancedExperienceImageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppClipAdvancedExperienceImagesGetInstance

> AppClipAdvancedExperienceImageResponse AppClipAdvancedExperienceImagesGetInstance(ctx, id).FieldsAppClipAdvancedExperienceImages(fieldsAppClipAdvancedExperienceImages).Execute()



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
    fieldsAppClipAdvancedExperienceImages := []string{"FieldsAppClipAdvancedExperienceImages_example"} // []string | the fields to include for returned resources of type appClipAdvancedExperienceImages (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppClipAdvancedExperienceImagesAPI.AppClipAdvancedExperienceImagesGetInstance(context.Background(), id).FieldsAppClipAdvancedExperienceImages(fieldsAppClipAdvancedExperienceImages).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppClipAdvancedExperienceImagesAPI.AppClipAdvancedExperienceImagesGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppClipAdvancedExperienceImagesGetInstance`: AppClipAdvancedExperienceImageResponse
    fmt.Fprintf(os.Stdout, "Response from `AppClipAdvancedExperienceImagesAPI.AppClipAdvancedExperienceImagesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppClipAdvancedExperienceImagesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppClipAdvancedExperienceImages** | **[]string** | the fields to include for returned resources of type appClipAdvancedExperienceImages | 

### Return type

[**AppClipAdvancedExperienceImageResponse**](AppClipAdvancedExperienceImageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppClipAdvancedExperienceImagesUpdateInstance

> AppClipAdvancedExperienceImageResponse AppClipAdvancedExperienceImagesUpdateInstance(ctx, id).AppClipAdvancedExperienceImageUpdateRequest(appClipAdvancedExperienceImageUpdateRequest).Execute()



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
    appClipAdvancedExperienceImageUpdateRequest := *openapiclient.NewAppClipAdvancedExperienceImageUpdateRequest(*openapiclient.NewAppClipAdvancedExperienceImageUpdateRequestData("Type_example", "Id_example")) // AppClipAdvancedExperienceImageUpdateRequest | AppClipAdvancedExperienceImage representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppClipAdvancedExperienceImagesAPI.AppClipAdvancedExperienceImagesUpdateInstance(context.Background(), id).AppClipAdvancedExperienceImageUpdateRequest(appClipAdvancedExperienceImageUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppClipAdvancedExperienceImagesAPI.AppClipAdvancedExperienceImagesUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppClipAdvancedExperienceImagesUpdateInstance`: AppClipAdvancedExperienceImageResponse
    fmt.Fprintf(os.Stdout, "Response from `AppClipAdvancedExperienceImagesAPI.AppClipAdvancedExperienceImagesUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppClipAdvancedExperienceImagesUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appClipAdvancedExperienceImageUpdateRequest** | [**AppClipAdvancedExperienceImageUpdateRequest**](AppClipAdvancedExperienceImageUpdateRequest.md) | AppClipAdvancedExperienceImage representation | 

### Return type

[**AppClipAdvancedExperienceImageResponse**](AppClipAdvancedExperienceImageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

