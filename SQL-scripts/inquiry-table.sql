CREATE TYPE inquiry_type AS ENUM(
	'booking_question',
	'adding_question',
	'rescheduling_canceling',
	'other'
);

CREATE TABLE chalet_inquiries(
	id SERIAL PRIMARY KEY,
	user_id UUID NOT NULL,
	chalet_id BIGINT,
	type inquiry_type NOT NULL,
	message TEXT NOT NULL,
	status VARCHAR(50) DEFAULT 'pending',
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

	CONSTRAINT fk_inquiry_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
	CONSTRAINT fk_inquiry_chalet FOREIGN KEY (chalet_id) REFERENCES chalets(id) ON DELETE CASCADE
);