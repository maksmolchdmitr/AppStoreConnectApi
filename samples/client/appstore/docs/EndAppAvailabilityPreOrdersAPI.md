# \EndAppAvailabilityPreOrdersAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EndAppAvailabilityPreOrdersCreateInstance**](EndAppAvailabilityPreOrdersAPI.md#EndAppAvailabilityPreOrdersCreateInstance) | **Post** /v1/endAppAvailabilityPreOrders | 



## EndAppAvailabilityPreOrdersCreateInstance

> EndAppAvailabilityPreOrderResponse EndAppAvailabilityPreOrdersCreateInstance(ctx).EndAppAvailabilityPreOrderCreateRequest(endAppAvailabilityPreOrderCreateRequest).Execute()



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
    endAppAvailabilityPreOrderCreateRequest := *openapiclient.NewEndAppAvailabilityPreOrderCreateRequest(*openapiclient.NewEndAppAvailabilityPreOrderCreateRequestData("Type_example", *openapiclient.NewEndAppAvailabilityPreOrderCreateRequestDataRelationships(*openapiclient.NewAppAvailabilityV2CreateRequestDataRelationshipsTerritoryAvailabilities([]openapiclient.AppAvailabilityV2RelationshipsTerritoryAvailabilitiesDataInner{*openapiclient.NewAppAvailabilityV2RelationshipsTerritoryAvailabilitiesDataInner("Type_example", "Id_example")})))) // EndAppAvailabilityPreOrderCreateRequest | EndAppAvailabilityPreOrder representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EndAppAvailabilityPreOrdersAPI.EndAppAvailabilityPreOrdersCreateInstance(context.Background()).EndAppAvailabilityPreOrderCreateRequest(endAppAvailabilityPreOrderCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EndAppAvailabilityPreOrdersAPI.EndAppAvailabilityPreOrdersCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EndAppAvailabilityPreOrdersCreateInstance`: EndAppAvailabilityPreOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `EndAppAvailabilityPreOrdersAPI.EndAppAvailabilityPreOrdersCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEndAppAvailabilityPreOrdersCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endAppAvailabilityPreOrderCreateRequest** | [**EndAppAvailabilityPreOrderCreateRequest**](EndAppAvailabilityPreOrderCreateRequest.md) | EndAppAvailabilityPreOrder representation | 

### Return type

[**EndAppAvailabilityPreOrderResponse**](EndAppAvailabilityPreOrderResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

