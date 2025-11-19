create type status as ENUM ('completed', 'inprogress', 'hold', 'not-started');

create table if not exists tbl_tasks (
id serial Primary key,
title text not null,
description text,
task_status status not null, 
deadline TIMESTAMP WITHOUT TIME ZONE, 
owner_id int not null references tbl_users(id) on delete cascade,
created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);

create index idc_tasks_owner_id on tbl_tasks(owner_id);