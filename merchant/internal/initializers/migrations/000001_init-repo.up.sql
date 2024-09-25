-- Create Merchants table
CREATE TABLE IF NOT EXISTS merchants (
    id BIGSERIAL PRIMARY KEY ,
    name VARCHAR(255) NOT NULL,
    merchant_id VARCHAR(255) UNIQUE NOT NULL,
    description VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Create Categories table
CREATE TABLE IF NOT EXISTS categories (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    parent_id BIGINT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (parent_id) REFERENCES categories(id)
);

-- Create Products table
CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    stock INTEGER NOT NULL,
    sku VARCHAR(255) UNIQUE NOT NULL,
    merchant_id BIGINT,
    category_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (merchant_id) REFERENCES merchants(id),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

-- Create Images table
CREATE TABLE IF NOT EXISTS images (
    id BIGSERIAL PRIMARY KEY,
    url VARCHAR(255) NOT NULL,
    product_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id)
);

-- Create Variants table
CREATE TABLE IF NOT EXISTS variants (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    sku VARCHAR(255) UNIQUE NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    stock INTEGER NOT NULL,
    product_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (product_id) REFERENCES products(id)
);