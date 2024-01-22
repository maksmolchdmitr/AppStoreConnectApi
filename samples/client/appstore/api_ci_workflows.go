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


// CiWorkflowsAPIService CiWorkflowsAPI service
type CiWorkflowsAPIService service

type ApiCiWorkflowsBuildRunsGetToManyRelatedRequest struct {
	ctx context.Context
	ApiService *CiWorkflowsAPIService
	id string
	filterBuilds *[]string
	fieldsScmGitReferences *[]string
	fieldsCiBuildRuns *[]string
	fieldsCiWorkflows *[]string
	fieldsScmPullRequests *[]string
	fieldsCiProducts *[]string
	fieldsBuilds *[]string
	limit *int32
	limitBuilds *int32
	include *[]string
}

// filter by id(s) of related &#39;builds&#39;
func (r ApiCiWorkflowsBuildRunsGetToManyRelatedRequest) FilterBuilds(filterBuilds []string) ApiCiWorkflowsBuildRunsGetToManyRelatedRequest {
	r.filterBuilds = &filterBuilds
	return r
}

// the fields to include for returned resources of type scmGitReferences
func (r ApiCiWorkflowsBuildRunsGetToManyRelatedRequest) FieldsScmGitReferences(fieldsScmGitReferences []string) ApiCiWorkflowsBuildRunsGetToManyRelatedRequest {
	r.fieldsScmGitReferences = &fieldsScmGitReferences
	return r
}

// the fields to include for returned resources of type ciBuildRuns
func (r ApiCiWorkflowsBuildRunsGetToManyRelatedRequest) FieldsCiBuildRuns(fieldsCiBuildRuns []string) ApiCiWorkflowsBuildRunsGetToManyRelatedRequest {
	r.fieldsCiBuildRuns = &fieldsCiBuildRuns
	return r
}

// the fields to include for returned resources of type ciWorkflows
func (r ApiCiWorkflowsBuildRunsGetToManyRelatedRequest) FieldsCiWorkflows(fieldsCiWorkflows []string) ApiCiWorkflowsBuildRunsGetToManyRelatedRequest {
	r.fieldsCiWorkflows = &fieldsCiWorkflows
	return r
}

// the fields to include for returned resources of type scmPullRequests
func (r ApiCiWorkflowsBuildRunsGetToManyRelatedRequest) FieldsScmPullRequests(fieldsScmPullRequests []string) ApiCiWorkflowsBuildRunsGetToManyRelatedRequest {
	r.fieldsScmPullRequests = &fieldsScmPullRequests
	return r
}

// the fields to include for returned resources of type ciProducts
func (r ApiCiWorkflowsBuildRunsGetToManyRelatedRequest) FieldsCiProducts(fieldsCiProducts []string) ApiCiWorkflowsBuildRunsGetToManyRelatedRequest {
	r.fieldsCiProducts = &fieldsCiProducts
	return r
}

// the fields to include for returned resources of type builds
func (r ApiCiWorkflowsBuildRunsGetToManyRelatedRequest) FieldsBuilds(fieldsBuilds []string) ApiCiWorkflowsBuildRunsGetToManyRelatedRequest {
	r.fieldsBuilds = &fieldsBuilds
	return r
}

// maximum resources per page
func (r ApiCiWorkflowsBuildRunsGetToManyRelatedRequest) Limit(limit int32) ApiCiWorkflowsBuildRunsGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

// maximum number of related builds returned (when they are included)
func (r ApiCiWorkflowsBuildRunsGetToManyRelatedRequest) LimitBuilds(limitBuilds int32) ApiCiWorkflowsBuildRunsGetToManyRelatedRequest {
	r.limitBuilds = &limitBuilds
	return r
}

// comma-separated list of relationships to include
func (r ApiCiWorkflowsBuildRunsGetToManyRelatedRequest) Include(include []string) ApiCiWorkflowsBuildRunsGetToManyRelatedRequest {
	r.include = &include
	return r
}

