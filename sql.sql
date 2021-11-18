create table users(
id int primary key AUTO_INCREMENT,
username varchar(100) not null UNIQUE,
password varchar(100) not null,
email varchar(100)
)