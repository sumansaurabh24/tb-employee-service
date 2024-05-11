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