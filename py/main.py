import datetime
import time
import celery
from celery import Task
from celery.result import ResultSet

app = celery.Celery(
	"publisher",
	broker="amqp://guest:guest@localhost:5545//",
	backend="rpc://"
)
# app.conf.worker_prefetch_multiplier = 1  # Только одна задача на процесс
# app.conf.worker_concurrency = 1          # Количество процессов
# app.conf.worker_optimization = 'fair'    # Включаем честное распределение

class CustomTask(Task):
    def __init__(self, name, *args, **kwargs):
        super(Task, self).__init__(*args, **kwargs)
        self.name = name

@app.task(name="add", queue="qwer")
def add(x, y):
    # time.sleep(1)
    return x + y

@app.task(name="counter", queue="qwer")
def counter(c):
    return c+1

def pub_message():
    
    
    
    ####################################################
    ####################################################
    # Много задач с ожиданием 
    ####################################################
    ####################################################
    hub = []
    ch = celery.chain(add.s(), counter.s())
    for i in range(10):
        res = ch(1, 2)# .get()
        print(res)
        hub.append(res)
    
    res = ResultSet(hub).join()
    print(res)
    ####################################################
    ####################################################
    ####################################################
    ####################################################
    
    
    ####################################################
    ####################################################
    # Кастомная задача
    ####################################################
    ####################################################
    # custom_task = CustomTask("test_task").s().set(queue="qwer")

    # res = add.delay(1, 2)
    # print(res.get())
    
    res = custom_task.delay({"message": "this is message", "time": datetime.datetime.now(), "sleep_time": 1})#.get()
    # print("Задача отправлена")
    # print(res)
    # print("Встала на ожидание")
    # print(res.get())
    ####################################################
    ####################################################
    ####################################################
    ####################################################

if __name__ == "__main__":
    pub_message()