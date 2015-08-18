#! /bin/bash

curl "http://192.168.22.128/hello.php"  > /dev/null

sleep 1

curl "http://192.168.22.128/run-tests.php"  > /dev/null