func (r ApiCiWorkflowsBuildRunsGetToManyRelatedRequest) Execute() (*CiBuildRunsResponse, *http.Response, error) {
	return r.ApiService.CiWorkflowsBuildRunsGetToManyRelatedExecute(r)
}

/*
CiWorkflowsBuildRunsGetToManyRelated Method for CiWorkflowsBuildRunsGetToManyRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiCiWorkflowsBuildRunsGetToManyRelatedRequest
*/
func (a *CiWorkflowsAPIService) CiWorkflowsBuildRunsGetToManyRelated(ctx context.Context, id string) ApiCiWorkflowsBuildRunsGetToManyRelatedRequest {
	return ApiCiWorkflowsBuildRunsGetToManyRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CiBuildRunsResponse
func (a *CiWorkflowsAPIService) CiWorkflowsBuildRunsGetToManyRelatedExecute(r ApiCiWorkflowsBuildRunsGetToManyRelatedRequest) (*CiBuildRunsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CiBuildRunsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CiWorkflowsAPIService.CiWorkflowsBuildRunsGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/ciWorkflows/{id}/buildRuns"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterBuilds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[builds]", r.filterBuilds, "csv")
	}
	if r.fieldsScmGitReferences != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[scmGitReferences]", r.fieldsScmGitReferences, "csv")
	}
	if r.fieldsCiBuildRuns != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[ciBuildRuns]", r.fieldsCiBuildRuns, "csv")
	}
	if r.fieldsCiWorkflows != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[ciWorkflows]", r.fieldsCiWorkflows, "csv")
	}
	if r.fieldsScmPullRequests != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[scmPullRequests]", r.fieldsScmPullRequests, "csv")
	}
	if r.fieldsCiProducts != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[ciProducts]", r.fieldsCiProducts, "csv")
	}
	if r.fieldsBuilds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[builds]", r.fieldsBuilds, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.limitBuilds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[builds]", r.limitBuilds, "")
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

type ApiCiWorkflowsCreateInstanceRequest struct {
	ctx context.Context
	ApiService *CiWorkflowsAPIService
	ciWorkflowCreateRequest *CiWorkflowCreateRequest
}

// CiWorkflow representation
func (r ApiCiWorkflowsCreateInstanceRequest) CiWorkflowCreateRequest(ciWorkflowCreateRequest CiWorkflowCreateRequest) ApiCiWorkflowsCreateInstanceRequest {
	r.ciWorkflowCreateRequest = &ciWorkflowCreateRequest
	return r
}

func (r ApiCiWorkflowsCreateInstanceRequest) Execute() (*CiWorkflowResponse, *http.Response, error) {
	return r.ApiService.CiWorkflowsCreateInstanceExecute(r)
}

