CREATE SCHEMA users;
CREATE SCHEMA codes;

    create table codes.code_test
    (
        id     serial
            constraint code_test_pk
                primary key,
        amount double precision not null,
        date   timestamp          not null
    );

alter table codes.code_test
    owner to postgres;

create unique index code_test_id_uindex
    on codes.code_test (id);


create table codes.codes
(
    id   serial
        constraint codes_pk
            primary key,
    name text not null,
    code text not null
);

alter table codes.codes
    owner to postgres;

create unique index codes_code_uindex
    on codes.codes (code);

create unique index codes_id_uindex
    on codes.codes (id);

create unique index codes_name_uindex
    on codes.codes (name);



create unique index table_name_code_uindex
    on codes.codes (code);

create unique index table_name_id_uindex
    on codes.codes (id);

create unique index table_name_name_uindex
    on codes.codes (name);


create table users.users
(
    id       serial
        constraint users_pk
            primary key,
    name     text not null,
    password text not null
);

alter table users.users
    owner to postgres;

create unique index users_id_uindex
    on users.users (id);

create unique index users_name_uindex
    on users.users (name);



create table users.sessions
(
    id          serial
        constraint sessions_pk
            primary key,
    session_key text                 not null,
    authed      boolean default true not null,
    user_id     integer              not null
        constraint sessions_users_id_fk
            references users.users
);

alter table users.sessions
    owner to postgres;

create unique index sessions_id_uindex
    on users.sessions (id);

create unique index sessions_session_key_uindex
    on users.sessions (session_key);

insert into codes.codes (name, code) values('Austria ECU-EUR exchange rates','AME_A_AUT_1_0_99_0_XNE');
insert into codes.codes (name, code) values('Belgium ECU-EUR exchange rates','AME_A_BEL_1_0_99_0_XNE');
insert into codes.codes (name, code) values('FR. Germany ECU-EUR exchange rates','AME_A_DEU_1_0_99_0_XNE');
insert into codes.codes (name, code) values('West Germany ECU-EUR exchange rates','AME_A_D_W_1_0_99_0_XNE');
insert into codes.codes (name, code) values('Spain ECU-EUR exchange rates','AME_A_ESP_1_0_99_0_XNE');
insert into codes.codes (name, code) values('Finland ECU-EUR exchange rates ','AME_A_FIN_1_0_99_0_XNE');
insert into codes.codes (name, code) values('France ECU-EUR exchange rates','AME_A_FRA_1_0_99_0_XNE');
insert into codes.codes (name, code) values('Greece ECU-EUR exchange rates','AME_A_GRC_1_0_99_0_XNE');
insert into codes.codes (name, code) values('Ireland ECU-EUR exchange rates','AME_A_IRL_1_0_99_0_XNE');
insert into codes.codes (name, code) values('Italy ECU-EUR exchange rates','AME_A_ITA_1_0_99_0_XNE');

insert into users.users (name, password) values('admin','admin');