-- url table --
CREATE TABLE IF NOT EXISTS `links` (
	`short_url` TEXT NOT NULL PRIMARY KEY,
	`target_url` TEXT NOT NULL
);