Engine
======

The underlying Set engine has eight primary operations it executes:
 - Insert an object
 - Ignore an object
 - Create a subset
 - Destroy a subset
 - Execute a query
 - Objects (hashes) inserted since logical timestamp
 - Fetch Objects by hash
 - Execute third-party module

These operations are all that are needed to operate.


1. Insert an object
-------------------
This section outlines the general algorithm that the engine follows when inserting an object into the store.

1. Sort the object by key, denoted **O**
2. Compute a hash of the object, **H**
3. Lock the metadata file **F<sub>M</sub>**
4. Write **O** data to end of the object file **F<sub>O</sub>**, retaining:
    4.1. **L** the length of the object data in bytes
    4.2. **S** the offset from the start of **F<sub>O</sub>** where data was written
5. Write metadata **H , M** to **F<sub>M</sub>**, where:
    5.1. **T<sub>R</sub>** the real-world time in UTC
    5.2. **T<sub>L</sub>** the logical timestamp
    3.3. **M = { T<sub>L</sub> , T<sub>R</sub> , L , S }**
6. Increment **T<sub>L</sub>**
7. Unlock **F<sub>M</sub>**
8. Return **H , T<sub>R</sub>** to client.


2. Ignore an object
-------------------
This section outlines the algorithm that the engine follows when directed to ignore an object in the store.
