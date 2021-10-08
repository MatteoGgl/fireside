CREATE TABLE IF NOT EXISTS links (
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    type text NOT NULL,
    url text,
    content text,
    likes integer NOT NULL DEFAULT 0,
    tags text[] NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
)