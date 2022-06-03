CREATE TABLE user (
	id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
	openid VARCHAR(255) NOT NULL,
	permission INT NOT NULL,
	phone_number VARCHAR(255),
	nick_name VARCHAR(255),
	avatar_url VARCHAR(255),
)ENGINE=InnoDB;
alter table user add index index1(openid(4));