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

 Date: 23/01/2018 16:54:04
*/


-- ----------------------------
-- Table structure for doc_alias
-- ----------------------------
DROP TABLE IF EXISTS "doc_alias";
CREATE TABLE "doc_alias" (
  "id" varchar(36) COLLATE "pg_catalog"."default" DEFAULT NULL,
  "alias_id" varchar(36) COLLATE "pg_catalog"."default" DEFAULT NULL,
  "doc_id" varchar(36) COLLATE "pg_catalog"."default" DEFAULT NULL,
  "created_at" timestamp(6) DEFAULT NULL,
  "updated_at" timestamp(6) DEFAULT NULL,
  "deleted_at" timestamp(6) DEFAULT NULL
)
;
ALTER TABLE "doc_alias" OWNER TO "postgres";

-- ----------------------------
-- Foreign Keys structure for table doc_alias
-- ----------------------------
ALTER TABLE "doc_alias" ADD CONSTRAINT "doc_alias_alias_id_fk" FOREIGN KEY ("alias_id") REFERENCES "alias" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "doc_alias" ADD CONSTRAINT "doc_alias_doc_id_fk" FOREIGN KEY ("doc_id") REFERENCES "doc" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
