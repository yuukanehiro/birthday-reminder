CREATE TABLE users
(
  id SERIAL,
  insert_date TIMESTAMP,
  update_date TIMESTAMP,
  PRIMARY KEY (id)
);
INSERT INTO users (id) VALUES (1),(2),(3);

CREATE TABLE user_accounts
(
  id SERIAL,
  user_id BIGINT,
  password VARCHAR(255),
  insert_date TIMESTAMP,
  update_date TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY(user_id) REFERENCES users(id)
);
INSERT INTO user_accounts (id, user_id, password) VALUES
  (1, 1, 'secret'),
  (2, 1, 'secret'),
  (3, 2, 'secret'),
  (4, 3, 'secret');

CREATE TABLE birth_days
(
  id SERIAL,
  user_id BIGINT,
  name VARCHAR(255),
  date DATE,
  insert_date TIMESTAMP,
  update_date TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY(user_id) REFERENCES users(id)
);
INSERT INTO birth_days (id, user_id, name, date) VALUES
  (1, 1, 'A太郎', '2018-10-20'),
  (2, 1, 'B子', '2018-10-21'),
  (3, 2, 'C助', '2018-11-01'),
  (4, 2, 'D浪', '2018-11-02');