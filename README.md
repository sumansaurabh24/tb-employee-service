### Employee service

CRUD operation for employee service

#### APIs Supported

* POST `api/v1/employees`
* GET `api/v1/employees/:id`
* PUT `api/v1/employees/:id`
* DELETE `api/v1/employees/:id`

#### Employee Request & Response

```
{
    "id" : "int",
    "name" : "string",
    "position" : "string",
    "salary" : "flat64",
}
```

#### Error Format

```
{
    "traceId": "UUID",
    "code": 1001,
    "message": "failed fetching employee",
    "info": {
        "error": {
            "Func": "Atoi",
            "Num": "as",
            "Err": {}
        }
    }
}

{
    "traceId": "UUID",
    "code": 2001,
    "message": "entity not found"
}
```

#### Tech Stack & Features

* [Echo golang framework](https://echo.labstack.com/)
* Internal DB with multi-threading support
* Auto-Incremented ids

#### Logs Structure

```
GOROOT=/usr/local/opt/go/libexec #gosetup
GOPATH=/Users/suman.saurabh/go #gosetup
/usr/local/opt/go/libexec/bin/go build -o /Users/suman.saurabh/Library/Caches/JetBrains/GoLand2023.3/tmp/GoLand/___1go_build_github_com_sumansaurabh24_tb_employee_service -gcflags all=-N -l github.com/sumansaurabh24/tb-employee-service #gosetup
/Applications/GoLand.app/Contents/plugins/go-plugin/lib/dlv/mac/dlv --listen=127.0.0.1:56125 --headless=true --api-version=2 --check-go-version=false --only-same-user=false exec /Users/suman.saurabh/Library/Caches/JetBrains/GoLand2023.3/tmp/GoLand/___1go_build_github_com_sumansaurabh24_tb_employee_service --
API server listening at: 127.0.0.1:56125
debugserver-@(#)PROGRAM:LLDB  PROJECT:lldb-1500.0.404.7
 for x86_64.
Got a connection, launched process /Users/suman.saurabh/Library/Caches/JetBrains/GoLand2023.3/tmp/GoLand/___1go_build_github_com_sumansaurabh24_tb_employee_service (pid = 90485).
WARNING: undefined behavior - version of Delve is too old for Go version go1.22.1 (maximum supported version 1.21)

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.12.0
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:8989
{"level":"info","ts":1715420966.599351,"caller":"handlers/employee.go:31","msg":"invoked employee handler"}
{"level":"info","ts":1715420966.599537,"caller":"handlers/request.go:10","msg":"binding request body to object"}
{"level":"info","ts":1715420966.5998452,"caller":"handlers/employee.go:37","msg":"serialized employee value","employee":{"id":0,"name":"Suman Saurabh","position":"LMTS","salary":100000}}
{"level":"info","ts":1715420966.5999098,"caller":"repositories/db.go:51","msg":"db set","value":{"id":1,"name":"Suman Saurabh","position":"LMTS","salary":100000}}
{"level":"info","ts":1715420966.60005,"caller":"tb-employee-service/main.go:37","msg":"request received","uri":"/api/v1/employees","status":201,"host":"","method":"","duration":0,"headers":null,"ip":"","trace_id":""}
{"level":"info","ts":1715420972.105302,"caller":"handlers/employee.go:31","msg":"invoked employee handler"}
{"level":"info","ts":1715420972.105388,"caller":"handlers/request.go:10","msg":"binding request body to object"}
{"level":"info","ts":1715420972.105501,"caller":"handlers/employee.go:37","msg":"serialized employee value","employee":{"id":0,"name":"Suman Saurabh","position":"LMTS","salary":100000}}
{"level":"info","ts":1715420972.105557,"caller":"repositories/db.go:51","msg":"db set","value":{"id":2,"name":"Suman Saurabh","position":"LMTS","salary":100000}}
{"level":"info","ts":1715420972.105669,"caller":"tb-employee-service/main.go:37","msg":"request received","uri":"/api/v1/employees","status":201,"host":"","method":"","duration":0,"headers":null,"ip":"","trace_id":""}
{"level":"info","ts":1715420980.75012,"caller":"services/service.go:39","msg":"get by id","id":1}
{"level":"info","ts":1715420980.75023,"caller":"repositories/db.go:35","msg":"db get","value":{"id":1,"name":"Suman Saurabh","position":"LMTS","salary":100000},"key":1}
{"level":"info","ts":1715420980.750292,"caller":"tb-employee-service/main.go:37","msg":"request received","uri":"/api/v1/employees/1","status":200,"host":"","method":"","duration":0,"headers":null,"ip":"","trace_id":""}
{"level":"info","ts":1715420994.058063,"caller":"handlers/request.go:10","msg":"binding request body to object"}
{"level":"info","ts":1715420994.058339,"caller":"services/service.go:46","msg":"update by id","id":1}
{"level":"info","ts":1715420994.0584168,"caller":"repositories/db.go:65","msg":"db update","value":{"id":0,"name":"Suman Saurabh Updated","position":"LMTS","salary":100000},"key":1}
{"level":"info","ts":1715420994.0584419,"caller":"repositories/db.go:71","msg":"db set","value":{"id":1,"name":"Suman Saurabh Updated","position":"LMTS","salary":100000}}
{"level":"info","ts":1715420994.058511,"caller":"tb-employee-service/main.go:37","msg":"request received","uri":"/api/v1/employees/1","status":200,"host":"","method":"","duration":0,"headers":null,"ip":"","trace_id":""}
{"level":"info","ts":1715421000.956779,"caller":"services/service.go:39","msg":"get by id","id":1}
{"level":"info","ts":1715421000.956872,"caller":"repositories/db.go:35","msg":"db get","value":{"id":1,"name":"Suman Saurabh Updated","position":"LMTS","salary":100000},"key":1}
{"level":"info","ts":1715421000.956917,"caller":"tb-employee-service/main.go:37","msg":"request received","uri":"/api/v1/employees/1","status":200,"host":"","method":"","duration":0,"headers":null,"ip":"","trace_id":""}
{"level":"info","ts":1715421009.4896598,"caller":"services/service.go:53","msg":"delete by id","id":1}
{"level":"info","ts":1715421009.4897642,"caller":"tb-employee-service/main.go:37","msg":"request received","uri":"/api/v1/employees/1","status":204,"host":"","method":"","duration":0,"headers":null,"ip":"","trace_id":""}
{"level":"info","ts":1715421009.4898078,"caller":"handlers/error.go:18","msg":"intercepted in global error handling"}
{"level":"warn","ts":1715421009.489844,"caller":"handlers/error.go:30","msg":"incoming error cannot be parsed into app error"}
{"level":"info","ts":1715421016.534325,"caller":"services/service.go:39","msg":"get by id","id":1}
{"level":"info","ts":1715421016.534451,"caller":"repositories/db.go:35","msg":"db get","value":{"id":0,"name":"","position":"","salary":0},"key":1}
{"level":"info","ts":1715421016.534506,"caller":"tb-employee-service/main.go:37","msg":"request received","uri":"/api/v1/employees/1","status":200,"host":"","method":"","duration":0,"headers":null,"ip":"","trace_id":""}
{"level":"info","ts":1715421016.534534,"caller":"handlers/error.go:18","msg":"intercepted in global error handling"}
{"level":"error","ts":1715421016.534559,"caller":"handlers/error.go:22","msg":"entity not found","stacktrace":"github.com/sumansaurabh24/tb-employee-service/pkg/handlers.(*ErrorHandler).Global\n\t/Users/suman.saurabh/Projects/poc/tb-employee-service/pkg/handlers/error.go:22\ngithub.com/labstack/echo/v4.(*Echo).ServeHTTP\n\t/Users/suman.saurabh/go/pkg/mod/github.com/labstack/echo/v4@v4.12.0/echo.go:675\nnet/http.serverHandler.ServeHTTP\n\t/usr/local/opt/go/libexec/src/net/http/server.go:3137\nnet/http.(*conn).serve\n\t/usr/local/opt/go/libexec/src/net/http/server.go:2039"}
{"level":"info","ts":1715421026.0719302,"caller":"services/service.go:39","msg":"get by id","id":2}
{"level":"info","ts":1715421026.072105,"caller":"repositories/db.go:35","msg":"db get","value":{"id":2,"name":"Suman Saurabh","position":"LMTS","salary":100000},"key":2}
{"level":"info","ts":1715421026.072167,"caller":"tb-employee-service/main.go:37","msg":"request received","uri":"/api/v1/employees/2","status":200,"host":"","method":"","duration":0,"headers":null,"ip":"","trace_id":""}
{"level":"info","ts":1715421044.0010371,"caller":"handlers/employee.go:31","msg":"invoked employee handler"}
{"level":"info","ts":1715421044.001077,"caller":"handlers/request.go:10","msg":"binding request body to object"}
{"level":"info","ts":1715421044.001143,"caller":"handlers/employee.go:37","msg":"serialized employee value","employee":{"id":0,"name":"Suman Saurabh Updated","position":"LMTS","salary":100000}}
{"level":"info","ts":1715421044.001195,"caller":"repositories/db.go:51","msg":"db set","value":{"id":3,"name":"Suman Saurabh Updated","position":"LMTS","salary":100000}}
{"level":"info","ts":1715421044.001249,"caller":"tb-employee-service/main.go:37","msg":"request received","uri":"/api/v1/employees","status":201,"host":"","method":"","duration":0,"headers":null,"ip":"","trace_id":""}
{"level":"info","ts":1715421696.902078,"caller":"handlers/employee.go:31","msg":"invoked employee handler"}
{"level":"info","ts":1715421696.902111,"caller":"handlers/request.go:10","msg":"binding request body to object"}
{"level":"info","ts":1715421696.902163,"caller":"handlers/employee.go:37","msg":"serialized employee value","employee":{"id":0,"name":"Suman Saurabh Updated","position":"LMTS","salary":100000}}
{"level":"info","ts":1715421696.902199,"caller":"repositories/db.go:51","msg":"db set","value":{"id":4,"name":"Suman Saurabh Updated","position":"LMTS","salary":100000}}
{"level":"info","ts":1715421696.902227,"caller":"tb-employee-service/main.go:37","msg":"request received","uri":"/api/v1/employees","status":201,"host":"","method":"","duration":0,"headers":null,"ip":"","trace_id":""}
{"level":"info","ts":1715421698.826187,"caller":"handlers/employee.go:31","msg":"invoked employee handler"}
{"level":"info","ts":1715421698.8262222,"caller":"handlers/request.go:10","msg":"binding request body to object"}
{"level":"info","ts":1715421698.826283,"caller":"handlers/employee.go:37","msg":"serialized employee value","employee":{"id":0,"name":"Suman Saurabh Updated","position":"LMTS","salary":100000}}
{"level":"info","ts":1715421698.826338,"caller":"repositories/db.go:51","msg":"db set","value":{"id":5,"name":"Suman Saurabh Updated","position":"LMTS","salary":100000}}
{"level":"info","ts":1715421698.826402,"caller":"tb-employee-service/main.go:37","msg":"request received","uri":"/api/v1/employees","status":201,"host":"","method":"","duration":0,"headers":null,"ip":"","trace_id":""}
{"level":"info","ts":1715421699.212152,"caller":"handlers/employee.go:31","msg":"invoked employee handler"}
{"level":"info","ts":1715421699.212184,"caller":"handlers/request.go:10","msg":"binding request body to object"}
{"level":"info","ts":1715421699.212234,"caller":"handlers/employee.go:37","msg":"serialized employee value","employee":{"id":0,"name":"Suman Saurabh Updated","position":"LMTS","salary":100000}}
{"level":"info","ts":1715421699.21227,"caller":"repositories/db.go:51","msg":"db set","value":{"id":6,"name":"Suman Saurabh Updated","position":"LMTS","salary":100000}}
{"level":"info","ts":1715421699.212298,"caller":"tb-employee-service/main.go:37","msg":"request received","uri":"/api/v1/employees","status":201,"host":"","method":"","duration":0,"headers":null,"ip":"","trace_id":""}
{"level":"info","ts":1715421699.7005131,"caller":"handlers/employee.go:31","msg":"invoked employee handler"}
{"level":"info","ts":1715421699.700543,"caller":"handlers/request.go:10","msg":"binding request body to object"}
{"level":"info","ts":1715421699.700594,"caller":"handlers/employee.go:37","msg":"serialized employee value","employee":{"id":0,"name":"Suman Saurabh Updated","position":"LMTS","salary":100000}}
{"level":"info","ts":1715421699.700624,"caller":"repositories/db.go:51","msg":"db set","value":{"id":7,"name":"Suman Saurabh Updated","position":"LMTS","salary":100000}}
{"level":"info","ts":1715421699.700651,"caller":"tb-employee-service/main.go:37","msg":"request received","uri":"/api/v1/employees","status":201,"host":"","method":"","duration":0,"headers":null,"ip":"","trace_id":""}
```