# \BundleIdsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BundleIdsAppGetToOneRelated**](BundleIdsAPI.md#BundleIdsAppGetToOneRelated) | **Get** /v1/bundleIds/{id}/app | 
[**BundleIdsBundleIdCapabilitiesGetToManyRelated**](BundleIdsAPI.md#BundleIdsBundleIdCapabilitiesGetToManyRelated) | **Get** /v1/bundleIds/{id}/bundleIdCapabilities | 
[**BundleIdsCreateInstance**](BundleIdsAPI.md#BundleIdsCreateInstance) | **Post** /v1/bundleIds | 
[**BundleIdsDeleteInstance**](BundleIdsAPI.md#BundleIdsDeleteInstance) | **Delete** /v1/bundleIds/{id} | 
[**BundleIdsGetCollection**](BundleIdsAPI.md#BundleIdsGetCollection) | **Get** /v1/bundleIds | 
[**BundleIdsGetInstance**](BundleIdsAPI.md#BundleIdsGetInstance) | **Get** /v1/bundleIds/{id} | 
[**BundleIdsProfilesGetToManyRelated**](BundleIdsAPI.md#BundleIdsProfilesGetToManyRelated) | **Get** /v1/bundleIds/{id}/profiles | 
[**BundleIdsUpdateInstance**](BundleIdsAPI.md#BundleIdsUpdateInstance) | **Patch** /v1/bundleIds/{id} | 



## BundleIdsAppGetToOneRelated

> AppWithoutIncludesResponse BundleIdsAppGetToOneRelated(ctx, id).FieldsApps(fieldsApps).Execute()



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
    resp, r, err := apiClient.BundleIdsAPI.BundleIdsAppGetToOneRelated(context.Background(), id).FieldsApps(fieldsApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundleIdsAPI.BundleIdsAppGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BundleIdsAppGetToOneRelated`: AppWithoutIncludesResponse
    fmt.Fprintf(os.Stdout, "Response from `BundleIdsAPI.BundleIdsAppGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBundleIdsAppGetToOneRelatedRequest struct via the builder pattern


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


## BundleIdsBundleIdCapabilitiesGetToManyRelated

> BundleIdCapabilitiesWithoutIncludesResponse BundleIdsBundleIdCapabilitiesGetToManyRelated(ctx, id).FieldsBundleIdCapabilities(fieldsBundleIdCapabilities).Limit(limit).Execute()



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
    fieldsBundleIdCapabilities := []string{"FieldsBundleIdCapabilities_example"} // []string | the fields to include for returned resources of type bundleIdCapabilities (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BundleIdsAPI.BundleIdsBundleIdCapabilitiesGetToManyRelated(context.Background(), id).FieldsBundleIdCapabilities(fieldsBundleIdCapabilities).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundleIdsAPI.BundleIdsBundleIdCapabilitiesGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BundleIdsBundleIdCapabilitiesGetToManyRelated`: BundleIdCapabilitiesWithoutIncludesResponse
    fmt.Fprintf(os.Stdout, "Response from `BundleIdsAPI.BundleIdsBundleIdCapabilitiesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBundleIdsBundleIdCapabilitiesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBundleIdCapabilities** | **[]string** | the fields to include for returned resources of type bundleIdCapabilities | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**BundleIdCapabilitiesWithoutIncludesResponse**](BundleIdCapabilitiesWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BundleIdsCreateInstance

> BundleIdResponse BundleIdsCreateInstance(ctx).BundleIdCreateRequest(bundleIdCreateRequest).Execute()



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
    bundleIdCreateRequest := *openapiclient.NewBundleIdCreateRequest(*openapiclient.NewBundleIdCreateRequestData("Type_example", *openapiclient.NewBundleIdCreateRequestDataAttributes("Name_example", openapiclient.BundleIdPlatform("IOS"), "Identifier_example"))) // BundleIdCreateRequest | BundleId representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BundleIdsAPI.BundleIdsCreateInstance(context.Background()).BundleIdCreateRequest(bundleIdCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundleIdsAPI.BundleIdsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BundleIdsCreateInstance`: BundleIdResponse
    fmt.Fprintf(os.Stdout, "Response from `BundleIdsAPI.BundleIdsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBundleIdsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bundleIdCreateRequest** | [**BundleIdCreateRequest**](BundleIdCreateRequest.md) | BundleId representation | 

### Return type

[**BundleIdResponse**](BundleIdResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BundleIdsDeleteInstance

> BundleIdsDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.BundleIdsAPI.BundleIdsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundleIdsAPI.BundleIdsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBundleIdsDeleteInstanceRequest struct via the builder pattern


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


## BundleIdsGetCollection

> BundleIdsResponse BundleIdsGetCollection(ctx).FilterIdentifier(filterIdentifier).FilterName(filterName).FilterPlatform(filterPlatform).FilterSeedId(filterSeedId).FilterId(filterId).Sort(sort).FieldsBundleIds(fieldsBundleIds).Limit(limit).Include(include).FieldsBundleIdCapabilities(fieldsBundleIdCapabilities).FieldsProfiles(fieldsProfiles).FieldsApps(fieldsApps).LimitBundleIdCapabilities(limitBundleIdCapabilities).LimitProfiles(limitProfiles).Execute()



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
    filterIdentifier := []string{"Inner_example"} // []string | filter by attribute 'identifier' (optional)
    filterName := []string{"Inner_example"} // []string | filter by attribute 'name' (optional)
    filterPlatform := []string{"FilterPlatform_example"} // []string | filter by attribute 'platform' (optional)
    filterSeedId := []string{"Inner_example"} // []string | filter by attribute 'seedId' (optional)
    filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; resources will be sorted as specified (optional)
    fieldsBundleIds := []string{"FieldsBundleIds_example"} // []string | the fields to include for returned resources of type bundleIds (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsBundleIdCapabilities := []string{"FieldsBundleIdCapabilities_example"} // []string | the fields to include for returned resources of type bundleIdCapabilities (optional)
    fieldsProfiles := []string{"FieldsProfiles_example"} // []string | the fields to include for returned resources of type profiles (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limitBundleIdCapabilities := int32(56) // int32 | maximum number of related bundleIdCapabilities returned (when they are included) (optional)
    limitProfiles := int32(56) // int32 | maximum number of related profiles returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BundleIdsAPI.BundleIdsGetCollection(context.Background()).FilterIdentifier(filterIdentifier).FilterName(filterName).FilterPlatform(filterPlatform).FilterSeedId(filterSeedId).FilterId(filterId).Sort(sort).FieldsBundleIds(fieldsBundleIds).Limit(limit).Include(include).FieldsBundleIdCapabilities(fieldsBundleIdCapabilities).FieldsProfiles(fieldsProfiles).FieldsApps(fieldsApps).LimitBundleIdCapabilities(limitBundleIdCapabilities).LimitProfiles(limitProfiles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundleIdsAPI.BundleIdsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BundleIdsGetCollection`: BundleIdsResponse
    fmt.Fprintf(os.Stdout, "Response from `BundleIdsAPI.BundleIdsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBundleIdsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterIdentifier** | **[]string** | filter by attribute &#39;identifier&#39; | 
 **filterName** | **[]string** | filter by attribute &#39;name&#39; | 
 **filterPlatform** | **[]string** | filter by attribute &#39;platform&#39; | 
 **filterSeedId** | **[]string** | filter by attribute &#39;seedId&#39; | 
 **filterId** | **[]string** | filter by id(s) | 
 **sort** | **[]string** | comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsBundleIds** | **[]string** | the fields to include for returned resources of type bundleIds | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsBundleIdCapabilities** | **[]string** | the fields to include for returned resources of type bundleIdCapabilities | 
 **fieldsProfiles** | **[]string** | the fields to include for returned resources of type profiles | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limitBundleIdCapabilities** | **int32** | maximum number of related bundleIdCapabilities returned (when they are included) | 
 **limitProfiles** | **int32** | maximum number of related profiles returned (when they are included) | 

### Return type

[**BundleIdsResponse**](BundleIdsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BundleIdsGetInstance

> BundleIdResponse BundleIdsGetInstance(ctx, id).FieldsBundleIds(fieldsBundleIds).Include(include).FieldsBundleIdCapabilities(fieldsBundleIdCapabilities).FieldsProfiles(fieldsProfiles).FieldsApps(fieldsApps).LimitBundleIdCapabilities(limitBundleIdCapabilities).LimitProfiles(limitProfiles).Execute()



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
    fieldsBundleIds := []string{"FieldsBundleIds_example"} // []string | the fields to include for returned resources of type bundleIds (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsBundleIdCapabilities := []string{"FieldsBundleIdCapabilities_example"} // []string | the fields to include for returned resources of type bundleIdCapabilities (optional)
    fieldsProfiles := []string{"FieldsProfiles_example"} // []string | the fields to include for returned resources of type profiles (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limitBundleIdCapabilities := int32(56) // int32 | maximum number of related bundleIdCapabilities returned (when they are included) (optional)
    limitProfiles := int32(56) // int32 | maximum number of related profiles returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BundleIdsAPI.BundleIdsGetInstance(context.Background(), id).FieldsBundleIds(fieldsBundleIds).Include(include).FieldsBundleIdCapabilities(fieldsBundleIdCapabilities).FieldsProfiles(fieldsProfiles).FieldsApps(fieldsApps).LimitBundleIdCapabilities(limitBundleIdCapabilities).LimitProfiles(limitProfiles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundleIdsAPI.BundleIdsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BundleIdsGetInstance`: BundleIdResponse
    fmt.Fprintf(os.Stdout, "Response from `BundleIdsAPI.BundleIdsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBundleIdsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBundleIds** | **[]string** | the fields to include for returned resources of type bundleIds | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsBundleIdCapabilities** | **[]string** | the fields to include for returned resources of type bundleIdCapabilities | 
 **fieldsProfiles** | **[]string** | the fields to include for returned resources of type profiles | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limitBundleIdCapabilities** | **int32** | maximum number of related bundleIdCapabilities returned (when they are included) | 
 **limitProfiles** | **int32** | maximum number of related profiles returned (when they are included) | 

### Return type

[**BundleIdResponse**](BundleIdResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BundleIdsProfilesGetToManyRelated

> ProfilesWithoutIncludesResponse BundleIdsProfilesGetToManyRelated(ctx, id).FieldsProfiles(fieldsProfiles).Limit(limit).Execute()



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
    fieldsProfiles := []string{"FieldsProfiles_example"} // []string | the fields to include for returned resources of type profiles (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BundleIdsAPI.BundleIdsProfilesGetToManyRelated(context.Background(), id).FieldsProfiles(fieldsProfiles).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundleIdsAPI.BundleIdsProfilesGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BundleIdsProfilesGetToManyRelated`: ProfilesWithoutIncludesResponse
    fmt.Fprintf(os.Stdout, "Response from `BundleIdsAPI.BundleIdsProfilesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBundleIdsProfilesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsProfiles** | **[]string** | the fields to include for returned resources of type profiles | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**ProfilesWithoutIncludesResponse**](ProfilesWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BundleIdsUpdateInstance

> BundleIdResponse BundleIdsUpdateInstance(ctx, id).BundleIdUpdateRequest(bundleIdUpdateRequest).Execute()



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
    bundleIdUpdateRequest := *openapiclient.NewBundleIdUpdateRequest(*openapiclient.NewBundleIdUpdateRequestData("Type_example", "Id_example")) // BundleIdUpdateRequest | BundleId representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BundleIdsAPI.BundleIdsUpdateInstance(context.Background(), id).BundleIdUpdateRequest(bundleIdUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BundleIdsAPI.BundleIdsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BundleIdsUpdateInstance`: BundleIdResponse
    fmt.Fprintf(os.Stdout, "Response from `BundleIdsAPI.BundleIdsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBundleIdsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bundleIdUpdateRequest** | [**BundleIdUpdateRequest**](BundleIdUpdateRequest.md) | BundleId representation | 

### Return type

[**BundleIdResponse**](BundleIdResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

