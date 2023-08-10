CREATE TABLE "user"
(
    id                  uuid PRIMARY KEY      DEFAULT gen_random_uuid(),
    full_name           text         NOT NULL,
    email               text UNIQUE  NOT NULL,
    phone_number        varchar(128) NOT NULL,
    profile_picture_url text         NOT NULL,
    username            text         NOT NULL,
    hashed_password     text         NOT NULL,
    created_at          TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);
