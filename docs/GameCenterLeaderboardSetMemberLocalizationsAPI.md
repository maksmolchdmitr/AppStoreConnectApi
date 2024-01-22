# \GameCenterLeaderboardSetMemberLocalizationsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterLeaderboardSetMemberLocalizationsCreateInstance**](GameCenterLeaderboardSetMemberLocalizationsAPI.md#GameCenterLeaderboardSetMemberLocalizationsCreateInstance) | **Post** /v1/gameCenterLeaderboardSetMemberLocalizations | 
[**GameCenterLeaderboardSetMemberLocalizationsDeleteInstance**](GameCenterLeaderboardSetMemberLocalizationsAPI.md#GameCenterLeaderboardSetMemberLocalizationsDeleteInstance) | **Delete** /v1/gameCenterLeaderboardSetMemberLocalizations/{id} | 
[**GameCenterLeaderboardSetMemberLocalizationsGameCenterLeaderboardGetToOneRelated**](GameCenterLeaderboardSetMemberLocalizationsAPI.md#GameCenterLeaderboardSetMemberLocalizationsGameCenterLeaderboardGetToOneRelated) | **Get** /v1/gameCenterLeaderboardSetMemberLocalizations/{id}/gameCenterLeaderboard | 
[**GameCenterLeaderboardSetMemberLocalizationsGameCenterLeaderboardSetGetToOneRelated**](GameCenterLeaderboardSetMemberLocalizationsAPI.md#GameCenterLeaderboardSetMemberLocalizationsGameCenterLeaderboardSetGetToOneRelated) | **Get** /v1/gameCenterLeaderboardSetMemberLocalizations/{id}/gameCenterLeaderboardSet | 
[**GameCenterLeaderboardSetMemberLocalizationsGetCollection**](GameCenterLeaderboardSetMemberLocalizationsAPI.md#GameCenterLeaderboardSetMemberLocalizationsGetCollection) | **Get** /v1/gameCenterLeaderboardSetMemberLocalizations | 
[**GameCenterLeaderboardSetMemberLocalizationsUpdateInstance**](GameCenterLeaderboardSetMemberLocalizationsAPI.md#GameCenterLeaderboardSetMemberLocalizationsUpdateInstance) | **Patch** /v1/gameCenterLeaderboardSetMemberLocalizations/{id} | 



## GameCenterLeaderboardSetMemberLocalizationsCreateInstance

> GameCenterLeaderboardSetMemberLocalizationResponse GameCenterLeaderboardSetMemberLocalizationsCreateInstance(ctx).GameCenterLeaderboardSetMemberLocalizationCreateRequest(gameCenterLeaderboardSetMemberLocalizationCreateRequest).Execute()



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
    gameCenterLeaderboardSetMemberLocalizationCreateRequest := *openapiclient.NewGameCenterLeaderboardSetMemberLocalizationCreateRequest(*openapiclient.NewGameCenterLeaderboardSetMemberLocalizationCreateRequestData("Type_example", *openapiclient.NewGameCenterLeaderboardSetMemberLocalizationCreateRequestDataRelationships(*openapiclient.NewGameCenterLeaderboardSetLocalizationCreateRequestDataRelationshipsGameCenterLeaderboardSet(*openapiclient.NewGameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner("Type_example", "Id_example")), *openapiclient.NewGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard(*openapiclient.NewGameCenterDetailRelationshipsGameCenterLeaderboardsDataInner("Type_example", "Id_example"))))) // GameCenterLeaderboardSetMemberLocalizationCreateRequest | GameCenterLeaderboardSetMemberLocalization representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardSetMemberLocalizationsAPI.GameCenterLeaderboardSetMemberLocalizationsCreateInstance(context.Background()).GameCenterLeaderboardSetMemberLocalizationCreateRequest(gameCenterLeaderboardSetMemberLocalizationCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetMemberLocalizationsAPI.GameCenterLeaderboardSetMemberLocalizationsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardSetMemberLocalizationsCreateInstance`: GameCenterLeaderboardSetMemberLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardSetMemberLocalizationsAPI.GameCenterLeaderboardSetMemberLocalizationsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetMemberLocalizationsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameCenterLeaderboardSetMemberLocalizationCreateRequest** | [**GameCenterLeaderboardSetMemberLocalizationCreateRequest**](GameCenterLeaderboardSetMemberLocalizationCreateRequest.md) | GameCenterLeaderboardSetMemberLocalization representation | 

### Return type

[**GameCenterLeaderboardSetMemberLocalizationResponse**](GameCenterLeaderboardSetMemberLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardSetMemberLocalizationsDeleteInstance

> GameCenterLeaderboardSetMemberLocalizationsDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.GameCenterLeaderboardSetMemberLocalizationsAPI.GameCenterLeaderboardSetMemberLocalizationsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetMemberLocalizationsAPI.GameCenterLeaderboardSetMemberLocalizationsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetMemberLocalizationsDeleteInstanceRequest struct via the builder pattern


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


## GameCenterLeaderboardSetMemberLocalizationsGameCenterLeaderboardGetToOneRelated

> GameCenterLeaderboardResponse GameCenterLeaderboardSetMemberLocalizationsGameCenterLeaderboardGetToOneRelated(ctx, id).FieldsGameCenterLeaderboardLocalizations(fieldsGameCenterLeaderboardLocalizations).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterLeaderboardReleases(fieldsGameCenterLeaderboardReleases).LimitGameCenterLeaderboardSets(limitGameCenterLeaderboardSets).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Include(include).Execute()



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
    resp, r, err := apiClient.GameCenterLeaderboardSetMemberLocalizationsAPI.GameCenterLeaderboardSetMemberLocalizationsGameCenterLeaderboardGetToOneRelated(context.Background(), id).FieldsGameCenterLeaderboardLocalizations(fieldsGameCenterLeaderboardLocalizations).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterLeaderboardReleases(fieldsGameCenterLeaderboardReleases).LimitGameCenterLeaderboardSets(limitGameCenterLeaderboardSets).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetMemberLocalizationsAPI.GameCenterLeaderboardSetMemberLocalizationsGameCenterLeaderboardGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardSetMemberLocalizationsGameCenterLeaderboardGetToOneRelated`: GameCenterLeaderboardResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardSetMemberLocalizationsAPI.GameCenterLeaderboardSetMemberLocalizationsGameCenterLeaderboardGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetMemberLocalizationsGameCenterLeaderboardGetToOneRelatedRequest struct via the builder pattern


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


## GameCenterLeaderboardSetMemberLocalizationsGameCenterLeaderboardSetGetToOneRelated

> GameCenterLeaderboardSetResponse GameCenterLeaderboardSetMemberLocalizationsGameCenterLeaderboardSetGetToOneRelated(ctx, id).FieldsGameCenterLeaderboardSetLocalizations(fieldsGameCenterLeaderboardSetLocalizations).FieldsGameCenterLeaderboardSetReleases(fieldsGameCenterLeaderboardSetReleases).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).LimitLocalizations(limitLocalizations).LimitGameCenterLeaderboards(limitGameCenterLeaderboards).LimitReleases(limitReleases).Include(include).Execute()



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
    fieldsGameCenterLeaderboardSetReleases := []string{"FieldsGameCenterLeaderboardSetReleases_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSetReleases (optional)
    fieldsGameCenterLeaderboardSets := []string{"FieldsGameCenterLeaderboardSets_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSets (optional)
    fieldsGameCenterGroups := []string{"FieldsGameCenterGroups_example"} // []string | the fields to include for returned resources of type gameCenterGroups (optional)
    fieldsGameCenterDetails := []string{"FieldsGameCenterDetails_example"} // []string | the fields to include for returned resources of type gameCenterDetails (optional)
    fieldsGameCenterLeaderboards := []string{"FieldsGameCenterLeaderboards_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboards (optional)
    limitLocalizations := int32(56) // int32 | maximum number of related localizations returned (when they are included) (optional)
    limitGameCenterLeaderboards := int32(56) // int32 | maximum number of related gameCenterLeaderboards returned (when they are included) (optional)
    limitReleases := int32(56) // int32 | maximum number of related releases returned (when they are included) (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardSetMemberLocalizationsAPI.GameCenterLeaderboardSetMemberLocalizationsGameCenterLeaderboardSetGetToOneRelated(context.Background(), id).FieldsGameCenterLeaderboardSetLocalizations(fieldsGameCenterLeaderboardSetLocalizations).FieldsGameCenterLeaderboardSetReleases(fieldsGameCenterLeaderboardSetReleases).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).LimitLocalizations(limitLocalizations).LimitGameCenterLeaderboards(limitGameCenterLeaderboards).LimitReleases(limitReleases).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetMemberLocalizationsAPI.GameCenterLeaderboardSetMemberLocalizationsGameCenterLeaderboardSetGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardSetMemberLocalizationsGameCenterLeaderboardSetGetToOneRelated`: GameCenterLeaderboardSetResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardSetMemberLocalizationsAPI.GameCenterLeaderboardSetMemberLocalizationsGameCenterLeaderboardSetGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetMemberLocalizationsGameCenterLeaderboardSetGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterLeaderboardSetLocalizations** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSetLocalizations | 
 **fieldsGameCenterLeaderboardSetReleases** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSetReleases | 
 **fieldsGameCenterLeaderboardSets** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSets | 
 **fieldsGameCenterGroups** | **[]string** | the fields to include for returned resources of type gameCenterGroups | 
 **fieldsGameCenterDetails** | **[]string** | the fields to include for returned resources of type gameCenterDetails | 
 **fieldsGameCenterLeaderboards** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboards | 
 **limitLocalizations** | **int32** | maximum number of related localizations returned (when they are included) | 
 **limitGameCenterLeaderboards** | **int32** | maximum number of related gameCenterLeaderboards returned (when they are included) | 
 **limitReleases** | **int32** | maximum number of related releases returned (when they are included) | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterLeaderboardSetResponse**](GameCenterLeaderboardSetResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardSetMemberLocalizationsGetCollection

> GameCenterLeaderboardSetMemberLocalizationsResponse GameCenterLeaderboardSetMemberLocalizationsGetCollection(ctx).FilterGameCenterLeaderboard(filterGameCenterLeaderboard).FilterGameCenterLeaderboardSet(filterGameCenterLeaderboardSet).FieldsGameCenterLeaderboardSetMemberLocalizations(fieldsGameCenterLeaderboardSetMemberLocalizations).Limit(limit).Include(include).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).Execute()



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
    filterGameCenterLeaderboard := []string{"Inner_example"} // []string | filter by id(s) of related 'gameCenterLeaderboard'
    filterGameCenterLeaderboardSet := []string{"Inner_example"} // []string | filter by id(s) of related 'gameCenterLeaderboardSet'
    fieldsGameCenterLeaderboardSetMemberLocalizations := []string{"FieldsGameCenterLeaderboardSetMemberLocalizations_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSetMemberLocalizations (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsGameCenterLeaderboardSets := []string{"FieldsGameCenterLeaderboardSets_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSets (optional)
    fieldsGameCenterLeaderboards := []string{"FieldsGameCenterLeaderboards_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboards (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardSetMemberLocalizationsAPI.GameCenterLeaderboardSetMemberLocalizationsGetCollection(context.Background()).FilterGameCenterLeaderboard(filterGameCenterLeaderboard).FilterGameCenterLeaderboardSet(filterGameCenterLeaderboardSet).FieldsGameCenterLeaderboardSetMemberLocalizations(fieldsGameCenterLeaderboardSetMemberLocalizations).Limit(limit).Include(include).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetMemberLocalizationsAPI.GameCenterLeaderboardSetMemberLocalizationsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardSetMemberLocalizationsGetCollection`: GameCenterLeaderboardSetMemberLocalizationsResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardSetMemberLocalizationsAPI.GameCenterLeaderboardSetMemberLocalizationsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetMemberLocalizationsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterGameCenterLeaderboard** | **[]string** | filter by id(s) of related &#39;gameCenterLeaderboard&#39; | 
 **filterGameCenterLeaderboardSet** | **[]string** | filter by id(s) of related &#39;gameCenterLeaderboardSet&#39; | 
 **fieldsGameCenterLeaderboardSetMemberLocalizations** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSetMemberLocalizations | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsGameCenterLeaderboardSets** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSets | 
 **fieldsGameCenterLeaderboards** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboards | 

### Return type

[**GameCenterLeaderboardSetMemberLocalizationsResponse**](GameCenterLeaderboardSetMemberLocalizationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardSetMemberLocalizationsUpdateInstance

> GameCenterLeaderboardSetMemberLocalizationResponse GameCenterLeaderboardSetMemberLocalizationsUpdateInstance(ctx, id).GameCenterLeaderboardSetMemberLocalizationUpdateRequest(gameCenterLeaderboardSetMemberLocalizationUpdateRequest).Execute()



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
    gameCenterLeaderboardSetMemberLocalizationUpdateRequest := *openapiclient.NewGameCenterLeaderboardSetMemberLocalizationUpdateRequest(*openapiclient.NewGameCenterLeaderboardSetMemberLocalizationUpdateRequestData("Type_example", "Id_example")) // GameCenterLeaderboardSetMemberLocalizationUpdateRequest | GameCenterLeaderboardSetMemberLocalization representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardSetMemberLocalizationsAPI.GameCenterLeaderboardSetMemberLocalizationsUpdateInstance(context.Background(), id).GameCenterLeaderboardSetMemberLocalizationUpdateRequest(gameCenterLeaderboardSetMemberLocalizationUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetMemberLocalizationsAPI.GameCenterLeaderboardSetMemberLocalizationsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardSetMemberLocalizationsUpdateInstance`: GameCenterLeaderboardSetMemberLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardSetMemberLocalizationsAPI.GameCenterLeaderboardSetMemberLocalizationsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetMemberLocalizationsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterLeaderboardSetMemberLocalizationUpdateRequest** | [**GameCenterLeaderboardSetMemberLocalizationUpdateRequest**](GameCenterLeaderboardSetMemberLocalizationUpdateRequest.md) | GameCenterLeaderboardSetMemberLocalization representation | 

### Return type

[**GameCenterLeaderboardSetMemberLocalizationResponse**](GameCenterLeaderboardSetMemberLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

