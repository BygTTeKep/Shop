package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Products struct {
	Id          uint    `json:"id_product"`
	Name        string  `json:"name_product"`
	Description string  `json:"description_product"`
	Price       float32 `json:"price_product"`
}

func (p *Products) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Description, validation.Required),
		validation.Field(&p.Price, validation.By(requiredIf(p.Price >= 0.00000001))),
	)
}
