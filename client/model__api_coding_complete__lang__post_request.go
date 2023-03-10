/*
ai-proto

ai proto for coder

API version: 0.0.2
Contact: panleiming@linksaas.pro
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ApiCodingCompleteLangPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiCodingCompleteLangPostRequest{}

// ApiCodingCompleteLangPostRequest struct for ApiCodingCompleteLangPostRequest
type ApiCodingCompleteLangPostRequest struct {
	// 编辑器光标前面的代码
	BeforeContent string `json:"beforeContent"`
	// 编辑器光标后面的代码
	AfterContent string `json:"afterContent"`
}

// NewApiCodingCompleteLangPostRequest instantiates a new ApiCodingCompleteLangPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiCodingCompleteLangPostRequest(beforeContent string, afterContent string) *ApiCodingCompleteLangPostRequest {
	this := ApiCodingCompleteLangPostRequest{}
	this.BeforeContent = beforeContent
	this.AfterContent = afterContent
	return &this
}

// NewApiCodingCompleteLangPostRequestWithDefaults instantiates a new ApiCodingCompleteLangPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiCodingCompleteLangPostRequestWithDefaults() *ApiCodingCompleteLangPostRequest {
	this := ApiCodingCompleteLangPostRequest{}
	return &this
}

// GetBeforeContent returns the BeforeContent field value
func (o *ApiCodingCompleteLangPostRequest) GetBeforeContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BeforeContent
}

// GetBeforeContentOk returns a tuple with the BeforeContent field value
// and a boolean to check if the value has been set.
func (o *ApiCodingCompleteLangPostRequest) GetBeforeContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BeforeContent, true
}

// SetBeforeContent sets field value
func (o *ApiCodingCompleteLangPostRequest) SetBeforeContent(v string) {
	o.BeforeContent = v
}

// GetAfterContent returns the AfterContent field value
func (o *ApiCodingCompleteLangPostRequest) GetAfterContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AfterContent
}

// GetAfterContentOk returns a tuple with the AfterContent field value
// and a boolean to check if the value has been set.
func (o *ApiCodingCompleteLangPostRequest) GetAfterContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AfterContent, true
}

// SetAfterContent sets field value
func (o *ApiCodingCompleteLangPostRequest) SetAfterContent(v string) {
	o.AfterContent = v
}

func (o ApiCodingCompleteLangPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiCodingCompleteLangPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["beforeContent"] = o.BeforeContent
	toSerialize["afterContent"] = o.AfterContent
	return toSerialize, nil
}

type NullableApiCodingCompleteLangPostRequest struct {
	value *ApiCodingCompleteLangPostRequest
	isSet bool
}

func (v NullableApiCodingCompleteLangPostRequest) Get() *ApiCodingCompleteLangPostRequest {
	return v.value
}

func (v *NullableApiCodingCompleteLangPostRequest) Set(val *ApiCodingCompleteLangPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiCodingCompleteLangPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiCodingCompleteLangPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiCodingCompleteLangPostRequest(val *ApiCodingCompleteLangPostRequest) *NullableApiCodingCompleteLangPostRequest {
	return &NullableApiCodingCompleteLangPostRequest{value: val, isSet: true}
}

func (v NullableApiCodingCompleteLangPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiCodingCompleteLangPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


