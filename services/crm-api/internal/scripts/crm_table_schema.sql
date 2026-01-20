create database crm_pltfrm_db;

create table crm_pltfrm_mgmt_users(
id varchar(255) not null primary key,
username varchar(255) not null unique,
password varchar(255) not null,
credt_at varchar(255) not null,
updtd_at varchar(255)
);

create table crm_pltfrm_mgmt_contact_category(
id varchar(255) not null primary key,
contact_category varchar(255) not null unique
);

create table crm_pltfrm_mgmt_contacts(
id varchar(255) not null primary key,
name varchar(255) not null,
phone varchar(255) not null,
email varchar(255) not null,
crm_id varchar(255) not null,
contact_category_id varchar(255) not null
);

create table crm_pltfrm_mgmt_email_templates(
id varchar(255) not null primary key,
template_name varchar(255) not null unique,
subject varchar(255) not null,
body varchar(1000) not null,
crm_id varchar(255) not null
)

