package application

import (
	"testing"

	"github.com/stretchr/testify/require"

	uuid "github.com/satori/go.uuid"
)

func TestProduct_Enabled(t *testing.T) {
	product := Product{}
	product.Name = "testName"
	product.Status = DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disabled(t *testing.T) {
	product := Product{}
	product.Name = "testName"
	product.Status = DISABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "test"
	product.Status = DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "WRONG_STATUS"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -1
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal to zero", err.Error())

}
