-- +goose Up
CREATE TABLE IF NOT EXISTS comments (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	report_id INT,
	user_uuid UUID,
	comment VARCHAR(255),
	created_at DATE NOT NULL DEFAULT CURRENT_DATE,
	updated_at DATE NOT NULL DEFAULT CURRENT_DATE
	FOREIGN KEY (report_id) REFERENCES reports (Id)
);

-- +goose Down
DROP TABLE IF EXISTS comments;
