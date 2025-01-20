create table if not exists public.status (
  status_id uuid primary key default uuid_generate_v4(),
  name varchar(200) COLLATE not null unique,
  created_at timestamp default current_timestamp,
  updated_at timestamp default current_timestamp,
  is_delete boolean default false
);
