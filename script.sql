CREATE DATABASE IF NOT EXISTS `todo`;
USE `todo`;

CREATE TABLE IF NOT EXISTS `todo` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT NULL,
  `description` varchar(300) DEFAULT NULL,
  `done` int(11) DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `todo` (`id`, `title`, `description`, `done`) VALUES
	(1, 'Cabelo', 'cortar cabelo', 0),
	(2, 'Exercicio', 'ir treinar 19:30', 0),
	(3, 'Janta', 'fazer janta', 1),
	(4, 'Lanche', 'fazer lanche', 1),
	(5, 'Fumar', 'dar uns trago', 1),
	(6, 'Almoço', 'Fazer almoço', 0);