CREATE TABLE IF NOT EXISTS "user"(
   id serial PRIMARY KEY,
   email TEXT NOT NULL UNIQUE,
   password TEXT NOT null
);