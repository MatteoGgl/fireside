CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    email text NOT NULL,
    hashed_password text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
)