from pwn import *

con = remote("u32.tasks.2gis.fun", 1778)

num_bytes = con.recv(4)
num = str(u32(num_bytes))
con.send(num + "\n")

print(con.recv())