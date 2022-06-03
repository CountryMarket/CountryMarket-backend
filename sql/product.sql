CREATE TABLE product (
	id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
	owner_user_id INT NOT NULL,
	price DOUBLE NOT NULL DEFAULT 0,
	title VARCHAR(255),
	description VARCHAR(255),
	picture_number INT NOT NULL DEFAULT 0,
	
	created_at DATETIME(3),
	updated_at DATETIME(3),
	deleted_at DATETIME(3)
)ENGINE=InnoDB;
ALTER TABLE product add index index1(owner_user);