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

// BetaBuildLocalizationsAPIService BetaBuildLocalizationsAPI service
type BetaBuildLocalizationsAPIService service

type ApiBetaBuildLocalizationsBuildGetToOneRelatedRequest struct {
	ctx          context.Context
	ApiService   *BetaBuildLocalizationsAPIService
	id           string
	fieldsBuilds *[]string
}

// the fields to include for returned resources of type builds
func (r ApiBetaBuildLocalizationsBuildGetToOneRelatedRequest) FieldsBuilds(fieldsBuilds []string) ApiBetaBuildLocalizationsBuildGetToOneRelatedRequest {
	r.fieldsBuilds = &fieldsBuilds
	return r
}

func (r ApiBetaBuildLocalizationsBuildGetToOneRelatedRequest) Execute() (*BuildWithoutIncludesResponse, *http.Response, error) {
	return r.ApiService.BetaBuildLocalizationsBuildGetToOneRelatedExecute(r)
}

/*
BetaBuildLocalizationsBuildGetToOneRelated Method for BetaBuildLocalizationsBuildGetToOneRelated

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiBetaBuildLocalizationsBuildGetToOneRelatedRequest
*/
func (a *BetaBuildLocalizationsAPIService) BetaBuildLocalizationsBuildGetToOneRelated(ctx context.Context, id string) ApiBetaBuildLocalizationsBuildGetToOneRelatedRequest {
	return ApiBetaBuildLocalizationsBuildGetToOneRelatedRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return BuildWithoutIncludesResponse
func (a *BetaBuildLocalizationsAPIService) BetaBuildLocalizationsBuildGetToOneRelatedExecute(r ApiBetaBuildLocalizationsBuildGetToOneRelatedRequest) (*BuildWithoutIncludesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BuildWithoutIncludesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaBuildLocalizationsAPIService.BetaBuildLocalizationsBuildGetToOneRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaBuildLocalizations/{id}/build"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsBuilds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[builds]", r.fieldsBuilds, "csv")
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

type ApiBetaBuildLocalizationsCreateInstanceRequest struct {
	ctx                                context.Context
	ApiService                         *BetaBuildLocalizationsAPIService
	betaBuildLocalizationCreateRequest *BetaBuildLocalizationCreateRequest
}

// BetaBuildLocalization representation
func (r ApiBetaBuildLocalizationsCreateInstanceRequest) BetaBuildLocalizationCreateRequest(betaBuildLocalizationCreateRequest BetaBuildLocalizationCreateRequest) ApiBetaBuildLocalizationsCreateInstanceRequest {
	r.betaBuildLocalizationCreateRequest = &betaBuildLocalizationCreateRequest
	return r
}

func (r ApiBetaBuildLocalizationsCreateInstanceRequest) Execute() (*BetaBuildLocalizationResponse, *http.Response, error) {
	return r.ApiService.BetaBuildLocalizationsCreateInstanceExecute(r)
}

/*
BetaBuildLocalizationsCreateInstance Method for BetaBuildLocalizationsCreateInstance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiBetaBuildLocalizationsCreateInstanceRequest
*/
func (a *BetaBuildLocalizationsAPIService) BetaBuildLocalizationsCreateInstance(ctx context.Context) ApiBetaBuildLocalizationsCreateInstanceRequest {
	return ApiBetaBuildLocalizationsCreateInstanceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BetaBuildLocalizationResponse
func (a *BetaBuildLocalizationsAPIService) BetaBuildLocalizationsCreateInstanceExecute(r ApiBetaBuildLocalizationsCreateInstanceRequest) (*BetaBuildLocalizationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BetaBuildLocalizationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaBuildLocalizationsAPIService.BetaBuildLocalizationsCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaBuildLocalizations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.betaBuildLocalizationCreateRequest == nil {
		return localVarReturnValue, nil, reportError("betaBuildLocalizationCreateRequest is required and must be specified")
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
	localVarPostBody = r.betaBuildLocalizationCreateRequest
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

type ApiBetaBuildLocalizationsDeleteInstanceRequest struct {
	ctx        context.Context
	ApiService *BetaBuildLocalizationsAPIService
	id         string
}

func (r ApiBetaBuildLocalizationsDeleteInstanceRequest) Execute() (*http.Response, error) {
	return r.ApiService.BetaBuildLocalizationsDeleteInstanceExecute(r)
}

/*
BetaBuildLocalizationsDeleteInstance Method for BetaBuildLocalizationsDeleteInstance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiBetaBuildLocalizationsDeleteInstanceRequest
*/
func (a *BetaBuildLocalizationsAPIService) BetaBuildLocalizationsDeleteInstance(ctx context.Context, id string) ApiBetaBuildLocalizationsDeleteInstanceRequest {
	return ApiBetaBuildLocalizationsDeleteInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *BetaBuildLocalizationsAPIService) BetaBuildLocalizationsDeleteInstanceExecute(r ApiBetaBuildLocalizationsDeleteInstanceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaBuildLocalizationsAPIService.BetaBuildLocalizationsDeleteInstance")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaBuildLocalizations/{id}"
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

type ApiBetaBuildLocalizationsGetCollectionRequest struct {
	ctx                          context.Context
	ApiService                   *BetaBuildLocalizationsAPIService
	filterLocale                 *[]string
	filterBuild                  *[]string
	fieldsBetaBuildLocalizations *[]string
	limit                        *int32
	include                      *[]string
	fieldsBuilds                 *[]string
}

// filter by attribute &#39;locale&#39;
func (r ApiBetaBuildLocalizationsGetCollectionRequest) FilterLocale(filterLocale []string) ApiBetaBuildLocalizationsGetCollectionRequest {
	r.filterLocale = &filterLocale
	return r
}

// filter by id(s) of related &#39;build&#39;
func (r ApiBetaBuildLocalizationsGetCollectionRequest) FilterBuild(filterBuild []string) ApiBetaBuildLocalizationsGetCollectionRequest {
	r.filterBuild = &filterBuild
	return r
}

// the fields to include for returned resources of type betaBuildLocalizations
func (r ApiBetaBuildLocalizationsGetCollectionRequest) FieldsBetaBuildLocalizations(fieldsBetaBuildLocalizations []string) ApiBetaBuildLocalizationsGetCollectionRequest {
	r.fieldsBetaBuildLocalizations = &fieldsBetaBuildLocalizations
	return r
}

// maximum resources per page
func (r ApiBetaBuildLocalizationsGetCollectionRequest) Limit(limit int32) ApiBetaBuildLocalizationsGetCollectionRequest {
	r.limit = &limit
	return r
}

// comma-separated list of relationships to include
func (r ApiBetaBuildLocalizationsGetCollectionRequest) Include(include []string) ApiBetaBuildLocalizationsGetCollectionRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type builds
func (r ApiBetaBuildLocalizationsGetCollectionRequest) FieldsBuilds(fieldsBuilds []string) ApiBetaBuildLocalizationsGetCollectionRequest {
	r.fieldsBuilds = &fieldsBuilds
	return r
}

func (r ApiBetaBuildLocalizationsGetCollectionRequest) Execute() (*BetaBuildLocalizationsResponse, *http.Response, error) {
	return r.ApiService.BetaBuildLocalizationsGetCollectionExecute(r)
}

/*
BetaBuildLocalizationsGetCollection Method for BetaBuildLocalizationsGetCollection

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiBetaBuildLocalizationsGetCollectionRequest
*/
func (a *BetaBuildLocalizationsAPIService) BetaBuildLocalizationsGetCollection(ctx context.Context) ApiBetaBuildLocalizationsGetCollectionRequest {
	return ApiBetaBuildLocalizationsGetCollectionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return BetaBuildLocalizationsResponse
func (a *BetaBuildLocalizationsAPIService) BetaBuildLocalizationsGetCollectionExecute(r ApiBetaBuildLocalizationsGetCollectionRequest) (*BetaBuildLocalizationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BetaBuildLocalizationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaBuildLocalizationsAPIService.BetaBuildLocalizationsGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaBuildLocalizations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterLocale != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[locale]", r.filterLocale, "csv")
	}
	if r.filterBuild != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[build]", r.filterBuild, "csv")
	}
	if r.fieldsBetaBuildLocalizations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[betaBuildLocalizations]", r.fieldsBetaBuildLocalizations, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.fieldsBuilds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[builds]", r.fieldsBuilds, "csv")
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

type ApiBetaBuildLocalizationsGetInstanceRequest struct {
	ctx                          context.Context
	ApiService                   *BetaBuildLocalizationsAPIService
	id                           string
	fieldsBetaBuildLocalizations *[]string
	include                      *[]string
	fieldsBuilds                 *[]string
}

// the fields to include for returned resources of type betaBuildLocalizations
func (r ApiBetaBuildLocalizationsGetInstanceRequest) FieldsBetaBuildLocalizations(fieldsBetaBuildLocalizations []string) ApiBetaBuildLocalizationsGetInstanceRequest {
	r.fieldsBetaBuildLocalizations = &fieldsBetaBuildLocalizations
	return r
}

// comma-separated list of relationships to include
func (r ApiBetaBuildLocalizationsGetInstanceRequest) Include(include []string) ApiBetaBuildLocalizationsGetInstanceRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type builds
func (r ApiBetaBuildLocalizationsGetInstanceRequest) FieldsBuilds(fieldsBuilds []string) ApiBetaBuildLocalizationsGetInstanceRequest {
	r.fieldsBuilds = &fieldsBuilds
	return r
}

func (r ApiBetaBuildLocalizationsGetInstanceRequest) Execute() (*BetaBuildLocalizationResponse, *http.Response, error) {
	return r.ApiService.BetaBuildLocalizationsGetInstanceExecute(r)
}

/*
BetaBuildLocalizationsGetInstance Method for BetaBuildLocalizationsGetInstance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiBetaBuildLocalizationsGetInstanceRequest
*/
func (a *BetaBuildLocalizationsAPIService) BetaBuildLocalizationsGetInstance(ctx context.Context, id string) ApiBetaBuildLocalizationsGetInstanceRequest {
	return ApiBetaBuildLocalizationsGetInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return BetaBuildLocalizationResponse
func (a *BetaBuildLocalizationsAPIService) BetaBuildLocalizationsGetInstanceExecute(r ApiBetaBuildLocalizationsGetInstanceRequest) (*BetaBuildLocalizationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BetaBuildLocalizationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaBuildLocalizationsAPIService.BetaBuildLocalizationsGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaBuildLocalizations/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsBetaBuildLocalizations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[betaBuildLocalizations]", r.fieldsBetaBuildLocalizations, "csv")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.fieldsBuilds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[builds]", r.fieldsBuilds, "csv")
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

type ApiBetaBuildLocalizationsUpdateInstanceRequest struct {
	ctx                                context.Context
	ApiService                         *BetaBuildLocalizationsAPIService
	id                                 string
	betaBuildLocalizationUpdateRequest *BetaBuildLocalizationUpdateRequest
}

// BetaBuildLocalization representation
func (r ApiBetaBuildLocalizationsUpdateInstanceRequest) BetaBuildLocalizationUpdateRequest(betaBuildLocalizationUpdateRequest BetaBuildLocalizationUpdateRequest) ApiBetaBuildLocalizationsUpdateInstanceRequest {
	r.betaBuildLocalizationUpdateRequest = &betaBuildLocalizationUpdateRequest
	return r
}

func (r ApiBetaBuildLocalizationsUpdateInstanceRequest) Execute() (*BetaBuildLocalizationResponse, *http.Response, error) {
	return r.ApiService.BetaBuildLocalizationsUpdateInstanceExecute(r)
}

/*
BetaBuildLocalizationsUpdateInstance Method for BetaBuildLocalizationsUpdateInstance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiBetaBuildLocalizationsUpdateInstanceRequest
*/
func (a *BetaBuildLocalizationsAPIService) BetaBuildLocalizationsUpdateInstance(ctx context.Context, id string) ApiBetaBuildLocalizationsUpdateInstanceRequest {
	return ApiBetaBuildLocalizationsUpdateInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return BetaBuildLocalizationResponse
func (a *BetaBuildLocalizationsAPIService) BetaBuildLocalizationsUpdateInstanceExecute(r ApiBetaBuildLocalizationsUpdateInstanceRequest) (*BetaBuildLocalizationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BetaBuildLocalizationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaBuildLocalizationsAPIService.BetaBuildLocalizationsUpdateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaBuildLocalizations/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.betaBuildLocalizationUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("betaBuildLocalizationUpdateRequest is required and must be specified")
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
	localVarPostBody = r.betaBuildLocalizationUpdateRequest
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
