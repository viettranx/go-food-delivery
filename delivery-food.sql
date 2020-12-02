drop table if exists addresses;

drop table if exists categories;

drop table if exists category_restaurants;

drop table if exists cities;

drop table if exists food_likes;

drop table if exists food_ratings;

drop table if exists foods;

drop table if exists order_details;

drop table if exists orders;

drop table if exists restaurant_likes;

drop table if exists restaurant_ratings;

drop table if exists restaurants;

drop table if exists users;

create table addresses
(
    id         int auto_increment primary key,
    user_id    int                                                             not null,
    city_id    int                                                             not null,
    addr       varchar(255)                                                    not null,
    lat        double                                                          null,
    lng        double                                                          null,
    status     int       default 1                                             not null,
    created_at timestamp default current_timestamp                             null,
    updated_at timestamp default current_timestamp on update current_timestamp null
);

create table categories
(
    id          int auto_increment primary key,
    name        varchar(100)                                                    not null,
    description text                                                            null,
    icon        json                                                            null,
    status      int       default 1                                             not null,
    created_at  timestamp default current_timestamp                             null,
    updated_at  timestamp default current_timestamp on update current_timestamp null
);

create table category_restaurants
(
    category_id   int                                                             not null,
    restaurant_id int                                                             not null,
    status        int       default 1                                             not null,
    created_at    timestamp default current_timestamp                             null,
    updated_at    timestamp default current_timestamp on update current_timestamp null,
    primary key (category_id, restaurant_id),
    index (restaurant_id)
);

create table cities
(
    id         int auto_increment primary key,
    title      varchar(100)                                                    not null,
    status     int       default 1                                             not null,
    created_at timestamp default current_timestamp                             null,
    updated_at timestamp default current_timestamp on update current_timestamp null
);

create table food_likes
(
    user_id    int                                 not null,
    food_id    int                                 not null,
    created_at timestamp default current_timestamp null,
    primary key (user_id, food_id),
    index (food_id)
);

create table food_ratings
(
    id         int auto_increment primary key,
    user_id    int                                                             not null,
    food_id    int                                                             not null,
    point      float     default 0                                             null,
    comment    text                                                            null,
    status     int       default 1                                             not null,
    created_at timestamp default current_timestamp                             null,
    updated_at timestamp default current_timestamp on update current_timestamp null
);

create table foods
(
    id          int auto_increment primary key,
    name        varchar(255)                                                    not null,
    description text                                                            null,
    price       float                                                           not null,
    images      json                                                            not null,
    status      int       default 1                                             not null,
    created_at  timestamp default current_timestamp                             null,
    updated_at  timestamp default current_timestamp on update current_timestamp null
);

create table order_details
(
    id          int auto_increment primary key,
    order_id    int                                                             not null,
    food_origin json                                                            null,
    price       float                                                           not null,
    quantity    int                                                             not null,
    discount    float     default 0                                             null,
    status      int       default 1                                             not null,
    created_at  timestamp default current_timestamp                             null,
    updated_at  timestamp default current_timestamp on update current_timestamp null
);

create table orders
(
    id          int auto_increment primary key,
    user_id     int                                                             not null,
    total_price float                                                           not null,
    status      int       default 1                                             not null,
    created_at  timestamp default current_timestamp                             null,
    updated_at  timestamp default current_timestamp on update current_timestamp null
);

create table restaurant_likes
(
    restaurant_id int                                 not null,
    user_id       int                                 not null,
    created_at    timestamp default current_timestamp null,
    primary key (restaurant_id, user_id),
    index (user_id)
);

create table restaurant_ratings
(
    id            int auto_increment primary key,
    user_id       int                                                             not null,
    restaurant_id int                                                             not null,
    point         float     default 0                                             not null,
    comment       text                                                            null,
    status        int       default 1                                             not null,
    created_at    timestamp default current_timestamp                             null,
    updated_at    timestamp default current_timestamp on update current_timestamp null
);

create table restaurants
(
    id         int auto_increment primary key,
    name       varchar(50)                                                     not null,
    addr       varchar(255)                                                    not null,
    lat        double                                                          null,
    lng        double                                                          null,
    logo       json                                                            not null,
    status     int       default 1                                             not null,
    created_at timestamp default current_timestamp                             null,
    updated_at timestamp default current_timestamp on update current_timestamp null
);

create table users
(
    id         int auto_increment primary key,
    email      varchar(50)                                                                  not null,
    password   varchar(50)                                                                  not null,
    last_name  varchar(50)                                                                  not null,
    first_name varchar(50)                                                                  not null,
    phone      varchar(20)                                                                  null,
    roles      enum ('user', 'admin') default 'user'                                        not null,
    salt       varchar(50)                                                                  null,
    avatar     json                                                                         null,
    status     int                    default 1                                             not null,
    created_at timestamp              default current_timestamp                             null,
    updated_at timestamp              default current_timestamp on update current_timestamp null,
    unique (email)
);

create table images
(
	id      int auto_increment primary key,
	url     text not null,
	width   int not null,
	height  int not null
);

CREATE TABLE carts
(
    user_id    INT(11)   NOT NULL,
    food_id    INT(11)   NOT NULL,
    quantity   int            default 1 NOT NULL,
    status     BOOLEAN   NOT NULL,
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    primary key (food_id, user_id)
);
