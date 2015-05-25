Console JSON formatter with query feature.

Install:

```$ go get github.com/miolini/jsonf```

Usage:

```
$ jsonf -h
Usage of jsonf:
  -d=false: debug output to stderr
  -f=true: format output json to stdout
  -q="": json query
  -s=true: syntax hightlight
  -v=false: verbose output to stderr
```

Examples:

```
$ echo '{"uid":1,"email":"user@gmail.com","address":{"city":"New-York","country":"US"}}' | jsonf
```
 
![ScreenShot](https://cdn.rawgit.com/miolini/jsonf/master/output.png "Screenshot")

 
```
$ echo '{"uid":1,"email":"user@gmail.com","address":{"city":"New-York","country":"US"}}' | jsonf -q 'value["address"]["country"]'
"US"
```
