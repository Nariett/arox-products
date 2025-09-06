CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    slug TEXT NOT NULL
);

CREATE TABLE product(
    id SERIAL PRIMARY KEY,
    brand TEXT NOT NULL,
    name TEXT NOT NULL,
    category INT REFERENCES categories (id),
    price BIGINT NOT NULL,
    description TEXT,
    sizes JSONB NOT NULL,
    is_active BOOL NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE images(
    id SERIAL PRIMARY KEY,
    id_product INT REFERENCES product (id),
    url TEXT NOT NULL,
    is_main BOOL,
    is_active BOOL
);

INSERT INTO categories (name, slug)
VALUES
    ('Одежда', 'clothes'),
    ('Обувь', 'shoes'),
    ('Аксессуары', 'accessories');
