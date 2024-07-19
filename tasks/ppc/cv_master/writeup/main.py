import re
import requests
import pytesseract
from PIL import Image

host = "cv-master.tasks.2gis.fun"

url = f"https://{host}"

last_success_cookie = {}
while True:
    response = requests.get(f"{url}/", cookies=last_success_cookie).text
    score = re.findall(r"score: \d+", response)
    print(score[0])
    images = re.findall(r"\./static/\w{32}\.png", response)

    image_bytes = requests.get(f"{url}/{images[0]}").content

    with open("image.png", "wb") as binary_file:
        binary_file.write(image_bytes)

    print(images[0])
    image = Image.open('image.png')
    captcha = pytesseract.image_to_string(image, config="-c tessedit_char_whitelist=0123456789abcdef")
    print(captcha)
    response = requests.post(f"{url}/confirm", data={"id": re.findall(r"\w{32}", images[0])[0], "captcha": captcha.replace("\n", "")}, cookies=last_success_cookie)
    print(response.text)
    if "2GIS.CTF" in response.text:
        exit(0)
    elif "done" in response.text:
        last_success_cookie = {"progress": response.cookies.get_dict()["progress"]}
