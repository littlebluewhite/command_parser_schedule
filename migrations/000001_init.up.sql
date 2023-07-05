CREATE TABLE `time_template` (
                                 `id` int PRIMARY KEY AUTO_INCREMENT,
                                 `name` varchar(255) UNIQUE NOT NULL,
                                 `time_data_id` int UNIQUE NOT NULL,
                                 `updated_at` datetime,
                                 `created_at` datetime DEFAULT (now())
);

CREATE TABLE `time_data` (
                             `id` int PRIMARY KEY AUTO_INCREMENT,
                             `repeat_type` ENUM ('daily', 'weekly', 'monthly'),
                             `start_date` date NOT NULL,
                             `end_date` date,
                             `start_time` time NOT NULL,
                             `end_time` time NOT NULL,
                             `interval_seconds` int,
                             `condition_type` ENUM ('monthly_day', 'weekly_day', 'weekly_first', 'weekly_second', 'weekly_third', 'weekly_fourth'),
                             `t_condition` json
);

CREATE TABLE `schedule` (
                            `id` int PRIMARY KEY AUTO_INCREMENT,
                            `name` varchar(255) UNIQUE NOT NULL,
                            `description` varchar(255),
                            `time_data_id` int UNIQUE NOT NULL,
                            `task_id` int,
                            `enabled` boolean DEFAULT false,
                            `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `created_at` datetime DEFAULT (now())
);

CREATE TABLE `command_template` (
                                    `id` int PRIMARY KEY AUTO_INCREMENT,
                                    `name` varchar(255) UNIQUE NOT NULL,
                                    `protocol` ENUM ('http', 'websocket', 'mqtt', 'redis_topic') NOT NULL,
                                    `description` varchar(255),
                                    `host` varchar(255) NOT NULL,
                                    `port` varchar(255) NOT NULL,
                                    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                    `created_at` datetime DEFAULT (now())
);

CREATE TABLE `https_command` (
                                 `id` int PRIMARY KEY AUTO_INCREMENT,
                                 `command_id` int UNIQUE,
                                 `method` ENUM ('GET', 'POST', 'PATCH', 'PUT', 'DELETE') NOT NULL,
                                 `url` varchar(255) NOT NULL,
                                 `authorization_type` ENUM ('basic', 'token'),
                                 `params` json,
                                 `header` json,
                                 `body_type` ENUM ('text', 'html', 'xml', 'form_data', 'x_www_form_urlencoded', 'json'),
                                 `body` json
);

CREATE TABLE `header_template` (
                                   `id` int PRIMARY KEY AUTO_INCREMENT,
                                   `name` varchar(255) UNIQUE NOT NULL,
                                   `data` json
);

CREATE TABLE `websocket_command` (
                                     `id` int PRIMARY KEY AUTO_INCREMENT,
                                     `command_id` int UNIQUE,
                                     `url` varchar(255) NOT NULL,
                                     `header` json,
                                     `message` varchar(255)
);

CREATE TABLE `mqtt_command` (
                                `id` int PRIMARY KEY AUTO_INCREMENT,
                                `command_id` int UNIQUE,
                                `topic` varchar(255) NOT NULL,
                                `header` json,
                                `message` json,
                                `type` ENUM ('publish', 'subscribe') NOT NULL
);

CREATE TABLE `redis_command` (
                                 `id` int PRIMARY KEY AUTO_INCREMENT,
                                 `command_id` int UNIQUE,
                                 `password` varchar(255),
                                 `db` int DEFAULT 0,
                                 `topic` varchar(255),
                                 `message` json,
                                 `type` ENUM ('publish', 'subscribe') NOT NULL
);

CREATE TABLE `monitor` (
                           `id` int PRIMARY KEY AUTO_INCREMENT,
                           `column` ENUM ('status', 'data') NOT NULL,
                           `timeout` int NOT NULL,
                           `interval` int,
                           `command_id` int UNIQUE NOT NULL
);

CREATE TABLE `m_condition` (
                               `id` int PRIMARY KEY AUTO_INCREMENT,
                               `order` int,
                               `calculate_type` ENUM ('=', '!=', '<', '>', '<=', '>=', 'include', 'exclude'),
                               `pre_logic_type` ENUM ('and', 'or'),
                               `value` varchar(255),
                               `search_rule` varchar(255) COMMENT 'ex: person.item.[]array.name',
                               `monitor_id` int
);

CREATE TABLE `task_template` (
                                 `id` int PRIMARY KEY AUTO_INCREMENT,
                                 `name` varchar(255),
                                 `variable` json,
                                 `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                 `created_at` datetime DEFAULT (now())
);

CREATE TABLE `task_template_stage` (
                                       `id` int PRIMARY KEY AUTO_INCREMENT,
                                       `task_template_id` int,
                                       `stage_id` int
);

CREATE TABLE `task_stage` (
                              `id` int PRIMARY KEY AUTO_INCREMENT,
                              `name` varchar(255),
                              `stage_number` int,
                              `mode` ENUM ('monitor', 'execute'),
                              `command_template_id` int,
                              `tag` json
);

ALTER TABLE `time_template` ADD FOREIGN KEY (`time_data_id`) REFERENCES `time_data` (`id`);

ALTER TABLE `schedule` ADD FOREIGN KEY (`time_data_id`) REFERENCES `time_data` (`id`);

ALTER TABLE `schedule` ADD FOREIGN KEY (`task_id`) REFERENCES `task_template` (`id`);

ALTER TABLE `https_command` ADD FOREIGN KEY (`command_id`) REFERENCES `command_template` (`id`) ON DELETE CASCADE;

ALTER TABLE `websocket_command` ADD FOREIGN KEY (`command_id`) REFERENCES `command_template` (`id`) ON DELETE CASCADE;

ALTER TABLE `mqtt_command` ADD FOREIGN KEY (`command_id`) REFERENCES `command_template` (`id`) ON DELETE CASCADE;

ALTER TABLE `redis_command` ADD FOREIGN KEY (`command_id`) REFERENCES `command_template` (`id`) ON DELETE CASCADE;

ALTER TABLE `monitor` ADD FOREIGN KEY (`command_id`) REFERENCES `command_template` (`id`) ON DELETE CASCADE;

ALTER TABLE `m_condition` ADD FOREIGN KEY (`monitor_id`) REFERENCES `monitor` (`id`) ON DELETE CASCADE;

ALTER TABLE `task_template_stage` ADD FOREIGN KEY (`task_template_id`) REFERENCES `task_template` (`id`);

ALTER TABLE `task_template_stage` ADD FOREIGN KEY (`stage_id`) REFERENCES `task_stage` (`id`);

ALTER TABLE `task_stage` ADD FOREIGN KEY (`command_template_id`) REFERENCES `command_template` (`id`);