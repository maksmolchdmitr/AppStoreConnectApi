# \GameCenterLeaderboardSetLocalizationsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterLeaderboardSetLocalizationsCreateInstance**](GameCenterLeaderboardSetLocalizationsAPI.md#GameCenterLeaderboardSetLocalizationsCreateInstance) | **Post** /v1/gameCenterLeaderboardSetLocalizations | 
[**GameCenterLeaderboardSetLocalizationsDeleteInstance**](GameCenterLeaderboardSetLocalizationsAPI.md#GameCenterLeaderboardSetLocalizationsDeleteInstance) | **Delete** /v1/gameCenterLeaderboardSetLocalizations/{id} | 
[**GameCenterLeaderboardSetLocalizationsGameCenterLeaderboardSetImageGetToOneRelated**](GameCenterLeaderboardSetLocalizationsAPI.md#GameCenterLeaderboardSetLocalizationsGameCenterLeaderboardSetImageGetToOneRelated) | **Get** /v1/gameCenterLeaderboardSetLocalizations/{id}/gameCenterLeaderboardSetImage | 
[**GameCenterLeaderboardSetLocalizationsGetInstance**](GameCenterLeaderboardSetLocalizationsAPI.md#GameCenterLeaderboardSetLocalizationsGetInstance) | **Get** /v1/gameCenterLeaderboardSetLocalizations/{id} | 
[**GameCenterLeaderboardSetLocalizationsUpdateInstance**](GameCenterLeaderboardSetLocalizationsAPI.md#GameCenterLeaderboardSetLocalizationsUpdateInstance) | **Patch** /v1/gameCenterLeaderboardSetLocalizations/{id} | 



## GameCenterLeaderboardSetLocalizationsCreateInstance

> GameCenterLeaderboardSetLocalizationResponse GameCenterLeaderboardSetLocalizationsCreateInstance(ctx).GameCenterLeaderboardSetLocalizationCreateRequest(gameCenterLeaderboardSetLocalizationCreateRequest).Execute()



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
    gameCenterLeaderboardSetLocalizationCreateRequest := *openapiclient.NewGameCenterLeaderboardSetLocalizationCreateRequest(*openapiclient.NewGameCenterLeaderboardSetLocalizationCreateRequestData("Type_example", *openapiclient.NewGameCenterLeaderboardSetLocalizationCreateRequestDataAttributes("Locale_example", "Name_example"), *openapiclient.NewGameCenterLeaderboardSetLocalizationCreateRequestDataRelationships(*openapiclient.NewGameCenterLeaderboardSetLocalizationCreateRequestDataRelationshipsGameCenterLeaderboardSet(*openapiclient.NewGameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner("Type_example", "Id_example"))))) // GameCenterLeaderboardSetLocalizationCreateRequest | GameCenterLeaderboardSetLocalization representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardSetLocalizationsAPI.GameCenterLeaderboardSetLocalizationsCreateInstance(context.Background()).GameCenterLeaderboardSetLocalizationCreateRequest(gameCenterLeaderboardSetLocalizationCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetLocalizationsAPI.GameCenterLeaderboardSetLocalizationsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardSetLocalizationsCreateInstance`: GameCenterLeaderboardSetLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardSetLocalizationsAPI.GameCenterLeaderboardSetLocalizationsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetLocalizationsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameCenterLeaderboardSetLocalizationCreateRequest** | [**GameCenterLeaderboardSetLocalizationCreateRequest**](GameCenterLeaderboardSetLocalizationCreateRequest.md) | GameCenterLeaderboardSetLocalization representation | 

### Return type

[**GameCenterLeaderboardSetLocalizationResponse**](GameCenterLeaderboardSetLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardSetLocalizationsDeleteInstance

> GameCenterLeaderboardSetLocalizationsDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.GameCenterLeaderboardSetLocalizationsAPI.GameCenterLeaderboardSetLocalizationsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetLocalizationsAPI.GameCenterLeaderboardSetLocalizationsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetLocalizationsDeleteInstanceRequest struct via the builder pattern


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


## GameCenterLeaderboardSetLocalizationsGameCenterLeaderboardSetImageGetToOneRelated

> GameCenterLeaderboardSetImageResponse GameCenterLeaderboardSetLocalizationsGameCenterLeaderboardSetImageGetToOneRelated(ctx, id).FieldsGameCenterLeaderboardSetLocalizations(fieldsGameCenterLeaderboardSetLocalizations).FieldsGameCenterLeaderboardSetImages(fieldsGameCenterLeaderboardSetImages).Include(include).Execute()



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
    fieldsGameCenterLeaderboardSetLocalizations := []string{"FieldsGameCenterLeaderboardSetLocalizations_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSetLocalizations (optional)
    fieldsGameCenterLeaderboardSetImages := []string{"FieldsGameCenterLeaderboardSetImages_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSetImages (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardSetLocalizationsAPI.GameCenterLeaderboardSetLocalizationsGameCenterLeaderboardSetImageGetToOneRelated(context.Background(), id).FieldsGameCenterLeaderboardSetLocalizations(fieldsGameCenterLeaderboardSetLocalizations).FieldsGameCenterLeaderboardSetImages(fieldsGameCenterLeaderboardSetImages).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetLocalizationsAPI.GameCenterLeaderboardSetLocalizationsGameCenterLeaderboardSetImageGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardSetLocalizationsGameCenterLeaderboardSetImageGetToOneRelated`: GameCenterLeaderboardSetImageResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardSetLocalizationsAPI.GameCenterLeaderboardSetLocalizationsGameCenterLeaderboardSetImageGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetLocalizationsGameCenterLeaderboardSetImageGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterLeaderboardSetLocalizations** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSetLocalizations | 
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


## GameCenterLeaderboardSetLocalizationsGetInstance

> GameCenterLeaderboardSetLocalizationResponse GameCenterLeaderboardSetLocalizationsGetInstance(ctx, id).FieldsGameCenterLeaderboardSetLocalizations(fieldsGameCenterLeaderboardSetLocalizations).Include(include).FieldsGameCenterLeaderboardSetImages(fieldsGameCenterLeaderboardSetImages).Execute()



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
    fieldsGameCenterLeaderboardSetLocalizations := []string{"FieldsGameCenterLeaderboardSetLocalizations_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSetLocalizations (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsGameCenterLeaderboardSetImages := []string{"FieldsGameCenterLeaderboardSetImages_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSetImages (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardSetLocalizationsAPI.GameCenterLeaderboardSetLocalizationsGetInstance(context.Background(), id).FieldsGameCenterLeaderboardSetLocalizations(fieldsGameCenterLeaderboardSetLocalizations).Include(include).FieldsGameCenterLeaderboardSetImages(fieldsGameCenterLeaderboardSetImages).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetLocalizationsAPI.GameCenterLeaderboardSetLocalizationsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardSetLocalizationsGetInstance`: GameCenterLeaderboardSetLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardSetLocalizationsAPI.GameCenterLeaderboardSetLocalizationsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetLocalizationsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterLeaderboardSetLocalizations** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSetLocalizations | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsGameCenterLeaderboardSetImages** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSetImages | 

### Return type

[**GameCenterLeaderboardSetLocalizationResponse**](GameCenterLeaderboardSetLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardSetLocalizationsUpdateInstance

> GameCenterLeaderboardSetLocalizationResponse GameCenterLeaderboardSetLocalizationsUpdateInstance(ctx, id).GameCenterLeaderboardSetLocalizationUpdateRequest(gameCenterLeaderboardSetLocalizationUpdateRequest).Execute()



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
    gameCenterLeaderboardSetLocalizationUpdateRequest := *openapiclient.NewGameCenterLeaderboardSetLocalizationUpdateRequest(*openapiclient.NewGameCenterLeaderboardSetLocalizationUpdateRequestData("Type_example", "Id_example")) // GameCenterLeaderboardSetLocalizationUpdateRequest | GameCenterLeaderboardSetLocalization representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardSetLocalizationsAPI.GameCenterLeaderboardSetLocalizationsUpdateInstance(context.Background(), id).GameCenterLeaderboardSetLocalizationUpdateRequest(gameCenterLeaderboardSetLocalizationUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetLocalizationsAPI.GameCenterLeaderboardSetLocalizationsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardSetLocalizationsUpdateInstance`: GameCenterLeaderboardSetLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardSetLocalizationsAPI.GameCenterLeaderboardSetLocalizationsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetLocalizationsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterLeaderboardSetLocalizationUpdateRequest** | [**GameCenterLeaderboardSetLocalizationUpdateRequest**](GameCenterLeaderboardSetLocalizationUpdateRequest.md) | GameCenterLeaderboardSetLocalization representation | 

### Return type

[**GameCenterLeaderboardSetLocalizationResponse**](GameCenterLeaderboardSetLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

