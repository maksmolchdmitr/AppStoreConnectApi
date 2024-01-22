# \UsersAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersDeleteInstance**](UsersAPI.md#UsersDeleteInstance) | **Delete** /v1/users/{id} | 
[**UsersGetCollection**](UsersAPI.md#UsersGetCollection) | **Get** /v1/users | 
[**UsersGetInstance**](UsersAPI.md#UsersGetInstance) | **Get** /v1/users/{id} | 
[**UsersUpdateInstance**](UsersAPI.md#UsersUpdateInstance) | **Patch** /v1/users/{id} | 
[**UsersVisibleAppsCreateToManyRelationship**](UsersAPI.md#UsersVisibleAppsCreateToManyRelationship) | **Post** /v1/users/{id}/relationships/visibleApps | 
[**UsersVisibleAppsDeleteToManyRelationship**](UsersAPI.md#UsersVisibleAppsDeleteToManyRelationship) | **Delete** /v1/users/{id}/relationships/visibleApps | 
[**UsersVisibleAppsGetToManyRelated**](UsersAPI.md#UsersVisibleAppsGetToManyRelated) | **Get** /v1/users/{id}/visibleApps | 
[**UsersVisibleAppsGetToManyRelationship**](UsersAPI.md#UsersVisibleAppsGetToManyRelationship) | **Get** /v1/users/{id}/relationships/visibleApps | 
[**UsersVisibleAppsReplaceToManyRelationship**](UsersAPI.md#UsersVisibleAppsReplaceToManyRelationship) | **Patch** /v1/users/{id}/relationships/visibleApps | 



## UsersDeleteInstance

> UsersDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.UsersAPI.UsersDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUsersDeleteInstanceRequest struct via the builder pattern


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


## UsersGetCollection

> UsersResponse UsersGetCollection(ctx).FilterRoles(filterRoles).FilterUsername(filterUsername).FilterVisibleApps(filterVisibleApps).Sort(sort).FieldsUsers(fieldsUsers).Limit(limit).Include(include).FieldsApps(fieldsApps).LimitVisibleApps(limitVisibleApps).Execute()



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
    filterRoles := []string{"FilterRoles_example"} // []string | filter by attribute 'roles' (optional)
    filterUsername := []string{"Inner_example"} // []string | filter by attribute 'username' (optional)
    filterVisibleApps := []string{"Inner_example"} // []string | filter by id(s) of related 'visibleApps' (optional)
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; resources will be sorted as specified (optional)
    fieldsUsers := []string{"FieldsUsers_example"} // []string | the fields to include for returned resources of type users (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limitVisibleApps := int32(56) // int32 | maximum number of related visibleApps returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.UsersGetCollection(context.Background()).FilterRoles(filterRoles).FilterUsername(filterUsername).FilterVisibleApps(filterVisibleApps).Sort(sort).FieldsUsers(fieldsUsers).Limit(limit).Include(include).FieldsApps(fieldsApps).LimitVisibleApps(limitVisibleApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetCollection`: UsersResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterRoles** | **[]string** | filter by attribute &#39;roles&#39; | 
 **filterUsername** | **[]string** | filter by attribute &#39;username&#39; | 
 **filterVisibleApps** | **[]string** | filter by id(s) of related &#39;visibleApps&#39; | 
 **sort** | **[]string** | comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsUsers** | **[]string** | the fields to include for returned resources of type users | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limitVisibleApps** | **int32** | maximum number of related visibleApps returned (when they are included) | 

### Return type

[**UsersResponse**](UsersResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersGetInstance

> UserResponse UsersGetInstance(ctx, id).FieldsUsers(fieldsUsers).Include(include).FieldsApps(fieldsApps).LimitVisibleApps(limitVisibleApps).Execute()



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
    fieldsUsers := []string{"FieldsUsers_example"} // []string | the fields to include for returned resources of type users (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limitVisibleApps := int32(56) // int32 | maximum number of related visibleApps returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.UsersGetInstance(context.Background(), id).FieldsUsers(fieldsUsers).Include(include).FieldsApps(fieldsApps).LimitVisibleApps(limitVisibleApps).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersGetInstance`: UserResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsUsers** | **[]string** | the fields to include for returned resources of type users | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limitVisibleApps** | **int32** | maximum number of related visibleApps returned (when they are included) | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUpdateInstance

> UserResponse UsersUpdateInstance(ctx, id).UserUpdateRequest(userUpdateRequest).Execute()



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
    userUpdateRequest := *openapiclient.NewUserUpdateRequest(*openapiclient.NewUserUpdateRequestData("Type_example", "Id_example")) // UserUpdateRequest | User representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.UsersUpdateInstance(context.Background(), id).UserUpdateRequest(userUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersUpdateInstance`: UserResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userUpdateRequest** | [**UserUpdateRequest**](UserUpdateRequest.md) | User representation | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersVisibleAppsCreateToManyRelationship

> UsersVisibleAppsCreateToManyRelationship(ctx, id).UserVisibleAppsLinkagesRequest(userVisibleAppsLinkagesRequest).Execute()



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
    userVisibleAppsLinkagesRequest := *openapiclient.NewUserVisibleAppsLinkagesRequest([]openapiclient.AppAvailabilityV2CreateRequestDataRelationshipsAppData{*openapiclient.NewAppAvailabilityV2CreateRequestDataRelationshipsAppData("Type_example", "Id_example")}) // UserVisibleAppsLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersAPI.UsersVisibleAppsCreateToManyRelationship(context.Background(), id).UserVisibleAppsLinkagesRequest(userVisibleAppsLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersVisibleAppsCreateToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUsersVisibleAppsCreateToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userVisibleAppsLinkagesRequest** | [**UserVisibleAppsLinkagesRequest**](UserVisibleAppsLinkagesRequest.md) | List of related linkages | 

### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersVisibleAppsDeleteToManyRelationship

> UsersVisibleAppsDeleteToManyRelationship(ctx, id).UserVisibleAppsLinkagesRequest(userVisibleAppsLinkagesRequest).Execute()



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
    userVisibleAppsLinkagesRequest := *openapiclient.NewUserVisibleAppsLinkagesRequest([]openapiclient.AppAvailabilityV2CreateRequestDataRelationshipsAppData{*openapiclient.NewAppAvailabilityV2CreateRequestDataRelationshipsAppData("Type_example", "Id_example")}) // UserVisibleAppsLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersAPI.UsersVisibleAppsDeleteToManyRelationship(context.Background(), id).UserVisibleAppsLinkagesRequest(userVisibleAppsLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersVisibleAppsDeleteToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUsersVisibleAppsDeleteToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userVisibleAppsLinkagesRequest** | [**UserVisibleAppsLinkagesRequest**](UserVisibleAppsLinkagesRequest.md) | List of related linkages | 

### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersVisibleAppsGetToManyRelated

> AppsWithoutIncludesResponse UsersVisibleAppsGetToManyRelated(ctx, id).FieldsApps(fieldsApps).Limit(limit).Execute()



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
    fieldsApps := []string{"FieldsApps_example"} // []string | the fields to include for returned resources of type apps (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.UsersVisibleAppsGetToManyRelated(context.Background(), id).FieldsApps(fieldsApps).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersVisibleAppsGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersVisibleAppsGetToManyRelated`: AppsWithoutIncludesResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersVisibleAppsGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersVisibleAppsGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsApps** | **[]string** | the fields to include for returned resources of type apps | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**AppsWithoutIncludesResponse**](AppsWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersVisibleAppsGetToManyRelationship

> UserVisibleAppsLinkagesResponse UsersVisibleAppsGetToManyRelationship(ctx, id).Limit(limit).Execute()



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
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.UsersVisibleAppsGetToManyRelationship(context.Background(), id).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersVisibleAppsGetToManyRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersVisibleAppsGetToManyRelationship`: UserVisibleAppsLinkagesResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersVisibleAppsGetToManyRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersVisibleAppsGetToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | maximum resources per page | 

### Return type

[**UserVisibleAppsLinkagesResponse**](UserVisibleAppsLinkagesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersVisibleAppsReplaceToManyRelationship

> UsersVisibleAppsReplaceToManyRelationship(ctx, id).UserVisibleAppsLinkagesRequest(userVisibleAppsLinkagesRequest).Execute()



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
    userVisibleAppsLinkagesRequest := *openapiclient.NewUserVisibleAppsLinkagesRequest([]openapiclient.AppAvailabilityV2CreateRequestDataRelationshipsAppData{*openapiclient.NewAppAvailabilityV2CreateRequestDataRelationshipsAppData("Type_example", "Id_example")}) // UserVisibleAppsLinkagesRequest | List of related linkages

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersAPI.UsersVisibleAppsReplaceToManyRelationship(context.Background(), id).UserVisibleAppsLinkagesRequest(userVisibleAppsLinkagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersVisibleAppsReplaceToManyRelationship``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUsersVisibleAppsReplaceToManyRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userVisibleAppsLinkagesRequest** | [**UserVisibleAppsLinkagesRequest**](UserVisibleAppsLinkagesRequest.md) | List of related linkages | 

### Return type

 (empty response body)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

