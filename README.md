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
### Response (tested with [ab tool of Apache](https://httpd.apache.org/docs/2.4/programs/ab.html))
```bash
{
    "totalWords": 351075,
    "totalRequests": 10000,
    "avgProcessingTimeNs": 17528
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
* Run benchmark test
```bash
go test ./api -bench=.
```

## The algorithm
I found out that 2 permutations are equal when you sort them. So on the init DB phase, I created a map[string][]string where the key is a sorted word, and the value is a list of the original words that are permutation of the sorted word.

I looped through the database lines, and for each word (line), I sorted the word, used it as a key and appended the original word to the value (strings list).  

After the initialization of the service, the server is ready to accept and handle requests.

GET similar word request: Before the request, I'm recording the start time of the request using middleware of Gin, Calling c.Next() and starting to handle the request.  
I'm extracting the word from the query param, sorting it, accessing the map, and retrieving the value which is a list of permutations of word.  
I'm removing the word from the list and returning similar:[list,of,words,that,are,similar,to,provided,word] json response, finishing the request handling and the middleware calls the stats update function with the start time.  
Inside the update func, I'm calculating the latency and updating the TotalRequests and AvgProc values with WLock of sync package.

GET stats request: locking RLock (RWMutex of sync package) so that the readers won't block each other,  
and returning the current stats struct.