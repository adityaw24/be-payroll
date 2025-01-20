CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

Create table if not exists public.users (
  user_id uuid primary key default uuid_generate_v4(),
  username varchar(200) COLLATE not null unique,
  name TEXT COLLATE not null,
  password varchar(200) not null,
  email varchar(200) COLLATE not null unique,
  nik varchar(200) not null unique,
  role_id uuid not null,
  position_id uuid not null,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp,
  is_delete boolean default false

  constraint fk_role_id foreign key (role_id) references public.roles (role_id) match simple on update cascade on delete restrict,
  constraint fk_position_id foreign key (position_id) references public.positions (position_id) match simple on update cascade on delete restrict
);