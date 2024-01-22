# \GameCenterMatchmakingRuleSetsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterMatchmakingRuleSetsCreateInstance**](GameCenterMatchmakingRuleSetsAPI.md#GameCenterMatchmakingRuleSetsCreateInstance) | **Post** /v1/gameCenterMatchmakingRuleSets | 
[**GameCenterMatchmakingRuleSetsDeleteInstance**](GameCenterMatchmakingRuleSetsAPI.md#GameCenterMatchmakingRuleSetsDeleteInstance) | **Delete** /v1/gameCenterMatchmakingRuleSets/{id} | 
[**GameCenterMatchmakingRuleSetsGetCollection**](GameCenterMatchmakingRuleSetsAPI.md#GameCenterMatchmakingRuleSetsGetCollection) | **Get** /v1/gameCenterMatchmakingRuleSets | 
[**GameCenterMatchmakingRuleSetsGetInstance**](GameCenterMatchmakingRuleSetsAPI.md#GameCenterMatchmakingRuleSetsGetInstance) | **Get** /v1/gameCenterMatchmakingRuleSets/{id} | 
[**GameCenterMatchmakingRuleSetsMatchmakingQueuesGetToManyRelated**](GameCenterMatchmakingRuleSetsAPI.md#GameCenterMatchmakingRuleSetsMatchmakingQueuesGetToManyRelated) | **Get** /v1/gameCenterMatchmakingRuleSets/{id}/matchmakingQueues | 
[**GameCenterMatchmakingRuleSetsRulesGetToManyRelated**](GameCenterMatchmakingRuleSetsAPI.md#GameCenterMatchmakingRuleSetsRulesGetToManyRelated) | **Get** /v1/gameCenterMatchmakingRuleSets/{id}/rules | 
[**GameCenterMatchmakingRuleSetsTeamsGetToManyRelated**](GameCenterMatchmakingRuleSetsAPI.md#GameCenterMatchmakingRuleSetsTeamsGetToManyRelated) | **Get** /v1/gameCenterMatchmakingRuleSets/{id}/teams | 
[**GameCenterMatchmakingRuleSetsUpdateInstance**](GameCenterMatchmakingRuleSetsAPI.md#GameCenterMatchmakingRuleSetsUpdateInstance) | **Patch** /v1/gameCenterMatchmakingRuleSets/{id} | 



## GameCenterMatchmakingRuleSetsCreateInstance

> GameCenterMatchmakingRuleSetResponse GameCenterMatchmakingRuleSetsCreateInstance(ctx).GameCenterMatchmakingRuleSetCreateRequest(gameCenterMatchmakingRuleSetCreateRequest).Execute()



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
    gameCenterMatchmakingRuleSetCreateRequest := *openapiclient.NewGameCenterMatchmakingRuleSetCreateRequest(*openapiclient.NewGameCenterMatchmakingRuleSetCreateRequestData("Type_example", *openapiclient.NewGameCenterMatchmakingRuleSetCreateRequestDataAttributes("ReferenceName_example", int32(123), int32(123), int32(123)))) // GameCenterMatchmakingRuleSetCreateRequest | GameCenterMatchmakingRuleSet representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsCreateInstance(context.Background()).GameCenterMatchmakingRuleSetCreateRequest(gameCenterMatchmakingRuleSetCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingRuleSetsCreateInstance`: GameCenterMatchmakingRuleSetResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingRuleSetsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameCenterMatchmakingRuleSetCreateRequest** | [**GameCenterMatchmakingRuleSetCreateRequest**](GameCenterMatchmakingRuleSetCreateRequest.md) | GameCenterMatchmakingRuleSet representation | 

### Return type

[**GameCenterMatchmakingRuleSetResponse**](GameCenterMatchmakingRuleSetResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingRuleSetsDeleteInstance

> GameCenterMatchmakingRuleSetsDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterMatchmakingRuleSetsDeleteInstanceRequest struct via the builder pattern


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


## GameCenterMatchmakingRuleSetsGetCollection

> GameCenterMatchmakingRuleSetsResponse GameCenterMatchmakingRuleSetsGetCollection(ctx).FieldsGameCenterMatchmakingRuleSets(fieldsGameCenterMatchmakingRuleSets).Limit(limit).Include(include).FieldsGameCenterMatchmakingQueues(fieldsGameCenterMatchmakingQueues).FieldsGameCenterMatchmakingTeams(fieldsGameCenterMatchmakingTeams).FieldsGameCenterMatchmakingRules(fieldsGameCenterMatchmakingRules).LimitMatchmakingQueues(limitMatchmakingQueues).LimitRules(limitRules).LimitTeams(limitTeams).Execute()



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
    fieldsGameCenterMatchmakingRuleSets := []string{"FieldsGameCenterMatchmakingRuleSets_example"} // []string | the fields to include for returned resources of type gameCenterMatchmakingRuleSets (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsGameCenterMatchmakingQueues := []string{"FieldsGameCenterMatchmakingQueues_example"} // []string | the fields to include for returned resources of type gameCenterMatchmakingQueues (optional)
    fieldsGameCenterMatchmakingTeams := []string{"FieldsGameCenterMatchmakingTeams_example"} // []string | the fields to include for returned resources of type gameCenterMatchmakingTeams (optional)
    fieldsGameCenterMatchmakingRules := []string{"FieldsGameCenterMatchmakingRules_example"} // []string | the fields to include for returned resources of type gameCenterMatchmakingRules (optional)
    limitMatchmakingQueues := int32(56) // int32 | maximum number of related matchmakingQueues returned (when they are included) (optional)
    limitRules := int32(56) // int32 | maximum number of related rules returned (when they are included) (optional)
    limitTeams := int32(56) // int32 | maximum number of related teams returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsGetCollection(context.Background()).FieldsGameCenterMatchmakingRuleSets(fieldsGameCenterMatchmakingRuleSets).Limit(limit).Include(include).FieldsGameCenterMatchmakingQueues(fieldsGameCenterMatchmakingQueues).FieldsGameCenterMatchmakingTeams(fieldsGameCenterMatchmakingTeams).FieldsGameCenterMatchmakingRules(fieldsGameCenterMatchmakingRules).LimitMatchmakingQueues(limitMatchmakingQueues).LimitRules(limitRules).LimitTeams(limitTeams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingRuleSetsGetCollection`: GameCenterMatchmakingRuleSetsResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingRuleSetsGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fieldsGameCenterMatchmakingRuleSets** | **[]string** | the fields to include for returned resources of type gameCenterMatchmakingRuleSets | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsGameCenterMatchmakingQueues** | **[]string** | the fields to include for returned resources of type gameCenterMatchmakingQueues | 
 **fieldsGameCenterMatchmakingTeams** | **[]string** | the fields to include for returned resources of type gameCenterMatchmakingTeams | 
 **fieldsGameCenterMatchmakingRules** | **[]string** | the fields to include for returned resources of type gameCenterMatchmakingRules | 
 **limitMatchmakingQueues** | **int32** | maximum number of related matchmakingQueues returned (when they are included) | 
 **limitRules** | **int32** | maximum number of related rules returned (when they are included) | 
 **limitTeams** | **int32** | maximum number of related teams returned (when they are included) | 

### Return type

[**GameCenterMatchmakingRuleSetsResponse**](GameCenterMatchmakingRuleSetsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingRuleSetsGetInstance

> GameCenterMatchmakingRuleSetResponse GameCenterMatchmakingRuleSetsGetInstance(ctx, id).FieldsGameCenterMatchmakingRuleSets(fieldsGameCenterMatchmakingRuleSets).Include(include).FieldsGameCenterMatchmakingQueues(fieldsGameCenterMatchmakingQueues).FieldsGameCenterMatchmakingTeams(fieldsGameCenterMatchmakingTeams).FieldsGameCenterMatchmakingRules(fieldsGameCenterMatchmakingRules).LimitMatchmakingQueues(limitMatchmakingQueues).LimitRules(limitRules).LimitTeams(limitTeams).Execute()



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
    fieldsGameCenterMatchmakingRuleSets := []string{"FieldsGameCenterMatchmakingRuleSets_example"} // []string | the fields to include for returned resources of type gameCenterMatchmakingRuleSets (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsGameCenterMatchmakingQueues := []string{"FieldsGameCenterMatchmakingQueues_example"} // []string | the fields to include for returned resources of type gameCenterMatchmakingQueues (optional)
    fieldsGameCenterMatchmakingTeams := []string{"FieldsGameCenterMatchmakingTeams_example"} // []string | the fields to include for returned resources of type gameCenterMatchmakingTeams (optional)
    fieldsGameCenterMatchmakingRules := []string{"FieldsGameCenterMatchmakingRules_example"} // []string | the fields to include for returned resources of type gameCenterMatchmakingRules (optional)
    limitMatchmakingQueues := int32(56) // int32 | maximum number of related matchmakingQueues returned (when they are included) (optional)
    limitRules := int32(56) // int32 | maximum number of related rules returned (when they are included) (optional)
    limitTeams := int32(56) // int32 | maximum number of related teams returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsGetInstance(context.Background(), id).FieldsGameCenterMatchmakingRuleSets(fieldsGameCenterMatchmakingRuleSets).Include(include).FieldsGameCenterMatchmakingQueues(fieldsGameCenterMatchmakingQueues).FieldsGameCenterMatchmakingTeams(fieldsGameCenterMatchmakingTeams).FieldsGameCenterMatchmakingRules(fieldsGameCenterMatchmakingRules).LimitMatchmakingQueues(limitMatchmakingQueues).LimitRules(limitRules).LimitTeams(limitTeams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingRuleSetsGetInstance`: GameCenterMatchmakingRuleSetResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingRuleSetsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterMatchmakingRuleSets** | **[]string** | the fields to include for returned resources of type gameCenterMatchmakingRuleSets | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsGameCenterMatchmakingQueues** | **[]string** | the fields to include for returned resources of type gameCenterMatchmakingQueues | 
 **fieldsGameCenterMatchmakingTeams** | **[]string** | the fields to include for returned resources of type gameCenterMatchmakingTeams | 
 **fieldsGameCenterMatchmakingRules** | **[]string** | the fields to include for returned resources of type gameCenterMatchmakingRules | 
 **limitMatchmakingQueues** | **int32** | maximum number of related matchmakingQueues returned (when they are included) | 
 **limitRules** | **int32** | maximum number of related rules returned (when they are included) | 
 **limitTeams** | **int32** | maximum number of related teams returned (when they are included) | 

### Return type

[**GameCenterMatchmakingRuleSetResponse**](GameCenterMatchmakingRuleSetResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingRuleSetsMatchmakingQueuesGetToManyRelated

> GameCenterMatchmakingQueuesResponse GameCenterMatchmakingRuleSetsMatchmakingQueuesGetToManyRelated(ctx, id).FieldsGameCenterMatchmakingQueues(fieldsGameCenterMatchmakingQueues).FieldsGameCenterMatchmakingRuleSets(fieldsGameCenterMatchmakingRuleSets).Limit(limit).Include(include).Execute()



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
    fieldsGameCenterMatchmakingRuleSets := []string{"FieldsGameCenterMatchmakingRuleSets_example"} // []string | the fields to include for returned resources of type gameCenterMatchmakingRuleSets (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsMatchmakingQueuesGetToManyRelated(context.Background(), id).FieldsGameCenterMatchmakingQueues(fieldsGameCenterMatchmakingQueues).FieldsGameCenterMatchmakingRuleSets(fieldsGameCenterMatchmakingRuleSets).Limit(limit).Include(include).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsMatchmakingQueuesGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingRuleSetsMatchmakingQueuesGetToManyRelated`: GameCenterMatchmakingQueuesResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsMatchmakingQueuesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingRuleSetsMatchmakingQueuesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterMatchmakingQueues** | **[]string** | the fields to include for returned resources of type gameCenterMatchmakingQueues | 
 **fieldsGameCenterMatchmakingRuleSets** | **[]string** | the fields to include for returned resources of type gameCenterMatchmakingRuleSets | 
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


## GameCenterMatchmakingRuleSetsRulesGetToManyRelated

> GameCenterMatchmakingRulesResponse GameCenterMatchmakingRuleSetsRulesGetToManyRelated(ctx, id).FieldsGameCenterMatchmakingRules(fieldsGameCenterMatchmakingRules).Limit(limit).Execute()



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
    fieldsGameCenterMatchmakingRules := []string{"FieldsGameCenterMatchmakingRules_example"} // []string | the fields to include for returned resources of type gameCenterMatchmakingRules (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsRulesGetToManyRelated(context.Background(), id).FieldsGameCenterMatchmakingRules(fieldsGameCenterMatchmakingRules).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsRulesGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingRuleSetsRulesGetToManyRelated`: GameCenterMatchmakingRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsRulesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingRuleSetsRulesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterMatchmakingRules** | **[]string** | the fields to include for returned resources of type gameCenterMatchmakingRules | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**GameCenterMatchmakingRulesResponse**](GameCenterMatchmakingRulesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingRuleSetsTeamsGetToManyRelated

> GameCenterMatchmakingTeamsResponse GameCenterMatchmakingRuleSetsTeamsGetToManyRelated(ctx, id).FieldsGameCenterMatchmakingTeams(fieldsGameCenterMatchmakingTeams).Limit(limit).Execute()



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
    fieldsGameCenterMatchmakingTeams := []string{"FieldsGameCenterMatchmakingTeams_example"} // []string | the fields to include for returned resources of type gameCenterMatchmakingTeams (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsTeamsGetToManyRelated(context.Background(), id).FieldsGameCenterMatchmakingTeams(fieldsGameCenterMatchmakingTeams).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsTeamsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingRuleSetsTeamsGetToManyRelated`: GameCenterMatchmakingTeamsResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsTeamsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingRuleSetsTeamsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsGameCenterMatchmakingTeams** | **[]string** | the fields to include for returned resources of type gameCenterMatchmakingTeams | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**GameCenterMatchmakingTeamsResponse**](GameCenterMatchmakingTeamsResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingRuleSetsUpdateInstance

> GameCenterMatchmakingRuleSetResponse GameCenterMatchmakingRuleSetsUpdateInstance(ctx, id).GameCenterMatchmakingRuleSetUpdateRequest(gameCenterMatchmakingRuleSetUpdateRequest).Execute()



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
    gameCenterMatchmakingRuleSetUpdateRequest := *openapiclient.NewGameCenterMatchmakingRuleSetUpdateRequest(*openapiclient.NewGameCenterMatchmakingRuleSetUpdateRequestData("Type_example", "Id_example")) // GameCenterMatchmakingRuleSetUpdateRequest | GameCenterMatchmakingRuleSet representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsUpdateInstance(context.Background(), id).GameCenterMatchmakingRuleSetUpdateRequest(gameCenterMatchmakingRuleSetUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingRuleSetsUpdateInstance`: GameCenterMatchmakingRuleSetResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingRuleSetsAPI.GameCenterMatchmakingRuleSetsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingRuleSetsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterMatchmakingRuleSetUpdateRequest** | [**GameCenterMatchmakingRuleSetUpdateRequest**](GameCenterMatchmakingRuleSetUpdateRequest.md) | GameCenterMatchmakingRuleSet representation | 

### Return type

[**GameCenterMatchmakingRuleSetResponse**](GameCenterMatchmakingRuleSetResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

