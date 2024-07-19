#!/bin/bash
while [ 1 ]
do
  socat TCP-LISTEN:1492,reuseaddr,fork EXEC:"./main"
done
