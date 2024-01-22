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


// BetaLicenseAgreementsAPIService BetaLicenseAgreementsAPI service
type BetaLicenseAgreementsAPIService service

type ApiBetaLicenseAgreementsAppGetToOneRelatedRequest struct {
	ctx context.Context
	ApiService *BetaLicenseAgreementsAPIService
	id string
	fieldsApps *[]string
}

// the fields to include for returned resources of type apps
func (r ApiBetaLicenseAgreementsAppGetToOneRelatedRequest) FieldsApps(fieldsApps []string) ApiBetaLicenseAgreementsAppGetToOneRelatedRequest {
	r.fieldsApps = &fieldsApps
	return r
}

func (r ApiBetaLicenseAgreementsAppGetToOneRelatedRequest) Execute() (*AppWithoutIncludesResponse, *http.Response, error) {
	return r.ApiService.BetaLicenseAgreementsAppGetToOneRelatedExecute(r)
}

/*
BetaLicenseAgreementsAppGetToOneRelated Method for BetaLicenseAgreementsAppGetToOneRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiBetaLicenseAgreementsAppGetToOneRelatedRequest
*/
func (a *BetaLicenseAgreementsAPIService) BetaLicenseAgreementsAppGetToOneRelated(ctx context.Context, id string) ApiBetaLicenseAgreementsAppGetToOneRelatedRequest {
	return ApiBetaLicenseAgreementsAppGetToOneRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AppWithoutIncludesResponse
func (a *BetaLicenseAgreementsAPIService) BetaLicenseAgreementsAppGetToOneRelatedExecute(r ApiBetaLicenseAgreementsAppGetToOneRelatedRequest) (*AppWithoutIncludesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AppWithoutIncludesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaLicenseAgreementsAPIService.BetaLicenseAgreementsAppGetToOneRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaLicenseAgreements/{id}/app"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsApps != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[apps]", r.fieldsApps, "csv")
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

type ApiBetaLicenseAgreementsGetCollectionRequest struct {
	ctx context.Context
	ApiService *BetaLicenseAgreementsAPIService
	filterApp *[]string
	fieldsBetaLicenseAgreements *[]string
	limit *int32
	include *[]string
	fieldsApps *[]string
}

// filter by id(s) of related &#39;app&#39;
func (r ApiBetaLicenseAgreementsGetCollectionRequest) FilterApp(filterApp []string) ApiBetaLicenseAgreementsGetCollectionRequest {
	r.filterApp = &filterApp
	return r
}

// the fields to include for returned resources of type betaLicenseAgreements
func (r ApiBetaLicenseAgreementsGetCollectionRequest) FieldsBetaLicenseAgreements(fieldsBetaLicenseAgreements []string) ApiBetaLicenseAgreementsGetCollectionRequest {
	r.fieldsBetaLicenseAgreements = &fieldsBetaLicenseAgreements
	return r
}

// maximum resources per page
func (r ApiBetaLicenseAgreementsGetCollectionRequest) Limit(limit int32) ApiBetaLicenseAgreementsGetCollectionRequest {
	r.limit = &limit
	return r
}

// comma-separated list of relationships to include
func (r ApiBetaLicenseAgreementsGetCollectionRequest) Include(include []string) ApiBetaLicenseAgreementsGetCollectionRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type apps
func (r ApiBetaLicenseAgreementsGetCollectionRequest) FieldsApps(fieldsApps []string) ApiBetaLicenseAgreementsGetCollectionRequest {
	r.fieldsApps = &fieldsApps
	return r
}

func (r ApiBetaLicenseAgreementsGetCollectionRequest) Execute() (*BetaLicenseAgreementsResponse, *http.Response, error) {
	return r.ApiService.BetaLicenseAgreementsGetCollectionExecute(r)
}

/*
BetaLicenseAgreementsGetCollection Method for BetaLicenseAgreementsGetCollection

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiBetaLicenseAgreementsGetCollectionRequest
*/
func (a *BetaLicenseAgreementsAPIService) BetaLicenseAgreementsGetCollection(ctx context.Context) ApiBetaLicenseAgreementsGetCollectionRequest {
	return ApiBetaLicenseAgreementsGetCollectionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BetaLicenseAgreementsResponse
func (a *BetaLicenseAgreementsAPIService) BetaLicenseAgreementsGetCollectionExecute(r ApiBetaLicenseAgreementsGetCollectionRequest) (*BetaLicenseAgreementsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BetaLicenseAgreementsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaLicenseAgreementsAPIService.BetaLicenseAgreementsGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaLicenseAgreements"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterApp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[app]", r.filterApp, "csv")
	}
	if r.fieldsBetaLicenseAgreements != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[betaLicenseAgreements]", r.fieldsBetaLicenseAgreements, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.fieldsApps != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[apps]", r.fieldsApps, "csv")
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

