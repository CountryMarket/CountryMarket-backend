CREATE TABLE product_order (
	id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
	owner_user_id INT NOT NULL,
	owner_shop_user_id INT NOT NULL,
	now_status INT NOT NULL,
	total_price DOUBLE NOT NULL DEFAULT 0,
	transportation_price DOUBLE NOT NULL DEFAULT 0,
	discount_price DOUBLE NOT NULL DEFAULT 0,
	product_and_count VARCHAR(511) NOT NULL,
	person_name VARCHAR(255) NOT NULL,
	person_phone_number VARCHAR(255) NOT NULL,
	person_address VARCHAR(255) NOT NULL,
	pay_time DATETIME(3),
	verify_time DATETIME(3),
	tracking_number VARCHAR(255) NOT NULL,
	message VARCHAR(255) NOT NULL
);
ALTER TABLE product_order add index index1(owner_user_id);