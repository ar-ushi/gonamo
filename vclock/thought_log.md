initially i thought of using -
https://labix.org/vclock

but i don't think i understand vector clocks well enough to use an abstraction - might as well implement myself

vclocks are data structures which don't use a physical timestamp but assign logical timestamp to each event which are used to determine if events causally related [ancestor/descendant] or concurrent

Each node in the system maintains a vector of counters, one for each node. When an event occurs, the node increments its own counter. When sending messages, nodes include their vector clock. When receiving messages, nodes merge the incoming vector with their own.


what does vclock need to do -
1. create new clock
2. increment it
3. retrieve it
4. find if descendant of
5. find if equal
6. merge
 
