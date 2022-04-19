# middleware

Simple middleware example with reading request and response body twice (in the middleware and handler). This is
achieved by reading request and then setting it back and the response is using multi writer.
