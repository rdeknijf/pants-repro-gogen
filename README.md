# `go-generate` problem reproduction repo

This repo is an example that reproduces the following:

    $ pants go-generate ::

    13:56:18.02 [INFO] Completed: Process `go generate` directives in file: pkg/something.go
    13:56:18.02 [ERROR] 1 Exception encountered:

    Engine traceback:
    in `go-generate` goal

    ProcessExecutionFailure: Process 'Process `go generate` directives in file: pkg/something.go' failed with exit code 1.
    stdout:

    stderr:
    2024/10/09 13:56:18 Loading input failed: GOPATH is not set

    Use `--keep-sandboxes=on_failure` to preserve the process chroot for inspection.

Basically `go-generate` doesn't work.