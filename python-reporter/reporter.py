import requests
import json

def fetch_data():
    try:
        res = requests.get('http://go-agent:9000/metrics')
        data = res.json()
        print('Fetched metrics:', data)
    except Exception as e:
        print('Failed to fetch:', e)

if __name__ == '__main__':
    fetch_data()
