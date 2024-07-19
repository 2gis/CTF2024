#!/bin/bash

i=137
while [ "$i" -ge 1 ]; do
    j=1000
    while [ "$j" -le 9999 ]; do
        zipfile="zips/archive_$i.zip"
        unzip -u -P $j $zipfile
        if [ $(echo $?) -ne "0" ]; then
            tmp=$(( i - 1 ))
            rm -rf "archive_$tmp.zip"
        else
            break
        fi
        j=$(( j + 1 ))
        echo "Пароль $j"
    done
    i=$(( i - 1 ))
done
