CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    slug TEXT NOT NULL UNIQUE
);

CREATE TABLE products(
    id SERIAL PRIMARY KEY,
    brand TEXT NOT NULL,
    name TEXT NOT NULL,
    category_id INT REFERENCES categories (id),
    price BIGINT NOT NULL,
    description TEXT,
    sizes JSONB NOT NULL,
    is_active BOOL NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE images(
    id SERIAL PRIMARY KEY,
    id_product INT REFERENCES products (id),
    url TEXT NOT NULL,
    is_main BOOL,
    is_active BOOL
);
