package model_test

import (
	"testing"

	"github.com/g91TeJl/Shop/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestProductPhoto_Validate(t *testing.T) {
	testCase := []struct {
		name    string
		ph      func() *model.ProductPhoto
		isValid bool
	}{
		{
			name: "valid",
			ph: func() *model.ProductPhoto {
				return model.TestProductPhoto(t)
			},
			isValid: true,
		},
		{
			name: "empty url",
			ph: func() *model.ProductPhoto {
				ph := model.TestProductPhoto(t)
				ph.Url = ""
				return ph
			},
			isValid: false,
		},
		{
			name: "empty product id",
			ph: func() *model.ProductPhoto {
				ph := model.TestProductPhoto(t)
				ph.Product_id = 0
				return ph
			},
			isValid: false,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.ph().Validate())
			} else {
				assert.Error(t, tc.ph().Validate())
			}
		})
	}
}
