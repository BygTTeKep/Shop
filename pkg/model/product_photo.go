package model

import validation "github.com/go-ozzo/ozzo-validation"

type ProductPhoto struct {
	Id         int    `json:"id"`
	Url        string `json:"url"`
	Product_id uint   `json:"product_id"`
}

func (ph *ProductPhoto) Validate() error {
	return validation.ValidateStruct(
		ph,
		validation.Field(&ph.Url, validation.Required),
		validation.Field(&ph.Product_id, validation.By(requiredIf(ph.Product_id > 0))),
	)
}
