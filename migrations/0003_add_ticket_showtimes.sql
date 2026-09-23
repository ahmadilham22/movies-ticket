BEGIN;

ALTER TABLE tickets
    ADD COLUMN movie_id UUID NOT NULL,
    ADD COLUMN starts_at TIMESTAMPTZ NOT NULL,

    ALTER COLUMN price SET NOT NULL,
    ALTER COLUMN quota SET NOT NULL,

    DROP COLUMN event_name,

    ADD CONSTRAINT fk_tickets_movie
        FOREIGN KEY (movie_id)
        REFERENCES movies(id)
        ON DELETE RESTRICT,

    ADD CONSTRAINT chk_tickets_price_positive
        CHECK (price > 0),

    ADD CONSTRAINT chk_tickets_quota_non_negative
        CHECK (quota >= 0);

CREATE INDEX idx_tickets_movie_id_starts_at
    ON tickets (movie_id, starts_at);

COMMIT;