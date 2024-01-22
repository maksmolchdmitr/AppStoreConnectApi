# \GameCenterMatchmakingQueuesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterMatchmakingQueuesCreateInstance**](GameCenterMatchmakingQueuesAPI.md#GameCenterMatchmakingQueuesCreateInstance) | **Post** /v1/gameCenterMatchmakingQueues | 
[**GameCenterMatchmakingQueuesDeleteInstance**](GameCenterMatchmakingQueuesAPI.md#GameCenterMatchmakingQueuesDeleteInstance) | **Delete** /v1/gameCenterMatchmakingQueues/{id} | 
[**GameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetrics**](GameCenterMatchmakingQueuesAPI.md#GameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetrics) | **Get** /v1/gameCenterMatchmakingQueues/{id}/metrics/experimentMatchmakingQueueSizes | 
[**GameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetrics**](GameCenterMatchmakingQueuesAPI.md#GameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetrics) | **Get** /v1/gameCenterMatchmakingQueues/{id}/metrics/experimentMatchmakingRequests | 
[**GameCenterMatchmakingQueuesGetCollection**](GameCenterMatchmakingQueuesAPI.md#GameCenterMatchmakingQueuesGetCollection) | **Get** /v1/gameCenterMatchmakingQueues | 
[**GameCenterMatchmakingQueuesGetInstance**](GameCenterMatchmakingQueuesAPI.md#GameCenterMatchmakingQueuesGetInstance) | **Get** /v1/gameCenterMatchmakingQueues/{id} | 
[**GameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetrics**](GameCenterMatchmakingQueuesAPI.md#GameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetrics) | **Get** /v1/gameCenterMatchmakingQueues/{id}/metrics/matchmakingQueueSizes | 
[**GameCenterMatchmakingQueuesMatchmakingRequestsGetMetrics**](GameCenterMatchmakingQueuesAPI.md#GameCenterMatchmakingQueuesMatchmakingRequestsGetMetrics) | **Get** /v1/gameCenterMatchmakingQueues/{id}/metrics/matchmakingRequests | 
[**GameCenterMatchmakingQueuesMatchmakingSessionsGetMetrics**](GameCenterMatchmakingQueuesAPI.md#GameCenterMatchmakingQueuesMatchmakingSessionsGetMetrics) | **Get** /v1/gameCenterMatchmakingQueues/{id}/metrics/matchmakingSessions | 
[**GameCenterMatchmakingQueuesUpdateInstance**](GameCenterMatchmakingQueuesAPI.md#GameCenterMatchmakingQueuesUpdateInstance) | **Patch** /v1/gameCenterMatchmakingQueues/{id} | 



## GameCenterMatchmakingQueuesCreateInstance

> GameCenterMatchmakingQueueResponse GameCenterMatchmakingQueuesCreateInstance(ctx).GameCenterMatchmakingQueueCreateRequest(gameCenterMatchmakingQueueCreateRequest).Execute()



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
    gameCenterMatchmakingQueueCreateRequest := *openapiclient.NewGameCenterMatchmakingQueueCreateRequest(*openapiclient.NewGameCenterMatchmakingQueueCreateRequestData("Type_example", *openapiclient.NewGameCenterMatchmakingQueueCreateRequestDataAttributes("ReferenceName_example"), *openapiclient.NewGameCenterMatchmakingQueueCreateRequestDataRelationships(*openapiclient.NewGameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet(*openapiclient.NewGameCenterMatchmakingQueueRelationshipsRuleSetData("Type_example", "Id_example"))))) // GameCenterMatchmakingQueueCreateRequest | GameCenterMatchmakingQueue representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesCreateInstance(context.Background()).GameCenterMatchmakingQueueCreateRequest(gameCenterMatchmakingQueueCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingQueuesCreateInstance`: GameCenterMatchmakingQueueResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingQueuesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameCenterMatchmakingQueueCreateRequest** | [**GameCenterMatchmakingQueueCreateRequest**](GameCenterMatchmakingQueueCreateRequest.md) | GameCenterMatchmakingQueue representation | 

### Return type

[**GameCenterMatchmakingQueueResponse**](GameCenterMatchmakingQueueResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingQueuesDeleteInstance

> GameCenterMatchmakingQueuesDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterMatchmakingQueuesDeleteInstanceRequest struct via the builder pattern


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


## GameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetrics

> GameCenterMatchmakingQueueSizesV1MetricResponse GameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetrics(ctx, id).Granularity(granularity).Limit(limit).Sort(sort).Execute()



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
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; metrics will be sorted as specified (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetrics(context.Background(), id).Granularity(granularity).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetrics`: GameCenterMatchmakingQueueSizesV1MetricResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **granularity** | **[]string** | the granularity of the per-group dataset | 
 **limit** | **int32** | maximum number of groups to return per page | 
 **sort** | **[]string** | comma-separated list of sort expressions; metrics will be sorted as specified | 

### Return type

[**GameCenterMatchmakingQueueSizesV1MetricResponse**](GameCenterMatchmakingQueueSizesV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetrics

> GameCenterMatchmakingQueueRequestsV1MetricResponse GameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetrics(ctx, id).Granularity(granularity).Limit(limit).GroupBy(groupBy).FilterResult(filterResult).FilterGameCenterDetail(filterGameCenterDetail).Sort(sort).Execute()



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
    filterGameCenterDetail := "filterGameCenterDetail_example" // string | filter by 'gameCenterDetail' relationship dimension (optional)
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; metrics will be sorted as specified (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetrics(context.Background(), id).Granularity(granularity).Limit(limit).GroupBy(groupBy).FilterResult(filterResult).FilterGameCenterDetail(filterGameCenterDetail).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetrics`: GameCenterMatchmakingQueueRequestsV1MetricResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **granularity** | **[]string** | the granularity of the per-group dataset | 
 **limit** | **int32** | maximum number of groups to return per page | 
 **groupBy** | **[]string** | the dimension by which to group the results | 
 **filterResult** | **string** | filter by &#39;result&#39; attribute dimension | 
 **filterGameCenterDetail** | **string** | filter by &#39;gameCenterDetail&#39; relationship dimension | 
 **sort** | **[]string** | comma-separated list of sort expressions; metrics will be sorted as specified | 

### Return type

[**GameCenterMatchmakingQueueRequestsV1MetricResponse**](GameCenterMatchmakingQueueRequestsV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingQueuesGetCollection

> GameCenterMatchmakingQueuesResponse GameCenterMatchmakingQueuesGetCollection(ctx).FieldsGameCenterMatchmakingQueues(fieldsGameCenterMatchmakingQueues).Limit(limit).Include(include).Execute()



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
    fieldsGameCenterMatchmakingQueues := []string{"FieldsGameCenterMatchmakingQueues_example"} // []string | the fields to include for returned resources of type gameCenterMatchmakingQueues (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesGetCollection(context.Background()).FieldsGameCenterMatchmakingQueues(fieldsGameCenterMatchmakingQueues).Limit(limit).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingQueuesGetCollection`: GameCenterMatchmakingQueuesResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingQueuesGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fieldsGameCenterMatchmakingQueues** | **[]string** | the fields to include for returned resources of type gameCenterMatchmakingQueues | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterMatchmakingQueuesResponse**](GameCenterMatchmakingQueuesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingQueuesGetInstance

> GameCenterMatchmakingQueueResponse GameCenterMatchmakingQueuesGetInstance(ctx, id).FieldsGameCenterMatchmakingQueues(fieldsGameCenterMatchmakingQueues).Include(include).Execute()



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
    fieldsGameCenterMatchmakingQueues := []string{"FieldsGameCenterMatchmakingQueues_example"} // []string | the fields to include for returned resources of type gameCenterMatchmakingQueues (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesGetInstance(context.Background(), id).FieldsGameCenterMatchmakingQueues(fieldsGameCenterMatchmakingQueues).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingQueuesGetInstance`: GameCenterMatchmakingQueueResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingQueuesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterMatchmakingQueues** | **[]string** | the fields to include for returned resources of type gameCenterMatchmakingQueues | 
 **include** | **[]string** | comma-separated list of relationships to include | 

### Return type

[**GameCenterMatchmakingQueueResponse**](GameCenterMatchmakingQueueResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetrics

> GameCenterMatchmakingQueueSizesV1MetricResponse GameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetrics(ctx, id).Granularity(granularity).Limit(limit).Sort(sort).Execute()



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
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; metrics will be sorted as specified (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetrics(context.Background(), id).Granularity(granularity).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetrics`: GameCenterMatchmakingQueueSizesV1MetricResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **granularity** | **[]string** | the granularity of the per-group dataset | 
 **limit** | **int32** | maximum number of groups to return per page | 
 **sort** | **[]string** | comma-separated list of sort expressions; metrics will be sorted as specified | 

### Return type

[**GameCenterMatchmakingQueueSizesV1MetricResponse**](GameCenterMatchmakingQueueSizesV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingQueuesMatchmakingRequestsGetMetrics

> GameCenterMatchmakingQueueRequestsV1MetricResponse GameCenterMatchmakingQueuesMatchmakingRequestsGetMetrics(ctx, id).Granularity(granularity).Limit(limit).GroupBy(groupBy).FilterResult(filterResult).FilterGameCenterDetail(filterGameCenterDetail).Sort(sort).Execute()



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
    filterGameCenterDetail := "filterGameCenterDetail_example" // string | filter by 'gameCenterDetail' relationship dimension (optional)
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; metrics will be sorted as specified (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesMatchmakingRequestsGetMetrics(context.Background(), id).Granularity(granularity).Limit(limit).GroupBy(groupBy).FilterResult(filterResult).FilterGameCenterDetail(filterGameCenterDetail).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesMatchmakingRequestsGetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingQueuesMatchmakingRequestsGetMetrics`: GameCenterMatchmakingQueueRequestsV1MetricResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesMatchmakingRequestsGetMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingQueuesMatchmakingRequestsGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **granularity** | **[]string** | the granularity of the per-group dataset | 
 **limit** | **int32** | maximum number of groups to return per page | 
 **groupBy** | **[]string** | the dimension by which to group the results | 
 **filterResult** | **string** | filter by &#39;result&#39; attribute dimension | 
 **filterGameCenterDetail** | **string** | filter by &#39;gameCenterDetail&#39; relationship dimension | 
 **sort** | **[]string** | comma-separated list of sort expressions; metrics will be sorted as specified | 

### Return type

[**GameCenterMatchmakingQueueRequestsV1MetricResponse**](GameCenterMatchmakingQueueRequestsV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingQueuesMatchmakingSessionsGetMetrics

> GameCenterMatchmakingSessionsV1MetricResponse GameCenterMatchmakingQueuesMatchmakingSessionsGetMetrics(ctx, id).Granularity(granularity).Limit(limit).Sort(sort).Execute()



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
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; metrics will be sorted as specified (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesMatchmakingSessionsGetMetrics(context.Background(), id).Granularity(granularity).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesMatchmakingSessionsGetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingQueuesMatchmakingSessionsGetMetrics`: GameCenterMatchmakingSessionsV1MetricResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesMatchmakingSessionsGetMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingQueuesMatchmakingSessionsGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **granularity** | **[]string** | the granularity of the per-group dataset | 
 **limit** | **int32** | maximum number of groups to return per page | 
 **sort** | **[]string** | comma-separated list of sort expressions; metrics will be sorted as specified | 

### Return type

[**GameCenterMatchmakingSessionsV1MetricResponse**](GameCenterMatchmakingSessionsV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingQueuesUpdateInstance

> GameCenterMatchmakingQueueResponse GameCenterMatchmakingQueuesUpdateInstance(ctx, id).GameCenterMatchmakingQueueUpdateRequest(gameCenterMatchmakingQueueUpdateRequest).Execute()



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
    gameCenterMatchmakingQueueUpdateRequest := *openapiclient.NewGameCenterMatchmakingQueueUpdateRequest(*openapiclient.NewGameCenterMatchmakingQueueUpdateRequestData("Type_example", "Id_example")) // GameCenterMatchmakingQueueUpdateRequest | GameCenterMatchmakingQueue representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesUpdateInstance(context.Background(), id).GameCenterMatchmakingQueueUpdateRequest(gameCenterMatchmakingQueueUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingQueuesUpdateInstance`: GameCenterMatchmakingQueueResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingQueuesAPI.GameCenterMatchmakingQueuesUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingQueuesUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterMatchmakingQueueUpdateRequest** | [**GameCenterMatchmakingQueueUpdateRequest**](GameCenterMatchmakingQueueUpdateRequest.md) | GameCenterMatchmakingQueue representation | 

### Return type

[**GameCenterMatchmakingQueueResponse**](GameCenterMatchmakingQueueResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

