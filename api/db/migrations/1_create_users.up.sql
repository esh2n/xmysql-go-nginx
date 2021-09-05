CREATE TABLE `users` (
  `id` varchar(26) NOT NULL,
  `name` varchar(50) NOT NULL,
  `token` varchar(500) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime,
  PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8mb4;
