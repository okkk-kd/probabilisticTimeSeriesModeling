package repository

const (
	queryGetCodesList = `
select id, name, code from codes.codes
`
	queryCreateCodesDataTable = `
create table if not exists codes.%s
(
    id     serial
        constraint %s_pk
            primary key,
    amount double precision not null,
    date   timestamp          not null
);

alter table codes.%s
    owner to postgres;

create unique index %s_id_uindex
    on codes.%s (id);
`
	queryInsertTableIntoTableList = `
insert into codes.codes(name, code) values ($1, $2)
`
	queryInsertDataIntoCodesDataTable = `
insert into codes.%s (amount, date) values($1, $2)
`
	queryGetCodeDataByID = `
select amount, date from codes.%s where id = $1
`
	queryDeleteCodeDataByID = `
delete from codes.%s where id = $1
`
	queryUpdateCodeDataByID = `
update codes.%s set amount = $1, date = $2 where id = $3
`
	queryAddCodeData = `
insert into codes.%s (amount, date)
values ($1, $2);
`
	querySelectDataByCode = `
select id, date, amount from codes.%s
`
)
