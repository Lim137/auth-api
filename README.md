The project implements two routes for JWT authorization.<br>
Route 1: access and refresh tokens are created based on the user's GUID. The access token type is JWT. The refresh token type is string.<br>
Route 2: Performs a Refresh operation on a pair of access and refresh tokens.<br><br>

Input data for routes:<br>
Route 1: URL - your_server/v1/tokens/user_guid. For example: http://localhost:8000/v1/tokens/123 . The request method is Get.<br>
Route 2: URL - your_server/v1/tokens/refresh. For example: http://localhost:8000/v1/tokens/refresh. The request method is Post. In the request body, you must specify access_token and refresh_token in JSON format. For example:<br>
{<br>
"access_token": "qwe.rty.uiop",<br>
"refresh_token": "asdfghjk"<br>
}<br><br>

Output:<br>
Both routes return a pair of access_token and refresh_token tokens in JSON format. For example:<br>
{<br>
"access_token": "zxc.vbn.ghjkl",<br>
"refresh_token": "lkjhgf"<br>
}<br>
An error may also be returned if it occurs. For example:<br>
{<br>
"error": "problem with verifying the uniqueness of the GUID: This user already has refresh token"<br>
}<br><br>

Inside the access token, information is stored about the ID of the user for whom the token was issued and the validity period of the token in UNIX format.<br><br>

The project uses the MongoDB database. It stores information about the issued refresh tokens. This information contains: the user ID (GUID), the hash of the refresh token, and the expiration date of the token in UNIX format.<br><br>

To run this project, follow these steps:<br>
1. Change the data in the .env.example file to the current ones<br>
2. Change the name of the .env.example file to .env<br>
3. Go to the project folder and run the go build command.<br>
4. Run the created file (its name will match the folder name)<br>
---

Проект реализует два маршрута для JWT авторизации.<br>
Маршрут 1: создаются access и refresh токены по GUID пользователя. Тип access токена - JWT. Тип refresh токена - string.<br>
Маршрут 2: выполняет операцию Refresh на пару токенов access и refresh.<br><br>

Входные данные для маршрутов:<br>
Маршрут 1: URL - your_server/v1/tokens/user_guid. Например: http://localhost:8000/v1/tokens/123. Метод запроса - Get.<br>
Маршрут 2: URL - your_server/v1/tokens/refresh. Например: http://localhost:8000/v1/tokens/refresh. Метод запроса - Post. В теле запроса необходимо указать access_token и refresh_token в формате JSON. Например:<br>
{<br>
  "access_token": "qwe.rty.uiop",<br>
  "refresh_token": "asdfghjk"<br>
}<br><br>

Выходные данные:<br>
Оба маршрута возвращают пару токенов access_token и refresh_token в формате JSON. Например:<br>
{<br>
  "access_token": "zxc.vbn.ghjkl",<br>
  "refresh_token": "lkjhgf"<br>
}<br>
Также может возващаться ошибка, если она возникнет. Например:<br>
{<br>
  "error": "problem with verifying the uniqueness of the GUID: This user already has refresh token"<br>
}<br><br>

Внутри access токена хранится информация о id пользователя, для которого был выдан токен и срок действия токена в формате UNIX.<br><br>

В проекте используется база данных MongoDB. В ней хранится информация о выданных refresh токенах. Эта информация содержит: id пользователя (GUID), хэш refresh токена и срок действия токена в формате UNIX.<br><br>

Для запуска этого проекта нужно выполнить следующие шаги:<br>
1. Поменять данные в файле .env.example на актуальные<br>
2. Поменять название файла .env.example на .env<br>
3. Перейти в папку с проектом и выполнить команду go build.<br>
4. Запустить созданный файл (его название будет совпадать с названием папки)
