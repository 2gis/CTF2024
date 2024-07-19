#!/bin/bash
while [ 1 ]
do
  socat TCP-LISTEN:1777,reuseaddr,fork EXEC:"./pwn"
done
