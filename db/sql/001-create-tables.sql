CREATE TABLE task_mysql.task
(
  id int(10) AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(35) NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
) DEFAULT CHARACTER SET utf8 COLLATE utf8_bin;