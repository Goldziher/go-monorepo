-- name: UpsertUser :one
INSERT INTO "user" (display_name, email, phone_number, photo_url, provider_id)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (email)
    DO UPDATE SET display_name = $2,
                  phone_number = $4,
                  photo_url    = $5,
                  provider_id  = $6
RETURNING "user".id;

-- name: GetUserById :one
SELECT *
FROM "user"
WHERE id = $1
LIMIT 1;
