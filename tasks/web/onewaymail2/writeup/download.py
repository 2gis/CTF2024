import urllib.request

arr = []
with open('slinks.txt') as fl:
    for line in fl:
        arr.append(line)

print(arr)
for link in arr:
    tmp = link.replace("\\n", "").replace("?", "")
    print(f"https://gz.blockchair.com/bitcoin/blocks/{tmp}")
    try:
        urllib.request.urlretrieve(f'https://gz.blockchair.com/bitcoin/blocks/{tmp}', tmp)
    except:
        pass
