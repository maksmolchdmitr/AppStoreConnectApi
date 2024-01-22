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


// GameCenterLeaderboardLocalizationsAPIService GameCenterLeaderboardLocalizationsAPI service
type GameCenterLeaderboardLocalizationsAPIService service

type ApiGameCenterLeaderboardLocalizationsCreateInstanceRequest struct {
	ctx context.Context
	ApiService *GameCenterLeaderboardLocalizationsAPIService
	gameCenterLeaderboardLocalizationCreateRequest *GameCenterLeaderboardLocalizationCreateRequest
}

// GameCenterLeaderboardLocalization representation
func (r ApiGameCenterLeaderboardLocalizationsCreateInstanceRequest) GameCenterLeaderboardLocalizationCreateRequest(gameCenterLeaderboardLocalizationCreateRequest GameCenterLeaderboardLocalizationCreateRequest) ApiGameCenterLeaderboardLocalizationsCreateInstanceRequest {
	r.gameCenterLeaderboardLocalizationCreateRequest = &gameCenterLeaderboardLocalizationCreateRequest
	return r
}

func (r ApiGameCenterLeaderboardLocalizationsCreateInstanceRequest) Execute() (*GameCenterLeaderboardLocalizationResponse, *http.Response, error) {
	return r.ApiService.GameCenterLeaderboardLocalizationsCreateInstanceExecute(r)
}

