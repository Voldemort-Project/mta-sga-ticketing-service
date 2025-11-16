create table if not exists roles (
	created_at timestamp with time zone not null default now(),
	updated_at timestamp with time zone not null default now(),
	deleted_at timestamp with time zone null,
	id text not null,
	name text not null,
	description text null,
	constraint pk_roles primary key (id)
);

create table if not exists organizations (
	created_at timestamp with time zone not null default now(),
	updated_at timestamp with time zone not null default now(),
	deleted_at timestamp with time zone null,
	id text not null,
	name text not null,
	address text null,
	constraint pk_organizations primary key (id)
);

create table if not exists divisions (
	created_at timestamp with time zone not null default now(),
	updated_at timestamp with time zone not null default now(),
	deleted_at timestamp with time zone null,
	id text not null,
	name text not null,
	description text null,
	organization_id text,
	constraint pk_divisions primary key (id),
	constraint fk_divisions_organizations foreign key (organization_id) references organizations (id)
		on delete set null
		on update set null
);

create table if not exists users (
	created_at timestamp with time zone not null default now(),
	updated_at timestamp with time zone not null default now(),
	deleted_at timestamp with time zone null,
	id text not null,
	name text not null,
	email text null,
	phone_number text null,
	id_card_number text null,
	organization_id text,
	division_id text,
	constraint pk_users primary key (id),
	constraint fk_users_organizations foreign key (organization_id) references organizations (id)
		on delete set null
		on update set null,
	constraint fk_users_divisions foreign key (division_id) references divisions (id)
		on delete set null
		on update set null
);

create table if not exists user_roles (
	created_at timestamp with time zone not null default now(),
	updated_at timestamp with time zone not null default now(),
	deleted_at timestamp with time zone null,
	id text not null,
	user_id text not null,
	role_id text not null,
	constraint pk_user_roles primary key (id),
	constraint fk_user_roles_users foreign key (user_id) references users (id)
		on delete cascade
		on update cascade,
	constraint fk_user_roles_roles foreign key (role_id) references roles (id)
		on delete cascade
		on update cascade,
	constraint uk_user_roles_user_id_role_id unique (user_id, role_id)
);

create table if not exists rooms (
	created_at timestamp with time zone not null default now(),
	updated_at timestamp with time zone not null default now(),
	deleted_at timestamp with time zone null,
	id text not null,
	name text not null,
	description text null,
	is_available boolean not null default true,
	constraint pk_rooms primary key (id)
);

create table if not exists checkins (
	created_at timestamp with time zone not null default now(),
	updated_at timestamp with time zone not null default now(),
	deleted_at timestamp with time zone null,
	id text not null,
	org_id text,
	user_id text,
	room_id text,
	check_in_time timestamp with time zone not null,
	check_out_time timestamp with time zone null,
	guest_name text not null,
	guest_email text null,
	guest_phone text null,
	guest_id_card text null,
	constraint pk_checkins primary key (id),
	constraint fk_checkins_organizations foreign key (org_id) references organizations (id)
		on delete set null
		on update set null,
	constraint fk_checkins_users foreign key (user_id) references users (id)
		on delete set null
		on update set null,
	constraint fk_checkins_rooms foreign key (room_id) references rooms (id)
		on delete set null
		on update set null
);

create index if not exists idx_roles_name on roles (name) where deleted_at is null;
create index if not exists idx_organizations_name on organizations (name) where deleted_at is null;
create index if not exists idx_divisions_name on divisions (name) where deleted_at is null;
create index if not exists idx_users_email on users (email) where deleted_at is null;
create index if not exists idx_users_phone_number on users (phone_number) where deleted_at is null;
create index if not exists idx_users_id_card_number on users (id_card_number) where deleted_at is null;
create index if not exists idx_users_organization_id on users (organization_id) where deleted_at is null;
create index if not exists idx_users_division_id on users (division_id) where deleted_at is null;
create index if not exists idx_rooms_name on rooms (name) where deleted_at is null;
create index if not exists idx_checkins_guest_name on checkins (guest_name) where deleted_at is null;
create index if not exists idx_checkins_guest_email on checkins (guest_email) where deleted_at is null;
create index if not exists idx_checkins_guest_phone on checkins (guest_phone) where deleted_at is null;
create index if not exists idx_checkins_guest_id_card on checkins (guest_id_card) where deleted_at is null;
create index if not exists idx_checkins_org_id on checkins (org_id) where deleted_at is null;
create index if not exists idx_checkins_user_id on checkins (user_id) where deleted_at is null;
create index if not exists idx_checkins_room_id on checkins (room_id) where deleted_at is null;
create index if not exists idx_checkins_check_in_time on checkins (check_in_time) where deleted_at is null;
create index if not exists idx_checkins_check_out_time on checkins (check_out_time) where deleted_at is null;