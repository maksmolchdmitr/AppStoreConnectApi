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
)


// InAppPurchaseSubmissionsAPIService InAppPurchaseSubmissionsAPI service
type InAppPurchaseSubmissionsAPIService service

type ApiInAppPurchaseSubmissionsCreateInstanceRequest struct {
	ctx context.Context
	ApiService *InAppPurchaseSubmissionsAPIService
	inAppPurchaseSubmissionCreateRequest *InAppPurchaseSubmissionCreateRequest
}

// InAppPurchaseSubmission representation
func (r ApiInAppPurchaseSubmissionsCreateInstanceRequest) InAppPurchaseSubmissionCreateRequest(inAppPurchaseSubmissionCreateRequest InAppPurchaseSubmissionCreateRequest) ApiInAppPurchaseSubmissionsCreateInstanceRequest {
	r.inAppPurchaseSubmissionCreateRequest = &inAppPurchaseSubmissionCreateRequest
	return r
}

func (r ApiInAppPurchaseSubmissionsCreateInstanceRequest) Execute() (*InAppPurchaseSubmissionResponse, *http.Response, error) {
	return r.ApiService.InAppPurchaseSubmissionsCreateInstanceExecute(r)
}

/*
InAppPurchaseSubmissionsCreateInstance Method for InAppPurchaseSubmissionsCreateInstance

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiInAppPurchaseSubmissionsCreateInstanceRequest
*/
func (a *InAppPurchaseSubmissionsAPIService) InAppPurchaseSubmissionsCreateInstance(ctx context.Context) ApiInAppPurchaseSubmissionsCreateInstanceRequest {
	return ApiInAppPurchaseSubmissionsCreateInstanceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InAppPurchaseSubmissionResponse
func (a *InAppPurchaseSubmissionsAPIService) InAppPurchaseSubmissionsCreateInstanceExecute(r ApiInAppPurchaseSubmissionsCreateInstanceRequest) (*InAppPurchaseSubmissionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InAppPurchaseSubmissionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InAppPurchaseSubmissionsAPIService.InAppPurchaseSubmissionsCreateInstance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/inAppPurchaseSubmissions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.inAppPurchaseSubmissionCreateRequest == nil {
		return localVarReturnValue, nil, reportError("inAppPurchaseSubmissionCreateRequest is required and must be specified")
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
	localVarPostBody = r.inAppPurchaseSubmissionCreateRequest
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