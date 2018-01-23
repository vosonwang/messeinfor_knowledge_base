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

 Date: 23/01/2018 16:53:54
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
  "number" int4 DEFAULT nextval('doc_number_seq'::regclass),
  "creator" varchar(36) COLLATE "pg_catalog"."default" DEFAULT NULL,
  "updater" varchar(36) COLLATE "pg_catalog"."default" DEFAULT NULL
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
ALTER TABLE "doc" ADD CONSTRAINT "doc_creater_user_id_fk" FOREIGN KEY ("creator") REFERENCES "user" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "doc" ADD CONSTRAINT "doc_updater_user_id_fk" FOREIGN KEY ("updater") REFERENCES "user" ("id") ON DELETE NO ACTION ON UPDATE NO ACTION;
