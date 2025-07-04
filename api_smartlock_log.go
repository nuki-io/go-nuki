/*
Nuki API

The Nuki Web Api

API version: 3.10.1
Contact: contact@nuki.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// SmartlockLogAPIService SmartlockLogAPI service
type SmartlockLogAPIService service

type ApiGetSmartlockLogsRequest struct {
	ctx context.Context
	ApiService *SmartlockLogAPIService
	smartlockId int32
	authId *string
	accountUserId *int32
	fromDate *string
	toDate *string
	action *int32
	id *string
	limit *int32
}

// Filter for auths
func (r ApiGetSmartlockLogsRequest) AuthId(authId string) ApiGetSmartlockLogsRequest {
	r.authId = &authId
	return r
}

// Filter for account users
func (r ApiGetSmartlockLogsRequest) AccountUserId(accountUserId int32) ApiGetSmartlockLogsRequest {
	r.accountUserId = &accountUserId
	return r
}

// Filter for date (RFC3339)
func (r ApiGetSmartlockLogsRequest) FromDate(fromDate string) ApiGetSmartlockLogsRequest {
	r.fromDate = &fromDate
	return r
}

// Filter for date (RFC3339)
func (r ApiGetSmartlockLogsRequest) ToDate(toDate string) ApiGetSmartlockLogsRequest {
	r.toDate = &toDate
	return r
}

// Filter for action
func (r ApiGetSmartlockLogsRequest) Action(action int32) ApiGetSmartlockLogsRequest {
	r.action = &action
	return r
}

// Filter for older logs
func (r ApiGetSmartlockLogsRequest) Id(id string) ApiGetSmartlockLogsRequest {
	r.id = &id
	return r
}

// Amount of logs (max: 50)
func (r ApiGetSmartlockLogsRequest) Limit(limit int32) ApiGetSmartlockLogsRequest {
	r.limit = &limit
	return r
}

func (r ApiGetSmartlockLogsRequest) Execute() ([]SmartlockLog, *http.Response, error) {
	return r.ApiService.GetSmartlockLogsExecute(r)
}

/*
GetSmartlockLogs Get a list of smartlock logs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param smartlockId The smartlock id
 @return ApiGetSmartlockLogsRequest
*/
func (a *SmartlockLogAPIService) GetSmartlockLogs(ctx context.Context, smartlockId int32) ApiGetSmartlockLogsRequest {
	return ApiGetSmartlockLogsRequest{
		ApiService: a,
		ctx: ctx,
		smartlockId: smartlockId,
	}
}

// Execute executes the request
//  @return []SmartlockLog
func (a *SmartlockLogAPIService) GetSmartlockLogsExecute(r ApiGetSmartlockLogsRequest) ([]SmartlockLog, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SmartlockLog
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmartlockLogAPIService.GetSmartlockLogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/smartlock/{smartlockId}/log"
	localVarPath = strings.Replace(localVarPath, "{"+"smartlockId"+"}", url.PathEscape(parameterValueToString(r.smartlockId, "smartlockId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.authId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "authId", r.authId, "", "")
	}
	if r.accountUserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "accountUserId", r.accountUserId, "", "")
	}
	if r.fromDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fromDate", r.fromDate, "", "")
	}
	if r.toDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "toDate", r.toDate, "", "")
	}
	if r.action != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "action", r.action, "", "")
	}
	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "", "")
	} else {
		var defaultValue int32 = 20
		r.limit = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

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

type ApiGetSmartlocksLogsRequest struct {
	ctx context.Context
	ApiService *SmartlockLogAPIService
	accountUserId *int32
	fromDate *string
	toDate *string
	action *int32
	id *string
	limit *int32
}

// Filter for account users
func (r ApiGetSmartlocksLogsRequest) AccountUserId(accountUserId int32) ApiGetSmartlocksLogsRequest {
	r.accountUserId = &accountUserId
	return r
}

// Filter for date (RFC3339)
func (r ApiGetSmartlocksLogsRequest) FromDate(fromDate string) ApiGetSmartlocksLogsRequest {
	r.fromDate = &fromDate
	return r
}

// Filter for date (RFC3339)
func (r ApiGetSmartlocksLogsRequest) ToDate(toDate string) ApiGetSmartlocksLogsRequest {
	r.toDate = &toDate
	return r
}

// Filter for action
func (r ApiGetSmartlocksLogsRequest) Action(action int32) ApiGetSmartlocksLogsRequest {
	r.action = &action
	return r
}

// Filter for older logs
func (r ApiGetSmartlocksLogsRequest) Id(id string) ApiGetSmartlocksLogsRequest {
	r.id = &id
	return r
}

// Amount of logs (max: 50)
func (r ApiGetSmartlocksLogsRequest) Limit(limit int32) ApiGetSmartlocksLogsRequest {
	r.limit = &limit
	return r
}

func (r ApiGetSmartlocksLogsRequest) Execute() ([]SmartlockLog, *http.Response, error) {
	return r.ApiService.GetSmartlocksLogsExecute(r)
}

/*
GetSmartlocksLogs Get a list of smartlock logs for all of your smartlocks

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSmartlocksLogsRequest
*/
func (a *SmartlockLogAPIService) GetSmartlocksLogs(ctx context.Context) ApiGetSmartlocksLogsRequest {
	return ApiGetSmartlocksLogsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SmartlockLog
func (a *SmartlockLogAPIService) GetSmartlocksLogsExecute(r ApiGetSmartlocksLogsRequest) ([]SmartlockLog, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SmartlockLog
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SmartlockLogAPIService.GetSmartlocksLogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/smartlock/log"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.accountUserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "accountUserId", r.accountUserId, "", "")
	}
	if r.fromDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fromDate", r.fromDate, "", "")
	}
	if r.toDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "toDate", r.toDate, "", "")
	}
	if r.action != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "action", r.action, "", "")
	}
	if r.id != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id", r.id, "", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "", "")
	} else {
		var defaultValue int32 = 20
		r.limit = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

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
