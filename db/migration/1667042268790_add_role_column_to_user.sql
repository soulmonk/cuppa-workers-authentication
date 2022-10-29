alter table "user"
add role varchar(32) default null;

update "user" set role='user';

alter table "user"
alter column role set not null;