package products

import "errors"

type ProductRepository struct {
    ProductsDb []Product
}

// todo: mock database -- replace with a real db instance
var db = ProductRepository{
    ProductsDb: []Product{
        {
            Id:          1,
            Name:        "T-Shirt",
            ImageUrl:    "https://i.stack.imgur.com/GNhxO.png",
            Description: "Just another t-shirt",
            Price:       10.99,
        },
        {
            Id:          2,
            Name:        "Jacket",
            ImageUrl:    "https://i.stack.imgur.com/GNhxO.png",
            Description: "A jacket",
            Price:       12.99,
        },
        {
            Id:          3,
            Name:        "Shoes",
            ImageUrl:    "https://i.stack.imgur.com/GNhxO.png",
            Description: "White Snickers",
            Price:       99.99,
        }},
}

func getAllProducts() []Product {
    products := db.ProductsDb

    return products
}

func getProduct(productId int) (Product, error) {
    for _, product := range db.ProductsDb {
        if product.Id == productId {
            return product, nil
        }
    }
    return Product{}, errors.New("not found product")
}
