# \GameCenterMatchmakingTeamsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterMatchmakingTeamsCreateInstance**](GameCenterMatchmakingTeamsAPI.md#GameCenterMatchmakingTeamsCreateInstance) | **Post** /v1/gameCenterMatchmakingTeams | 
[**GameCenterMatchmakingTeamsDeleteInstance**](GameCenterMatchmakingTeamsAPI.md#GameCenterMatchmakingTeamsDeleteInstance) | **Delete** /v1/gameCenterMatchmakingTeams/{id} | 
[**GameCenterMatchmakingTeamsUpdateInstance**](GameCenterMatchmakingTeamsAPI.md#GameCenterMatchmakingTeamsUpdateInstance) | **Patch** /v1/gameCenterMatchmakingTeams/{id} | 



## GameCenterMatchmakingTeamsCreateInstance

> GameCenterMatchmakingTeamResponse GameCenterMatchmakingTeamsCreateInstance(ctx).GameCenterMatchmakingTeamCreateRequest(gameCenterMatchmakingTeamCreateRequest).Execute()



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
    gameCenterMatchmakingTeamCreateRequest := *openapiclient.NewGameCenterMatchmakingTeamCreateRequest(*openapiclient.NewGameCenterMatchmakingTeamCreateRequestData("Type_example", *openapiclient.NewGameCenterMatchmakingTeamCreateRequestDataAttributes("ReferenceName_example", int32(123), int32(123)), *openapiclient.NewGameCenterMatchmakingRuleCreateRequestDataRelationships(*openapiclient.NewGameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet(*openapiclient.NewGameCenterMatchmakingQueueRelationshipsRuleSetData("Type_example", "Id_example"))))) // GameCenterMatchmakingTeamCreateRequest | GameCenterMatchmakingTeam representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingTeamsAPI.GameCenterMatchmakingTeamsCreateInstance(context.Background()).GameCenterMatchmakingTeamCreateRequest(gameCenterMatchmakingTeamCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingTeamsAPI.GameCenterMatchmakingTeamsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingTeamsCreateInstance`: GameCenterMatchmakingTeamResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingTeamsAPI.GameCenterMatchmakingTeamsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingTeamsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameCenterMatchmakingTeamCreateRequest** | [**GameCenterMatchmakingTeamCreateRequest**](GameCenterMatchmakingTeamCreateRequest.md) | GameCenterMatchmakingTeam representation | 

### Return type

[**GameCenterMatchmakingTeamResponse**](GameCenterMatchmakingTeamResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GameCenterMatchmakingTeamsDeleteInstance

> GameCenterMatchmakingTeamsDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.GameCenterMatchmakingTeamsAPI.GameCenterMatchmakingTeamsDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingTeamsAPI.GameCenterMatchmakingTeamsDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGameCenterMatchmakingTeamsDeleteInstanceRequest struct via the builder pattern


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


## GameCenterMatchmakingTeamsUpdateInstance

> GameCenterMatchmakingTeamResponse GameCenterMatchmakingTeamsUpdateInstance(ctx, id).GameCenterMatchmakingTeamUpdateRequest(gameCenterMatchmakingTeamUpdateRequest).Execute()



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
    gameCenterMatchmakingTeamUpdateRequest := *openapiclient.NewGameCenterMatchmakingTeamUpdateRequest(*openapiclient.NewGameCenterMatchmakingTeamUpdateRequestData("Type_example", "Id_example")) // GameCenterMatchmakingTeamUpdateRequest | GameCenterMatchmakingTeam representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingTeamsAPI.GameCenterMatchmakingTeamsUpdateInstance(context.Background(), id).GameCenterMatchmakingTeamUpdateRequest(gameCenterMatchmakingTeamUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingTeamsAPI.GameCenterMatchmakingTeamsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingTeamsUpdateInstance`: GameCenterMatchmakingTeamResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingTeamsAPI.GameCenterMatchmakingTeamsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingTeamsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gameCenterMatchmakingTeamUpdateRequest** | [**GameCenterMatchmakingTeamUpdateRequest**](GameCenterMatchmakingTeamUpdateRequest.md) | GameCenterMatchmakingTeam representation | 

### Return type

[**GameCenterMatchmakingTeamResponse**](GameCenterMatchmakingTeamResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

