CREATE DATABASE IF NOT EXISTS `gomeeting`  DEFAULT CHARACTER SET utf8mb4;

USE `gomeeting`;

--
-- Table structure for table `meeting`
--

CREATE TABLE IF NOT EXISTS `meeting` (
  `room_id` int(10) unsigned NOT NULL,
  `start_time` int(10) unsigned NOT NULL,
  `end_time` int(10) unsigned NOT NULL,
  `maker` int(10) unsigned NOT NULL,
  `memo` varchar(100) NOT NULL,
  `make_date` int(8) unsigned NOT NULL,
  UNIQUE KEY `room_id` (`room_id`,`start_time`,`end_time`,`make_date`),
  KEY `make_date` (`make_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Table structure for table `org`
--

CREATE TABLE IF NOT EXISTS `org` (
  `id` int(10) unsigned NOT NULL,
  `name` varchar(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


--
-- Table structure for table `room`
--

CREATE TABLE IF NOT EXISTS `room` (
  `id` int(11) NOT NULL,
  `name` varchar(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Table structure for table `user`
--

CREATE TABLE IF NOT EXISTS `user` (
  `id` int(10) unsigned NOT NULL,
  `username` varchar(20) NOT NULL COMMENT 'login name',
  `pw` varchar(100) NOT NULL,
  `level` int(10) unsigned NOT NULL,
  `org_id` int(10) unsigned NOT NULL,
  `name` varchar(20) NOT NULL COMMENT 'display name',
  `email` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Table structure for table `notification`
--

CREATE TABLE IF NOT EXISTS `notification` (
  `message` varchar(3072) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


--
-- Insert root/root account into `user`
--

INSERT INTO `user`
  SELECT 0 , 'root', md5( 'root' ) , 0, 0, 'root', 'root@example.com'
  FROM dual
  WHERE NOT EXISTS (
    SELECT id
    FROM `user`
    WHERE id = 0
  );

--
-- Insert 0/root org into `org`
--

INSERT INTO `org`
  SELECT 0 , 'root'
  FROM dual
  WHERE NOT EXISTS (
    SELECT id
    FROM `org`
    WHERE id = 0
  );