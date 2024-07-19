#!/bin/bash
while [ 1 ]
do
  socat TCP-LISTEN:1489,reuseaddr,fork EXEC:"./main"
done
