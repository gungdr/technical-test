CREATE TABLE `user` (
  ID INT AUTO_INCREMENT PRIMARY KEY,
  UserName VARCHAR(10) NOT NULL,
  Parent INT NOT NULL DEFAULT 0
);
