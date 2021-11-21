# mail

Example of using go smtp library. This example uses gmail smtp. You will need gmail account and generate
[application password](https://support.google.com/accounts/answer/185833?p=InvalidSecondFactor)
(normal user password will fail with `Application-specific password required` message).

## example

 - build `go build`
 - run `./mail`

```go
from email: <user>@gmail.com
password: 
to email: <user>@gmail.com
message: hello
```