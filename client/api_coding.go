/*
ai-proto

ai proto for coder

API version: 0.0.2
Contact: panleiming@linksaas.pro
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


// CodingApiService CodingApi service
type CodingApiService service

type ApiApiCodingCompleteLangPostRequest struct {
	ctx context.Context
	ApiService *CodingApiService
	xAuthToken *string
	lang string
	apiCodingCompleteLangPostRequest *ApiCodingCompleteLangPostRequest
}

func (r ApiApiCodingCompleteLangPostRequest) XAuthToken(xAuthToken string) ApiApiCodingCompleteLangPostRequest {
	r.xAuthToken = &xAuthToken
	return r
}

func (r ApiApiCodingCompleteLangPostRequest) ApiCodingCompleteLangPostRequest(apiCodingCompleteLangPostRequest ApiCodingCompleteLangPostRequest) ApiApiCodingCompleteLangPostRequest {
	r.apiCodingCompleteLangPostRequest = &apiCodingCompleteLangPostRequest
	return r
}

func (r ApiApiCodingCompleteLangPostRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.ApiCodingCompleteLangPostExecute(r)
}

/*
ApiCodingCompleteLangPost 根据上下文补全代码

根据上下文补全代码

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param lang
 @return ApiApiCodingCompleteLangPostRequest
*/
func (a *CodingApiService) ApiCodingCompleteLangPost(ctx context.Context, lang string) ApiApiCodingCompleteLangPostRequest {
	return ApiApiCodingCompleteLangPostRequest{
		ApiService: a,
		ctx: ctx,
		lang: lang,
	}
}

