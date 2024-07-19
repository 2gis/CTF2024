import qrcode
flag = "2GIS.CTF{cf6e596105238ac5bfb2a483c8989f28efeb884a29900e798e290c6662259af10256897fe1c44b319eb2a570c5ec329b1170f0f3ad9b67b163f4ebd2e0627fa5}"
ind = 0
for i in flag:
    img = qrcode.make(i)
    img.save(f"qrs/{ind}.png")
    ind += 1