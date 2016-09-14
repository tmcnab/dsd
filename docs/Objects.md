Objects
=======

Using the classic set theory terminology, objects are individual values in the
set. Objects are accepted as a dictionary of key:value pairs, where values may
be:
    - binary
    - boolean
    - dictionary
    - number
    - sequence
    - string

This might look familiar because it's basically a restricted subset of JSON,
without the null type. Key components of any dictionary are encoded as UTF-8
strings. Value components immediately follow their keys.

`binary`
--------
Binary values are encoded as:
    - first byte containing the data type (binary, 0x01)
    - unsigned 56-bit integer containing the length (bytes) of the binary blob
    - binary data

`boolean`
---------
Boolean values are encoded as a single byte value with the most-significant
bit indicating truth (0 → false, 1 → true).

`dictionary`
------------
TBA.

`number`
--------
Number values are encoded as:
    - first byte containing the data type (number, 0x08)
    - eight bytes containing the floating number itself

`sequence`
----------
TBA. 0x10

`string`
--------
String values are encoded as:
    - byte containing the data type (string: 0x20)
    - seven bytes containing the length of the string
    - N bytes of UTF-8 string data
