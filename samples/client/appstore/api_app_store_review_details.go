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


// AppStoreReviewDetailsAPIService AppStoreReviewDetailsAPI service
type AppStoreReviewDetailsAPIService service

type ApiAppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelatedRequest struct {
	ctx context.Context
	ApiService *AppStoreReviewDetailsAPIService
	id string
	fieldsAppStoreReviewDetails *[]string
	fieldsAppStoreReviewAttachments *[]string
	limit *int32
	include *[]string
}

// the fields to include for returned resources of type appStoreReviewDetails
func (r ApiAppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelatedRequest) FieldsAppStoreReviewDetails(fieldsAppStoreReviewDetails []string) ApiAppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelatedRequest {
	r.fieldsAppStoreReviewDetails = &fieldsAppStoreReviewDetails
	return r
}

// the fields to include for returned resources of type appStoreReviewAttachments
func (r ApiAppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelatedRequest) FieldsAppStoreReviewAttachments(fieldsAppStoreReviewAttachments []string) ApiAppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelatedRequest {
	r.fieldsAppStoreReviewAttachments = &fieldsAppStoreReviewAttachments
	return r
}

// maximum resources per page
func (r ApiAppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelatedRequest) Limit(limit int32) ApiAppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

// comma-separated list of relationships to include
func (r ApiAppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelatedRequest) Include(include []string) ApiAppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelatedRequest {
	r.include = &include
	return r
}

func (r ApiAppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelatedRequest) Execute() (*AppStoreReviewAttachmentsResponse, *http.Response, error) {
	return r.ApiService.AppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelatedExecute(r)
}

/*
AppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelated Method for AppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiAppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelatedRequest
*/
func (a *AppStoreReviewDetailsAPIService) AppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelated(ctx context.Context, id string) ApiAppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelatedRequest {
	return ApiAppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AppStoreReviewAttachmentsResponse
func (a *AppStoreReviewDetailsAPIService) AppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelatedExecute(r ApiAppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelatedRequest) (*AppStoreReviewAttachmentsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppStoreReviewAttachmentsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppStoreReviewDetailsAPIService.AppStoreReviewDetailsAppStoreReviewAttachmentsGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appStoreReviewDetails/{id}/appStoreReviewAttachments"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsAppStoreReviewDetails != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appStoreReviewDetails]", r.fieldsAppStoreReviewDetails, "csv")
	}
	if r.fieldsAppStoreReviewAttachments != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appStoreReviewAttachments]", r.fieldsAppStoreReviewAttachments, "csv")
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

type ApiAppStoreReviewDetailsCreateInstanceRequest struct {
	ctx context.Context
	ApiService *AppStoreReviewDetailsAPIService
	appStoreReviewDetailCreateRequest *AppStoreReviewDetailCreateRequest
}

// AppStoreReviewDetail representation
func (r ApiAppStoreReviewDetailsCreateInstanceRequest) AppStoreReviewDetailCreateRequest(appStoreReviewDetailCreateRequest AppStoreReviewDetailCreateRequest) ApiAppStoreReviewDetailsCreateInstanceRequest {
	r.appStoreReviewDetailCreateRequest = &appStoreReviewDetailCreateRequest
	return r
}

func (r ApiAppStoreReviewDetailsCreateInstanceRequest) Execute() (*AppStoreReviewDetailResponse, *http.Response, error) {
	return r.ApiService.AppStoreReviewDetailsCreateInstanceExecute(r)
}

