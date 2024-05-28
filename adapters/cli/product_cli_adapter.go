package cli

import (
	"fmt"

	"github.com/PedrodeAlmeidaFreitas/go-hex/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float32) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with name %s with price %f has been created with the status %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with name %s with price %f has been enabled",
			res.GetID(), res.GetName(), res.GetPrice())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with name %s with price %f has been disabled",
			res.GetID(), res.GetName(), res.GetPrice())
	default:
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	}
	return result, nil
}
