worker:
	celery -A main worker --loglevel=INFO --pool=solo -Q qwer --concurrency=1 --prefetch-multiplier=1

worker2:
	celery -A main worker --loglevel=INFO -Q qwer
