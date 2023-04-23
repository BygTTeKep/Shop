CREATE TABLE users
(
    id serial not null unique,
    username varchar(255) not null unique,
    first_name varchar(255) not null,
    last_name varchar(255) not null,
    surname varchar(255),
    email varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE products
(
    id serial not null unique,
    name varchar(255) not null,
    description text,
    price float not null
);

CREATE TABLE cart
(
    id serial not null unique,
    customer_id int references users(id) on delete cascade not null
);

CREATE TABLE cart_product
(
    cart_id int references cart(id) on delete cascade not null,
    product_id int references products(id) on delete cascade not null
);

CREATE TABLE product_photo
(
    id serial not null unique,
    url varchar(255) not null,
    product_id int references products(id) on delete cascade not null
);