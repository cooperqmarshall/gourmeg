create table if not exists chef (
    created_at timestamp with time zone default now(),
    id serial primary key,
    username text
);

create table if not exists recipe (
    created_at timestamp with time zone default now(),
    id serial primary key,
    chef_id integer,
    url text not null,
    name text,
    ingredients text[],
    instructions text[],
    domain text,
    image_url text,
    thumbnail_url text
);

create table if not exists list (
    created_at timestamp with time zone default now(),
    id serial primary key,
    chef_id integer,
    name text not null
);

create table if not exists link (
    parent_id integer not null,
    child_id integer not null,
    child_type character varying(8) not null
);
