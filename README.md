# pgen-simple

Xkcd password generator go wrapper

# Build

```go build -o ~/bin/get-me-a-password main.go```


# Purpose

I spend a good amount of time generating the Xkcd password from the site.
I'll convert this to a go program, so I can also spend some time researching go.


# Structure

[Read this](https://github.com/golang-standards/project-layout)

# Internal

As I'm learning, what I did was to go build first, then I could import.
The usage of capital letter or lower case letter affects the visibility
of the module.

Another thing I did, was to go mod init jvazquez.com/password-gen-xkcd.
After that, I could import the function. 