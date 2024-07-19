#!/bin/bash
while [ 1 ]
do
  socat TCP-LISTEN:1488,reuseaddr,fork EXEC:"python3 main.py"
done
