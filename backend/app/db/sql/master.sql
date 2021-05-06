SET foreign_key_checks=0;

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_name` VARCHAR(255) NOT NULL,
    `auth_id` VARCHAR(191) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE `auth_id_idx` (`auth_id`),
    UNIQUE `email_idx` (`email`),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARACTER SET utf8mb4;


DROP TABLE IF EXISTS `todos`;

CREATE TABLE `todos` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL,
    `title` VARCHAR(255) NOT NULL,
    `content` LONGTEXT NOT NULL,
    `check` TINYINT(1) NOT NULL DEFAULT false,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX `user_id_idx` (`user_id`),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARACTER SET utf8mb4;

SET foreign_key_checks=1;
