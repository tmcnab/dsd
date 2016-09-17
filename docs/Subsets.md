Subsets
=======
Subsets are parts of the larger set differentiated by a _query_. Subsets can be either **_stored_** or **_ad-hoc_** determined by how they are executed. A client may create a stored query (static subset?) of which the objects it contains are computed ahead of time to maximize speed. Subsets which are computed dynamically (ad-hoc querie => dynamic subset?) will be subject to performance penalties due to a full scan of the master set.


Queries
-------
Queries operate on a sequence of objects and follow standard set theory operators. The following examples refer to a theoretical database where a variety of objects have been previously added.

To get down to basics, a query is one statement of one or more expressions that return zero or more objects. Clients can assign expressions to identifiers which are just printable letters &dagger; and re-use the expressions to build a complicated result. Comments are just lines that do not start with an assignment or the return statement.

```
Create expression of objects where they're type Person and location not null
people : {} ∩ {type = 'Person'} ∩ {location ≠ ø}

Create a "locations" expression
locations : {} ∩ {type = 'Location'}

Find all locations in California (zip codes between 90000 and 96100)
california : locations ∩ {zip > 90000} ∪ locations ∩ {zip > 96100}

Return all Californians (people with location intersecting locations)
people ∩ { location ∩ california }

```

**Intersection** clauses allow the client to select objects which match a given condition. For example, imagine we want to select all the people who are employed:

```
a : S ∩ { {type: 'Person'} }
b : a ∩ { {employed: true} }
// or
=> S ∩
```

**Union** clauses allow the client to join previously computed subsets together into another set.
