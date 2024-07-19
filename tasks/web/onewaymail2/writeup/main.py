import requests
import urllib.parse
import csv
import datetime
import gzip
import hashlib
import os
import pytz
import urllib.parse
from binascii import unhexlify, hexlify

url = "http://localhost:5007"

def swap_endianness(data: str) -> str:
    return "".join(data[idx:idx + 2] for idx in range(len(data) - 2, -1, -2))


def calculate_block_hash(previous_block_hash: str, block: dict) -> str:
    version = int(block["version"])
    merkle_root = block["merkle_root"]
    nonce = int(block["nonce"])
    bits = int(block["bits"])
    time = datetime.datetime.strptime(block["time"], "%Y-%m-%d %H:%M:%S")
    timestamp = int(time.replace(tzinfo=pytz.UTC).timestamp())

    header_parts = [
        "%08x" % version,
        previous_block_hash,
        merkle_root,
        "%08x" % timestamp,
        "%08x" % bits,
        "%08x" % nonce
    ]
    header_hex = "".join(map(swap_endianness, header_parts))
    header_bin = unhexlify(header_hex)
    calculated_hash_a = hashlib.sha256(header_bin).digest()
    calculated_hash_b = hashlib.sha256(calculated_hash_a).digest()
    return hexlify(calculated_hash_a).decode(), hexlify(calculated_hash_b).decode(), hexlify(
        calculated_hash_b[::-1]).decode()


def find_block(prefix):
    previous_block_hash = None
    for filename in os.listdir("./blocks/"):
        filepath = os.path.join("./blocks/", filename)
        if "gz" not in filepath:
            continue
        with gzip.open(filepath, 'rt') as file:
            tsv_file = csv.DictReader(file, delimiter="\t")
            for block in tsv_file:
                block_hash = block["hash"]
                try:
                    if previous_block_hash is None:
                        continue
                    if not block_hash.startswith('0' * 16):
                        continue
                    calculated_block_hash_a, calculated_block_hash_b, calculated_block_hash_b_reversed = calculate_block_hash(
                        previous_block_hash, block)
                    if calculated_block_hash_a.startswith(prefix):
                        if calculated_block_hash_b_reversed != block_hash:
                            continue
                        return filepath, calculated_block_hash_a, calculated_block_hash_b, block_hash
                finally:
                    previous_block_hash = block_hash


response = requests.post(f"{url}/create", data={"title": "empty", "content": "empty"})
auth = response.cookies.get_dict()["auth"]
id = auth.split("_")[0]
hash_prefix = f"{int(id)+1:x}"
hash_prefix = hash_prefix.zfill(4)
block = find_block(hash_prefix)
filename_hex = block[1][4:]
title = unhexlify(filename_hex)
title_encoded = urllib.parse.quote(title)
response = requests.post(f"{url}/create", data={"title": title, "content": "empty"})
cookies = response.cookies
print(requests.get(f"{url}/", cookies=cookies).text) #Flag in html code