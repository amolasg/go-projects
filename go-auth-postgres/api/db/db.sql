

CREATE TABLE users(
id serial PRIMARY KEY,
first_name text,
last_name text,
email_id text NOT NULL UNIQUE,
mobile_number text,
password_hash text,
created_date text
);
