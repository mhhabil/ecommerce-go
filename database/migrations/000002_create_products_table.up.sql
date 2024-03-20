CREATE TABLE IF NOT EXISTS products(
   product_id SERIAL PRIMARY KEY,
   name VARCHAR (200) NOT NULL,
   price DECIMAL (15, 0) NOT NULL,
   image_url VARCHAR (300) NOT NULL,
   stock DECIMAL (15, 0) NOT NULL,
   condition VARCHAR (15) NOT NULL,
   tags VARCHAR ARRAY NOT NULL,
   is_purchasable BOOLEAN NOT NULL
);