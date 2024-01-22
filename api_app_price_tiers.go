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

// AppPriceTiersAPIService AppPriceTiersAPI service
type AppPriceTiersAPIService service

type ApiAppPriceTiersGetCollectionRequest struct {
	ctx                  context.Context
	ApiService           *AppPriceTiersAPIService
	filterId             *[]string
	fieldsAppPriceTiers  *[]string
	limit                *int32
	include              *[]string
	fieldsAppPricePoints *[]string
	limitPricePoints     *int32
}

// filter by id(s)
func (r ApiAppPriceTiersGetCollectionRequest) FilterId(filterId []string) ApiAppPriceTiersGetCollectionRequest {
	r.filterId = &filterId
	return r
}

// the fields to include for returned resources of type appPriceTiers
func (r ApiAppPriceTiersGetCollectionRequest) FieldsAppPriceTiers(fieldsAppPriceTiers []string) ApiAppPriceTiersGetCollectionRequest {
	r.fieldsAppPriceTiers = &fieldsAppPriceTiers
	return r
}

// maximum resources per page
func (r ApiAppPriceTiersGetCollectionRequest) Limit(limit int32) ApiAppPriceTiersGetCollectionRequest {
	r.limit = &limit
	return r
}

// comma-separated list of relationships to include
func (r ApiAppPriceTiersGetCollectionRequest) Include(include []string) ApiAppPriceTiersGetCollectionRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type appPricePoints
// Deprecated
func (r ApiAppPriceTiersGetCollectionRequest) FieldsAppPricePoints(fieldsAppPricePoints []string) ApiAppPriceTiersGetCollectionRequest {
	r.fieldsAppPricePoints = &fieldsAppPricePoints
	return r
}

// maximum number of related pricePoints returned (when they are included)
// Deprecated
func (r ApiAppPriceTiersGetCollectionRequest) LimitPricePoints(limitPricePoints int32) ApiAppPriceTiersGetCollectionRequest {
	r.limitPricePoints = &limitPricePoints
	return r
}

func (r ApiAppPriceTiersGetCollectionRequest) Execute() (*AppPriceTiersResponse, *http.Response, error) {
	return r.ApiService.AppPriceTiersGetCollectionExecute(r)
}

/*
AppPriceTiersGetCollection Method for AppPriceTiersGetCollection

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAppPriceTiersGetCollectionRequest

Deprecated
*/
func (a *AppPriceTiersAPIService) AppPriceTiersGetCollection(ctx context.Context) ApiAppPriceTiersGetCollectionRequest {
	return ApiAppPriceTiersGetCollectionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return AppPriceTiersResponse
//
// Deprecated
func (a *AppPriceTiersAPIService) AppPriceTiersGetCollectionExecute(r ApiAppPriceTiersGetCollectionRequest) (*AppPriceTiersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AppPriceTiersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppPriceTiersAPIService.AppPriceTiersGetCollection")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appPriceTiers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[id]", r.filterId, "csv")
	}
	if r.fieldsAppPriceTiers != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appPriceTiers]", r.fieldsAppPriceTiers, "csv")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.fieldsAppPricePoints != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appPricePoints]", r.fieldsAppPricePoints, "csv")
	}
	if r.limitPricePoints != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[pricePoints]", r.limitPricePoints, "")
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

type ApiAppPriceTiersGetInstanceRequest struct {
	ctx                  context.Context
	ApiService           *AppPriceTiersAPIService
	id                   string
	fieldsAppPriceTiers  *[]string
	include              *[]string
	fieldsAppPricePoints *[]string
	limitPricePoints     *int32
}

// the fields to include for returned resources of type appPriceTiers
func (r ApiAppPriceTiersGetInstanceRequest) FieldsAppPriceTiers(fieldsAppPriceTiers []string) ApiAppPriceTiersGetInstanceRequest {
	r.fieldsAppPriceTiers = &fieldsAppPriceTiers
	return r
}

// comma-separated list of relationships to include
func (r ApiAppPriceTiersGetInstanceRequest) Include(include []string) ApiAppPriceTiersGetInstanceRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type appPricePoints
// Deprecated
func (r ApiAppPriceTiersGetInstanceRequest) FieldsAppPricePoints(fieldsAppPricePoints []string) ApiAppPriceTiersGetInstanceRequest {
	r.fieldsAppPricePoints = &fieldsAppPricePoints
	return r
}

// maximum number of related pricePoints returned (when they are included)
// Deprecated
func (r ApiAppPriceTiersGetInstanceRequest) LimitPricePoints(limitPricePoints int32) ApiAppPriceTiersGetInstanceRequest {
	r.limitPricePoints = &limitPricePoints
	return r
}

func (r ApiAppPriceTiersGetInstanceRequest) Execute() (*AppPriceTierResponse, *http.Response, error) {
	return r.ApiService.AppPriceTiersGetInstanceExecute(r)
}

