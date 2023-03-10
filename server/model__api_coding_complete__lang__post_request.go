/*
 * ai-proto
 *
 * ai proto for coder
 *
 * API version: 0.0.2
 * Contact: panleiming@linksaas.pro
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

type ApiCodingCompleteLangPostRequest struct {

	// 编辑器光标前面的代码
	BeforeContent string `json:"beforeContent"`

	// 编辑器光标后面的代码
	AfterContent string `json:"afterContent"`
}
