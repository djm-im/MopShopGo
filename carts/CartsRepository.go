package carts

import (
    "MopShopGo/repository"
)

func insertCartItem(cartItem CartItems) error {
    _, err := repository.GetMysql().Query(""+
        "INSERT INTO cartItems "+
        "(userId, productId, quantity) "+
        "VALUES "+
        "(?, ?, ?)",
        cartItem.UserId, cartItem.ProductId, cartItem.Quantity)

    return err
}
