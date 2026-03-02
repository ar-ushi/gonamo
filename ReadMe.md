# Layer 0 — Project Skeleton

Goal: runnable node process.

Decide:

* node configuration format
* logging
* metrics
* request protocol (HTTP / gRPC / TCP)

A node should start and expose a **health endpoint**.

Checklist

* node ID
* port
* config loader
* graceful shutdown
* structured logs

Test

* start 3 nodes locally

---

# Layer 1 — Local Storage Engine

Goal: single-node KV store.

Features

* put
* get
* delete
* multiple versions per key

Concepts to support

* siblings
* opaque values
* version metadata

Test cases

* write / read
* overwrite
* store two conflicting versions

Do **not distribute yet**.

---

# Layer 2 — Versioning Model

Goal: represent causality.

Add

* vector clocks
* comparison logic

Operations

* increment
* merge
* descends-from
* concurrent

Test

Cases:

```
A -> B
A || B
merge(A,B)
```

If this layer is wrong the whole system breaks.

---

# Layer 3 — Consistent Hash Ring

Goal: deterministic data ownership.

Add

* node hash
* key hash
* clockwise lookup
* replication preference list

Features

* virtual nodes
* stable ordering

Test

* add node
* remove node
* same key always maps same node

---

# Layer 4 — Cluster Membership

Goal: nodes know about each other.

Start simple:

* static node list
* periodic health checks

Then later upgrade to gossip.

Track

* node id
* address
* status

Test

* node joins
* node leaves
* node marked down

---

# Layer 5 — Request Routing

Goal: any node can accept requests.

Add concept:

**coordinator node**

Workflow

```
client → random node
node → finds replicas
node → coordinates operation
```

Test

* send request to wrong node
* ensure correct nodes handle data

---

# Layer 6 — Replication

Goal: durability and availability.

Introduce:

```
N = replication factor
```

Coordinator sends writes to N nodes.

Test

* kill one node
* data still available

---

# Layer 7 — Quorum Logic

Goal: tunable consistency.

Add

```
R = read quorum
W = write quorum
```

Rules

```
W + R > N → strong consistency
```

Test matrix

```
node slow
node down
partial success
```

---

# Layer 8 — Conflict Handling

Goal: allow concurrent writes.

System must:

* detect concurrency
* store siblings
* return multiple values to client

Client merges.

Test

Simulate partition:

```
write A
disconnect nodes
write B
reconnect
```

Both versions must exist.

---

# Layer 9 — Read Repair

Goal: replicas converge naturally.

On reads:

1. compare versions
2. send updates to stale replicas

Test

* intentionally corrupt one replica
* read should fix it

---

# Layer 10 — Failure Detection

Goal: cluster resilience.

Track

* heartbeats
* timeouts
* suspicion states

Test

* kill process
* ensure routing avoids dead nodes

---

# Layer 11 — Sloppy Quorum

Goal: maintain availability during failures.

If a replica is down:

* write to next healthy node on ring.

This temporarily violates placement rules.

Test

* shut down primary replica
* write still succeeds

---

# Layer 12 — Hinted Handoff

Goal: restore correct replica placement.

Temporary node stores a **hint**.

When original node returns:

```
replay hints
```

Test

* node down
* writes occur
* node back
* data transferred

---

# Layer 13 — Anti-Entropy

Goal: repair silent divergence.

Add background sync using **Merkle trees**.

Process

```
compare hash trees
locate divergent ranges
sync keys
```

Test

* manually modify replica
* anti-entropy fixes it

---

# Layer 14 — Gossip Protocol

Goal: scalable membership.

Nodes periodically:

```
pick random node
exchange state
merge updates
```

Properties

* eventually consistent
* scalable

Test

* start node late
* it learns cluster automatically

---

# Layer 15 — Operational Safety

Goal: production behavior.

Add

* backpressure
* timeouts
* hinted handoff limits
* clock pruning
* tombstones
* metrics

---

# Recommended Build Timeline

```
1  storage
2  vector clocks
3  consistent hashing
4  replication
5  quorum
6  conflicts
7  failure handling
8  repair mechanisms
```

Everything after quorum is where Dynamo becomes interesting.

---

# Milestone goals

### Milestone 1

Single node KV with vector clocks.

### Milestone 2

Cluster with replication.

### Milestone 3

Quorum reads/writes.

### Milestone 4

Partition tolerant system.

### Milestone 5

Self-healing cluster.

---

# Difficulty spikes (expect these)

Most implementations get stuck at:

1. vector clock comparisons
2. sloppy quorum routing
3. Merkle tree repair
4. gossip convergence

---

# One mental model that helps

Dynamo is basically:

```
consistent hashing
+ versioned storage
+ quorum voting
+ background repair
```

Not a traditional database.

---

If you want, I can also show the **mistakes almost everyone makes when implementing Dynamo the first time** (they're surprisingly consistent).
