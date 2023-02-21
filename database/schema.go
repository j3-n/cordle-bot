package database

var schema = `
	create table users (
		id primary key int not null,
		name text not null,
		wins int not null,
		losses int not null,
		draws int not null,
		games int not null,
		elo int not null,
		level int not null
	);
`
