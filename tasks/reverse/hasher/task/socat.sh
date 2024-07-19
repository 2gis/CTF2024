#!/bin/bash
while [ 1 ]
do
  socat TCP-LISTEN:1822,reuseaddr,fork EXEC:"./main"
done
