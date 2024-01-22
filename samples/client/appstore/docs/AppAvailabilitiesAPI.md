# \AppAvailabilitiesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppAvailabilitiesAvailableTerritoriesGetToManyRelated**](AppAvailabilitiesAPI.md#AppAvailabilitiesAvailableTerritoriesGetToManyRelated) | **Get** /v1/appAvailabilities/{id}/availableTerritories | 
[**AppAvailabilitiesCreateInstance**](AppAvailabilitiesAPI.md#AppAvailabilitiesCreateInstance) | **Post** /v1/appAvailabilities | 
[**AppAvailabilitiesGetInstance**](AppAvailabilitiesAPI.md#AppAvailabilitiesGetInstance) | **Get** /v1/appAvailabilities/{id} | 
[**AppAvailabilitiesV2CreateInstance**](AppAvailabilitiesAPI.md#AppAvailabilitiesV2CreateInstance) | **Post** /v2/appAvailabilities | 
[**AppAvailabilitiesV2GetInstance**](AppAvailabilitiesAPI.md#AppAvailabilitiesV2GetInstance) | **Get** /v2/appAvailabilities/{id} | 
[**AppAvailabilitiesV2TerritoryAvailabilitiesGetToManyRelated**](AppAvailabilitiesAPI.md#AppAvailabilitiesV2TerritoryAvailabilitiesGetToManyRelated) | **Get** /v2/appAvailabilities/{id}/territoryAvailabilities | 



## AppAvailabilitiesAvailableTerritoriesGetToManyRelated

> TerritoriesResponse AppAvailabilitiesAvailableTerritoriesGetToManyRelated(ctx, id).FieldsTerritories(fieldsTerritories).Limit(limit).Execute()



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
    fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppAvailabilitiesAPI.AppAvailabilitiesAvailableTerritoriesGetToManyRelated(context.Background(), id).FieldsTerritories(fieldsTerritories).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppAvailabilitiesAPI.AppAvailabilitiesAvailableTerritoriesGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppAvailabilitiesAvailableTerritoriesGetToManyRelated`: TerritoriesResponse
    fmt.Fprintf(os.Stdout, "Response from `AppAvailabilitiesAPI.AppAvailabilitiesAvailableTerritoriesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**TerritoriesResponse**](TerritoriesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppAvailabilitiesCreateInstance

> AppAvailabilityResponse AppAvailabilitiesCreateInstance(ctx).AppAvailabilityCreateRequest(appAvailabilityCreateRequest).Execute()



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
    appAvailabilityCreateRequest := *openapiclient.NewAppAvailabilityCreateRequest(*openapiclient.NewAppAvailabilityCreateRequestData("Type_example", *openapiclient.NewAppAvailabilityV2CreateRequestDataAttributes(false), *openapiclient.NewAppAvailabilityCreateRequestDataRelationships(*openapiclient.NewAppAvailabilityV2CreateRequestDataRelationshipsApp(*openapiclient.NewAppAvailabilityV2CreateRequestDataRelationshipsAppData("Type_example", "Id_example")), *openapiclient.NewAppAvailabilityCreateRequestDataRelationshipsAvailableTerritories([]openapiclient.AppAvailabilityRelationshipsAvailableTerritoriesDataInner{*openapiclient.NewAppAvailabilityRelationshipsAvailableTerritoriesDataInner("Type_example", "Id_example")})))) // AppAvailabilityCreateRequest | AppAvailability representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppAvailabilitiesAPI.AppAvailabilitiesCreateInstance(context.Background()).AppAvailabilityCreateRequest(appAvailabilityCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppAvailabilitiesAPI.AppAvailabilitiesCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppAvailabilitiesCreateInstance`: AppAvailabilityResponse
    fmt.Fprintf(os.Stdout, "Response from `AppAvailabilitiesAPI.AppAvailabilitiesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppAvailabilitiesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appAvailabilityCreateRequest** | [**AppAvailabilityCreateRequest**](AppAvailabilityCreateRequest.md) | AppAvailability representation | 

### Return type

[**AppAvailabilityResponse**](AppAvailabilityResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppAvailabilitiesGetInstance

> AppAvailabilityResponse AppAvailabilitiesGetInstance(ctx, id).FieldsAppAvailabilities(fieldsAppAvailabilities).Include(include).FieldsTerritories(fieldsTerritories).LimitAvailableTerritories(limitAvailableTerritories).Execute()



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
    fieldsAppAvailabilities := []string{"FieldsAppAvailabilities_example"} // []string | the fields to include for returned resources of type appAvailabilities (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
    limitAvailableTerritories := int32(56) // int32 | maximum number of related availableTerritories returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppAvailabilitiesAPI.AppAvailabilitiesGetInstance(context.Background(), id).FieldsAppAvailabilities(fieldsAppAvailabilities).Include(include).FieldsTerritories(fieldsTerritories).LimitAvailableTerritories(limitAvailableTerritories).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppAvailabilitiesAPI.AppAvailabilitiesGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppAvailabilitiesGetInstance`: AppAvailabilityResponse
    fmt.Fprintf(os.Stdout, "Response from `AppAvailabilitiesAPI.AppAvailabilitiesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppAvailabilitiesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppAvailabilities** | **[]string** | the fields to include for returned resources of type appAvailabilities | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 
 **limitAvailableTerritories** | **int32** | maximum number of related availableTerritories returned (when they are included) | 

### Return type

[**AppAvailabilityResponse**](AppAvailabilityResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppAvailabilitiesV2CreateInstance

> AppAvailabilityV2Response AppAvailabilitiesV2CreateInstance(ctx).AppAvailabilityV2CreateRequest(appAvailabilityV2CreateRequest).Execute()



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
    appAvailabilityV2CreateRequest := *openapiclient.NewAppAvailabilityV2CreateRequest(*openapiclient.NewAppAvailabilityV2CreateRequestData("Type_example", *openapiclient.NewAppAvailabilityV2CreateRequestDataAttributes(false), *openapiclient.NewAppAvailabilityV2CreateRequestDataRelationships(*openapiclient.NewAppAvailabilityV2CreateRequestDataRelationshipsApp(*openapiclient.NewAppAvailabilityV2CreateRequestDataRelationshipsAppData("Type_example", "Id_example")), *openapiclient.NewAppAvailabilityV2CreateRequestDataRelationshipsTerritoryAvailabilities([]openapiclient.AppAvailabilityV2RelationshipsTerritoryAvailabilitiesDataInner{*openapiclient.NewAppAvailabilityV2RelationshipsTerritoryAvailabilitiesDataInner("Type_example", "Id_example")})))) // AppAvailabilityV2CreateRequest | AppAvailability representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppAvailabilitiesAPI.AppAvailabilitiesV2CreateInstance(context.Background()).AppAvailabilityV2CreateRequest(appAvailabilityV2CreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppAvailabilitiesAPI.AppAvailabilitiesV2CreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppAvailabilitiesV2CreateInstance`: AppAvailabilityV2Response
    fmt.Fprintf(os.Stdout, "Response from `AppAvailabilitiesAPI.AppAvailabilitiesV2CreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppAvailabilitiesV2CreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appAvailabilityV2CreateRequest** | [**AppAvailabilityV2CreateRequest**](AppAvailabilityV2CreateRequest.md) | AppAvailability representation | 

### Return type

[**AppAvailabilityV2Response**](AppAvailabilityV2Response.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppAvailabilitiesV2GetInstance

> AppAvailabilityV2Response AppAvailabilitiesV2GetInstance(ctx, id).FieldsAppAvailabilities(fieldsAppAvailabilities).Include(include).FieldsTerritoryAvailabilities(fieldsTerritoryAvailabilities).LimitTerritoryAvailabilities(limitTerritoryAvailabilities).Execute()



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
    fieldsAppAvailabilities := []string{"FieldsAppAvailabilities_example"} // []string | the fields to include for returned resources of type appAvailabilities (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsTerritoryAvailabilities := []string{"FieldsTerritoryAvailabilities_example"} // []string | the fields to include for returned resources of type territoryAvailabilities (optional)
    limitTerritoryAvailabilities := int32(56) // int32 | maximum number of related territoryAvailabilities returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppAvailabilitiesAPI.AppAvailabilitiesV2GetInstance(context.Background(), id).FieldsAppAvailabilities(fieldsAppAvailabilities).Include(include).FieldsTerritoryAvailabilities(fieldsTerritoryAvailabilities).LimitTerritoryAvailabilities(limitTerritoryAvailabilities).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppAvailabilitiesAPI.AppAvailabilitiesV2GetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppAvailabilitiesV2GetInstance`: AppAvailabilityV2Response
    fmt.Fprintf(os.Stdout, "Response from `AppAvailabilitiesAPI.AppAvailabilitiesV2GetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppAvailabilitiesV2GetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppAvailabilities** | **[]string** | the fields to include for returned resources of type appAvailabilities | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsTerritoryAvailabilities** | **[]string** | the fields to include for returned resources of type territoryAvailabilities | 
 **limitTerritoryAvailabilities** | **int32** | maximum number of related territoryAvailabilities returned (when they are included) | 

### Return type

[**AppAvailabilityV2Response**](AppAvailabilityV2Response.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppAvailabilitiesV2TerritoryAvailabilitiesGetToManyRelated

> TerritoryAvailabilitiesResponse AppAvailabilitiesV2TerritoryAvailabilitiesGetToManyRelated(ctx, id).FieldsTerritoryAvailabilities(fieldsTerritoryAvailabilities).FieldsTerritories(fieldsTerritories).Limit(limit).Include(include).Execute()



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
    fieldsTerritoryAvailabilities := []string{"FieldsTerritoryAvailabilities_example"} // []string | the fields to include for returned resources of type territoryAvailabilities (optional)
    fieldsTerritories := []string{"FieldsTerritories_example"} // []string | the fields to include for returned resources of type territories (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppAvailabilitiesAPI.AppAvailabilitiesV2TerritoryAvailabilitiesGetToManyRelated(context.Background(), id).FieldsTerritoryAvailabilities(fieldsTerritoryAvailabilities).FieldsTerritories(fieldsTerritories).Limit(limit).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppAvailabilitiesAPI.AppAvailabilitiesV2TerritoryAvailabilitiesGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppAvailabilitiesV2TerritoryAvailabilitiesGetToManyRelated`: TerritoryAvailabilitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `AppAvailabilitiesAPI.AppAvailabilitiesV2TerritoryAvailabilitiesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppAvailabilitiesV2TerritoryAvailabilitiesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsTerritoryAvailabilities** | **[]string** | the fields to include for returned resources of type territoryAvailabilities | 
 **fieldsTerritories** | **[]string** | the fields to include for returned resources of type territories | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**TerritoryAvailabilitiesResponse**](TerritoryAvailabilitiesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

