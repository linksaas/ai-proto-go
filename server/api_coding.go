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

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ApiCodingCompleteLangPost - 根据上下文补全代码
func ApiCodingCompleteLangPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ApiCodingConvertLangPost - 对选中代码转换成其他编程语言
func ApiCodingConvertLangPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ApiCodingExplainLangPost - 解释选择代码
func ApiCodingExplainLangPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ApiCodingFixErrorLangPost - 根据错误提示给出解决方案
func ApiCodingFixErrorLangPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ApiCodingGenTestLangPost - 对选中函数生成测试代码
func ApiCodingGenTestLangPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}