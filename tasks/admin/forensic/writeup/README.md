Если мы войдём в виртуалку, при попытке ввода `ls` система будет нас выкидывать.

Просмотрим файлы через многократное нажатие TAB и увидим файл shadow в домашней директории.

В описании задания, про слабый root пароль, поэтому попробуем сбрутить его

hashcat -m 1800 -a 0 hash.txt rockyou.txt

Получаем root пароль `sunshine`

При входе в root, нам сообщают, что для получения флага, нужно ввести "pwnd"

Вводим и получаем флаг.