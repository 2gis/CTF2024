from qreader import QReader
import cv2

for i in range(1, 138):
    qreader = QReader()
    image = cv2.cvtColor(cv2.imread(f"./qrs/{i}.png"), cv2.COLOR_BGR2RGB)
    decoded_text = qreader.detect_and_decode(image=image)
    print(decoded_text[0], end='')
