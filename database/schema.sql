SET FOREIGN_KEY_CHECKS = 0;

DROP DATABASE IF EXISTS `test`;
CREATE DATABASE `test`;
USE `test`;

# password not necessary
CREATE TABLE `user` (
    `id`  BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(10) NOT NULL,
    `password` VARCHAR(20) NOT NULL,
    PRIMARY KEY (`id`)
);

# every user starts with his own group
CREATE TABLE `user_group` (
    `user_id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `group_id` BIGINT(20) NOT NULL,
    PRIMARY KEY (`user_id`, `group_id`),
    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`group_id`) REFERENCES `group`(`id`) ON DELETE CASCADE,
    INDEX(`user_id`),
    INDEX(`group_id`)
);

CREATE TABLE `group` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(10) DEFAULT "",
    PRIMARY KEY (`id`, `name`),
    INDEX (`name`)
);

# contentS would be stored in other db
# assert(user publisher_id in group group_id)
# assert(start_date <= end_date)
CREATE TABLE `task` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `group_id` BIGINT(20) NOT NULL,
    `name` VARCHAR(20) DEFAULT "",
    `publisher_id` BIGINT(20) NOT NULL,
    `start_date` BIGINT(20) DEFAULT CURRENT_DATE,
    `end_date` BIGINT(20) DEFAULT CURRENT_DATE,
    `type` VARCHAR(20) DEFAULT "",
    `description` TEXT DEFAULT "",
    CHECK (end_date <= start_date),
    PRIMARY KEY (`id`),
    UNIQUE INDEX (`group_id`, `name`),
    INDEX (`group_id`),
    INDEX (`publisher_id`),
    INDEX (`start_date`, `end_date`),
    FOREIGN KEY (`publisher_id`, `group_id`) references `user_group`(`user_id`, `group_id`) ON DELETE CASCADE
);

SET FOREIGN_KEY_CHECKS = 1;