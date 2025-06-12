import requests
import time
import json
import logging
from queue import Queue
from threading import Thread
from requests.exceptions import RequestException

logging.basicConfig(level=logging.INFO, format='%(asctime)s %(levelname)s:%(message)s')

class MessageQueue:
    def __init__(self):
        self.queue = Queue()

    def send(self, data):
        self.queue.put(data)
        logging.info(f"Data queued: {data}")

    def process(self):
        while True:
            data = self.queue.get()
            if data is None:
                break
            # Simulate sending data to Kafka or RabbitMQ
            logging.info(f"Processing data: {data}")
            time.sleep(1)  # Simulate processing delay
            self.queue.task_done()

def fetch_metrics():
    url = "http://go-agent:9000"
    for attempt in range(5):
        try:
            response = requests.get(url, timeout=2)
            response.raise_for_status()
            data = response.json()
            logging.info(f"Fetched metrics: {data}")
            return data
        except RequestException as e:
            logging.warning(f"Attempt {attempt+1}: Failed to fetch metrics: {e}")
            time.sleep(2 ** attempt)  # exponential backoff
    logging.error("Failed to fetch metrics after retries")
    return None

def transform_data(data):
    data['MemoryGB'] = round(data.get('Memory', 0) / (1024**3), 2)
    data['ProcessedTimestamp'] = time.time()
    return data

if __name__ == "__main__":
    mq = MessageQueue()
    worker = Thread(target=mq.process, daemon=True)
    worker.start()

    while True:
        raw_data = fetch_metrics()
        if raw_data:
            processed_data = transform_data(raw_data)
            mq.send(processed_data)
        time.sleep(10)  