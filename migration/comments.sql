/*
По каждой жалобе записываем историю в отдельную таблицу.
Исторические данные удаляется спустя год, основная таблица - без очистки.

Согласно ТЗ следует отражать uuid жалобы в названии таблицы
*/

CREATE TABLE IF NOT EXISTS comments (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	user_uuid UUID,
	comment VARCHAR(255),
	created_at DATE NOT NULL DEFAULT CURRENT_DATE,
	updated_at DATE NOT NULL DEFAULT CURRENT_DATE
);

/*
Общий вариант БД комментариев для хранения ВСЕХ коментариев ко всем жалобам. 
Этот вариант может упростить поиск информации в комментрариях, т е не нужно 
пробегать по всем файлам комментов для поиска.
Во второй таблице дополнительное поле compliance_uuid, т к таблица целиковая на все
жалобы. Для различий по жалобам это поле и нужно. Эта таблица не удовлетворяем требованиям ТЗ, 
но возможно облегчит поиск нужной информации в комментариях. Если это лишнее, то таблица будет удалена.
*/
CREATE TABLE IF NOT EXISTS comments (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	FOREIGN KEY (compliance_id) REFERENCES REPORTS (id),
	user_uuid UUID,
	comment VARCHAR(255),
	created_at DATE NOT NULL DEFAULT CURRENT_DATE,
	updated_at DATE NOT NULL DEFAULT CURRENT_DATE
);

/*
Какой вариант выбрать и оставить в PR???
*/
