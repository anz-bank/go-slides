# Webserver Comparison - Java Spring Boot vs Go

The purpose of this project is to demonstrate Go's "out-of-the-box" concurrency.  
It implements a simple webserver in Java Spring Boot and Go respectively with a single route:
`/ping` which returns `pong` after a two second delay.

## Java Spring Boot implementation

The Java implementation is located at [/java](java) and was built with

- JDK 1.8
- Spring Boot 2.1.2
- Gradle 5.0

Build and run from the command line with

    cd java
    ./gradlew clean bootRun

and test it with

    curl localhost:8080/ping

## Go implementation

The go implementation is located at [/go](go) and was built with

- go 1.11
- go modules

Build and run from the command line with

    cd go
    go run main.go

and test it with

    curl localhost:9090/ping

## Benchmark

A benchmark has been created with Apache's [`ab`](https://httpd.apache.org/docs/2.4/programs/ab.html) command,
install with

    brew install httpd

After starting Java Spring Boot and go webservers in separate terminals as described above the following response times stats were created with

    ab -n 7000 -c 7000 -s 200 localhost:8080/ping  # Java
    ab -n 7000 -c 7000 -s 200 localhost:9090/ping  # go

with `-n` referring to the number of total requests, `-c` the number of requests started concurrently and `-s` the timeout in seconds.

### Detailed response times

On a 2017 MacBook Pro (2.9 GHz Intel Core i7, 16 GB 2133 MHz LPDDR3) `ab` reported the following
response time stats when issuing 7,000 concurrent requests:

#### Java Spring Boot implementation

```
Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0  328  99.4    335     470
Processing:  2007 36010 20196.8  36009   70075
Waiting:     2006 36009 20195.4  36009   70075
Total:       2475 36338 20097.8  36344   70223

Percentage of the requests served within a certain time (ms)
  50%  36344
  66%  48270
  75%  54246
  80%  58213
  90%  64190
  95%  68162
  98%  70153
  99%  70160
 100%  70223 (longest request)
```

#### Go implementation

```
Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0  267  74.9    266     390
Processing:  2001 2576 312.8   2769    2931
Waiting:     2001 2576 312.9   2769    2931
Total:       2388 2843 245.0   3019    3124

Percentage of the requests served within a certain time (ms)
  50%   3019
  66%   3030
  75%   3035
  80%   3039
  90%   3049
  95%   3055
  98%   3066
  99%   3073
 100%   3124 (longest request)
```

### Response time summary

7,000 concurrent requests:

```
Java mean: 36.3 sec
Go mean:    2.8 sec
```
