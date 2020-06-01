# System Design

Course link: https://www.hiredintech.com/classrooms/system-design/lesson/52

DO NOT:
- Approach questions in a chaotic way and get ratholed, or
- Lack solid understanding of how to properly design architectures that scale.

Use cases:
  - Describe main usecases
    follow up by other usecases
  - ask which usecease to consider
  - ask about HA, fault tolerance, retry, downtime etc requirements.

- Constraints:
  Traffic estimation. no or read/write and storage
  Network I/O:
  - start by asking for a month of requests
  - estimate if its read heavy or write heavy.
  - based on percentage desribution come up with request per second.
  - calculate requests per second using. total requests per month/30 days/24 hours/3600 = ex. 400/s
  - convert into read/write percentage. if 10%read 90%write then: 40 read, 360 write/second

Storage:
 - Decide what objects you are storing
 - avg length of the object.
 - Decide what objects you are creating 
 - decide avg lengh of those objects
 - Decide how long you wanna store them. ex. 5 years.
 - calculate total storage for 5 years : 5 years * 24 months * (100k requests/month * total size of the objects)
 -  x bln bytes =  x GB. ex. 240 billion bytes = 240 GB  
 - 1 GB has six zeros if base unit is byte
 - 1 GB has 3 zeros if base unit is MB (1000 MB = 1 GB)
 - 1000 GB = 1 TB
 througthput: 
 how much data will be read/write per second
 - calculae by  size of all objects * No read per seconds (ex. 100 bytes * 360 read/second) = 36kB/second
- same for write ex. (100 bytes * 40 read/second) = 4kB/second

Verify with the Interviewer

Abstract Design:
- Decide which searvices/workers you will need to handle the usecases
- Decide databases to store what objects

Discuss with the interviewer about the services and database. 
Do not give technology names yet. just say type of DB you can use. 
Do not give how many instances of a serivice you will need at this point. 


Bottlenecks:
- think about how many instances of each servie might require.
- think if it require load balancer, api gateway
- If data needs to be distributed among machines
- if data needs to be cached
- consider trade-offs of the choice and describe impact of those trade offs

- if read requests increases, load increase on db. one idea is to use cache or use replication of the data
- fix bottlenecks from low to high level. from DB, application layer, load balancer
- to fix load balancer use dns between multiple load balancers

-  If we are using a relational database with the proper indexes running on a single machine it could very quickly become unable to handle the loads that our application experiences. One approach mentioned above is to add an in-memory cache solution in front of the database with the goal to not send repeated read requests to the database itself. We could use a key-value store like memcached to handle that.

This will also be really useful for handling situations in which a tweet becomes viral and people start accessing it or the same thing happens to a given user’s profile.

- But this could be insufficient if our data grows too quickly. In such cases we may have to start thinking about partitioning this data and storing it on separate servers to increase availability and spread the load. 


Sharding:
https://medium.com/@jeeyoungk/how-sharding-works-b4dec46b3f6

- If your application is bound by read performance, you can add caches or database replicas. 
- Not everything may need to be sharded. Often times, only few tables occupy a majority of the disk space. Very little is gained by sharding small tables with hundreds of rows. Focus on the large tables.
- Shard or Partition Key is a portion of primary key which determines how data should be distributed. A partition key allows you to retrieve and modify data efficiently by routing operations to the correct database. Entries with the same partition key are stored in the same node. A logical shard is a collection of data sharing the same partition key. A database node, sometimes referred as a physical shard, contains multiple logical shards.
- One way to categorize sharding is algorithmic versus dynamic. In algorithmic sharding, the client can determine a given partition’s database without any help. 
- Algorithmically sharded databases use a sharding function (partition_key) -> database_id to locate data. A simple sharding function may be “hash(key) % NUM_DB”.-
- In dynamic sharding, a separate locator service tracks the partitions amongst the nodes.
- Algorithmic sharding distributes data by its sharding function only. It doesn’t consider the payload size or space utilization. To uniformly distribute data, each partition should be similarly sized. Fine grained partitions reduce hotspots 
- For this reason, algorithmic sharding is suitable for key-value databases with homogeneous values.
- Resharding data can be challenging. It requires updating the sharding function and moving data around the cluster. Doing both at the same time while maintaining consistency and availability is hard.

- In dynamic sharding, an external locator service determines the location of entries
- To read and write data, clients need to consult the locator service first. Operation by primary key becomes fairly trivial.


consistent hashing;
https://www.paperplanes.de/2011/12/9/the-magic-of-consistent-hashing.html



