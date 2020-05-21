
/* This script sets up tables in the database and
 * should be run by the courselect user like:
 *
 * sudo mysql -u courselect -p < schema.sql
 * 
 * Attention: running this script would remove all
 * existing data in your courselect database. */

drop table if exists student;
drop table if exists faculty;
drop table if exists major;
drop table if exists course;
drop table if exists select_request;
drop table if exists select_result;
drop table if exists select_time;

create table major (
  id   int primary key auto_increment,
  name text
);

create table student (
  id       char(10) primary key,
  name     text,
  password text,
  grade    char(4),
  class    int,
  major    int,
  email    text,
  phone    text,
  foreign key (major) references major(id)
);

create table faculty (
  id       char(10) primary key,
  name     text,
  password text,
  position enum('Prof', 'Associate Prof', 'Assistant Prof')
);

create table course (
  id      int primary key auto_increment,
  name    text,
  teacher char(10),
  credit  int,
  sdate   date,
  edate   date,
  day     enum('Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'),
  stime   time,
  etime   time,
  valid   enum('Passed', 'Rejected'),
  foreign key (teacher) references faculty(id)
);

create table select_request (
  id      int primary key auto_increment,
  course  int,
  student char(10),
  time    datetime,
  foreign key (course) references course(id),
  foreign key (student) references student(id)
);

create table select_result (
  id      int primary key auto_increment,
  course  int,
  student char(10),
  grade   int,
  foreign key (course) references course(id),
  foreign key (student) references student(id)
);

create table select_time (
  id      int primary key auto_increment,
  major   int,
  stime   datetime,
  etime   datetime,
  foreign key (major) references major(id)
);