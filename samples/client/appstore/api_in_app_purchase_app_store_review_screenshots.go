/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// InAppPurchaseAppStoreReviewScreenshotsAPIService InAppPurchaseAppStoreReviewScreenshotsAPI service
type InAppPurchaseAppStoreReviewScreenshotsAPIService service

type ApiInAppPurchaseAppStoreReviewScreenshotsCreateInstanceRequest struct {
	ctx context.Context
	ApiService *InAppPurchaseAppStoreReviewScreenshotsAPIService
	inAppPurchaseAppStoreReviewScreenshotCreateRequest *InAppPurchaseAppStoreReviewScreenshotCreateRequest
}

// InAppPurchaseAppStoreReviewScreenshot representation
func (r ApiInAppPurchaseAppStoreReviewScreenshotsCreateInstanceRequest) InAppPurchaseAppStoreReviewScreenshotCreateRequest(inAppPurchaseAppStoreReviewScreenshotCreateRequest InAppPurchaseAppStoreReviewScreenshotCreateRequest) ApiInAppPurchaseAppStoreReviewScreenshotsCreateInstanceRequest {
	r.inAppPurchaseAppStoreReviewScreenshotCreateRequest = &inAppPurchaseAppStoreReviewScreenshotCreateRequest
	return r
}

func (r ApiInAppPurchaseAppStoreReviewScreenshotsCreateInstanceRequest) Execute() (*InAppPurchaseAppStoreReviewScreenshotResponse, *http.Response, error) {
	return r.ApiService.InAppPurchaseAppStoreReviewScreenshotsCreateInstanceExecute(r)
}

