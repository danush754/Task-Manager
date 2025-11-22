-- name: CreateUser :one
insert into tbl_users(name, email, password_hash) 
values ($1, $2, $3)
RETURNING *;

-- name: GetUserById :one
select * from tbl_users where id = $1 limit 1;

-- name: GetUserByEmail :one
select * from tbl_users where email = $1 limit 1;

-- name: GetUsersList :many
select * from tbl_users order by id desc;