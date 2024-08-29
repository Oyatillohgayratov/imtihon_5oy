CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    hotel_id INT REFERENCES hotels(id),
    room_type VARCHAR(255) NOT NULL,
    check_in_date DATE NOT NULL,
    check_out_date DATE NOT NULL,
    total_amount NUMERIC(10, 2) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'Confirmed',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
