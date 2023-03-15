-- name: CreateVerifyEmail :one
INSERT INTO verify_emails (
    username, email, secret_code
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetVerifyEmail :one
SELECT * FROM verify_emails
WHERE 
    id = $1 AND secret_code = $2
LIMIT 1;

-- name: VerifyEmail :one
UPDATE verify_emails 
SET
    is_used = $3
WHERE 
    id = $1 AND secret_code = $2
RETURNING *;