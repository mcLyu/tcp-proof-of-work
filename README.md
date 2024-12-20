# TCP client-server app with Proof-of-Work hashcash implementation 

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://github.com/mcLyu/tcp-proof-of-work/actions)

This repository contains an golang solution of task about TCP client-server application with DDoS protection via proof-of-work
hashcash implementation. The server returns quote from 'Book of Wisdom' book if the client has solved a proof-of-work challenge.

## Task description
```
Design and implement “Word of Wisdom” tcp server.
• TCP server should be protected from DDOS attacks with the Proof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.
• The choice of the POW algorithm should be explained.
• After Proof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.
• Docker file should be provided both for the server and for the client that solves the POW challenge
```

## Clone the project

```
$ git clone https://github.com/mcLyu/tcp-proof-of-work.git
```

## Requirements
+ [Go 1.23+](https://go.dev/dl/) 
+ [Docker](https://docs.docker.com/engine/install/) 
+ [Make utility](https://www.gnu.org/software/make/) in case you want to use the Makefile

## Run the project

```
make build
```
Executes docker-compose build command to build the server and client images.

```
make start
```
Executes docker-compose up command to start the server and client containers.

```
make start_ddos n={number_of_clients}
```
Executes docker-compose up command to start the server and {number_of_clients} client containers to emulate ddos attack.

## Description
### Choice of PoW algorithm
The PoW algorithm is based on the [hashcash](https://en.wikipedia.org/wiki/Hashcash) algorithm. This algorithm was chosen because it's most popular and used in BTC PoW algorithm.
There were other options like [Cuckoo Cycle](https://medium.com/codechain/cuckoo-cycle-c337e30c6c99) or [Equihash](https://en.bitcoinwiki.org/wiki/Equihash), but they are less popular so I decided to choose hashcash.

Concrete algorithm implementation is based on the [article](https://therootcompany.com/blog/http-hashcash/). 
In  this article HTTP headers are used to pass the PoW challenge and solution, 
but in this implementation, the PoW challenge and solution are passed as a part of the TCP message.

### DDoS protection
Main goal of 'DDoS protection' here is to show how PoW hashcash can be used to protect from DDoS attacks.
This repository contains a simple implementation of DDoS protection based on the number of active connections, which increases difficulty of PoW challenge.  
It uses active connection number instead of RPS just to make local tests easier and show impact of increasing of PoW difficulty.