// Execute executes the request
//  @return []string
func (a *CodingApiService) ApiCodingCompleteLangPostExecute(r ApiApiCodingCompleteLangPostRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CodingApiService.ApiCodingCompleteLangPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/coding/complete/{lang}"
	localVarPath = strings.Replace(localVarPath, "{"+"lang"+"}", url.PathEscape(parameterValueToString(r.lang, "lang")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAuthToken == nil {
		return localVarReturnValue, nil, reportError("xAuthToken is required and must be specified")
	}
	if r.apiCodingCompleteLangPostRequest == nil {
		return localVarReturnValue, nil, reportError("apiCodingCompleteLangPostRequest is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-AuthToken", r.xAuthToken, "")
	// body params
	localVarPostBody = r.apiCodingCompleteLangPostRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrInfo
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrInfo
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

type ApiApiCodingConvertLangPostRequest struct {
	ctx context.Context
	ApiService *CodingApiService
	xAuthToken *string
	lang string
	apiCodingConvertLangPostRequest *ApiCodingConvertLangPostRequest
}

func (r ApiApiCodingConvertLangPostRequest) XAuthToken(xAuthToken string) ApiApiCodingConvertLangPostRequest {
	r.xAuthToken = &xAuthToken
	return r
}

func (r ApiApiCodingConvertLangPostRequest) ApiCodingConvertLangPostRequest(apiCodingConvertLangPostRequest ApiCodingConvertLangPostRequest) ApiApiCodingConvertLangPostRequest {
	r.apiCodingConvertLangPostRequest = &apiCodingConvertLangPostRequest
	return r
}

func (r ApiApiCodingConvertLangPostRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.ApiCodingConvertLangPostExecute(r)
}

/*
ApiCodingConvertLangPost 对选中代码转换成其他编程语言

对选中代码转换成其他编程语言

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param lang
 @return ApiApiCodingConvertLangPostRequest
*/
func (a *CodingApiService) ApiCodingConvertLangPost(ctx context.Context, lang string) ApiApiCodingConvertLangPostRequest {
	return ApiApiCodingConvertLangPostRequest{
		ApiService: a,
		ctx: ctx,
		lang: lang,
	}
}

// Execute executes the request
//  @return []string
func (a *CodingApiService) ApiCodingConvertLangPostExecute(r ApiApiCodingConvertLangPostRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CodingApiService.ApiCodingConvertLangPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/coding/convert/{lang}"
	localVarPath = strings.Replace(localVarPath, "{"+"lang"+"}", url.PathEscape(parameterValueToString(r.lang, "lang")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAuthToken == nil {
		return localVarReturnValue, nil, reportError("xAuthToken is required and must be specified")
	}
	if r.apiCodingConvertLangPostRequest == nil {
		return localVarReturnValue, nil, reportError("apiCodingConvertLangPostRequest is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-AuthToken", r.xAuthToken, "")
	// body params
	localVarPostBody = r.apiCodingConvertLangPostRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrInfo
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrInfo
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

type ApiApiCodingExplainLangPostRequest struct {
	ctx context.Context
	ApiService *CodingApiService
	xAuthToken *string
	lang string
	apiCodingExplainLangPostRequest *ApiCodingExplainLangPostRequest
}

func (r ApiApiCodingExplainLangPostRequest) XAuthToken(xAuthToken string) ApiApiCodingExplainLangPostRequest {
	r.xAuthToken = &xAuthToken
	return r
}

func (r ApiApiCodingExplainLangPostRequest) ApiCodingExplainLangPostRequest(apiCodingExplainLangPostRequest ApiCodingExplainLangPostRequest) ApiApiCodingExplainLangPostRequest {
	r.apiCodingExplainLangPostRequest = &apiCodingExplainLangPostRequest
	return r
}

func (r ApiApiCodingExplainLangPostRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.ApiCodingExplainLangPostExecute(r)
}

/*
ApiCodingExplainLangPost 解释选择代码

解释选择代码

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param lang
 @return ApiApiCodingExplainLangPostRequest
*/
func (a *CodingApiService) ApiCodingExplainLangPost(ctx context.Context, lang string) ApiApiCodingExplainLangPostRequest {
	return ApiApiCodingExplainLangPostRequest{
		ApiService: a,
		ctx: ctx,
		lang: lang,
	}
}

// Execute executes the request
//  @return []string
func (a *CodingApiService) ApiCodingExplainLangPostExecute(r ApiApiCodingExplainLangPostRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CodingApiService.ApiCodingExplainLangPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/coding/explain/{lang}"
	localVarPath = strings.Replace(localVarPath, "{"+"lang"+"}", url.PathEscape(parameterValueToString(r.lang, "lang")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAuthToken == nil {
		return localVarReturnValue, nil, reportError("xAuthToken is required and must be specified")
	}
	if r.apiCodingExplainLangPostRequest == nil {
		return localVarReturnValue, nil, reportError("apiCodingExplainLangPostRequest is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-AuthToken", r.xAuthToken, "")
	// body params
	localVarPostBody = r.apiCodingExplainLangPostRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrInfo
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrInfo
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

type ApiApiCodingFixErrorLangPostRequest struct {
	ctx context.Context
	ApiService *CodingApiService
	xAuthToken *string
	lang string
	apiCodingFixErrorLangPostRequest *ApiCodingFixErrorLangPostRequest
}

func (r ApiApiCodingFixErrorLangPostRequest) XAuthToken(xAuthToken string) ApiApiCodingFixErrorLangPostRequest {
	r.xAuthToken = &xAuthToken
	return r
}

func (r ApiApiCodingFixErrorLangPostRequest) ApiCodingFixErrorLangPostRequest(apiCodingFixErrorLangPostRequest ApiCodingFixErrorLangPostRequest) ApiApiCodingFixErrorLangPostRequest {
	r.apiCodingFixErrorLangPostRequest = &apiCodingFixErrorLangPostRequest
	return r
}

func (r ApiApiCodingFixErrorLangPostRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.ApiCodingFixErrorLangPostExecute(r)
}

/*
ApiCodingFixErrorLangPost 根据错误提示给出解决方案

根据错误提示给出解决方案

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param lang
 @return ApiApiCodingFixErrorLangPostRequest
*/
func (a *CodingApiService) ApiCodingFixErrorLangPost(ctx context.Context, lang string) ApiApiCodingFixErrorLangPostRequest {
	return ApiApiCodingFixErrorLangPostRequest{
		ApiService: a,
		ctx: ctx,
		lang: lang,
	}
}

// Execute executes the request
//  @return []string
func (a *CodingApiService) ApiCodingFixErrorLangPostExecute(r ApiApiCodingFixErrorLangPostRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CodingApiService.ApiCodingFixErrorLangPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/coding/fixError/{lang}"
	localVarPath = strings.Replace(localVarPath, "{"+"lang"+"}", url.PathEscape(parameterValueToString(r.lang, "lang")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAuthToken == nil {
		return localVarReturnValue, nil, reportError("xAuthToken is required and must be specified")
	}
	if r.apiCodingFixErrorLangPostRequest == nil {
		return localVarReturnValue, nil, reportError("apiCodingFixErrorLangPostRequest is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-AuthToken", r.xAuthToken, "")
	// body params
	localVarPostBody = r.apiCodingFixErrorLangPostRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrInfo
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrInfo
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

type ApiApiCodingGenTestLangPostRequest struct {
	ctx context.Context
	ApiService *CodingApiService
	xAuthToken *string
	lang string
	apiCodingExplainLangPostRequest *ApiCodingExplainLangPostRequest
}

func (r ApiApiCodingGenTestLangPostRequest) XAuthToken(xAuthToken string) ApiApiCodingGenTestLangPostRequest {
	r.xAuthToken = &xAuthToken
	return r
}

func (r ApiApiCodingGenTestLangPostRequest) ApiCodingExplainLangPostRequest(apiCodingExplainLangPostRequest ApiCodingExplainLangPostRequest) ApiApiCodingGenTestLangPostRequest {
	r.apiCodingExplainLangPostRequest = &apiCodingExplainLangPostRequest
	return r
}

func (r ApiApiCodingGenTestLangPostRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.ApiCodingGenTestLangPostExecute(r)
}

/*
ApiCodingGenTestLangPost 对选中函数生成测试代码

对选中函数生成测试代码

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param lang
 @return ApiApiCodingGenTestLangPostRequest
*/
func (a *CodingApiService) ApiCodingGenTestLangPost(ctx context.Context, lang string) ApiApiCodingGenTestLangPostRequest {
	return ApiApiCodingGenTestLangPostRequest{
		ApiService: a,
		ctx: ctx,
		lang: lang,
	}
}

// Execute executes the request
//  @return []string
func (a *CodingApiService) ApiCodingGenTestLangPostExecute(r ApiApiCodingGenTestLangPostRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CodingApiService.ApiCodingGenTestLangPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/coding/genTest/{lang}"
	localVarPath = strings.Replace(localVarPath, "{"+"lang"+"}", url.PathEscape(parameterValueToString(r.lang, "lang")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAuthToken == nil {
		return localVarReturnValue, nil, reportError("xAuthToken is required and must be specified")
	}
	if r.apiCodingExplainLangPostRequest == nil {
		return localVarReturnValue, nil, reportError("apiCodingExplainLangPostRequest is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-AuthToken", r.xAuthToken, "")
	// body params
	localVarPostBody = r.apiCodingExplainLangPostRequest
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
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrInfo
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrInfo
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