/*
CiWorkflowsCreateInstance Method for CiWorkflowsCreateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCiWorkflowsCreateInstanceRequest
*/
func (a *CiWorkflowsAPIService) CiWorkflowsCreateInstance(ctx context.Context) ApiCiWorkflowsCreateInstanceRequest {
	return ApiCiWorkflowsCreateInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CiWorkflowResponse
func (a *CiWorkflowsAPIService) CiWorkflowsCreateInstanceExecute(r ApiCiWorkflowsCreateInstanceRequest) (*CiWorkflowResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CiWorkflowResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CiWorkflowsAPIService.CiWorkflowsCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/ciWorkflows"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ciWorkflowCreateRequest == nil {
		return localVarReturnValue, nil, reportError("ciWorkflowCreateRequest is required and must be specified")
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
	localVarPostBody = r.ciWorkflowCreateRequest
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

type ApiCiWorkflowsDeleteInstanceRequest struct {
	ctx context.Context
	ApiService *CiWorkflowsAPIService
	id string
}

func (r ApiCiWorkflowsDeleteInstanceRequest) Execute() (*http.Response, error) {
	return r.ApiService.CiWorkflowsDeleteInstanceExecute(r)
}

/*
CiWorkflowsDeleteInstance Method for CiWorkflowsDeleteInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiCiWorkflowsDeleteInstanceRequest
*/
func (a *CiWorkflowsAPIService) CiWorkflowsDeleteInstance(ctx context.Context, id string) ApiCiWorkflowsDeleteInstanceRequest {
	return ApiCiWorkflowsDeleteInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *CiWorkflowsAPIService) CiWorkflowsDeleteInstanceExecute(r ApiCiWorkflowsDeleteInstanceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CiWorkflowsAPIService.CiWorkflowsDeleteInstance")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/ciWorkflows/{id}"
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

type ApiCiWorkflowsGetInstanceRequest struct {
	ctx context.Context
	ApiService *CiWorkflowsAPIService
	id string
	fieldsCiWorkflows *[]string
	include *[]string
	fieldsCiBuildRuns *[]string
	fieldsScmRepositories *[]string
}

// the fields to include for returned resources of type ciWorkflows
func (r ApiCiWorkflowsGetInstanceRequest) FieldsCiWorkflows(fieldsCiWorkflows []string) ApiCiWorkflowsGetInstanceRequest {
	r.fieldsCiWorkflows = &fieldsCiWorkflows
	return r
}

// comma-separated list of relationships to include
func (r ApiCiWorkflowsGetInstanceRequest) Include(include []string) ApiCiWorkflowsGetInstanceRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type ciBuildRuns
func (r ApiCiWorkflowsGetInstanceRequest) FieldsCiBuildRuns(fieldsCiBuildRuns []string) ApiCiWorkflowsGetInstanceRequest {
	r.fieldsCiBuildRuns = &fieldsCiBuildRuns
	return r
}

// the fields to include for returned resources of type scmRepositories
func (r ApiCiWorkflowsGetInstanceRequest) FieldsScmRepositories(fieldsScmRepositories []string) ApiCiWorkflowsGetInstanceRequest {
	r.fieldsScmRepositories = &fieldsScmRepositories
	return r
}

func (r ApiCiWorkflowsGetInstanceRequest) Execute() (*CiWorkflowResponse, *http.Response, error) {
	return r.ApiService.CiWorkflowsGetInstanceExecute(r)
}

/*
CiWorkflowsGetInstance Method for CiWorkflowsGetInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiCiWorkflowsGetInstanceRequest
*/
func (a *CiWorkflowsAPIService) CiWorkflowsGetInstance(ctx context.Context, id string) ApiCiWorkflowsGetInstanceRequest {
	return ApiCiWorkflowsGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CiWorkflowResponse
func (a *CiWorkflowsAPIService) CiWorkflowsGetInstanceExecute(r ApiCiWorkflowsGetInstanceRequest) (*CiWorkflowResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CiWorkflowResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CiWorkflowsAPIService.CiWorkflowsGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/ciWorkflows/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsCiWorkflows != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[ciWorkflows]", r.fieldsCiWorkflows, "csv")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.fieldsCiBuildRuns != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[ciBuildRuns]", r.fieldsCiBuildRuns, "csv")
	}
	if r.fieldsScmRepositories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[scmRepositories]", r.fieldsScmRepositories, "csv")
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

type ApiCiWorkflowsRepositoryGetToOneRelatedRequest struct {
	ctx context.Context
	ApiService *CiWorkflowsAPIService
	id string
	fieldsScmGitReferences *[]string
	fieldsScmProviders *[]string
	fieldsScmRepositories *[]string
	include *[]string
}

// the fields to include for returned resources of type scmGitReferences
func (r ApiCiWorkflowsRepositoryGetToOneRelatedRequest) FieldsScmGitReferences(fieldsScmGitReferences []string) ApiCiWorkflowsRepositoryGetToOneRelatedRequest {
	r.fieldsScmGitReferences = &fieldsScmGitReferences
	return r
}

// the fields to include for returned resources of type scmProviders
func (r ApiCiWorkflowsRepositoryGetToOneRelatedRequest) FieldsScmProviders(fieldsScmProviders []string) ApiCiWorkflowsRepositoryGetToOneRelatedRequest {
	r.fieldsScmProviders = &fieldsScmProviders
	return r
}

// the fields to include for returned resources of type scmRepositories
func (r ApiCiWorkflowsRepositoryGetToOneRelatedRequest) FieldsScmRepositories(fieldsScmRepositories []string) ApiCiWorkflowsRepositoryGetToOneRelatedRequest {
	r.fieldsScmRepositories = &fieldsScmRepositories
	return r
}

// comma-separated list of relationships to include
func (r ApiCiWorkflowsRepositoryGetToOneRelatedRequest) Include(include []string) ApiCiWorkflowsRepositoryGetToOneRelatedRequest {
	r.include = &include
	return r
}

func (r ApiCiWorkflowsRepositoryGetToOneRelatedRequest) Execute() (*ScmRepositoryResponse, *http.Response, error) {
	return r.ApiService.CiWorkflowsRepositoryGetToOneRelatedExecute(r)
}

/*
CiWorkflowsRepositoryGetToOneRelated Method for CiWorkflowsRepositoryGetToOneRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiCiWorkflowsRepositoryGetToOneRelatedRequest
*/
func (a *CiWorkflowsAPIService) CiWorkflowsRepositoryGetToOneRelated(ctx context.Context, id string) ApiCiWorkflowsRepositoryGetToOneRelatedRequest {
	return ApiCiWorkflowsRepositoryGetToOneRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ScmRepositoryResponse
func (a *CiWorkflowsAPIService) CiWorkflowsRepositoryGetToOneRelatedExecute(r ApiCiWorkflowsRepositoryGetToOneRelatedRequest) (*ScmRepositoryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ScmRepositoryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CiWorkflowsAPIService.CiWorkflowsRepositoryGetToOneRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/ciWorkflows/{id}/repository"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsScmGitReferences != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[scmGitReferences]", r.fieldsScmGitReferences, "csv")
	}
	if r.fieldsScmProviders != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[scmProviders]", r.fieldsScmProviders, "csv")
	}
	if r.fieldsScmRepositories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[scmRepositories]", r.fieldsScmRepositories, "csv")
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

type ApiCiWorkflowsUpdateInstanceRequest struct {
	ctx context.Context
	ApiService *CiWorkflowsAPIService
	id string
	ciWorkflowUpdateRequest *CiWorkflowUpdateRequest
}

// CiWorkflow representation
func (r ApiCiWorkflowsUpdateInstanceRequest) CiWorkflowUpdateRequest(ciWorkflowUpdateRequest CiWorkflowUpdateRequest) ApiCiWorkflowsUpdateInstanceRequest {
	r.ciWorkflowUpdateRequest = &ciWorkflowUpdateRequest
	return r
}

func (r ApiCiWorkflowsUpdateInstanceRequest) Execute() (*CiWorkflowResponse, *http.Response, error) {
	return r.ApiService.CiWorkflowsUpdateInstanceExecute(r)
}

/*
CiWorkflowsUpdateInstance Method for CiWorkflowsUpdateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiCiWorkflowsUpdateInstanceRequest
*/
func (a *CiWorkflowsAPIService) CiWorkflowsUpdateInstance(ctx context.Context, id string) ApiCiWorkflowsUpdateInstanceRequest {
	return ApiCiWorkflowsUpdateInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CiWorkflowResponse
func (a *CiWorkflowsAPIService) CiWorkflowsUpdateInstanceExecute(r ApiCiWorkflowsUpdateInstanceRequest) (*CiWorkflowResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CiWorkflowResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CiWorkflowsAPIService.CiWorkflowsUpdateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/ciWorkflows/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ciWorkflowUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("ciWorkflowUpdateRequest is required and must be specified")
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
	localVarPostBody = r.ciWorkflowUpdateRequest
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
