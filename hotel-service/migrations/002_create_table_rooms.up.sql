CREATE TABLE rooms (
    id SERIAL PRIMARY KEY,
    hotel_id INT REFERENCES hotels(id),
    room_type VARCHAR(255) NOT NULL,
    price_per_night NUMERIC(10, 2),
    availability BOOLEAN NOT NULL default true
);