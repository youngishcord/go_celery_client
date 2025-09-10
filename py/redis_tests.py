import datetime
import random
import time
import celery
import numpy as np
from celery import Task
from celery.result import ResultSet
import redis_tests

app = celery.Celery(
    "publisher",
    broker='redis://localhost:5546/0',
    backend='redis://localhost:5546/0'
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

@app.task(name="test", queue="asdf")
def test(message):
    print(message)
    return "YES"

@app.task(name="counter", queue="qwer")
def counter(c):
    return c+1

def pub_message():

    # chain
    # t1 = CustomTask("test_task1").s().set(queue="qwer")
    # t2 = CustomTask("test_task2").s().set(queue="asdf")
    # t3 = CustomTask("test_task3").s().set(queue="qwer")

    # ch = t1 | t2 | t3

    # ch.delay("q1w2e3r4")

    ####################################################
    ####################################################
    # Много задач с ожиданием
    ####################################################
    ####################################################
    # hub = []
    # ch = celery.chain(add.s(), counter.s())
    # for i in range(10):
    #     res = ch(1, 2)# .get()
    #     print(res)
    #     hub.append(res)
    #
    # res = ResultSet(hub).join()
    # print(res)
    ####################################################
    ####################################################
    ####################################################
    ####################################################


    ####################################################
    ####################################################
    # Кастомная задача
    ####################################################
    ####################################################

    res = add.apply_async((1, 2,), ignore_result=True)
    # while 1:
    #     time.sleep(5)
    # print(res.get())

    # ret = test.delay("это тестовое сообщение")

    # while 1:
    #     time.sleep(random.random())
    #     custom_task = CustomTask("test_task").s().set(queue="qwer")
    #     res = custom_task.delay({"message": "this is message", "time": datetime.datetime.now(), "sleep_time": 1})#.get()

    #     time.sleep(random.random())
    #     custom_task = CustomTask("test_task").s().set(queue="asdf")
    #     res = custom_task.delay(1, 2, 3)#.get()
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