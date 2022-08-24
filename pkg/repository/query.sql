-- name: ListBooks :many
SELECT * FROM books
ORDER BY id;

-- name: CreateBook :one
INSERT INTO books (
    title, description, author_name, price
) VALUES (
             $1, $2, $3, $4
         )
RETURNING *;

-- name: GetBookById :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: UpdateBookById :exec
UPDATE books
SET title = $2,
    description = $3,
    author_name = $4,
    price = $5
WHERE id = $1;

-- name: DeleteBookById :exec
DELETE FROM books
WHERE id = $1;