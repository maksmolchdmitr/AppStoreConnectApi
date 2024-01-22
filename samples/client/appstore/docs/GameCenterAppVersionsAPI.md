# \GameCenterAppVersionsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterAppVersionsAppStoreVersionGetToOneRelated**](GameCenterAppVersionsAPI.md#GameCenterAppVersionsAppStoreVersionGetToOneRelated) | **Get** /v1/gameCenterAppVersions/{id}/appStoreVersion | 
[**GameCenterAppVersionsCompatibilityVersionsCreateToManyRelationship**](GameCenterAppVersionsAPI.md#GameCenterAppVersionsCompatibilityVersionsCreateToManyRelationship) | **Post** /v1/gameCenterAppVersions/{id}/relationships/compatibilityVersions | 
[**GameCenterAppVersionsCompatibilityVersionsDeleteToManyRelationship**](GameCenterAppVersionsAPI.md#GameCenterAppVersionsCompatibilityVersionsDeleteToManyRelationship) | **Delete** /v1/gameCenterAppVersions/{id}/relationships/compatibilityVersions | 
[**GameCenterAppVersionsCompatibilityVersionsGetToManyRelated**](GameCenterAppVersionsAPI.md#GameCenterAppVersionsCompatibilityVersionsGetToManyRelated) | **Get** /v1/gameCenterAppVersions/{id}/compatibilityVersions | 
[**GameCenterAppVersionsCompatibilityVersionsGetToManyRelationship**](GameCenterAppVersionsAPI.md#GameCenterAppVersionsCompatibilityVersionsGetToManyRelationship) | **Get** /v1/gameCenterAppVersions/{id}/relationships/compatibilityVersions | 
[**GameCenterAppVersionsCreateInstance**](GameCenterAppVersionsAPI.md#GameCenterAppVersionsCreateInstance) | **Post** /v1/gameCenterAppVersions | 
[**GameCenterAppVersionsGetInstance**](GameCenterAppVersionsAPI.md#GameCenterAppVersionsGetInstance) | **Get** /v1/gameCenterAppVersions/{id} | 
[**GameCenterAppVersionsUpdateInstance**](GameCenterAppVersionsAPI.md#GameCenterAppVersionsUpdateInstance) | **Patch** /v1/gameCenterAppVersions/{id} | 



## GameCenterAppVersionsAppStoreVersionGetToOneRelated

> AppStoreVersionResponse GameCenterAppVersionsAppStoreVersionGetToOneRelated(ctx, id).FieldsAppStoreVersionExperiments(fieldsAppStoreVersionExperiments).FieldsAgeRatingDeclarations(fieldsAgeRatingDeclarations).FieldsAppStoreVersionSubmissions(fieldsAppStoreVersionSubmissions).FieldsAppStoreReviewDetails(fieldsAppStoreReviewDetails).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsApps(fieldsApps).FieldsRoutingAppCoverages(fieldsRoutingAppCoverages).FieldsAppClipDefaultExperiences(fieldsAppClipDefaultExperiences).FieldsAppStoreVersionPhasedReleases(fieldsAppStoreVersionPhasedReleases).FieldsBuilds(fieldsBuilds).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).LimitAppStoreVersionLocalizations(limitAppStoreVersionLocalizations).LimitAppStoreVersionExperiments(limitAppStoreVersionExperiments).LimitAppStoreVersionExperimentsV2(limitAppStoreVersionExperimentsV2).Include(include).Execute()



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
    fieldsAppStoreVersionExperiments := []string{"FieldsAppStoreVersionExperiments_example"} // []string | the fields to include for returned resources of type appStoreVersionExperiments (optional)
    fieldsAgeRatingDeclarations := []string{"FieldsAgeRatingDeclarations_example"} // []string | the fields to include for returned resources of type ageRatingDeclarations (optional)
    fieldsAppStoreVersionSubmissions := []string{"FieldsAppStoreVersionSubmissions_example"} // []string | the fields to include for returned resources of type appStoreVersionSubmissions (optional)
    fieldsAppStoreReviewDetails := []string{"FieldsAppStoreReviewDetails_example"} // []string | the fields to include for returned resources of type appStoreReviewDetails (optional)
    fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    fieldsRoutingAppCoverages := []string{"FieldsRoutingAppCoverages_example"} // []string | the fields to include for returned resources of type routingAppCoverages (optional)
    fieldsAppClipDefaultExperiences := []string{"FieldsAppClipDefaultExperiences_example"} // []string | the fields to include for returned resources of type appClipDefaultExperiences (optional)
    fieldsAppStoreVersionPhasedReleases := []string{"FieldsAppStoreVersionPhasedReleases_example"} // []string | the fields to include for returned resources of type appStoreVersionPhasedReleases (optional)
    fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)
    fieldsAppStoreVersionLocalizations := []string{"FieldsAppStoreVersionLocalizations_example"} // []string | the fields to include for returned resources of type appStoreVersionLocalizations (optional)
    limitAppStoreVersionLocalizations := int32(56) // int32 | maximum number of related appStoreVersionLocalizations returned (when they are included) (optional)
    limitAppStoreVersionExperiments := int32(56) // int32 | maximum number of related appStoreVersionExperiments returned (when they are included) (optional)
    limitAppStoreVersionExperimentsV2 := int32(56) // int32 | maximum number of related appStoreVersionExperimentsV2 returned (when they are included) (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterAppVersionsAPI.GameCenterAppVersionsAppStoreVersionGetToOneRelated(context.Background(), id).FieldsAppStoreVersionExperiments(fieldsAppStoreVersionExperiments).FieldsAgeRatingDeclarations(fieldsAgeRatingDeclarations).FieldsAppStoreVersionSubmissions(fieldsAppStoreVersionSubmissions).FieldsAppStoreReviewDetails(fieldsAppStoreReviewDetails).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsApps(fieldsApps).FieldsRoutingAppCoverages(fieldsRoutingAppCoverages).FieldsAppClipDefaultExperiences(fieldsAppClipDefaultExperiences).FieldsAppStoreVersionPhasedReleases(fieldsAppStoreVersionPhasedReleases).FieldsBuilds(fieldsBuilds).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).LimitAppStoreVersionLocalizations(limitAppStoreVersionLocalizations).LimitAppStoreVersionExperiments(limitAppStoreVersionExperiments).LimitAppStoreVersionExperimentsV2(limitAppStoreVersionExperimentsV2).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAppVersionsAPI.GameCenterAppVersionsAppStoreVersionGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterAppVersionsAppStoreVersionGetToOneRelated`: AppStoreVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterAppVersionsAPI.GameCenterAppVersionsAppStoreVersionGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterAppVersionsAppStoreVersionGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppStoreVersionExperiments** | **[]string** | the fields to include for returned resources of type appStoreVersionExperiments | 
 **fieldsAgeRatingDeclarations** | **[]string** | the fields to include for returned resources of type ageRatingDeclarations | 
 **fieldsAppStoreVersionSubmissions** | **[]string** | the fields to include for returned resources of type appStoreVersionSubmissions | 
 **fieldsAppStoreReviewDetails** | **[]string** | the fields to include for returned resources of type appStoreReviewDetails | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **fieldsRoutingAppCoverages** | **[]string** | the fields to include for returned resources of type routingAppCoverages | 
 **fieldsAppClipDefaultExperiences** | **[]string** | the fields to include for returned resources of type appClipDefaultExperiences | 
 **fieldsAppStoreVersionPhasedReleases** | **[]string** | the fields to include for returned resources of type appStoreVersionPhasedReleases | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 
 **fieldsAppStoreVersionLocalizations** | **[]string** | the fields to include for returned resources of type appStoreVersionLocalizations | 
 **limitAppStoreVersionLocalizations** | **int32** | maximum number of related appStoreVersionLocalizations returned (when they are included) | 
 **limitAppStoreVersionExperiments** | **int32** | maximum number of related appStoreVersionExperiments returned (when they are included) | 
 **limitAppStoreVersionExperimentsV2** | **int32** | maximum number of related appStoreVersionExperimentsV2 returned (when they are included) | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**AppStoreVersionResponse**](AppStoreVersionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterAppVersionsCompatibilityVersionsCreateToManyRelationship

> GameCenterAppVersionsCompatibilityVersionsCreateToManyRelationship(ctx, id).GameCenterAppVersionCompatibilityVersionsLinkagesRequest(gameCenterAppVersionCompatibilityVersionsLinkagesRequest).Execute()



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
    gameCenterAppVersionCompatibilityVersionsLinkagesRequest := *openapiclient.NewGameCenterAppVersionCompatibilityVersionsLinkagesRequest([]openapiclient.GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner{*openapiclient.NewGameCenterAppVersionRelationshipsCompatibilityVersionsDataInner("Type_example", "Id_example")}) // GameCenterAppVersionCompatibilityVersionsLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GameCenterAppVersionsAPI.GameCenterAppVersionsCompatibilityVersionsCreateToManyRelationship(context.Background(), id).GameCenterAppVersionCompatibilityVersionsLinkagesRequest(gameCenterAppVersionCompatibilityVersionsLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAppVersionsAPI.GameCenterAppVersionsCompatibilityVersionsCreateToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterAppVersionsCompatibilityVersionsCreateToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterAppVersionCompatibilityVersionsLinkagesRequest** | [**GameCenterAppVersionCompatibilityVersionsLinkagesRequest**](GameCenterAppVersionCompatibilityVersionsLinkagesRequest.md) | List of related linkages | 

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


## GameCenterAppVersionsCompatibilityVersionsDeleteToManyRelationship

> GameCenterAppVersionsCompatibilityVersionsDeleteToManyRelationship(ctx, id).GameCenterAppVersionCompatibilityVersionsLinkagesRequest(gameCenterAppVersionCompatibilityVersionsLinkagesRequest).Execute()



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
    gameCenterAppVersionCompatibilityVersionsLinkagesRequest := *openapiclient.NewGameCenterAppVersionCompatibilityVersionsLinkagesRequest([]openapiclient.GameCenterAppVersionRelationshipsCompatibilityVersionsDataInner{*openapiclient.NewGameCenterAppVersionRelationshipsCompatibilityVersionsDataInner("Type_example", "Id_example")}) // GameCenterAppVersionCompatibilityVersionsLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GameCenterAppVersionsAPI.GameCenterAppVersionsCompatibilityVersionsDeleteToManyRelationship(context.Background(), id).GameCenterAppVersionCompatibilityVersionsLinkagesRequest(gameCenterAppVersionCompatibilityVersionsLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAppVersionsAPI.GameCenterAppVersionsCompatibilityVersionsDeleteToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterAppVersionsCompatibilityVersionsDeleteToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterAppVersionCompatibilityVersionsLinkagesRequest** | [**GameCenterAppVersionCompatibilityVersionsLinkagesRequest**](GameCenterAppVersionCompatibilityVersionsLinkagesRequest.md) | List of related linkages | 

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


## GameCenterAppVersionsCompatibilityVersionsGetToManyRelated

> GameCenterAppVersionsResponse GameCenterAppVersionsCompatibilityVersionsGetToManyRelated(ctx, id).FilterEnabled(filterEnabled).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsGameCenterAppVersions(fieldsGameCenterAppVersions).Limit(limit).LimitCompatibilityVersions(limitCompatibilityVersions).Include(include).Execute()



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
    filterEnabled := []string{"Inner_example"} // []string | filter by attribute 'enabled' (optional)
    fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
    fieldsGameCenterAppVersions := []string{"FieldsGameCenterAppVersions_example"} // []string | the fields to include for returned resources of type gameCenterAppVersions (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    limitCompatibilityVersions := int32(56) // int32 | maximum number of related compatibilityVersions returned (when they are included) (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterAppVersionsAPI.GameCenterAppVersionsCompatibilityVersionsGetToManyRelated(context.Background(), id).FilterEnabled(filterEnabled).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsGameCenterAppVersions(fieldsGameCenterAppVersions).Limit(limit).LimitCompatibilityVersions(limitCompatibilityVersions).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAppVersionsAPI.GameCenterAppVersionsCompatibilityVersionsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterAppVersionsCompatibilityVersionsGetToManyRelated`: GameCenterAppVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterAppVersionsAPI.GameCenterAppVersionsCompatibilityVersionsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterAppVersionsCompatibilityVersionsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterEnabled** | **[]string** | filter by attribute &#39;enabled&#39; | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **fieldsGameCenterAppVersions** | **[]string** | the fields to include for returned resources of type gameCenterAppVersions | 
 **limit** | **int32** | maximum resources per page | 
 **limitCompatibilityVersions** | **int32** | maximum number of related compatibilityVersions returned (when they are included) | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterAppVersionsResponse**](GameCenterAppVersionsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterAppVersionsCompatibilityVersionsGetToManyRelationship

> GameCenterAppVersionCompatibilityVersionsLinkagesResponse GameCenterAppVersionsCompatibilityVersionsGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
    resp, r, err := apiClient.GameCenterAppVersionsAPI.GameCenterAppVersionsCompatibilityVersionsGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAppVersionsAPI.GameCenterAppVersionsCompatibilityVersionsGetToManyRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterAppVersionsCompatibilityVersionsGetToManyRelationship`: GameCenterAppVersionCompatibilityVersionsLinkagesResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterAppVersionsAPI.GameCenterAppVersionsCompatibilityVersionsGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterAppVersionsCompatibilityVersionsGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**GameCenterAppVersionCompatibilityVersionsLinkagesResponse**](GameCenterAppVersionCompatibilityVersionsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterAppVersionsCreateInstance

> GameCenterAppVersionResponse GameCenterAppVersionsCreateInstance(ctx).GameCenterAppVersionCreateRequest(gameCenterAppVersionCreateRequest).Execute()



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
    gameCenterAppVersionCreateRequest := *openapiclient.NewGameCenterAppVersionCreateRequest(*openapiclient.NewGameCenterAppVersionCreateRequestData("Type_example", *openapiclient.NewAppStoreReviewDetailCreateRequestDataRelationships(*openapiclient.NewAppStoreReviewDetailCreateRequestDataRelationshipsAppStoreVersion(*openapiclient.NewAppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersionData("Type_example", "Id_example"))))) // GameCenterAppVersionCreateRequest | GameCenterAppVersion representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterAppVersionsAPI.GameCenterAppVersionsCreateInstance(context.Background()).GameCenterAppVersionCreateRequest(gameCenterAppVersionCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAppVersionsAPI.GameCenterAppVersionsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterAppVersionsCreateInstance`: GameCenterAppVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterAppVersionsAPI.GameCenterAppVersionsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterAppVersionsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameCenterAppVersionCreateRequest** | [**GameCenterAppVersionCreateRequest**](GameCenterAppVersionCreateRequest.md) | GameCenterAppVersion representation | 

### Return type

[**GameCenterAppVersionResponse**](GameCenterAppVersionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterAppVersionsGetInstance

> GameCenterAppVersionResponse GameCenterAppVersionsGetInstance(ctx, id).FieldsGameCenterAppVersions(fieldsGameCenterAppVersions).Include(include).FieldsAppStoreVersions(fieldsAppStoreVersions).LimitCompatibilityVersions(limitCompatibilityVersions).Execute()



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
    fieldsGameCenterAppVersions := []string{"FieldsGameCenterAppVersions_example"} // []string | the fields to include for returned resources of type gameCenterAppVersions (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
    limitCompatibilityVersions := int32(56) // int32 | maximum number of related compatibilityVersions returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterAppVersionsAPI.GameCenterAppVersionsGetInstance(context.Background(), id).FieldsGameCenterAppVersions(fieldsGameCenterAppVersions).Include(include).FieldsAppStoreVersions(fieldsAppStoreVersions).LimitCompatibilityVersions(limitCompatibilityVersions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAppVersionsAPI.GameCenterAppVersionsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterAppVersionsGetInstance`: GameCenterAppVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterAppVersionsAPI.GameCenterAppVersionsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterAppVersionsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterAppVersions** | **[]string** | the fields to include for returned resources of type gameCenterAppVersions | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **limitCompatibilityVersions** | **int32** | maximum number of related compatibilityVersions returned (when they are included) | 

### Return type

[**GameCenterAppVersionResponse**](GameCenterAppVersionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterAppVersionsUpdateInstance

> GameCenterAppVersionResponse GameCenterAppVersionsUpdateInstance(ctx, id).GameCenterAppVersionUpdateRequest(gameCenterAppVersionUpdateRequest).Execute()



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
    gameCenterAppVersionUpdateRequest := *openapiclient.NewGameCenterAppVersionUpdateRequest(*openapiclient.NewGameCenterAppVersionUpdateRequestData("Type_example", "Id_example")) // GameCenterAppVersionUpdateRequest | GameCenterAppVersion representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterAppVersionsAPI.GameCenterAppVersionsUpdateInstance(context.Background(), id).GameCenterAppVersionUpdateRequest(gameCenterAppVersionUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAppVersionsAPI.GameCenterAppVersionsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterAppVersionsUpdateInstance`: GameCenterAppVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterAppVersionsAPI.GameCenterAppVersionsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterAppVersionsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterAppVersionUpdateRequest** | [**GameCenterAppVersionUpdateRequest**](GameCenterAppVersionUpdateRequest.md) | GameCenterAppVersion representation | 

### Return type

[**GameCenterAppVersionResponse**](GameCenterAppVersionResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

