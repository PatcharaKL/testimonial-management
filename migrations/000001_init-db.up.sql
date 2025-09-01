CREATE TABLE testimonials (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    role VARCHAR(255),
    company VARCHAR(255),
    testimonial TEXT NOT NULL,
    photo_url VARCHAR(255),
    is_public BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
