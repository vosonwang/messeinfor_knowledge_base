/*
 Navicat Premium Data Transfer

 Source Server         : postgres
 Source Server Type    : PostgreSQL
 Source Server Version : 100000
 Source Host           : localhost:32769
 Source Catalog        : messeinfor_kownledge
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 100000
 File Encoding         : 65001

 Date: 09/01/2018 18:55:54
*/


-- ----------------------------
-- Table structure for docs
-- ----------------------------
DROP TABLE IF EXISTS "docs";
CREATE TABLE "docs" (
  "id" varchar(36) COLLATE "pg_catalog"."default" NOT NULL DEFAULT NULL,
  "lang" int4 NOT NULL DEFAULT NULL,
  "title" varchar(300) COLLATE "pg_catalog"."default" DEFAULT NULL,
  "doc" text COLLATE "pg_catalog"."default" DEFAULT NULL,
  "node_key" int4 NOT NULL DEFAULT nextval('docs_node_key_seq'::regclass),
  "created_at" timestamp(6) DEFAULT NULL,
  "updated_at" timestamp(6) DEFAULT NULL,
  "deleted_at" timestamp(6) DEFAULT NULL,
  "parent_id" varchar(36) COLLATE "pg_catalog"."default" DEFAULT NULL
)
;
ALTER TABLE "docs" OWNER TO "postgres";

-- ----------------------------
-- Indexes structure for table docs
-- ----------------------------
CREATE INDEX "fk_id_lang" ON "docs" USING btree (
  "id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "lang" "pg_catalog"."int4_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table docs
-- ----------------------------
ALTER TABLE "docs" ADD CONSTRAINT "docs_pkey" PRIMARY KEY ("id", "lang");
