#!/bin/bash
while [ 1 ]
do
  socat TCP-LISTEN:1495,reuseaddr,fork EXEC:"./a.out"
done
