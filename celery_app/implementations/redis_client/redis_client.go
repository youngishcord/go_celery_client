package redis_client

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// TODO: У результата в Redis и у получаемой задачи другой формат передачи сообщений.
//  вот пример структуры результата.
//  {"status": "SUCCESS", "result": 3, "traceback": null, "children": [],
//  "date_done": "2025-09-18T21:24:13.006162+00:00", "task_id": "1a07f5af-ee69-4179-83ae-789181db5916"}
//  хотя и без этой структуры результат получаем

// TODO: сделать настройку ttl через конфиг
type RedisClient struct {
	conn *redis.Client

	ttl time.Duration
}

// host string, port string, user string, password string, db int
// Тут можно отдельно настроить размер буфера для чтения и записи
func NewRedisClient() *RedisClient {
	conn := redis.NewClient(&redis.Options{
		Addr:     "localhost:5546",
		Password: "",
		DB:       0,
	})
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := conn.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	return &RedisClient{
		conn: conn,
		//ctx:  ctx,
		ttl: 15 * time.Minute,
	}
}
