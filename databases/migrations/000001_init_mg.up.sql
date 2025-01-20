begin;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists public.roles (
  role_id uuid primary key default uuid_generate_v4(),
  name varchar(200) not null unique,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp,
  is_delete boolean default false
);

create table if not exists public.positions (
  position_id uuid primary key default uuid_generate_v4(),
  name varchar(200) not null unique,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp,
  is_delete boolean default false
);

create table if not exists public.status (
  status_id uuid primary key default uuid_generate_v4(),
  name varchar(200) not null unique,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp,
  is_delete boolean default false
);

Create table if not exists public.users (
  user_id uuid primary key default uuid_generate_v4(),
  username varchar(200) not null,
  name TEXT not null,
  password varchar(200) not null,
  email varchar(200) not null,
  nik varchar(200) not null unique,
  role_id uuid not null,
  position_id uuid not null,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp,
  is_delete boolean default false,

  constraint username_unique unique (username),
  constraint email_unique unique (email),
  constraint nik_unique unique (nik),
  constraint fk_role_id foreign key (role_id) references public.roles (role_id) match simple on update cascade on delete restrict,
  constraint fk_position_id foreign key (position_id) references public.positions (position_id) match simple on update cascade on delete restrict
);

create table if not exists public.leave_balances (
    leave_id serial primary key not null,
    leave_year varchar(200) not null,
    leave_balance int not null default 0,
    cuti_tahunan int not null default 0,
    cuti_izin int not null default 0,
    cuti_sakit int not null default 0,
    user_id uuid not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    is_delete boolean default false,

    constraint fk_user_id foreign key (user_id) references public.users (user_id) match simple on update cascade on delete restrict
);

create table if not exists public.leave_records (
  request_id uuid primary key default uuid_generate_v4(),
  request_on timestamp not null,
  from_date timestamp not null,
  to_date timestamp not null,
  return_date timestamp not null,
  amount int not null,
  reason varchar(200) not null,
  mobile varchar(200) not null,
  address varchar(200) not null,
  status_id uuid not null,
  leave_id int not null,
  user_id uuid not null,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp,
  is_delete boolean default false,

  constraint fk_status_id foreign key (status_id) references public.status (status_id) match simple on update cascade on delete restrict,
  constraint fk_user_id foreign key (user_id) references public.users (user_id) match simple on update cascade on delete restrict,
  constraint fk_leave_id foreign key (leave_id) references public.leave_balances (leave_id) match simple on update cascade on delete restrict
);

create table if not exists public.payroll_records (
  payroll_id uuid primary key default uuid_generate_v4(),
  payment_period varchar(200),
  payment_date timestamp not null,
  basic_salary int not null,
  allowance int,
  bpjs int,
  tax int,
  total_salary int not null,
  status_id uuid not null,
  user_id uuid not null,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp,
  is_delete boolean default false,

  constraint fk_user_id foreign key (user_id) references public.users (user_id) match simple on update cascade on delete restrict
);

alter table if exists public.users
  alter column password type varchar(200),
  alter column nik type varchar(200);

commit;