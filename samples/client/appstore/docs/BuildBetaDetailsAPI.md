# \BuildBetaDetailsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuildBetaDetailsBuildGetToOneRelated**](BuildBetaDetailsAPI.md#BuildBetaDetailsBuildGetToOneRelated) | **Get** /v1/buildBetaDetails/{id}/build | 
[**BuildBetaDetailsGetCollection**](BuildBetaDetailsAPI.md#BuildBetaDetailsGetCollection) | **Get** /v1/buildBetaDetails | 
[**BuildBetaDetailsGetInstance**](BuildBetaDetailsAPI.md#BuildBetaDetailsGetInstance) | **Get** /v1/buildBetaDetails/{id} | 
[**BuildBetaDetailsUpdateInstance**](BuildBetaDetailsAPI.md#BuildBetaDetailsUpdateInstance) | **Patch** /v1/buildBetaDetails/{id} | 



## BuildBetaDetailsBuildGetToOneRelated

> BuildWithoutIncludesResponse BuildBetaDetailsBuildGetToOneRelated(ctx, id).FieldsBuilds(fieldsBuilds).Execute()



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
    fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildBetaDetailsAPI.BuildBetaDetailsBuildGetToOneRelated(context.Background(), id).FieldsBuilds(fieldsBuilds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildBetaDetailsAPI.BuildBetaDetailsBuildGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildBetaDetailsBuildGetToOneRelated`: BuildWithoutIncludesResponse
    fmt.Fprintf(os.Stdout, "Response from `BuildBetaDetailsAPI.BuildBetaDetailsBuildGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildBetaDetailsBuildGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 

### Return type

[**BuildWithoutIncludesResponse**](BuildWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildBetaDetailsGetCollection

> BuildBetaDetailsResponse BuildBetaDetailsGetCollection(ctx).FilterBuild(filterBuild).FilterId(filterId).FieldsBuildBetaDetails(fieldsBuildBetaDetails).Limit(limit).Include(include).FieldsBuilds(fieldsBuilds).Execute()



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
    filterBuild := []string{"Inner_example"} // []string | filter by id(s) of related 'build' (optional)
    filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
    fieldsBuildBetaDetails := []string{"FieldsBuildBetaDetails_example"} // []string | the fields to include for returned resources of type buildBetaDetails (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildBetaDetailsAPI.BuildBetaDetailsGetCollection(context.Background()).FilterBuild(filterBuild).FilterId(filterId).FieldsBuildBetaDetails(fieldsBuildBetaDetails).Limit(limit).Include(include).FieldsBuilds(fieldsBuilds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildBetaDetailsAPI.BuildBetaDetailsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildBetaDetailsGetCollection`: BuildBetaDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuildBetaDetailsAPI.BuildBetaDetailsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBuildBetaDetailsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterBuild** | **[]string** | filter by id(s) of related &#39;build&#39; | 
 **filterId** | **[]string** | filter by id(s) | 
 **fieldsBuildBetaDetails** | **[]string** | the fields to include for returned resources of type buildBetaDetails | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 

### Return type

[**BuildBetaDetailsResponse**](BuildBetaDetailsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildBetaDetailsGetInstance

> BuildBetaDetailResponse BuildBetaDetailsGetInstance(ctx, id).FieldsBuildBetaDetails(fieldsBuildBetaDetails).Include(include).FieldsBuilds(fieldsBuilds).Execute()



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
    fieldsBuildBetaDetails := []string{"FieldsBuildBetaDetails_example"} // []string | the fields to include for returned resources of type buildBetaDetails (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsBuilds := []string{"FieldsBuilds_example"} // []string | the fields to include for returned resources of type builds (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildBetaDetailsAPI.BuildBetaDetailsGetInstance(context.Background(), id).FieldsBuildBetaDetails(fieldsBuildBetaDetails).Include(include).FieldsBuilds(fieldsBuilds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildBetaDetailsAPI.BuildBetaDetailsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildBetaDetailsGetInstance`: BuildBetaDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `BuildBetaDetailsAPI.BuildBetaDetailsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildBetaDetailsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBuildBetaDetails** | **[]string** | the fields to include for returned resources of type buildBetaDetails | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsBuilds** | **[]string** | the fields to include for returned resources of type builds | 

### Return type

[**BuildBetaDetailResponse**](BuildBetaDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildBetaDetailsUpdateInstance

> BuildBetaDetailResponse BuildBetaDetailsUpdateInstance(ctx, id).BuildBetaDetailUpdateRequest(buildBetaDetailUpdateRequest).Execute()



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
    buildBetaDetailUpdateRequest := *openapiclient.NewBuildBetaDetailUpdateRequest(*openapiclient.NewBuildBetaDetailUpdateRequestData("Type_example", "Id_example")) // BuildBetaDetailUpdateRequest | BuildBetaDetail representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuildBetaDetailsAPI.BuildBetaDetailsUpdateInstance(context.Background(), id).BuildBetaDetailUpdateRequest(buildBetaDetailUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuildBetaDetailsAPI.BuildBetaDetailsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildBetaDetailsUpdateInstance`: BuildBetaDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `BuildBetaDetailsAPI.BuildBetaDetailsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildBetaDetailsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buildBetaDetailUpdateRequest** | [**BuildBetaDetailUpdateRequest**](BuildBetaDetailUpdateRequest.md) | BuildBetaDetail representation | 

### Return type

[**BuildBetaDetailResponse**](BuildBetaDetailResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

