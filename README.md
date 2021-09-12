# MopShopGo

MoP Shop implementation in GoLang.

### Important notes and assumptions

Before starting this project, it is needed to run the MySQL container.

For testing open Postman and import endpoints.

**Assumptions**

* Endless supply of products. That means no limit on the number of products.
* The project contains hard-coded values for path variables and passwords.  
  This style is terrible for leas two reasons—the first, security.  
  The second reason, it is not flexible and configurable.  
  This trade-off was made for the sake of faster development.
* The field `Product.Name` is chosen to be a unique value.  
  However, a better approach is to have a product code and set it as a unique value.  
  This trade-off is made because of simplicity.
* Better logging - more details.

**Improvements**

* Database migration tool (like Java FlywayDb)
* Add search of products by different fields

#### Build and run the project

```bash
 go get .
 go build
 go test
 go run main.go
```

### Completed

- [x] `v0.5.1` DRY refactoring
  - Extracted a method for building responses with error
- [x] `v0.5` Admin Panel
    - Implemented method for adding products
- [x] `v0.4` Get products from database
    - Read data from database (supported only reading of products)
- [x] `v0.3` Convert indents in code
    - Use spaces instead tabs
- [x] `v0.2` Added mock Products Repository
    - used in memory "database" for products (slice)
- [x] `v0.1` Initialized project
    - Create a Go project
    - Added Service (business logic) Layer with planned method
    - Git initialize
    - Git ignore: `.idea` (IDE dir), and `MopShopGo` (binary fine)

### TODO

- [ ] Add admin for create, delete, and update products
- [ ] Admin/Products delete
- [ ] Admin/Products update
- [ ] Admin/Products enable image upload
- [ ] Products, add a new product - price have to be greater than zero
- [ ] Admin/Products when insert a new product chek is name a unique value (currently returns MySQL `Error 1062: Duplicate entry`)
- [ ] Add MySql Docker container
- [ ] Add a Postman collection
- [ ] Add link to MySql container repository
- [ ] Add tests (unit and integration tests)
- [ ] Move the project into a Docker container

### Project Description 