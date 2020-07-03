#!/usr/bin/env bash
set -e
set -x

req () {
    printf "$1\n"
    shift
    curl -siv "http://127.0.0.1:8080" "$@" | head -n1
    printf "\n"
}

#req "Simple GET /"
req "SQLi in query param" -G --data-urlencode "a=b" --data-urlencode "foo=bar' or 1=1 --"
req "SQLi in body" --data-urlencode "a=b" --data-urlencode "foo=bar' or 'a'='a"
req "XSS in query param" -G --data-urlencode "a=b" --data-urlencode "foo=<script>alert(1);</script>"
req "XSS in body" --data-urlencode "a=b" --data-urlencode "foo=<img onload=\"alert(1);\"/>"
req "File upload" -F "someRandomField=@run-tests.sh"
