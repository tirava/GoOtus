create table events (
	id UUID primary key,
	createdat timestamp not null,
	updatedat timestamp,
	deletedat timestamp,
	occursat timestamp,
	subject text not null,
	body text,
	duration bigint,
	location text,
	userid UUID not null
)