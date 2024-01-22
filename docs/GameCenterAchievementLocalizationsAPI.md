# \GameCenterAchievementLocalizationsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterAchievementLocalizationsCreateInstance**](GameCenterAchievementLocalizationsAPI.md#GameCenterAchievementLocalizationsCreateInstance) | **Post** /v1/gameCenterAchievementLocalizations | 
[**GameCenterAchievementLocalizationsDeleteInstance**](GameCenterAchievementLocalizationsAPI.md#GameCenterAchievementLocalizationsDeleteInstance) | **Delete** /v1/gameCenterAchievementLocalizations/{id} | 
[**GameCenterAchievementLocalizationsGameCenterAchievementGetToOneRelated**](GameCenterAchievementLocalizationsAPI.md#GameCenterAchievementLocalizationsGameCenterAchievementGetToOneRelated) | **Get** /v1/gameCenterAchievementLocalizations/{id}/gameCenterAchievement | 
[**GameCenterAchievementLocalizationsGameCenterAchievementImageGetToOneRelated**](GameCenterAchievementLocalizationsAPI.md#GameCenterAchievementLocalizationsGameCenterAchievementImageGetToOneRelated) | **Get** /v1/gameCenterAchievementLocalizations/{id}/gameCenterAchievementImage | 
[**GameCenterAchievementLocalizationsGetInstance**](GameCenterAchievementLocalizationsAPI.md#GameCenterAchievementLocalizationsGetInstance) | **Get** /v1/gameCenterAchievementLocalizations/{id} | 
[**GameCenterAchievementLocalizationsUpdateInstance**](GameCenterAchievementLocalizationsAPI.md#GameCenterAchievementLocalizationsUpdateInstance) | **Patch** /v1/gameCenterAchievementLocalizations/{id} | 



## GameCenterAchievementLocalizationsCreateInstance

> GameCenterAchievementLocalizationResponse GameCenterAchievementLocalizationsCreateInstance(ctx).GameCenterAchievementLocalizationCreateRequest(gameCenterAchievementLocalizationCreateRequest).Execute()



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
    gameCenterAchievementLocalizationCreateRequest := *openapiclient.NewGameCenterAchievementLocalizationCreateRequest(*openapiclient.NewGameCenterAchievementLocalizationCreateRequestData("Type_example", *openapiclient.NewGameCenterAchievementLocalizationCreateRequestDataAttributes("Locale_example", "Name_example", "BeforeEarnedDescription_example", "AfterEarnedDescription_example"), *openapiclient.NewGameCenterAchievementLocalizationCreateRequestDataRelationships(*openapiclient.NewGameCenterAchievementLocalizationCreateRequestDataRelationshipsGameCenterAchievement(*openapiclient.NewGameCenterAchievementLocalizationRelationshipsGameCenterAchievementData("Type_example", "Id_example"))))) // GameCenterAchievementLocalizationCreateRequest | GameCenterAchievementLocalization representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterAchievementLocalizationsAPI.GameCenterAchievementLocalizationsCreateInstance(context.Background()).GameCenterAchievementLocalizationCreateRequest(gameCenterAchievementLocalizationCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAchievementLocalizationsAPI.GameCenterAchievementLocalizationsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterAchievementLocalizationsCreateInstance`: GameCenterAchievementLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterAchievementLocalizationsAPI.GameCenterAchievementLocalizationsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterAchievementLocalizationsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameCenterAchievementLocalizationCreateRequest** | [**GameCenterAchievementLocalizationCreateRequest**](GameCenterAchievementLocalizationCreateRequest.md) | GameCenterAchievementLocalization representation | 

### Return type

[**GameCenterAchievementLocalizationResponse**](GameCenterAchievementLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterAchievementLocalizationsDeleteInstance

> GameCenterAchievementLocalizationsDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.GameCenterAchievementLocalizationsAPI.GameCenterAchievementLocalizationsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAchievementLocalizationsAPI.GameCenterAchievementLocalizationsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterAchievementLocalizationsDeleteInstanceRequest struct via the builder pattern


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


## GameCenterAchievementLocalizationsGameCenterAchievementGetToOneRelated

> GameCenterAchievementResponse GameCenterAchievementLocalizationsGameCenterAchievementGetToOneRelated(ctx, id).FieldsGameCenterAchievementReleases(fieldsGameCenterAchievementReleases).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterAchievementLocalizations(fieldsGameCenterAchievementLocalizations).FieldsGameCenterAchievements(fieldsGameCenterAchievements).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Include(include).Execute()



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
    resp, r, err := apiClient.GameCenterAchievementLocalizationsAPI.GameCenterAchievementLocalizationsGameCenterAchievementGetToOneRelated(context.Background(), id).FieldsGameCenterAchievementReleases(fieldsGameCenterAchievementReleases).FieldsGameCenterGroups(fieldsGameCenterGroups).FieldsGameCenterDetails(fieldsGameCenterDetails).FieldsGameCenterAchievementLocalizations(fieldsGameCenterAchievementLocalizations).FieldsGameCenterAchievements(fieldsGameCenterAchievements).LimitLocalizations(limitLocalizations).LimitReleases(limitReleases).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAchievementLocalizationsAPI.GameCenterAchievementLocalizationsGameCenterAchievementGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterAchievementLocalizationsGameCenterAchievementGetToOneRelated`: GameCenterAchievementResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterAchievementLocalizationsAPI.GameCenterAchievementLocalizationsGameCenterAchievementGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterAchievementLocalizationsGameCenterAchievementGetToOneRelatedRequest struct via the builder pattern


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


## GameCenterAchievementLocalizationsGameCenterAchievementImageGetToOneRelated

> GameCenterAchievementImageResponse GameCenterAchievementLocalizationsGameCenterAchievementImageGetToOneRelated(ctx, id).FieldsGameCenterAchievementImages(fieldsGameCenterAchievementImages).FieldsGameCenterAchievementLocalizations(fieldsGameCenterAchievementLocalizations).Include(include).Execute()



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
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterAchievementLocalizationsAPI.GameCenterAchievementLocalizationsGameCenterAchievementImageGetToOneRelated(context.Background(), id).FieldsGameCenterAchievementImages(fieldsGameCenterAchievementImages).FieldsGameCenterAchievementLocalizations(fieldsGameCenterAchievementLocalizations).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAchievementLocalizationsAPI.GameCenterAchievementLocalizationsGameCenterAchievementImageGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterAchievementLocalizationsGameCenterAchievementImageGetToOneRelated`: GameCenterAchievementImageResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterAchievementLocalizationsAPI.GameCenterAchievementLocalizationsGameCenterAchievementImageGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterAchievementLocalizationsGameCenterAchievementImageGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterAchievementImages** | **[]string** | the fields to include for returned resources of type gameCenterAchievementImages | 
 **fieldsGameCenterAchievementLocalizations** | **[]string** | the fields to include for returned resources of type gameCenterAchievementLocalizations | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterAchievementImageResponse**](GameCenterAchievementImageResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterAchievementLocalizationsGetInstance

> GameCenterAchievementLocalizationResponse GameCenterAchievementLocalizationsGetInstance(ctx, id).FieldsGameCenterAchievementLocalizations(fieldsGameCenterAchievementLocalizations).Include(include).FieldsGameCenterAchievementImages(fieldsGameCenterAchievementImages).FieldsGameCenterAchievements(fieldsGameCenterAchievements).Execute()



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
    fieldsGameCenterAchievementLocalizations := []string{"FieldsGameCenterAchievementLocalizations_example"} // []string | the fields to include for returned resources of type gameCenterAchievementLocalizations (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsGameCenterAchievementImages := []string{"FieldsGameCenterAchievementImages_example"} // []string | the fields to include for returned resources of type gameCenterAchievementImages (optional)
    fieldsGameCenterAchievements := []string{"FieldsGameCenterAchievements_example"} // []string | the fields to include for returned resources of type gameCenterAchievements (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterAchievementLocalizationsAPI.GameCenterAchievementLocalizationsGetInstance(context.Background(), id).FieldsGameCenterAchievementLocalizations(fieldsGameCenterAchievementLocalizations).Include(include).FieldsGameCenterAchievementImages(fieldsGameCenterAchievementImages).FieldsGameCenterAchievements(fieldsGameCenterAchievements).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAchievementLocalizationsAPI.GameCenterAchievementLocalizationsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterAchievementLocalizationsGetInstance`: GameCenterAchievementLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterAchievementLocalizationsAPI.GameCenterAchievementLocalizationsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterAchievementLocalizationsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterAchievementLocalizations** | **[]string** | the fields to include for returned resources of type gameCenterAchievementLocalizations | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsGameCenterAchievementImages** | **[]string** | the fields to include for returned resources of type gameCenterAchievementImages | 
 **fieldsGameCenterAchievements** | **[]string** | the fields to include for returned resources of type gameCenterAchievements | 

### Return type

[**GameCenterAchievementLocalizationResponse**](GameCenterAchievementLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterAchievementLocalizationsUpdateInstance

> GameCenterAchievementLocalizationResponse GameCenterAchievementLocalizationsUpdateInstance(ctx, id).GameCenterAchievementLocalizationUpdateRequest(gameCenterAchievementLocalizationUpdateRequest).Execute()



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
    gameCenterAchievementLocalizationUpdateRequest := *openapiclient.NewGameCenterAchievementLocalizationUpdateRequest(*openapiclient.NewGameCenterAchievementLocalizationUpdateRequestData("Type_example", "Id_example")) // GameCenterAchievementLocalizationUpdateRequest | GameCenterAchievementLocalization representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterAchievementLocalizationsAPI.GameCenterAchievementLocalizationsUpdateInstance(context.Background(), id).GameCenterAchievementLocalizationUpdateRequest(gameCenterAchievementLocalizationUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterAchievementLocalizationsAPI.GameCenterAchievementLocalizationsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterAchievementLocalizationsUpdateInstance`: GameCenterAchievementLocalizationResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterAchievementLocalizationsAPI.GameCenterAchievementLocalizationsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterAchievementLocalizationsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterAchievementLocalizationUpdateRequest** | [**GameCenterAchievementLocalizationUpdateRequest**](GameCenterAchievementLocalizationUpdateRequest.md) | GameCenterAchievementLocalization representation | 

### Return type

[**GameCenterAchievementLocalizationResponse**](GameCenterAchievementLocalizationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

