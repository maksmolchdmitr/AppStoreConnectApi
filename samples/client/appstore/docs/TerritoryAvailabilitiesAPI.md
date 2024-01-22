# \TerritoryAvailabilitiesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TerritoryAvailabilitiesUpdateInstance**](TerritoryAvailabilitiesAPI.md#TerritoryAvailabilitiesUpdateInstance) | **Patch** /v1/territoryAvailabilities/{id} | 



## TerritoryAvailabilitiesUpdateInstance

> TerritoryAvailabilityResponse TerritoryAvailabilitiesUpdateInstance(ctx, id).TerritoryAvailabilityUpdateRequest(territoryAvailabilityUpdateRequest).Execute()



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
    territoryAvailabilityUpdateRequest := *openapiclient.NewTerritoryAvailabilityUpdateRequest(*openapiclient.NewTerritoryAvailabilityUpdateRequestData("Type_example", "Id_example")) // TerritoryAvailabilityUpdateRequest | TerritoryAvailability representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TerritoryAvailabilitiesAPI.TerritoryAvailabilitiesUpdateInstance(context.Background(), id).TerritoryAvailabilityUpdateRequest(territoryAvailabilityUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TerritoryAvailabilitiesAPI.TerritoryAvailabilitiesUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TerritoryAvailabilitiesUpdateInstance`: TerritoryAvailabilityResponse
    fmt.Fprintf(os.Stdout, "Response from `TerritoryAvailabilitiesAPI.TerritoryAvailabilitiesUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerritoryAvailabilitiesUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **territoryAvailabilityUpdateRequest** | [**TerritoryAvailabilityUpdateRequest**](TerritoryAvailabilityUpdateRequest.md) | TerritoryAvailability representation | 

### Return type

[**TerritoryAvailabilityResponse**](TerritoryAvailabilityResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

