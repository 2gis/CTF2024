from pwn import *

con = remote("numbers.tasks.2gis.fun", 1777)

sum = 0
for i in range(8):
	sum += u64(con.recv(8))

sum &= 0xffffffffffffffff
con.send(p64(sum))

print(con.recv())
