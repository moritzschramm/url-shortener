-- url table --
CREATE TABLE IF NOT EXISTS `links` (
	`id` INTEGER PRIMARY KEY AUTOINCREMENT,
	`short_url` TEXT NOT NULL,
	`target_url` TEXT NOT NULL
);