-- drop
DROP TABLE IF EXISTS `test_table`;

-- create
CREATE TABLE IF NOT EXISTS `test_table`(
  `id` INT(20) AUTO_INCREMENT,
  `testname` VARCHAR(20) NOT NULL,
  PRIMARY KEY(`id`)
);


CREATE TABLE IF NOT EXISTS `jokes` (
  `id` INT(20) AUTO_INCREMENT,
  `joke` VARCHAR(100) NOT NULL,
  PRIMARY KEY(`id`)
);
