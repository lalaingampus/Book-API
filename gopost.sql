-- Table: buku

-- DROP TABLE buku;

create table buku
(
    id serial not null,
    judul_buku character varying not null,
    penulis character varying,
    tgl_publikasi date,
    constraint pk_buku primary key (id)
)

with (
    OIDS=FALSE
);
ALTER TABLE buku 
 OWNER TO postgres;