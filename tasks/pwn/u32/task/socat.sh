#!/bin/bash
while [ 1 ]
do
  socat TCP-LISTEN:1778,reuseaddr,fork EXEC:"./pwn"
done
