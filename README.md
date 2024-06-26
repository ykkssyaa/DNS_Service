# DNS Service
## Описание работы
Данный проект представляет собой сервер и клиент 
для выставления имени хоста в Linux, а также изменения списка DNS серверов.

Сервис предоставляет grpc gateway для общения между сервисами, 
а также REST API, которое позволяет работать с сервисом напрямую. 
Сервис для работы с именем хоста использует утилиты `hostname` и `hostnamectl set-hostname`. 

Для работы со списком DNS сервис напрямую обращается к файлу конфигурации `/etc/resolv.conf`. Для изменения данного файла
необходимы права суперпользователя, поэтому сервис нужно запускать через `sudo`.

Клиент предоставляет CLI интерфейс, реализованный с помощью фреймворка [cobra](https://github.com/spf13/cobra).


При разработке использовались Ubuntu 22.04.4 LTS и Golang 1.21 

## Запуск
Для запуска сервиса необходимо сбилдить бинарный файл и запустить его с правами суперпользователя 
1. `go build -v ./server/cmd/main.go` - билд сервиса
2. `sudo ./main` - запуск сервиса с правами суперпользователя

Для работы с клиентом можно также сбилдить бинарный файл и работать с ним напрямую
1. `go build -v -o ./cl ./client/main.go` - билд клиента
2. `./cl` - просмотр информации о командах CLI 

### Реализованные команды в CLI
1. `./cl hostname` - Получение имени хоста 
2. `./cl setHostname [hostname]` - Изменение имени хоста Linux
3. `./cl dnsList` - Получение списка DNS серверов
4. `./cl addDns [address]` - Добавление нового адресса DNS
5. `./cl removeDns [address]` - Удаление адресса DNS из списка

### Работа с REST API сервиса
С помощью [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) 
был сгенерирован [файл SwaggerAPI](./server/api/dns_service.swagger.json). 
С помощью него я сгенерировал коллекцию в Postman, которой можно воспользоваться по [ссылке](https://www.postman.com/joint-operations-operator-99149269/workspace/dns/collection/28284200-387f82b9-ceae-483a-8688-7a31d495443a?action=share&creator=28284200).

