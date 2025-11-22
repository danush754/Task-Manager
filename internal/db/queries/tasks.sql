-- name: CreateTask :one
insert into tbl_tasks(title, description, task_status, deadline, owner_id)
values($1,$2,$3,$4,$5)
RETURNING *;

-- name: ListTaskByUser :many
select * from tbl_tasks where owner_id = $1;

-- name: GetTaskById :one
select * from tbl_tasks where id = $1 limit 1;

-- name: ExtendDeadlineById :one 
update tbl_tasks set deadline = $1, updated_at = NOW() where id = $2
RETURNING *;

-- name: ChangeTaskStatusById :one
update tbl_tasks set task_status = $1, updated_at = NOW() where id = $2
RETURNING *;

-- name: DeleteTask :exec
delete from tbl_tasks where id = $1;
