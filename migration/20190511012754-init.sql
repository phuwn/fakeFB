
-- +migrate Up

-- ----------------------------
-- Table structure for table users
-- ----------------------------
DROP TABLE IF EXISTS "users" CASCADE;
CREATE TABLE "users" (
    "id"  SERIAL NOT NULL,
    "email" VARCHAR(50),
    "password" VARCHAR(50)
);

-- +migrate Down
DROP TABLE IF EXISTS "users" CASCADE;