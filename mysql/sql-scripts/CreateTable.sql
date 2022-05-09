SET NAMES utf8mb4;
SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL';

DROP SCHEMA IF EXISTS movie_db;
CREATE SCHEMA movie_db;
USE movie_db;

--
-- Table structure for table `actors`
--

CREATE TABLE `actors` (
  `actor_id` SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(150) NOT NULL,
  PRIMARY KEY (`actor_id`),
  KEY idx_actor_last_name (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Table structure for table `movies`
--

CREATE TABLE `movies` (
  `movie_id` SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` varchar(128) NOT NULL,
  `release_date` DATE DEFAULT NULL,
  `description` TEXT DEFAULT NULL,
  PRIMARY KEY (`movie_id`),
  KEY idx_title(`title`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Table structure for table `movie_casts`
--

CREATE TABLE `movie_casts` (
  `actor_id` SMALLINT UNSIGNED NOT NULL,
  `movie_id` SMALLINT UNSIGNED NOT NULL,
  PRIMARY KEY (`actor_id`,`movie_id`),
  KEY idx_fk_movie_id (`movie_id`),
  CONSTRAINT fk_movie_actor_actor FOREIGN KEY (`actor_id`) REFERENCES actors (`actor_id`) ON DELETE RESTRICT ON UPDATE CASCADE,
  CONSTRAINT fk_movie_actor_movie FOREIGN KEY (movie_id) REFERENCES movies (movie_id) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Table structure for table `ratings`
--

CREATE TABLE `ratings` (
  `rating_id` SMALLINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `movie_id` SMALLINT UNSIGNED NOT NULL,
  `rating` TINYINT UNSIGNED NOT NULL,
  `num_of_ratings` SMALLINT UNSIGNED NOT NULL,
  PRIMARY KEY (`rating_id`,`movie_id`),
  KEY idx_fk_movie_id (`movie_id`),
  CONSTRAINT fk_movie_rating_movie FOREIGN KEY (movie_id) REFERENCES movies (movie_id) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;