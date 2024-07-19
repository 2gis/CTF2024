#!/bin/bash
while [ 1 ]
do
  socat TCP-LISTEN:1800,reuseaddr,fork EXEC:"./main"
done