type ApiBetaLicenseAgreementsGetInstanceRequest struct {
	ctx context.Context
	ApiService *BetaLicenseAgreementsAPIService
	id string
	fieldsBetaLicenseAgreements *[]string
	include *[]string
	fieldsApps *[]string
}

// the fields to include for returned resources of type betaLicenseAgreements
func (r ApiBetaLicenseAgreementsGetInstanceRequest) FieldsBetaLicenseAgreements(fieldsBetaLicenseAgreements []string) ApiBetaLicenseAgreementsGetInstanceRequest {
	r.fieldsBetaLicenseAgreements = &fieldsBetaLicenseAgreements
	return r
}

// comma-separated list of relationships to include
func (r ApiBetaLicenseAgreementsGetInstanceRequest) Include(include []string) ApiBetaLicenseAgreementsGetInstanceRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type apps
func (r ApiBetaLicenseAgreementsGetInstanceRequest) FieldsApps(fieldsApps []string) ApiBetaLicenseAgreementsGetInstanceRequest {
	r.fieldsApps = &fieldsApps
	return r
}

func (r ApiBetaLicenseAgreementsGetInstanceRequest) Execute() (*BetaLicenseAgreementResponse, *http.Response, error) {
	return r.ApiService.BetaLicenseAgreementsGetInstanceExecute(r)
}

/*
BetaLicenseAgreementsGetInstance Method for BetaLicenseAgreementsGetInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiBetaLicenseAgreementsGetInstanceRequest
*/
func (a *BetaLicenseAgreementsAPIService) BetaLicenseAgreementsGetInstance(ctx context.Context, id string) ApiBetaLicenseAgreementsGetInstanceRequest {
	return ApiBetaLicenseAgreementsGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return BetaLicenseAgreementResponse
func (a *BetaLicenseAgreementsAPIService) BetaLicenseAgreementsGetInstanceExecute(r ApiBetaLicenseAgreementsGetInstanceRequest) (*BetaLicenseAgreementResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BetaLicenseAgreementResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaLicenseAgreementsAPIService.BetaLicenseAgreementsGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaLicenseAgreements/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsBetaLicenseAgreements != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[betaLicenseAgreements]", r.fieldsBetaLicenseAgreements, "csv")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.fieldsApps != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[apps]", r.fieldsApps, "csv")
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

type ApiBetaLicenseAgreementsUpdateInstanceRequest struct {
	ctx context.Context
	ApiService *BetaLicenseAgreementsAPIService
	id string
	betaLicenseAgreementUpdateRequest *BetaLicenseAgreementUpdateRequest
}

// BetaLicenseAgreement representation
func (r ApiBetaLicenseAgreementsUpdateInstanceRequest) BetaLicenseAgreementUpdateRequest(betaLicenseAgreementUpdateRequest BetaLicenseAgreementUpdateRequest) ApiBetaLicenseAgreementsUpdateInstanceRequest {
	r.betaLicenseAgreementUpdateRequest = &betaLicenseAgreementUpdateRequest
	return r
}

func (r ApiBetaLicenseAgreementsUpdateInstanceRequest) Execute() (*BetaLicenseAgreementResponse, *http.Response, error) {
	return r.ApiService.BetaLicenseAgreementsUpdateInstanceExecute(r)
}

/*
BetaLicenseAgreementsUpdateInstance Method for BetaLicenseAgreementsUpdateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiBetaLicenseAgreementsUpdateInstanceRequest
*/
func (a *BetaLicenseAgreementsAPIService) BetaLicenseAgreementsUpdateInstance(ctx context.Context, id string) ApiBetaLicenseAgreementsUpdateInstanceRequest {
	return ApiBetaLicenseAgreementsUpdateInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return BetaLicenseAgreementResponse
func (a *BetaLicenseAgreementsAPIService) BetaLicenseAgreementsUpdateInstanceExecute(r ApiBetaLicenseAgreementsUpdateInstanceRequest) (*BetaLicenseAgreementResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BetaLicenseAgreementResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BetaLicenseAgreementsAPIService.BetaLicenseAgreementsUpdateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/betaLicenseAgreements/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.betaLicenseAgreementUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("betaLicenseAgreementUpdateRequest is required and must be specified")
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
	localVarPostBody = r.betaLicenseAgreementUpdateRequest
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