/*
GameCenterLeaderboardLocalizationsCreateInstance Method for GameCenterLeaderboardLocalizationsCreateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGameCenterLeaderboardLocalizationsCreateInstanceRequest
*/
func (a *GameCenterLeaderboardLocalizationsAPIService) GameCenterLeaderboardLocalizationsCreateInstance(ctx context.Context) ApiGameCenterLeaderboardLocalizationsCreateInstanceRequest {
	return ApiGameCenterLeaderboardLocalizationsCreateInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GameCenterLeaderboardLocalizationResponse
func (a *GameCenterLeaderboardLocalizationsAPIService) GameCenterLeaderboardLocalizationsCreateInstanceExecute(r ApiGameCenterLeaderboardLocalizationsCreateInstanceRequest) (*GameCenterLeaderboardLocalizationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GameCenterLeaderboardLocalizationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GameCenterLeaderboardLocalizationsAPIService.GameCenterLeaderboardLocalizationsCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/gameCenterLeaderboardLocalizations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.gameCenterLeaderboardLocalizationCreateRequest == nil {
		return localVarReturnValue, nil, reportError("gameCenterLeaderboardLocalizationCreateRequest is required and must be specified")
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
	localVarPostBody = r.gameCenterLeaderboardLocalizationCreateRequest
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

type ApiGameCenterLeaderboardLocalizationsDeleteInstanceRequest struct {
	ctx context.Context
	ApiService *GameCenterLeaderboardLocalizationsAPIService
	id string
}

func (r ApiGameCenterLeaderboardLocalizationsDeleteInstanceRequest) Execute() (*http.Response, error) {
	return r.ApiService.GameCenterLeaderboardLocalizationsDeleteInstanceExecute(r)
}

/*
GameCenterLeaderboardLocalizationsDeleteInstance Method for GameCenterLeaderboardLocalizationsDeleteInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiGameCenterLeaderboardLocalizationsDeleteInstanceRequest
*/
func (a *GameCenterLeaderboardLocalizationsAPIService) GameCenterLeaderboardLocalizationsDeleteInstance(ctx context.Context, id string) ApiGameCenterLeaderboardLocalizationsDeleteInstanceRequest {
	return ApiGameCenterLeaderboardLocalizationsDeleteInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *GameCenterLeaderboardLocalizationsAPIService) GameCenterLeaderboardLocalizationsDeleteInstanceExecute(r ApiGameCenterLeaderboardLocalizationsDeleteInstanceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GameCenterLeaderboardLocalizationsAPIService.GameCenterLeaderboardLocalizationsDeleteInstance")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/gameCenterLeaderboardLocalizations/{id}"
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

type ApiGameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelatedRequest struct {
	ctx context.Context
	ApiService *GameCenterLeaderboardLocalizationsAPIService
	id string
	fieldsGameCenterLeaderboardImages *[]string
	fieldsGameCenterLeaderboardLocalizations *[]string
	include *[]string
}

// the fields to include for returned resources of type gameCenterLeaderboardImages
func (r ApiGameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelatedRequest) FieldsGameCenterLeaderboardImages(fieldsGameCenterLeaderboardImages []string) ApiGameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelatedRequest {
	r.fieldsGameCenterLeaderboardImages = &fieldsGameCenterLeaderboardImages
	return r
}

// the fields to include for returned resources of type gameCenterLeaderboardLocalizations
func (r ApiGameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelatedRequest) FieldsGameCenterLeaderboardLocalizations(fieldsGameCenterLeaderboardLocalizations []string) ApiGameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelatedRequest {
	r.fieldsGameCenterLeaderboardLocalizations = &fieldsGameCenterLeaderboardLocalizations
	return r
}

// comma-separated list of relationships to include
func (r ApiGameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelatedRequest) Include(include []string) ApiGameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelatedRequest {
	r.include = &include
	return r
}

func (r ApiGameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelatedRequest) Execute() (*GameCenterLeaderboardImageResponse, *http.Response, error) {
	return r.ApiService.GameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelatedExecute(r)
}

/*
GameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelated Method for GameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelated

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiGameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelatedRequest
*/
func (a *GameCenterLeaderboardLocalizationsAPIService) GameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelated(ctx context.Context, id string) ApiGameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelatedRequest {
	return ApiGameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelatedRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GameCenterLeaderboardImageResponse
func (a *GameCenterLeaderboardLocalizationsAPIService) GameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelatedExecute(r ApiGameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelatedRequest) (*GameCenterLeaderboardImageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GameCenterLeaderboardImageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GameCenterLeaderboardLocalizationsAPIService.GameCenterLeaderboardLocalizationsGameCenterLeaderboardImageGetToOneRelated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/gameCenterLeaderboardLocalizations/{id}/gameCenterLeaderboardImage"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsGameCenterLeaderboardImages != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[gameCenterLeaderboardImages]", r.fieldsGameCenterLeaderboardImages, "csv")
	}
	if r.fieldsGameCenterLeaderboardLocalizations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[gameCenterLeaderboardLocalizations]", r.fieldsGameCenterLeaderboardLocalizations, "csv")
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

type ApiGameCenterLeaderboardLocalizationsGetInstanceRequest struct {
	ctx context.Context
	ApiService *GameCenterLeaderboardLocalizationsAPIService
	id string
	fieldsGameCenterLeaderboardLocalizations *[]string
	include *[]string
	fieldsGameCenterLeaderboardImages *[]string
}

// the fields to include for returned resources of type gameCenterLeaderboardLocalizations
func (r ApiGameCenterLeaderboardLocalizationsGetInstanceRequest) FieldsGameCenterLeaderboardLocalizations(fieldsGameCenterLeaderboardLocalizations []string) ApiGameCenterLeaderboardLocalizationsGetInstanceRequest {
	r.fieldsGameCenterLeaderboardLocalizations = &fieldsGameCenterLeaderboardLocalizations
	return r
}

// comma-separated list of relationships to include
func (r ApiGameCenterLeaderboardLocalizationsGetInstanceRequest) Include(include []string) ApiGameCenterLeaderboardLocalizationsGetInstanceRequest {
	r.include = &include
	return r
}

// the fields to include for returned resources of type gameCenterLeaderboardImages
func (r ApiGameCenterLeaderboardLocalizationsGetInstanceRequest) FieldsGameCenterLeaderboardImages(fieldsGameCenterLeaderboardImages []string) ApiGameCenterLeaderboardLocalizationsGetInstanceRequest {
	r.fieldsGameCenterLeaderboardImages = &fieldsGameCenterLeaderboardImages
	return r
}

func (r ApiGameCenterLeaderboardLocalizationsGetInstanceRequest) Execute() (*GameCenterLeaderboardLocalizationResponse, *http.Response, error) {
	return r.ApiService.GameCenterLeaderboardLocalizationsGetInstanceExecute(r)
}

/*
GameCenterLeaderboardLocalizationsGetInstance Method for GameCenterLeaderboardLocalizationsGetInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiGameCenterLeaderboardLocalizationsGetInstanceRequest
*/
func (a *GameCenterLeaderboardLocalizationsAPIService) GameCenterLeaderboardLocalizationsGetInstance(ctx context.Context, id string) ApiGameCenterLeaderboardLocalizationsGetInstanceRequest {
	return ApiGameCenterLeaderboardLocalizationsGetInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GameCenterLeaderboardLocalizationResponse
func (a *GameCenterLeaderboardLocalizationsAPIService) GameCenterLeaderboardLocalizationsGetInstanceExecute(r ApiGameCenterLeaderboardLocalizationsGetInstanceRequest) (*GameCenterLeaderboardLocalizationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GameCenterLeaderboardLocalizationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GameCenterLeaderboardLocalizationsAPIService.GameCenterLeaderboardLocalizationsGetInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/gameCenterLeaderboardLocalizations/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fieldsGameCenterLeaderboardLocalizations != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[gameCenterLeaderboardLocalizations]", r.fieldsGameCenterLeaderboardLocalizations, "csv")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "csv")
	}
	if r.fieldsGameCenterLeaderboardImages != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields[gameCenterLeaderboardImages]", r.fieldsGameCenterLeaderboardImages, "csv")
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

type ApiGameCenterLeaderboardLocalizationsUpdateInstanceRequest struct {
	ctx context.Context
	ApiService *GameCenterLeaderboardLocalizationsAPIService
	id string
	gameCenterLeaderboardLocalizationUpdateRequest *GameCenterLeaderboardLocalizationUpdateRequest
}

// GameCenterLeaderboardLocalization representation
func (r ApiGameCenterLeaderboardLocalizationsUpdateInstanceRequest) GameCenterLeaderboardLocalizationUpdateRequest(gameCenterLeaderboardLocalizationUpdateRequest GameCenterLeaderboardLocalizationUpdateRequest) ApiGameCenterLeaderboardLocalizationsUpdateInstanceRequest {
	r.gameCenterLeaderboardLocalizationUpdateRequest = &gameCenterLeaderboardLocalizationUpdateRequest
	return r
}

func (r ApiGameCenterLeaderboardLocalizationsUpdateInstanceRequest) Execute() (*GameCenterLeaderboardLocalizationResponse, *http.Response, error) {
	return r.ApiService.GameCenterLeaderboardLocalizationsUpdateInstanceExecute(r)
}

/*
GameCenterLeaderboardLocalizationsUpdateInstance Method for GameCenterLeaderboardLocalizationsUpdateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id the id of the requested resource
 @return ApiGameCenterLeaderboardLocalizationsUpdateInstanceRequest
*/
func (a *GameCenterLeaderboardLocalizationsAPIService) GameCenterLeaderboardLocalizationsUpdateInstance(ctx context.Context, id string) ApiGameCenterLeaderboardLocalizationsUpdateInstanceRequest {
	return ApiGameCenterLeaderboardLocalizationsUpdateInstanceRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GameCenterLeaderboardLocalizationResponse
func (a *GameCenterLeaderboardLocalizationsAPIService) GameCenterLeaderboardLocalizationsUpdateInstanceExecute(r ApiGameCenterLeaderboardLocalizationsUpdateInstanceRequest) (*GameCenterLeaderboardLocalizationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GameCenterLeaderboardLocalizationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GameCenterLeaderboardLocalizationsAPIService.GameCenterLeaderboardLocalizationsUpdateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/gameCenterLeaderboardLocalizations/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.gameCenterLeaderboardLocalizationUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("gameCenterLeaderboardLocalizationUpdateRequest is required and must be specified")
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
	localVarPostBody = r.gameCenterLeaderboardLocalizationUpdateRequest
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
