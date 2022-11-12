alter table "user"
  drop column refresh_token;

drop table "user_verification";

create table "refresh_token"
(
  id         BIGSERIAL               not null
    constraint refresh_token_pk
      primary key,
  user_id    bigint                 not null
    references "user",
  created_at timestamp default now() not null,
  source     varchar(64)             not null,
  token      uuid                    not null
)
