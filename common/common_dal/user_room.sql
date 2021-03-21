
CREATE TABLE `user_room` (
	`roomId`  bigint(20) NOT NULL DEFAULT '0',
	`userId`  bigint(20) NOT NULL DEFAULT '0',
	`status`  int(20) NOT NULL DEFAULT '0',
    PRIMARY KEY (`roomId`, `userId`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
