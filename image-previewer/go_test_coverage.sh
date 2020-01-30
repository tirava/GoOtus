#!/bin/bash

go test -cover `go list ./... | grep -v integration_tests` | awk '{if ($1 != "?") print $5; else print "0.0";}' | sed 's/\%//g' | awk '{s+=$1; if ($1 != "0.0") s2+=1} END {printf "coverage: %.2f%% of statements\n", s/s2}'


