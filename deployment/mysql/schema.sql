create database cordle_test;
use cordle_test;

create table users (
	id     varchar(50)     not null primary key,
	wins   int default 0   not null,
	losses int default 0   not null,
	draws  int default 0   not null,
	elo    int default 1000 not null,
	constraint id
	    unique (id)
);

create database cordle;
use cordle;

create table users (
	id     varchar(50)     not null primary key,
	wins   int default 0   not null,
	losses int default 0   not null,
	draws  int default 0   not null,
	elo    int default 1000 not null,
	constraint id
	    unique (id)
);
