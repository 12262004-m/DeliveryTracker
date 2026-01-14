<p align="center">
  <img src="img/delivery_banner.svg" alt="Delivery Tracker banner"/>
</p>

<div align="center">
  <div>
    <img src="https://img.shields.io/badge/Backend-Go-002A48?style=for-the-badge&logo=go&logoColor=white"/>
    <img src="https://img.shields.io/badge/Frontend-HTML%20%7C%20Tailwind-000026?style=for-the-badge&logo=tailwindcss&logoColor=white"/>
  </div>
</div>

---

## Описание
**Delivery Tracker** — проект для отслеживания заказов доставки в реальном времени. Он реализует полный цикл работы с заказом: создание заказа клиентом, управление заказами администратором, мгновенные обновления статуса без перезагрузки страницы и система уведомлений

---

## Стек

<h4> FRONT-END (HTML templates, Tailwind CSS, EventSource (SSE))</h4>

[![My Skills](https://skillicons.dev/icons?i=html,js,tailwind)](https://skillicons.dev)

<h4> BACK-END (Go, Gin, Server-Sent Events (SSE), Sessions, Concurrency (sync.RWMutex))</h4>

[![My Skills](https://skillicons.dev/icons?i=go)](https://skillicons.dev)

<h4> DB (SQLite, normalized schema, relations, CRUD)</h4>

[![My Skills](https://skillicons.dev/icons?i=sqlite)](https://skillicons.dev)

---

## Функциональные особенности

- Real-time обновление данных без WebSocket (Server-Sent Events)
- Собственный `NotificationManager` для работы с SSE
- Потокобезопасная система уведомлений
- Разделение ролей: клиент / администратор
- Серверный рендеринг шаблонов
- Чистая и расширяемая backend-архитектура

---

## Пользовательские роли

### Client
- Создание нового заказа через форму
- Просмотр статуса заказа
- Автоматическое обновление страницы при изменении статуса (SSE)
- Просмотр подробной информации о заказах

### Admin
- Авторизация в админ-панели
- Просмотр всех активных заказов
- Изменение статуса заказа пользователя
- Удаление заказов
- Система уведомлений о новых заказах в реальном времени
- Отображение времени последнего обновления

---

## Запуск проекта
После запуска приложение будет доступно по адресу `http://localhost:8080`
```bash
go run ./cmd
```

---

## Интерфейс

### Админ-панель
<p align="center">
  <img src="img/admin.png" alt="Admin"/>
</p>

### Страница отслеживания заказа
<p align="center">
  <img src="img/client.png" alt="Customer order page"/>
</p>