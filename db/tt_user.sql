
DROP TABLE IF EXISTS "public"."user";
CREATE TABLE "public"."user" (
  "id" bigint NOT NULL,
  "email" varchar(50) COLLATE "pg_catalog"."default",
  "password" varchar(255) COLLATE "pg_catalog"."default",
  "is_premium" boolean DEFAULT false,
  "verified" boolean DEFAULT false,
  "quota_swipe" integer DEFAULT 10,
  "created_at" timestamp with time zone DEFAULT now() NOT NULL,
)
;


ALTER TABLE "public"."user" ADD CONSTRAINT "user_pkey" PRIMARY KEY ("id");

CREATE SEQUENCE public.user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE user_id_seq OWNED BY user.id;

ALTER TABLE ONLY public.user ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);
