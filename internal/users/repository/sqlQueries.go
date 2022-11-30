package repository

const (
	queryCreateUser = `
insert into users.users(name, password) values($1, $2)
`
	queryUpdatePassword = `
update users.users
set password = $1
where id = (select id from users.users where name = $2) and password = $3;
`
	queryAuthorizationInsertSessionKey = `
insert into users.sessions(session_key, authed, user_id)
values ($1, $2, (select id from users.users where name = $3));
`
	queryUpdateSessionKey = `
update users.sessions
set authed = false
where user_id = (select id from users.users where name = $1)
  and authed = true;
`
	queryCheckUserParams = `
select case when name = $1 and password = $2 then true else false end as authed
from users.users
where name = $1;
`
)
