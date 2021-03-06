package main

import (
    "MopShopGo/admin"
    "MopShopGo/carts"
    "MopShopGo/products"
    "MopShopGo/purchases"
    "MopShopGo/users"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

func main() {
    log.Println("Mop Shop Go started...")

    router := mux.NewRouter()

    // admin
    router.HandleFunc("/admin/users", admin.GetAllUsers).Methods("GET")
    router.HandleFunc("/admin/products/add", admin.AddProduct).Methods("POST")
    router.HandleFunc("/admin/products/update/{productId}", admin.UpdateProduct).Methods("PUT")
    router.HandleFunc("/admin/products/upload/image/{productId}", admin.UploadImage).Methods("POST")
    router.HandleFunc("/admin/products/delete/{productId}", admin.DeleteProduct).Methods("DELETE")
    router.HandleFunc("/admin/products/statistics", admin.ProductStatistics).Methods("GET")
    router.HandleFunc("/admin/products/statistics/{productId}", admin.ProductStatistics).Methods("GET")
    router.HandleFunc("/admin/products/statistics/user/{userId}", admin.ProductStatistics).Methods("GET")

    // users
    router.HandleFunc("/users/signup", users.Signup).Methods("POST")
    router.HandleFunc("/users/login", users.Login).Methods("POST")
    router.HandleFunc("/user/logout", users.Logout).Methods("POST")

    // products
    router.HandleFunc("/products/all", products.GetAll).Methods("GET")
    router.HandleFunc("/products/get/{productId}", products.GetProduct).Methods("GET")

    // cart
    router.HandleFunc("/carts/user/{userId}/get", carts.GetCart).Methods("GET")
    router.HandleFunc("/carts/user/{userId}/add/product/{productId}/quantity/{quantity}", carts.AddProduct).Methods("PATCH")
    router.HandleFunc("/carts/user/{userId}remove/product/{productId}", carts.RemoveProduct).Methods("PATCH")
    router.HandleFunc("/carts/user/{userId}/delete", carts.DeleteCart).Methods("DELETE")

    // Purchase
    router.HandleFunc("/purchases/buy/user/{userId}", purchases.Buy).Methods("POST")

    log.Fatal(http.ListenAndServe(":8000", router))
}
