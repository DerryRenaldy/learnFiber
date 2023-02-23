CREATE DATABASE `customers` /*!40100 DEFAULT CHARACTER SET utf8 */ /*!80016 DEFAULT ENCRYPTION='N' */;

CREATE TABLE IF NOT EXISTS customers.customers_status (
	id TINYINT AUTO_INCREMENT NOT NULL,
	code TINYINT(1) UNSIGNED NOT NULL,
	description VARCHAR(12) NOT NULL,
	created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	CONSTRAINT customer_status_pk PRIMARY KEY(id),
	CONSTRAINT customers_status_UN UNIQUE KEY (code)
)
ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

CREATE TABLE IF NOT EXISTS customers.customers (
  id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  code CHAR(32) NOT NULL,
  status TINYINT(1) UNSIGNED DEFAULT 0 NOT NULL,
  created timestamp DEFAULT CURRENT_TIMESTAMP,
  updated timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT customers_PK PRIMARY KEY (id),
  CONSTRAINT customers_UN UNIQUE KEY (code),
  CONSTRAINT customers_status_FK FOREIGN KEY (status) REFERENCES customers.customers_status(code)

)
ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

CREATE TABLE IF NOT EXISTS customers.customers_merchants (
  id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  customerCode CHAR(32) NOT NULL,
  merchantCode CHAR(6) NOT NULL,
  created timestamp DEFAULT CURRENT_TIMESTAMP,
  updated timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT customers_merchants_PK PRIMARY KEY (id),
  CONSTRAINT customers_merchants_FK FOREIGN KEY (customerCode) REFERENCES customers.customers(code)
)
ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

CREATE TABLE IF NOT EXISTS customers.customers_emails (
  id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  customerCode CHAR(32) NOT NULL,
  email VARCHAR(100) NOT NULL,
  created timestamp DEFAULT CURRENT_TIMESTAMP,
  updated timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT customers_emails_PK PRIMARY KEY (id),
  CONSTRAINT customers_emails_FK FOREIGN KEY (customerCode) REFERENCES customers.customers(code)
)
ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

CREATE INDEX idx_status_code ON customers.customers_status (code);

CREATE TABLE IF NOT EXISTS customers.customers_phone_numbers (
  id BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  customerCode CHAR(32) NOT NULL,
  phoneNumber VARCHAR(20) NOT NULL,
  created timestamp DEFAULT CURRENT_TIMESTAMP,
  updated timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT customers_phone_numbers_PK PRIMARY KEY (id),
  CONSTRAINT customers_phone_numbers_FK FOREIGN KEY (customerCode) REFERENCES customers.customers(code)
)
ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

INSERT INTO customers.customers_status (code, `description`) VALUES (0, 'INACTIVE');
INSERT INTO customers.customers_status (code, `description`) VALUES (1, 'ACTIVE');
INSERT INTO customers.customers_status (code, `description`) VALUES (2, 'DISABLED');
INSERT INTO customers.customers_status (code, `description`) VALUES (3, 'BLOCKED');

CREATE UNIQUE INDEX customers_status_IDX ON customers.customers_status(`code`);

ALTER TABLE customers.customers_emails DROP FOREIGN KEY customers_emails_FK;
ALTER TABLE customers.customers_merchants DROP FOREIGN KEY customers_merchants_FK;
ALTER TABLE customers.customers_phone_numbers DROP FOREIGN KEY customers_phone_numbers_FK;

ALTER TABLE customers.customers MODIFY COLUMN code char(16) NOT NULL;
ALTER TABLE customers.customers_merchants MODIFY COLUMN customerCode char(16) NOT NULL;
ALTER TABLE customers.customers_emails MODIFY COLUMN customerCode char(16) NOT NULL;
ALTER TABLE customers.customers_phone_numbers MODIFY COLUMN customerCode char(16) NOT NULL;

ALTER TABLE customers.customers_phone_numbers ADD FOREIGN KEY customers_phone_numbers_FK (customerCode) REFERENCES customers.customers(code);
ALTER TABLE customers.customers_emails ADD FOREIGN KEY customers_emails_FK (customerCode) REFERENCES customers.customers(code);
ALTER TABLE customers.customers_merchants ADD FOREIGN KEY customers_merchants_FK (customerCode) REFERENCES customers.customers(code);

INSERT INTO customers.customers(code, `status`, created, updated) VALUES('AYOCON-11S3R691', 1, current_timestamp, current_timestamp);
INSERT INTO customers.customers_merchants(customerCode, merchantCode, created, updated) VALUES('AYOCON-11S3R691', 'AYOCON', current_timestamp, current_timestamp);
INSERT INTO customers.customers_emails(customerCode, email, created, updated) VALUES('AYOCON-11S3R691', 'anthonlotus9@gmail.com', current_timestamp, current_timestamp);
INSERT INTO customers.customers_phone_numbers(customerCode, phoneNumber, created, updated) VALUES('AYOCON-11S3R691', '6282240306021', current_timestamp, current_timestamp);