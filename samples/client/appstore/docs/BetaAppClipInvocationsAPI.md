# \BetaAppClipInvocationsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BetaAppClipInvocationsCreateInstance**](BetaAppClipInvocationsAPI.md#BetaAppClipInvocationsCreateInstance) | **Post** /v1/betaAppClipInvocations | 
[**BetaAppClipInvocationsDeleteInstance**](BetaAppClipInvocationsAPI.md#BetaAppClipInvocationsDeleteInstance) | **Delete** /v1/betaAppClipInvocations/{id} | 
[**BetaAppClipInvocationsGetInstance**](BetaAppClipInvocationsAPI.md#BetaAppClipInvocationsGetInstance) | **Get** /v1/betaAppClipInvocations/{id} | 
[**BetaAppClipInvocationsUpdateInstance**](BetaAppClipInvocationsAPI.md#BetaAppClipInvocationsUpdateInstance) | **Patch** /v1/betaAppClipInvocations/{id} | 



## BetaAppClipInvocationsCreateInstance

> BetaAppClipInvocationResponse BetaAppClipInvocationsCreateInstance(ctx).BetaAppClipInvocationCreateRequest(betaAppClipInvocationCreateRequest).Execute()



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
    betaAppClipInvocationCreateRequest := *openapiclient.NewBetaAppClipInvocationCreateRequest(*openapiclient.NewBetaAppClipInvocationCreateRequestData("Type_example", *openapiclient.NewBetaAppClipInvocationCreateRequestDataAttributes("Url_example"), *openapiclient.NewBetaAppClipInvocationCreateRequestDataRelationships(*openapiclient.NewBetaAppClipInvocationCreateRequestDataRelationshipsBuildBundle(*openapiclient.NewBetaAppClipInvocationCreateRequestDataRelationshipsBuildBundleData("Type_example", "Id_example")), *openapiclient.NewBetaAppClipInvocationCreateRequestDataRelationshipsBetaAppClipInvocationLocalizations([]openapiclient.BetaAppClipInvocationRelationshipsBetaAppClipInvocationLocalizationsDataInner{*openapiclient.NewBetaAppClipInvocationRelationshipsBetaAppClipInvocationLocalizationsDataInner("Type_example", "Id_example")})))) // BetaAppClipInvocationCreateRequest | BetaAppClipInvocation representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BetaAppClipInvocationsAPI.BetaAppClipInvocationsCreateInstance(context.Background()).BetaAppClipInvocationCreateRequest(betaAppClipInvocationCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppClipInvocationsAPI.BetaAppClipInvocationsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaAppClipInvocationsCreateInstance`: BetaAppClipInvocationResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaAppClipInvocationsAPI.BetaAppClipInvocationsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBetaAppClipInvocationsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **betaAppClipInvocationCreateRequest** | [**BetaAppClipInvocationCreateRequest**](BetaAppClipInvocationCreateRequest.md) | BetaAppClipInvocation representation | 

### Return type

[**BetaAppClipInvocationResponse**](BetaAppClipInvocationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaAppClipInvocationsDeleteInstance

> BetaAppClipInvocationsDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.BetaAppClipInvocationsAPI.BetaAppClipInvocationsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppClipInvocationsAPI.BetaAppClipInvocationsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBetaAppClipInvocationsDeleteInstanceRequest struct via the builder pattern


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


## BetaAppClipInvocationsGetInstance

> BetaAppClipInvocationResponse BetaAppClipInvocationsGetInstance(ctx, id).FieldsBetaAppClipInvocations(fieldsBetaAppClipInvocations).Include(include).LimitBetaAppClipInvocationLocalizations(limitBetaAppClipInvocationLocalizations).Execute()



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
    fieldsBetaAppClipInvocations := []string{"FieldsBetaAppClipInvocations_example"} // []string | the fields to include for returned resources of type betaAppClipInvocations (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    limitBetaAppClipInvocationLocalizations := int32(56) // int32 | maximum number of related betaAppClipInvocationLocalizations returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BetaAppClipInvocationsAPI.BetaAppClipInvocationsGetInstance(context.Background(), id).FieldsBetaAppClipInvocations(fieldsBetaAppClipInvocations).Include(include).LimitBetaAppClipInvocationLocalizations(limitBetaAppClipInvocationLocalizations).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppClipInvocationsAPI.BetaAppClipInvocationsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaAppClipInvocationsGetInstance`: BetaAppClipInvocationResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaAppClipInvocationsAPI.BetaAppClipInvocationsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaAppClipInvocationsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBetaAppClipInvocations** | **[]string** | the fields to include for returned resources of type betaAppClipInvocations | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **limitBetaAppClipInvocationLocalizations** | **int32** | maximum number of related betaAppClipInvocationLocalizations returned (when they are included) | 

### Return type

[**BetaAppClipInvocationResponse**](BetaAppClipInvocationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaAppClipInvocationsUpdateInstance

> BetaAppClipInvocationResponse BetaAppClipInvocationsUpdateInstance(ctx, id).BetaAppClipInvocationUpdateRequest(betaAppClipInvocationUpdateRequest).Execute()



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
    betaAppClipInvocationUpdateRequest := *openapiclient.NewBetaAppClipInvocationUpdateRequest(*openapiclient.NewBetaAppClipInvocationUpdateRequestData("Type_example", "Id_example")) // BetaAppClipInvocationUpdateRequest | BetaAppClipInvocation representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BetaAppClipInvocationsAPI.BetaAppClipInvocationsUpdateInstance(context.Background(), id).BetaAppClipInvocationUpdateRequest(betaAppClipInvocationUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BetaAppClipInvocationsAPI.BetaAppClipInvocationsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaAppClipInvocationsUpdateInstance`: BetaAppClipInvocationResponse
    fmt.Fprintf(os.Stdout, "Response from `BetaAppClipInvocationsAPI.BetaAppClipInvocationsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaAppClipInvocationsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **betaAppClipInvocationUpdateRequest** | [**BetaAppClipInvocationUpdateRequest**](BetaAppClipInvocationUpdateRequest.md) | BetaAppClipInvocation representation | 

### Return type

[**BetaAppClipInvocationResponse**](BetaAppClipInvocationResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

