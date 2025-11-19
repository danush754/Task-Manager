create table if not exists tbl_users (id serial Primary Key,
name varchar(100) not null, 
email text unique not null, 
password_hash text not null, 
created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW(),
updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()  
);