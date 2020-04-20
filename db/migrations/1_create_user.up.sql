CREATE TABLE public."user"
(
    id serial NOT NULL,
    firstname  text NOT NULL,
    lastname  text NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE public."user"
    OWNER to steven;