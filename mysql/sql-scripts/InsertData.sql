SET NAMES utf8mb4;
SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL';
SET @old_autocommit=@@autocommit;

USE movie_db;

--
-- Dumping data for table actors
--

SET AUTOCOMMIT=0;
INSERT INTO actors VALUES (1,'Penelope Guiness'),
(2,'Nick Wahlberg'),
(3,'Ed Chase'),
(4,'Jennifer Davis'),
(5,'Johnny Lollobrigida'),
(6,'Bette Nicholson'),
(7,'Grace Mostel'),
(8,'Mathhew Johansson'),
(9,'Joe Swank'),
(10,'Christian Gable');
COMMIT;

--
-- Dumping data for table movies
--

SET AUTOCOMMIT=0;
INSERT INTO movies VALUES (1,'ACADEMY DINOSAUR',2006,'A Epic Drama of a Feminist And a Mad Scientist who must Battle a Teacher in The Canadian Rockies'),
(2,'ADAPTATION HOLES',2006,'A Astounding Reflection of a Lumberjack And a Car who must Sink a Lumberjack in A Baloon Factory'),
(3,'ALABAMA DEVIL',2006,'A Thoughtful Panorama of a Database Administrator And a Mad Scientist who must Outgun a Mad Scientist in A Jet Boat'),
(4,'ALONE TRIP',2006,'A Fast-Paced Character Study of a Composer And a Dog who must Outgun a Boat in An Abandoned Fun House'),
(5,'AMADEUS HOLY',2006,'A Emotional Display of a Pioneer And a Technical Writer who must Battle a Man in A Baloon'),
(6,'ANACONDA CONFESSIONS',2006,'A Lacklusture Display of a Dentist And a Dentist who must Fight a Girl in Australia'),
(7,'ANGELS LIFE',2006,'A Thoughtful Display of a Woman And a Astronaut who must Battle a Robot in Berlin'),
(8,'ANONYMOUS HUMAN',2006,'A Amazing Reflection of a Database Administrator And a Astronaut who must Outrace a Database Administrator in A Shark Tank'),
(9,'ANTITRUST TOMATOES',2006,'A Fateful Yarn of a Womanizer And a Feminist who must Succumb a Database Administrator in Ancient India'),
(10,'ANYTHING SAVANNAH',2006,'A Epic Story of a Pastry Chef And a Woman who must Chase a Feminist in An Abandoned Fun House'),
(11,'APACHE DIVINE',2006,'A Awe-Inspiring Reflection of a Pastry Chef And a Teacher who must Overcome a Sumo Wrestler in A U-Boat'),
(12,'ARACHNOPHOBIA ROLLERCOASTER',2006,'A Action-Packed Reflection of a Pastry Chef And a Composer who must Discover a Mad Scientist in The First Manned Space Station'),
(13,'ARMY FLINTSTONES',2006,'A Boring Saga of a Database Administrator And a Womanizer who must Battle a Waitress in Nigeria'),
(14,'ARTIST COLDBLOODED',2006,'A Stunning Reflection of a Robot And a Moose who must Challenge a Woman in California'),
(15,'BABY HALL',2006,'A Boring Character Study of a A Shark And a Girl who must Outrace a Feminist in An Abandoned Mine Shaft'),
(16,'BANG KWAI',2006,'A Epic Drama of a Madman And a Cat who must Face a A Shark in An Abandoned Amusement Park'),
(17,'BANGER PINOCCHIO',2006,'A Awe-Inspiring Drama of a Car And a Pastry Chef who must Chase a Crocodile in The First Manned Space Station'),
(18,'BAREFOOT MANCHURIAN',2006,'A Intrepid Story of a Cat And a Student who must Vanquish a Girl in An Abandoned Amusement Park'),
(19,'BEAST HUNCHBACK',2006,'A Awe-Inspiring Epistle of a Student And a Squirrel who must Defeat a Boy in Ancient China'),
(20,'BIRCH ANTITRUST',2006,'A Fanciful Panorama of a Husband And a Pioneer who must Outgun a Dog in A Baloon'),
(21,'BONNIE HOLOCAUST',2006,'A Fast-Paced Story of a Crocodile And a Robot who must Find a Moose in Ancient Japan'),
(22,'CAMPUS REMEMBER',2006,'A Astounding Drama of a Crocodile And a Mad Cow who must Build a Robot in A Jet Boat'),
(23,'CHOCOLAT HARRY',2006,'A Action-Packed Epistle of a Dentist And a Moose who must Meet a Mad Cow in Ancient Japan'),
(24,'CLONES PINOCCHIO',2006,'A Amazing Drama of a Car And a Robot who must Pursue a Dentist in New Orleans'),
(25,'CROOKED FROGMEN',2006,'A Unbelieveable Drama of a Hunter And a Database Administrator who must Battle a Crocodile in An Abandoned Amusement Park');
COMMIT;

--
-- Dumping data for table movie_casts
--

SET AUTOCOMMIT=0;
INSERT INTO movie_casts VALUES(1,1),
(1,6),
(1,7),
(2,2),
(2,11),
(2,15),
(3,4),
(3,13),
(3,3),
(4,6),
(4,7),
(4,18),
(5,5),
(5,17),
(5,21),
(6,9),
(6,16),
(6,19),
(7,7),
(7,8),
(7,12),
(8,15),
(8,22),
(8,24),
(9,10),
(9,20),
(9,23),
(10,1),
(10,3),
(10,25);
COMMIT;

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
SET autocommit=@old_autocommit;