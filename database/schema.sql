// TODO: ADD TRIGGERS
SET FOREIGN_KEY_CHECKS = 0;

DROP DATABASE IF EXISTS `test`;
CREATE DATABASE `test`;
USE `test`;

# password not necessary
# trigger insert after insert self group
CREATE TABLE `user` (
    `id`  BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(10) NOT NULL,
    `password` VARCHAR(20) NOT NULL,
    `self_group_id` DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8;

# every user starts with his own group
# trigger insert update user_cnt of group
CREATE TABLE `user_group` (
    `user_id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `group_id` BIGINT(20) NOT NULL,
    PRIMARY KEY (`user_id`, `group_id`),
    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`group_id`) REFERENCES `group`(`id`) ON DELETE CASCADE,
    INDEX(`user_id`),
    INDEX(`group_id`)
) DEFAULT CHARSET=utf8;


CREATE TABLE `group` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(10) DEFAULT "",
    PRIMARY KEY (`id`, `name`),
    INDEX (`name`)
) DEFAULT CHARSET=utf8;

# contents would be stored in other db
# assert(user publisher_id in group group_id)
# assert(start_date <= end_date)
CREATE TABLE `task` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `group_id` BIGINT(20) NOT NULL,
    `name` VARCHAR(20) DEFAULT "",
    `publisher_id` BIGINT(20) NOT NULL,
    `start_date` DATE DEFAULT NULL,
    `end_date` DATE DEFAULT NULL,
    `type` ENUM('GROUP', 'INDIVIDUAL') DEFAULT 'GROUP',
    `description` TEXT,
    CHECK (end_date >= start_date),
    PRIMARY KEY (`id`),
    UNIQUE INDEX (`group_id`, `name`),
    INDEX (`group_id`),
    INDEX (`publisher_id`),
    INDEX (`start_date`, `end_date`),
    INDEX (`done`),
    FOREIGN KEY (`publisher_id`, `group_id`) references `user_group`(`user_id`, `group_id`) ON DELETE CASCADE
    CHECK (`done_cnt` >= 0)
    CHECK (`user_cnt` >= 1)
) DEFAULT CHARSET=utf8;

CREATE TABLE `task_user` (
    `task_id` BIGINT(20) NOT NULL,
    `user_id` BIGINT(20) NOT NULL,
    `done` BOOL DEFAULT FALSE NOT NULL,
    `done_time` TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (`task_id`, `user_id`),
    FOREIGN KEY (`task_id`) REFERENCES `task`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE,
    INDEX (`user_id`),
    INDEX (`task_id`)
);

SET FOREIGN_KEY_CHECKS = 1;