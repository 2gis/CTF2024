import random
import time
from colorama import init
import sys
import os
from termcolor import cprint
from pyfiglet import figlet_format

i = 1

while i < 100:
    a = random.randint(1, 999)
    b = random.randint(1, 999)
    sym = "+"
    res = a + b
    sym_c = random.randint(0, 1)
    if sym_c == 0:
        sym = "-"
        res = a - b

    init(strip=not sys.stdout.isatty())
    calc = " "
    for csym in f'{a}{sym}{b}':
        calc += csym + " "
    cprint(figlet_format(calc, font='ntgreek'), attrs=['bold'])

    start = time.time()
    answer = input()
    if int(answer) != res:
        print("wrong!")
        exit(0)
    if time.time() - start > 2:
        print("time left!")
        exit(0)
    time.sleep(1)
    i += 1
print(os.getenv("FLAG"))
