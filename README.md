# url-shortener — HTTP service for URL shortening 
### Running the application without Docker
```
go run .
```
## HTTP Methods
```
/ping — Checking the server connection
```
```
/generate — Generate a short URL
options: 
    url — URL which must be shortened
    input — Short link text (optional)
    
example: http://localhost:8080/generate?url=google.com&input=ggl
output:  example.com/ggl
```
```
/get — Getting the original url from the short
options: 
    url — Short url 
    
example: http://localhost:8080/get?url=example.com/ggl
output:  google.com
```
### Params for ```.env``` file
```
URL=example.com/
HOST=localhost
PORT=8080
MONGO=mongodb://localhost:27017
```
