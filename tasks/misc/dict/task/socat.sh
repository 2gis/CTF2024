#!/bin/bash
while [ 1 ]
do
  socat TCP-LISTEN:1490,reuseaddr,fork EXEC:"./main"
done
