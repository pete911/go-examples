# binary

## left shift
```
'1 << 0' -   1
'1 << 1' -  10
'1 << 2' - 100
```

```go
out := byte(0b0001) << byte(0b0000)
fmt.Printf("%04b\n", out)
out = byte(0b0001) << byte(0b0001)
fmt.Printf("%04b\n", out)
...
```

## right shift
```
'4 >> 0' - 100
'4 >> 1' -  10
'4 >> 2' -   1
```

```go
out := byte(0b0100) >> byte(0b0000)
fmt.Printf("%04b\n", out)
out = byte(0b0100) >> byte(0b0001)
fmt.Printf("%04b\n", out)
...
```

## bitwise operators
### AND
`1010 & 1001 - 1000`

```
1010
1001
----
1000
```

```go
out := byte(0b1010) & byte(0b1001)
fmt.Printf("%04b\n", out)
```

### OR
`1010 | 1001 - 1011`

```
1010
1001
----
1011
```

```go
out := byte(0b1010) | byte(0b1001)
fmt.Printf("%04b\n", out)
```

### XOR
`1010 ^ 1001 - 11`

```
1010
1001
----
0011
```

```go
out := byte(0b1010) ^ byte(0b1001)
fmt.Printf("%04b\n", out)
```

### clear (AND NOT)
`1010 &^ 1001 - 10`

```
1010
1001
----
0010
```

```go
out := byte(0b1010) &^ byte(0b1001)
fmt.Printf("%04b\n", out)
```

## setting bits
`bit |= 1 << index`

x is 0
 - set bit indexed 0 `x |= 1 << 0 - 1`
 - set bit indexed 1 `x |= 1 << 1 - 10`
 - set bit indexed 2 `x |= 1 << 2 - 100`
 
```go
out := byte(0b0000)
out = out | byte(0b0001) << byte(0b0000)
fmt.Printf("%04b\n", out)
out = byte(0b0000)
out = out | byte(0b0001) << byte(0b0001)
fmt.Printf("%04b\n", out)
...
```

## clearing bits
`bit &= ^(1 << index)`

x is 3 (11)
 - clear bit indexed 0 `x &= ^(1 << 0) - 10`
 - clear bit indexed 1 `x &= ^(1 << 1) - 1`
 
```go
out := byte(0b0011)
out = out & ^(byte(0b0001) << byte(0b0000))
fmt.Printf("%04b\n", out)
out = byte(0b0011)
out = out & ^(byte(0b0001) << byte(0b0001))
fmt.Printf("%04b\n", out)
```

## check if bit is set
`bit & (1 << index) != 0`
 - if true, bit at index 0 is set `x & (1 << 0) != 0`
 - if true, bit at index 1 is set `x & (1 << 1) != 0`
 
```go
out := byte(0x0001) & (byte(0x0001) << byte(0x0000)) != byte(0x0000)
fmt.Printf("%t\n", out) // true, bit at index 0 is set
out = byte(0x0001) & (byte(0x0001) << byte(0x0001)) != byte(0x0000)
fmt.Printf("%t\n", out) // false, bit at index 1 is not set
```
