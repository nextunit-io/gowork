#!/bin/sh

alias mkgowork="./gowork create"

# This 'hack' is needed because the environment variables needs to be changed.
# alias goworkuse="eval `./gowork use $@`"

goworkuse() {
    eval `./gowork use $@`
}
