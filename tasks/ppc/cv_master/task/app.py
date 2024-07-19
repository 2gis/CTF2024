import hashlib
import os
import random
import time
from threading import Thread
from time import sleep
import jwt
from PIL import Image, ImageFont
from flask import send_from_directory
from flask import Flask, render_template, request, make_response

app = Flask(__name__)

captcha = {}
times = {}

SECRET_KEY = os.getenv("SECRET_KEY")

def sha512(data):
    hash=hashlib.md5(data.encode('utf-8')).hexdigest()
    return hash

@app.route('/static/<path:path>')
def send_report(path):
    return send_from_directory('static', path)

@app.route('/')
def home():
    cookie = request.cookies.get("progress")
    score = 0
    if cookie != None:
        try:
            decoded_jwt = jwt.decode(cookie, SECRET_KEY, algorithms=["HS256"])
        except:
            resp = make_response("corrupted cookie")
            resp.delete_cookie('progress')
            return resp
        score = decoded_jwt["score"]
    random.seed(time.time_ns() - 0xAB23121)
    text = sha512(f"{random.randint(-0x7FFFFFFFFF, 0x7FFFFFFFFF)}{SECRET_KEY}")
    print(f"Generated captcha: {text}")
    font_size = 36
    font_filepath = "./noto.ttc"
    color = (0, 0, 0, 155)
    font = ImageFont.truetype(font_filepath, size=font_size)
    mask_image = font.getmask(text, "L")
    img = Image.new("RGBA", mask_image.size)
    img.im.paste(color, (0, 0) + mask_image.size, mask_image)
    filename = sha512(f'{random.randint(-0x7FFFFFFFFF, 0x7FFFFFFFFF)}{SECRET_KEY}')
    img.save(f"./static/{filename}.png")
    captcha[filename] = text
    times[filename] = time.time()
    return render_template('index.html', static_cv=f"./static/{filename}.png", score=score, captcha_id=filename)


@app.route('/confirm', methods=['POST'])
def create_note():
    solve = request.form['captcha']
    id = request.form['id']
    if id in captcha:
        if captcha[id] == solve:
            if time.time() - times[id] <= 2:
                cookie = request.cookies.get("progress")
                score = 0
                if cookie != None:
                    try:
                        decoded_jwt = jwt.decode(cookie, SECRET_KEY, algorithms=["HS256"])
                    except:
                        resp = make_response("corrupted cookie")
                        resp.delete_cookie('progress')
                        return resp
                    score = decoded_jwt["score"]
                encoded_jwt = jwt.encode({"score": score + 1}, SECRET_KEY, algorithm="HS256")
                del captcha[id]
                del times[id]
                if score + 1 == 100:
                    resp = make_response(os.getenv("FLAG"))
                    resp.delete_cookie('progress')
                    return resp
                resp = make_response("done. +1 score")
                resp.set_cookie('progress', encoded_jwt)
                os.remove(f"./static/{id}.png")
                return resp
            else:
                resp = make_response("time left")
                resp.delete_cookie('progress')
                return resp
    resp = make_response("invalid")
    resp.delete_cookie('progress')
    return resp

def cleaner():
    while True:
        to_del = []
        for id in times:
            if time.time() - times[id] > 20:
                to_del.append(id)
                os.remove(f"./static/{id}.png")
        for id in to_del:
            del captcha[id]
            del times[id]
        sleep(1)


thread = Thread(target=cleaner)
thread.start()
os.makedirs("static", exist_ok=True)
app.run(debug=False, host='0.0.0.0', port=5000)
