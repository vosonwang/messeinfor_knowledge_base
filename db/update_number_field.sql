CREATE SEQUENCE public.alias_number_seq NO MINVALUE NO MAXVALUE NO CYCLE;
ALTER TABLE public.alias ALTER COLUMN number SET DEFAULT nextval('public.alias_number_seq');
ALTER SEQUENCE public.alias_number_seq OWNED BY public.alias.number;


CREATE SEQUENCE public.doc_number_seq NO MINVALUE NO MAXVALUE NO CYCLE;
ALTER TABLE public.doc ALTER COLUMN number SET DEFAULT nextval('public.doc_number_seq');
ALTER SEQUENCE public.doc_number_seq OWNED BY public.doc.number;