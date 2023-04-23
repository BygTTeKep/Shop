package model_test

import (
	"testing"

	"github.com/g91TeJl/Shop/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestProductsValidate(t *testing.T) {
	testCase := []struct {
		name    string
		p       func() *model.Products
		isValid bool
	}{
		{
			name: "valid",
			p: func() *model.Products {
				return model.TestProduct(t)
			},
			isValid: true,
		},
		{
			name: "empty name",
			p: func() *model.Products {
				p := model.TestProduct(t)
				p.Name = ""
				return p
			},
			isValid: false,
		},
		{
			name: "incorect price",
			p: func() *model.Products {
				p := model.TestProduct(t)
				p.Price = 0
				return p
			},
			isValid: false,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.p().Validate())
			} else {
				assert.Error(t, tc.p().Validate())
			}
		})
	}
}
