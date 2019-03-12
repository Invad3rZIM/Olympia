CREATE TABLE Users (
user_id int PRIMARY KEY NOT NULL ,
first_name varchar(30) NOT NULL,
last_name varchar(30) NOT NULL,
email varchar(30) NOT NULL,
phone varchar(30) NOT NULL,
password varchar NOT NULL,
salt int NOT NULL,
user_type varchar(30) NOT NULL);

CREATE TABLE Arenas (
arena_id int PRIMARY KEY NOT NULL ,
arena_name varchar(50) NOT NULL,
arena_capacity int NOT NULL,
IsOutdoor BOOLEAN DEFAULT TRUE,
IsAquatic BOOLEAN DEFAULT FALSE,

HasField BOOLEAN DEFAULT TRUE,
HasTrack BOOLEAN DEFAULT FALSE,

HasPool BOOLEAN DEFAULT FALSE);

CREATE TABLE Events (
event_id serial PRIMARY KEY NOT NULL ,

event_min_start_day int NOT NULL,
event_max_start_day int NOT NULL,
event_min_start_month int NOT NULL,
event_max_start_month int NOT NULL,
event_min_start_time int NOT NULL,
event_max_start_time int NOT NULL,
event_duration int NOT NULL,
arena_id int NOT NULL, 

security_required int NOT NULL,
security_obtained int NOT NULL,

tickets_sold int NOT NULL,
ticket_price float8 NOT NULL,
ticket_discount_price float8 NOT NULL

schedule_id int NOT NULL,
manually_scheduled BOOLEAN DEFAULT FALSE,
scheduled_time int NOT NULL,
scheduled_day int NOT NULL,
scheduled_month int NOT NULL);

CREATE TABLE Tickets(
ticket_id int PRIMARY KEY NOT NULL,
event_id int NOT NULL,
user_id int NOT NULL,
PricePaid float8 NOT NULL);

CREATE TABLE AthletesSchedules (
athlete_id int NOT NULL,
obligation_type int NOT NULL,
day int NOT NULL,
start_time int NOT NULL,
month int NOT NULL,
duration int NOT NULL);

CREATE TABLE Security(
guard_id int PRIMARY KEY NOT NULL,
obligation_arena_id int NOT NULL,
priority int NOT NULL,
assigned_officer_id int NOT NULL);
