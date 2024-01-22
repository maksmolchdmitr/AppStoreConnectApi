# \GameCenterAchievementReleasesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterAchievementReleasesCreateInstance**](GameCenterAchievementReleasesAPI.md#GameCenterAchievementReleasesCreateInstance) | **Post** /v1/gameCenterAchievementReleases | 
[**GameCenterAchievementReleasesDeleteInstance**](GameCenterAchievementReleasesAPI.md#GameCenterAchievementReleasesDeleteInstance) | **Delete** /v1/gameCenterAchievementReleases/{id} | 
[**GameCenterAchievementReleasesGetInstance**](GameCenterAchievementReleasesAPI.md#GameCenterAchievementReleasesGetInstance) | **Get** /v1/gameCenterAchievementReleases/{id} | 



## GameCenterAchievementReleasesCreateInstance

> GameCenterAchievementReleaseResponse GameCenterAchievementReleasesCreateInstance(ctx).GameCenterAchievementReleaseCreateRequest(gameCenterAchievementReleaseCreateRequest).Execute()



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
    gameCenterAchievementReleaseCreateRequest := *openapiclient.NewGameCenterAchievementReleaseCreateRequest(*openapiclient.NewGameCenterAchievementReleaseCreateRequestData("Type_example", *openapiclient.NewGameCenterAchievementReleaseCreateRequestDataRelationships(*openapiclient.NewGameCenterAchievementReleaseCreateRequestDataRelationshipsGameCenterDetail(*openapiclient.NewAppRelationshipsGameCenterDetailData("Type_example", "Id_example")), *openapiclient.NewGameCenterAchievementLocalizationCreateRequestDataRelationshipsGameCenterAchievement(*openapiclient.NewGameCenterAchievementLocalizationRelationshipsGameCenterAchievementData("Type_example", "Id_example"))))) // GameCenterAchievementReleaseCreateRequest | GameCenterAchievementRelease representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterAchievementReleasesAPI.GameCenterAchievementReleasesCreateInstance(context.Background()).GameCenterAchievementReleaseCreateRequest(gameCenterAchievementReleaseCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAchievementReleasesAPI.GameCenterAchievementReleasesCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterAchievementReleasesCreateInstance`: GameCenterAchievementReleaseResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterAchievementReleasesAPI.GameCenterAchievementReleasesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterAchievementReleasesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameCenterAchievementReleaseCreateRequest** | [**GameCenterAchievementReleaseCreateRequest**](GameCenterAchievementReleaseCreateRequest.md) | GameCenterAchievementRelease representation | 

### Return type

[**GameCenterAchievementReleaseResponse**](GameCenterAchievementReleaseResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterAchievementReleasesDeleteInstance

> GameCenterAchievementReleasesDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.GameCenterAchievementReleasesAPI.GameCenterAchievementReleasesDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAchievementReleasesAPI.GameCenterAchievementReleasesDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterAchievementReleasesDeleteInstanceRequest struct via the builder pattern


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


## GameCenterAchievementReleasesGetInstance

> GameCenterAchievementReleaseResponse GameCenterAchievementReleasesGetInstance(ctx, id).FieldsGameCenterAchievementReleases(fieldsGameCenterAchievementReleases).Include(include).Execute()



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
    fieldsGameCenterAchievementReleases := []string{"FieldsGameCenterAchievementReleases_example"} // []string | the fields to include for returned resources of type gameCenterAchievementReleases (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterAchievementReleasesAPI.GameCenterAchievementReleasesGetInstance(context.Background(), id).FieldsGameCenterAchievementReleases(fieldsGameCenterAchievementReleases).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAchievementReleasesAPI.GameCenterAchievementReleasesGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterAchievementReleasesGetInstance`: GameCenterAchievementReleaseResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterAchievementReleasesAPI.GameCenterAchievementReleasesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterAchievementReleasesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterAchievementReleases** | **[]string** | the fields to include for returned resources of type gameCenterAchievementReleases | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterAchievementReleaseResponse**](GameCenterAchievementReleaseResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

