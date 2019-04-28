create table if not exists post
(
  filename  varchar(256) not null,
  timestamp datetime     not null,
  hash      char(32)     not null,
  title     varchar(256) not null,
  body      text         not null,
  primary key (filename)
) row_format = dynamic;

create table if not exists post_label
(
  filename varchar(256),
  label    varchar(256)
) row_format = dynamic;
