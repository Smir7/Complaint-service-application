# Сервис подачи жалоб



- [*Описание проекта*](#описание_проекта)
- [*Установка*](#установка)
- [*Использование*](#использование)
- [*Опции приложения*](#опции_приложения)
- [*Контрибьюция*](#контрибьюция)


## 1. *Описание проекта*
Приложение предназначено для подачи жалоб и сопровождения их на пути от регистрации до решения. Сервис позволяет пользователю зарегистрироваться, авторизовываться и отслеживать статус своей жалобы. 
Цель - предоставление широкому кругу пользователей простого и удобного интерфейса, что бы они могли делиться своими проблемами. 

## 2. *Установка*
Для успешного запуска приложения необходимо установить [Docker](https://www.docker.com/), [GNU Make](https://www.gnu.org/software/make/)

## 3. *Использование*
Для упрощения взаимодействия с приложением используется утилита make. В терминале вы должны находится в корневом каталоге проекта.<br><br>
3.1. Собрать образ через Docker-Compose:
```bash
make build
```
3.2. Запустить контейнеры:
```bash
make run-local
```
3.3. Если контейнеры на компьютере запущены впервые и не создан сервер СУБД, то нужно создать сервер СУБД. Это можно сделать через консольную утилиту psql или веб-интерфейс pgAdmin, который является одним из запущенных контейнеров.<br><br>
Для доступа к pgAdmin и созданию сервера, необходима информация о контейнерах pgAdmin и PostgreSQL, указанная в файле docker-compose и конфигурации приложения, заданных в переменных окружения.<br><br>
Далее предлагается инструкция по созданию сервера СУБД с параметрами по-умолчанию (могут быть изменены) для локального компьютера через pgAdmin:<br>
 
>3.3.1. Открыть браузер.<br>
3.3.2. Ввести в адресную строку localhost:5050 и нажать Enter.<br>
3.3.3. Появится веб-интерфейс pgAdmin. Введите мастер-пароль, например postgres.<br>
3.3.4. В веб-интерфейсе pgAdmin с левой стороны должен быть значок Servers. Щёлкните по нему правой кнопкой мыши.<br>
3.3.5. В появившемся окне выберите Register > Server...<br>
3.3.6. Появится окно создания сервера с несколькими вкладками. В первой вкладке General в поле Name введите произвольное имя сервера СУБД, например PostgreSQL_server_practicum.<br>
3.3.7. Перейдите во вкладку Connection. В поле Host name/adress введите наименование контейнера PostgreSQL из docker-compose-файла, а именно: postgres_complaint_service_1.<br>
Для поля Port: 5432<br>
Для полей Maintaince database, Username и Password: postgres<br>
3.3.8. Остальные настройки оставьте по-умолчанию и нажимите кнопку Save внизу справа.<br>
3.3.9. Сервер СУБД готов к работе.<br>
<br>

3.4. Применить миграции схемы БД<br>
В случае запуска docker-контейнеров без фонового режима (поведение по-умолчанию), потребуется запуск отдельного терминала (например, сочетанием клавишь Ctrl+Enter в VSCode) и выполнения команды в нём:
```bash
make migrate-up
```
Приложение готово к работе.

## 4. *Опции приложения*
4.1. Посмотреть статус миграций в БД
```bash
make migrate-status:
```
4.2. Откатить последнюю миграцию
```bash
make migrate-down
```
4.3. Создать файл миграции с заданным форматом<br>
Используется для изменения схемы БД с соблюдением стандарта префикса наименования файла миграции для очерёдности применения миграций.
```bash
make migrate-create name=table
```
где table - желаемое смысловое имя миграции, например first_admin.<br>
Далее в созданную таблицу нужно внести sql-код с аннотациями библиотеки миграций. Далее миграция готова к применению.

4.4. В приложении реализованы:

 - Метод регистрации для пользователя. 

Пользователь может зарегистрироваться, указав свое имя и пароль. Пароль хешируется с использованием алгоритма SHA-256 для повышения безопасности.

Пример ввода данных при регистрации.

```bash
{
  "username": "ваше_имя",
  "password": "ваш_пароль"
}
```
- Метод авторизации для пользователя.

Пользователь может авторизоваться с использованием своего имени и пароля. В случае успешной аутентификации возвращается JWT токен, который будет использован при последующих запросах.

Пример ввода данных при авторизации

```bash
{
  "username": "ваше_имя",
  "password": "ваш_пароль"
}
```
- Метод подачи жалобы

Зарегистрированные и авторизованные пользователи могут подавать жалобы. Каждая жалоба содержит описание проблемы и приоритет.

Пример приема жалобы в приложении. 

```bash
{
  "description": "Описание вашей жалобы",
  "priority": "high/medium/low"
}
```

## №5. *Контрибьюция*
Если вы хотите внести свой вклад в данное приложение, пожалуйста, создайте форк репозитория, внесите ваши изменения и создайте pull request. Мы будем рады вашим идеям и предложениям!

______
На этом этапе сервис подачи жалобы готов к использованию. Если у вас есть какие-либо вопросы или предложения, вы можете обратиться к разработчиками открыв issue в нашем репозитории, и мы постараемся помочь. Спасибо за использование нашего приложения!
