CREATE TABLE cart (
	id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
	owner_user_id INT NOT NULL,
	product_id INT NOT NULL,
	product_count INT NOT NULL
)ENGINE=InnoDB;
ALTER TABLE cart add index index1(owner_user_id, product_id);