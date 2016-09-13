Replication
===========

- `dsd` clusters over a shared port (`--cluster_port=46374`) using TCP
- Each node has a unique id (`uid`) which is created on first start
- Each node keeps an incrementing logical timestamp (`lts`)
- Each node which connects to the cluster keeps track of other node's `uid` and last known `lts`
- When a node inserts an object into the distributed set it is given a unique id of `{uid,rts,hsh}`, where:
    - `rts` is a timestamp in GMT
    - `hsh` is the non-cryptographic hash of the object
- When node N′ wants to synchronize with another node N″ in the cluster it will perform the following steps:
    1. N′ will send the last known logical timestamp that N′ has received from N″, denoted T′
    2. N″ sends N′ a list of all object identifiers from T′ onwards noninclusive of T′, denoted I″<sub>1..N</sub>
    3. For each identifier in I″ in I<sub>1..N</sub>:
        1. Select `rts`, `hsh` from I″, denoted R″ and H″ respectively
        2. If `hsh` exists in I′<sub>1..N</sub>, if R′ < R″ discard I″ else queue I″
        3. If `hsh` does not exist in I′<sub>1..N</sub> queue I″ for fetching: F″ = F″ ∪ I″
    4. N′ asks N″ for objects specified in F″, denoted O″
    5. N′ integrates O″ locally and overrides `uid`,`rts` values on conflict
    4. N′ updates last known timestamp of O″
