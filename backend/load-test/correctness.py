import base64

from locust import HttpLocust, TaskSet, task, constant, constant_pacing
from random import randint, choice
from os import getenv 
from time import time


class WebTasks(TaskSet):
    def on_start(self):
        self.username = "user%d" %(int(time()*1000.0))
        self.password = "passwd"
        self.client.post("/register", json={
            "username": self.username, 
            "password": self.password, 
            "firstName": "first", 
            "lastName": "last", 
            "email": "user@void.com"})
        base64string = base64.encodestring('%s:%s' % (self.username, self.password)).replace('\n', '')
        self.client.get("/login", headers={"Authorization":"Basic %s" % base64string})
        self.client.post("/addresses", json={"number":"800","street":"Dongchuan","city":"Shanghai","postcode":"200240","country":"China"})
        self.client.post("/cards", json={"longNum":"123443211234","expires":"11/11","ccv":"1234"})

    @task
    def load(self):
        catalogue = self.client.get("/catalogue").json()
        category_item = choice(catalogue)
        item_id = category_item["id"]
        self.client.get("/")
        self.client.get("/category.html")
        self.client.get("/detail.html?id={}".format(item_id))
        self.client.delete("/cart")
        self.client.post("/cart", json={"id": item_id, "quantity": 1})
        base64string = base64.encodestring('%s:%s' % (self.username, self.password)).replace('\n', '')
        self.client.get("/login", headers={"Authorization":"Basic %s" % base64string})
        self.client.get("/basket.html")
        response = self.client.post("/orders")
        if (response.json()["items"][0]["itemId"] != item_id):
            raise Exception("Ordered item and selected item does not match!") 

class Web(HttpLocust):
    task_set = WebTasks
    wait_time = constant_pacing(1)
