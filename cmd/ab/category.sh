#!/bin/bash

ab -v 1 -l -n 10000 -c 30 http://localhost:8000/c/supermarket/11
