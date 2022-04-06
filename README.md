# Similar words generator

Similar words generator is a Golang backend service for generating similar words.

## Installation and Run
### Setting up the service
* Download [Golang programming language](https://go.dev/doc/install).
* Clone repository
```bash
git clone https://github.com/ranavital/pan_task.git
```

* Move into the project directory:
```bash
cd pan_task
```

* Build the service
```bash
go build
```

* Copy your words list database (txt file) into the main repository folder (same folder as main.go)

* Change local.json (in config folder) to your configuration (set the database file name to 'db_path' value).


### Running the service
Run:
```bash
./pan_task
```


## Usage
### Get similar words
for host=localhost and port=8000:
```bash
curl http://localhost:8000/api/v1/similar?word=apple
```
### Response
```bash
{"similar":["appel","pepla"]}
```
The case where there are no similar words:
```bash
curl http://localhost:8000/api/v1/similar?word=nosimilarwords
```
### Response
```bash
{
    "similar": []
}
```
### Stats on the server for 'similar' request
for host=localhost and port=8000:
```bash
curl http://localhost:8000/api/v1/stats
```
### Response
```bash
{
    "totalWords": 351075,
    "totalRequests": 1,
    "avgProcessingTimeNs": 57296
}
```

## Tests
### Run tests:
* Change test.json (in config folder) to your configuration (set the database file name to 'db_path' value).
* Simple test run:
```bash
go test ./api -v
```
* Run test with coverage
```bash
go test ./api -v -cover
```

