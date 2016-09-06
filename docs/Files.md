Files
=====

All database files live in the $CWD/.data directory when dsd is opened. All
files are binary and inspection via text editor will be fruitless.

| Filename | Description                                                     |
| -------- | --------------------------------------------------------------- |
| data     | Object data file. Each entry consists of a 64-bit length and a UTF-8 string |
| store    | Index file of object hashes to file byte offsets in `data`      |
| *.index  | Query index files. Stores both the matching objects and the predicates that were used in initializing them |
