# Debugging a Go executable in a Docker Container with VS Code

This repository illustrates how to debug into a Docker container
running a Go binary.

Prerequisites:

* [Go](https://golang.org/doc/install)
* [Visual Studio Code](https://code.visualstudio.com/)
* [Delve](https://github.com/go-delve/delve/blob/master/README.md), the debugger for Go
* [Docker](https://www.docker.com/get-started) and [Docker Compose](https://docs.docker.com/compose/)

## In a nutshell

1. Build the Go binary (for Linux) by running `make`
2. Build the Docker container by running `docker-compose up --build`
3. Open the code in Visual Studio by running `code .`
4. Set some breakpoints in the code
5. Start `Debug into Docker` in the Debug section of VS Code
6. Open `http://localhost:8080` in your browser
7. See the breakpoints being hit
8. Profit!

## Details

Now, the application is a simply webserver that prints the
current time on every page refresh. By default, it runs on
`http://localhost:8080`.

The `Dockerfile` however won't start the executable. Instead,
it will install the `delve` debugger and start it in
headless mode. In addition to the application port `8080`, it
will open a debugger port at `2345`.

When starting a VS Code debugging session, it will connect to
that port `2345` and will make the debugger start the application.
So now you can hit `http://localhost:8080` and breakpoints 
should work as expected.

Notice that when you stop the debugger, the container will also
be shut down and you might want to restart it.

## References

[1] Remote debugging Go code with VS Code:
    https://github.com/Microsoft/vscode-go/wiki/Debugging-Go-code-using-VS-Code#remote-debugging

# License

MIT.
