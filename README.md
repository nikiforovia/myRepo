# myRepo
Данная програма является rest-api сервером.
Для создания веб сервера использовался Postgresql версии 12.5 и go1.15.6.windows-amd64.msi а также pgAdmin версии 4.28.
С помощью простой формы обращения в HTML файле программа делает запрос на добавление, удаление, изменение данных в Postgresql.
С помощью пакета "github.com/lib/pq" осуществляется подключение к базе данныхю.
С помощью фреймворка "github.com/gin-gonic/gin" развертывается веб сервер, который имеет get запрос к HTML файлу и 4 post запроса на изменение базы данных.




Сначала, открывается файл конфигурации (settings.cfg). В нем находятся переменные ,относительно данного репозитория. 
Затем, с помощью этих переменных происходит подключение к postgresql (db.go). В этом же файле подготовленны функции с помощью которых будет происходить post запрос.
В main.go начинается развертывание сервера. В данном проекте использовалось http соединение посредством функции gin.Default().
Далее задаются пост запросы и выделяются горутины для многопоточности. Ниже функции main находится функции этих запросов.
© 2020 GitHub, Inc.
                                                                                      p.s в запросе на сортировку пользователей есть ошибка. Работают 3 запроса из 4.
