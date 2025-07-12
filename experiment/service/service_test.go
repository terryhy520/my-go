// service_test.go
package service

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
	"github.com/golang/mock/gomock"
	"github.com/terryhy520/my-go/mocks"
)

const (
	// 设置测试数据
	scope                 = "/subscriptions/123/resourceGroups/test"
	roleAssignmentNameStr = "test-assignment"
	roleDefinitionIDStr   = "/subscriptions/123/providers/Microsoft.Authorization/roleDefinitions/reader"
	principalIDStr        = "user-principal-id"
)

var (
	roleAssignmentName = roleAssignmentNameStr
	roleDefinitionID   = roleDefinitionIDStr
	principalID        = principalIDStr
)

func TestRoleAssignmentService_CreateRoleAssignment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRoleAssignmentsClientInterface(ctrl)
	service := NewRoleAssignmentService(mockClient)

	// 定义预期的调用和返回值
	mockClient.EXPECT().Create(
		gomock.Any(), // 匹配任何 context
		scope,
		roleAssignmentName,
		armauthorization.RoleAssignmentCreateParameters{
			Properties: &armauthorization.RoleAssignmentProperties{
				RoleDefinitionID: &roleDefinitionID,
				PrincipalID:      &principalID,
			},
		},
		gomock.Any(), // 匹配任何选项
	).Return(armauthorization.RoleAssignmentsClientCreateResponse{}, nil)

	// 执行测试
	err := service.CreateRoleAssignment(context.Background(), scope, roleAssignmentName, roleDefinitionID, principalID)

	// 验证结果
	if err != nil {
		t.Errorf("CreateRoleAssignment() error = %v", err)
	}
}

//func TestRoleAssignmentService_ListRoleAssignments(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockClient := mocks.NewMockRoleAssignmentsClientInterface(ctrl)
//	service := NewRoleAssignmentService(mockClient)
//
//	// 设置测试数据
//	scope := "/subscriptions/123/resourceGroups/test"
//
//	// 创建一个模拟的分页器
//	pager := &mocks.MockPager[armauthorization.RoleAssignmentsClientListForScopeResponse](ctrl)
//
//	// 设置分页器的行为
//	mockClient.EXPECT().NewListForScopePager(
//		scope,
//		gomock.Any(), // 匹配任何选项
//	).Return(pager)
//
//	// 设置分页器的 NextPage 和 Value 方法
//	pager.EXPECT().NextPage(gomock.Any()).Return(true, nil).Times(1)
//	pager.EXPECT().NextPage(gomock.Any()).Return(false, nil).Times(1)
//
//	// 设置分页器返回的数据
//	pager.EXPECT().Value().Return(armauthorization.RoleAssignmentsClientListForScopeResponse{
//		RoleAssignmentListResult: armauthorization.RoleAssignmentListResult{
//			Value: []*armauthorization.RoleAssignment{
//				{
//					Name: &roleAssignmentName,
//					Properties: &armauthorization.RoleAssignmentProperties{
//						RoleDefinitionID: &roleDefinitionID,
//						PrincipalID:      &principalID,
//					},
//				},
//			},
//		},
//	}).Times(1)
//
//	// 执行测试
//	assignments, err := service.ListRoleAssignments(context.Background(), scope)
//
//	// 验证结果
//	if err != nil {
//		t.Errorf("ListRoleAssignments() error = %v", err)
//	}
//
//	if len(assignments) != 1 {
//		t.Errorf("Expected 1 role assignment, got %d", len(assignments))
//	}
//}
