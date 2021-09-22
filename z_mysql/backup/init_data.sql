# **********************************************************************************************************************
# Users

INSERT INTO schema_MopShop.users (email, passwordHash, name, address)
VALUES ('customer@email.com', '$2a$04$PsN7IkNay9spBN0EIdzdZeVaBeQgXCGo/QnFKKdvNNuS6i6aqhMs.', '', '');

# **********************************************************************************************************************
# Products

INSERT INTO schema_MopShop.products (name, imageUrl, description, price)
VALUES ('T-Shirt', 'https://i.stack.imgur.com/GNhxO.png', 'Just a t-shirt', 10.99);

INSERT INTO schema_MopShop.products (name, imageUrl, description, price)
VALUES ('Jacket', 'https://i.stack.imgur.com/GNhxO.png', 'This is a jacket', 12.99);
