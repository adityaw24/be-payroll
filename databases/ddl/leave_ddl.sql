create table if not exists public.leave_records (
  request_id uuid primary key default uuid_generate_v4(),
  request_on timestamp not null,
  from_date timestamp not null,
  to_date timestamp not null,
  return_date timestamp not null,
  amount int not null,
  reason varchar(200) COLLATE not null,
  mobile varchar(200) COLLATE not null,
  address varchar(200) COLLATE not null,
  status_id uuid not null,
  leave_id int not null,
  user_id uuid not null,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp,
  is_delete boolean default false

  constraint fk_status_id foreign key (status_id) references public.status (status_id) match simple on update cascade on delete restrict,
  constraint fk_user_id foreign key (user_id) references public.users (user_id) match simple on update cascade on delete restrict,
  constraint fk_leave_id foreign key (leave_id) references public.leave_balances (leave_id) match simple on update cascade on delete restrict
);

create table if not exists public.leave_balances (
    leave_id serial primary key not null,
    leave_year varchar(200) COLLATE not null,
    leave_balance int not null default 0,
    cuti_tahunan int not null default 0,
    cuti_izin int not null default 0,
    cuti_sakit int not null default 0,
    user_id uuid not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    is_delete boolean default false

    constraint fk_user_id foreign key (user_id) references public.users (user_id) match simple on update cascade on delete restrict
);