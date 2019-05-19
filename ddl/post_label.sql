create table if not exists post_label
(
  filename varchar(256),
  label    varchar(256)
) row_format = dynamic;
