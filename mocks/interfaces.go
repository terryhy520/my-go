package mocks

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// RoleAssignmentsClientInterface 定义了我们要 Mock 的接口
// 便于在测试中对 Azure 授权相关操作进行模拟
// 可通过 mockgen 工具生成对应的 mock 实现
//go:generate mockgen -source=interfaces.go -destination=mock_authorization_client.go -package=mocks

type RoleAssignmentsClientInterface interface {
	Create(ctx context.Context, scope string, roleAssignmentName string, parameters armauthorization.RoleAssignmentCreateParameters, options *armauthorization.RoleAssignmentsClientCreateOptions) (armauthorization.RoleAssignmentsClientCreateResponse, error)
	Delete(ctx context.Context, scope string, roleAssignmentName string, options *armauthorization.RoleAssignmentsClientDeleteOptions) (armauthorization.RoleAssignmentsClientDeleteResponse, error)
}
