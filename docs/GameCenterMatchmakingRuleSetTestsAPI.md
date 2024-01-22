# \GameCenterMatchmakingRuleSetTestsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameCenterMatchmakingRuleSetTestsCreateInstance**](GameCenterMatchmakingRuleSetTestsAPI.md#GameCenterMatchmakingRuleSetTestsCreateInstance) | **Post** /v1/gameCenterMatchmakingRuleSetTests | 



## GameCenterMatchmakingRuleSetTestsCreateInstance

> GameCenterMatchmakingRuleSetTestResponse GameCenterMatchmakingRuleSetTestsCreateInstance(ctx).GameCenterMatchmakingRuleSetTestCreateRequest(gameCenterMatchmakingRuleSetTestCreateRequest).Execute()



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
    gameCenterMatchmakingRuleSetTestCreateRequest := *openapiclient.NewGameCenterMatchmakingRuleSetTestCreateRequest(*openapiclient.NewGameCenterMatchmakingRuleSetTestCreateRequestData("Type_example", *openapiclient.NewGameCenterMatchmakingRuleSetTestCreateRequestDataRelationships(*openapiclient.NewGameCenterMatchmakingQueueCreateRequestDataRelationshipsRuleSet(*openapiclient.NewGameCenterMatchmakingQueueRelationshipsRuleSetData("Type_example", "Id_example")), *openapiclient.NewGameCenterMatchmakingRuleSetTestCreateRequestDataRelationshipsMatchmakingRequests([]openapiclient.GameCenterMatchmakingRuleSetTestCreateRequestDataRelationshipsMatchmakingRequestsDataInner{*openapiclient.NewGameCenterMatchmakingRuleSetTestCreateRequestDataRelationshipsMatchmakingRequestsDataInner("Type_example", "Id_example")})))) // GameCenterMatchmakingRuleSetTestCreateRequest | GameCenterMatchmakingRuleSetTest representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GameCenterMatchmakingRuleSetTestsAPI.GameCenterMatchmakingRuleSetTestsCreateInstance(context.Background()).GameCenterMatchmakingRuleSetTestCreateRequest(gameCenterMatchmakingRuleSetTestCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GameCenterMatchmakingRuleSetTestsAPI.GameCenterMatchmakingRuleSetTestsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GameCenterMatchmakingRuleSetTestsCreateInstance`: GameCenterMatchmakingRuleSetTestResponse
    fmt.Fprintf(os.Stdout, "Response from `GameCenterMatchmakingRuleSetTestsAPI.GameCenterMatchmakingRuleSetTestsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGameCenterMatchmakingRuleSetTestsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gameCenterMatchmakingRuleSetTestCreateRequest** | [**GameCenterMatchmakingRuleSetTestCreateRequest**](GameCenterMatchmakingRuleSetTestCreateRequest.md) | GameCenterMatchmakingRuleSetTest representation | 

### Return type

[**GameCenterMatchmakingRuleSetTestResponse**](GameCenterMatchmakingRuleSetTestResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

