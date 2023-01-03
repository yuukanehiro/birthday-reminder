CREATE TABLE users
(
  id BIGINT UNIQUE,
  password TEXT,
  insert_date TIMESTAMP WITH TIME ZONE,
  update_date TIMESTAMP WITH TIME ZONE,
  PRIMARY KEY (id)
);
INSERT INTO users (id, password) VALUES (1, 'secret');

CREATE TABLE birth_days
(
  id BIGINT UNIQUE,
  user_id BIGINT,
  date DATE,
  insert_date TIMESTAMP WITH TIME ZONE,
  update_date TIMESTAMP WITH TIME ZONE,
  PRIMARY KEY (id),
  FOREIGN KEY(user_id) REFERENCES users(id)
);
INSERT INTO birth_days (id, user_id, date) VALUES (1, 1, '2018-10-20');
INSERT INTO birth_days (id, user_id, date) VALUES (2, 1, '2018-10-21');