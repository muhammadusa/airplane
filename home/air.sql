create table air(
    airplane_id serial primary key not null,
    terminal smallint not null,
    flight varchar(50) not null,
    flytime time not null,
    togo varchar(50) not null,
    gate smallint not null,
    landtime time not null,
    remark varchar(50) not null
);

insert into air (terminal, flight, flytime, togo, gate, landtime, remark) values (3, 'BK345', '08:00:00', 'London', 12, '12:01:00', 'onetime');
insert into air (terminal, flight, flytime, togo, gate, landtime, remark) values (1, 'HI264', '09:20:00', 'Samarqand', 3, '12:34:00', 'chek-in');
insert into air (terminal, flight, flytime, togo, gate, landtime, remark) values (2, 'UY362', '10:26:00','Turkish', 5, '12:25:00', 'delated');
insert into air (terminal, flight, flytime, togo, gate, landtime, remark) values (1, 'ER264', '18:30:00','Namangan', 3, '12:25:00', 'then');
insert into air (terminal, flight, flytime, togo, gate, landtime, remark) values (2, 'TU923', '20:50:00','Qozogiston', 1, '12:25:00', 'then');


insert into air (terminal, flight, flytime, togo, gate, landtime, remark) values (3, 'ED111', '10:00:00', 'London', 9, '12:01:00', 'chec-in');
insert into air (terminal, flight, flytime, togo, gate, landtime, remark) values (3, 'PO163', '09:30:00', 'Shvetsariya', 5, '12:01:00', 'chek-in');
insert into air (terminal, flight, flytime, togo, gate, landtime, remark) values (1, 'WA737', '23:16:00','Azarbayjan', 3, '12:25:00', 'then');
insert into air (terminal, flight, flytime, togo, gate, landtime, remark) values (1, 'TB639', '14:35:00','Buxoro', 6, '12:25:00', 'then');
insert into air (terminal, flight, flytime, togo, gate, landtime, remark) values (2, 'DD923', '22:53:00','Qozogiston', 2, '12:25:00', 'deleted');

select
	terminal as id
;


select
	terminal,
	airplane_id ,
	flight,
	to_char(flytime, 'HH24:MI'),	
	togo,
	gate,
	to_char(landtime, 'HH24:MI'),
	remark
from 
	air;