
SET FOREIGN_KEY_CHECKS = 0;

DROP DATABASE IF EXISTS `dolphin`;
CREATE DATABASE `dolphin`;
USE `dolphin`;

# password not necessary
# trigger insert after insert self group
CREATE TABLE `user` (
    `id`  varchar(30) NOT NULL,
    `name` VARCHAR(10) NOT NULL,
    `gender` ENUM('F', 'M'),
# not exposed to clients
    `self_group_id` BIGINT(20),
    PRIMARY KEY (`id`),
    FOREIGN KEY (`self_group_id`) REFERENCES `group`(`id`) ON DELETE SET NULL ON UPDATE CASCADE
) DEFAULT CHARSET=utf8;


# every user starts with his own group
# but the self group relationship is not included here
CREATE TABLE `user_group` (
    `user_id` VARCHAR(30) NOT NULL,
    `group_id` BIGINT(20) NOT NULL,
    PRIMARY KEY (`user_id`, `group_id`),
    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`group_id`) REFERENCES `group`(`id`) ON DELETE CASCADE
) DEFAULT CHARSET=utf8;

# when insert check whether the same user
CREATE TABLE `group` (
    `id` BIGINT(20) AUTO_INCREMENT,
    `name` VARCHAR(10) DEFAULT "",
# also used for self group
    `creator_id` VARCHAR(30) NOT NULL,
    `type` ENUM('GROUP', 'INDIVIDUAL') DEFAULT 'GROUP',
    PRIMARY KEY (`id`),
    UNIQUE (`name`, `creator_id`),
    FOREIGN KEY (`creator_id`) REFERENCES `user` (`id`) ON DELETE CASCADE
) DEFAULT CHARSET=utf8;

# need review
# contents would be stored in other db
# assert(user publisher_id in group group_id)
# assert(start_date <= end_date)
# assert(leader_id in task worker)
CREATE TABLE `task` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `group_id` BIGINT(20) NOT NULL,
    `name` VARCHAR(20) DEFAULT "",
    `publisher_id` VARCHAR(30) NOT NULL,
    # only for group work, can be null
    `leader_id` VARCHAR(30),
    `start_date` DATE DEFAULT NULL,
    `end_date` DATE DEFAULT NULL,
    # if readonly, only the publisher can revise the task
    `readonly` BOOL DEFAULT FALSE NOT NULL,
    `type` ENUM('GROUP', 'INDIVIDUAL') DEFAULT 'GROUP',
    `description` TEXT,
    CHECK (end_date >= start_date),
    PRIMARY KEY (`id`),
    UNIQUE INDEX (`group_id`, `name`),
    FOREIGN KEY (`publisher_id`, `group_id`) REFERENCES `user_group`(`user_id`, `group_id`) ON DELETE CASCADE,
    FOREIGN KEY (`group_id`) references `group`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`leader_id`) references `user`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`publisher_id`) references `user`(`id`) ON DELETE CASCADE
) DEFAULT CHARSET=utf8;

CREATE TABLE `task_user` (
    `task_id` BIGINT(20) NOT NULL,
    `user_id` VARCHAR(30) NOT NULL,
    `done` BOOL DEFAULT FALSE NOT NULL,
    `done_time` TIMESTAMP DEFAULT NULL,
    PRIMARY KEY (`task_id`, `user_id`),
    FOREIGN KEY (`task_id`) REFERENCES `task`(`id`) ON DELETE CASCADE
);


SET FOREIGN_KEY_CHECKS = 1;