DROP DATABASE IF EXISTS orders;
CREATE DATABASE orders;
USE orders;
CREATE TABLE `orders` (
  `id` varchar(36) NOT NULL,
  `price` decimal(10,2) NOT NULL,
  `tax` decimal(10,2) NOT NULL,
  `final_price` decimal(10,2) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
