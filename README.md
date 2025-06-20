# Concurrency_Task

A simple Go project exploring concurrent task execution and dynamic task discovery.

---
##  Описание

Это учебный проект, демонстрирующий:
- Использование `goroutines` и конкурентности в Go
- Применение интерфейсов (`interface`) для абстракции задач
- Динамическое обнаружение новых задач через файловый мониторинг
- Хэширование файлов для отслеживания изменений

---

## Структура проекта
```txt
.
├── cmd
│    └── Concurrency_Task
│        └── main.go
├── go.mod
├── internal
│      ├── tasks
│      │     ├── interface.go
│      │     ├── task_impl
│      │     │   └── task_1.go
│      │     └── task_storage
│      │         └── taskStorage.go
│      └── utils
│          ├── capabilityChecker
│          │     └── Checker.go
│          └── general
│              └── utils.go
├── Makefile
└── README.md

```