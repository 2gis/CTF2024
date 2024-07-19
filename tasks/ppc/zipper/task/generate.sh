 #!/bin/bash

i=0
while [ "$i" -le 137 ]; do
    last="qrs/$i.png"
    password=$(head /dev/urandom | tr -dc 1-9 | head -c 4 ; echo '')
    zipfile="zips/archive_$i.zip"
    zip -P $password $zipfile $last
    last=$zipfile
    i=$(( i + 1 ))
    echo "Архив $zipfile создан с паролем $password"
done
