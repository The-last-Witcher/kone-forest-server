CREATE TYPE inquiry_topics AS ENUM(
	'booking_question',
	'adding_chalet',
	'rescheduling_canceling',
	'other'
);

CREATE TABLE contact_inquiries(
	id SERIAL PRIMARY KEY,
	first_name VARCHAR(100) NOT NULL,
	last_name VARCHAR(100) NOT NULL,
	email VARCHAR(150) NOT NULL,
	topic inquiry_topics NOT NULL,
	message TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)