# \AppEncryptionDeclarationsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppEncryptionDeclarationsAppEncryptionDeclarationDocumentGetToOneRelated**](AppEncryptionDeclarationsAPI.md#AppEncryptionDeclarationsAppEncryptionDeclarationDocumentGetToOneRelated) | **Get** /v1/appEncryptionDeclarations/{id}/appEncryptionDeclarationDocument | 
[**AppEncryptionDeclarationsAppGetToOneRelated**](AppEncryptionDeclarationsAPI.md#AppEncryptionDeclarationsAppGetToOneRelated) | **Get** /v1/appEncryptionDeclarations/{id}/app | 
[**AppEncryptionDeclarationsBuildsCreateToManyRelationship**](AppEncryptionDeclarationsAPI.md#AppEncryptionDeclarationsBuildsCreateToManyRelationship) | **Post** /v1/appEncryptionDeclarations/{id}/relationships/builds | 
[**AppEncryptionDeclarationsGetCollection**](AppEncryptionDeclarationsAPI.md#AppEncryptionDeclarationsGetCollection) | **Get** /v1/appEncryptionDeclarations | 
[**AppEncryptionDeclarationsGetInstance**](AppEncryptionDeclarationsAPI.md#AppEncryptionDeclarationsGetInstance) | **Get** /v1/appEncryptionDeclarations/{id} | 



## AppEncryptionDeclarationsAppEncryptionDeclarationDocumentGetToOneRelated

> AppEncryptionDeclarationDocumentResponse AppEncryptionDeclarationsAppEncryptionDeclarationDocumentGetToOneRelated(ctx, id).FieldsAppEncryptionDeclarationDocuments(fieldsAppEncryptionDeclarationDocuments).Execute()



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
    fieldsAppEncryptionDeclarationDocuments := []string{"FieldsAppEncryptionDeclarationDocuments_example"} // []string | the fields to include for returned resources of type appEncryptionDeclarationDocuments (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppEncryptionDeclarationsAPI.AppEncryptionDeclarationsAppEncryptionDeclarationDocumentGetToOneRelated(context.Background(), id).FieldsAppEncryptionDeclarationDocuments(fieldsAppEncryptionDeclarationDocuments).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppEncryptionDeclarationsAPI.AppEncryptionDeclarationsAppEncryptionDeclarationDocumentGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppEncryptionDeclarationsAppEncryptionDeclarationDocumentGetToOneRelated`: AppEncryptionDeclarationDocumentResponse
    fmt.Fprintf(os.Stdout, "Response from `AppEncryptionDeclarationsAPI.AppEncryptionDeclarationsAppEncryptionDeclarationDocumentGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppEncryptionDeclarationsAppEncryptionDeclarationDocumentGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppEncryptionDeclarationDocuments** | **[]string** | the fields to include for returned resources of type appEncryptionDeclarationDocuments | 

### Return type

[**AppEncryptionDeclarationDocumentResponse**](AppEncryptionDeclarationDocumentResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppEncryptionDeclarationsAppGetToOneRelated

> AppWithoutIncludesResponse AppEncryptionDeclarationsAppGetToOneRelated(ctx, id).FieldsApps(fieldsApps).Execute()



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
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppEncryptionDeclarationsAPI.AppEncryptionDeclarationsAppGetToOneRelated(context.Background(), id).FieldsApps(fieldsApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppEncryptionDeclarationsAPI.AppEncryptionDeclarationsAppGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppEncryptionDeclarationsAppGetToOneRelated`: AppWithoutIncludesResponse
    fmt.Fprintf(os.Stdout, "Response from `AppEncryptionDeclarationsAPI.AppEncryptionDeclarationsAppGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppEncryptionDeclarationsAppGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 

### Return type

[**AppWithoutIncludesResponse**](AppWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppEncryptionDeclarationsBuildsCreateToManyRelationship

> AppEncryptionDeclarationsBuildsCreateToManyRelationship(ctx, id).AppEncryptionDeclarationBuildsLinkagesRequest(appEncryptionDeclarationBuildsLinkagesRequest).Execute()



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
    appEncryptionDeclarationBuildsLinkagesRequest := *openapiclient.NewAppEncryptionDeclarationBuildsLinkagesRequest([]openapiclient.AppEncryptionDeclarationRelationshipsBuildsDataInner{*openapiclient.NewAppEncryptionDeclarationRelationshipsBuildsDataInner("Type_example", "Id_example")}) // AppEncryptionDeclarationBuildsLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AppEncryptionDeclarationsAPI.AppEncryptionDeclarationsBuildsCreateToManyRelationship(context.Background(), id).AppEncryptionDeclarationBuildsLinkagesRequest(appEncryptionDeclarationBuildsLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppEncryptionDeclarationsAPI.AppEncryptionDeclarationsBuildsCreateToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAppEncryptionDeclarationsBuildsCreateToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appEncryptionDeclarationBuildsLinkagesRequest** | [**AppEncryptionDeclarationBuildsLinkagesRequest**](AppEncryptionDeclarationBuildsLinkagesRequest.md) | List of related linkages | 

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


## AppEncryptionDeclarationsGetCollection

> AppEncryptionDeclarationsResponse AppEncryptionDeclarationsGetCollection(ctx).FilterPlatform(filterPlatform).FilterApp(filterApp).FilterBuilds(filterBuilds).FieldsAppEncryptionDeclarations(fieldsAppEncryptionDeclarations).Limit(limit).Include(include).FieldsAppEncryptionDeclarationDocuments(fieldsAppEncryptionDeclarationDocuments).FieldsApps(fieldsApps).LimitBuilds(limitBuilds).Execute()



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
    filterPlatform := []string{"FilterPlatform_example"} // []string | filter by attribute 'platform' (optional)
    filterApp := []string{"Inner_example"} // []string | filter by id(s) of related 'app' (optional)
    filterBuilds := []string{"Inner_example"} // []string | filter by id(s) of related 'builds' (optional)
    fieldsAppEncryptionDeclarations := []string{"FieldsAppEncryptionDeclarations_example"} // []string | the fields to include for returned resources of type appEncryptionDeclarations (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsAppEncryptionDeclarationDocuments := []string{"FieldsAppEncryptionDeclarationDocuments_example"} // []string | the fields to include for returned resources of type appEncryptionDeclarationDocuments (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limitBuilds := int32(56) // int32 | maximum number of related builds returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppEncryptionDeclarationsAPI.AppEncryptionDeclarationsGetCollection(context.Background()).FilterPlatform(filterPlatform).FilterApp(filterApp).FilterBuilds(filterBuilds).FieldsAppEncryptionDeclarations(fieldsAppEncryptionDeclarations).Limit(limit).Include(include).FieldsAppEncryptionDeclarationDocuments(fieldsAppEncryptionDeclarationDocuments).FieldsApps(fieldsApps).LimitBuilds(limitBuilds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppEncryptionDeclarationsAPI.AppEncryptionDeclarationsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppEncryptionDeclarationsGetCollection`: AppEncryptionDeclarationsResponse
    fmt.Fprintf(os.Stdout, "Response from `AppEncryptionDeclarationsAPI.AppEncryptionDeclarationsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppEncryptionDeclarationsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterPlatform** | **[]string** | filter by attribute &#39;platform&#39; | 
 **filterApp** | **[]string** | filter by id(s) of related &#39;app&#39; | 
 **filterBuilds** | **[]string** | filter by id(s) of related &#39;builds&#39; | 
 **fieldsAppEncryptionDeclarations** | **[]string** | the fields to include for returned resources of type appEncryptionDeclarations | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsAppEncryptionDeclarationDocuments** | **[]string** | the fields to include for returned resources of type appEncryptionDeclarationDocuments | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limitBuilds** | **int32** | maximum number of related builds returned (when they are included) | 

### Return type

[**AppEncryptionDeclarationsResponse**](AppEncryptionDeclarationsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppEncryptionDeclarationsGetInstance

> AppEncryptionDeclarationResponse AppEncryptionDeclarationsGetInstance(ctx, id).FieldsAppEncryptionDeclarations(fieldsAppEncryptionDeclarations).Include(include).FieldsAppEncryptionDeclarationDocuments(fieldsAppEncryptionDeclarationDocuments).FieldsApps(fieldsApps).LimitBuilds(limitBuilds).Execute()



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
    fieldsAppEncryptionDeclarations := []string{"FieldsAppEncryptionDeclarations_example"} // []string | the fields to include for returned resources of type appEncryptionDeclarations (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsAppEncryptionDeclarationDocuments := []string{"FieldsAppEncryptionDeclarationDocuments_example"} // []string | the fields to include for returned resources of type appEncryptionDeclarationDocuments (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limitBuilds := int32(56) // int32 | maximum number of related builds returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppEncryptionDeclarationsAPI.AppEncryptionDeclarationsGetInstance(context.Background(), id).FieldsAppEncryptionDeclarations(fieldsAppEncryptionDeclarations).Include(include).FieldsAppEncryptionDeclarationDocuments(fieldsAppEncryptionDeclarationDocuments).FieldsApps(fieldsApps).LimitBuilds(limitBuilds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppEncryptionDeclarationsAPI.AppEncryptionDeclarationsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppEncryptionDeclarationsGetInstance`: AppEncryptionDeclarationResponse
    fmt.Fprintf(os.Stdout, "Response from `AppEncryptionDeclarationsAPI.AppEncryptionDeclarationsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppEncryptionDeclarationsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppEncryptionDeclarations** | **[]string** | the fields to include for returned resources of type appEncryptionDeclarations | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsAppEncryptionDeclarationDocuments** | **[]string** | the fields to include for returned resources of type appEncryptionDeclarationDocuments | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limitBuilds** | **int32** | maximum number of related builds returned (when they are included) | 

### Return type

[**AppEncryptionDeclarationResponse**](AppEncryptionDeclarationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

