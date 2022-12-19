
CREATE TABLE IF NOT EXISTS public.referred
(
    id integer NOT NULL DEFAULT nextval('referred_id_seq'::regclass),
    agent_id character varying(255) COLLATE pg_catalog."default" NOT NULL,
    job_url character varying(255) COLLATE pg_catalog."default",
    request_referer character varying(255) COLLATE pg_catalog."default",
    request_ip character varying(255) COLLATE pg_catalog."default",
    request_user_agent character varying(255) COLLATE pg_catalog."default",
    created_at date,
    updated_at date,
    deleted_at date,
    CONSTRAINT referred_pkey PRIMARY KEY (id)
)