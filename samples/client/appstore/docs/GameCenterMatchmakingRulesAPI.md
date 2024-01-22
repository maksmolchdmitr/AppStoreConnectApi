# \GameCenterMatchmakingRulesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterMatchmakingRulesCreateInstance**](GameCenterMatchmakingRulesAPI.md#GameCenterMatchmakingRulesCreateInstance) | **Post** /v1/gameCenterMatchmakingRules | 
[**GameCenterMatchmakingRulesDeleteInstance**](GameCenterMatchmakingRulesAPI.md#GameCenterMatchmakingRulesDeleteInstance) | **Delete** /v1/gameCenterMatchmakingRules/{id} | 
[**GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics**](GameCenterMatchmakingRulesAPI.md#GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics) | **Get** /v1/gameCenterMatchmakingRules/{id}/metrics/matchmakingBooleanRuleResults | 
[**GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics**](GameCenterMatchmakingRulesAPI.md#GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics) | **Get** /v1/gameCenterMatchmakingRules/{id}/metrics/matchmakingNumberRuleResults | 
[**GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics**](GameCenterMatchmakingRulesAPI.md#GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics) | **Get** /v1/gameCenterMatchmakingRules/{id}/metrics/matchmakingRuleErrors | 
[**GameCenterMatchmakingRulesUpdateInstance**](GameCenterMatchmakingRulesAPI.md#GameCenterMatchmakingRulesUpdateInstance) | **Patch** /v1/gameCenterMatchmakingRules/{id} | 



## GameCenterMatchmakingRulesCreateInstance

> GameCenterMatchmakingRuleResponse GameCenterMatchmakingRulesCreateInstance(ctx).GameCenterMatchmakingRuleCreateRequest(gameCenterMatchmakingRuleCreateRequest).Execute()



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
    gameCenterMatchmakingRuleCreateRequest := *openapiclient.NewGameCenterMatchmakingRuleCreateRequest(*openapiclient.NewGameCenterMatchmakingRuleCreateRequestData("Type_example", *openapiclient.NewGameCenterMatchmakingRuleCreateRequestDataAttributes("ReferenceName_example", "Description_example", "Type_example", "Expression_example"), *openapiclient.NewGameCenterMatchmakingRuleCreateRequestDataRelationships(*openapiclient.NewGameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet(*openapiclient.NewGameCenterMatchmakingQueueRelationshipsRuleSetData("Type_example", "Id_example"))))) // GameCenterMatchmakingRuleCreateRequest | GameCenterMatchmakingRule representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingRulesAPI.GameCenterMatchmakingRulesCreateInstance(context.Background()).GameCenterMatchmakingRuleCreateRequest(gameCenterMatchmakingRuleCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingRulesAPI.GameCenterMatchmakingRulesCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingRulesCreateInstance`: GameCenterMatchmakingRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingRulesAPI.GameCenterMatchmakingRulesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingRulesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameCenterMatchmakingRuleCreateRequest** | [**GameCenterMatchmakingRuleCreateRequest**](GameCenterMatchmakingRuleCreateRequest.md) | GameCenterMatchmakingRule representation | 

### Return type

[**GameCenterMatchmakingRuleResponse**](GameCenterMatchmakingRuleResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingRulesDeleteInstance

> GameCenterMatchmakingRulesDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.GameCenterMatchmakingRulesAPI.GameCenterMatchmakingRulesDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingRulesAPI.GameCenterMatchmakingRulesDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterMatchmakingRulesDeleteInstanceRequest struct via the builder pattern


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


## GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics

> GameCenterMatchmakingBooleanRuleResultsV1MetricResponse GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics(ctx, id).Granularity(granularity).Limit(limit).GroupBy(groupBy).FilterResult(filterResult).FilterGameCenterMatchmakingQueue(filterGameCenterMatchmakingQueue).Sort(sort).Execute()



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
    granularity := []string{"Granularity_example"} // []string | the granularity of the per-group dataset
    limit := int32(56) // int32 | maximum number of groups to return per page (optional)
    groupBy := []string{"GroupBy_example"} // []string | the dimension by which to group the results (optional)
    filterResult := "filterResult_example" // string | filter by 'result' attribute dimension (optional)
    filterGameCenterMatchmakingQueue := "filterGameCenterMatchmakingQueue_example" // string | filter by 'gameCenterMatchmakingQueue' relationship dimension (optional)
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; metrics will be sorted as specified (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingRulesAPI.GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics(context.Background(), id).Granularity(granularity).Limit(limit).GroupBy(groupBy).FilterResult(filterResult).FilterGameCenterMatchmakingQueue(filterGameCenterMatchmakingQueue).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingRulesAPI.GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics`: GameCenterMatchmakingBooleanRuleResultsV1MetricResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingRulesAPI.GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **granularity** | **[]string** | the granularity of the per-group dataset | 
 **limit** | **int32** | maximum number of groups to return per page | 
 **groupBy** | **[]string** | the dimension by which to group the results | 
 **filterResult** | **string** | filter by &#39;result&#39; attribute dimension | 
 **filterGameCenterMatchmakingQueue** | **string** | filter by &#39;gameCenterMatchmakingQueue&#39; relationship dimension | 
 **sort** | **[]string** | comma-separated list of sort expressions; metrics will be sorted as specified | 

### Return type

[**GameCenterMatchmakingBooleanRuleResultsV1MetricResponse**](GameCenterMatchmakingBooleanRuleResultsV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics

> GameCenterMatchmakingNumberRuleResultsV1MetricResponse GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics(ctx, id).Granularity(granularity).Limit(limit).GroupBy(groupBy).FilterGameCenterMatchmakingQueue(filterGameCenterMatchmakingQueue).Sort(sort).Execute()



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
    granularity := []string{"Granularity_example"} // []string | the granularity of the per-group dataset
    limit := int32(56) // int32 | maximum number of groups to return per page (optional)
    groupBy := []string{"GroupBy_example"} // []string | the dimension by which to group the results (optional)
    filterGameCenterMatchmakingQueue := "filterGameCenterMatchmakingQueue_example" // string | filter by 'gameCenterMatchmakingQueue' relationship dimension (optional)
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; metrics will be sorted as specified (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingRulesAPI.GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics(context.Background(), id).Granularity(granularity).Limit(limit).GroupBy(groupBy).FilterGameCenterMatchmakingQueue(filterGameCenterMatchmakingQueue).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingRulesAPI.GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics`: GameCenterMatchmakingNumberRuleResultsV1MetricResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingRulesAPI.GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **granularity** | **[]string** | the granularity of the per-group dataset | 
 **limit** | **int32** | maximum number of groups to return per page | 
 **groupBy** | **[]string** | the dimension by which to group the results | 
 **filterGameCenterMatchmakingQueue** | **string** | filter by &#39;gameCenterMatchmakingQueue&#39; relationship dimension | 
 **sort** | **[]string** | comma-separated list of sort expressions; metrics will be sorted as specified | 

### Return type

[**GameCenterMatchmakingNumberRuleResultsV1MetricResponse**](GameCenterMatchmakingNumberRuleResultsV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics

> GameCenterMatchmakingRuleErrorsV1MetricResponse GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics(ctx, id).Granularity(granularity).Limit(limit).GroupBy(groupBy).FilterGameCenterMatchmakingQueue(filterGameCenterMatchmakingQueue).Sort(sort).Execute()



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
    granularity := []string{"Granularity_example"} // []string | the granularity of the per-group dataset
    limit := int32(56) // int32 | maximum number of groups to return per page (optional)
    groupBy := []string{"GroupBy_example"} // []string | the dimension by which to group the results (optional)
    filterGameCenterMatchmakingQueue := "filterGameCenterMatchmakingQueue_example" // string | filter by 'gameCenterMatchmakingQueue' relationship dimension (optional)
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; metrics will be sorted as specified (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingRulesAPI.GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics(context.Background(), id).Granularity(granularity).Limit(limit).GroupBy(groupBy).FilterGameCenterMatchmakingQueue(filterGameCenterMatchmakingQueue).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingRulesAPI.GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics`: GameCenterMatchmakingRuleErrorsV1MetricResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingRulesAPI.GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **granularity** | **[]string** | the granularity of the per-group dataset | 
 **limit** | **int32** | maximum number of groups to return per page | 
 **groupBy** | **[]string** | the dimension by which to group the results | 
 **filterGameCenterMatchmakingQueue** | **string** | filter by &#39;gameCenterMatchmakingQueue&#39; relationship dimension | 
 **sort** | **[]string** | comma-separated list of sort expressions; metrics will be sorted as specified | 

### Return type

[**GameCenterMatchmakingRuleErrorsV1MetricResponse**](GameCenterMatchmakingRuleErrorsV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingRulesUpdateInstance

> GameCenterMatchmakingRuleResponse GameCenterMatchmakingRulesUpdateInstance(ctx, id).GameCenterMatchmakingRuleUpdateRequest(gameCenterMatchmakingRuleUpdateRequest).Execute()



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
    gameCenterMatchmakingRuleUpdateRequest := *openapiclient.NewGameCenterMatchmakingRuleUpdateRequest(*openapiclient.NewGameCenterMatchmakingRuleUpdateRequestData("Type_example", "Id_example")) // GameCenterMatchmakingRuleUpdateRequest | GameCenterMatchmakingRule representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingRulesAPI.GameCenterMatchmakingRulesUpdateInstance(context.Background(), id).GameCenterMatchmakingRuleUpdateRequest(gameCenterMatchmakingRuleUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingRulesAPI.GameCenterMatchmakingRulesUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingRulesUpdateInstance`: GameCenterMatchmakingRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingRulesAPI.GameCenterMatchmakingRulesUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingRulesUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterMatchmakingRuleUpdateRequest** | [**GameCenterMatchmakingRuleUpdateRequest**](GameCenterMatchmakingRuleUpdateRequest.md) | GameCenterMatchmakingRule representation | 

### Return type

[**GameCenterMatchmakingRuleResponse**](GameCenterMatchmakingRuleResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

