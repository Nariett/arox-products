CREATE TYPE gender AS ENUM(
    'female',
    'male',
    'unisex'
);

CREATE TYPE season AS ENUM (
    'winter',
    'spring',
    'summer',
    'autumn',
    'all_season'
);


CREATE TABLE product(
    id SERIAL PRIMARY KEY,
    brand TEXT NOT NULL,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    price BIGINT NOT NULL,
    description TEXT,
    gender gender NOT NULL,
    season season,
    sizes JSONB NOT NULL,
    is_active BOOL NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE TABLE images(
    id SERIAL PRIMARY KEY,
    id_product INT REFERENCES product (id),
    url TEXT NOT NULL,
    is_main BOOL
)
