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

// AppPreviewSetsAPIService AppPreviewSetsAPI service
type AppPreviewSetsAPIService service

type ApiAppPreviewSetsAppPreviewsGetToManyRelatedRequest struct {
	ctx                  context.Context
	ApiService           *AppPreviewSetsAPIService
	id                   string
	fieldsAppPreviews    *[]string
	fieldsAppPreviewSets *[]string
	limit                *int32
	include              *[]string
}

// the fields to include for returned resources of type appPreviews
func (r ApiAppPreviewSetsAppPreviewsGetToManyRelatedRequest) FieldsAppPreviews(fieldsAppPreviews []string) ApiAppPreviewSetsAppPreviewsGetToManyRelatedRequest {
	r.fieldsAppPreviews = &fieldsAppPreviews
	return r
}

// the fields to include for returned resources of type appPreviewSets
func (r ApiAppPreviewSetsAppPreviewsGetToManyRelatedRequest) FieldsAppPreviewSets(fieldsAppPreviewSets []string) ApiAppPreviewSetsAppPreviewsGetToManyRelatedRequest {
	r.fieldsAppPreviewSets = &fieldsAppPreviewSets
	return r
}

// maximum resources per page
func (r ApiAppPreviewSetsAppPreviewsGetToManyRelatedRequest) Limit(limit int32) ApiAppPreviewSetsAppPreviewsGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

// comma-separated list of relationships to include
func (r ApiAppPreviewSetsAppPreviewsGetToManyRelatedRequest) Include(include []string) ApiAppPreviewSetsAppPreviewsGetToManyRelatedRequest {
	r.include = &include
	return r
}

func (r ApiAppPreviewSetsAppPreviewsGetToManyRelatedRequest) Execute() (*AppPreviewsResponse, *http.Response, error) {
	return r.ApiService.AppPreviewSetsAppPreviewsGetToManyRelatedExecute(r)
}

/*
AppPreviewSetsAppPreviewsGetToManyRelated Method for AppPreviewSetsAppPreviewsGetToManyRelated

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiAppPreviewSetsAppPreviewsGetToManyRelatedRequest
*/
func (a *AppPreviewSetsAPIService) AppPreviewSetsAppPreviewsGetToManyRelated(ctx context.Context, id string) ApiAppPreviewSetsAppPreviewsGetToManyRelatedRequest {
	return ApiAppPreviewSetsAppPreviewsGetToManyRelatedRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AppPreviewsResponse
func (a *AppPreviewSetsAPIService) AppPreviewSetsAppPreviewsGetToManyRelatedExecute(r ApiAppPreviewSetsAppPreviewsGetToManyRelatedRequest) (*AppPreviewsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AppPreviewsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppPreviewSetsAPIService.AppPreviewSetsAppPreviewsGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appPreviewSets/{id}/appPreviews"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsAppPreviews != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appPreviews]", r.fieldsAppPreviews, "csv")
	}
	if r.fieldsAppPreviewSets != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appPreviewSets]", r.fieldsAppPreviewSets, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type ApiAppPreviewSetsAppPreviewsGetToManyRelationshipRequest struct {
	ctx        context.Context
	ApiService *AppPreviewSetsAPIService
	id         string
	limit      *int32
}

// maximum resources per page
func (r ApiAppPreviewSetsAppPreviewsGetToManyRelationshipRequest) Limit(limit int32) ApiAppPreviewSetsAppPreviewsGetToManyRelationshipRequest {
	r.limit = &limit
	return r
}

func (r ApiAppPreviewSetsAppPreviewsGetToManyRelationshipRequest) Execute() (*AppPreviewSetAppPreviewsLinkagesResponse, *http.Response, error) {
	return r.ApiService.AppPreviewSetsAppPreviewsGetToManyRelationshipExecute(r)
}

