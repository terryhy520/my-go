// Package mocks 提供了用于模拟 Azure Authorization 客户端的接口和实现
// 以便在单元测试中使用。它包含了对 `armauthorization.RoleAssignmentsClient` 的 Mock 实现。
// 通过使用 `go:generate` 指令，可以自动生成 Mock 代码，便于测试和验证 Azure 相关操作的正确性。
// generate.go
//go:generate mockgen -source=./mocks/interfaces.go -destination=./mocks/mock_authorization_client.go -package=mocks

package mocks

import (
	"context"

	runtime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// RoleAssignmentsClientInterface 定义了我们要 Mock 的接口
type RoleAssignmentsClientInterface interface {
	NewListForScopePager(scope string, options *armauthorization.RoleAssignmentsClientListForScopeOptions) *runtime.Pager[armauthorization.RoleAssignmentsClientListForScopeResponse]
	Create(context.Context, string, string, armauthorization.RoleAssignmentCreateParameters, *armauthorization.RoleAssignmentsClientCreateOptions) (armauthorization.RoleAssignmentsClientCreateResponse, error)
	Delete(context.Context, string, string, *armauthorization.RoleAssignmentsClientDeleteOptions) (armauthorization.RoleAssignmentsClientDeleteResponse, error)
}

// 将实际的客户端转换为接口类型
func NewRoleAssignmentsClientInterface(client *armauthorization.RoleAssignmentsClient) RoleAssignmentsClientInterface {
	return client
}
