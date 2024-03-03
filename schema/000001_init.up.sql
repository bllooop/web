CREATE TABLE shops
(
    id serial not null unique,
    shopname varchar(255) not null unique,
    password varchar(255) not null
);
CREATE TABLE productlist
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255)
);

CREATE TABLE shoplist
(
    id serial not null unique,
    shop_id int references shops(id) on delete cascade not null,
    list_id int references productlist(id) on delete cascade not null
);
CREATE TABLE productItem
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255),
    price int not null,
    expiration varchar(255) not null

);
CREATE TABLE listsItem
(
    id serial not null unique,
    list_id int references productlist(id) on delete cascade not null,
    item_id int references productItem(id) on delete cascade not null
);
