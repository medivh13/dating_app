DROP TABLE IF EXISTS "public"."profile";
CREATE TABLE  "public."."profile" (
    id bigint NOT NULL PRIMARY KEY,
    user_id bigint,
    name VARCHAR(255),
    gender VARCHAR(10),
    age INT,
    lat DOUBLE PRECISION,
    long DOUBLE PRECISION,
    picture varchar(255),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.user(id)

);

CREATE SEQUENCE public.profile_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE public.profile_id_seq OWNED BY public.profile.id;

ALTER TABLE ONLY public.profile ALTER COLUMN id SET DEFAULT nextval('public.profile_id_seq'::regclass);