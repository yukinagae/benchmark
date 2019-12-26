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

## Case 1

### Setting

- 100 users
- about 3,000 requests

### Results

- python
  - RPS: 66
  - latency (90%ile): 6ms
  - failures: 0%
- go
  - RPS: 65
  - latency (90%ile): 5ms
  - failures: 0%

## Case 2

### Setting

- 1,000 users
- about 50,000 requests

### Results

- python
  - RPS: 500
  - latency (90%ile): 890ms
  - failures: 6%
- go
  - RPS: 610
  - latency (90%ile): 360ms
  - failures: 1%
