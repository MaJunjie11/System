CREATE TABLE `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `name` varchar(45) NOT NULL DEFAULT '',
    `age` int(11) NOT NULL DEFAULT '0',
    `sex` tinyint(3) NOT NULL DEFAULT '0',
    `phone` varchar(45) NOT NULL DEFAULT '',
	`passworld` varchar(300) NOT NULL DEFAULT '',
	`uid`  bigint(20) NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
