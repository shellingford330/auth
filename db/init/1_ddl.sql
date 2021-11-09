-- -----------------------------------------------------
-- Schema shellingford
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `shellingford` DEFAULT CHARACTER SET utf8mb4 ;
USE `shellingford` ;

SET CHARSET utf8mb4;

-- -----------------------------------------------------
-- Table `shellingford`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `shellingford`.`users` (
  `id` VARCHAR(26) NOT NULL COMMENT 'ユーザID',
  `name` VARCHAR(64) NOT NULL COMMENT 'ユーザ名',
  `email` VARCHAR(64) NOT NULL COMMENT 'メールアドレス',
  `password` VARCHAR(128) COMMENT 'パスワードハッシュ',
  `image` VARCHAR(256) COMMENT 'アイコンURL',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = 'ユーザ';

-- -----------------------------------------------------
-- Table `shellingford`.`accounts`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `shellingford`.`accounts` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `provider_id` VARCHAR(64) NOT NULL COMMENT 'プロバイダID',
  `provider_type` VARCHAR(64) NOT NULL COMMENT 'プロバイダタイプ',
  `provider_account_id` VARCHAR(128) COMMENT 'プロバイダアカウントID',
  `user_id` VARCHAR(26) NOT NULL COMMENT 'ユーザID',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_account_user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `shellingford`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
COMMENT = 'アカウント';

-- -----------------------------------------------------
-- Table `shellingford`.`sessions`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `shellingford`.`sessions` (
  `session_token` VARCHAR(64) NOT NULL COMMENT 'セッショントークン(ulid)',
  `access_token` VARCHAR(64) NOT NULL COMMENT 'アクセストークン(ulid)',
  `expires` TIMESTAMP NOT NULL COMMENT '期限',
  `user_id` VARCHAR(26) NOT NULL COMMENT 'ユーザID',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`session_token`),
  CONSTRAINT `fk_session_user_id`
    FOREIGN KEY (`user_id`)
    REFERENCES `shellingford`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
COMMENT = 'セッション';
