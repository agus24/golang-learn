ALTER TABLE promotions
ADD COLUMN status TINYINT(1) UNSIGNED DEFAULT 1 NOT NULL COMMENT '0: inactive, 1: active' AFTER discount_amount;
