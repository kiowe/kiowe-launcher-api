CREATE TABLE IF NOT EXISTS users_groups
(
    id          INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name        VARCHAR(50)  NOT NULL UNIQUE,
    description VARCHAR(300) NULL
);

CREATE TABLE IF NOT EXISTS users_statuses
(
    id   INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(20) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS games_categories
(
    id          INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name        VARCHAR(50)  NOT NULL UNIQUE,
    description VARCHAR(300) NULL
);

CREATE TABLE IF NOT EXISTS dev_pub_account
(
    id          INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    login       VARCHAR(26)  NOT NULL UNIQUE,
    password    VARCHAR(48)  NOT NULL UNIQUE,
    email       VARCHAR(64)  NOT NULL UNIQUE,
    name        VARCHAR(64)  NOT NULL,
    description VARCHAR(300) NULL
);

CREATE TABLE IF NOT EXISTS games
(
    id                  INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name                VARCHAR(60)  NOT NULL,
    price               DECIMAL      NOT NULL CHECK (price >= 0),
    id_developers       INT          NOT NULL REFERENCES dev_pub_account (id) ON DELETE CASCADE,
    id_publishers       INT          NOT NULL REFERENCES dev_pub_account (id) ON DELETE CASCADE,
    id_categories       INT          NULL REFERENCES games_categories (id) ON DELETE SET NULL,
    system_requirements VARCHAR(300) NOT NULL,
    age_limit           VARCHAR(3)   NOT NULL,
    description         VARCHAR(300) NULL,
    release_date        DATE         NOT NULL,
    version             VARCHAR(500) NOT NULL,
    rating              DECIMAL      NOT NULL CHECK (rating >= 0 AND rating <= 5)
);

CREATE TABLE IF NOT EXISTS users_libraries
(
    id       INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    id_games INT NOT NULL REFERENCES games (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS purchases_cheques
(
    id               INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    date_of_purchase DATE NOT NULL DEFAULT CURRENT_DATE
);

CREATE TABLE IF NOT EXISTS users_wish_lists
(
    id       INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    id_games INT NOT NULL REFERENCES games (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS users
(
    id                INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    login             VARCHAR(26)  NOT NULL UNIQUE,
    password          VARCHAR(48)  NOT NULL UNIQUE,
    email             VARCHAR(64)  NOT NULL UNIQUE,
    phone_number      VARCHAR(20)  NULL UNIQUE,
    nickname          VARCHAR(90)  NOT NULL,
    description       VARCHAR(300) NULL,
    registration_date DATE         NOT NULL,
    id_library        INT          NOT NULL REFERENCES users_libraries (id) ON DELETE NO ACTION,
    id_groups         INT          NULL REFERENCES users_groups (id) ON DELETE SET NULL,
    id_status         INT          NOT NULL REFERENCES users_statuses (id) ON DELETE NO ACTION,
    id_wishlist       INT          NULL REFERENCES users_wish_lists (id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS games_purchases
(
    id        INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    id_user   INT NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    id_game   INT NOT NULL REFERENCES games (id) ON DELETE CASCADE,
    id_cheque INT NOT NULL REFERENCES purchases_cheques (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS users_friends
(
    id        INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    id_user   INT NULL REFERENCES users (id) ON DELETE CASCADE,
    id_friend INT NULL REFERENCES users (id) ON DELETE CASCADE
);

INSERT INTO users_statuses(name)
VALUES ('В сети'),
       ('В игре'),
       ('Не на месте'),
       ('Неведимка'),
       ('Не в сети');