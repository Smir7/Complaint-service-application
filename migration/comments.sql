/*
По каждой жалобе записываем историю в отдельную таблицу.
Исторические данные удаляется спустя год, основная таблица - без очистки.

Согласно ТЗ следует отражать uuid жалобы в названии таблицы.
*/

CREATE TABLE IF NOT EXISTS comment%s (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	user_uuid UUID,
	comment TEXT,
	created_at DATE NOT NULL DEFAULT CURRENT_DATE,
	updated_at DATE NOT NULL DEFAULT CURRENT_DATE
);

/*
Общий вариант БД комментариев для хранения ВСЕХ коментариев ко всем жалобам. 
Этот вариант может упростить поиск информации в комментрариях, т е не нужно 
пробегать по всем файлам комментов для поиска.
*/
CREATE TABLE IF NOT EXISTS comments (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	compliance_uuid UUID,
	user_uuid UUID,
	comment TEXT,
	created_at DATE NOT NULL DEFAULT CURRENT_DATE,
	updated_at DATE NOT NULL DEFAULT CURRENT_DATE
);

/*
Какой вариант выбрать и оставить в PR???
*/
