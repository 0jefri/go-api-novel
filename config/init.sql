CREATE DATABASE novel_api

CREATE TABLE novel (
  id varchar(100) primary key,
  judul varchar(100),
  penerbit varchar(100),
  tahun_terbit varchar(20),
  penulis varchar(100)
)