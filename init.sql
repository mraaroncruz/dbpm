CREATE TABLE picks (
  id SERIAL PRIMARY KEY,
  episode_id int,
  host VARCHAR(100),
  name VARCHAR(300),
  link VARCHAR(300),
  description VARCHAR(400),
  content TEXT
);

CREATE INDEX episode_id_idx ON picks(episode_id);
CREATE INDEX link_idx ON picks(link);
CREATE INDEX search_idx ON picks USING gin(to_tsvector('english', name || ' ' || description || ' ' || content));

CREATE TABLE episodes (
  id SERIAL PRIMARY KEY,
  title VARCHAR(100),
  slug VARCHAR(300),
  description VARCHAR(400),
  number integer,
  published_at date
);
