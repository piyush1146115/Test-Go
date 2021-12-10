package SOLID

import (
	"github.com/ahmetb/go-linq"
	"github.com/gin-gonic/gin"
)

type PermissionProvider interface {
	Type() string
	GetPermissions(ctx *gin.Context) []string
}

type PermissionChecker struct {
	providers []PermissionProvider
	//
	// some fields
	//
}

func (c *PermissionChecker) HasPermission(ctx *gin.Context, name string) bool {
	var permissions []string
	for _, provider := range c.providers {
		if ctx.GetString("authType") != provider.Type() {
			continue
		}

		permissions = provider.GetPermissions(ctx)
		break
	}

	var result []string
	linq.From(permissions).
		Where(func(permission interface{}) bool {
			return permission.(string) == name
		}).ToSlice(&result)

	return len(result) > 0
}

//When we do not respect the Open/Closed Principle
//type PermissionChecker struct {
//	//
//	// some fields
//	//
//}
//
//func (c *PermissionChecker) HasPermission(ctx *gin.Context, name string) bool {
//	var permissions []string
//	switch ctx.GetString("authType") {
//	case "jwt":
//		permissions = c.extractPermissionsFromJwt(ctx.Request.Header)
//	case "basic":
//		permissions = c.getPermissionsForBasicAuth(ctx.Request.Header)
//	case "applicationKey":
//		permissions = c.getPermissionsForApplicationKey(ctx.Query("applicationKey"))
//	}
//
//	var result []string
//	linq.From(permissions).
//		Where(func(permission interface{}) bool {
//			return permission.(string) == name
//		}).ToSlice(&result)
//
//	return len(result) > 0
//}
//
//func (c *PermissionChecker) getPermissionsForApplicationKey(key string) []string {
//	var result []string
//	//
//	// extract JWT from the request header
//	//
//	return result
//}
//
//func (c *PermissionChecker) getPermissionsForBasicAuth(h http.Header) []string {
//	var result []string
//	//
//	// extract JWT from the request header
//	//
//	return result
//}
//
//func (c *PermissionChecker) extractPermissionsFromJwt(h http.Header) []string {
//	var result []string
//	//
//	// extract JWT from the request header
//	//
//	return result
//}

//
//Such implementation brings the explosion of issues:
//PermissionChecker keeps the logic initially handled somewhere else.
//Any adaptation of authorization logic, which may be a different module, requires adaptation in PermissionChecker.
//To add a new way of extracting permissions, we always need to modify PermissionChecker.
//Logic inside PermissionChecker inevitably grows with each new authorization flow.
//Unit testing for PermissionChecker contains too many technical details about different extractions of permission.
