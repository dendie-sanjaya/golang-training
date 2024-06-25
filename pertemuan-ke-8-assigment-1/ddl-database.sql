-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id int4 NOT NULL DEFAULT nextval('user_user_id_seq'::regclass),
	name varchar(50) NOT NULL,
	email varchar(50) NOT NULL,
	created_at date NULL,
	updated_at date NULL,
	CONSTRAINT user_pkey PRIMARY KEY (id)
);


-- public.submissions definition

-- Drop table

-- DROP TABLE public.submissions;

CREATE TABLE public.submissions (
	id int4 NOT NULL DEFAULT nextval('submissions_submission_id_seq'::regclass),
	user_id int4 NULL,
	answers json NULL,
	risk_score int4 NULL,
	risk_category varchar NULL,
	created_at date NULL,
	updated_at date NULL,
	CONSTRAINT submissions_pkey PRIMARY KEY (id),
	CONSTRAINT submissions_fk FOREIGN KEY (user_id) REFERENCES public.users(id)
);