/*
AppPriceTiersGetInstance Method for AppPriceTiersGetInstance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiAppPriceTiersGetInstanceRequest

Deprecated
*/
func (a *AppPriceTiersAPIService) AppPriceTiersGetInstance(ctx context.Context, id string) ApiAppPriceTiersGetInstanceRequest {
	return ApiAppPriceTiersGetInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AppPriceTierResponse
//
// Deprecated
func (a *AppPriceTiersAPIService) AppPriceTiersGetInstanceExecute(r ApiAppPriceTiersGetInstanceRequest) (*AppPriceTierResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AppPriceTierResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppPriceTiersAPIService.AppPriceTiersGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appPriceTiers/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsAppPriceTiers != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appPriceTiers]", r.fieldsAppPriceTiers, "csv")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.fieldsAppPricePoints != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appPricePoints]", r.fieldsAppPricePoints, "csv")
	}
	if r.limitPricePoints != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[pricePoints]", r.limitPricePoints, "")
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

type ApiAppPriceTiersPricePointsGetToManyRelatedRequest struct {
	ctx                  context.Context
	ApiService           *AppPriceTiersAPIService
	id                   string
	filterTerritory      *[]string
	fieldsAppPriceTiers  *[]string
	fieldsAppPricePoints *[]string
	fieldsTerritories    *[]string
	limit                *int32
	include              *[]string
}

// filter by id(s) of related &#39;territory&#39;
func (r ApiAppPriceTiersPricePointsGetToManyRelatedRequest) FilterTerritory(filterTerritory []string) ApiAppPriceTiersPricePointsGetToManyRelatedRequest {
	r.filterTerritory = &filterTerritory
	return r
}

// the fields to include for returned resources of type appPriceTiers
func (r ApiAppPriceTiersPricePointsGetToManyRelatedRequest) FieldsAppPriceTiers(fieldsAppPriceTiers []string) ApiAppPriceTiersPricePointsGetToManyRelatedRequest {
	r.fieldsAppPriceTiers = &fieldsAppPriceTiers
	return r
}

// the fields to include for returned resources of type appPricePoints
func (r ApiAppPriceTiersPricePointsGetToManyRelatedRequest) FieldsAppPricePoints(fieldsAppPricePoints []string) ApiAppPriceTiersPricePointsGetToManyRelatedRequest {
	r.fieldsAppPricePoints = &fieldsAppPricePoints
	return r
}

// the fields to include for returned resources of type territories
func (r ApiAppPriceTiersPricePointsGetToManyRelatedRequest) FieldsTerritories(fieldsTerritories []string) ApiAppPriceTiersPricePointsGetToManyRelatedRequest {
	r.fieldsTerritories = &fieldsTerritories
	return r
}

// maximum resources per page
func (r ApiAppPriceTiersPricePointsGetToManyRelatedRequest) Limit(limit int32) ApiAppPriceTiersPricePointsGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

// comma-separated list of relationships to include
func (r ApiAppPriceTiersPricePointsGetToManyRelatedRequest) Include(include []string) ApiAppPriceTiersPricePointsGetToManyRelatedRequest {
	r.include = &include
	return r
}

func (r ApiAppPriceTiersPricePointsGetToManyRelatedRequest) Execute() (*AppPricePointsResponse, *http.Response, error) {
	return r.ApiService.AppPriceTiersPricePointsGetToManyRelatedExecute(r)
}

/*
AppPriceTiersPricePointsGetToManyRelated Method for AppPriceTiersPricePointsGetToManyRelated

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiAppPriceTiersPricePointsGetToManyRelatedRequest

Deprecated
*/
func (a *AppPriceTiersAPIService) AppPriceTiersPricePointsGetToManyRelated(ctx context.Context, id string) ApiAppPriceTiersPricePointsGetToManyRelatedRequest {
	return ApiAppPriceTiersPricePointsGetToManyRelatedRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return AppPricePointsResponse
//
// Deprecated
func (a *AppPriceTiersAPIService) AppPriceTiersPricePointsGetToManyRelatedExecute(r ApiAppPriceTiersPricePointsGetToManyRelatedRequest) (*AppPricePointsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AppPricePointsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppPriceTiersAPIService.AppPriceTiersPricePointsGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/appPriceTiers/{id}/pricePoints"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterTerritory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[territory]", r.filterTerritory, "csv")
	}
	if r.fieldsAppPriceTiers != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appPriceTiers]", r.fieldsAppPriceTiers, "csv")
	}
	if r.fieldsAppPricePoints != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[appPricePoints]", r.fieldsAppPricePoints, "csv")
	}
	if r.fieldsTerritories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[territories]", r.fieldsTerritories, "csv")
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
