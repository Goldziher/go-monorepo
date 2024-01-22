-- name: UpsertUser :one
INSERT INTO "user_account" (
    full_name,
    email,
    phone_number,
    profile_picture_url,
    username,
    hashed_password
)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (email)
DO UPDATE SET full_name = $1,
phone_number = $3,
profile_picture_url = $4,
username = $5
RETURNING id;

-- name: GetUserById :one
SELECT
    full_name,
    email,
    phone_number,
    profile_picture_url,
    username
FROM "user_account"
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT
    full_name,
    email,
    phone_number,
    profile_picture_url,
    username
FROM "user_account"
WHERE email = $1
LIMIT 1;
