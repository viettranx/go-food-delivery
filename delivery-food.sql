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
	id int auto_increment primary key,
	user_id int not null,
	city_id int not null,
	addr varchar(255) not null,
	status int default 1 not null,
	created_at timestamp default CURRENT_TIMESTAMP null,
	updated_at timestamp null
);

create table categories
(
	id int auto_increment primary key,
	name varchar(100) not null,
	description text null,
	status int default 1 not null,
	created_at timestamp default CURRENT_TIMESTAMP null,
	updated_at timestamp null
);

create table category_restaurants
(
	category_id int not null,
	restaurant_id int not null,
	status int default 1 not null,
	created_at timestamp default CURRENT_TIMESTAMP null,
	updated_at timestamp null,
	primary key (restaurant_id, category_id)
);

create table cities
(
	id int auto_increment primary key,
	title varchar(100) not null,
	status int default 1 not null,
	created_at timestamp default CURRENT_TIMESTAMP null,
	updated_at timestamp null
);

create table food_likes
(
	user_id int not null,
	food_id int not null,
	primary key (user_id, food_id)
);

create table food_ratings
(
	id int auto_increment primary key,
	user_id int not null,
	food_id int not null,
	point float default 0 null,
	comment text null,
	status int default 1 not null,
	created_at timestamp default CURRENT_TIMESTAMP null,
	updated_at timestamp null
);

create table foods
(
	id int auto_increment primary key,
	name varchar(255) not null,
	description text null,
	price float not null,
	images json not null,
	status int default 1 not null,
	created_at timestamp default CURRENT_TIMESTAMP null,
	updated_at int null
);

create table order_details
(
	id int auto_increment primary key,
	order_id int not null,
	food_origin json null,
	price float not null,
	quantity int not null,
	discount float default 0 null,
    status int default 1 not null,
	created_at timestamp default CURRENT_TIMESTAMP null,
	updated_at timestamp null
);

create table orders
(
	id int auto_increment primary key,
	user_id int not null,
	total_price float not null,
	status int default 1 not null,
	created_at timestamp default CURRENT_TIMESTAMP null,
	updated_at timestamp null
);

create table restaurant_likes
(
	user_id int not null,
	restaurant_id int not null,
	constraint restaurant_likes_pk unique (restaurant_id, user_id)
);

create table restaurant_ratings
(
	id int auto_increment primary key,
	restaurant_id int not null,
	point float default 0 not null,
	comment text null,
	status int default 1 not null,
	created_at timestamp default CURRENT_TIMESTAMP null,
	updated_at timestamp null
);

create table restaurants
(
	id int auto_increment primary key,
	name varchar(50) not null,
	addr varchar(255) not null,
	logo json not null,
	status int default 1 not null,
	created_at timestamp default CURRENT_TIMESTAMP null,
	updated_at timestamp null
);

create table users
(
	id int auto_increment primary key,
	email varchar(50) not null,
	password varchar(50) not null,
	last_name varchar(50) not null,
	first_name varchar(50) not null,
	phone varchar(20) null,
	avatar json null,
	status int default 1 not null,
	created_at timestamp default CURRENT_TIMESTAMP null,
	updated_at timestamp null,
	constraint users_email_uindex unique (email)
);

alter table restaurant_ratings add user_id int not null after id;
alter table food_likes add created_at timestamp default current_timestamp null;
create index food_likes_food_id_index on food_likes (food_id);
alter table restaurant_likes add created_at timestamp default current_timestamp null;
create index restaurant_likes_user_id_index on restaurant_likes (user_id);