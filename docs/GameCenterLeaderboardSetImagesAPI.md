# \GameCenterLeaderboardSetImagesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterLeaderboardSetImagesCreateInstance**](GameCenterLeaderboardSetImagesAPI.md#GameCenterLeaderboardSetImagesCreateInstance) | **Post** /v1/gameCenterLeaderboardSetImages | 
[**GameCenterLeaderboardSetImagesDeleteInstance**](GameCenterLeaderboardSetImagesAPI.md#GameCenterLeaderboardSetImagesDeleteInstance) | **Delete** /v1/gameCenterLeaderboardSetImages/{id} | 
[**GameCenterLeaderboardSetImagesGetInstance**](GameCenterLeaderboardSetImagesAPI.md#GameCenterLeaderboardSetImagesGetInstance) | **Get** /v1/gameCenterLeaderboardSetImages/{id} | 
[**GameCenterLeaderboardSetImagesUpdateInstance**](GameCenterLeaderboardSetImagesAPI.md#GameCenterLeaderboardSetImagesUpdateInstance) | **Patch** /v1/gameCenterLeaderboardSetImages/{id} | 



## GameCenterLeaderboardSetImagesCreateInstance

> GameCenterLeaderboardSetImageResponse GameCenterLeaderboardSetImagesCreateInstance(ctx).GameCenterLeaderboardSetImageCreateRequest(gameCenterLeaderboardSetImageCreateRequest).Execute()



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
    gameCenterLeaderboardSetImageCreateRequest := *openapiclient.NewGameCenterLeaderboardSetImageCreateRequest(*openapiclient.NewGameCenterLeaderboardSetImageCreateRequestData("Type_example", *openapiclient.NewAppClipAdvancedExperienceImageCreateRequestDataAttributes(int32(123), "FileName_example"), *openapiclient.NewGameCenterLeaderboardSetImageCreateRequestDataRelationships(*openapiclient.NewGameCenterLeaderboardSetImageCreateRequestDataRelationshipsGameCenterLeaderboardSetLocalization(*openapiclient.NewGameCenterLeaderboardSetImageRelationshipsGameCenterLeaderboardSetLocalizationData("Type_example", "Id_example"))))) // GameCenterLeaderboardSetImageCreateRequest | GameCenterLeaderboardSetImage representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardSetImagesAPI.GameCenterLeaderboardSetImagesCreateInstance(context.Background()).GameCenterLeaderboardSetImageCreateRequest(gameCenterLeaderboardSetImageCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetImagesAPI.GameCenterLeaderboardSetImagesCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardSetImagesCreateInstance`: GameCenterLeaderboardSetImageResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardSetImagesAPI.GameCenterLeaderboardSetImagesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetImagesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameCenterLeaderboardSetImageCreateRequest** | [**GameCenterLeaderboardSetImageCreateRequest**](GameCenterLeaderboardSetImageCreateRequest.md) | GameCenterLeaderboardSetImage representation | 

### Return type

[**GameCenterLeaderboardSetImageResponse**](GameCenterLeaderboardSetImageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardSetImagesDeleteInstance

> GameCenterLeaderboardSetImagesDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.GameCenterLeaderboardSetImagesAPI.GameCenterLeaderboardSetImagesDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetImagesAPI.GameCenterLeaderboardSetImagesDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetImagesDeleteInstanceRequest struct via the builder pattern


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


## GameCenterLeaderboardSetImagesGetInstance

> GameCenterLeaderboardSetImageResponse GameCenterLeaderboardSetImagesGetInstance(ctx, id).FieldsGameCenterLeaderboardSetImages(fieldsGameCenterLeaderboardSetImages).Include(include).Execute()



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
    fieldsGameCenterLeaderboardSetImages := []string{"FieldsGameCenterLeaderboardSetImages_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSetImages (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardSetImagesAPI.GameCenterLeaderboardSetImagesGetInstance(context.Background(), id).FieldsGameCenterLeaderboardSetImages(fieldsGameCenterLeaderboardSetImages).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetImagesAPI.GameCenterLeaderboardSetImagesGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardSetImagesGetInstance`: GameCenterLeaderboardSetImageResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardSetImagesAPI.GameCenterLeaderboardSetImagesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetImagesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterLeaderboardSetImages** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSetImages | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterLeaderboardSetImageResponse**](GameCenterLeaderboardSetImageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardSetImagesUpdateInstance

> GameCenterLeaderboardSetImageResponse GameCenterLeaderboardSetImagesUpdateInstance(ctx, id).GameCenterLeaderboardSetImageUpdateRequest(gameCenterLeaderboardSetImageUpdateRequest).Execute()



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
    gameCenterLeaderboardSetImageUpdateRequest := *openapiclient.NewGameCenterLeaderboardSetImageUpdateRequest(*openapiclient.NewGameCenterLeaderboardSetImageUpdateRequestData("Type_example", "Id_example")) // GameCenterLeaderboardSetImageUpdateRequest | GameCenterLeaderboardSetImage representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardSetImagesAPI.GameCenterLeaderboardSetImagesUpdateInstance(context.Background(), id).GameCenterLeaderboardSetImageUpdateRequest(gameCenterLeaderboardSetImageUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetImagesAPI.GameCenterLeaderboardSetImagesUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardSetImagesUpdateInstance`: GameCenterLeaderboardSetImageResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardSetImagesAPI.GameCenterLeaderboardSetImagesUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetImagesUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterLeaderboardSetImageUpdateRequest** | [**GameCenterLeaderboardSetImageUpdateRequest**](GameCenterLeaderboardSetImageUpdateRequest.md) | GameCenterLeaderboardSetImage representation | 

### Return type

[**GameCenterLeaderboardSetImageResponse**](GameCenterLeaderboardSetImageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

