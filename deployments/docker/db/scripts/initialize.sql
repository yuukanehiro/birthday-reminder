create table users
(
  id bigint UNIQUE,
  password TEXT,
  insert_date timestamp with time zone,
  update_date timestamp with time zone
);
INSERT INTO users (id, password) VALUES (1, 'secret');

create table birth_days
(
  id bigint UNIQUE,
  user_id bigint,
  date DATE,
  insert_date timestamp with time zone,
  update_date timestamp with time zone
);
INSERT INTO birth_days (id, user_id, date) VALUES (1, 1, '2018-10-20');
INSERT INTO birth_days (id, user_id, date) VALUES (2, 1, '2018-10-21');