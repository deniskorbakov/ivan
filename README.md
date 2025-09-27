# Ivan

![Go Report Card](https://goreportcard.com/badge/github.com/deniskorbakov/ivan)
![Release](https://img.shields.io/github/release/deniskorbakov/ivan?status.svg)
![Action Lint](https://github.com/deniskorbakov/ivan/actions/workflows/lint.yml/badge.svg)
![GitHub Repo stars](https://img.shields.io/github/stars/deniskorbakov/ivan)

Ivan - Auto build and deploy pipline on Go with Ansible

Made using data from packages:

* [cobra](https://github.com/spf13/cobra)
* [fang](http://github.com/charmbracelet/fang)
* [huh](https://github.com/charmbracelet/huh)

## ✨ Install

You will need [Go](https://go.dev/doc/install) and [Ansible](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html) to work.

You can also install the package from the release via the [link](https://github.com/deniskorbakov/ivan/releases)

Local:

clone the repository

```bash
git clone https://github.com/deniskorbakov/ivan.git
````

go to the project folder

```bash
cd ivan
````

build the app

```bash
make build
```

launch the app

```bash
./ivan
```

## 📖 Examples & Usage

The list of commands that you can use in this Build and Deploy 

Утилита корректно работает для следующих типов проектов:


1\) Любой проект на языке Go, PHP, Python, имеющий compose-файл и запускающийся с помощью одной команды `docker compose up -d`. В таком случае стек и прочая специфика проекта не учитывается в Ansible playbook.

Примеры проектов:
- https://github.com/zhaba-team/python-docker-compose-example
- https://github.com/profatsky/freebots

2\) Python-проекты без Docker Compose
В случае если нет возможности запустить проект с помощью Docker Compose, выполняется ручная установка зависимостей (при наличии) через пакетные менеджеры pip или uv. 

2.1) Проект с использованием фреймворков Django или FastAPI.
В случае если проект использует рассматриваемый нами фреймворк, мы учитываем специфику и выполняем запуск: для Django - `python <путь к точке входа> runserver`, для FastAPI - `uvicorn <путь к точке входа>`.

Примеры Django-проектов:
- https://github.com/zhaba-team/django-pip-example
- https://github.com/zhaba-team/django-pip-with-nested-manage-example
Пример FastAPI-проектов:
- https://github.com/zhaba-team/fastapi-uv-example


2.2) Проект, запускаемый стандартным способом.
В случае если проект не использует рассматриваемый нами фреймворк, мы выполняем запуск стандартным способом с помощью `python <путь к точке входа>`.

3\) Go-проекты без Docker Compose
Сборка бинарника и запуск его на сервере с помощью стандартных средств языка.

### 🔌 Build

The command to get rep info and build without ansible

<img src=".assets/build.svg" width="650">

## 🧪 Testing

The command to launch the linter:

```bash
make lint
```

## 🤝 Feedback

We appreciate your support and look forward to making our product even better with your help!

[@Denis Korbakov](https://github.com/deniskorbakov)
