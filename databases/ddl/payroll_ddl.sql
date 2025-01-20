create table if not exists public.payroll_records (
  payroll_id uuid primary key default uuid_generate_v4(),
  payment_period varchar(200) COLLATE,
  payment_date timestamp not null,
  basic_salary int not null,
  allowance int,
  bpjs int,
  user_id uuid not null,
  tax int,
  total_salary int not null,
  status_id uuid not null,
  user_id uuid not null,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp,
  is_delete boolean default false

  constraint fk_user_id foreign key (user_id) references public.users (user_id) match simple on update cascade on delete restrict
);