import datetime
import celery
from celery import Task

app = celery.Celery(
	"publisher",
	broker="amqp://guest:guest@localhost:5545//",
	backend="rpc://"
)

class CustomTask(Task):
    def __init__(self, name, *args, **kwargs):
        super(Task, self).__init__(*args, **kwargs)
        self.name = name

@app.task(name="add", queue="qwer")
def add(x, y):
    return x + y

def pub_message():
    custom_task = CustomTask("test_task").s().set(queue="qwer")

    # res = add.delay(1, 2)
    # print(res.get())

    res = custom_task.delay({"message": "this is message", "time": datetime.datetime.now(), "sleep_time": 1})#.get()
    print("Задача отправлена")
    print(res)
    print("Встала на ожидание")
    print(res.get())

if __name__ == "__main__":
    pub_message()