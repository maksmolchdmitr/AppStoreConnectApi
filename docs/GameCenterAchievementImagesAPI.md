# \GameCenterAchievementImagesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterAchievementImagesCreateInstance**](GameCenterAchievementImagesAPI.md#GameCenterAchievementImagesCreateInstance) | **Post** /v1/gameCenterAchievementImages | 
[**GameCenterAchievementImagesDeleteInstance**](GameCenterAchievementImagesAPI.md#GameCenterAchievementImagesDeleteInstance) | **Delete** /v1/gameCenterAchievementImages/{id} | 
[**GameCenterAchievementImagesGetInstance**](GameCenterAchievementImagesAPI.md#GameCenterAchievementImagesGetInstance) | **Get** /v1/gameCenterAchievementImages/{id} | 
[**GameCenterAchievementImagesUpdateInstance**](GameCenterAchievementImagesAPI.md#GameCenterAchievementImagesUpdateInstance) | **Patch** /v1/gameCenterAchievementImages/{id} | 



## GameCenterAchievementImagesCreateInstance

> GameCenterAchievementImageResponse GameCenterAchievementImagesCreateInstance(ctx).GameCenterAchievementImageCreateRequest(gameCenterAchievementImageCreateRequest).Execute()



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
    gameCenterAchievementImageCreateRequest := *openapiclient.NewGameCenterAchievementImageCreateRequest(*openapiclient.NewGameCenterAchievementImageCreateRequestData("Type_example", *openapiclient.NewAppClipAdvancedExperienceImageCreateRequestDataAttributes(int32(123), "FileName_example"), *openapiclient.NewGameCenterAchievementImageCreateRequestDataRelationships(*openapiclient.NewGameCenterAchievementImageCreateRequestDataRelationshipsGameCenterAchievementLocalization(*openapiclient.NewGameCenterAchievementImageRelationshipsGameCenterAchievementLocalizationData("Type_example", "Id_example"))))) // GameCenterAchievementImageCreateRequest | GameCenterAchievementImage representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterAchievementImagesAPI.GameCenterAchievementImagesCreateInstance(context.Background()).GameCenterAchievementImageCreateRequest(gameCenterAchievementImageCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAchievementImagesAPI.GameCenterAchievementImagesCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterAchievementImagesCreateInstance`: GameCenterAchievementImageResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterAchievementImagesAPI.GameCenterAchievementImagesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterAchievementImagesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameCenterAchievementImageCreateRequest** | [**GameCenterAchievementImageCreateRequest**](GameCenterAchievementImageCreateRequest.md) | GameCenterAchievementImage representation | 

### Return type

[**GameCenterAchievementImageResponse**](GameCenterAchievementImageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterAchievementImagesDeleteInstance

> GameCenterAchievementImagesDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.GameCenterAchievementImagesAPI.GameCenterAchievementImagesDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAchievementImagesAPI.GameCenterAchievementImagesDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterAchievementImagesDeleteInstanceRequest struct via the builder pattern


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


## GameCenterAchievementImagesGetInstance

> GameCenterAchievementImageResponse GameCenterAchievementImagesGetInstance(ctx, id).FieldsGameCenterAchievementImages(fieldsGameCenterAchievementImages).Include(include).Execute()



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
    fieldsGameCenterAchievementImages := []string{"FieldsGameCenterAchievementImages_example"} // []string | the fields to include for returned resources of type gameCenterAchievementImages (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterAchievementImagesAPI.GameCenterAchievementImagesGetInstance(context.Background(), id).FieldsGameCenterAchievementImages(fieldsGameCenterAchievementImages).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAchievementImagesAPI.GameCenterAchievementImagesGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterAchievementImagesGetInstance`: GameCenterAchievementImageResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterAchievementImagesAPI.GameCenterAchievementImagesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterAchievementImagesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterAchievementImages** | **[]string** | the fields to include for returned resources of type gameCenterAchievementImages | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterAchievementImageResponse**](GameCenterAchievementImageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterAchievementImagesUpdateInstance

> GameCenterAchievementImageResponse GameCenterAchievementImagesUpdateInstance(ctx, id).GameCenterAchievementImageUpdateRequest(gameCenterAchievementImageUpdateRequest).Execute()



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
    gameCenterAchievementImageUpdateRequest := *openapiclient.NewGameCenterAchievementImageUpdateRequest(*openapiclient.NewGameCenterAchievementImageUpdateRequestData("Type_example", "Id_example")) // GameCenterAchievementImageUpdateRequest | GameCenterAchievementImage representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterAchievementImagesAPI.GameCenterAchievementImagesUpdateInstance(context.Background(), id).GameCenterAchievementImageUpdateRequest(gameCenterAchievementImageUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAchievementImagesAPI.GameCenterAchievementImagesUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterAchievementImagesUpdateInstance`: GameCenterAchievementImageResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterAchievementImagesAPI.GameCenterAchievementImagesUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterAchievementImagesUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterAchievementImageUpdateRequest** | [**GameCenterAchievementImageUpdateRequest**](GameCenterAchievementImageUpdateRequest.md) | GameCenterAchievementImage representation | 

### Return type

[**GameCenterAchievementImageResponse**](GameCenterAchievementImageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

