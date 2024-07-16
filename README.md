# Time Tracker

## Описание
Тестовое задание от компании Effective Mobile. Если коротко, то это приложение представляет из мебя реализацию некоторых API-методов для достижения цели.
Изначально приложение собиралось в Docker, поэтому рекомендую использовать именно его. Для того, чтобы приложение функционировало как надо я предлагаю использовать вам
те же версии инструментов, что и я

## Инструменты

1. Docker version 27.0.3
2. go version go1.22.5
3. GNU Make 4.3

## WARNING
Для начала определитесь, что вам нужно: вы хотите увидеть тесты, вы хотите зпустить тесты каждого метода по отдельности, или же вы хотите просто включить все end pointы?
Для того, чтобы запустить все тесты сразу вам ничего не нужно делать, сразу же можите переходить к разделу установка и запуск. Если вы хотите запустить тесты каждого метода
по отдельности, тогда вам нужо склонировать репозиторий, а потом изменить файл timetracker.go в главной дириктории.

## А именно:

1. Откройте файл timetracker.go в вашем редакторе кода
2. Найдите функцию main()
3. Под коментарием "Сервера по отдельности" будет список функций, которые начинают тестирование Сервера
4. Раскомментируйте вызов той функции, который выбрали
5. Закомоентируйте все остальные testAll() и turnAllOn(), если они еще не закомментированны
6. Следуйте по инструкции для запуска

Если вы хотите просто запустить все методы тогда сделайте

## Это:

1. Откройте файл timetracker.go в вашем редакторе кода
2. Найдите функцию main()
3. Закомоентируйте testAll() и раскоментируйте turnAllOn(), а так же закоментируйте все отсальные функции, если они незакоменчены
4. Следйуте инструкции запуска


## Установка и запуск приложения

1. Клонируйте репозиторий:  
    ```bash
    git clone https://github.com/l1qwie/TimeTracker.git

2. Перейдите в директорию приложения:
    ```bash
    cd TimeTracker

3. Скачайте все зависимости:
    ```bash
    go mod download

4. Скомпилирейте контейнер в Docker:
    ``bash
    make rb
    или
    docker build . -t timetracker-app

5. Запустите docker-compose файл:
    ```bash
    make up
    или
    docker compose -f docker-compose.yml up --force-recreate

6. Готово! Приложение работает!
