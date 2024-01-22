# \GameCenterLeaderboardsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterLeaderboardsCreateInstance**](GameCenterLeaderboardsAPI.md#GameCenterLeaderboardsCreateInstance) | **Post** /v1/gameCenterLeaderboards | 
[**GameCenterLeaderboardsDeleteInstance**](GameCenterLeaderboardsAPI.md#GameCenterLeaderboardsDeleteInstance) | **Delete** /v1/gameCenterLeaderboards/{id} | 
[**GameCenterLeaderboardsGetInstance**](GameCenterLeaderboardsAPI.md#GameCenterLeaderboardsGetInstance) | **Get** /v1/gameCenterLeaderboards/{id} | 
[**GameCenterLeaderboardsGroupLeaderboardGetToOneRelated**](GameCenterLeaderboardsAPI.md#GameCenterLeaderboardsGroupLeaderboardGetToOneRelated) | **Get** /v1/gameCenterLeaderboards/{id}/groupLeaderboard | 
[**GameCenterLeaderboardsGroupLeaderboardGetToOneRelationship**](GameCenterLeaderboardsAPI.md#GameCenterLeaderboardsGroupLeaderboardGetToOneRelationship) | **Get** /v1/gameCenterLeaderboards/{id}/relationships/groupLeaderboard | 
[**GameCenterLeaderboardsGroupLeaderboardUpdateToOneRelationship**](GameCenterLeaderboardsAPI.md#GameCenterLeaderboardsGroupLeaderboardUpdateToOneRelationship) | **Patch** /v1/gameCenterLeaderboards/{id}/relationships/groupLeaderboard | 
[**GameCenterLeaderboardsLocalizationsGetToManyRelated**](GameCenterLeaderboardsAPI.md#GameCenterLeaderboardsLocalizationsGetToManyRelated) | **Get** /v1/gameCenterLeaderboards/{id}/localizations | 
[**GameCenterLeaderboardsReleasesGetToManyRelated**](GameCenterLeaderboardsAPI.md#GameCenterLeaderboardsReleasesGetToManyRelated) | **Get** /v1/gameCenterLeaderboards/{id}/releases | 
[**GameCenterLeaderboardsUpdateInstance**](GameCenterLeaderboardsAPI.md#GameCenterLeaderboardsUpdateInstance) | **Patch** /v1/gameCenterLeaderboards/{id} | 



## GameCenterLeaderboardsCreateInstance

> GameCenterLeaderboardResponse GameCenterLeaderboardsCreateInstance(ctx).GameCenterLeaderboardCreateRequest(gameCenterLeaderboardCreateRequest).Execute()



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
    gameCenterLeaderboardCreateRequest := *openapiclient.NewGameCenterLeaderboardCreateRequest(*openapiclient.NewGameCenterLeaderboardCreateRequestData("Type_example", *openapiclient.NewGameCenterLeaderboardCreateRequestDataAttributes(openapiclient.GameCenterLeaderboardFormatter("INTEGER"), "ReferenceName_example", "VendorIdentifier_example", "SubmissionType_example", "ScoreSortType_example"))) // GameCenterLeaderboardCreateRequest | GameCenterLeaderboard representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardsAPI.GameCenterLeaderboardsCreateInstance(context.Background()).GameCenterLeaderboardCreateRequest(gameCenterLeaderboardCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardsAPI.GameCenterLeaderboardsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardsCreateInstance`: GameCenterLeaderboardResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardsAPI.GameCenterLeaderboardsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameCenterLeaderboardCreateRequest** | [**GameCenterLeaderboardCreateRequest**](GameCenterLeaderboardCreateRequest.md) | GameCenterLeaderboard representation | 

### Return type

[**GameCenterLeaderboardResponse**](GameCenterLeaderboardResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardsDeleteInstance

> GameCenterLeaderboardsDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.GameCenterLeaderboardsAPI.GameCenterLeaderboardsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardsAPI.GameCenterLeaderboardsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterLeaderboardsDeleteInstanceRequest struct via the builder pattern


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


## GameCenterLeaderboardsGetInstance

> GameCenterLeaderboardResponse GameCenterLeaderboardsGetInstance(ctx, id).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).Include(include).FieldsGameCenterLeaderboardLocalizations(fieldsGameCenterLeaderboardLocalizations).FieldsGameCenterLeaderboardReleases(fieldsGameCenterLeaderboardReleases).LimitGameCenterLeaderboardSets(limitGameCenterLeaderboardSets).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Execute()



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
    fieldsGameCenterLeaderboards := []string{"FieldsGameCenterLeaderboards_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboards (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsGameCenterLeaderboardLocalizations := []string{"FieldsGameCenterLeaderboardLocalizations_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardLocalizations (optional)
    fieldsGameCenterLeaderboardReleases := []string{"FieldsGameCenterLeaderboardReleases_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardReleases (optional)
    limitGameCenterLeaderboardSets := int32(56) // int32 | maximum number of related gameCenterLeaderboardSets returned (when they are included) (optional)
    limitLocalizations := int32(56) // int32 | maximum number of related localizations returned (when they are included) (optional)
    limitReleases := int32(56) // int32 | maximum number of related releases returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardsAPI.GameCenterLeaderboardsGetInstance(context.Background(), id).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).Include(include).FieldsGameCenterLeaderboardLocalizations(fieldsGameCenterLeaderboardLocalizations).FieldsGameCenterLeaderboardReleases(fieldsGameCenterLeaderboardReleases).LimitGameCenterLeaderboardSets(limitGameCenterLeaderboardSets).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardsAPI.GameCenterLeaderboardsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardsGetInstance`: GameCenterLeaderboardResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardsAPI.GameCenterLeaderboardsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterLeaderboards** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboards | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsGameCenterLeaderboardLocalizations** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardLocalizations | 
 **fieldsGameCenterLeaderboardReleases** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardReleases | 
 **limitGameCenterLeaderboardSets** | **int32** | maximum number of related gameCenterLeaderboardSets returned (when they are included) | 
 **limitLocalizations** | **int32** | maximum number of related localizations returned (when they are included) | 
 **limitReleases** | **int32** | maximum number of related releases returned (when they are included) | 

### Return type

[**GameCenterLeaderboardResponse**](GameCenterLeaderboardResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardsGroupLeaderboardGetToOneRelated

> GameCenterLeaderboardResponse GameCenterLeaderboardsGroupLeaderboardGetToOneRelated(ctx, id).FieldsGameCenterLeaderboardLocalizations(fieldsGameCenterLeaderboardLocalizations).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterLeaderboardReleases(fieldsGameCenterLeaderboardReleases).LimitGameCenterLeaderboardSets(limitGameCenterLeaderboardSets).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Include(include).Execute()



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
    fieldsGameCenterLeaderboardSets := []string{"FieldsGameCenterLeaderboardSets_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSets (optional)
    fieldsGameCenterGroups := []string{"FieldsGameCenterGroups_example"} // []string | the fields to include for returned resources of type gameCenterGroups (optional)
    fieldsGameCenterDetails := []string{"FieldsGameCenterDetails_example"} // []string | the fields to include for returned resources of type gameCenterDetails (optional)
    fieldsGameCenterLeaderboards := []string{"FieldsGameCenterLeaderboards_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboards (optional)
    fieldsGameCenterLeaderboardReleases := []string{"FieldsGameCenterLeaderboardReleases_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardReleases (optional)
    limitGameCenterLeaderboardSets := int32(56) // int32 | maximum number of related gameCenterLeaderboardSets returned (when they are included) (optional)
    limitLocalizations := int32(56) // int32 | maximum number of related localizations returned (when they are included) (optional)
    limitReleases := int32(56) // int32 | maximum number of related releases returned (when they are included) (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardsAPI.GameCenterLeaderboardsGroupLeaderboardGetToOneRelated(context.Background(), id).FieldsGameCenterLeaderboardLocalizations(fieldsGameCenterLeaderboardLocalizations).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterLeaderboardReleases(fieldsGameCenterLeaderboardReleases).LimitGameCenterLeaderboardSets(limitGameCenterLeaderboardSets).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardsAPI.GameCenterLeaderboardsGroupLeaderboardGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardsGroupLeaderboardGetToOneRelated`: GameCenterLeaderboardResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardsAPI.GameCenterLeaderboardsGroupLeaderboardGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardsGroupLeaderboardGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterLeaderboardLocalizations** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardLocalizations | 
 **fieldsGameCenterLeaderboardSets** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSets | 
 **fieldsGameCenterGroups** | **[]string** | the fields to include for returned resources of type gameCenterGroups | 
 **fieldsGameCenterDetails** | **[]string** | the fields to include for returned resources of type gameCenterDetails | 
 **fieldsGameCenterLeaderboards** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboards | 
 **fieldsGameCenterLeaderboardReleases** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardReleases | 
 **limitGameCenterLeaderboardSets** | **int32** | maximum number of related gameCenterLeaderboardSets returned (when they are included) | 
 **limitLocalizations** | **int32** | maximum number of related localizations returned (when they are included) | 
 **limitReleases** | **int32** | maximum number of related releases returned (when they are included) | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterLeaderboardResponse**](GameCenterLeaderboardResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardsGroupLeaderboardGetToOneRelationship

> GameCenterLeaderboardGroupLeaderboardLinkageResponse GameCenterLeaderboardsGroupLeaderboardGetToOneRelationship(ctx, id).Execute()



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
    resp, r, err := apiClient.GameCenterLeaderboardsAPI.GameCenterLeaderboardsGroupLeaderboardGetToOneRelationship(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardsAPI.GameCenterLeaderboardsGroupLeaderboardGetToOneRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardsGroupLeaderboardGetToOneRelationship`: GameCenterLeaderboardGroupLeaderboardLinkageResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardsAPI.GameCenterLeaderboardsGroupLeaderboardGetToOneRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardsGroupLeaderboardGetToOneRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GameCenterLeaderboardGroupLeaderboardLinkageResponse**](GameCenterLeaderboardGroupLeaderboardLinkageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardsGroupLeaderboardUpdateToOneRelationship

> GameCenterLeaderboardsGroupLeaderboardUpdateToOneRelationship(ctx, id).GameCenterLeaderboardGroupLeaderboardLinkageRequest(gameCenterLeaderboardGroupLeaderboardLinkageRequest).Execute()



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
    gameCenterLeaderboardGroupLeaderboardLinkageRequest := *openapiclient.NewGameCenterLeaderboardGroupLeaderboardLinkageRequest(*openapiclient.NewGameCenterDetailRelationshipsGameCenterLeaderboardsDataInner("Type_example", "Id_example")) // GameCenterLeaderboardGroupLeaderboardLinkageRequest | Related linkage

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GameCenterLeaderboardsAPI.GameCenterLeaderboardsGroupLeaderboardUpdateToOneRelationship(context.Background(), id).GameCenterLeaderboardGroupLeaderboardLinkageRequest(gameCenterLeaderboardGroupLeaderboardLinkageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardsAPI.GameCenterLeaderboardsGroupLeaderboardUpdateToOneRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterLeaderboardsGroupLeaderboardUpdateToOneRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterLeaderboardGroupLeaderboardLinkageRequest** | [**GameCenterLeaderboardGroupLeaderboardLinkageRequest**](GameCenterLeaderboardGroupLeaderboardLinkageRequest.md) | Related linkage | 

### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardsLocalizationsGetToManyRelated

> GameCenterLeaderboardLocalizationsResponse GameCenterLeaderboardsLocalizationsGetToManyRelated(ctx, id).FieldsGameCenterLeaderboardLocalizations(fieldsGameCenterLeaderboardLocalizations).FieldsGameCenterLeaderboardImages(fieldsGameCenterLeaderboardImages).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).Limit(limit).Include(include).Execute()



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
    fieldsGameCenterLeaderboardImages := []string{"FieldsGameCenterLeaderboardImages_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardImages (optional)
    fieldsGameCenterLeaderboards := []string{"FieldsGameCenterLeaderboards_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboards (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardsAPI.GameCenterLeaderboardsLocalizationsGetToManyRelated(context.Background(), id).FieldsGameCenterLeaderboardLocalizations(fieldsGameCenterLeaderboardLocalizations).FieldsGameCenterLeaderboardImages(fieldsGameCenterLeaderboardImages).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).Limit(limit).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardsAPI.GameCenterLeaderboardsLocalizationsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardsLocalizationsGetToManyRelated`: GameCenterLeaderboardLocalizationsResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardsAPI.GameCenterLeaderboardsLocalizationsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardsLocalizationsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterLeaderboardLocalizations** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardLocalizations | 
 **fieldsGameCenterLeaderboardImages** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardImages | 
 **fieldsGameCenterLeaderboards** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboards | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterLeaderboardLocalizationsResponse**](GameCenterLeaderboardLocalizationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardsReleasesGetToManyRelated

> GameCenterLeaderboardReleasesResponse GameCenterLeaderboardsReleasesGetToManyRelated(ctx, id).FilterLive(filterLive).FilterGameCenterDetail(filterGameCenterDetail).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterLeaderboardReleases(fieldsGameCenterLeaderboardReleases).Limit(limit).Include(include).Execute()



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
    filterLive := []string{"Inner_example"} // []string | filter by attribute 'live' (optional)
    filterGameCenterDetail := []string{"Inner_example"} // []string | filter by id(s) of related 'gameCenterDetail' (optional)
    fieldsGameCenterDetails := []string{"FieldsGameCenterDetails_example"} // []string | the fields to include for returned resources of type gameCenterDetails (optional)
    fieldsGameCenterLeaderboards := []string{"FieldsGameCenterLeaderboards_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboards (optional)
    fieldsGameCenterLeaderboardReleases := []string{"FieldsGameCenterLeaderboardReleases_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardReleases (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardsAPI.GameCenterLeaderboardsReleasesGetToManyRelated(context.Background(), id).FilterLive(filterLive).FilterGameCenterDetail(filterGameCenterDetail).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterLeaderboardReleases(fieldsGameCenterLeaderboardReleases).Limit(limit).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardsAPI.GameCenterLeaderboardsReleasesGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardsReleasesGetToManyRelated`: GameCenterLeaderboardReleasesResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardsAPI.GameCenterLeaderboardsReleasesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardsReleasesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterLive** | **[]string** | filter by attribute &#39;live&#39; | 
 **filterGameCenterDetail** | **[]string** | filter by id(s) of related &#39;gameCenterDetail&#39; | 
 **fieldsGameCenterDetails** | **[]string** | the fields to include for returned resources of type gameCenterDetails | 
 **fieldsGameCenterLeaderboards** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboards | 
 **fieldsGameCenterLeaderboardReleases** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardReleases | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterLeaderboardReleasesResponse**](GameCenterLeaderboardReleasesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardsUpdateInstance

> GameCenterLeaderboardResponse GameCenterLeaderboardsUpdateInstance(ctx, id).GameCenterLeaderboardUpdateRequest(gameCenterLeaderboardUpdateRequest).Execute()



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
    gameCenterLeaderboardUpdateRequest := *openapiclient.NewGameCenterLeaderboardUpdateRequest(*openapiclient.NewGameCenterLeaderboardUpdateRequestData("Type_example", "Id_example")) // GameCenterLeaderboardUpdateRequest | GameCenterLeaderboard representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardsAPI.GameCenterLeaderboardsUpdateInstance(context.Background(), id).GameCenterLeaderboardUpdateRequest(gameCenterLeaderboardUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardsAPI.GameCenterLeaderboardsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardsUpdateInstance`: GameCenterLeaderboardResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardsAPI.GameCenterLeaderboardsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterLeaderboardUpdateRequest** | [**GameCenterLeaderboardUpdateRequest**](GameCenterLeaderboardUpdateRequest.md) | GameCenterLeaderboard representation | 

### Return type

[**GameCenterLeaderboardResponse**](GameCenterLeaderboardResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

