/*修改日期字段为unix时间*/
ALTER TABLE public."user" ADD created_at_bak INTEGER NULL;
ALTER TABLE public."user" ADD updated_at_bak INTEGER NULL;
ALTER TABLE public."user" ADD deleted_at_bak INTEGER NULL;

/*修改将日期字段转化为unix时间*/
update mkb.public.user set created_at_bak = extract(epoch from created_at);
update mkb.public.user set updated_at_bak = extract(epoch from updated_at);
update mkb.public.user set deleted_at_bak = extract(epoch from deleted_at);

ALTER TABLE public."user" DROP created_at;
ALTER TABLE public."user" DROP updated_at;
ALTER TABLE public."user" DROP deleted_at;

ALTER TABLE public."user" RENAME COLUMN created_at_bak TO created_at;
ALTER TABLE public."user" RENAME COLUMN updated_at_bak TO updated_at;
ALTER TABLE public."user" RENAME COLUMN deleted_at_bak TO deleted_at;




ALTER TABLE public."doc" ADD created_at_bak INTEGER NULL;
ALTER TABLE public."doc" ADD updated_at_bak INTEGER NULL;
ALTER TABLE public."doc" ADD deleted_at_bak INTEGER NULL;

update mkb.public.doc set created_at_bak = extract(epoch from created_at);
update mkb.public.doc set updated_at_bak = extract(epoch from updated_at);
update mkb.public.doc set deleted_at_bak = extract(epoch from deleted_at);

ALTER TABLE public."doc" DROP created_at;
ALTER TABLE public."doc" DROP updated_at;
ALTER TABLE public."doc" DROP deleted_at;

ALTER TABLE public."doc" RENAME COLUMN created_at_bak TO created_at;
ALTER TABLE public."doc" RENAME COLUMN updated_at_bak TO updated_at;
ALTER TABLE public."doc" RENAME COLUMN deleted_at_bak TO deleted_at;



ALTER TABLE public."alias" ADD created_at_bak INTEGER NULL;
ALTER TABLE public."alias" ADD updated_at_bak INTEGER NULL;
ALTER TABLE public."alias" ADD deleted_at_bak INTEGER NULL;

update mkb.public.alias set created_at_bak = extract(epoch from created_at);
update mkb.public.alias set updated_at_bak = extract(epoch from updated_at);
update mkb.public.alias set deleted_at_bak = extract(epoch from deleted_at);

ALTER TABLE public."alias" DROP created_at;
ALTER TABLE public."alias" DROP updated_at;
ALTER TABLE public."alias" DROP deleted_at;

ALTER TABLE public."alias" RENAME COLUMN created_at_bak TO created_at;
ALTER TABLE public."alias" RENAME COLUMN updated_at_bak TO updated_at;
ALTER TABLE public."alias" RENAME COLUMN deleted_at_bak TO deleted_at;





