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

// SubscriptionPromotionalOffersAPIService SubscriptionPromotionalOffersAPI service
type SubscriptionPromotionalOffersAPIService service

type ApiSubscriptionPromotionalOffersCreateInstanceRequest struct {
	ctx                                       context.Context
	ApiService                                *SubscriptionPromotionalOffersAPIService
	subscriptionPromotionalOfferCreateRequest *SubscriptionPromotionalOfferCreateRequest
}

// SubscriptionPromotionalOffer representation
func (r ApiSubscriptionPromotionalOffersCreateInstanceRequest) SubscriptionPromotionalOfferCreateRequest(subscriptionPromotionalOfferCreateRequest SubscriptionPromotionalOfferCreateRequest) ApiSubscriptionPromotionalOffersCreateInstanceRequest {
	r.subscriptionPromotionalOfferCreateRequest = &subscriptionPromotionalOfferCreateRequest
	return r
}

func (r ApiSubscriptionPromotionalOffersCreateInstanceRequest) Execute() (*SubscriptionPromotionalOfferResponse, *http.Response, error) {
	return r.ApiService.SubscriptionPromotionalOffersCreateInstanceExecute(r)
}

/*
SubscriptionPromotionalOffersCreateInstance Method for SubscriptionPromotionalOffersCreateInstance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSubscriptionPromotionalOffersCreateInstanceRequest
*/
func (a *SubscriptionPromotionalOffersAPIService) SubscriptionPromotionalOffersCreateInstance(ctx context.Context) ApiSubscriptionPromotionalOffersCreateInstanceRequest {
	return ApiSubscriptionPromotionalOffersCreateInstanceRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SubscriptionPromotionalOfferResponse
func (a *SubscriptionPromotionalOffersAPIService) SubscriptionPromotionalOffersCreateInstanceExecute(r ApiSubscriptionPromotionalOffersCreateInstanceRequest) (*SubscriptionPromotionalOfferResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SubscriptionPromotionalOfferResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionPromotionalOffersAPIService.SubscriptionPromotionalOffersCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/subscriptionPromotionalOffers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.subscriptionPromotionalOfferCreateRequest == nil {
		return localVarReturnValue, nil, reportError("subscriptionPromotionalOfferCreateRequest is required and must be specified")
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
	localVarPostBody = r.subscriptionPromotionalOfferCreateRequest
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

type ApiSubscriptionPromotionalOffersDeleteInstanceRequest struct {
	ctx        context.Context
	ApiService *SubscriptionPromotionalOffersAPIService
	id         string
}

func (r ApiSubscriptionPromotionalOffersDeleteInstanceRequest) Execute() (*http.Response, error) {
	return r.ApiService.SubscriptionPromotionalOffersDeleteInstanceExecute(r)
}

/*
SubscriptionPromotionalOffersDeleteInstance Method for SubscriptionPromotionalOffersDeleteInstance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiSubscriptionPromotionalOffersDeleteInstanceRequest
*/
func (a *SubscriptionPromotionalOffersAPIService) SubscriptionPromotionalOffersDeleteInstance(ctx context.Context, id string) ApiSubscriptionPromotionalOffersDeleteInstanceRequest {
	return ApiSubscriptionPromotionalOffersDeleteInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *SubscriptionPromotionalOffersAPIService) SubscriptionPromotionalOffersDeleteInstanceExecute(r ApiSubscriptionPromotionalOffersDeleteInstanceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionPromotionalOffersAPIService.SubscriptionPromotionalOffersDeleteInstance")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/subscriptionPromotionalOffers/{id}"
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

type ApiSubscriptionPromotionalOffersGetInstanceRequest struct {
	ctx                                      context.Context
	ApiService                               *SubscriptionPromotionalOffersAPIService
	id                                       string
	fieldsSubscriptionPromotionalOffers      *[]string
	include                                  *[]string
	fieldsSubscriptionPromotionalOfferPrices *[]string
	limitPrices                              *int32
}

// the fields to include for returned resources of type subscriptionPromotionalOffers
func (r ApiSubscriptionPromotionalOffersGetInstanceRequest) FieldsSubscriptionPromotionalOffers(fieldsSubscriptionPromotionalOffers []string) ApiSubscriptionPromotionalOffersGetInstanceRequest {
	r.fieldsSubscriptionPromotionalOffers = &fieldsSubscriptionPromotionalOffers
	return r
}

// comma-separated list of relationships to include
func (r ApiSubscriptionPromotionalOffersGetInstanceRequest) Include(include []string) ApiSubscriptionPromotionalOffersGetInstanceRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type subscriptionPromotionalOfferPrices
func (r ApiSubscriptionPromotionalOffersGetInstanceRequest) FieldsSubscriptionPromotionalOfferPrices(fieldsSubscriptionPromotionalOfferPrices []string) ApiSubscriptionPromotionalOffersGetInstanceRequest {
	r.fieldsSubscriptionPromotionalOfferPrices = &fieldsSubscriptionPromotionalOfferPrices
	return r
}

// maximum number of related prices returned (when they are included)
func (r ApiSubscriptionPromotionalOffersGetInstanceRequest) LimitPrices(limitPrices int32) ApiSubscriptionPromotionalOffersGetInstanceRequest {
	r.limitPrices = &limitPrices
	return r
}

func (r ApiSubscriptionPromotionalOffersGetInstanceRequest) Execute() (*SubscriptionPromotionalOfferResponse, *http.Response, error) {
	return r.ApiService.SubscriptionPromotionalOffersGetInstanceExecute(r)
}

/*
SubscriptionPromotionalOffersGetInstance Method for SubscriptionPromotionalOffersGetInstance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiSubscriptionPromotionalOffersGetInstanceRequest
*/
func (a *SubscriptionPromotionalOffersAPIService) SubscriptionPromotionalOffersGetInstance(ctx context.Context, id string) ApiSubscriptionPromotionalOffersGetInstanceRequest {
	return ApiSubscriptionPromotionalOffersGetInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return SubscriptionPromotionalOfferResponse
func (a *SubscriptionPromotionalOffersAPIService) SubscriptionPromotionalOffersGetInstanceExecute(r ApiSubscriptionPromotionalOffersGetInstanceRequest) (*SubscriptionPromotionalOfferResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SubscriptionPromotionalOfferResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionPromotionalOffersAPIService.SubscriptionPromotionalOffersGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/subscriptionPromotionalOffers/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsSubscriptionPromotionalOffers != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[subscriptionPromotionalOffers]", r.fieldsSubscriptionPromotionalOffers, "csv")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.fieldsSubscriptionPromotionalOfferPrices != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[subscriptionPromotionalOfferPrices]", r.fieldsSubscriptionPromotionalOfferPrices, "csv")
	}
	if r.limitPrices != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit[prices]", r.limitPrices, "")
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

type ApiSubscriptionPromotionalOffersPricesGetToManyRelatedRequest struct {
	ctx                                      context.Context
	ApiService                               *SubscriptionPromotionalOffersAPIService
	id                                       string
	filterTerritory                          *[]string
	fieldsSubscriptionPricePoints            *[]string
	fieldsTerritories                        *[]string
	fieldsSubscriptionPromotionalOfferPrices *[]string
	limit                                    *int32
	include                                  *[]string
}

// filter by id(s) of related &#39;territory&#39;
func (r ApiSubscriptionPromotionalOffersPricesGetToManyRelatedRequest) FilterTerritory(filterTerritory []string) ApiSubscriptionPromotionalOffersPricesGetToManyRelatedRequest {
	r.filterTerritory = &filterTerritory
	return r
}

// the fields to include for returned resources of type subscriptionPricePoints
func (r ApiSubscriptionPromotionalOffersPricesGetToManyRelatedRequest) FieldsSubscriptionPricePoints(fieldsSubscriptionPricePoints []string) ApiSubscriptionPromotionalOffersPricesGetToManyRelatedRequest {
	r.fieldsSubscriptionPricePoints = &fieldsSubscriptionPricePoints
	return r
}

// the fields to include for returned resources of type territories
func (r ApiSubscriptionPromotionalOffersPricesGetToManyRelatedRequest) FieldsTerritories(fieldsTerritories []string) ApiSubscriptionPromotionalOffersPricesGetToManyRelatedRequest {
	r.fieldsTerritories = &fieldsTerritories
	return r
}

// the fields to include for returned resources of type subscriptionPromotionalOfferPrices
func (r ApiSubscriptionPromotionalOffersPricesGetToManyRelatedRequest) FieldsSubscriptionPromotionalOfferPrices(fieldsSubscriptionPromotionalOfferPrices []string) ApiSubscriptionPromotionalOffersPricesGetToManyRelatedRequest {
	r.fieldsSubscriptionPromotionalOfferPrices = &fieldsSubscriptionPromotionalOfferPrices
	return r
}

// maximum resources per page
func (r ApiSubscriptionPromotionalOffersPricesGetToManyRelatedRequest) Limit(limit int32) ApiSubscriptionPromotionalOffersPricesGetToManyRelatedRequest {
	r.limit = &limit
	return r
}

// comma-separated list of relationships to include
func (r ApiSubscriptionPromotionalOffersPricesGetToManyRelatedRequest) Include(include []string) ApiSubscriptionPromotionalOffersPricesGetToManyRelatedRequest {
	r.include = &include
	return r
}

func (r ApiSubscriptionPromotionalOffersPricesGetToManyRelatedRequest) Execute() (*SubscriptionPromotionalOfferPricesResponse, *http.Response, error) {
	return r.ApiService.SubscriptionPromotionalOffersPricesGetToManyRelatedExecute(r)
}

/*
SubscriptionPromotionalOffersPricesGetToManyRelated Method for SubscriptionPromotionalOffersPricesGetToManyRelated

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiSubscriptionPromotionalOffersPricesGetToManyRelatedRequest
*/
func (a *SubscriptionPromotionalOffersAPIService) SubscriptionPromotionalOffersPricesGetToManyRelated(ctx context.Context, id string) ApiSubscriptionPromotionalOffersPricesGetToManyRelatedRequest {
	return ApiSubscriptionPromotionalOffersPricesGetToManyRelatedRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return SubscriptionPromotionalOfferPricesResponse
func (a *SubscriptionPromotionalOffersAPIService) SubscriptionPromotionalOffersPricesGetToManyRelatedExecute(r ApiSubscriptionPromotionalOffersPricesGetToManyRelatedRequest) (*SubscriptionPromotionalOfferPricesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SubscriptionPromotionalOfferPricesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionPromotionalOffersAPIService.SubscriptionPromotionalOffersPricesGetToManyRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/subscriptionPromotionalOffers/{id}/prices"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filterTerritory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter[territory]", r.filterTerritory, "csv")
	}
	if r.fieldsSubscriptionPricePoints != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[subscriptionPricePoints]", r.fieldsSubscriptionPricePoints, "csv")
	}
	if r.fieldsTerritories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[territories]", r.fieldsTerritories, "csv")
	}
	if r.fieldsSubscriptionPromotionalOfferPrices != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[subscriptionPromotionalOfferPrices]", r.fieldsSubscriptionPromotionalOfferPrices, "csv")
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

type ApiSubscriptionPromotionalOffersUpdateInstanceRequest struct {
	ctx                                       context.Context
	ApiService                                *SubscriptionPromotionalOffersAPIService
	id                                        string
	subscriptionPromotionalOfferUpdateRequest *SubscriptionPromotionalOfferUpdateRequest
}

// SubscriptionPromotionalOffer representation
func (r ApiSubscriptionPromotionalOffersUpdateInstanceRequest) SubscriptionPromotionalOfferUpdateRequest(subscriptionPromotionalOfferUpdateRequest SubscriptionPromotionalOfferUpdateRequest) ApiSubscriptionPromotionalOffersUpdateInstanceRequest {
	r.subscriptionPromotionalOfferUpdateRequest = &subscriptionPromotionalOfferUpdateRequest
	return r
}

func (r ApiSubscriptionPromotionalOffersUpdateInstanceRequest) Execute() (*SubscriptionPromotionalOfferResponse, *http.Response, error) {
	return r.ApiService.SubscriptionPromotionalOffersUpdateInstanceExecute(r)
}

/*
SubscriptionPromotionalOffersUpdateInstance Method for SubscriptionPromotionalOffersUpdateInstance

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id the id of the requested resource
	@return ApiSubscriptionPromotionalOffersUpdateInstanceRequest
*/
func (a *SubscriptionPromotionalOffersAPIService) SubscriptionPromotionalOffersUpdateInstance(ctx context.Context, id string) ApiSubscriptionPromotionalOffersUpdateInstanceRequest {
	return ApiSubscriptionPromotionalOffersUpdateInstanceRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return SubscriptionPromotionalOfferResponse
func (a *SubscriptionPromotionalOffersAPIService) SubscriptionPromotionalOffersUpdateInstanceExecute(r ApiSubscriptionPromotionalOffersUpdateInstanceRequest) (*SubscriptionPromotionalOfferResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SubscriptionPromotionalOfferResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionPromotionalOffersAPIService.SubscriptionPromotionalOffersUpdateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/subscriptionPromotionalOffers/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.subscriptionPromotionalOfferUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("subscriptionPromotionalOfferUpdateRequest is required and must be specified")
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
	localVarPostBody = r.subscriptionPromotionalOfferUpdateRequest
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
