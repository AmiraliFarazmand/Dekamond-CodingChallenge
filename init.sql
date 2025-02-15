CREATE TABLE IF NOT EXISTS tables (
    id SERIAL PRIMARY KEY,
    seats_number INT NOT NULL CHECK (seats_number >= 4 AND seats_number <= 10),
    is_reserved BOOLEAN DEFAULT FALSE
);

INSERT INTO tables (seats_number, is_reserved) VALUES
(4, FALSE),
(4, FALSE),
(6, FALSE),
(6, FALSE),
(6, FALSE),
(6, FALSE),
(8, FALSE),
(8, FALSE),
(10, FALSE),
(10, FALSE);