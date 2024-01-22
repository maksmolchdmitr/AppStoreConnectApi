# \AppEncryptionDeclarationDocumentsAPI

All URIs are relative to *https://api.appstoreconnect.apple.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppEncryptionDeclarationDocumentsCreateInstance**](AppEncryptionDeclarationDocumentsAPI.md#AppEncryptionDeclarationDocumentsCreateInstance) | **Post** /v1/appEncryptionDeclarationDocuments | 
[**AppEncryptionDeclarationDocumentsGetInstance**](AppEncryptionDeclarationDocumentsAPI.md#AppEncryptionDeclarationDocumentsGetInstance) | **Get** /v1/appEncryptionDeclarationDocuments/{id} | 
[**AppEncryptionDeclarationDocumentsUpdateInstance**](AppEncryptionDeclarationDocumentsAPI.md#AppEncryptionDeclarationDocumentsUpdateInstance) | **Patch** /v1/appEncryptionDeclarationDocuments/{id} | 



## AppEncryptionDeclarationDocumentsCreateInstance

> AppEncryptionDeclarationDocumentResponse AppEncryptionDeclarationDocumentsCreateInstance(ctx).AppEncryptionDeclarationDocumentCreateRequest(appEncryptionDeclarationDocumentCreateRequest).Execute()



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
    appEncryptionDeclarationDocumentCreateRequest := *openapiclient.NewAppEncryptionDeclarationDocumentCreateRequest(*openapiclient.NewAppEncryptionDeclarationDocumentCreateRequestData("Type_example", *openapiclient.NewAppClipAdvancedExperienceImageCreateRequestDataAttributes(int32(123), "FileName_example"), *openapiclient.NewAppEncryptionDeclarationDocumentCreateRequestDataRelationships(*openapiclient.NewAppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclaration(*openapiclient.NewAppEncryptionDeclarationDocumentCreateRequestDataRelationshipsAppEncryptionDeclarationData("Type_example", "Id_example"))))) // AppEncryptionDeclarationDocumentCreateRequest | AppEncryptionDeclarationDocument representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppEncryptionDeclarationDocumentsAPI.AppEncryptionDeclarationDocumentsCreateInstance(context.Background()).AppEncryptionDeclarationDocumentCreateRequest(appEncryptionDeclarationDocumentCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppEncryptionDeclarationDocumentsAPI.AppEncryptionDeclarationDocumentsCreateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppEncryptionDeclarationDocumentsCreateInstance`: AppEncryptionDeclarationDocumentResponse
    fmt.Fprintf(os.Stdout, "Response from `AppEncryptionDeclarationDocumentsAPI.AppEncryptionDeclarationDocumentsCreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppEncryptionDeclarationDocumentsCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appEncryptionDeclarationDocumentCreateRequest** | [**AppEncryptionDeclarationDocumentCreateRequest**](AppEncryptionDeclarationDocumentCreateRequest.md) | AppEncryptionDeclarationDocument representation | 

### Return type

[**AppEncryptionDeclarationDocumentResponse**](AppEncryptionDeclarationDocumentResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppEncryptionDeclarationDocumentsGetInstance

> AppEncryptionDeclarationDocumentResponse AppEncryptionDeclarationDocumentsGetInstance(ctx, id).FieldsAppEncryptionDeclarationDocuments(fieldsAppEncryptionDeclarationDocuments).Execute()



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
    fieldsAppEncryptionDeclarationDocuments := []string{"FieldsAppEncryptionDeclarationDocuments_example"} // []string | the fields to include for returned resources of type appEncryptionDeclarationDocuments (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppEncryptionDeclarationDocumentsAPI.AppEncryptionDeclarationDocumentsGetInstance(context.Background(), id).FieldsAppEncryptionDeclarationDocuments(fieldsAppEncryptionDeclarationDocuments).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppEncryptionDeclarationDocumentsAPI.AppEncryptionDeclarationDocumentsGetInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppEncryptionDeclarationDocumentsGetInstance`: AppEncryptionDeclarationDocumentResponse
    fmt.Fprintf(os.Stdout, "Response from `AppEncryptionDeclarationDocumentsAPI.AppEncryptionDeclarationDocumentsGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppEncryptionDeclarationDocumentsGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldsAppEncryptionDeclarationDocuments** | **[]string** | the fields to include for returned resources of type appEncryptionDeclarationDocuments | 

### Return type

[**AppEncryptionDeclarationDocumentResponse**](AppEncryptionDeclarationDocumentResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppEncryptionDeclarationDocumentsUpdateInstance

> AppEncryptionDeclarationDocumentResponse AppEncryptionDeclarationDocumentsUpdateInstance(ctx, id).AppEncryptionDeclarationDocumentUpdateRequest(appEncryptionDeclarationDocumentUpdateRequest).Execute()



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
    appEncryptionDeclarationDocumentUpdateRequest := *openapiclient.NewAppEncryptionDeclarationDocumentUpdateRequest(*openapiclient.NewAppEncryptionDeclarationDocumentUpdateRequestData("Type_example", "Id_example")) // AppEncryptionDeclarationDocumentUpdateRequest | AppEncryptionDeclarationDocument representation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppEncryptionDeclarationDocumentsAPI.AppEncryptionDeclarationDocumentsUpdateInstance(context.Background(), id).AppEncryptionDeclarationDocumentUpdateRequest(appEncryptionDeclarationDocumentUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppEncryptionDeclarationDocumentsAPI.AppEncryptionDeclarationDocumentsUpdateInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppEncryptionDeclarationDocumentsUpdateInstance`: AppEncryptionDeclarationDocumentResponse
    fmt.Fprintf(os.Stdout, "Response from `AppEncryptionDeclarationDocumentsAPI.AppEncryptionDeclarationDocumentsUpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the id of the requested resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppEncryptionDeclarationDocumentsUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **appEncryptionDeclarationDocumentUpdateRequest** | [**AppEncryptionDeclarationDocumentUpdateRequest**](AppEncryptionDeclarationDocumentUpdateRequest.md) | AppEncryptionDeclarationDocument representation | 

### Return type

[**AppEncryptionDeclarationDocumentResponse**](AppEncryptionDeclarationDocumentResponse.md)

### Authorization

[itc-bearer-token](../README.md#itc-bearer-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

