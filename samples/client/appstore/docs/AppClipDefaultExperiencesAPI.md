# \AppClipDefaultExperiencesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppClipDefaultExperiencesAppClipAppStoreReviewDetailGetToOneRelated**](AppClipDefaultExperiencesAPI.md#AppClipDefaultExperiencesAppClipAppStoreReviewDetailGetToOneRelated) | **Get** /v1/appClipDefaultExperiences/{id}/appClipAppStoreReviewDetail | 
[**AppClipDefaultExperiencesAppClipDefaultExperienceLocalizationsGetToManyRelated**](AppClipDefaultExperiencesAPI.md#AppClipDefaultExperiencesAppClipDefaultExperienceLocalizationsGetToManyRelated) | **Get** /v1/appClipDefaultExperiences/{id}/appClipDefaultExperienceLocalizations | 
[**AppClipDefaultExperiencesCreateInstance**](AppClipDefaultExperiencesAPI.md#AppClipDefaultExperiencesCreateInstance) | **Post** /v1/appClipDefaultExperiences | 
[**AppClipDefaultExperiencesDeleteInstance**](AppClipDefaultExperiencesAPI.md#AppClipDefaultExperiencesDeleteInstance) | **Delete** /v1/appClipDefaultExperiences/{id} | 
[**AppClipDefaultExperiencesGetInstance**](AppClipDefaultExperiencesAPI.md#AppClipDefaultExperiencesGetInstance) | **Get** /v1/appClipDefaultExperiences/{id} | 
[**AppClipDefaultExperiencesReleaseWithAppStoreVersionGetToOneRelated**](AppClipDefaultExperiencesAPI.md#AppClipDefaultExperiencesReleaseWithAppStoreVersionGetToOneRelated) | **Get** /v1/appClipDefaultExperiences/{id}/releaseWithAppStoreVersion | 
[**AppClipDefaultExperiencesReleaseWithAppStoreVersionGetToOneRelationship**](AppClipDefaultExperiencesAPI.md#AppClipDefaultExperiencesReleaseWithAppStoreVersionGetToOneRelationship) | **Get** /v1/appClipDefaultExperiences/{id}/relationships/releaseWithAppStoreVersion | 
[**AppClipDefaultExperiencesReleaseWithAppStoreVersionUpdateToOneRelationship**](AppClipDefaultExperiencesAPI.md#AppClipDefaultExperiencesReleaseWithAppStoreVersionUpdateToOneRelationship) | **Patch** /v1/appClipDefaultExperiences/{id}/relationships/releaseWithAppStoreVersion | 
[**AppClipDefaultExperiencesUpdateInstance**](AppClipDefaultExperiencesAPI.md#AppClipDefaultExperiencesUpdateInstance) | **Patch** /v1/appClipDefaultExperiences/{id} | 



## AppClipDefaultExperiencesAppClipAppStoreReviewDetailGetToOneRelated

> AppClipAppStoreReviewDetailResponse AppClipDefaultExperiencesAppClipAppStoreReviewDetailGetToOneRelated(ctx, id).FieldsAppClipAppStoreReviewDetails(fieldsAppClipAppStoreReviewDetails).FieldsAppClipDefaultExperiences(fieldsAppClipDefaultExperiences).Include(include).Execute()



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
    fieldsAppClipAppStoreReviewDetails := []string{"FieldsAppClipAppStoreReviewDetails_example"} // []string | the fields to include for returned resources of type appClipAppStoreReviewDetails (optional)
    fieldsAppClipDefaultExperiences := []string{"FieldsAppClipDefaultExperiences_example"} // []string | the fields to include for returned resources of type appClipDefaultExperiences (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesAppClipAppStoreReviewDetailGetToOneRelated(context.Background(), id).FieldsAppClipAppStoreReviewDetails(fieldsAppClipAppStoreReviewDetails).FieldsAppClipDefaultExperiences(fieldsAppClipDefaultExperiences).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesAppClipAppStoreReviewDetailGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppClipDefaultExperiencesAppClipAppStoreReviewDetailGetToOneRelated`: AppClipAppStoreReviewDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesAppClipAppStoreReviewDetailGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppClipDefaultExperiencesAppClipAppStoreReviewDetailGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppClipAppStoreReviewDetails** | **[]string** | the fields to include for returned resources of type appClipAppStoreReviewDetails | 
 **fieldsAppClipDefaultExperiences** | **[]string** | the fields to include for returned resources of type appClipDefaultExperiences | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**AppClipAppStoreReviewDetailResponse**](AppClipAppStoreReviewDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppClipDefaultExperiencesAppClipDefaultExperienceLocalizationsGetToManyRelated

> AppClipDefaultExperienceLocalizationsResponse AppClipDefaultExperiencesAppClipDefaultExperienceLocalizationsGetToManyRelated(ctx, id).FilterLocale(filterLocale).FieldsAppClipDefaultExperiences(fieldsAppClipDefaultExperiences).FieldsAppClipDefaultExperienceLocalizations(fieldsAppClipDefaultExperienceLocalizations).FieldsAppClipHeaderImages(fieldsAppClipHeaderImages).Limit(limit).Include(include).Execute()



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
    filterLocale := []string{"Inner_example"} // []string | filter by attribute 'locale' (optional)
    fieldsAppClipDefaultExperiences := []string{"FieldsAppClipDefaultExperiences_example"} // []string | the fields to include for returned resources of type appClipDefaultExperiences (optional)
    fieldsAppClipDefaultExperienceLocalizations := []string{"FieldsAppClipDefaultExperienceLocalizations_example"} // []string | the fields to include for returned resources of type appClipDefaultExperienceLocalizations (optional)
    fieldsAppClipHeaderImages := []string{"FieldsAppClipHeaderImages_example"} // []string | the fields to include for returned resources of type appClipHeaderImages (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesAppClipDefaultExperienceLocalizationsGetToManyRelated(context.Background(), id).FilterLocale(filterLocale).FieldsAppClipDefaultExperiences(fieldsAppClipDefaultExperiences).FieldsAppClipDefaultExperienceLocalizations(fieldsAppClipDefaultExperienceLocalizations).FieldsAppClipHeaderImages(fieldsAppClipHeaderImages).Limit(limit).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesAppClipDefaultExperienceLocalizationsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppClipDefaultExperiencesAppClipDefaultExperienceLocalizationsGetToManyRelated`: AppClipDefaultExperienceLocalizationsResponse
    fmt.Fprintf(os.Stdout, "Response from `AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesAppClipDefaultExperienceLocalizationsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppClipDefaultExperiencesAppClipDefaultExperienceLocalizationsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterLocale** | **[]string** | filter by attribute &#39;locale&#39; | 
 **fieldsAppClipDefaultExperiences** | **[]string** | the fields to include for returned resources of type appClipDefaultExperiences | 
 **fieldsAppClipDefaultExperienceLocalizations** | **[]string** | the fields to include for returned resources of type appClipDefaultExperienceLocalizations | 
 **fieldsAppClipHeaderImages** | **[]string** | the fields to include for returned resources of type appClipHeaderImages | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**AppClipDefaultExperienceLocalizationsResponse**](AppClipDefaultExperienceLocalizationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppClipDefaultExperiencesCreateInstance

> AppClipDefaultExperienceResponse AppClipDefaultExperiencesCreateInstance(ctx).AppClipDefaultExperienceCreateRequest(appClipDefaultExperienceCreateRequest).Execute()



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
    appClipDefaultExperienceCreateRequest := *openapiclient.NewAppClipDefaultExperienceCreateRequest(*openapiclient.NewAppClipDefaultExperienceCreateRequestData("Type_example", *openapiclient.NewAppClipDefaultExperienceCreateRequestDataRelationships(*openapiclient.NewAppClipAdvancedExperienceCreateRequestDataRelationshipsAppClip(*openapiclient.NewAppClipAdvancedExperienceRelationshipsAppClipData("Type_example", "Id_example"))))) // AppClipDefaultExperienceCreateRequest | AppClipDefaultExperience representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesCreateInstance(context.Background()).AppClipDefaultExperienceCreateRequest(appClipDefaultExperienceCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppClipDefaultExperiencesCreateInstance`: AppClipDefaultExperienceResponse
    fmt.Fprintf(os.Stdout, "Response from `AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppClipDefaultExperiencesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appClipDefaultExperienceCreateRequest** | [**AppClipDefaultExperienceCreateRequest**](AppClipDefaultExperienceCreateRequest.md) | AppClipDefaultExperience representation | 

### Return type

[**AppClipDefaultExperienceResponse**](AppClipDefaultExperienceResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppClipDefaultExperiencesDeleteInstance

> AppClipDefaultExperiencesDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppClipDefaultExperiencesDeleteInstanceRequest struct via the builder pattern


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


## AppClipDefaultExperiencesGetInstance

> AppClipDefaultExperienceResponse AppClipDefaultExperiencesGetInstance(ctx, id).FieldsAppClipDefaultExperiences(fieldsAppClipDefaultExperiences).Include(include).FieldsAppClipAppStoreReviewDetails(fieldsAppClipAppStoreReviewDetails).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsAppClipDefaultExperienceLocalizations(fieldsAppClipDefaultExperienceLocalizations).LimitAppClipDefaultExperienceLocalizations(limitAppClipDefaultExperienceLocalizations).Execute()



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
    fieldsAppClipDefaultExperiences := []string{"FieldsAppClipDefaultExperiences_example"} // []string | the fields to include for returned resources of type appClipDefaultExperiences (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsAppClipAppStoreReviewDetails := []string{"FieldsAppClipAppStoreReviewDetails_example"} // []string | the fields to include for returned resources of type appClipAppStoreReviewDetails (optional)
    fieldsAppStoreVersions := []string{"FieldsAppStoreVersions_example"} // []string | the fields to include for returned resources of type appStoreVersions (optional)
    fieldsAppClipDefaultExperienceLocalizations := []string{"FieldsAppClipDefaultExperienceLocalizations_example"} // []string | the fields to include for returned resources of type appClipDefaultExperienceLocalizations (optional)
    limitAppClipDefaultExperienceLocalizations := int32(56) // int32 | maximum number of related appClipDefaultExperienceLocalizations returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesGetInstance(context.Background(), id).FieldsAppClipDefaultExperiences(fieldsAppClipDefaultExperiences).Include(include).FieldsAppClipAppStoreReviewDetails(fieldsAppClipAppStoreReviewDetails).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsAppClipDefaultExperienceLocalizations(fieldsAppClipDefaultExperienceLocalizations).LimitAppClipDefaultExperienceLocalizations(limitAppClipDefaultExperienceLocalizations).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppClipDefaultExperiencesGetInstance`: AppClipDefaultExperienceResponse
    fmt.Fprintf(os.Stdout, "Response from `AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppClipDefaultExperiencesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppClipDefaultExperiences** | **[]string** | the fields to include for returned resources of type appClipDefaultExperiences | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsAppClipAppStoreReviewDetails** | **[]string** | the fields to include for returned resources of type appClipAppStoreReviewDetails | 
 **fieldsAppStoreVersions** | **[]string** | the fields to include for returned resources of type appStoreVersions | 
 **fieldsAppClipDefaultExperienceLocalizations** | **[]string** | the fields to include for returned resources of type appClipDefaultExperienceLocalizations | 
 **limitAppClipDefaultExperienceLocalizations** | **int32** | maximum number of related appClipDefaultExperienceLocalizations returned (when they are included) | 

### Return type

[**AppClipDefaultExperienceResponse**](AppClipDefaultExperienceResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppClipDefaultExperiencesReleaseWithAppStoreVersionGetToOneRelated

> AppStoreVersionResponse AppClipDefaultExperiencesReleaseWithAppStoreVersionGetToOneRelated(ctx, id).FieldsAppStoreVersionExperiments(fieldsAppStoreVersionExperiments).FieldsAgeRatingDeclarations(fieldsAgeRatingDeclarations).FieldsAppStoreVersionSubmissions(fieldsAppStoreVersionSubmissions).FieldsAppStoreReviewDetails(fieldsAppStoreReviewDetails).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsApps(fieldsApps).FieldsRoutingAppCoverages(fieldsRoutingAppCoverages).FieldsAppClipDefaultExperiences(fieldsAppClipDefaultExperiences).FieldsAppStoreVersionPhasedReleases(fieldsAppStoreVersionPhasedReleases).FieldsBuilds(fieldsBuilds).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).LimitAppStoreVersionLocalizations(limitAppStoreVersionLocalizations).LimitAppStoreVersionExperiments(limitAppStoreVersionExperiments).LimitAppStoreVersionExperimentsV2(limitAppStoreVersionExperimentsV2).Include(include).Execute()



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
    resp, r, err := apiClient.AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesReleaseWithAppStoreVersionGetToOneRelated(context.Background(), id).FieldsAppStoreVersionExperiments(fieldsAppStoreVersionExperiments).FieldsAgeRatingDeclarations(fieldsAgeRatingDeclarations).FieldsAppStoreVersionSubmissions(fieldsAppStoreVersionSubmissions).FieldsAppStoreReviewDetails(fieldsAppStoreReviewDetails).FieldsAppStoreVersions(fieldsAppStoreVersions).FieldsApps(fieldsApps).FieldsRoutingAppCoverages(fieldsRoutingAppCoverages).FieldsAppClipDefaultExperiences(fieldsAppClipDefaultExperiences).FieldsAppStoreVersionPhasedReleases(fieldsAppStoreVersionPhasedReleases).FieldsBuilds(fieldsBuilds).FieldsAppStoreVersionLocalizations(fieldsAppStoreVersionLocalizations).LimitAppStoreVersionLocalizations(limitAppStoreVersionLocalizations).LimitAppStoreVersionExperiments(limitAppStoreVersionExperiments).LimitAppStoreVersionExperimentsV2(limitAppStoreVersionExperimentsV2).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesReleaseWithAppStoreVersionGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppClipDefaultExperiencesReleaseWithAppStoreVersionGetToOneRelated`: AppStoreVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesReleaseWithAppStoreVersionGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppClipDefaultExperiencesReleaseWithAppStoreVersionGetToOneRelatedRequest struct via the builder pattern


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


## AppClipDefaultExperiencesReleaseWithAppStoreVersionGetToOneRelationship

> AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageResponse AppClipDefaultExperiencesReleaseWithAppStoreVersionGetToOneRelationship(ctx, id).Execute()



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
    resp, r, err := apiClient.AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesReleaseWithAppStoreVersionGetToOneRelationship(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesReleaseWithAppStoreVersionGetToOneRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppClipDefaultExperiencesReleaseWithAppStoreVersionGetToOneRelationship`: AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageResponse
    fmt.Fprintf(os.Stdout, "Response from `AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesReleaseWithAppStoreVersionGetToOneRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppClipDefaultExperiencesReleaseWithAppStoreVersionGetToOneRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageResponse**](AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppClipDefaultExperiencesReleaseWithAppStoreVersionUpdateToOneRelationship

> AppClipDefaultExperiencesReleaseWithAppStoreVersionUpdateToOneRelationship(ctx, id).AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest(appClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest).Execute()



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
    appClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest := *openapiclient.NewAppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest(*openapiclient.NewAppClipDefaultExperienceRelationshipsReleaseWithAppStoreVersionData("Type_example", "Id_example")) // AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest | Related linkage

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesReleaseWithAppStoreVersionUpdateToOneRelationship(context.Background(), id).AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest(appClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesReleaseWithAppStoreVersionUpdateToOneRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppClipDefaultExperiencesReleaseWithAppStoreVersionUpdateToOneRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest** | [**AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest**](AppClipDefaultExperienceReleaseWithAppStoreVersionLinkageRequest.md) | Related linkage | 

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


## AppClipDefaultExperiencesUpdateInstance

> AppClipDefaultExperienceResponse AppClipDefaultExperiencesUpdateInstance(ctx, id).AppClipDefaultExperienceUpdateRequest(appClipDefaultExperienceUpdateRequest).Execute()



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
    appClipDefaultExperienceUpdateRequest := *openapiclient.NewAppClipDefaultExperienceUpdateRequest(*openapiclient.NewAppClipDefaultExperienceUpdateRequestData("Type_example", "Id_example")) // AppClipDefaultExperienceUpdateRequest | AppClipDefaultExperience representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesUpdateInstance(context.Background(), id).AppClipDefaultExperienceUpdateRequest(appClipDefaultExperienceUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppClipDefaultExperiencesUpdateInstance`: AppClipDefaultExperienceResponse
    fmt.Fprintf(os.Stdout, "Response from `AppClipDefaultExperiencesAPI.AppClipDefaultExperiencesUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppClipDefaultExperiencesUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appClipDefaultExperienceUpdateRequest** | [**AppClipDefaultExperienceUpdateRequest**](AppClipDefaultExperienceUpdateRequest.md) | AppClipDefaultExperience representation | 

### Return type

[**AppClipDefaultExperienceResponse**](AppClipDefaultExperienceResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

