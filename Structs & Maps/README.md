# Structs VS Maps
|                     **Structs**                     |                            **Maps**                           |
|:---------------------------------------------------:|:-------------------------------------------------------------:|
|            All keys must be the same type           |                              ---                              |
|           All values must be the same type          |                Values can be of different types               |
|     Keys are indexed so we can iterate over them    |                  Keys don't support indexing                  |
| Use to represent a collection of related properties | Use to represent a "thing" with a lot of different properties |
|   Don't need to know all the keys at compile time   |   You need to know all the different fields at compile time   |
|                   Reference type!                   |                          Value Type!                          |

### Test it out by running this command `go run main.go`