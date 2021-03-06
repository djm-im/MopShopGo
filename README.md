# MopShopGo

MoP Shop implementation in GoLang.

**Important Note**: I don't have a prior GoLang professional experience

## Important notes and assumptions

Before starting this project, it is needed to run the MySQL container.

For testing open Postman and import endpoints.

### Assumptions

* Endless supply of products. That means no limit on the number of products.
* The project contains hard-coded values for path variables and passwords.  
  This style is terrible for leas two reasons — the first, security.  
  The second reason, it is not flexible and configurable.  
  This trade-off was made for the sake of faster development.
* Carts have a simple live-cycle. No abandoned carts (endless life).
* The field `Product.Name` is chosen to be a unique value.  
  However, a better approach is to have a product code and set it as a unique value.  
  This trade-off is made because of simplicity.
* Payments are in USD.

### Improvements

* Database migration tool (like Java FlywayDb)
* Add search of products by different fields
* Add roles for users - Admin and User (customer)
* Better logging - more log messages, and more details.
* Add SwaggerUI for endpoints
* Split "Service layer" (`*Service.go` files) - it's a combination of Controller and Service layer.
* Store session tokens into a mem-cache database
* Add anonymous carts - when a user is not logged in, enable to create a cart and continue shopping.
* Use JWT tokens.

## Build and run the project

**Mysql**

To run MySQL container go to project home directory

```bash
cd .../MopShopGo # go to the project directory
mkdir -p z_data/data
docker-compose up
```

On starting the MySQL container (the first time) it creates MySQL schema and imports some predefined values. For more
details see files `.../z_mysql/backup/init.sql` and `.../z_mysql/backup/init_data.sql`.

For table users there is inserted a default customer.

```text
    email:      customer@email.com
    password:   customer
```

**Go Project**

To run the Go project

```bash
cd .../MopShopGo # go to the project directory
go get .
go build
go test
go run main.go
```

**Postman**

Hardcoded value for a user

For testing endpoints form Postman import files from project.
`.../z_postman/MopShop Postman Collection.postman_collection.json` and `.../z_postman/MopShop.postman_environment.json`.
The files contain collection of endpoints and variables which are used as variables in the collection of endpoints.

## Completed

- [x] `v0.11` Postman endpoints
    - Added Postman collection and postman variables
- [x] `v0.10` Added MySQL `docker-compose.yml` file
    - Added running database from local
    - Added `init.sql` a script to create schema
    - Added `init_data.sql` a script to insert data into schema
- [x] `v0.9` Stripe payment
    - Send stripe payment request
- [x] `v0.8` Cart Items
    - Added adding cart items
    - Improvements group cart items to cart with pre-calculated total price
- [x] `v0.7.1` Store session tokens in MySql database
    - This is a temporary solution
- [x] `v0.7` User login
    - Use session token (store it into Cookies)
    - Create Session service and repository
    - SessionRepository is in-memory "database" (more precisely, it is a map)
        - That means, after restart the server each session will be lost
        - Therefore, user have again to create a session token
- [x] `v0.6.7` `UserRepository` move code around - better organized and renaming
- [x] `v0.6.6` Implemented a method which check does user with an email exists
- [x] `v0.6.5` Postman added environment - parametrized requests
- [x] `v0.6.4` Basic Auth - parameters as constants
- [x] `v0.6.3` DRY - Admin service
    - Repeating code for checking basic authentication extracted into a func
- [x] `v0.6.2` Parametrized repository
    - Added constants with DB credentials (still hard-coded values)
- [x] `v0.6.1` DRY Repository
    - Make connection to database through a single point
- [x] `v0.6` User signup
    - Implemented user signup endpoint
    - Admin panel - fetch all users
- [x] `v0.5.2` Admin added response for Unauthorized requests
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

## Backlog (Todo list)

- [x] Postman export - collection and environment variables to a repository
- [ ] Confirm payment was successful
- [ ] Calculate total amount
- [ ] Check token validity - SessionsService
- [ ] `CartsService` split it into Controller and Service
- [ ] Carts & quantity: should we replace the path variable with a query parameter
- [ ] Add Carts (class/table) to group `CartItems`
- [ ] Store session tokens in a more persistent database like MySQL / mem-cache alternative
- [ ] Better error handling in `SessionsRepository.saveSessionToken()`
- [ ] Add admin for create, delete, and update products
- [ ] Admin/Products delete
- [ ] Admin/Products update
- [ ] Admin/Products enable image upload
- [ ] Products, add a new product - price have to be greater than zero
- [ ] Admin/Products when insert a new product chek is name a unique value (currently returns
  MySQL `Error 1062: Duplicate entry`)
- [ ] Add tests (unit and integration tests)
- [ ] Move the project into a Docker container
- [ ] MySQL don't use the root account for connections with database
- [ ] Check English wording and style

## Project Description 