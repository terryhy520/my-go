// service.go
package service

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

// RoleAssignmentService 提供角色分配管理功能
type RoleAssignmentService struct {
	client RoleAssignmentsClientInterface
}

// RoleAssignmentsClientInterface 定义客户端接口
// 仅包含 mockgen 支持的方法
type RoleAssignmentsClientInterface interface {
	Create(context.Context, string, string, armauthorization.RoleAssignmentCreateParameters, *armauthorization.RoleAssignmentsClientCreateOptions) (armauthorization.RoleAssignmentsClientCreateResponse, error)
	Delete(context.Context, string, string, *armauthorization.RoleAssignmentsClientDeleteOptions) (armauthorization.RoleAssignmentsClientDeleteResponse, error)
}

// NewRoleAssignmentService 创建一个新的 RoleAssignmentService 实例
func NewRoleAssignmentService(client RoleAssignmentsClientInterface) *RoleAssignmentService {
	return &RoleAssignmentService{client: client}
}

// CreateRoleAssignment 创建新的角色分配
func (s *RoleAssignmentService) CreateRoleAssignment(ctx context.Context, scope, name, roleDefinitionID, principalID string) error {
	_, err := s.client.Create(ctx, scope, name, armauthorization.RoleAssignmentCreateParameters{
		Properties: &armauthorization.RoleAssignmentProperties{
			RoleDefinitionID: &roleDefinitionID,
			PrincipalID:      &principalID,
		},
	}, nil)
	return err
}

// ListRoleAssignments 获取指定范围内的所有角色分配
// func (s *RoleAssignmentService) ListRoleAssignments(ctx context.Context, scope string) ([]*armauthorization.RoleAssignment, error) {
// 	pager := s.client.NewListForScopePager(scope, nil)
//
// 	var assignments []*armauthorization.RoleAssignment
// 	for pager.More() {
// 		nextResult, err := pager.NextPage(ctx)
// 		if err != nil {
// 			return nil, err
// 		}
// 		assignments = append(assignments, nextResult.Value...)
// 	}
//
// 	return assignments, nil
// }
