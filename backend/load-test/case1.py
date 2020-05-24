import base64

from locust import HttpLocust, TaskSet, task, constant_pacing, constant
from locust.contrib.fasthttp import FastHttpLocust
from random import randint, choice
from os import getenv 
from time import time
import sys


class WebTasks(TaskSet):
    @task
    def load(self):
        self.client.get("/items/holy_1.jpg")

class Web(FastHttpLocust):
    wait_time = constant_pacing(0.5)
    task_set = WebTasks
