# \GameCenterLeaderboardLocalizationsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterLeaderboardLocalizationsCreateInstance**](GameCenterLeaderboardLocalizationsAPI.md#GameCenterLeaderboardLocalizationsCreateInstance) | **Post** /v1/gameCenterLeaderboardLocalizations | 
[**GameCenterLeaderboardLocalizationsDeleteInstance**](GameCenterLeaderboardLocalizationsAPI.md#GameCenterLeaderboardLocalizationsDeleteInstance) | **Delete** /v1/gameCenterLeaderboardLocalizations/{id} | 
[**GameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelated**](GameCenterLeaderboardLocalizationsAPI.md#GameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelated) | **Get** /v1/gameCenterLeaderboardLocalizations/{id}/gameCenterLeaderboardImage | 
[**GameCenterLeaderboardLocalizationsGetInstance**](GameCenterLeaderboardLocalizationsAPI.md#GameCenterLeaderboardLocalizationsGetInstance) | **Get** /v1/gameCenterLeaderboardLocalizations/{id} | 
[**GameCenterLeaderboardLocalizationsUpdateInstance**](GameCenterLeaderboardLocalizationsAPI.md#GameCenterLeaderboardLocalizationsUpdateInstance) | **Patch** /v1/gameCenterLeaderboardLocalizations/{id} | 



## GameCenterLeaderboardLocalizationsCreateInstance

> GameCenterLeaderboardLocalizationResponse GameCenterLeaderboardLocalizationsCreateInstance(ctx).GameCenterLeaderboardLocalizationCreateRequest(gameCenterLeaderboardLocalizationCreateRequest).Execute()



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
    gameCenterLeaderboardLocalizationCreateRequest := *openapiclient.NewGameCenterLeaderboardLocalizationCreateRequest(*openapiclient.NewGameCenterLeaderboardLocalizationCreateRequestData("Type_example", *openapiclient.NewGameCenterLeaderboardLocalizationCreateRequestDataAttributes("Locale_example", "Name_example"), *openapiclient.NewGameCenterLeaderboardLocalizationCreateRequestDataRelationships(*openapiclient.NewGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard(*openapiclient.NewGameCenterDetailRelationshipsGameCenterLeaderboardsDataInner("Type_example", "Id_example"))))) // GameCenterLeaderboardLocalizationCreateRequest | GameCenterLeaderboardLocalization representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardLocalizationsAPI.GameCenterLeaderboardLocalizationsCreateInstance(context.Background()).GameCenterLeaderboardLocalizationCreateRequest(gameCenterLeaderboardLocalizationCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardLocalizationsAPI.GameCenterLeaderboardLocalizationsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardLocalizationsCreateInstance`: GameCenterLeaderboardLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardLocalizationsAPI.GameCenterLeaderboardLocalizationsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardLocalizationsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameCenterLeaderboardLocalizationCreateRequest** | [**GameCenterLeaderboardLocalizationCreateRequest**](GameCenterLeaderboardLocalizationCreateRequest.md) | GameCenterLeaderboardLocalization representation | 

### Return type

[**GameCenterLeaderboardLocalizationResponse**](GameCenterLeaderboardLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardLocalizationsDeleteInstance

> GameCenterLeaderboardLocalizationsDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.GameCenterLeaderboardLocalizationsAPI.GameCenterLeaderboardLocalizationsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardLocalizationsAPI.GameCenterLeaderboardLocalizationsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterLeaderboardLocalizationsDeleteInstanceRequest struct via the builder pattern


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


## GameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelated

> GameCenterLeaderboardImageResponse GameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelated(ctx, id).FieldsGameCenterLeaderboardImages(fieldsGameCenterLeaderboardImages).FieldsGameCenterLeaderboardLocalizations(fieldsGameCenterLeaderboardLocalizations).Include(include).Execute()



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
    fieldsGameCenterLeaderboardLocalizations := []string{"FieldsGameCenterLeaderboardLocalizations_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardLocalizations (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardLocalizationsAPI.GameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelated(context.Background(), id).FieldsGameCenterLeaderboardImages(fieldsGameCenterLeaderboardImages).FieldsGameCenterLeaderboardLocalizations(fieldsGameCenterLeaderboardLocalizations).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardLocalizationsAPI.GameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelated`: GameCenterLeaderboardImageResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardLocalizationsAPI.GameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterLeaderboardImages** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardImages | 
 **fieldsGameCenterLeaderboardLocalizations** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardLocalizations | 
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


## GameCenterLeaderboardLocalizationsGetInstance

> GameCenterLeaderboardLocalizationResponse GameCenterLeaderboardLocalizationsGetInstance(ctx, id).FieldsGameCenterLeaderboardLocalizations(fieldsGameCenterLeaderboardLocalizations).Include(include).FieldsGameCenterLeaderboardImages(fieldsGameCenterLeaderboardImages).Execute()



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
    fieldsGameCenterLeaderboardLocalizations := []string{"FieldsGameCenterLeaderboardLocalizations_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardLocalizations (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsGameCenterLeaderboardImages := []string{"FieldsGameCenterLeaderboardImages_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardImages (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardLocalizationsAPI.GameCenterLeaderboardLocalizationsGetInstance(context.Background(), id).FieldsGameCenterLeaderboardLocalizations(fieldsGameCenterLeaderboardLocalizations).Include(include).FieldsGameCenterLeaderboardImages(fieldsGameCenterLeaderboardImages).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardLocalizationsAPI.GameCenterLeaderboardLocalizationsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardLocalizationsGetInstance`: GameCenterLeaderboardLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardLocalizationsAPI.GameCenterLeaderboardLocalizationsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardLocalizationsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterLeaderboardLocalizations** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardLocalizations | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsGameCenterLeaderboardImages** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardImages | 

### Return type

[**GameCenterLeaderboardLocalizationResponse**](GameCenterLeaderboardLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardLocalizationsUpdateInstance

> GameCenterLeaderboardLocalizationResponse GameCenterLeaderboardLocalizationsUpdateInstance(ctx, id).GameCenterLeaderboardLocalizationUpdateRequest(gameCenterLeaderboardLocalizationUpdateRequest).Execute()



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
    gameCenterLeaderboardLocalizationUpdateRequest := *openapiclient.NewGameCenterLeaderboardLocalizationUpdateRequest(*openapiclient.NewGameCenterLeaderboardLocalizationUpdateRequestData("Type_example", "Id_example")) // GameCenterLeaderboardLocalizationUpdateRequest | GameCenterLeaderboardLocalization representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardLocalizationsAPI.GameCenterLeaderboardLocalizationsUpdateInstance(context.Background(), id).GameCenterLeaderboardLocalizationUpdateRequest(gameCenterLeaderboardLocalizationUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardLocalizationsAPI.GameCenterLeaderboardLocalizationsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardLocalizationsUpdateInstance`: GameCenterLeaderboardLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardLocalizationsAPI.GameCenterLeaderboardLocalizationsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardLocalizationsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterLeaderboardLocalizationUpdateRequest** | [**GameCenterLeaderboardLocalizationUpdateRequest**](GameCenterLeaderboardLocalizationUpdateRequest.md) | GameCenterLeaderboardLocalization representation | 

### Return type

[**GameCenterLeaderboardLocalizationResponse**](GameCenterLeaderboardLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

