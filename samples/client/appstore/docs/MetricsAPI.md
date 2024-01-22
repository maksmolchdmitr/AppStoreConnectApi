# \MetricsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsBetaTesterUsagesGetMetrics**](MetricsAPI.md#AppsBetaTesterUsagesGetMetrics) | **Get** /v1/apps/{id}/metrics/betaTesterUsages | 
[**BetaGroupsBetaTesterUsagesGetMetrics**](MetricsAPI.md#BetaGroupsBetaTesterUsagesGetMetrics) | **Get** /v1/betaGroups/{id}/metrics/betaTesterUsages | 
[**BetaTestersBetaTesterUsagesGetMetrics**](MetricsAPI.md#BetaTestersBetaTesterUsagesGetMetrics) | **Get** /v1/betaTesters/{id}/metrics/betaTesterUsages | 
[**BuildsBetaBuildUsagesGetMetrics**](MetricsAPI.md#BuildsBetaBuildUsagesGetMetrics) | **Get** /v1/builds/{id}/metrics/betaBuildUsages | 
[**GameCenterDetailsClassicMatchmakingRequestsGetMetrics**](MetricsAPI.md#GameCenterDetailsClassicMatchmakingRequestsGetMetrics) | **Get** /v1/gameCenterDetails/{id}/metrics/classicMatchmakingRequests | 
[**GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics**](MetricsAPI.md#GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics) | **Get** /v1/gameCenterDetails/{id}/metrics/ruleBasedMatchmakingRequests | 
[**GameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetrics**](MetricsAPI.md#GameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetrics) | **Get** /v1/gameCenterMatchmakingQueues/{id}/metrics/experimentMatchmakingQueueSizes | 
[**GameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetrics**](MetricsAPI.md#GameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetrics) | **Get** /v1/gameCenterMatchmakingQueues/{id}/metrics/experimentMatchmakingRequests | 
[**GameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetrics**](MetricsAPI.md#GameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetrics) | **Get** /v1/gameCenterMatchmakingQueues/{id}/metrics/matchmakingQueueSizes | 
[**GameCenterMatchmakingQueuesMatchmakingRequestsGetMetrics**](MetricsAPI.md#GameCenterMatchmakingQueuesMatchmakingRequestsGetMetrics) | **Get** /v1/gameCenterMatchmakingQueues/{id}/metrics/matchmakingRequests | 
[**GameCenterMatchmakingQueuesMatchmakingSessionsGetMetrics**](MetricsAPI.md#GameCenterMatchmakingQueuesMatchmakingSessionsGetMetrics) | **Get** /v1/gameCenterMatchmakingQueues/{id}/metrics/matchmakingSessions | 
[**GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics**](MetricsAPI.md#GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics) | **Get** /v1/gameCenterMatchmakingRules/{id}/metrics/matchmakingBooleanRuleResults | 
[**GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics**](MetricsAPI.md#GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics) | **Get** /v1/gameCenterMatchmakingRules/{id}/metrics/matchmakingNumberRuleResults | 
[**GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics**](MetricsAPI.md#GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics) | **Get** /v1/gameCenterMatchmakingRules/{id}/metrics/matchmakingRuleErrors | 



## AppsBetaTesterUsagesGetMetrics

> AppsBetaTesterUsagesV1MetricResponse AppsBetaTesterUsagesGetMetrics(ctx, id).Limit(limit).GroupBy(groupBy).FilterBetaTesters(filterBetaTesters).Period(period).Execute()



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
    limit := int32(56) // int32 | maximum number of groups to return per page (optional)
    groupBy := []string{"GroupBy_example"} // []string | the dimension by which to group the results (optional)
    filterBetaTesters := "filterBetaTesters_example" // string | filter by 'betaTesters' relationship dimension (optional)
    period := "PT10M" // string | the duration of the reporting period (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsAPI.AppsBetaTesterUsagesGetMetrics(context.Background(), id).Limit(limit).GroupBy(groupBy).FilterBetaTesters(filterBetaTesters).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.AppsBetaTesterUsagesGetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppsBetaTesterUsagesGetMetrics`: AppsBetaTesterUsagesV1MetricResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.AppsBetaTesterUsagesGetMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsBetaTesterUsagesGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum number of groups to return per page | 
 **groupBy** | **[]string** | the dimension by which to group the results | 
 **filterBetaTesters** | **string** | filter by &#39;betaTesters&#39; relationship dimension | 
 **period** | **string** | the duration of the reporting period | 

### Return type

[**AppsBetaTesterUsagesV1MetricResponse**](AppsBetaTesterUsagesV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaGroupsBetaTesterUsagesGetMetrics

> AppsBetaTesterUsagesV1MetricResponse BetaGroupsBetaTesterUsagesGetMetrics(ctx, id).Limit(limit).GroupBy(groupBy).FilterBetaTesters(filterBetaTesters).Period(period).Execute()



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
    limit := int32(56) // int32 | maximum number of groups to return per page (optional)
    groupBy := []string{"GroupBy_example"} // []string | the dimension by which to group the results (optional)
    filterBetaTesters := "filterBetaTesters_example" // string | filter by 'betaTesters' relationship dimension (optional)
    period := "PT10M" // string | the duration of the reporting period (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsAPI.BetaGroupsBetaTesterUsagesGetMetrics(context.Background(), id).Limit(limit).GroupBy(groupBy).FilterBetaTesters(filterBetaTesters).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.BetaGroupsBetaTesterUsagesGetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaGroupsBetaTesterUsagesGetMetrics`: AppsBetaTesterUsagesV1MetricResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.BetaGroupsBetaTesterUsagesGetMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaGroupsBetaTesterUsagesGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum number of groups to return per page | 
 **groupBy** | **[]string** | the dimension by which to group the results | 
 **filterBetaTesters** | **string** | filter by &#39;betaTesters&#39; relationship dimension | 
 **period** | **string** | the duration of the reporting period | 

### Return type

[**AppsBetaTesterUsagesV1MetricResponse**](AppsBetaTesterUsagesV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BetaTestersBetaTesterUsagesGetMetrics

> BetaTesterUsagesV1MetricResponse BetaTestersBetaTesterUsagesGetMetrics(ctx, id).Limit(limit).FilterApps(filterApps).Period(period).Execute()



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
    limit := int32(56) // int32 | maximum number of groups to return per page (optional)
    filterApps := "filterApps_example" // string | filter by 'apps' relationship dimension (optional)
    period := "PT10M" // string | the duration of the reporting period (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsAPI.BetaTestersBetaTesterUsagesGetMetrics(context.Background(), id).Limit(limit).FilterApps(filterApps).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.BetaTestersBetaTesterUsagesGetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BetaTestersBetaTesterUsagesGetMetrics`: BetaTesterUsagesV1MetricResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.BetaTestersBetaTesterUsagesGetMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBetaTestersBetaTesterUsagesGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum number of groups to return per page | 
 **filterApps** | **string** | filter by &#39;apps&#39; relationship dimension | 
 **period** | **string** | the duration of the reporting period | 

### Return type

[**BetaTesterUsagesV1MetricResponse**](BetaTesterUsagesV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsBetaBuildUsagesGetMetrics

> BetaBuildUsagesV1MetricResponse BuildsBetaBuildUsagesGetMetrics(ctx, id).Limit(limit).Execute()



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
    limit := int32(56) // int32 | maximum number of groups to return per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsAPI.BuildsBetaBuildUsagesGetMetrics(context.Background(), id).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.BuildsBetaBuildUsagesGetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BuildsBetaBuildUsagesGetMetrics`: BetaBuildUsagesV1MetricResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.BuildsBetaBuildUsagesGetMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsBetaBuildUsagesGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum number of groups to return per page | 

### Return type

[**BetaBuildUsagesV1MetricResponse**](BetaBuildUsagesV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterDetailsClassicMatchmakingRequestsGetMetrics

> GameCenterMatchmakingAppRequestsV1MetricResponse GameCenterDetailsClassicMatchmakingRequestsGetMetrics(ctx, id).Granularity(granularity).Limit(limit).GroupBy(groupBy).FilterResult(filterResult).Sort(sort).Execute()



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
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; metrics will be sorted as specified (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsAPI.GameCenterDetailsClassicMatchmakingRequestsGetMetrics(context.Background(), id).Granularity(granularity).Limit(limit).GroupBy(groupBy).FilterResult(filterResult).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GameCenterDetailsClassicMatchmakingRequestsGetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterDetailsClassicMatchmakingRequestsGetMetrics`: GameCenterMatchmakingAppRequestsV1MetricResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GameCenterDetailsClassicMatchmakingRequestsGetMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterDetailsClassicMatchmakingRequestsGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **granularity** | **[]string** | the granularity of the per-group dataset | 
 **limit** | **int32** | maximum number of groups to return per page | 
 **groupBy** | **[]string** | the dimension by which to group the results | 
 **filterResult** | **string** | filter by &#39;result&#39; attribute dimension | 
 **sort** | **[]string** | comma-separated list of sort expressions; metrics will be sorted as specified | 

### Return type

[**GameCenterMatchmakingAppRequestsV1MetricResponse**](GameCenterMatchmakingAppRequestsV1MetricResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics

> GameCenterMatchmakingAppRequestsV1MetricResponse GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics(ctx, id).Granularity(granularity).Limit(limit).GroupBy(groupBy).FilterResult(filterResult).Sort(sort).Execute()



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
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; metrics will be sorted as specified (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsAPI.GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics(context.Background(), id).Granularity(granularity).Limit(limit).GroupBy(groupBy).FilterResult(filterResult).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics`: GameCenterMatchmakingAppRequestsV1MetricResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GameCenterDetailsRuleBasedMatchmakingRequestsGetMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterDetailsRuleBasedMatchmakingRequestsGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **granularity** | **[]string** | the granularity of the per-group dataset | 
 **limit** | **int32** | maximum number of groups to return per page | 
 **groupBy** | **[]string** | the dimension by which to group the results | 
 **filterResult** | **string** | filter by &#39;result&#39; attribute dimension | 
 **sort** | **[]string** | comma-separated list of sort expressions; metrics will be sorted as specified | 

### Return type

[**GameCenterMatchmakingAppRequestsV1MetricResponse**](GameCenterMatchmakingAppRequestsV1MetricResponse.md)

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
    resp, r, err := apiClient.MetricsAPI.GameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetrics(context.Background(), id).Granularity(granularity).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetrics`: GameCenterMatchmakingQueueSizesV1MetricResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GameCenterMatchmakingQueuesExperimentMatchmakingQueueSizesGetMetrics`: %v\n", resp)
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
    resp, r, err := apiClient.MetricsAPI.GameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetrics(context.Background(), id).Granularity(granularity).Limit(limit).GroupBy(groupBy).FilterResult(filterResult).FilterGameCenterDetail(filterGameCenterDetail).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetrics`: GameCenterMatchmakingQueueRequestsV1MetricResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GameCenterMatchmakingQueuesExperimentMatchmakingRequestsGetMetrics`: %v\n", resp)
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
    resp, r, err := apiClient.MetricsAPI.GameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetrics(context.Background(), id).Granularity(granularity).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetrics`: GameCenterMatchmakingQueueSizesV1MetricResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GameCenterMatchmakingQueuesMatchmakingQueueSizesGetMetrics`: %v\n", resp)
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
    resp, r, err := apiClient.MetricsAPI.GameCenterMatchmakingQueuesMatchmakingRequestsGetMetrics(context.Background(), id).Granularity(granularity).Limit(limit).GroupBy(groupBy).FilterResult(filterResult).FilterGameCenterDetail(filterGameCenterDetail).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GameCenterMatchmakingQueuesMatchmakingRequestsGetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingQueuesMatchmakingRequestsGetMetrics`: GameCenterMatchmakingQueueRequestsV1MetricResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GameCenterMatchmakingQueuesMatchmakingRequestsGetMetrics`: %v\n", resp)
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
    resp, r, err := apiClient.MetricsAPI.GameCenterMatchmakingQueuesMatchmakingSessionsGetMetrics(context.Background(), id).Granularity(granularity).Limit(limit).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GameCenterMatchmakingQueuesMatchmakingSessionsGetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingQueuesMatchmakingSessionsGetMetrics`: GameCenterMatchmakingSessionsV1MetricResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GameCenterMatchmakingQueuesMatchmakingSessionsGetMetrics`: %v\n", resp)
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
    resp, r, err := apiClient.MetricsAPI.GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics(context.Background(), id).Granularity(granularity).Limit(limit).GroupBy(groupBy).FilterResult(filterResult).FilterGameCenterMatchmakingQueue(filterGameCenterMatchmakingQueue).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics`: GameCenterMatchmakingBooleanRuleResultsV1MetricResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GameCenterMatchmakingRulesMatchmakingBooleanRuleResultsGetMetrics`: %v\n", resp)
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
    resp, r, err := apiClient.MetricsAPI.GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics(context.Background(), id).Granularity(granularity).Limit(limit).GroupBy(groupBy).FilterGameCenterMatchmakingQueue(filterGameCenterMatchmakingQueue).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics`: GameCenterMatchmakingNumberRuleResultsV1MetricResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GameCenterMatchmakingRulesMatchmakingNumberRuleResultsGetMetrics`: %v\n", resp)
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
    resp, r, err := apiClient.MetricsAPI.GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics(context.Background(), id).Granularity(granularity).Limit(limit).GroupBy(groupBy).FilterGameCenterMatchmakingQueue(filterGameCenterMatchmakingQueue).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsAPI.GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics`: GameCenterMatchmakingRuleErrorsV1MetricResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricsAPI.GameCenterMatchmakingRulesMatchmakingRuleErrorsGetMetrics`: %v\n", resp)
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

