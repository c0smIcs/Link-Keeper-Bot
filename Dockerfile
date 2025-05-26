# ---------- ЭТАП 1. Сборка приложения ----------
FROM golang:1.24.3 AS builder

# Устанавливает рабочую директорию внутри контейнера как /app
# (Более простыми словами: создаем папку /app где будут размещены файлы приложения)
WORKDIR /app

# Копируем go.mod и go.sum в папку /app
COPY go.mod go.sum ./

# Загружает все зависимости, указанные в go.mod 
# Убеждаемся, что все библиотеки доступны перед сборкой. 
RUN go mod download

# Копируем все файлы из текущей директории (где находится Dockerfile) в /app
COPY . .

# Компилируем приложение Go в исполняемый файл, названный "Link-Keeper-Bot"
RUN go build -o Link-Keeper-Bot


# ---------- ЭТАП 2. Запуск приложения ----------
FROM alpine:3.21

# Устанавливаем библиотеки SQLite, необходимые для приложения
RUN apk add --no-cache sqlite-libs

# Создаем группу "info_hub" и пользователя "info_vault" принадлежащего этой группе.
RUN addgroup -S info_hub && adduser -S info_vault -G info_hub

WORKDIR /app

# Копируем исполняемый файл "Link-Keeper-Bot" из этапа builder в /app этого этапа
COPY --from=builder /app/Link-Keeper-Bot /app/Link-Keeper-Bot

# Создаем директорию для бд /app/data/sqlite
RUN mkdir -p /app/data/sqlite

# Изменение владельца директории /app/data/sqlite на info_vault:info_hub
RUN chown info_vault:info_hub /app/data/sqlite

# Устанавливаем пользователя как info_vault для выполнения последующих команд
USER info_vault

# Указываем команду для выполнения при запуске контейнера, запускаем приложение Link-Keeper-Bot
CMD ["/app/Link-Keeper-Bot"]
