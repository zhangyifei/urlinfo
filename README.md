# Urlinfo Project
Url lookup service

Urlinfo is used provide a Url lookup Api service for validating a malware url and it also provide methods to update its malware url database. 


## Architecture

![Architecture](doc/../docs/images/architecture.svg "Architecture")

- Ngnix 
  
    The ngnix server is used as a load balancer for the application servers. Besides that, I also use it to do the rate limiting for controlling the traffic. 

- Application Servers
    
    The Url lookup service instances, created by using Go and Go-zero framework. We could deploy several instances to handle a large number of requests. Inside the application, the rate limiting logic is also added to control the requests. 

    Two kinds of apis provided in this service:
       
        1.  Url lookup 
        2.  Update or Batch update the Urls 

    Swagger API json [file](docs/urlinfo.json)


- Redis

    The redis instance or cluster is used to do the DB cache

- Mongo DB Shard Cluster

    In order to store data and make the queries efficiently, I create a Mongo DB Shard Cluster as the storage part. The Shard Cluster can handle large number of Url Data. Each Shard in the cluster is a replicat set and I use the read preference to split the Read and Write requests. This mechanism can ensure the performance of data reading and writing. 

## How it works


## Setting up a demo


## Future work