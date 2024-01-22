# \GameCenterLeaderboardSetReleasesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterLeaderboardSetReleasesCreateInstance**](GameCenterLeaderboardSetReleasesAPI.md#GameCenterLeaderboardSetReleasesCreateInstance) | **Post** /v1/gameCenterLeaderboardSetReleases | 
[**GameCenterLeaderboardSetReleasesDeleteInstance**](GameCenterLeaderboardSetReleasesAPI.md#GameCenterLeaderboardSetReleasesDeleteInstance) | **Delete** /v1/gameCenterLeaderboardSetReleases/{id} | 
[**GameCenterLeaderboardSetReleasesGetInstance**](GameCenterLeaderboardSetReleasesAPI.md#GameCenterLeaderboardSetReleasesGetInstance) | **Get** /v1/gameCenterLeaderboardSetReleases/{id} | 



## GameCenterLeaderboardSetReleasesCreateInstance

> GameCenterLeaderboardSetReleaseResponse GameCenterLeaderboardSetReleasesCreateInstance(ctx).GameCenterLeaderboardSetReleaseCreateRequest(gameCenterLeaderboardSetReleaseCreateRequest).Execute()



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
    gameCenterLeaderboardSetReleaseCreateRequest := *openapiclient.NewGameCenterLeaderboardSetReleaseCreateRequest(*openapiclient.NewGameCenterLeaderboardSetReleaseCreateRequestData("Type_example", *openapiclient.NewGameCenterLeaderboardSetReleaseCreateRequestDataRelationships(*openapiclient.NewGameCenterAchievementReleaseCreateRequestDataRelationshipsGameCenterDetail(*openapiclient.NewAppRelationshipsGameCenterDetailData("Type_example", "Id_example")), *openapiclient.NewGameCenterLeaderboardSetLocalizationCreateRequestDataRelationshipsGameCenterLeaderboardSet(*openapiclient.NewGameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner("Type_example", "Id_example"))))) // GameCenterLeaderboardSetReleaseCreateRequest | GameCenterLeaderboardSetRelease representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardSetReleasesAPI.GameCenterLeaderboardSetReleasesCreateInstance(context.Background()).GameCenterLeaderboardSetReleaseCreateRequest(gameCenterLeaderboardSetReleaseCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetReleasesAPI.GameCenterLeaderboardSetReleasesCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardSetReleasesCreateInstance`: GameCenterLeaderboardSetReleaseResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardSetReleasesAPI.GameCenterLeaderboardSetReleasesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetReleasesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameCenterLeaderboardSetReleaseCreateRequest** | [**GameCenterLeaderboardSetReleaseCreateRequest**](GameCenterLeaderboardSetReleaseCreateRequest.md) | GameCenterLeaderboardSetRelease representation | 

### Return type

[**GameCenterLeaderboardSetReleaseResponse**](GameCenterLeaderboardSetReleaseResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardSetReleasesDeleteInstance

> GameCenterLeaderboardSetReleasesDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.GameCenterLeaderboardSetReleasesAPI.GameCenterLeaderboardSetReleasesDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetReleasesAPI.GameCenterLeaderboardSetReleasesDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetReleasesDeleteInstanceRequest struct via the builder pattern


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


## GameCenterLeaderboardSetReleasesGetInstance

> GameCenterLeaderboardSetReleaseResponse GameCenterLeaderboardSetReleasesGetInstance(ctx, id).FieldsGameCenterLeaderboardSetReleases(fieldsGameCenterLeaderboardSetReleases).Include(include).Execute()



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
    fieldsGameCenterLeaderboardSetReleases := []string{"FieldsGameCenterLeaderboardSetReleases_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSetReleases (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardSetReleasesAPI.GameCenterLeaderboardSetReleasesGetInstance(context.Background(), id).FieldsGameCenterLeaderboardSetReleases(fieldsGameCenterLeaderboardSetReleases).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetReleasesAPI.GameCenterLeaderboardSetReleasesGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardSetReleasesGetInstance`: GameCenterLeaderboardSetReleaseResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardSetReleasesAPI.GameCenterLeaderboardSetReleasesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetReleasesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterLeaderboardSetReleases** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSetReleases | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterLeaderboardSetReleaseResponse**](GameCenterLeaderboardSetReleaseResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

