CREATE TABLE `room` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
	`roomInfo` varchar(1000) NOT NULL DEFAULT '',
	`roomId`  bigint(20) NOT NULL DEFAULT '0',
	`roomStatus`  int NOT NULL DEFAULT '0',
	`managerId`  bigint(20) NOT NULL DEFAULT '0',
	`teacherId`  bigint(20) NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
