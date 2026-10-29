# Kone-forest server
Перед вами проект, который представляет собой серверную часть приложения Kone-forest - сайт по бронированию шале.

## Технологии
В данном проекте использовались технологии:
- Go
- gorilla/mux
- net/http
- pgx5 + pgxpool

## Запуск
Для запуска на локальной машине, вам необходимо:
1. Установить Go, PostgreSQL (желательно с PgAdmin, но это необязательно) и Git
2. Следовать следующим инструкциям:
- git clone https://github.com/The-last-Witcher/kone-forest-server.git kone-forest-server
- cd kone-forest-server
- go mod tidy
3. Далее необходимо:
- Запустить PostgreSQL
- Получить данные о пользователе, пароль, создать БД и прописать необходимые данные в файле .env
- В терминале запустить команду go run .
