CREATE SCHEMA public AUTHORIZATION postgres;
CREATE TABLE public.credit_card (
	user_id INT NOT NULL,
	credit_number varchar NOT NULL,
	expiration_month VARCHAR NOT NULL,
	expiration_yaer VARCHAR NOT NULL,
	cvv varchar NOT NULL,
	customer_id VARCHAR,
	token_stripe VARCHAR,
	card_id VARCHAR
);
CREATE TABLE public."token" (
	user_id INT NULL,
	tokenstr VARCHAR NULL
);
CREATE TABLE public."user" (
	user_id INT NOT NULL,
	user_name VARCHAR(50) NOT NULL,
	email VARCHAR(50) NULL,
	password VARCHAR(50) NULL,
	role VARCHAR NULL
);