package cli_test

import (
	"fmt"
	"testing"

	"github.com/flaviomdutra/product-srv-hexagonal/adapters/cli"
	mock_application "github.com/flaviomdutra/product-srv-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productId := "abc"
	productName := "Product Test"
	productPrice := 25.99
	productStatus := "enabled"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	productServiceMock := mock_application.NewMockProductServiceInterface(ctrl)
	productServiceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s", productId, productName, productPrice, productStatus)

	result, err := cli.Run(productServiceMock, "create", "", productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been enabled", productName)

	result, err = cli.Run(productServiceMock, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been disabled", productName)

	result, err = cli.Run(productServiceMock, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s", productId, productName, productPrice, productStatus)

	result, err = cli.Run(productServiceMock, "get", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
