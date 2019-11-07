# Map

One of the most useful data structure in computer science is the hash table. Many hash table implementations exist with varying properties, but in general they offer fast lookups, adds, and deletes.

|   Operations  |                       |
|---------------|-----------------------|
| Construct     | m := map[key]value{}  |
| Insert        | m[k] = v              |
| Lookup        | v = m[k]              |
| Delete        | delete(m, k)          |
| Iterate       | for k, v := range m   |
| Size          | len(m)                |

Key may be any type that is comparable and value may be any type at all, including another map!

## Hash function

Choose a bucket for each key so that entries are distributed as evenly as possible.

## How Maps are Structured

Maps in Go are implemented as a hash table. The hash table for a Go map is structured as an array of buckets. The number of buckets is always equal to a power of 2. When a map operation is performed, such as (colors["Black"] = "#000000"), a hash key is generated against the key that is specified. In this case the string "Black" is used to generate the hash key. The low order bits (LOB) of the generated hash key is used to select a bucket.
