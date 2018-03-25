#!/bin/sh

alias mkgowork="./gowork create"

# This 'hack' is needed because the environment variables needs to be changed.
goworkuse() {
    eval `./gowork use $@`
}
