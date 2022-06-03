CREATE TABLE address (
	id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
	owner_user_id INT NOT NULL,
	name VARCHAR(255) NOT NULL,
	phone_number VARCHAR(255) NOT NULL,
	address VARCHAR(255) NOT NULL
)ENGINE=InnoDB;
ALTER TABLE cart add index index1(owner_user_id);