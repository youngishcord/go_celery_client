# Простой пример с интерактивной консолью

import code
from celery import Celery
from celery import Task as BaseTask

class Task(BaseTask):
    def __init__(self, name, *args, **kwargs):
        super(BaseTask, self).__init__(*args, **kwargs)
        self.name = name

app = Celery(
    "interactive_test",
    broker="amqp://admin:admin@localhost:5672//",
    backend = "rpc://",
)

# x = Task("add").s().set(queue="qwer")
# t = x.delay(1, 2)# .get() # -> 3
# print(t.get())
# x.delay(1, 2)# .get() # -> 3
# x.delay(1, 2)# .get() # -> 3
# x.delay(1, 2)# .get() # -> 3
# x.delay(1, 2)# .get() # -> 3

add = Task("add").s().set(queue="qwer")
panic = Task("panic").s().set(queue="qwer")
inf = Task("inf").s().set(queue="qwer")
err = Task("err").s().set(queue="qwer")

# inf.apply_async(time_limit=5)

con = code.interact(banner="Welcome to celery tester client!", local=locals())
con.interact()