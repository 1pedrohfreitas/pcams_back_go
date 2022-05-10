-- DROP SCHEMA pcam;

CREATE SCHEMA pcam AUTHORIZATION pedrohfreitas;

-- DROP SEQUENCE pcam.cam_access_permission_id_seq;

CREATE SEQUENCE pcam.cam_access_permission_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1;
-- DROP SEQUENCE pcam.cams_id_seq;

CREATE SEQUENCE pcam.cams_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1;
-- DROP SEQUENCE pcam.clients_id_seq;

CREATE SEQUENCE pcam.clients_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1;
-- DROP SEQUENCE pcam.installers_id_seq;

CREATE SEQUENCE pcam.installers_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1;
-- DROP SEQUENCE pcam.licences_id_seq;

CREATE SEQUENCE pcam.licences_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1;
-- DROP SEQUENCE pcam.users_id_seq;

CREATE SEQUENCE pcam.users_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START 1;-- pcam.cams definition

-- Drop table

-- DROP TABLE pcam.cams;

CREATE TABLE pcam.cams (
	id bigserial NOT NULL,
	devicename varchar(100) NOT NULL,
	alias varchar(100) NOT NULL,
	streamtype varchar(100) NOT NULL,
	urlcamstream varchar(100) NOT NULL,
	status int2 NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	CONSTRAINT cams_pkey PRIMARY KEY (id)
);


-- pcam.clients definition

-- Drop table

-- DROP TABLE pcam.clients;

CREATE TABLE pcam.clients (
	id bigserial NOT NULL,
	fullname varchar(100) NOT NULL,
	alias varchar(100) NOT NULL,
	status int2 NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	codclient varchar NOT NULL,
	CONSTRAINT clients_pkey PRIMARY KEY (id)
);


-- pcam.installers definition

-- Drop table

-- DROP TABLE pcam.installers;

CREATE TABLE pcam.installers (
	id bigserial NOT NULL,
	fullname varchar(100) NOT NULL,
	alias varchar(100) NOT NULL,
	username varchar(30) NOT NULL,
	usertype varchar(30) NOT NULL,
	status int2 NOT NULL,
	"password" varchar(20) NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	CONSTRAINT installers_pkey PRIMARY KEY (id)
);
CREATE INDEX installers_username_idx ON installers USING btree (username, usertype);


-- pcam.licences definition

-- Drop table

-- DROP TABLE pcam.licences;

CREATE TABLE pcam.licences (
	id bigserial NOT NULL,
	uniccode varchar(100) NOT NULL,
	clientid int4 NOT NULL,
	installerid int4 NOT NULL,
	licencetype varchar(20) NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	CONSTRAINT licences_pkey PRIMARY KEY (id)
);


-- pcam.users definition

-- Drop table

-- DROP TABLE pcam.users;

CREATE TABLE pcam.users (
	id bigserial NOT NULL,
	fullname varchar(100) NOT NULL,
	alias varchar(100) NOT NULL,
	username varchar(30) NOT NULL,
	usertype varchar(30) NOT NULL,
	status int2 NOT NULL,
	"password" varchar(20) NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);
CREATE INDEX users_username_idx ON users USING btree (username, usertype);


-- pcam.cam_access_permission definition

-- Drop table

-- DROP TABLE pcam.cam_access_permission;

CREATE TABLE pcam.cam_access_permission (
	id bigserial NOT NULL,
	camid int4 NOT NULL,
	alias varchar(100) NOT NULL,
	startpermissiondate timestamptz NULL,
	stoppermissiondate timestamptz NULL,
	durationpermitions int4 NULL,
	CONSTRAINT cam_access_permission_pkey PRIMARY KEY (id),
	CONSTRAINT cam_access_permission_fk FOREIGN KEY (camid) REFERENCES pcam.cams(id) ON DELETE CASCADE
);
CREATE INDEX cam_access_permission_camid_idx ON cam_access_permission USING btree (camid);


-- pcam.cams_clients definition

-- Drop table

-- DROP TABLE pcam.cams_clients;

CREATE TABLE pcam.cams_clients (
	camid int4 NOT NULL,
	clientid int4 NOT NULL,
	CONSTRAINT cams_clients_fk FOREIGN KEY (camid) REFERENCES pcam.cams(id) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT cams_clients_fk_1 FOREIGN KEY (clientid) REFERENCES pcam.clients(id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX cams_clients_clientid_idx ON cams_clients USING btree (clientid);


-- pcam.cams_installers definition

-- Drop table

-- DROP TABLE pcam.cams_installers;

CREATE TABLE pcam.cams_installers (
	camid int4 NOT NULL,
	installerid int4 NOT NULL,
	CONSTRAINT cams_installers_fk FOREIGN KEY (camid) REFERENCES pcam.cams(id) ON DELETE CASCADE,
	CONSTRAINT cams_installers_fk_1 FOREIGN KEY (installerid) REFERENCES pcam.installers(id) ON DELETE CASCADE
);
CREATE INDEX cams_installers_installerid_idx ON cams_installers USING btree (installerid);


-- pcam.clients_installers definition

-- Drop table

-- DROP TABLE pcam.clients_installers;

CREATE TABLE pcam.clients_installers (
	installerid int4 NOT NULL,
	clientid int4 NOT NULL,
	CONSTRAINT clients_installers_fk FOREIGN KEY (clientid) REFERENCES pcam.clients(id) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT clients_installers_fk_1 FOREIGN KEY (installerid) REFERENCES pcam.installers(id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX clients_installers_installerid_idx ON clients_installers USING btree (installerid);


-- pcam.clients_users definition

-- Drop table

-- DROP TABLE pcam.clients_users;

CREATE TABLE pcam.clients_users (
	userid int4 NOT NULL,
	clientid int4 NOT NULL,
	CONSTRAINT clients_users_fk FOREIGN KEY (userid) REFERENCES pcam.users(id) ON DELETE CASCADE ON UPDATE CASCADE,
	CONSTRAINT clients_users_fk_1 FOREIGN KEY (clientid) REFERENCES pcam.clients(id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX clients_users_clientid_idx ON clients_users USING btree (clientid);
CREATE INDEX clients_users_userid_idx ON clients_users USING btree (userid);
