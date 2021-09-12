# MopShopGo

MoP Shop implementation in GoLang.

### Important notes and assumptions

Before starting this project, it is needed to run the MySQL container.

For testing open Postman and import endpoints.

**Assumptions**

* Endless supply of products. That means no limit on the number of products.

**Improvement**

* Database migration tool (like Java FlywayDb)

#### Build and run the project

```bash
 go get .
 go build
 go test
 go run main.go
```

### Completed

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
- [ ] Add MySql Docker container
- [ ] Add a Postman collection
- [ ] Add link to MySql container repository
- [ ] Add service methods
- [ ] Add tests (unit and integration tests)
- [ ] Move the project into a Docker container

### Project Description 