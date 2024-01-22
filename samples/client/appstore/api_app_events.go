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


// AppEventsAPIService AppEventsAPI service
type AppEventsAPIService service

type ApiAppEventsCreateInstanceRequest struct {
	ctx context.Context
	ApiService *AppEventsAPIService
	appEventCreateRequest *AppEventCreateRequest
}

// AppEvent representation
func (r ApiAppEventsCreateInstanceRequest) AppEventCreateRequest(appEventCreateRequest AppEventCreateRequest) ApiAppEventsCreateInstanceRequest {
	r.appEventCreateRequest = &appEventCreateRequest
	return r
}

func (r ApiAppEventsCreateInstanceRequest) Execute() (*AppEventResponse, *http.Response, error) {
	return r.ApiService.AppEventsCreateInstanceExecute(r)
}

/*
AppEventsCreateInstance Method for AppEventsCreateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAppEventsCreateInstanceRequest
*/
func (a *AppEventsAPIService) AppEventsCreateInstance(ctx context.Context) ApiAppEventsCreateInstanceRequest {
	return ApiAppEventsCreateInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AppEventResponse
func (a *AppEventsAPIService) AppEventsCreateInstanceExecute(r ApiAppEventsCreateInstanceRequest) (*AppEventResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppEventResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppEventsAPIService.AppEventsCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appEvents"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.appEventCreateRequest == nil {
		return localVarReturnValue, nil, reportError("appEventCreateRequest is required and must be specified")
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
	localVarPostBody = r.appEventCreateRequest
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

type ApiAppEventsDeleteInstanceRequest struct {
	ctx context.Context
	ApiService *AppEventsAPIService
	id string
}

func (r ApiAppEventsDeleteInstanceRequest) Execute() (*http.Response, error) {
	return r.ApiService.AppEventsDeleteInstanceExecute(r)
}

/*
AppEventsDeleteInstance Method for AppEventsDeleteInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiAppEventsDeleteInstanceRequest
*/
func (a *AppEventsAPIService) AppEventsDeleteInstance(ctx context.Context, id string) ApiAppEventsDeleteInstanceRequest {
	return ApiAppEventsDeleteInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *AppEventsAPIService) AppEventsDeleteInstanceExecute(r ApiAppEventsDeleteInstanceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppEventsAPIService.AppEventsDeleteInstance")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appEvents/{id}"
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

type ApiAppEventsGetInstanceRequest struct {
	ctx context.Context
	ApiService *AppEventsAPIService
	id string
	fieldsAppEvents *[]string
	include *[]string
	fieldsAppEventLocalizations *[]string
	limitLocalizations *int32
}

// the fields to include for returned resources of type appEvents
func (r ApiAppEventsGetInstanceRequest) FieldsAppEvents(fieldsAppEvents []string) ApiAppEventsGetInstanceRequest {
	r.fieldsAppEvents = &fieldsAppEvents
	return r
}

// comma-separated list of relationships to include
func (r ApiAppEventsGetInstanceRequest) Include(include []string) ApiAppEventsGetInstanceRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type appEventLocalizations
func (r ApiAppEventsGetInstanceRequest) FieldsAppEventLocalizations(fieldsAppEventLocalizations []string) ApiAppEventsGetInstanceRequest {
	r.fieldsAppEventLocalizations = &fieldsAppEventLocalizations
	return r
}

// maximum number of related localizations returned (when they are included)
func (r ApiAppEventsGetInstanceRequest) LimitLocalizations(limitLocalizations int32) ApiAppEventsGetInstanceRequest {
	r.limitLocalizations = &limitLocalizations
	return r
}

func (r ApiAppEventsGetInstanceRequest) Execute() (*AppEventResponse, *http.Response, error) {
	return r.ApiService.AppEventsGetInstanceExecute(r)
}

/*
AppEventsGetInstance Method for AppEventsGetInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiAppEventsGetInstanceRequest
*/
func (a *AppEventsAPIService) AppEventsGetInstance(ctx context.Context, id string) ApiAppEventsGetInstanceRequest {
	return ApiAppEventsGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AppEventResponse
func (a *AppEventsAPIService) AppEventsGetInstanceExecute(r ApiAppEventsGetInstanceRequest) (*AppEventResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppEventResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppEventsAPIService.AppEventsGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appEvents/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsAppEvents != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appEvents]", r.fieldsAppEvents, "csv")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.fieldsAppEventLocalizations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appEventLocalizations]", r.fieldsAppEventLocalizations, "csv")
	}
	if r.limitLocalizations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[localizations]", r.limitLocalizations, "")
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

type ApiAppEventsLocalizationsGetToManyRelatedRequest struct {
	ctx context.Context
	ApiService *AppEventsAPIService
	id string
	fieldsAppEventScreenshots *[]string
	fieldsAppEventVideoClips *[]string
	fieldsAppEventLocalizations *[]string
	fieldsAppEvents *[]string
	limit *int32
	limitAppEventScreenshots *int32
	limitAppEventVideoClips *int32
	include *[]string
}

// the fields to include for returned resources of type appEventScreenshots
func (r ApiAppEventsLocalizationsGetToManyRelatedRequest) FieldsAppEventScreenshots(fieldsAppEventScreenshots []string) ApiAppEventsLocalizationsGetToManyRelatedRequest {
	r.fieldsAppEventScreenshots = &fieldsAppEventScreenshots
	return r
}

// the fields to include for returned resources of type appEventVideoClips
func (r ApiAppEventsLocalizationsGetToManyRelatedRequest) FieldsAppEventVideoClips(fieldsAppEventVideoClips []string) ApiAppEventsLocalizationsGetToManyRelatedRequest {
	r.fieldsAppEventVideoClips = &fieldsAppEventVideoClips
	return r
}

// the fields to include for returned resources of type appEventLocalizations
func (r ApiAppEventsLocalizationsGetToManyRelatedRequest) FieldsAppEventLocalizations(fieldsAppEventLocalizations []string) ApiAppEventsLocalizationsGetToManyRelatedRequest {
	r.fieldsAppEventLocalizations = &fieldsAppEventLocalizations
	return r
}

// the fields to include for returned resources of type appEvents
func (r ApiAppEventsLocalizationsGetToManyRelatedRequest) FieldsAppEvents(fieldsAppEvents []string) ApiAppEventsLocalizationsGetToManyRelatedRequest {
	r.fieldsAppEvents = &fieldsAppEvents
	return r
}

// maximum resources per page
func (r ApiAppEventsLocalizationsGetToManyRelatedRequest) Limit(limit int32) ApiAppEventsLocalizationsGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

// maximum number of related appEventScreenshots returned (when they are included)
func (r ApiAppEventsLocalizationsGetToManyRelatedRequest) LimitAppEventScreenshots(limitAppEventScreenshots int32) ApiAppEventsLocalizationsGetToManyRelatedRequest {
	r.limitAppEventScreenshots = &limitAppEventScreenshots
	return r
}

// maximum number of related appEventVideoClips returned (when they are included)
func (r ApiAppEventsLocalizationsGetToManyRelatedRequest) LimitAppEventVideoClips(limitAppEventVideoClips int32) ApiAppEventsLocalizationsGetToManyRelatedRequest {
	r.limitAppEventVideoClips = &limitAppEventVideoClips
	return r
}

// comma-separated list of relationships to include
func (r ApiAppEventsLocalizationsGetToManyRelatedRequest) Include(include []string) ApiAppEventsLocalizationsGetToManyRelatedRequest {
	r.include = &include
	return r
}

func (r ApiAppEventsLocalizationsGetToManyRelatedRequest) Execute() (*AppEventLocalizationsResponse, *http.Response, error) {
	return r.ApiService.AppEventsLocalizationsGetToManyRelatedExecute(r)
}

/*
AppEventsLocalizationsGetToManyRelated Method for AppEventsLocalizationsGetToManyRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiAppEventsLocalizationsGetToManyRelatedRequest
*/
func (a *AppEventsAPIService) AppEventsLocalizationsGetToManyRelated(ctx context.Context, id string) ApiAppEventsLocalizationsGetToManyRelatedRequest {
	return ApiAppEventsLocalizationsGetToManyRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AppEventLocalizationsResponse
func (a *AppEventsAPIService) AppEventsLocalizationsGetToManyRelatedExecute(r ApiAppEventsLocalizationsGetToManyRelatedRequest) (*AppEventLocalizationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppEventLocalizationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppEventsAPIService.AppEventsLocalizationsGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appEvents/{id}/localizations"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsAppEventScreenshots != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appEventScreenshots]", r.fieldsAppEventScreenshots, "csv")
	}
	if r.fieldsAppEventVideoClips != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appEventVideoClips]", r.fieldsAppEventVideoClips, "csv")
	}
	if r.fieldsAppEventLocalizations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appEventLocalizations]", r.fieldsAppEventLocalizations, "csv")
	}
	if r.fieldsAppEvents != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appEvents]", r.fieldsAppEvents, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.limitAppEventScreenshots != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[appEventScreenshots]", r.limitAppEventScreenshots, "")
	}
	if r.limitAppEventVideoClips != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[appEventVideoClips]", r.limitAppEventVideoClips, "")
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

