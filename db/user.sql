/*
 Navicat Premium Data Transfer

 Source Server         : postgres
 Source Server Type    : PostgreSQL
 Source Server Version : 100000
 Source Host           : localhost:32769
 Source Catalog        : messeinfor_knowledge
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 100000
 File Encoding         : 65001

 Date: 15/01/2018 19:02:07
*/


-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "user";
CREATE TABLE "user" (
  "id" varchar(36) COLLATE "pg_catalog"."default" NOT NULL DEFAULT NULL,
  "username" varchar(20) COLLATE "pg_catalog"."default" DEFAULT NULL,
  "password" varchar(20) COLLATE "pg_catalog"."default" DEFAULT NULL,
  "created_at" timestamp(6) DEFAULT NULL,
  "updated_at" timestamp(6) DEFAULT NULL,
  "deleted_at" timestamp(6) DEFAULT NULL
)
;
ALTER TABLE "user" OWNER TO "messeinfor";

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO "user" VALUES ('a8828fd9-a9b6-4709-bedd-90cc8370b389', 'voson', 'voson@2017', '2017-12-15 23:10:21', '2017-12-15 23:10:23', NULL);
COMMIT;

-- ----------------------------
-- Primary Key structure for table user
-- ----------------------------
ALTER TABLE "user" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("id");
