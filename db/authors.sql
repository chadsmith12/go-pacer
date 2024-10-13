-- name: GetAuthor :one
select * from authors
where id = $1 limit 1;

-- name: ListAuthors :many
select * from authors
order by name;

-- name: CreateAuthor :one
INSERT INTO authors (
    name, bio
) VALUES ( $1, $2 )
returning *;

-- name: UpdateAuthor :exec
UPDATE authors 
    SET name = $2,
    bio = $3
    WHERE id = $1;

-- name: DeleteAuthor :exec
DELETE FROM authors
    WHERE id = $1;
