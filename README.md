Console JSON formatter with query feature.

Install:

```$ go get github.com/miolini/jsonf```

Usage:

```
Usage of jsonf:
  -c=true: colorize output
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
 
![Output](https://cdn.rawgit.com/miolini/jsonf/master/output.png "Output")

 
```
$ echo '{"uid":1,"email":"user@gmail.com","address":{"city":"New-York","country":"US"}}' \
  | jsonf -q 'value["address"]["country"]'
```

```
"US"
```
