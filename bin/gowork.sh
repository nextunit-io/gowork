#!/bin/sh

SCRIPTPATH="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
GOWORK_ROOT_PATH=$(cd "${SCRIPTPATH}/.." && pwd)

alias mkgowork="${GOWORK_ROOT_PATH}/gowork create"

# This 'hack' is needed because the environment variables needs to be changed.
goworkuse() {
    eval `${GOWORK_ROOT_PATH}/gowork use $@`
}

_gowork_go() {
    eval `${GOWORK_ROOT_PATH}/gowork go $@`
}