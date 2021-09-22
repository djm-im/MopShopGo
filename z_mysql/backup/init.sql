create database if not exists
    schema_MopShop collate latin1_swedish_ci;

USE schema_MopShop;

create table if not exists products
(
    id          int auto_increment,
    name        varchar(20)              not null,
    imageUrl    varchar(1024) default '' not null,
    description varchar(2000) default '' not null,
    price       decimal(10, 2)           not null,
    constraint products_id_uindex
        unique (id),
    constraint products_name_uindex
        unique (name)
);

alter table products
    add primary key (id);


create table if not exists sessions
(
    email        varchar(50)                         not null,
    sessionToken varchar(100)                        null,
    created      timestamp default CURRENT_TIMESTAMP not null,
    expire       timestamp default CURRENT_TIMESTAMP not null,
    constraint sessions_email_uindex
        unique (email)
);

alter table sessions
    add primary key (email);

create table if not exists users
(
    id           int auto_increment,
    email        varchar(50)             not null,
    passwordHash varchar(76)  default '' not null,
    name         varchar(100) default '' not null,
    address      varchar(100) default '' not null,
    constraint users_email_uindex
        unique (email),
    constraint users_id_uindex
        unique (id)
);

alter table users
    add primary key (id);


create table if not exists cartItems
(
    cartItemId int auto_increment,
    userId     int           not null,
    productId  int           not null,
    quantity   int default 0 not null,
    constraint cartItems_cartId_uindex
        unique (cartItemId),
    constraint cartItems_products_id_fk
        foreign key (productId) references products (id),
    constraint cartItems_users_id_fk
        foreign key (userId) references users (id)
);

alter table cartItems
    add primary key (cartItemId);
