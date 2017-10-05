[![Go Report Card](https://goreportcard.com/badge/github.com/kanapuliAthavan/redistree)](https://goreportcard.com/report/github.com/kanapuliAthavan/redistree)
[![Build Status](https://travis-ci.org/kanapuliAthavan/redistree.svg?branch=master)](https://travis-ci.org/kanapuliAthavan/redistree)
[![codecov](https://codecov.io/gh/kanapuliAthavan/redistree/branch/master/graph/badge.svg)](https://codecov.io/gh/kanapuliAthavan/redistree)


# RedisTree
   A Client library written for the In-Memory database Redis in Go
   
# What is this name RedisTree
   The name came from my extra enthusiasm for environmental friendliness. So all the internal implementation methods and variables of the client bear some name that is related in some way to biological mechanisms. 
   So why call Connect as Connect, Let's put a Redis **Seed()** . Why call Quit as Close(), let it be **SweepWaste()**
   
# Why RedisTree
  - Designed for Redis 2.x and greater versions
  - Usage is simple
  - Marshalling and UnMarshalling of Go Types to hashes
  - Stable Api
  - The client code implementations are simple enough to understand even for the beginner Gophers
  
# Supported Features
   The list of supported methods by the Redistree client is ,
   - **Connection Command Group** - Seed(Connect), Ping, Echo, Close
   - **String Command Group**     - Append, Decr, Incr, Set, Get, StrLen, SetRange, SetNx,
                                    SetEx, SetBit, MSet, MGet, GetSet
   - **Keys Command Group**       - Exists, Del
     More Features coming ...
     
# Installation
    Run the following command in the terminal,
    go get github.com/kanapuliAthavan/redistree
    
# ToDo
  - Connection Pooling
  - Support for Pub Sub, Transactions, List commands, Set and Zset commands
  - Documentation
