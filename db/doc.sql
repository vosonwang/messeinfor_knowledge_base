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

 Date: 15/01/2018 19:01:20
*/


-- ----------------------------
-- Table structure for doc
-- ----------------------------
DROP TABLE IF EXISTS "doc";
CREATE TABLE "doc" (
  "id" varchar(36) COLLATE "pg_catalog"."default" NOT NULL DEFAULT NULL,
  "lang" int2 NOT NULL DEFAULT NULL,
  "title" varchar(300) COLLATE "pg_catalog"."default" DEFAULT NULL,
  "text" text COLLATE "pg_catalog"."default" DEFAULT NULL,
  "created_at" timestamp(6) DEFAULT NULL,
  "updated_at" timestamp(6) DEFAULT NULL,
  "deleted_at" timestamp(6) DEFAULT NULL,
  "parent_id" varchar(36) COLLATE "pg_catalog"."default" NOT NULL DEFAULT NULL,
  "alias_id" varchar(36) COLLATE "pg_catalog"."default" NOT NULL DEFAULT NULL
)
;
ALTER TABLE "doc" OWNER TO "postgres";

-- ----------------------------
-- Indexes structure for table doc
-- ----------------------------
CREATE UNIQUE INDEX CONCURRENTLY "docs_id_uindex" ON "doc" USING btree (
  "id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table doc
-- ----------------------------
ALTER TABLE "doc" ADD CONSTRAINT "docs_id_pk" PRIMARY KEY ("id");

-- ----------------------------
-- Foreign Keys structure for table doc
-- ----------------------------
ALTER TABLE "doc" ADD CONSTRAINT "doc_alias_id_fk" FOREIGN KEY ("alias_id") REFERENCES "alias" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
