CREATE TABLE `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
	`userInfo` varchar(1000) NOT NULL DEFAULT '',
	`uid`  bigint(20) NOT NULL DEFAULT '0',
    `phone` varchar(45) NOT NULL DEFAULT '',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
