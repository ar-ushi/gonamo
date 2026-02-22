thinking about the interface --> 
1. when doing GET, you will always get one value at the least and at max a slice of values when expecting application to perform conflict resolution between different versions. 
2. storage engine imo is the core of the node 


thinking about put -
Dynamo uses vector clocks [12] in order to capture causality
between different versions of the same object. A vector clock is
effectively a list of (node, counter) pairs. One vector clock is
associated with every version of every object. One can determine
whether two versions of an object are on parallel branches or have
a causal ordering, by examine their vector clocks. If the counters
on the first object’s clock are less-than-or-equal to all of the nodes
in the second clock, then the first is an ancestor of the second and
can be forgotten. Otherwise, the two changes are considered to be
in conflict and require reconciliation

1. If my value to be added (VTBA) is an ancestor, ignore it entirely
2. If old value's clock <= VTBA, delete old value
3. Siblings -> add both to the new map and map that to the key 

**IMPLEMENTATION DETAILS**
PUT
1. I've never used mutex locks prior to this but we want to prevent race conditions - concurrent operations is expected here. 
2. TIL defering a mutex unlock is "good practice" --> basically instead of defining a unlock at all returns in a function, we can just use defer to perform the function anytime the scope completes.
3. Core Idea -> reform a map of values every single time a new one is added.
    A. Why? We need to check ancestors every single time since dynamo allows concurrent writes b/w replica nodes
        mapping out a scenario ->
        1. User A works on `branch1` on their local machine
        2. User B also works on `branch1`on their local machine
        start state - branch1 -> {clock: {}}
        User A on Node 1 -> 
        value: "update readme"
        clock: {node1:1}
        User B on Node 2 -> 
        value: "delete xyz"
        clock: {node2:1}
        {node1:1}  ||  {node2:1}
    B. Dynamo is for write heavy situations
    

In practice a key usually has **1–3 siblings**, rarely more. The “rebuild the slice/map” step is effectively **O(k)** where **k is tiny**.

Dynamo also:

• truncates vector clocks (keeps only the most recent N entries)
• relies on read-repair and reconciliation to collapse siblings later
• pushes conflict resolution to clients in some cases


The real latency cost in Dynamo systems is typically **network + quorum coordination**, not the in-memory version filtering.

So reforming the value list is acceptable because it keeps the algorithm **simple, deterministic, and safe under concurrency**, while the bounded sibling count prevents it from becoming expensive.
