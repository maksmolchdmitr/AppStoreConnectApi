# \ProfilesAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProfilesBundleIdGetToOneRelated**](ProfilesAPI.md#ProfilesBundleIdGetToOneRelated) | **Get** /v1/profiles/{id}/bundleId | 
[**ProfilesCertificatesGetToManyRelated**](ProfilesAPI.md#ProfilesCertificatesGetToManyRelated) | **Get** /v1/profiles/{id}/certificates | 
[**ProfilesCreateInstance**](ProfilesAPI.md#ProfilesCreateInstance) | **Post** /v1/profiles | 
[**ProfilesDeleteInstance**](ProfilesAPI.md#ProfilesDeleteInstance) | **Delete** /v1/profiles/{id} | 
[**ProfilesDevicesGetToManyRelated**](ProfilesAPI.md#ProfilesDevicesGetToManyRelated) | **Get** /v1/profiles/{id}/devices | 
[**ProfilesGetCollection**](ProfilesAPI.md#ProfilesGetCollection) | **Get** /v1/profiles | 
[**ProfilesGetInstance**](ProfilesAPI.md#ProfilesGetInstance) | **Get** /v1/profiles/{id} | 



## ProfilesBundleIdGetToOneRelated

> BundleIdWithoutIncludesResponse ProfilesBundleIdGetToOneRelated(ctx, id).FieldsBundleIds(fieldsBundleIds).Execute()



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
    fieldsBundleIds := []string{"FieldsBundleIds_example"} // []string | the fields to include for returned resources of type bundleIds (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfilesAPI.ProfilesBundleIdGetToOneRelated(context.Background(), id).FieldsBundleIds(fieldsBundleIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ProfilesBundleIdGetToOneRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProfilesBundleIdGetToOneRelated`: BundleIdWithoutIncludesResponse
    fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ProfilesBundleIdGetToOneRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiProfilesBundleIdGetToOneRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsBundleIds** | **[]string** | the fields to include for returned resources of type bundleIds | 

### Return type

[**BundleIdWithoutIncludesResponse**](BundleIdWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProfilesCertificatesGetToManyRelated

> CertificatesWithoutIncludesResponse ProfilesCertificatesGetToManyRelated(ctx, id).FieldsCertificates(fieldsCertificates).Limit(limit).Execute()



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
    fieldsCertificates := []string{"FieldsCertificates_example"} // []string | the fields to include for returned resources of type certificates (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfilesAPI.ProfilesCertificatesGetToManyRelated(context.Background(), id).FieldsCertificates(fieldsCertificates).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ProfilesCertificatesGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProfilesCertificatesGetToManyRelated`: CertificatesWithoutIncludesResponse
    fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ProfilesCertificatesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiProfilesCertificatesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsCertificates** | **[]string** | the fields to include for returned resources of type certificates | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**CertificatesWithoutIncludesResponse**](CertificatesWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProfilesCreateInstance

> ProfileResponse ProfilesCreateInstance(ctx).ProfileCreateRequest(profileCreateRequest).Execute()



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
    profileCreateRequest := *openapiclient.NewProfileCreateRequest(*openapiclient.NewProfileCreateRequestData("Type_example", *openapiclient.NewProfileCreateRequestDataAttributes("Name_example", "ProfileType_example"), *openapiclient.NewProfileCreateRequestDataRelationships(*openapiclient.NewBundleIdCapabilityCreateRequestDataRelationshipsBundleId(*openapiclient.NewBundleIdCapabilityCreateRequestDataRelationshipsBundleIdData("Type_example", "Id_example")), *openapiclient.NewProfileCreateRequestDataRelationshipsCertificates([]openapiclient.ProfileRelationshipsCertificatesDataInner{*openapiclient.NewProfileRelationshipsCertificatesDataInner("Type_example", "Id_example")})))) // ProfileCreateRequest | Profile representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfilesAPI.ProfilesCreateInstance(context.Background()).ProfileCreateRequest(profileCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ProfilesCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProfilesCreateInstance`: ProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ProfilesCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProfilesCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profileCreateRequest** | [**ProfileCreateRequest**](ProfileCreateRequest.md) | Profile representation | 

### Return type

[**ProfileResponse**](ProfileResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProfilesDeleteInstance

> ProfilesDeleteInstance(ctx, id).Execute()



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
    r, err := apiClient.ProfilesAPI.ProfilesDeleteInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ProfilesDeleteInstance``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProfilesDeleteInstanceRequest struct via the builder pattern


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


## ProfilesDevicesGetToManyRelated

> DevicesWithoutIncludesResponse ProfilesDevicesGetToManyRelated(ctx, id).FieldsDevices(fieldsDevices).Limit(limit).Execute()



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
    fieldsDevices := []string{"FieldsDevices_example"} // []string | the fields to include for returned resources of type devices (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfilesAPI.ProfilesDevicesGetToManyRelated(context.Background(), id).FieldsDevices(fieldsDevices).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ProfilesDevicesGetToManyRelated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProfilesDevicesGetToManyRelated`: DevicesWithoutIncludesResponse
    fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ProfilesDevicesGetToManyRelated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiProfilesDevicesGetToManyRelatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsDevices** | **[]string** | the fields to include for returned resources of type devices | 
 **limit** | **int32** | maximum resources per page | 

### Return type

[**DevicesWithoutIncludesResponse**](DevicesWithoutIncludesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProfilesGetCollection

> ProfilesResponse ProfilesGetCollection(ctx).FilterName(filterName).FilterProfileState(filterProfileState).FilterProfileType(filterProfileType).FilterId(filterId).Sort(sort).FieldsProfiles(fieldsProfiles).Limit(limit).Include(include).FieldsCertificates(fieldsCertificates).FieldsDevices(fieldsDevices).FieldsBundleIds(fieldsBundleIds).LimitCertificates(limitCertificates).LimitDevices(limitDevices).Execute()



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
    filterName := []string{"Inner_example"} // []string | filter by attribute 'name' (optional)
    filterProfileState := []string{"FilterProfileState_example"} // []string | filter by attribute 'profileState' (optional)
    filterProfileType := []string{"FilterProfileType_example"} // []string | filter by attribute 'profileType' (optional)
    filterId := []string{"Inner_example"} // []string | filter by id(s) (optional)
    sort := []string{"Sort_example"} // []string | comma-separated list of sort expressions; resources will be sorted as specified (optional)
    fieldsProfiles := []string{"FieldsProfiles_example"} // []string | the fields to include for returned resources of type profiles (optional)
    limit := int32(56) // int32 | maximum resources per page (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsCertificates := []string{"FieldsCertificates_example"} // []string | the fields to include for returned resources of type certificates (optional)
    fieldsDevices := []string{"FieldsDevices_example"} // []string | the fields to include for returned resources of type devices (optional)
    fieldsBundleIds := []string{"FieldsBundleIds_example"} // []string | the fields to include for returned resources of type bundleIds (optional)
    limitCertificates := int32(56) // int32 | maximum number of related certificates returned (when they are included) (optional)
    limitDevices := int32(56) // int32 | maximum number of related devices returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfilesAPI.ProfilesGetCollection(context.Background()).FilterName(filterName).FilterProfileState(filterProfileState).FilterProfileType(filterProfileType).FilterId(filterId).Sort(sort).FieldsProfiles(fieldsProfiles).Limit(limit).Include(include).FieldsCertificates(fieldsCertificates).FieldsDevices(fieldsDevices).FieldsBundleIds(fieldsBundleIds).LimitCertificates(limitCertificates).LimitDevices(limitDevices).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ProfilesGetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProfilesGetCollection`: ProfilesResponse
    fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ProfilesGetCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProfilesGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterName** | **[]string** | filter by attribute &#39;name&#39; | 
 **filterProfileState** | **[]string** | filter by attribute &#39;profileState&#39; | 
 **filterProfileType** | **[]string** | filter by attribute &#39;profileType&#39; | 
 **filterId** | **[]string** | filter by id(s) | 
 **sort** | **[]string** | comma-separated list of sort expressions; resources will be sorted as specified | 
 **fieldsProfiles** | **[]string** | the fields to include for returned resources of type profiles | 
 **limit** | **int32** | maximum resources per page | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsCertificates** | **[]string** | the fields to include for returned resources of type certificates | 
 **fieldsDevices** | **[]string** | the fields to include for returned resources of type devices | 
 **fieldsBundleIds** | **[]string** | the fields to include for returned resources of type bundleIds | 
 **limitCertificates** | **int32** | maximum number of related certificates returned (when they are included) | 
 **limitDevices** | **int32** | maximum number of related devices returned (when they are included) | 

### Return type

[**ProfilesResponse**](ProfilesResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProfilesGetInstance

> ProfileResponse ProfilesGetInstance(ctx, id).FieldsProfiles(fieldsProfiles).Include(include).FieldsCertificates(fieldsCertificates).FieldsDevices(fieldsDevices).FieldsBundleIds(fieldsBundleIds).LimitCertificates(limitCertificates).LimitDevices(limitDevices).Execute()



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
    fieldsProfiles := []string{"FieldsProfiles_example"} // []string | the fields to include for returned resources of type profiles (optional)
    include := []string{"Include_example"} // []string | comma-separated list of relationships to include (optional)
    fieldsCertificates := []string{"FieldsCertificates_example"} // []string | the fields to include for returned resources of type certificates (optional)
    fieldsDevices := []string{"FieldsDevices_example"} // []string | the fields to include for returned resources of type devices (optional)
    fieldsBundleIds := []string{"FieldsBundleIds_example"} // []string | the fields to include for returned resources of type bundleIds (optional)
    limitCertificates := int32(56) // int32 | maximum number of related certificates returned (when they are included) (optional)
    limitDevices := int32(56) // int32 | maximum number of related devices returned (when they are included) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfilesAPI.ProfilesGetInstance(context.Background(), id).FieldsProfiles(fieldsProfiles).Include(include).FieldsCertificates(fieldsCertificates).FieldsDevices(fieldsDevices).FieldsBundleIds(fieldsBundleIds).LimitCertificates(limitCertificates).LimitDevices(limitDevices).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesAPI.ProfilesGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProfilesGetInstance`: ProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `ProfilesAPI.ProfilesGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiProfilesGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsProfiles** | **[]string** | the fields to include for returned resources of type profiles | 
 **include** | **[]string** | comma-separated list of relationships to include | 
 **fieldsCertificates** | **[]string** | the fields to include for returned resources of type certificates | 
 **fieldsDevices** | **[]string** | the fields to include for returned resources of type devices | 
 **fieldsBundleIds** | **[]string** | the fields to include for returned resources of type bundleIds | 
 **limitCertificates** | **int32** | maximum number of related certificates returned (when they are included) | 
 **limitDevices** | **int32** | maximum number of related devices returned (when they are included) | 

### Return type

[**ProfileResponse**](ProfileResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