/*
AppPreviewSetsAppPreviewsGetToManyRelationship Method for AppPreviewSetsAppPreviewsGetToManyRelationship

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiAppPreviewSetsAppPreviewsGetToManyRelationshipRequest
*/
func (a *AppPreviewSetsAPIService) AppPreviewSetsAppPreviewsGetToManyRelationship(ctx context.Context, id string) ApiAppPreviewSetsAppPreviewsGetToManyRelationshipRequest {
	return ApiAppPreviewSetsAppPreviewsGetToManyRelationshipRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AppPreviewSetAppPreviewsLinkagesResponse
func (a *AppPreviewSetsAPIService) AppPreviewSetsAppPreviewsGetToManyRelationshipExecute(r ApiAppPreviewSetsAppPreviewsGetToManyRelationshipRequest) (*AppPreviewSetAppPreviewsLinkagesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AppPreviewSetAppPreviewsLinkagesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppPreviewSetsAPIService.AppPreviewSetsAppPreviewsGetToManyRelationship")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appPreviewSets/{id}/relationships/appPreviews"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

type ApiAppPreviewSetsAppPreviewsReplaceToManyRelationshipRequest struct {
	ctx                                     context.Context
	ApiService                              *AppPreviewSetsAPIService
	id                                      string
	appPreviewSetAppPreviewsLinkagesRequest *AppPreviewSetAppPreviewsLinkagesRequest
}

// List of related linkages
func (r ApiAppPreviewSetsAppPreviewsReplaceToManyRelationshipRequest) AppPreviewSetAppPreviewsLinkagesRequest(appPreviewSetAppPreviewsLinkagesRequest AppPreviewSetAppPreviewsLinkagesRequest) ApiAppPreviewSetsAppPreviewsReplaceToManyRelationshipRequest {
	r.appPreviewSetAppPreviewsLinkagesRequest = &appPreviewSetAppPreviewsLinkagesRequest
	return r
}

func (r ApiAppPreviewSetsAppPreviewsReplaceToManyRelationshipRequest) Execute() (*http.Response, error) {
	return r.ApiService.AppPreviewSetsAppPreviewsReplaceToManyRelationshipExecute(r)
}

/*
AppPreviewSetsAppPreviewsReplaceToManyRelationship Method for AppPreviewSetsAppPreviewsReplaceToManyRelationship

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiAppPreviewSetsAppPreviewsReplaceToManyRelationshipRequest
*/
func (a *AppPreviewSetsAPIService) AppPreviewSetsAppPreviewsReplaceToManyRelationship(ctx context.Context, id string) ApiAppPreviewSetsAppPreviewsReplaceToManyRelationshipRequest {
	return ApiAppPreviewSetsAppPreviewsReplaceToManyRelationshipRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *AppPreviewSetsAPIService) AppPreviewSetsAppPreviewsReplaceToManyRelationshipExecute(r ApiAppPreviewSetsAppPreviewsReplaceToManyRelationshipRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppPreviewSetsAPIService.AppPreviewSetsAppPreviewsReplaceToManyRelationship")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appPreviewSets/{id}/relationships/appPreviews"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.appPreviewSetAppPreviewsLinkagesRequest == nil {
		return nil, reportError("appPreviewSetAppPreviewsLinkagesRequest is required and must be specified")
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
	localVarPostBody = r.appPreviewSetAppPreviewsLinkagesRequest
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

type ApiAppPreviewSetsCreateInstanceRequest struct {
	ctx                        context.Context
	ApiService                 *AppPreviewSetsAPIService
	appPreviewSetCreateRequest *AppPreviewSetCreateRequest
}

// AppPreviewSet representation
func (r ApiAppPreviewSetsCreateInstanceRequest) AppPreviewSetCreateRequest(appPreviewSetCreateRequest AppPreviewSetCreateRequest) ApiAppPreviewSetsCreateInstanceRequest {
	r.appPreviewSetCreateRequest = &appPreviewSetCreateRequest
	return r
}

func (r ApiAppPreviewSetsCreateInstanceRequest) Execute() (*AppPreviewSetResponse, *http.Response, error) {
	return r.ApiService.AppPreviewSetsCreateInstanceExecute(r)
}

/*
AppPreviewSetsCreateInstance Method for AppPreviewSetsCreateInstance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAppPreviewSetsCreateInstanceRequest
*/
func (a *AppPreviewSetsAPIService) AppPreviewSetsCreateInstance(ctx context.Context) ApiAppPreviewSetsCreateInstanceRequest {
	return ApiAppPreviewSetsCreateInstanceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AppPreviewSetResponse
func (a *AppPreviewSetsAPIService) AppPreviewSetsCreateInstanceExecute(r ApiAppPreviewSetsCreateInstanceRequest) (*AppPreviewSetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AppPreviewSetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppPreviewSetsAPIService.AppPreviewSetsCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appPreviewSets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.appPreviewSetCreateRequest == nil {
		return localVarReturnValue, nil, reportError("appPreviewSetCreateRequest is required and must be specified")
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
	localVarPostBody = r.appPreviewSetCreateRequest
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

type ApiAppPreviewSetsDeleteInstanceRequest struct {
	ctx        context.Context
	ApiService *AppPreviewSetsAPIService
	id         string
}

func (r ApiAppPreviewSetsDeleteInstanceRequest) Execute() (*http.Response, error) {
	return r.ApiService.AppPreviewSetsDeleteInstanceExecute(r)
}

/*
AppPreviewSetsDeleteInstance Method for AppPreviewSetsDeleteInstance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiAppPreviewSetsDeleteInstanceRequest
*/
func (a *AppPreviewSetsAPIService) AppPreviewSetsDeleteInstance(ctx context.Context, id string) ApiAppPreviewSetsDeleteInstanceRequest {
	return ApiAppPreviewSetsDeleteInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *AppPreviewSetsAPIService) AppPreviewSetsDeleteInstanceExecute(r ApiAppPreviewSetsDeleteInstanceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppPreviewSetsAPIService.AppPreviewSetsDeleteInstance")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appPreviewSets/{id}"
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

type ApiAppPreviewSetsGetInstanceRequest struct {
	ctx                  context.Context
	ApiService           *AppPreviewSetsAPIService
	id                   string
	fieldsAppPreviewSets *[]string
	include              *[]string
	fieldsAppPreviews    *[]string
	limitAppPreviews     *int32
}

// the fields to include for returned resources of type appPreviewSets
func (r ApiAppPreviewSetsGetInstanceRequest) FieldsAppPreviewSets(fieldsAppPreviewSets []string) ApiAppPreviewSetsGetInstanceRequest {
	r.fieldsAppPreviewSets = &fieldsAppPreviewSets
	return r
}

// comma-separated list of relationships to include
func (r ApiAppPreviewSetsGetInstanceRequest) Include(include []string) ApiAppPreviewSetsGetInstanceRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type appPreviews
func (r ApiAppPreviewSetsGetInstanceRequest) FieldsAppPreviews(fieldsAppPreviews []string) ApiAppPreviewSetsGetInstanceRequest {
	r.fieldsAppPreviews = &fieldsAppPreviews
	return r
}

// maximum number of related appPreviews returned (when they are included)
func (r ApiAppPreviewSetsGetInstanceRequest) LimitAppPreviews(limitAppPreviews int32) ApiAppPreviewSetsGetInstanceRequest {
	r.limitAppPreviews = &limitAppPreviews
	return r
}

func (r ApiAppPreviewSetsGetInstanceRequest) Execute() (*AppPreviewSetResponse, *http.Response, error) {
	return r.ApiService.AppPreviewSetsGetInstanceExecute(r)
}

/*
AppPreviewSetsGetInstance Method for AppPreviewSetsGetInstance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiAppPreviewSetsGetInstanceRequest
*/
func (a *AppPreviewSetsAPIService) AppPreviewSetsGetInstance(ctx context.Context, id string) ApiAppPreviewSetsGetInstanceRequest {
	return ApiAppPreviewSetsGetInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AppPreviewSetResponse
func (a *AppPreviewSetsAPIService) AppPreviewSetsGetInstanceExecute(r ApiAppPreviewSetsGetInstanceRequest) (*AppPreviewSetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AppPreviewSetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppPreviewSetsAPIService.AppPreviewSetsGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appPreviewSets/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsAppPreviewSets != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appPreviewSets]", r.fieldsAppPreviewSets, "csv")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.fieldsAppPreviews != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appPreviews]", r.fieldsAppPreviews, "csv")
	}
	if r.limitAppPreviews != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[appPreviews]", r.limitAppPreviews, "")
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
