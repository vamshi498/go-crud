CREATE TABLE public."user"
(
    id serial,
    firstname character varying(50)[] NOT NULL,
    lastname character varying(50)[] NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE public."user"
    OWNER to steven;