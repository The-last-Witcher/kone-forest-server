CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE Users(
	id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
	name VARCHAR(30) NOT NULL,
	email VARCHAR(100) NOT NULL UNIQUE,
	password VARCHAR(255) NOT NULL
);

CREATE TABLE Chalets(
	id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	
	title VARCHAR(100) NOT NULL,
	description TEXT NOT NULL,
	address VARCHAR(100) NOT NULL,
	
	price MONEY NOT NULL,
	rating NUMERIC(3,2) NOT NULL,
	
	guestsMax SMALLINT NOT NULL,
	bedrooms SMALLINT NOT NULL,
	bathrooms SMALLINT NOT NULL,

	is_active BOOLEAN DEFAULT TRUE NOT NULL
);

CREATE TABLE Reservation(
	id UUID DEFAULT gen_random_uuid() PRIMARY KEY,

	user_id UUID NOT NULL,
	chalet_id BIGINT NOT NULL,

	check_in DATE NOT NULL,
	check_out DATE NOT NULL,

	status VARCHAR(20) DEFAULT 'confirmed' NOT NULL,

	CONSTRAINT fk_reservation_user 
        FOREIGN KEY (user_id) 
        REFERENCES Users(id) 
        ON DELETE CASCADE,

	CONSTRAINT fk_reservation_hotel 
        FOREIGN KEY (chalet_id) 
        REFERENCES Chalets(id) 
        ON DELETE RESTRICT,

	CONSTRAINT check_reservation
        CHECK (check_out > check_in)
);

CREATE INDEX idx_reservation_hotel_dates 
ON Reservation (chalet_id, check_in, check_out) 
WHERE status = 'confirmed';