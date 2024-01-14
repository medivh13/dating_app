CREATE TABLE public.swipe_history (
    id bigint,
    user_id bigint,
    profile_id bigint,
    action_type VARCHAR(5) CHECK (action_type IN ('pass', 'like')),
    swipe_date DATE,
    CONSTRAINT fk_user_swipe FOREIGN KEY (user_id) REFERENCES public.user(id),
    CONSTRAINT fk_profile_swipe FOREIGN KEY (profile_id) REFERENCES public.profile(id)
);


ALTER TABLE "public"."swipe_history" ADD CONSTRAINT "swipe_pkey" PRIMARY KEY ("id");

CREATE SEQUENCE public.swipe_history_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE swipe_history_id_seq OWNED BY public.swipe_history.id;

ALTER TABLE ONLY public.swipe_history ALTER COLUMN id SET DEFAULT nextval('public.swipe_history_id_seq'::regclass);