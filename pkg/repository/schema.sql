CREATE TABLE books (
    id   BIGSERIAL PRIMARY KEY,
    title varchar NOT NULL,
    description  text NOT NULL,
    author_name varchar NOT NULL,
    price integer
);