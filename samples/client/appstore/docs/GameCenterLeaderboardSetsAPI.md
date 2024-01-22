# \GameCenterLeaderboardSetsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterLeaderboardSetsCreateInstance**](GameCenterLeaderboardSetsAPI.md#GameCenterLeaderboardSetsCreateInstance) | **Post** /v1/gameCenterLeaderboardSets | 
[**GameCenterLeaderboardSetsDeleteInstance**](GameCenterLeaderboardSetsAPI.md#GameCenterLeaderboardSetsDeleteInstance) | **Delete** /v1/gameCenterLeaderboardSets/{id} | 
[**GameCenterLeaderboardSetsGameCenterLeaderboardsCreateToManyRelationship**](GameCenterLeaderboardSetsAPI.md#GameCenterLeaderboardSetsGameCenterLeaderboardsCreateToManyRelationship) | **Post** /v1/gameCenterLeaderboardSets/{id}/relationships/gameCenterLeaderboards | 
[**GameCenterLeaderboardSetsGameCenterLeaderboardsDeleteToManyRelationship**](GameCenterLeaderboardSetsAPI.md#GameCenterLeaderboardSetsGameCenterLeaderboardsDeleteToManyRelationship) | **Delete** /v1/gameCenterLeaderboardSets/{id}/relationships/gameCenterLeaderboards | 
[**GameCenterLeaderboardSetsGameCenterLeaderboardsGetToManyRelated**](GameCenterLeaderboardSetsAPI.md#GameCenterLeaderboardSetsGameCenterLeaderboardsGetToManyRelated) | **Get** /v1/gameCenterLeaderboardSets/{id}/gameCenterLeaderboards | 
[**GameCenterLeaderboardSetsGameCenterLeaderboardsGetToManyRelationship**](GameCenterLeaderboardSetsAPI.md#GameCenterLeaderboardSetsGameCenterLeaderboardsGetToManyRelationship) | **Get** /v1/gameCenterLeaderboardSets/{id}/relationships/gameCenterLeaderboards | 
[**GameCenterLeaderboardSetsGameCenterLeaderboardsReplaceToManyRelationship**](GameCenterLeaderboardSetsAPI.md#GameCenterLeaderboardSetsGameCenterLeaderboardsReplaceToManyRelationship) | **Patch** /v1/gameCenterLeaderboardSets/{id}/relationships/gameCenterLeaderboards | 
[**GameCenterLeaderboardSetsGetInstance**](GameCenterLeaderboardSetsAPI.md#GameCenterLeaderboardSetsGetInstance) | **Get** /v1/gameCenterLeaderboardSets/{id} | 
[**GameCenterLeaderboardSetsGroupLeaderboardSetGetToOneRelated**](GameCenterLeaderboardSetsAPI.md#GameCenterLeaderboardSetsGroupLeaderboardSetGetToOneRelated) | **Get** /v1/gameCenterLeaderboardSets/{id}/groupLeaderboardSet | 
[**GameCenterLeaderboardSetsGroupLeaderboardSetGetToOneRelationship**](GameCenterLeaderboardSetsAPI.md#GameCenterLeaderboardSetsGroupLeaderboardSetGetToOneRelationship) | **Get** /v1/gameCenterLeaderboardSets/{id}/relationships/groupLeaderboardSet | 
[**GameCenterLeaderboardSetsGroupLeaderboardSetUpdateToOneRelationship**](GameCenterLeaderboardSetsAPI.md#GameCenterLeaderboardSetsGroupLeaderboardSetUpdateToOneRelationship) | **Patch** /v1/gameCenterLeaderboardSets/{id}/relationships/groupLeaderboardSet | 
[**GameCenterLeaderboardSetsLocalizationsGetToManyRelated**](GameCenterLeaderboardSetsAPI.md#GameCenterLeaderboardSetsLocalizationsGetToManyRelated) | **Get** /v1/gameCenterLeaderboardSets/{id}/localizations | 
[**GameCenterLeaderboardSetsReleasesGetToManyRelated**](GameCenterLeaderboardSetsAPI.md#GameCenterLeaderboardSetsReleasesGetToManyRelated) | **Get** /v1/gameCenterLeaderboardSets/{id}/releases | 
[**GameCenterLeaderboardSetsUpdateInstance**](GameCenterLeaderboardSetsAPI.md#GameCenterLeaderboardSetsUpdateInstance) | **Patch** /v1/gameCenterLeaderboardSets/{id} | 



## GameCenterLeaderboardSetsCreateInstance

> GameCenterLeaderboardSetResponse GameCenterLeaderboardSetsCreateInstance(ctx).GameCenterLeaderboardSetCreateRequest(gameCenterLeaderboardSetCreateRequest).Execute()



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
    gameCenterLeaderboardSetCreateRequest := *openapiclient.NewGameCenterLeaderboardSetCreateRequest(*openapiclient.NewGameCenterLeaderboardSetCreateRequestData("Type_example", *openapiclient.NewGameCenterLeaderboardSetCreateRequestDataAttributes("ReferenceName_example", "VendorIdentifier_example"))) // GameCenterLeaderboardSetCreateRequest | GameCenterLeaderboardSet representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsCreateInstance(context.Background()).GameCenterLeaderboardSetCreateRequest(gameCenterLeaderboardSetCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardSetsCreateInstance`: GameCenterLeaderboardSetResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameCenterLeaderboardSetCreateRequest** | [**GameCenterLeaderboardSetCreateRequest**](GameCenterLeaderboardSetCreateRequest.md) | GameCenterLeaderboardSet representation | 

### Return type

[**GameCenterLeaderboardSetResponse**](GameCenterLeaderboardSetResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardSetsDeleteInstance

> GameCenterLeaderboardSetsDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetsDeleteInstanceRequest struct via the builder pattern


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


## GameCenterLeaderboardSetsGameCenterLeaderboardsCreateToManyRelationship

> GameCenterLeaderboardSetsGameCenterLeaderboardsCreateToManyRelationship(ctx, id).GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest(gameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest).Execute()



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
    gameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest := *openapiclient.NewGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest([]openapiclient.GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner{*openapiclient.NewGameCenterDetailRelationshipsGameCenterLeaderboardsDataInner("Type_example", "Id_example")}) // GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGameCenterLeaderboardsCreateToManyRelationship(context.Background(), id).GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest(gameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGameCenterLeaderboardsCreateToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetsGameCenterLeaderboardsCreateToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest** | [**GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest**](GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest.md) | List of related linkages | 

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


## GameCenterLeaderboardSetsGameCenterLeaderboardsDeleteToManyRelationship

> GameCenterLeaderboardSetsGameCenterLeaderboardsDeleteToManyRelationship(ctx, id).GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest(gameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest).Execute()



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
    gameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest := *openapiclient.NewGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest([]openapiclient.GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner{*openapiclient.NewGameCenterDetailRelationshipsGameCenterLeaderboardsDataInner("Type_example", "Id_example")}) // GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGameCenterLeaderboardsDeleteToManyRelationship(context.Background(), id).GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest(gameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGameCenterLeaderboardsDeleteToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetsGameCenterLeaderboardsDeleteToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest** | [**GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest**](GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest.md) | List of related linkages | 

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


## GameCenterLeaderboardSetsGameCenterLeaderboardsGetToManyRelated

> GameCenterLeaderboardsResponse GameCenterLeaderboardSetsGameCenterLeaderboardsGetToManyRelated(ctx, id).FilterArchived(filterArchived).FilterReferenceName(filterReferenceName).FilterId(filterId).FieldsGameCenterLeaderboardLocalizations(fieldsGameCenterLeaderboardLocalizations).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterLeaderboardReleases(fieldsGameCenterLeaderboardReleases).Limit(limit).LimitGameCenterLeaderboardSets(limitGameCenterLeaderboardSets).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Include(include).Execute()



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
    filterArchived := []string{"Inner_example"} // []string | filter by attribute 'archived' (optional)
    filterReferenceName := []string{"Inner_example"} // []string | filter by attribute 'referenceName' (optional)
    filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
    fieldsGameCenterLeaderboardLocalizations := []string{"FieldsGameCenterLeaderboardLocalizations_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardLocalizations (optional)
    fieldsGameCenterLeaderboardSets := []string{"FieldsGameCenterLeaderboardSets_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSets (optional)
    fieldsGameCenterGroups := []string{"FieldsGameCenterGroups_example"} // []string | the fields to include for returned resources of type gameCenterGroups (optional)
    fieldsGameCenterDetails := []string{"FieldsGameCenterDetails_example"} // []string | the fields to include for returned resources of type gameCenterDetails (optional)
    fieldsGameCenterLeaderboards := []string{"FieldsGameCenterLeaderboards_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboards (optional)
    fieldsGameCenterLeaderboardReleases := []string{"FieldsGameCenterLeaderboardReleases_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardReleases (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    limitGameCenterLeaderboardSets := int32(56) // int32 | maximum number of related gameCenterLeaderboardSets returned (when they are included) (optional)
    limitLocalizations := int32(56) // int32 | maximum number of related localizations returned (when they are included) (optional)
    limitReleases := int32(56) // int32 | maximum number of related releases returned (when they are included) (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGameCenterLeaderboardsGetToManyRelated(context.Background(), id).FilterArchived(filterArchived).FilterReferenceName(filterReferenceName).FilterId(filterId).FieldsGameCenterLeaderboardLocalizations(fieldsGameCenterLeaderboardLocalizations).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).FieldsGameCenterLeaderboardReleases(fieldsGameCenterLeaderboardReleases).Limit(limit).LimitGameCenterLeaderboardSets(limitGameCenterLeaderboardSets).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGameCenterLeaderboardsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardSetsGameCenterLeaderboardsGetToManyRelated`: GameCenterLeaderboardsResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGameCenterLeaderboardsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetsGameCenterLeaderboardsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterArchived** | **[]string** | filter by attribute &#39;archived&#39; | 
 **filterReferenceName** | **[]string** | filter by attribute &#39;referenceName&#39; | 
 **filterId** | **[]string** | filter by id(s) | 
 **fieldsGameCenterLeaderboardLocalizations** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardLocalizations | 
 **fieldsGameCenterLeaderboardSets** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSets | 
 **fieldsGameCenterGroups** | **[]string** | the fields to include for returned resources of type gameCenterGroups | 
 **fieldsGameCenterDetails** | **[]string** | the fields to include for returned resources of type gameCenterDetails | 
 **fieldsGameCenterLeaderboards** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboards | 
 **fieldsGameCenterLeaderboardReleases** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardReleases | 
 **limit** | **int32** | maximum resources per page | 
 **limitGameCenterLeaderboardSets** | **int32** | maximum number of related gameCenterLeaderboardSets returned (when they are included) | 
 **limitLocalizations** | **int32** | maximum number of related localizations returned (when they are included) | 
 **limitReleases** | **int32** | maximum number of related releases returned (when they are included) | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterLeaderboardsResponse**](GameCenterLeaderboardsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardSetsGameCenterLeaderboardsGetToManyRelationship

> GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse GameCenterLeaderboardSetsGameCenterLeaderboardsGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGameCenterLeaderboardsGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGameCenterLeaderboardsGetToManyRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardSetsGameCenterLeaderboardsGetToManyRelationship`: GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGameCenterLeaderboardsGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetsGameCenterLeaderboardsGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse**](GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardSetsGameCenterLeaderboardsReplaceToManyRelationship

> GameCenterLeaderboardSetsGameCenterLeaderboardsReplaceToManyRelationship(ctx, id).GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest(gameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest).Execute()



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
    gameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest := *openapiclient.NewGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest([]openapiclient.GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner{*openapiclient.NewGameCenterDetailRelationshipsGameCenterLeaderboardsDataInner("Type_example", "Id_example")}) // GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGameCenterLeaderboardsReplaceToManyRelationship(context.Background(), id).GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest(gameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGameCenterLeaderboardsReplaceToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetsGameCenterLeaderboardsReplaceToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest** | [**GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest**](GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesRequest.md) | List of related linkages | 

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


## GameCenterLeaderboardSetsGetInstance

> GameCenterLeaderboardSetResponse GameCenterLeaderboardSetsGetInstance(ctx, id).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).Include(include).FieldsGameCenterLeaderboardSetLocalizations(fieldsGameCenterLeaderboardSetLocalizations).FieldsGameCenterLeaderboardSetReleases(fieldsGameCenterLeaderboardSetReleases).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).LimitGameCenterLeaderboards(limitGameCenterLeaderboards).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Execute()



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
    fieldsGameCenterLeaderboardSets := []string{"FieldsGameCenterLeaderboardSets_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSets (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsGameCenterLeaderboardSetLocalizations := []string{"FieldsGameCenterLeaderboardSetLocalizations_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSetLocalizations (optional)
    fieldsGameCenterLeaderboardSetReleases := []string{"FieldsGameCenterLeaderboardSetReleases_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSetReleases (optional)
    fieldsGameCenterLeaderboards := []string{"FieldsGameCenterLeaderboards_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboards (optional)
    limitGameCenterLeaderboards := int32(56) // int32 | maximum number of related gameCenterLeaderboards returned (when they are included) (optional)
    limitLocalizations := int32(56) // int32 | maximum number of related localizations returned (when they are included) (optional)
    limitReleases := int32(56) // int32 | maximum number of related releases returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGetInstance(context.Background(), id).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).Include(include).FieldsGameCenterLeaderboardSetLocalizations(fieldsGameCenterLeaderboardSetLocalizations).FieldsGameCenterLeaderboardSetReleases(fieldsGameCenterLeaderboardSetReleases).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).LimitGameCenterLeaderboards(limitGameCenterLeaderboards).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardSetsGetInstance`: GameCenterLeaderboardSetResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterLeaderboardSets** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSets | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsGameCenterLeaderboardSetLocalizations** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSetLocalizations | 
 **fieldsGameCenterLeaderboardSetReleases** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSetReleases | 
 **fieldsGameCenterLeaderboards** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboards | 
 **limitGameCenterLeaderboards** | **int32** | maximum number of related gameCenterLeaderboards returned (when they are included) | 
 **limitLocalizations** | **int32** | maximum number of related localizations returned (when they are included) | 
 **limitReleases** | **int32** | maximum number of related releases returned (when they are included) | 

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


## GameCenterLeaderboardSetsGroupLeaderboardSetGetToOneRelated

> GameCenterLeaderboardSetResponse GameCenterLeaderboardSetsGroupLeaderboardSetGetToOneRelated(ctx, id).FieldsGameCenterLeaderboardSetLocalizations(fieldsGameCenterLeaderboardSetLocalizations).FieldsGameCenterLeaderboardSetReleases(fieldsGameCenterLeaderboardSetReleases).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).LimitLocalizations(limitLocalizations).LimitGameCenterLeaderboards(limitGameCenterLeaderboards).LimitReleases(limitReleases).Include(include).Execute()



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
    resp, r, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGroupLeaderboardSetGetToOneRelated(context.Background(), id).FieldsGameCenterLeaderboardSetLocalizations(fieldsGameCenterLeaderboardSetLocalizations).FieldsGameCenterLeaderboardSetReleases(fieldsGameCenterLeaderboardSetReleases).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterLeaderboards(fieldsGameCenterLeaderboards).LimitLocalizations(limitLocalizations).LimitGameCenterLeaderboards(limitGameCenterLeaderboards).LimitReleases(limitReleases).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGroupLeaderboardSetGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardSetsGroupLeaderboardSetGetToOneRelated`: GameCenterLeaderboardSetResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGroupLeaderboardSetGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetsGroupLeaderboardSetGetToOneRelatedRequest struct via the builder pattern


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


## GameCenterLeaderboardSetsGroupLeaderboardSetGetToOneRelationship

> GameCenterLeaderboardSetGroupLeaderboardSetLinkageResponse GameCenterLeaderboardSetsGroupLeaderboardSetGetToOneRelationship(ctx, id).Execute()



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
    resp, r, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGroupLeaderboardSetGetToOneRelationship(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGroupLeaderboardSetGetToOneRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardSetsGroupLeaderboardSetGetToOneRelationship`: GameCenterLeaderboardSetGroupLeaderboardSetLinkageResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGroupLeaderboardSetGetToOneRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetsGroupLeaderboardSetGetToOneRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GameCenterLeaderboardSetGroupLeaderboardSetLinkageResponse**](GameCenterLeaderboardSetGroupLeaderboardSetLinkageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardSetsGroupLeaderboardSetUpdateToOneRelationship

> GameCenterLeaderboardSetsGroupLeaderboardSetUpdateToOneRelationship(ctx, id).GameCenterLeaderboardSetGroupLeaderboardSetLinkageRequest(gameCenterLeaderboardSetGroupLeaderboardSetLinkageRequest).Execute()



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
    gameCenterLeaderboardSetGroupLeaderboardSetLinkageRequest := *openapiclient.NewGameCenterLeaderboardSetGroupLeaderboardSetLinkageRequest(*openapiclient.NewGameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner("Type_example", "Id_example")) // GameCenterLeaderboardSetGroupLeaderboardSetLinkageRequest | Related linkage

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGroupLeaderboardSetUpdateToOneRelationship(context.Background(), id).GameCenterLeaderboardSetGroupLeaderboardSetLinkageRequest(gameCenterLeaderboardSetGroupLeaderboardSetLinkageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsGroupLeaderboardSetUpdateToOneRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetsGroupLeaderboardSetUpdateToOneRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterLeaderboardSetGroupLeaderboardSetLinkageRequest** | [**GameCenterLeaderboardSetGroupLeaderboardSetLinkageRequest**](GameCenterLeaderboardSetGroupLeaderboardSetLinkageRequest.md) | Related linkage | 

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


## GameCenterLeaderboardSetsLocalizationsGetToManyRelated

> GameCenterLeaderboardSetLocalizationsResponse GameCenterLeaderboardSetsLocalizationsGetToManyRelated(ctx, id).FieldsGameCenterLeaderboardSetLocalizations(fieldsGameCenterLeaderboardSetLocalizations).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterLeaderboardSetImages(fieldsGameCenterLeaderboardSetImages).Limit(limit).Include(include).Execute()



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
    fieldsGameCenterLeaderboardSets := []string{"FieldsGameCenterLeaderboardSets_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSets (optional)
    fieldsGameCenterLeaderboardSetImages := []string{"FieldsGameCenterLeaderboardSetImages_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSetImages (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsLocalizationsGetToManyRelated(context.Background(), id).FieldsGameCenterLeaderboardSetLocalizations(fieldsGameCenterLeaderboardSetLocalizations).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterLeaderboardSetImages(fieldsGameCenterLeaderboardSetImages).Limit(limit).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsLocalizationsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardSetsLocalizationsGetToManyRelated`: GameCenterLeaderboardSetLocalizationsResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsLocalizationsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetsLocalizationsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterLeaderboardSetLocalizations** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSetLocalizations | 
 **fieldsGameCenterLeaderboardSets** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSets | 
 **fieldsGameCenterLeaderboardSetImages** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSetImages | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterLeaderboardSetLocalizationsResponse**](GameCenterLeaderboardSetLocalizationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardSetsReleasesGetToManyRelated

> GameCenterLeaderboardSetReleasesResponse GameCenterLeaderboardSetsReleasesGetToManyRelated(ctx, id).FilterLive(filterLive).FilterGameCenterDetail(filterGameCenterDetail).FieldsGameCenterLeaderboardSetReleases(fieldsGameCenterLeaderboardSetReleases).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterDetails(fieldsGameCenterDetails).Limit(limit).Include(include).Execute()



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
    fieldsGameCenterLeaderboardSetReleases := []string{"FieldsGameCenterLeaderboardSetReleases_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSetReleases (optional)
    fieldsGameCenterLeaderboardSets := []string{"FieldsGameCenterLeaderboardSets_example"} // []string | the fields to include for returned resources of type gameCenterLeaderboardSets (optional)
    fieldsGameCenterDetails := []string{"FieldsGameCenterDetails_example"} // []string | the fields to include for returned resources of type gameCenterDetails (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsReleasesGetToManyRelated(context.Background(), id).FilterLive(filterLive).FilterGameCenterDetail(filterGameCenterDetail).FieldsGameCenterLeaderboardSetReleases(fieldsGameCenterLeaderboardSetReleases).FieldsGameCenterLeaderboardSets(fieldsGameCenterLeaderboardSets).FieldsGameCenterDetails(fieldsGameCenterDetails).Limit(limit).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsReleasesGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardSetsReleasesGetToManyRelated`: GameCenterLeaderboardSetReleasesResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsReleasesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetsReleasesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterLive** | **[]string** | filter by attribute &#39;live&#39; | 
 **filterGameCenterDetail** | **[]string** | filter by id(s) of related &#39;gameCenterDetail&#39; | 
 **fieldsGameCenterLeaderboardSetReleases** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSetReleases | 
 **fieldsGameCenterLeaderboardSets** | **[]string** | the fields to include for returned resources of type gameCenterLeaderboardSets | 
 **fieldsGameCenterDetails** | **[]string** | the fields to include for returned resources of type gameCenterDetails | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterLeaderboardSetReleasesResponse**](GameCenterLeaderboardSetReleasesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterLeaderboardSetsUpdateInstance

> GameCenterLeaderboardSetResponse GameCenterLeaderboardSetsUpdateInstance(ctx, id).GameCenterLeaderboardSetUpdateRequest(gameCenterLeaderboardSetUpdateRequest).Execute()



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
    gameCenterLeaderboardSetUpdateRequest := *openapiclient.NewGameCenterLeaderboardSetUpdateRequest(*openapiclient.NewGameCenterLeaderboardSetUpdateRequestData("Type_example", "Id_example")) // GameCenterLeaderboardSetUpdateRequest | GameCenterLeaderboardSet representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsUpdateInstance(context.Background(), id).GameCenterLeaderboardSetUpdateRequest(gameCenterLeaderboardSetUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterLeaderboardSetsUpdateInstance`: GameCenterLeaderboardSetResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterLeaderboardSetsAPI.GameCenterLeaderboardSetsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterLeaderboardSetsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterLeaderboardSetUpdateRequest** | [**GameCenterLeaderboardSetUpdateRequest**](GameCenterLeaderboardSetUpdateRequest.md) | GameCenterLeaderboardSet representation | 

### Return type

[**GameCenterLeaderboardSetResponse**](GameCenterLeaderboardSetResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

