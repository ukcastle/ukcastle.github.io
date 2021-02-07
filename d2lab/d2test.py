import requests
import cv2 
import urllib.request
import numpy as np

'''
static
'''

#dictionary to hold extra headers
HEADERS = {"X-API-Key":'f6f6701ecc654545b354e5d99d72cd04'}

#make request for Gjallarhorn
r = requests.get("https://www.bungie.net/platform/Destiny/Manifest/InventoryItem/1274330687/", headers=HEADERS)

#convert the json object we received into a Python dictionary object
#and print the name of the item
inventoryItem = r.json()

Gjallarhorn = inventoryItem['Response']['data']['inventoryItem']

imgPath = 'http://www.bungie.net'+Gjallarhorn['icon']


def url_to_image(url):
    # download the image, convert it to a NumPy array, and then read
    # it into OpenCV format
    resp = urllib.request.urlopen(url)
    image = np.asarray(bytearray(resp.read()), dtype="uint8")
    image = cv2.imdecode(image, cv2.IMREAD_COLOR)
    # return the image
    return image

#main
def main():
    img = url_to_image(imgPath)
    if img.any():
        print(img)
        cv2.imshow('image',img)
        cv2.waitKey(0)


if __name__ == "__main__":
    main()

