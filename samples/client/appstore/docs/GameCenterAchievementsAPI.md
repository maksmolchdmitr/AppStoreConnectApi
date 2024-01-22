# \GameCenterAchievementsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterAchievementsCreateInstance**](GameCenterAchievementsAPI.md#GameCenterAchievementsCreateInstance) | **Post** /v1/gameCenterAchievements | 
[**GameCenterAchievementsDeleteInstance**](GameCenterAchievementsAPI.md#GameCenterAchievementsDeleteInstance) | **Delete** /v1/gameCenterAchievements/{id} | 
[**GameCenterAchievementsGetInstance**](GameCenterAchievementsAPI.md#GameCenterAchievementsGetInstance) | **Get** /v1/gameCenterAchievements/{id} | 
[**GameCenterAchievementsGroupAchievementGetToOneRelated**](GameCenterAchievementsAPI.md#GameCenterAchievementsGroupAchievementGetToOneRelated) | **Get** /v1/gameCenterAchievements/{id}/groupAchievement | 
[**GameCenterAchievementsGroupAchievementGetToOneRelationship**](GameCenterAchievementsAPI.md#GameCenterAchievementsGroupAchievementGetToOneRelationship) | **Get** /v1/gameCenterAchievements/{id}/relationships/groupAchievement | 
[**GameCenterAchievementsGroupAchievementUpdateToOneRelationship**](GameCenterAchievementsAPI.md#GameCenterAchievementsGroupAchievementUpdateToOneRelationship) | **Patch** /v1/gameCenterAchievements/{id}/relationships/groupAchievement | 
[**GameCenterAchievementsLocalizationsGetToManyRelated**](GameCenterAchievementsAPI.md#GameCenterAchievementsLocalizationsGetToManyRelated) | **Get** /v1/gameCenterAchievements/{id}/localizations | 
[**GameCenterAchievementsReleasesGetToManyRelated**](GameCenterAchievementsAPI.md#GameCenterAchievementsReleasesGetToManyRelated) | **Get** /v1/gameCenterAchievements/{id}/releases | 
[**GameCenterAchievementsUpdateInstance**](GameCenterAchievementsAPI.md#GameCenterAchievementsUpdateInstance) | **Patch** /v1/gameCenterAchievements/{id} | 



## GameCenterAchievementsCreateInstance

> GameCenterAchievementResponse GameCenterAchievementsCreateInstance(ctx).GameCenterAchievementCreateRequest(gameCenterAchievementCreateRequest).Execute()



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
    gameCenterAchievementCreateRequest := *openapiclient.NewGameCenterAchievementCreateRequest(*openapiclient.NewGameCenterAchievementCreateRequestData("Type_example", *openapiclient.NewGameCenterAchievementCreateRequestDataAttributes("ReferenceName_example", "VendorIdentifier_example", int32(123), false, false))) // GameCenterAchievementCreateRequest | GameCenterAchievement representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterAchievementsAPI.GameCenterAchievementsCreateInstance(context.Background()).GameCenterAchievementCreateRequest(gameCenterAchievementCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAchievementsAPI.GameCenterAchievementsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterAchievementsCreateInstance`: GameCenterAchievementResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterAchievementsAPI.GameCenterAchievementsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterAchievementsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameCenterAchievementCreateRequest** | [**GameCenterAchievementCreateRequest**](GameCenterAchievementCreateRequest.md) | GameCenterAchievement representation | 

### Return type

[**GameCenterAchievementResponse**](GameCenterAchievementResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterAchievementsDeleteInstance

> GameCenterAchievementsDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.GameCenterAchievementsAPI.GameCenterAchievementsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAchievementsAPI.GameCenterAchievementsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterAchievementsDeleteInstanceRequest struct via the builder pattern


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


## GameCenterAchievementsGetInstance

> GameCenterAchievementResponse GameCenterAchievementsGetInstance(ctx, id).FieldsGameCenterAchievements(fieldsGameCenterAchievements).Include(include).FieldsGameCenterAchievementReleases(fieldsGameCenterAchievementReleases).FieldsGameCenterAchievementLocalizations(fieldsGameCenterAchievementLocalizations).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Execute()



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
    fieldsGameCenterAchievements := []string{"FieldsGameCenterAchievements_example"} // []string | the fields to include for returned resources of type gameCenterAchievements (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsGameCenterAchievementReleases := []string{"FieldsGameCenterAchievementReleases_example"} // []string | the fields to include for returned resources of type gameCenterAchievementReleases (optional)
    fieldsGameCenterAchievementLocalizations := []string{"FieldsGameCenterAchievementLocalizations_example"} // []string | the fields to include for returned resources of type gameCenterAchievementLocalizations (optional)
    limitLocalizations := int32(56) // int32 | maximum number of related localizations returned (when they are included) (optional)
    limitReleases := int32(56) // int32 | maximum number of related releases returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterAchievementsAPI.GameCenterAchievementsGetInstance(context.Background(), id).FieldsGameCenterAchievements(fieldsGameCenterAchievements).Include(include).FieldsGameCenterAchievementReleases(fieldsGameCenterAchievementReleases).FieldsGameCenterAchievementLocalizations(fieldsGameCenterAchievementLocalizations).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAchievementsAPI.GameCenterAchievementsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterAchievementsGetInstance`: GameCenterAchievementResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterAchievementsAPI.GameCenterAchievementsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterAchievementsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterAchievements** | **[]string** | the fields to include for returned resources of type gameCenterAchievements | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsGameCenterAchievementReleases** | **[]string** | the fields to include for returned resources of type gameCenterAchievementReleases | 
 **fieldsGameCenterAchievementLocalizations** | **[]string** | the fields to include for returned resources of type gameCenterAchievementLocalizations | 
 **limitLocalizations** | **int32** | maximum number of related localizations returned (when they are included) | 
 **limitReleases** | **int32** | maximum number of related releases returned (when they are included) | 

### Return type

[**GameCenterAchievementResponse**](GameCenterAchievementResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterAchievementsGroupAchievementGetToOneRelated

> GameCenterAchievementResponse GameCenterAchievementsGroupAchievementGetToOneRelated(ctx, id).FieldsGameCenterAchievementReleases(fieldsGameCenterAchievementReleases).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterAchievementLocalizations(fieldsGameCenterAchievementLocalizations).FieldsGameCenterAchievements(fieldsGameCenterAchievements).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Include(include).Execute()



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
    fieldsGameCenterGroups := []string{"FieldsGameCenterGroups_example"} // []string | the fields to include for returned resources of type gameCenterGroups (optional)
    fieldsGameCenterDetails := []string{"FieldsGameCenterDetails_example"} // []string | the fields to include for returned resources of type gameCenterDetails (optional)
    fieldsGameCenterAchievementLocalizations := []string{"FieldsGameCenterAchievementLocalizations_example"} // []string | the fields to include for returned resources of type gameCenterAchievementLocalizations (optional)
    fieldsGameCenterAchievements := []string{"FieldsGameCenterAchievements_example"} // []string | the fields to include for returned resources of type gameCenterAchievements (optional)
    limitLocalizations := int32(56) // int32 | maximum number of related localizations returned (when they are included) (optional)
    limitReleases := int32(56) // int32 | maximum number of related releases returned (when they are included) (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterAchievementsAPI.GameCenterAchievementsGroupAchievementGetToOneRelated(context.Background(), id).FieldsGameCenterAchievementReleases(fieldsGameCenterAchievementReleases).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterAchievementLocalizations(fieldsGameCenterAchievementLocalizations).FieldsGameCenterAchievements(fieldsGameCenterAchievements).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAchievementsAPI.GameCenterAchievementsGroupAchievementGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterAchievementsGroupAchievementGetToOneRelated`: GameCenterAchievementResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterAchievementsAPI.GameCenterAchievementsGroupAchievementGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterAchievementsGroupAchievementGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterAchievementReleases** | **[]string** | the fields to include for returned resources of type gameCenterAchievementReleases | 
 **fieldsGameCenterGroups** | **[]string** | the fields to include for returned resources of type gameCenterGroups | 
 **fieldsGameCenterDetails** | **[]string** | the fields to include for returned resources of type gameCenterDetails | 
 **fieldsGameCenterAchievementLocalizations** | **[]string** | the fields to include for returned resources of type gameCenterAchievementLocalizations | 
 **fieldsGameCenterAchievements** | **[]string** | the fields to include for returned resources of type gameCenterAchievements | 
 **limitLocalizations** | **int32** | maximum number of related localizations returned (when they are included) | 
 **limitReleases** | **int32** | maximum number of related releases returned (when they are included) | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterAchievementResponse**](GameCenterAchievementResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterAchievementsGroupAchievementGetToOneRelationship

> GameCenterAchievementGroupAchievementLinkageResponse GameCenterAchievementsGroupAchievementGetToOneRelationship(ctx, id).Execute()



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
    resp, r, err := apiClient.GameCenterAchievementsAPI.GameCenterAchievementsGroupAchievementGetToOneRelationship(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAchievementsAPI.GameCenterAchievementsGroupAchievementGetToOneRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterAchievementsGroupAchievementGetToOneRelationship`: GameCenterAchievementGroupAchievementLinkageResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterAchievementsAPI.GameCenterAchievementsGroupAchievementGetToOneRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterAchievementsGroupAchievementGetToOneRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GameCenterAchievementGroupAchievementLinkageResponse**](GameCenterAchievementGroupAchievementLinkageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterAchievementsGroupAchievementUpdateToOneRelationship

> GameCenterAchievementsGroupAchievementUpdateToOneRelationship(ctx, id).GameCenterAchievementGroupAchievementLinkageRequest(gameCenterAchievementGroupAchievementLinkageRequest).Execute()



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
    gameCenterAchievementGroupAchievementLinkageRequest := *openapiclient.NewGameCenterAchievementGroupAchievementLinkageRequest(*openapiclient.NewGameCenterAchievementLocalizationRelationshipsGameCenterAchievementData("Type_example", "Id_example")) // GameCenterAchievementGroupAchievementLinkageRequest | Related linkage

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GameCenterAchievementsAPI.GameCenterAchievementsGroupAchievementUpdateToOneRelationship(context.Background(), id).GameCenterAchievementGroupAchievementLinkageRequest(gameCenterAchievementGroupAchievementLinkageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAchievementsAPI.GameCenterAchievementsGroupAchievementUpdateToOneRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterAchievementsGroupAchievementUpdateToOneRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterAchievementGroupAchievementLinkageRequest** | [**GameCenterAchievementGroupAchievementLinkageRequest**](GameCenterAchievementGroupAchievementLinkageRequest.md) | Related linkage | 

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


## GameCenterAchievementsLocalizationsGetToManyRelated

> GameCenterAchievementLocalizationsResponse GameCenterAchievementsLocalizationsGetToManyRelated(ctx, id).FieldsGameCenterAchievementImages(fieldsGameCenterAchievementImages).FieldsGameCenterAchievementLocalizations(fieldsGameCenterAchievementLocalizations).FieldsGameCenterAchievements(fieldsGameCenterAchievements).Limit(limit).Include(include).Execute()



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
    fieldsGameCenterAchievementImages := []string{"FieldsGameCenterAchievementImages_example"} // []string | the fields to include for returned resources of type gameCenterAchievementImages (optional)
    fieldsGameCenterAchievementLocalizations := []string{"FieldsGameCenterAchievementLocalizations_example"} // []string | the fields to include for returned resources of type gameCenterAchievementLocalizations (optional)
    fieldsGameCenterAchievements := []string{"FieldsGameCenterAchievements_example"} // []string | the fields to include for returned resources of type gameCenterAchievements (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterAchievementsAPI.GameCenterAchievementsLocalizationsGetToManyRelated(context.Background(), id).FieldsGameCenterAchievementImages(fieldsGameCenterAchievementImages).FieldsGameCenterAchievementLocalizations(fieldsGameCenterAchievementLocalizations).FieldsGameCenterAchievements(fieldsGameCenterAchievements).Limit(limit).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAchievementsAPI.GameCenterAchievementsLocalizationsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterAchievementsLocalizationsGetToManyRelated`: GameCenterAchievementLocalizationsResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterAchievementsAPI.GameCenterAchievementsLocalizationsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterAchievementsLocalizationsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterAchievementImages** | **[]string** | the fields to include for returned resources of type gameCenterAchievementImages | 
 **fieldsGameCenterAchievementLocalizations** | **[]string** | the fields to include for returned resources of type gameCenterAchievementLocalizations | 
 **fieldsGameCenterAchievements** | **[]string** | the fields to include for returned resources of type gameCenterAchievements | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterAchievementLocalizationsResponse**](GameCenterAchievementLocalizationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterAchievementsReleasesGetToManyRelated

> GameCenterAchievementReleasesResponse GameCenterAchievementsReleasesGetToManyRelated(ctx, id).FilterLive(filterLive).FilterGameCenterDetail(filterGameCenterDetail).FieldsGameCenterAchievementReleases(fieldsGameCenterAchievementReleases).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterAchievements(fieldsGameCenterAchievements).Limit(limit).Include(include).Execute()



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
    fieldsGameCenterAchievementReleases := []string{"FieldsGameCenterAchievementReleases_example"} // []string | the fields to include for returned resources of type gameCenterAchievementReleases (optional)
    fieldsGameCenterDetails := []string{"FieldsGameCenterDetails_example"} // []string | the fields to include for returned resources of type gameCenterDetails (optional)
    fieldsGameCenterAchievements := []string{"FieldsGameCenterAchievements_example"} // []string | the fields to include for returned resources of type gameCenterAchievements (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterAchievementsAPI.GameCenterAchievementsReleasesGetToManyRelated(context.Background(), id).FilterLive(filterLive).FilterGameCenterDetail(filterGameCenterDetail).FieldsGameCenterAchievementReleases(fieldsGameCenterAchievementReleases).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterAchievements(fieldsGameCenterAchievements).Limit(limit).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAchievementsAPI.GameCenterAchievementsReleasesGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterAchievementsReleasesGetToManyRelated`: GameCenterAchievementReleasesResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterAchievementsAPI.GameCenterAchievementsReleasesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterAchievementsReleasesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterLive** | **[]string** | filter by attribute &#39;live&#39; | 
 **filterGameCenterDetail** | **[]string** | filter by id(s) of related &#39;gameCenterDetail&#39; | 
 **fieldsGameCenterAchievementReleases** | **[]string** | the fields to include for returned resources of type gameCenterAchievementReleases | 
 **fieldsGameCenterDetails** | **[]string** | the fields to include for returned resources of type gameCenterDetails | 
 **fieldsGameCenterAchievements** | **[]string** | the fields to include for returned resources of type gameCenterAchievements | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterAchievementReleasesResponse**](GameCenterAchievementReleasesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterAchievementsUpdateInstance

> GameCenterAchievementResponse GameCenterAchievementsUpdateInstance(ctx, id).GameCenterAchievementUpdateRequest(gameCenterAchievementUpdateRequest).Execute()



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
    gameCenterAchievementUpdateRequest := *openapiclient.NewGameCenterAchievementUpdateRequest(*openapiclient.NewGameCenterAchievementUpdateRequestData("Type_example", "Id_example")) // GameCenterAchievementUpdateRequest | GameCenterAchievement representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterAchievementsAPI.GameCenterAchievementsUpdateInstance(context.Background(), id).GameCenterAchievementUpdateRequest(gameCenterAchievementUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAchievementsAPI.GameCenterAchievementsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterAchievementsUpdateInstance`: GameCenterAchievementResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterAchievementsAPI.GameCenterAchievementsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterAchievementsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterAchievementUpdateRequest** | [**GameCenterAchievementUpdateRequest**](GameCenterAchievementUpdateRequest.md) | GameCenterAchievement representation | 

### Return type

[**GameCenterAchievementResponse**](GameCenterAchievementResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

