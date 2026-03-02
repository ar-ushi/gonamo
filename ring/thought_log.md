read up the paper a bit more - grossly underestimated how little i understood the first time around

the ring is just the hash space and we can define a Q where Q is the total virtual nodes in system 

the ring needs to define the mapping of the hash range to the partition to the actual node 



what i learnt about consistent hashing ->
1. hashing basically allows for flexibility in scaling your servers up/down
2. if your data storage mechanism is dependent on the server itself, you will almost always need to redistribute data 
3. w consistent hashing, the hash space is treated as a circle [hash space is just the values in the possible hash values that will map to your server]
4. why use vnodes instead of treating the servers as is? basic hashing enables random distribution - you don't control the split when servers are added, leading to some servers being underutilized. vnodes is just the load balanced ver of basic consistent hashing -- w basic, you just randomly assign positions instead we assign multiple v positions (tokens) to physical server. w that each node has higher chances of equal distribution. 
5. When a server fails, its load is distributed across many servers, not just one neighbor.

we'll use the ring as our layer to track the total number of partitions for each node/server and track total partitions which form our ring


each partition is a section of the hashspace taken up by a virtual node, it is assigned a token which is a random position 