DROP TABLE IF EXISTS `nearby_area_codes`;
CREATE TABLE `nearby_area_codes`
(
    `nearby_area_code_id` INTEGER PRIMARY KEY,
    `area_code`           int(3) NOT NULL,
    `nearby_area_code`    int(3) NOT NULL
);