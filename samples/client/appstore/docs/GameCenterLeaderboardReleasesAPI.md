# \GameCenterLeaderboardReleasesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterLeaderboardReleasesCreateInstance**](GameCenterLeaderboardReleasesAPI.md#GameCenterLeaderboardReleasesCreateInstance) | **Post** /v1/gameCenterLeaderboardReleases | 
[**GameCenterLeaderboardReleasesDeleteInstance**](GameCenterLeaderboardReleasesAPI.md#GameCenterLeaderboardReleasesDeleteInstance) | **Delete** /v1/gameCenterLeaderboardReleases/{id} | 
[**GameCenterLeaderboardReleasesGetInstance**](GameCenterLeaderboardReleasesAPI.md#GameCenterLeaderboardReleasesGetInstance) | **Get** /v1/gameCenterLeaderboardReleases/{id} | 



## GameCenterLeaderboardReleasesCreateInstance

> GameCenterLeaderboardReleaseResponse GameCenterLeaderboardReleasesCreateInstance(ctx).GameCenterLeaderboardReleaseCreateRequest(gameCenterLeaderboardReleaseCreateRequest).Execute()



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
    gameCenterLeaderboardReleaseCreateRequest := *openapiclient.NewGameCenterLeaderboardReleaseCreateRequest(*openapiclient.NewGameCenterLeaderboardReleaseCreateRequestData("Type_example", *openapiclient.NewGameCenterLeaderboardReleaseCreateRequestDataRelationships(*openapiclient.NewGameCenterAchievementReleaseCreateRequestDataRelationshipsGameCenterDetail(*openapiclient.NewAppRelationshipsGameCenterDetailData("Type_example", "Id_example")), *openapiclient.NewGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard(*openapiclient.NewGameCenterDetailRelationshipsGameCenterLeaderboardsDataInner("Type_example", "Id_example"))))) // GameCenterLeaderboardReleaseCreateRequest | GameCenterLeaderboardRelease representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardReleasesAPI.GameCenterLeaderboardReleasesCreateInstance(context.Background()).GameCenterLeaderboardReleaseCreateRequest(gameCenterLeaderboardReleaseCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardReleasesAPI.GameCenterLeaderboardReleasesCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardReleasesCreateInstance`: GameCenterLeaderboardReleaseResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardReleasesAPI.GameCenterLeaderboardReleasesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardReleasesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameCenterLeaderboardReleaseCreateRequest** | [**GameCenterLeaderboardReleaseCreateRequest**](GameCenterLeaderboardReleaseCreateRequest.md) | GameCenterLeaderboardRelease representation | 

### Return type

[**GameCenterLeaderboardReleaseResponse**](GameCenterLeaderboardReleaseResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardReleasesDeleteInstance

> GameCenterLeaderboardReleasesDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.GameCenterLeaderboardReleasesAPI.GameCenterLeaderboardReleasesDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardReleasesAPI.GameCenterLeaderboardReleasesDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterLeaderboardReleasesDeleteInstanceRequest struct via the builder pattern


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


## GameCenterLeaderboardReleasesGetInstance

> GameCenterLeaderboardReleaseResponse GameCenterLeaderboardReleasesGetInstance(ctx, id).FieldsGameCenterLeaderboardReleases(fieldsGameCenterLeaderboardReleases).Include(include).Execute()



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
    fieldsGameCenterLeaderboardReleases := []string{"FieldsGameCenterLeaderboardReleases_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardReleases (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardReleasesAPI.GameCenterLeaderboardReleasesGetInstance(context.Background(), id).FieldsGameCenterLeaderboardReleases(fieldsGameCenterLeaderboardReleases).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardReleasesAPI.GameCenterLeaderboardReleasesGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardReleasesGetInstance`: GameCenterLeaderboardReleaseResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardReleasesAPI.GameCenterLeaderboardReleasesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardReleasesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterLeaderboardReleases** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardReleases | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterLeaderboardReleaseResponse**](GameCenterLeaderboardReleaseResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

