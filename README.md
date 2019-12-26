# benchmark

## Usage

```bash
locust -f ./locust/bench.py
```

### API Start

- python

```bash
cd json/python
poetry install
poetry shell
gunicorn -w 4 -k uvicorn.workers.UvicornWorker main:app
```

- go

```bash
cd json/go
go build main.go
./main
```

## Setting

- 1,00 users
- about 30,00 requests

## Results

- python
  - RPS: 66
  - latency (90%ile): 6ms
  - failures: 0%
- go
  - RPS: 65
  - latency (90%ile): 5ms
  - failures: 0%
