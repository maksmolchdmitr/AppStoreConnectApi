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


// SubscriptionAvailabilitiesAPIService SubscriptionAvailabilitiesAPI service
type SubscriptionAvailabilitiesAPIService service

type ApiSubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest struct {
	ctx context.Context
	ApiService *SubscriptionAvailabilitiesAPIService
	id string
	fieldsTerritories *[]string
	limit *int32
}

// the fields to include for returned resources of type territories
func (r ApiSubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest) FieldsTerritories(fieldsTerritories []string) ApiSubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest {
	r.fieldsTerritories = &fieldsTerritories
	return r
}

// maximum resources per page
func (r ApiSubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest) Limit(limit int32) ApiSubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

func (r ApiSubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest) Execute() (*TerritoriesResponse, *http.Response, error) {
	return r.ApiService.SubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelatedExecute(r)
}

/*
SubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelated Method for SubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiSubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest
*/
func (a *SubscriptionAvailabilitiesAPIService) SubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelated(ctx context.Context, id string) ApiSubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest {
	return ApiSubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return TerritoriesResponse
func (a *SubscriptionAvailabilitiesAPIService) SubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelatedExecute(r ApiSubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelatedRequest) (*TerritoriesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TerritoriesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionAvailabilitiesAPIService.SubscriptionAvailabilitiesAvailableTerritoriesGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/subscriptionAvailabilities/{id}/availableTerritories"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsTerritories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[territories]", r.fieldsTerritories, "csv")
	}
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

type ApiSubscriptionAvailabilitiesCreateInstanceRequest struct {
	ctx context.Context
	ApiService *SubscriptionAvailabilitiesAPIService
	subscriptionAvailabilityCreateRequest *SubscriptionAvailabilityCreateRequest
}

// SubscriptionAvailability representation
func (r ApiSubscriptionAvailabilitiesCreateInstanceRequest) SubscriptionAvailabilityCreateRequest(subscriptionAvailabilityCreateRequest SubscriptionAvailabilityCreateRequest) ApiSubscriptionAvailabilitiesCreateInstanceRequest {
	r.subscriptionAvailabilityCreateRequest = &subscriptionAvailabilityCreateRequest
	return r
}

func (r ApiSubscriptionAvailabilitiesCreateInstanceRequest) Execute() (*SubscriptionAvailabilityResponse, *http.Response, error) {
	return r.ApiService.SubscriptionAvailabilitiesCreateInstanceExecute(r)
}

/*
SubscriptionAvailabilitiesCreateInstance Method for SubscriptionAvailabilitiesCreateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSubscriptionAvailabilitiesCreateInstanceRequest
*/
func (a *SubscriptionAvailabilitiesAPIService) SubscriptionAvailabilitiesCreateInstance(ctx context.Context) ApiSubscriptionAvailabilitiesCreateInstanceRequest {
	return ApiSubscriptionAvailabilitiesCreateInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SubscriptionAvailabilityResponse
func (a *SubscriptionAvailabilitiesAPIService) SubscriptionAvailabilitiesCreateInstanceExecute(r ApiSubscriptionAvailabilitiesCreateInstanceRequest) (*SubscriptionAvailabilityResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SubscriptionAvailabilityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionAvailabilitiesAPIService.SubscriptionAvailabilitiesCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/subscriptionAvailabilities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.subscriptionAvailabilityCreateRequest == nil {
		return localVarReturnValue, nil, reportError("subscriptionAvailabilityCreateRequest is required and must be specified")
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
	localVarPostBody = r.subscriptionAvailabilityCreateRequest
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

type ApiSubscriptionAvailabilitiesGetInstanceRequest struct {
	ctx context.Context
	ApiService *SubscriptionAvailabilitiesAPIService
	id string
	fieldsSubscriptionAvailabilities *[]string
	include *[]string
	fieldsTerritories *[]string
	limitAvailableTerritories *int32
}

// the fields to include for returned resources of type subscriptionAvailabilities
func (r ApiSubscriptionAvailabilitiesGetInstanceRequest) FieldsSubscriptionAvailabilities(fieldsSubscriptionAvailabilities []string) ApiSubscriptionAvailabilitiesGetInstanceRequest {
	r.fieldsSubscriptionAvailabilities = &fieldsSubscriptionAvailabilities
	return r
}

// comma-separated list of relationships to include
func (r ApiSubscriptionAvailabilitiesGetInstanceRequest) Include(include []string) ApiSubscriptionAvailabilitiesGetInstanceRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type territories
func (r ApiSubscriptionAvailabilitiesGetInstanceRequest) FieldsTerritories(fieldsTerritories []string) ApiSubscriptionAvailabilitiesGetInstanceRequest {
	r.fieldsTerritories = &fieldsTerritories
	return r
}

// maximum number of related availableTerritories returned (when they are included)
func (r ApiSubscriptionAvailabilitiesGetInstanceRequest) LimitAvailableTerritories(limitAvailableTerritories int32) ApiSubscriptionAvailabilitiesGetInstanceRequest {
	r.limitAvailableTerritories = &limitAvailableTerritories
	return r
}

func (r ApiSubscriptionAvailabilitiesGetInstanceRequest) Execute() (*SubscriptionAvailabilityResponse, *http.Response, error) {
	return r.ApiService.SubscriptionAvailabilitiesGetInstanceExecute(r)
}

/*
SubscriptionAvailabilitiesGetInstance Method for SubscriptionAvailabilitiesGetInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiSubscriptionAvailabilitiesGetInstanceRequest
*/
func (a *SubscriptionAvailabilitiesAPIService) SubscriptionAvailabilitiesGetInstance(ctx context.Context, id string) ApiSubscriptionAvailabilitiesGetInstanceRequest {
	return ApiSubscriptionAvailabilitiesGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SubscriptionAvailabilityResponse
func (a *SubscriptionAvailabilitiesAPIService) SubscriptionAvailabilitiesGetInstanceExecute(r ApiSubscriptionAvailabilitiesGetInstanceRequest) (*SubscriptionAvailabilityResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SubscriptionAvailabilityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionAvailabilitiesAPIService.SubscriptionAvailabilitiesGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/subscriptionAvailabilities/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsSubscriptionAvailabilities != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[subscriptionAvailabilities]", r.fieldsSubscriptionAvailabilities, "csv")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.fieldsTerritories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[territories]", r.fieldsTerritories, "csv")
	}
	if r.limitAvailableTerritories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[availableTerritories]", r.limitAvailableTerritories, "")
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