package products

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

type Repository struct {
    database *sql.DB
}

var mysqlDatabase = Repository{
    database: connectDb(),
}

func connectDb() *sql.DB {
    database, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/schema_MopShop")

    if err != nil {
        log.Fatalf("Cannot conntect to database. Terminated.")
    }

    return database
}

func getAllProducts() ([]Product, error) {
    result, err := mysqlDatabase.database.Query("" +
        "SELECT * " +
        "FROM products")
    defer result.Close()

    if err != nil {
        log.Printf("Cannot read data from database.")

        return []Product{}, err
    }

    var products []Product
    for result.Next() {
        var product Product
        err = result.Scan(&product.Id, &product.Name, &product.ImageUrl, &product.Description, &product.Price)

        if err != nil {
            log.Printf("Exception %s ", err.Error())

            return []Product{}, err
        }

        products = append(products, product)
    }

    return products, nil
}

func getProduct(productId int) (Product, error) {
    result, err := mysqlDatabase.database.Query(""+
        "SELECT * "+
        "FROM products "+
        "WHERE id = ?", productId)
    defer result.Close()

    if err != nil {

    }

    if result.Next() {
        var product Product
        err = result.Scan(&product.Id, &product.Name, &product.ImageUrl, &product.Description, &product.Price)

        if err != nil {
            log.Printf("Exception %s ", err.Error())

            return Product{}, err
        }

        return product, nil
    }

    return Product{}, fmt.Errorf("Cannot find a product with id %d.", productId)
}

func addProduct(product Product) (Product, error) {
    exec, err := mysqlDatabase.database.Exec(""+
        "INSERT INTO products"+
        "   (name, imageUrl, description, price)"+
        "   VALUE"+
        "   (?, ?, ?, ?)",
        product.Name, product.ImageUrl, product.Description, product.Price)
    if err != nil {
        return Product{}, err
    }

    id, _ := exec.LastInsertId()
    product.Id = id

    return product, nil
}