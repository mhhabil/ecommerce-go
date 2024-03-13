CREATE TABLE IF NOT EXISTS users(
   user_id serial PRIMARY KEY,
   username VARCHAR (20) UNIQUE NOT NULL,
   name VARCHAR (100) NOT NULL,
   password VARCHAR (150) NOT NULL
);