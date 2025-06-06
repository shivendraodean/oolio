package repository

import "oolio/api-ecommerce/src/internal/model"

type ProductRepository interface {
	FindProduct(productID int64) model.Product
	ListProducts() []model.Product
}

type ProductRepositoryImpl struct{}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (r *ProductRepositoryImpl) FindProduct(productID int64) model.Product {
	product := model.Product{
		ID:       "10",
		Name:     "Chicken Waffle",
		Price:    13.3,
		Category: "Waffle",
		Image: model.ProductImage{
			Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-waffle-thumbnail.jpg",
			Mobile:    "https://orderfoodonline.deno.dev/public/images/image-waffle-mobile.jpg",
			Tablet:    "https://orderfoodonline.deno.dev/public/images/image-waffle-tablet.jpg",
			Desktop:   "https://orderfoodonline.deno.dev/public/images/image-waffle-desktop.jpg",
		},
	}

	return product
}

func (r *ProductRepositoryImpl) ListProducts() []model.Product {
	products := []model.Product{
		{
			ID:       "10",
			Name:     "Chicken Waffle",
			Price:    13.3,
			Category: "Waffle",
			Image: model.ProductImage{
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-waffle-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-waffle-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-waffle-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-waffle-desktop.jpg",
			},
		},
		{
			ID:       "11",
			Name:     "Chocolate Waffle",
			Price:    15.5,
			Category: "Waffle",
			Image: model.ProductImage{
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-waffle-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-waffle-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-waffle-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-waffle-desktop.jpg",
			},
		},
		{
			ID:       "12",
			Name:     "Strawberry Waffle",
			Price:    14.2,
			Category: "Waffle",
			Image: model.ProductImage{
				Thumbnail: "https://orderfoodonline.deno.dev/public/images/image-waffle-thumbnail.jpg",
				Mobile:    "https://orderfoodonline.deno.dev/public/images/image-waffle-mobile.jpg",
				Tablet:    "https://orderfoodonline.deno.dev/public/images/image-waffle-tablet.jpg",
				Desktop:   "https://orderfoodonline.deno.dev/public/images/image-waffle-desktop.jpg",
			},
		},
	}

	return products
}
