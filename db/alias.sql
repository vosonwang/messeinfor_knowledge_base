/*
 Navicat Premium Data Transfer

 Source Server         : pg 10.1
 Source Server Type    : PostgreSQL
 Source Server Version : 100000
 Source Host           : localhost:32769
 Source Catalog        : messeinfor_knowledge
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 100000
 File Encoding         : 65001

 Date: 25/01/2018 17:16:45
*/


-- ----------------------------
-- Table structure for alias
-- ----------------------------
DROP TABLE IF EXISTS "alias";
CREATE TABLE "alias" (
  "id" varchar(36) COLLATE "pg_catalog"."default" NOT NULL DEFAULT NULL,
  "name" varchar(36) COLLATE "pg_catalog"."default" DEFAULT NULL,
  "created_at" timestamp(6) DEFAULT NULL,
  "updated_at" timestamp(6) DEFAULT NULL,
  "deleted_at" timestamp(6) DEFAULT NULL,
  "description" varchar(50) COLLATE "pg_catalog"."default" DEFAULT NULL,
  "parent_id" varchar(36) COLLATE "pg_catalog"."default" DEFAULT NULL,
  "number" int4 DEFAULT nextval('alias_number_seq'::regclass)
)
;
ALTER TABLE "alias" OWNER TO "messeinfor";

-- ----------------------------
-- Indexes structure for table alias
-- ----------------------------
CREATE UNIQUE INDEX CONCURRENTLY "alias_id_uindex" ON "alias" USING btree (
  "id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table alias
-- ----------------------------
ALTER TABLE "alias" ADD CONSTRAINT "alias_pkey" PRIMARY KEY ("id");
