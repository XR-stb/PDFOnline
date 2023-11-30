CREATE TABLE IF NOT EXISTS `pdfs` (
    `id` INT(11) NOT NULL,
    `title` VARCHAR(255) NOT NULL,
    `description` TEXT,
    `url` VARCHAR(255) NOT NULL,
    `cover_url` VARCHAR(255) NOT NULL,
    `created_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;

INSERT INTO `pdfs`(`id`, `title`, `description`, `url`, `cover_url`, `created_at`) VALUES ('9', 'Title 1', 'Description 1', '/static/pdf/深度学习-花书.pdf', 'http://localhost:8080/static/cover/深度学习-花书.jpg', '2023-11-26 14:23:06.000');
INSERT INTO `pdfs`(`id`, `title`, `description`, `url`, `cover_url`, `created_at`) VALUES ('2', 'Title 2', 'Description 2', '/static/pdf/深度学习-花书.pdf', 'http://localhost:8080/static/cover/深度学习-花书.jpg', '2023-11-26 14:23:06.000');
INSERT INTO `pdfs`(`id`, `title`, `description`, `url`, `cover_url`, `created_at`) VALUES ('3', 'Title 3', 'Description 3', '/static/pdf/深度学习-花书.pdf', 'http://localhost:8080/static/cover/深度学习-花书.jpg', '2023-11-26 14:23:06.000');
INSERT INTO `pdfs`(`id`, `title`, `description`, `url`, `cover_url`, `created_at`) VALUES ('4', 'Title 4', 'Description 4', '/static/pdf/深度学习-花书.pdf', 'http://localhost:8080/static/cover/深度学习-花书.jpg', '2023-11-26 14:23:06.000');
INSERT INTO `pdfs`(`id`, `title`, `description`, `url`, `cover_url`, `created_at`) VALUES ('5', 'Title 5', 'Description 5', '/static/pdf/深度学习-花书.pdf', 'http://localhost:8080/static/cover/深度学习-花书.jpg', '2023-11-26 14:23:06.000');
INSERT INTO `pdfs`(`id`, `title`, `description`, `url`, `cover_url`, `created_at`) VALUES ('6', 'Title 6', 'Description 6', '/static/pdf/深度学习-花书.pdf', 'http://localhost:8080/static/cover/深度学习-花书.jpg', '2023-11-26 14:23:06.000');
INSERT INTO `pdfs`(`id`, `title`, `description`, `url`, `cover_url`, `created_at`) VALUES ('7', 'Title 7', 'Description 7', '/static/pdf/深度学习-花书.pdf', 'http://localhost:8080/static/cover/深度学习-花书.jpg', '2023-11-26 14:23:06.000');
INSERT INTO `pdfs`(`id`, `title`, `description`, `url`, `cover_url`, `created_at`) VALUES ('8', 'Title 8', 'Description 8', '/static/pdf/深度学习-花书.pdf', 'http://localhost:8080/static/cover/深度学习-花书.jpg', '2023-11-26 14:23:06.000');
