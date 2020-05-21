
/* This script sets up the courselect database and
 * should be run by the root user like:
 *
 * sudo mysql -u root < setup.sql
 * 
 * Attention: running this script would remove all
 * existing data in your courselect database. */

-- Create the database;
drop database if exists courselect;
create database courselect;

-- Create the user;
drop user if exists courselect@localhost;
create user courselect@localhost identified by 'P@ssw0rd';
grant all on courselect.* to courselect@localhost;
