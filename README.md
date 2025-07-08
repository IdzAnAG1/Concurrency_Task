# Concurrency Task Runner (Go)

> - Механизм отслеживания изменений в директории с автоматическим выполнением задач.  
> - Реализовано на Go с использованием `goroutines`, `channels`, интерфейсов и динамической инициализации.

---

##  Возможности

-  **Отслеживание изменений файлов и структуры директории**
-  **Динамическое выполнение задач при изменениях**
-  **Гибкая система регистрации задач через интерфейсы**
-  **Проверка изменений по содержимому файлов (MD5 hash)**
-  **Concurrency-first архитектура (goroutines + channels)**

---

## 🛠️ Установка и запуск

1. Убедитесь, что у вас установлен Go (1.20+)

2. Клонируйте репозиторий:

```bash
    git clone https://github.com/your-username/concurrency-task.git
    cd concurrency-task
```
3. Запустите проект
```text
    make run
```

---

## Как добавить новую задачу
> На данный момент реализация добавления задачи очень сырая, но происходит это следующим образом
1. Создаем файл с расширением `.go` в `internal/tasks/task_impl/` 
2. Реализуем интерфейс `ConcurrencyTask`
3. Реализуем функцию init для того чтобы задача помещалась в пространство программы (В дальнейшем планируется реализовать автоматическую реализацию функции init)
4. Задача автоматически подтягивается в очередь на исполнение

---

## Как это работает

     capabilityChecker сканирует директорию с задачами (task_impl) каждые N миллисекунд

     Сохраняется хэш каждого файла (по содержимому)

     При появлении новых файлов или изменении — в канал отправляется сигнал

     main.go слушает канал и запускает задачи, зарегистрированные в task_storage

---


## 📁 Структура проекта
```text
.
├── cmd
│     └── Concurrency_Task
│         └── main.go
├── go.mod
├── go.sum
├── internal
│        ├── config
│        │       └── config.go
│        ├── file_verifier
│        │       ├── change_detector
│        │       │       └── ChaD.go
│        │       ├── file_readiness_detector
│        │       │   └── FiReD.go
│        │       ├── injection_of_function_init
│        │       │       └── InFinit.go
│        │       └── Verifier.go
│        ├── models
│        │       └── models.go
│        ├── tasks
│        │       ├── interface.go
│        │       ├── task_code_storage
│        │       │       └── TaskCode.go
│        │       ├── task_impl
│        │       │       └── task_1.go
│        │       └── task_storage
│        │           └── taskStorage.go
│        ├── utils
│        │       ├── file_handler
│        │       │       └── file_handler.go
│        │       ├── go_uuid
│        │       │       └── uuid.go
│        │       ├── hash
│        │       │       └── hash.go
│        │       └── regex
│        │           └── regexp.go
│        └── variables
│            └── variables.go
├── Makefile
└── README.md


```
