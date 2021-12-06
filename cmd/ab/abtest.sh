#!/bin/bash

ab -v 1 -n 100000 -c 125 localhost:8080/
