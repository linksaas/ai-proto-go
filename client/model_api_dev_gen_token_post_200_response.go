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

// checks if the ApiDevGenTokenPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiDevGenTokenPost200Response{}

// ApiDevGenTokenPost200Response struct for ApiDevGenTokenPost200Response
type ApiDevGenTokenPost200Response struct {
	// 登录凭证
	Token string `json:"token"`
}

// NewApiDevGenTokenPost200Response instantiates a new ApiDevGenTokenPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDevGenTokenPost200Response(token string) *ApiDevGenTokenPost200Response {
	this := ApiDevGenTokenPost200Response{}
	this.Token = token
	return &this
}

// NewApiDevGenTokenPost200ResponseWithDefaults instantiates a new ApiDevGenTokenPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDevGenTokenPost200ResponseWithDefaults() *ApiDevGenTokenPost200Response {
	this := ApiDevGenTokenPost200Response{}
	return &this
}

// GetToken returns the Token field value
func (o *ApiDevGenTokenPost200Response) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ApiDevGenTokenPost200Response) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ApiDevGenTokenPost200Response) SetToken(v string) {
	o.Token = v
}

func (o ApiDevGenTokenPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiDevGenTokenPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

type NullableApiDevGenTokenPost200Response struct {
	value *ApiDevGenTokenPost200Response
	isSet bool
}

func (v NullableApiDevGenTokenPost200Response) Get() *ApiDevGenTokenPost200Response {
	return v.value
}

func (v *NullableApiDevGenTokenPost200Response) Set(val *ApiDevGenTokenPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDevGenTokenPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDevGenTokenPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDevGenTokenPost200Response(val *ApiDevGenTokenPost200Response) *NullableApiDevGenTokenPost200Response {
	return &NullableApiDevGenTokenPost200Response{value: val, isSet: true}
}

func (v NullableApiDevGenTokenPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDevGenTokenPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


