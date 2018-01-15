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

 Date: 15/01/2018 19:01:11
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
  "node_key" int4 NOT NULL DEFAULT nextval('alias_node_key_seq'::regclass)
)
;
ALTER TABLE "alias" OWNER TO "postgres";

-- ----------------------------
-- Indexes structure for table alias
-- ----------------------------
CREATE UNIQUE INDEX "alias_id_uindex" ON "alias" USING btree (
  "id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table alias
-- ----------------------------
ALTER TABLE "alias" ADD CONSTRAINT "alias_pkey" PRIMARY KEY ("id");
