# Golang general guidelines

## panic and recover
 One possible counterexample is during initialization: if the library truly cannot set itself up, it might be reasonable to panic, so to speak.
 
 When panic is called, including implicitly for run-time errors such as indexing a slice out of bounds or failing a type assertion, it immediately stops execution of the current function and begins unwinding the stack of the goroutine, running any deferred functions along the way. If that unwinding reaches the top of the goroutine's stack, the program dies.
 
 A call to recover stops the unwinding and returns the argument passed to panic.
**Because the only code that runs while unwinding is inside deferred functions, recover is only useful inside deferred functions**

One application of recover is to shut down a failing goroutine inside a server without killing the other executing goroutines.

```go
func server(workChan <-chan *Work) {
    for work := range workChan {
        go safelyDo(work)
    }
}

func safelyDo(work *Work) {
    defer func() {
        if err := recover(); err != nil {
            log.Println("work failed:", err)
        }
    }()
    do(work)
}
```
In this example, if do(work) panics, the result will be logged and the goroutine will exit cleanly without disturbing the others. There's no need to do anything else in the deferred closure; calling recover handles the condition completely.


## context

## Go HTTP server
- create flag
```go
import "flag"

var addr = flag.String("addr", ":8080", "description)
func main(){
flag.Parse()
fmt.Println(*addr)
}
```
- create server
```go
import "net/http"

func main(){
  http.Handle("/", http.HandlerFunc(MyFunc))
  err := http.ListenAndServe(":8080", nil)
  if err != nil{
    log.Fatalf(err)
  }
}
func MyFunc(w http.ResponseWriter, r *http.Request){
}

```


```go

package main

import (
    "flag"
    "html/template"
    "log"
    "net/http"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

var templ = template.Must(template.New("qr").Parse(templateStr))

func main() {
    flag.Parse()
    http.Handle("/", http.HandlerFunc(QR))
    err := http.ListenAndServe(*addr, nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
func QR(w http.ResponseWriter, req *http.Request) {
    templ.Execute(w, req.FormValue("s"))
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET">
    <input maxLength=1024 size=70 name=s value="" title="Text to QR Encode">
    <input type=submit value="Show QR" name=qr>
</form>
</body>
</html>
`
```

## interface


## inheritance

## embedded types

## pass by value Vs reference


## error
```
type error interface {
    Error() string
}
```
