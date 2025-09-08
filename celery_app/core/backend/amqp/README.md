Что именно ожидает Celery

Брокер/очередь задач. По умолчанию задачи публикуются в очередь celery (можно роутить иначе через exchange/routing key).
docs.celeryq.dev

Сообщение задачи (AMQP). Заголовки содержат минимум:

id (UUID задачи),

task (имя задачи),

другие поля состояния;
Тело протокола v2 — список: [args, kwargs, embed], где args — позиционные аргументы, kwargs — именованные.
docs.celeryq.dev
+1

RPC-результат. При включённом result_backend="rpc://" клиент Celery при отправке задачи указывает свойства AMQP reply_to (имя свою «callback» очереди) и correlation_id (обычно task_id). Воркеры должны ответить в очередь reply_to и проставить тот же correlation_id. Это стандартный шаблон RabbitMQ RPC и именно так устроен Celery RPC-бэкенд.
docs.celeryq.dev
+1
rabbitmq.com

Форма полезной нагрузки ответа. Celery ожидает сериализованный JSON с метаданными состояния, например:

{"status":"SUCCESS","result": <ваш_результат>, "traceback": null, "task_id": "<id>", "children": []}