/*
InAppPurchaseAppStoreReviewScreenshotsCreateInstance Method for InAppPurchaseAppStoreReviewScreenshotsCreateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiInAppPurchaseAppStoreReviewScreenshotsCreateInstanceRequest
*/
func (a *InAppPurchaseAppStoreReviewScreenshotsAPIService) InAppPurchaseAppStoreReviewScreenshotsCreateInstance(ctx context.Context) ApiInAppPurchaseAppStoreReviewScreenshotsCreateInstanceRequest {
	return ApiInAppPurchaseAppStoreReviewScreenshotsCreateInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InAppPurchaseAppStoreReviewScreenshotResponse
func (a *InAppPurchaseAppStoreReviewScreenshotsAPIService) InAppPurchaseAppStoreReviewScreenshotsCreateInstanceExecute(r ApiInAppPurchaseAppStoreReviewScreenshotsCreateInstanceRequest) (*InAppPurchaseAppStoreReviewScreenshotResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InAppPurchaseAppStoreReviewScreenshotResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InAppPurchaseAppStoreReviewScreenshotsAPIService.InAppPurchaseAppStoreReviewScreenshotsCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/inAppPurchaseAppStoreReviewScreenshots"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.inAppPurchaseAppStoreReviewScreenshotCreateRequest == nil {
		return localVarReturnValue, nil, reportError("inAppPurchaseAppStoreReviewScreenshotCreateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.inAppPurchaseAppStoreReviewScreenshotCreateRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiInAppPurchaseAppStoreReviewScreenshotsDeleteInstanceRequest struct {
	ctx context.Context
	ApiService *InAppPurchaseAppStoreReviewScreenshotsAPIService
	id string
}

func (r ApiInAppPurchaseAppStoreReviewScreenshotsDeleteInstanceRequest) Execute() (*http.Response, error) {
	return r.ApiService.InAppPurchaseAppStoreReviewScreenshotsDeleteInstanceExecute(r)
}

/*
InAppPurchaseAppStoreReviewScreenshotsDeleteInstance Method for InAppPurchaseAppStoreReviewScreenshotsDeleteInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiInAppPurchaseAppStoreReviewScreenshotsDeleteInstanceRequest
*/
func (a *InAppPurchaseAppStoreReviewScreenshotsAPIService) InAppPurchaseAppStoreReviewScreenshotsDeleteInstance(ctx context.Context, id string) ApiInAppPurchaseAppStoreReviewScreenshotsDeleteInstanceRequest {
	return ApiInAppPurchaseAppStoreReviewScreenshotsDeleteInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *InAppPurchaseAppStoreReviewScreenshotsAPIService) InAppPurchaseAppStoreReviewScreenshotsDeleteInstanceExecute(r ApiInAppPurchaseAppStoreReviewScreenshotsDeleteInstanceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InAppPurchaseAppStoreReviewScreenshotsAPIService.InAppPurchaseAppStoreReviewScreenshotsDeleteInstance")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/inAppPurchaseAppStoreReviewScreenshots/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiInAppPurchaseAppStoreReviewScreenshotsGetInstanceRequest struct {
	ctx context.Context
	ApiService *InAppPurchaseAppStoreReviewScreenshotsAPIService
	id string
	fieldsInAppPurchaseAppStoreReviewScreenshots *[]string
	include *[]string
}

// the fields to include for returned resources of type inAppPurchaseAppStoreReviewScreenshots
func (r ApiInAppPurchaseAppStoreReviewScreenshotsGetInstanceRequest) FieldsInAppPurchaseAppStoreReviewScreenshots(fieldsInAppPurchaseAppStoreReviewScreenshots []string) ApiInAppPurchaseAppStoreReviewScreenshotsGetInstanceRequest {
	r.fieldsInAppPurchaseAppStoreReviewScreenshots = &fieldsInAppPurchaseAppStoreReviewScreenshots
	return r
}

// comma-separated list of relationships to include
func (r ApiInAppPurchaseAppStoreReviewScreenshotsGetInstanceRequest) Include(include []string) ApiInAppPurchaseAppStoreReviewScreenshotsGetInstanceRequest {
	r.include = &include
	return r
}

func (r ApiInAppPurchaseAppStoreReviewScreenshotsGetInstanceRequest) Execute() (*InAppPurchaseAppStoreReviewScreenshotResponse, *http.Response, error) {
	return r.ApiService.InAppPurchaseAppStoreReviewScreenshotsGetInstanceExecute(r)
}

/*
InAppPurchaseAppStoreReviewScreenshotsGetInstance Method for InAppPurchaseAppStoreReviewScreenshotsGetInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiInAppPurchaseAppStoreReviewScreenshotsGetInstanceRequest
*/
func (a *InAppPurchaseAppStoreReviewScreenshotsAPIService) InAppPurchaseAppStoreReviewScreenshotsGetInstance(ctx context.Context, id string) ApiInAppPurchaseAppStoreReviewScreenshotsGetInstanceRequest {
	return ApiInAppPurchaseAppStoreReviewScreenshotsGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return InAppPurchaseAppStoreReviewScreenshotResponse
func (a *InAppPurchaseAppStoreReviewScreenshotsAPIService) InAppPurchaseAppStoreReviewScreenshotsGetInstanceExecute(r ApiInAppPurchaseAppStoreReviewScreenshotsGetInstanceRequest) (*InAppPurchaseAppStoreReviewScreenshotResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InAppPurchaseAppStoreReviewScreenshotResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InAppPurchaseAppStoreReviewScreenshotsAPIService.InAppPurchaseAppStoreReviewScreenshotsGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/inAppPurchaseAppStoreReviewScreenshots/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsInAppPurchaseAppStoreReviewScreenshots != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[inAppPurchaseAppStoreReviewScreenshots]", r.fieldsInAppPurchaseAppStoreReviewScreenshots, "csv")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiInAppPurchaseAppStoreReviewScreenshotsUpdateInstanceRequest struct {
	ctx context.Context
	ApiService *InAppPurchaseAppStoreReviewScreenshotsAPIService
	id string
	inAppPurchaseAppStoreReviewScreenshotUpdateRequest *InAppPurchaseAppStoreReviewScreenshotUpdateRequest
}

// InAppPurchaseAppStoreReviewScreenshot representation
func (r ApiInAppPurchaseAppStoreReviewScreenshotsUpdateInstanceRequest) InAppPurchaseAppStoreReviewScreenshotUpdateRequest(inAppPurchaseAppStoreReviewScreenshotUpdateRequest InAppPurchaseAppStoreReviewScreenshotUpdateRequest) ApiInAppPurchaseAppStoreReviewScreenshotsUpdateInstanceRequest {
	r.inAppPurchaseAppStoreReviewScreenshotUpdateRequest = &inAppPurchaseAppStoreReviewScreenshotUpdateRequest
	return r
}

func (r ApiInAppPurchaseAppStoreReviewScreenshotsUpdateInstanceRequest) Execute() (*InAppPurchaseAppStoreReviewScreenshotResponse, *http.Response, error) {
	return r.ApiService.InAppPurchaseAppStoreReviewScreenshotsUpdateInstanceExecute(r)
}

/*
InAppPurchaseAppStoreReviewScreenshotsUpdateInstance Method for InAppPurchaseAppStoreReviewScreenshotsUpdateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiInAppPurchaseAppStoreReviewScreenshotsUpdateInstanceRequest
*/
func (a *InAppPurchaseAppStoreReviewScreenshotsAPIService) InAppPurchaseAppStoreReviewScreenshotsUpdateInstance(ctx context.Context, id string) ApiInAppPurchaseAppStoreReviewScreenshotsUpdateInstanceRequest {
	return ApiInAppPurchaseAppStoreReviewScreenshotsUpdateInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return InAppPurchaseAppStoreReviewScreenshotResponse
func (a *InAppPurchaseAppStoreReviewScreenshotsAPIService) InAppPurchaseAppStoreReviewScreenshotsUpdateInstanceExecute(r ApiInAppPurchaseAppStoreReviewScreenshotsUpdateInstanceRequest) (*InAppPurchaseAppStoreReviewScreenshotResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InAppPurchaseAppStoreReviewScreenshotResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InAppPurchaseAppStoreReviewScreenshotsAPIService.InAppPurchaseAppStoreReviewScreenshotsUpdateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/inAppPurchaseAppStoreReviewScreenshots/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.inAppPurchaseAppStoreReviewScreenshotUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("inAppPurchaseAppStoreReviewScreenshotUpdateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.inAppPurchaseAppStoreReviewScreenshotUpdateRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
