package mocks_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/terryhy520/my-go/mocks"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
)

func TestMockRoleAssignmentsClientInterface_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRoleAssignmentsClientInterface(ctrl)

	mockClient.EXPECT().Create(
		gomock.Any(),
		"scope",
		"name",
		armauthorization.RoleAssignmentCreateParameters{},
		gomock.Any(),
	).Return(armauthorization.RoleAssignmentsClientCreateResponse{}, nil)

	_, err := mockClient.Create(context.Background(), "scope", "name", armauthorization.RoleAssignmentCreateParameters{}, nil)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestMockRoleAssignmentsClientInterface_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockRoleAssignmentsClientInterface(ctrl)

	mockClient.EXPECT().Delete(
		gomock.Any(),
		"scope",
		"name",
		gomock.Any(),
	).Return(armauthorization.RoleAssignmentsClientDeleteResponse{}, nil)

	_, err := mockClient.Delete(context.Background(), "scope", "name", nil)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
