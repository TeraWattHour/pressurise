# Pressurise

**Pressurise is the Go fullstack framework with no unnecessary bloat.**

This repository contains the client for interfacing with generated Pressurise handlers.
The CLI is available [here](https://github.com/terawatthour/pressurise-cli).

This project is in early stages of development, any input on how to make it better is very much welcome.

## Concept

The main goal of this library is to simplify development of full-stack apps
in Go. To accomplish this we use the powerful `html/template` package together
with some simple code generation. Really, this native library is all you need
in most of the cases. Not many languages offer such a great tool, do they?

How your project looks is very similar to Astro javascript apps - but in Go.
Every `.html` file in the `app/` directory is mapped to its generated route - that's all what happens under the hood.

## Usage

- Install the CLI `go install github.com/terawatthour/pressurise-cli`
- Add the client to your module `go get github.com/terawatthour/pressurise`
- Move into your project's base directory `cd base_project_path`
- Build the routes `pressurise-cli build .`
- Use the generated handlers

```go
func main() {
    app := NewPressurise()
    app.RegisterPages(pressHandlers)
    app.Run("localhost:8080")
}
```

## App Directory

- app/layout.htm - bare in mind that files that are not routes may not
  contain code blocks

  ```html
  <!DOCTYPE html>
  <html lang="en">
    <head>
      <meta charset="UTF-8" />
      <meta name="viewport" content="width=device-width, initial-scale=1.0" />
      <title>Document</title>
    </head>
    <body>
      {{ block "content" . }} fallback {{ end }}
    </body>
  </html>
  ```

- app/index.html - this route will map to url `/` and will be an extension of
  `./layout.htm` HTML template

  ```
  ---
  // extends command takes in one argument which is a relative path
  // to the extended template, this file may not be a route
  !extends ./layout.htm

  import (
      "fmt"
  )

  // this is executed on every request,
  // in this code block w (http.ResponseWriter),
  // r (*http.Request) and every structure
  // declared in the same package `main` are available

  fmt.Println("hello from my page", r.Method)

  // `method` is available in the template
  // whilst r.Method is directly not,
  // in short: you can use every structure declared
  // in this code block in your template below
  method := r.Method

  ---

  {{ define "content" }}
    <h1>This was accessed with method: {{ .method }} </h1>
  {{ end }}

  ```

## TODOs

- Add tests (who are we kidding I'm not doing that)
- Add docs
