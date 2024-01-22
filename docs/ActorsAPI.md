# \ActorsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActorsGetCollection**](ActorsAPI.md#ActorsGetCollection) | **Get** /v1/actors | 
[**ActorsGetInstance**](ActorsAPI.md#ActorsGetInstance) | **Get** /v1/actors/{id} | 



## ActorsGetCollection

> ActorsResponse ActorsGetCollection(ctx).FilterId(filterId).FieldsActors(fieldsActors).Limit(limit).Execute()



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
    filterId := []string{"Inner_example"} // []string | filter by id(s)
    fieldsActors := []string{"FieldsActors_example"} // []string | the fields to include for returned resources of type actors (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActorsAPI.ActorsGetCollection(context.Background()).FilterId(filterId).FieldsActors(fieldsActors).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActorsAPI.ActorsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActorsGetCollection`: ActorsResponse
    fmt.Fprintf(os.Stdout, "Response from `ActorsAPI.ActorsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActorsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterId** | **[]string** | filter by id(s) | 
 **fieldsActors** | **[]string** | the fields to include for returned resources of type actors | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**ActorsResponse**](ActorsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActorsGetInstance

> ActorResponse ActorsGetInstance(ctx, id).FieldsActors(fieldsActors).Execute()



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
    fieldsActors := []string{"FieldsActors_example"} // []string | the fields to include for returned resources of type actors (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActorsAPI.ActorsGetInstance(context.Background(), id).FieldsActors(fieldsActors).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActorsAPI.ActorsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActorsGetInstance`: ActorResponse
    fmt.Fprintf(os.Stdout, "Response from `ActorsAPI.ActorsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiActorsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsActors** | **[]string** | the fields to include for returned resources of type actors | 

### Return type

[**ActorResponse**](ActorResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

