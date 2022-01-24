CREATE TABLE IF NOT EXISTS users (
    id serial NOT NULL,
    first_name VARCHAR(150) NOT NULL,
    last_name VARCHAR(150) NOT NULL,
    username VARCHAR(150) NOT NULL UNIQUE,
    password varchar(256) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    picture VARCHAR(256) DEFAULT 'https://placekitten.com/g/300/300',
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    CONSTRAINT pk_users PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS medicines (
    id serial NOT NULL,
    name VARCHAR(50) NOT NULL,
    price float NOT NULL,
    location VARCHAR(50) NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    CONSTRAINT pk_medicine PRIMARY KEY(id)    
);

CREATE TABLE IF NOT EXISTS promotions (
    id serial NOT NULL,
    description VARCHAR(100) NOT NULL,
    percentage float NOT NULL,
    start_date VARCHAR(10) NOT NULL,
    end_date VARCHAR(10) NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    CONSTRAINT pk_promotions PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS bills (
    id serial NOT NULL,
    created_at VARCHAR(10) DEFAULT now(),
    full_payment float NOT NULL,
    updated_at timestamp NOT NULL,
    CONSTRAINT pk_bills PRIMARY KEY(id)--,
    --CONSTRAINT fk_bills_promotions FOREIGN KEY(promotion_id) REFERENCES promotions(id)
);