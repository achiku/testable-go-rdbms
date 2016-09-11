BEGIN;


CREATE TABLE item (
    id BIGSERIAL NOT NULL, 
    name TEXT NOT NULL, 
    price NUMERIC, 
    description TEXT, 
    PRIMARY KEY (id)
);

CREATE TABLE user_account (
    id BIGSERIAL NOT NULL, 
    email TEXT NOT NULL, 
    gender TEXT NOT NULL, 
    birthday DATE NOT NULL, 
    password TEXT NOT NULL, 
    registered_at TIMESTAMP WITH TIME ZONE NOT NULL, 
    PRIMARY KEY (id)
);

CREATE TABLE access_token (
    account_id BIGINT NOT NULL, 
    token TEXT NOT NULL, 
    is_valid BOOLEAN NOT NULL, 
    generated_at TIMESTAMP WITH TIME ZONE NOT NULL, 
    PRIMARY KEY (account_id), 
    FOREIGN KEY(account_id) REFERENCES user_account (id), 
    UNIQUE (token)
);

CREATE TABLE sale (
    id BIGSERIAL NOT NULL, 
    account_id BIGINT NOT NULL, 
    item_id BIGINT NOT NULL, 
    paid_amount NUMERIC NOT NULL, 
    sold_at TIMESTAMP WITH TIME ZONE NOT NULL, 
    PRIMARY KEY (id), 
    FOREIGN KEY(account_id) REFERENCES user_account (id), 
    FOREIGN KEY(item_id) REFERENCES item (id)
);

CREATE TABLE username (
    account_id BIGINT NOT NULL, 
    lower_name TEXT NOT NULL, 
    display_name TEXT NOT NULL, 
    PRIMARY KEY (account_id), 
    FOREIGN KEY(account_id) REFERENCES user_account (id), 
    UNIQUE (lower_name)
);

COMMIT;