/*
AppStoreReviewDetailsCreateInstance Method for AppStoreReviewDetailsCreateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAppStoreReviewDetailsCreateInstanceRequest
*/
func (a *AppStoreReviewDetailsAPIService) AppStoreReviewDetailsCreateInstance(ctx context.Context) ApiAppStoreReviewDetailsCreateInstanceRequest {
	return ApiAppStoreReviewDetailsCreateInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AppStoreReviewDetailResponse
func (a *AppStoreReviewDetailsAPIService) AppStoreReviewDetailsCreateInstanceExecute(r ApiAppStoreReviewDetailsCreateInstanceRequest) (*AppStoreReviewDetailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppStoreReviewDetailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppStoreReviewDetailsAPIService.AppStoreReviewDetailsCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appStoreReviewDetails"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.appStoreReviewDetailCreateRequest == nil {
		return localVarReturnValue, nil, reportError("appStoreReviewDetailCreateRequest is required and must be specified")
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
	localVarPostBody = r.appStoreReviewDetailCreateRequest
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

type ApiAppStoreReviewDetailsGetInstanceRequest struct {
	ctx context.Context
	ApiService *AppStoreReviewDetailsAPIService
	id string
	fieldsAppStoreReviewDetails *[]string
	include *[]string
	fieldsAppStoreReviewAttachments *[]string
	limitAppStoreReviewAttachments *int32
}

// the fields to include for returned resources of type appStoreReviewDetails
func (r ApiAppStoreReviewDetailsGetInstanceRequest) FieldsAppStoreReviewDetails(fieldsAppStoreReviewDetails []string) ApiAppStoreReviewDetailsGetInstanceRequest {
	r.fieldsAppStoreReviewDetails = &fieldsAppStoreReviewDetails
	return r
}

// comma-separated list of relationships to include
func (r ApiAppStoreReviewDetailsGetInstanceRequest) Include(include []string) ApiAppStoreReviewDetailsGetInstanceRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type appStoreReviewAttachments
func (r ApiAppStoreReviewDetailsGetInstanceRequest) FieldsAppStoreReviewAttachments(fieldsAppStoreReviewAttachments []string) ApiAppStoreReviewDetailsGetInstanceRequest {
	r.fieldsAppStoreReviewAttachments = &fieldsAppStoreReviewAttachments
	return r
}

// maximum number of related appStoreReviewAttachments returned (when they are included)
func (r ApiAppStoreReviewDetailsGetInstanceRequest) LimitAppStoreReviewAttachments(limitAppStoreReviewAttachments int32) ApiAppStoreReviewDetailsGetInstanceRequest {
	r.limitAppStoreReviewAttachments = &limitAppStoreReviewAttachments
	return r
}

func (r ApiAppStoreReviewDetailsGetInstanceRequest) Execute() (*AppStoreReviewDetailResponse, *http.Response, error) {
	return r.ApiService.AppStoreReviewDetailsGetInstanceExecute(r)
}

/*
AppStoreReviewDetailsGetInstance Method for AppStoreReviewDetailsGetInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiAppStoreReviewDetailsGetInstanceRequest
*/
func (a *AppStoreReviewDetailsAPIService) AppStoreReviewDetailsGetInstance(ctx context.Context, id string) ApiAppStoreReviewDetailsGetInstanceRequest {
	return ApiAppStoreReviewDetailsGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AppStoreReviewDetailResponse
func (a *AppStoreReviewDetailsAPIService) AppStoreReviewDetailsGetInstanceExecute(r ApiAppStoreReviewDetailsGetInstanceRequest) (*AppStoreReviewDetailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppStoreReviewDetailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppStoreReviewDetailsAPIService.AppStoreReviewDetailsGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appStoreReviewDetails/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsAppStoreReviewDetails != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appStoreReviewDetails]", r.fieldsAppStoreReviewDetails, "csv")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.fieldsAppStoreReviewAttachments != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appStoreReviewAttachments]", r.fieldsAppStoreReviewAttachments, "csv")
	}
	if r.limitAppStoreReviewAttachments != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[appStoreReviewAttachments]", r.limitAppStoreReviewAttachments, "")
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

type ApiAppStoreReviewDetailsUpdateInstanceRequest struct {
	ctx context.Context
	ApiService *AppStoreReviewDetailsAPIService
	id string
	appStoreReviewDetailUpdateRequest *AppStoreReviewDetailUpdateRequest
}

// AppStoreReviewDetail representation
func (r ApiAppStoreReviewDetailsUpdateInstanceRequest) AppStoreReviewDetailUpdateRequest(appStoreReviewDetailUpdateRequest AppStoreReviewDetailUpdateRequest) ApiAppStoreReviewDetailsUpdateInstanceRequest {
	r.appStoreReviewDetailUpdateRequest = &appStoreReviewDetailUpdateRequest
	return r
}

func (r ApiAppStoreReviewDetailsUpdateInstanceRequest) Execute() (*AppStoreReviewDetailResponse, *http.Response, error) {
	return r.ApiService.AppStoreReviewDetailsUpdateInstanceExecute(r)
}

/*
AppStoreReviewDetailsUpdateInstance Method for AppStoreReviewDetailsUpdateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiAppStoreReviewDetailsUpdateInstanceRequest
*/
func (a *AppStoreReviewDetailsAPIService) AppStoreReviewDetailsUpdateInstance(ctx context.Context, id string) ApiAppStoreReviewDetailsUpdateInstanceRequest {
	return ApiAppStoreReviewDetailsUpdateInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AppStoreReviewDetailResponse
func (a *AppStoreReviewDetailsAPIService) AppStoreReviewDetailsUpdateInstanceExecute(r ApiAppStoreReviewDetailsUpdateInstanceRequest) (*AppStoreReviewDetailResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppStoreReviewDetailResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppStoreReviewDetailsAPIService.AppStoreReviewDetailsUpdateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appStoreReviewDetails/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.appStoreReviewDetailUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("appStoreReviewDetailUpdateRequest is required and must be specified")
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
	localVarPostBody = r.appStoreReviewDetailUpdateRequest
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
