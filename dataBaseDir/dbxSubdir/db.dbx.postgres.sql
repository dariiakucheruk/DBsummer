-- AUTOGENERATED BY storj.io/dbx
-- DO NOT EDIT
CREATE TABLE groups (
	cipher text NOT NULL,
	groupname text NOT NULL,
	educationalyear text NOT NULL,
	semester text NOT NULL,
	course text NOT NULL,
	subject integer NOT NULL,
	PRIMARY KEY ( cipher )
);
CREATE TABLE students (
	student_cipher text NOT NULL,
	firstname text NOT NULL,
	last_name text NOT NULL,
	middle_name text NOT NULL,
	record_book_number text NOT NULL,
	PRIMARY KEY ( student_cipher )
);
CREATE TABLE subjects (
	subjectid integer NOT NULL,
	subjectname text NOT NULL,
	educationallevel text NOT NULL,
	faculty text NOT NULL,
	PRIMARY KEY ( subjectid )
);
CREATE TABLE tableNews (
	id_t text NOT NULL,
	PRIMARY KEY ( id_t )
);
