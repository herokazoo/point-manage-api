CREATE TABLE IF NOT EXISTS `user_points` (
  `id`         BIGINT UNSIGNED           NOT NULL AUTO_INCREMENT COMMENT 'ユーザーポイントの識別子',
  `user_id`    BIGINT UNSIGNED           NOT NULL                COMMENT 'ユーザーの識別子',
  `total`      INT    UNSIGNED DEFAULT 0 NOT NULL                COMMENT 'ポイント残高',
  `created_at` DATETIME(6)               NOT NULL                COMMENT 'レコード作成日時',
  `updated_at` DATETIME(6)               NOT NULL                COMMENT 'レコード更新日時',

  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_user_id` (`user_id`) USING BTREE
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ポイント';

CREATE TABLE IF NOT EXISTS `user_point_histories` (
  `id`             BIGINT  UNSIGNED                        NOT NULL AUTO_INCREMENT COMMENT 'ユーザーポイントの操作履歴の識別子',
  `user_point_id`  BIGINT  UNSIGNED                        NOT NULL                COMMENT 'ユーザーポイントの識別子',
  `operation`      TINYINT UNSIGNED                        NOT NULL                COMMENT 'ポイント操作の識別子',
  `operation_date` DATE             DEFAULT (CURRENT_DATE) NOT NULL                COMMENT 'ポイント操作日',
  `expiry_date`    DATE                                                            COMMENT 'ポイント失効日',
  `amount`         INT                                     NOT NULL                COMMENT 'ポイント数',
  `remaining`      INT     UNSIGNED DEFAULT 0              NOT NULL                COMMENT 'ポイント残数',
  `created_at`     DATETIME(6)                             NOT NULL                COMMENT 'レコード作成日時',
  `updated_at`     DATETIME(6)                             NOT NULL                COMMENT 'レコード更新日時',

  PRIMARY KEY (`id`),
  CONSTRAINT `fk_user_point_id`
    FOREIGN KEY (`user_point_id`) REFERENCES `user_points` (`id`)
      ON DELETE RESTRICT ON UPDATE RESTRICT
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ポイント履歴';