type ApiAppEventsUpdateInstanceRequest struct {
	ctx context.Context
	ApiService *AppEventsAPIService
	id string
	appEventUpdateRequest *AppEventUpdateRequest
}

// AppEvent representation
func (r ApiAppEventsUpdateInstanceRequest) AppEventUpdateRequest(appEventUpdateRequest AppEventUpdateRequest) ApiAppEventsUpdateInstanceRequest {
	r.appEventUpdateRequest = &appEventUpdateRequest
	return r
}

func (r ApiAppEventsUpdateInstanceRequest) Execute() (*AppEventResponse, *http.Response, error) {
	return r.ApiService.AppEventsUpdateInstanceExecute(r)
}

/*
AppEventsUpdateInstance Method for AppEventsUpdateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiAppEventsUpdateInstanceRequest
*/
func (a *AppEventsAPIService) AppEventsUpdateInstance(ctx context.Context, id string) ApiAppEventsUpdateInstanceRequest {
	return ApiAppEventsUpdateInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AppEventResponse
func (a *AppEventsAPIService) AppEventsUpdateInstanceExecute(r ApiAppEventsUpdateInstanceRequest) (*AppEventResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppEventResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppEventsAPIService.AppEventsUpdateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appEvents/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.appEventUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("appEventUpdateRequest is required and must be specified")
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
	localVarPostBody = r.appEventUpdateRequest
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
