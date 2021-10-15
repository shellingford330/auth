-- -----------------------------------------------------
-- Schema shellingford
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `shellingford` DEFAULT CHARACTER SET utf8mb4 ;
USE `shellingford` ;

SET CHARSET utf8mb4;

-- -----------------------------------------------------
-- Table `shellingford`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `docs_api`.`users` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ユーザID',
  `name` VARCHAR(64) NOT NULL COMMENT 'ユーザ名',
  `email` VARCHAR(64) NOT NULL COMMENT 'メールアドレス',
  `password` VARCHAR(128) NOT NULL COMMENT 'パスワードハッシュ',
  `image` VARCHAR(256) COMMENT 'アイコンURL'
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = 'ユーザ';
