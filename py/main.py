import datetime
import time
import celery


app = celery.Celery(
	"publisher",
	broker="amqp://guest:guest@localhost:5545//",
	backend="rpc://"
)


def pub_message():
	task = celery.Task()
	# task.set(queue="echo")
	while 1:
		# to default queue "celery"
		res = task.delay({"message": "this is message", "time": datetime.datetime.now(), "sleep_time": 1})
		print(res)
		print("on sleep")
		time.sleep(10)


if __name__ == "__main__":
	pub_message()