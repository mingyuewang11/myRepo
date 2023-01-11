import pytesseract
from PIL import Image
# 读取图片信息
im = Image.open('/home/ts/图片/shazi.png')
# 把从图片识别的文字传给变量，打印到控制台
# string = pytesseract.image_to_string(im)
string = pytesseract.image_to_string(im, lang="chi_sim")
print(string)