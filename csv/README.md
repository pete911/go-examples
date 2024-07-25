# CSV

Example of custom tags (csv) and how to marshal using reflection.

 - run example `go run .`
 - slice of structs with csv tags is converted to CSV byte array

```
name,other_names,worker,created_at
Ryan,John Jane,true,2004-04-20T15:38:57+01:00
John,Josh Ryan,false,2019-12-20T15:38:57Z
Emily,Naomi Joseph,true,1992-03-05T15:38:57Z
```