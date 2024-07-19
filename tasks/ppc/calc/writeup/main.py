import socket
import time

host = "calc.tasks.2gis.fun"
port = 1488

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((host, port))

while True:
    result = s.recv(1024).strip().decode().replace("\\n", "\n")
    print(result)
    lines = result.split("\n")
    while len(lines) == 1:
        result = s.recv(1024).strip().decode().replace("\\n", "\n")
        print(result)
        lines = result.split("\n")
        time.sleep(1)
    calc = lines[2]
    calc = calc.replace(" - |", "1")
    calc = calc.replace(" ( (_) |", "9")
    calc = calc.replace(" \ O /", "8")
    calc = calc.replace(" _/ /", "7")
    calc = calc.replace(" /_", "6")
    calc = calc.replace(" |__", "5")
    calc = calc.replace(" / o |", "4")
    calc = calc.replace(" / /", "3")
    calc = calc.replace(" __)", "2")
    calc = calc.replace(" | | | |", "0")
    calc = calc.replace(" _| |_", "*")
    calc = calc.replace(" _____", "$")
    alphabet = "1234567890*$"
    execute = ""
    for sym in calc:
        if sym in alphabet:
            execute += sym
    execute = execute.replace("*", "+").replace("$", "-")
    print(execute)
    result = str(eval(execute)).encode()
    print(result)
    s.send(result + b"\n")
