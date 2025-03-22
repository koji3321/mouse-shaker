import pyautogui as pya
from pynput.keyboard import Listener,Key
import threading

#pya.move(xOffset=-5,yOffset=0,duration=0.5,)
#pya.dragTo(x=1920/2,y=1080/2,duration=5)

asd=True

def mose():
    while asd:
        pya.move(5,0)
        pya.move(-5,0)

a=threading.Thread(target=mose)
a.start()

def naber(key):
    try:
        if key==Key.esc:
            global asd
            asd=False
            return False
    except AttributeError:
        pass

with Listener(on_press=naber) as listener:
    listener.join()

a.join()