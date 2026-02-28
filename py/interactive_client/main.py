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
    broker="pyamqp://admin:admin@localhost:5672//",
    backend = "rpc://",
)

x = Task("add").s().set(queue="test")
x.delay(1, 2)# .get() # -> 3
x.delay(1, 2)# .get() # -> 3
x.delay(1, 2)# .get() # -> 3
x.delay(1, 2)# .get() # -> 3
x.delay(1, 2)# .get() # -> 3

# con = code.interact(banner="Welcome to celery tester client!", local=locals())
# con.interact()