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
  show_id int,
  title VARCHAR(100),
  slug VARCHAR(300),
  description VARCHAR(400),
  number integer,
  published_at date DEFAULT now()
);

CREATE INDEX episodes_show_id ON episodes(show_id);

CREATE TABLE shows (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100),
  slug VARCHAR(100)
);

INSERT INTO shows (name, slug) VALUES('Adventures in Angular', 'adventures-in-angular');
INSERT INTO shows (name, slug) VALUES('Ruby Rogues', 'ruby-rogues');
INSERT INTO shows (name, slug) VALUES('JavaScript Jabber', 'js-jabber');
INSERT INTO shows (name, slug) VALUES('iPhreaks', 'iphreaks');
INSERT INTO shows (name, slug) VALUES('Freelancers'' Show', 'freelancers');
