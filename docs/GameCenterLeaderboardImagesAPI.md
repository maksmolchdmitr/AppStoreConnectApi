# \GameCenterLeaderboardImagesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterLeaderboardImagesCreateInstance**](GameCenterLeaderboardImagesAPI.md#GameCenterLeaderboardImagesCreateInstance) | **Post** /v1/gameCenterLeaderboardImages | 
[**GameCenterLeaderboardImagesDeleteInstance**](GameCenterLeaderboardImagesAPI.md#GameCenterLeaderboardImagesDeleteInstance) | **Delete** /v1/gameCenterLeaderboardImages/{id} | 
[**GameCenterLeaderboardImagesGetInstance**](GameCenterLeaderboardImagesAPI.md#GameCenterLeaderboardImagesGetInstance) | **Get** /v1/gameCenterLeaderboardImages/{id} | 
[**GameCenterLeaderboardImagesUpdateInstance**](GameCenterLeaderboardImagesAPI.md#GameCenterLeaderboardImagesUpdateInstance) | **Patch** /v1/gameCenterLeaderboardImages/{id} | 



## GameCenterLeaderboardImagesCreateInstance

> GameCenterLeaderboardImageResponse GameCenterLeaderboardImagesCreateInstance(ctx).GameCenterLeaderboardImageCreateRequest(gameCenterLeaderboardImageCreateRequest).Execute()



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
    gameCenterLeaderboardImageCreateRequest := *openapiclient.NewGameCenterLeaderboardImageCreateRequest(*openapiclient.NewGameCenterLeaderboardImageCreateRequestData("Type_example", *openapiclient.NewAppClipAdvancedExperienceImageCreateRequestDataAttributes(int32(123), "FileName_example"), *openapiclient.NewGameCenterLeaderboardImageCreateRequestDataRelationships(*openapiclient.NewGameCenterLeaderboardImageCreateRequestDataRelationshipsGameCenterLeaderboardLocalization(*openapiclient.NewGameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData("Type_example", "Id_example"))))) // GameCenterLeaderboardImageCreateRequest | GameCenterLeaderboardImage representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardImagesAPI.GameCenterLeaderboardImagesCreateInstance(context.Background()).GameCenterLeaderboardImageCreateRequest(gameCenterLeaderboardImageCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardImagesAPI.GameCenterLeaderboardImagesCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardImagesCreateInstance`: GameCenterLeaderboardImageResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardImagesAPI.GameCenterLeaderboardImagesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardImagesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameCenterLeaderboardImageCreateRequest** | [**GameCenterLeaderboardImageCreateRequest**](GameCenterLeaderboardImageCreateRequest.md) | GameCenterLeaderboardImage representation | 

### Return type

[**GameCenterLeaderboardImageResponse**](GameCenterLeaderboardImageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardImagesDeleteInstance

> GameCenterLeaderboardImagesDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.GameCenterLeaderboardImagesAPI.GameCenterLeaderboardImagesDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardImagesAPI.GameCenterLeaderboardImagesDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterLeaderboardImagesDeleteInstanceRequest struct via the builder pattern


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


## GameCenterLeaderboardImagesGetInstance

> GameCenterLeaderboardImageResponse GameCenterLeaderboardImagesGetInstance(ctx, id).FieldsGameCenterLeaderboardImages(fieldsGameCenterLeaderboardImages).Include(include).Execute()



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
    fieldsGameCenterLeaderboardImages := []string{"FieldsGameCenterLeaderboardImages_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardImages (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardImagesAPI.GameCenterLeaderboardImagesGetInstance(context.Background(), id).FieldsGameCenterLeaderboardImages(fieldsGameCenterLeaderboardImages).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardImagesAPI.GameCenterLeaderboardImagesGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardImagesGetInstance`: GameCenterLeaderboardImageResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardImagesAPI.GameCenterLeaderboardImagesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardImagesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterLeaderboardImages** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardImages | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterLeaderboardImageResponse**](GameCenterLeaderboardImageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardImagesUpdateInstance

> GameCenterLeaderboardImageResponse GameCenterLeaderboardImagesUpdateInstance(ctx, id).GameCenterLeaderboardImageUpdateRequest(gameCenterLeaderboardImageUpdateRequest).Execute()



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
    gameCenterLeaderboardImageUpdateRequest := *openapiclient.NewGameCenterLeaderboardImageUpdateRequest(*openapiclient.NewGameCenterLeaderboardImageUpdateRequestData("Type_example", "Id_example")) // GameCenterLeaderboardImageUpdateRequest | GameCenterLeaderboardImage representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardImagesAPI.GameCenterLeaderboardImagesUpdateInstance(context.Background(), id).GameCenterLeaderboardImageUpdateRequest(gameCenterLeaderboardImageUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardImagesAPI.GameCenterLeaderboardImagesUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardImagesUpdateInstance`: GameCenterLeaderboardImageResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardImagesAPI.GameCenterLeaderboardImagesUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardImagesUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterLeaderboardImageUpdateRequest** | [**GameCenterLeaderboardImageUpdateRequest**](GameCenterLeaderboardImageUpdateRequest.md) | GameCenterLeaderboardImage representation | 

### Return type

[**GameCenterLeaderboardImageResponse**](GameCenterLeaderboardImageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

