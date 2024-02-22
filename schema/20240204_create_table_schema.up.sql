-- Create user table
CREATE TABLE user (
    user_id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
    username varchar(255) NOT NULL,
    email varchar(255) NOT NULL unique,
    password varchar(512) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP
);

-- Create category table
CREATE TABLE category (
    category_id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name varchar(255) NOT NULL,
    user_id INTEGER NOT NULL,
    short_description varchar(512),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user(user_id)
);

-- Create wiki table
CREATE TABLE wiki (
    wiki_id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
    title varchar(255) NOT NULL,
    category_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    body MEDIUMTEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES category(category_id),
    FOREIGN KEY (user_id) REFERENCES user(user_id)
);