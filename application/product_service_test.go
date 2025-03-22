package application_test

import (
	"testing"

	"github.com/flaviomdutra/product-srv-hexagonal/application"
	mock_application "github.com/flaviomdutra/product-srv-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductServiceGet(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockProductInterface(ctrl)
	persistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{Persistence: persistence}

	result, err := service.Get("abc")

	require.Nil(t, err)
	require.Equal(t, product, result)
}
