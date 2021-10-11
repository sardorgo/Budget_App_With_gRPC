-- The name of db is grpc_hw1

create extension "uuid-ossp";

create table users (
    user_id uuid not null primary key,
    first_name varchar(32) not null,
    last_name varchar(32) not null,
    total_money int default 0
);

create table cash (
    cash_id uuid not null primary key,
    cash_amount int not null,
    summary varchar(300) default null,
    created_at timestamp with time zone default current_timestamp,
    user_id uuid not null references users (user_id)
);

create table cost (
    cost_id uuid not null primary key,
    cost_amount int not null,
    summary varchar(300) default null,
    created_at timestamp with time zone default current_timestamp,
    user_id uuid not null references users (user_id)

);

--- Triggers

create or replace function cash_controller () returns trigger language plpgsql as
    $$
        begin
            update users
            set total_money = total_money + (select cash_amount from cash where cash_id = new.cash_id)
            where user_id = new.user_id;

            return new;
        end;
    $$
;

create trigger cash_controller_tg after insert on cash
for each row execute procedure cash_controller();


create or replace function cost_controller () returns trigger language plpgsql as
    $$
        begin
            update users
            set total_money = total_money - (select cost_amount from cost where cost_id = new.cost_id)
            where user_id = new.user_id;

            return new;
        end;
    $$
;

create trigger cost_controller_tg after insert on cost
for each row execute procedure cost_controller();

-- Mock data 

insert into users (user_id, first_name, last_name) values (uuid_generate_v4(), 'John', 'Doe');

insert into cash (cash_id, cash_amount, user_id) values (uuid_generate_v4(), 100000, '26380728-d8f2-4a5a-b3f9-ad543076bae9');
insert into cost (cost_id, cost_amount, user_id) values (uuid_generate_v4(), 50000, '26380728-d8f2-4a5a-b3f9-ad543076bae9');



update users
set total_money = total_money + 100000
where user_id = '39cad5bb-5f0f-4b2d-a238-1c244f6572bf';