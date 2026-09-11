CREATE TABLE genres (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name VARCHAR(64) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

	CONSTRAINT chk_genres_name_not_blank
	  CHECK (btrim(name) <> '')
);

CREATE UNIQUE INDEX uidx_genres_name_ci
ON genres (LOWER(btrim(name)));

CREATE TABLE movies (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title VARCHAR(255) NOT NULL,
  synopsis TEXT NOT NULL,
  duration_minutes INTEGER NOT NULL,
  release_date DATE NOT NULL,
  poster_url TEXT NOT NULL,
  trailer_url TEXT,
  age_rating VARCHAR(10) NOT NULL,
  language VARCHAR(64) NOT NULL,
  country VARCHAR(64) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT chk_movies_title_not_blank
    CHECK (btrim(title) <> ''),

  CONSTRAINT chk_movies_synopsis_not_blank
    CHECK (btrim(synopsis) <> ''),

  CONSTRAINT chk_movies_duration_positive
    CHECK (duration_minutes > 0),

  CONSTRAINT chk_movies_poster_url_not_blank
    CHECK (btrim(poster_url) <> ''),

  CONSTRAINT chk_movies_age_rating_not_blank
    CHECK (btrim(age_rating) <> ''),

  CONSTRAINT chk_movies_language_not_blank
    CHECK (btrim(language) <> ''),

  CONSTRAINT chk_movies_country_not_blank
    CHECK (btrim(country) <> '')
);


CREATE TABLE movie_genres (
  movie_id UUID NOT NULL,
  genre_id UUID NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT pk_movie_genres
    PRIMARY KEY (movie_id, genre_id),
  CONSTRAINT fk_movie_genres_movie
    FOREIGN KEY (movie_id) REFERENCES movies(id) ON DELETE CASCADE,
  CONSTRAINT fk_movie_genres_genre
    FOREIGN KEY (genre_id) REFERENCES genres(id) ON DELETE CASCADE
);

CREATE INDEX idx_movie_genres_genre_id
  ON movie_genres (genre_id);
