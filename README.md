# Анализатор размера диска (Go)

[![hexlet-check](https://github.com/EAplemakh/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/EAplemakh/go-project-242/actions)

Консольная утилита, которая считает размер файла или директории с флагами рекурсивного обхода, человекочитаемого формата и учёта скрытых файлов.

Учебный проект Хекслета: https://ru.hexlet.io/programs/go


## Стек

- Go

## Установка

### Требования

- [Go](https://go.dev/dl/) 1.27.2 или новее
- [golangci-lint](https://golangci-lint.run/) v2 — только для разработки (запуск линтера)

### Клонирование и зависимости

```bash
git clone https://github.com/EAplemakh/go-project-242.git
cd go-project-242
go mod download
```

Зависимости (`github.com/urfave/cli/v3`) также будут загружены автоматически при первой сборке.

### Сборка

```bash
make build
```

Исполняемый файл появится в `bin/hexlet-path-size`. Проверить, что всё работает:

```bash
./bin/hexlet-path-size -h
```

## Использование
[Демо по использованию утилиты hexlet-path-size](https://asciinema.org/a/ijbvRR1jLFEyApyN)

Вывод справки:
```bash
./bin/hexlet-path-size -h
```

Запуск со всеми флагами:
```bash
./bin/hexlet-path-size -H -a -r testdata
```

---

<details>
<summary>Автоматические тесты Хекслета</summary>

Тесты запускаются на каждый коммит. За запуск отвечает файл `.github/workflows/hexlet-check.yml` — не удаляйте и не переименовывайте ни его, ни репозиторий.

</details>

## О Хекслете

[Хекслет](https://ru.hexlet.io/) — школа программирования: авторские программы обучения с практикой, поддержкой наставников и реальными проектами, которые остаются в резюме. Этот репозиторий — один из таких проектов.
