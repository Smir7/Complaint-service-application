/*
По каждой жалобе записываем историю в отдельную таблицу.
Исторические данные удаляется спустя год, основная таблица - без очистки.
Согласно ТЗ следует отражать uuid жалобы в названии таблицы

Из коммента в ТГ следует, что вторая таблица не нужна. Удаляю.
*/

CREATE TABLE IF NOT EXISTS comments (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	user_uuid UUID,
	comment VARCHAR(255),
	created_at DATE NOT NULL DEFAULT CURRENT_DATE,
	updated_at DATE NOT NULL DEFAULT CURRENT_DATE
);

