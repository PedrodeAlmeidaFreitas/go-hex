package cli_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/PedrodeAlmeidaFreitas/go-hex/adapters/cli"
	"github.com/PedrodeAlmeidaFreitas/go-hex/application"
	mock_application "github.com/PedrodeAlmeidaFreitas/go-hex/application/mocks"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := float32(25.99)
	productStatus := application.ENABLED
	productId := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	expectedResult := fmt.Sprintf("Product ID %s with name %s with price %f has been created with the status %s",
		productId, productName, productPrice, productStatus)

	// Action Create
	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf("Product ID %s with name %s with price %f has been enabled",
		productId, productName, productPrice)

	// Action Enable
	result, err = cli.Run(service, "enable", "abc", "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf("Product ID %s with name %s with price %f has been disabled",
		productId, productName, productPrice)

	// Action Disable
	result, err = cli.Run(service, "disable", "abc", "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	// Action default
	expectedResult = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		productId, productName, productPrice, productStatus)

	result, err = cli.Run(service, "", "abc", "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)
}
