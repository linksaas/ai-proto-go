/*
 * ai-proto
 *
 * ai proto for coder
 *
 * API version: 0.0.1
 * Contact: panleiming@linksaas.pro
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

type ApiDevGenTokenPostRequest struct {

	// 上下文信息
	ContextValue string `json:"contextValue"`

	// 随机字符串，加密因子。需要32位长度以上
	RandomStr string `json:"randomStr"`

	// 共享密钥
	Secret string `json:"secret"`
}
