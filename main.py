import json
import asyncio
import websockets
# import pyautogui
import mouse
import pydirectinput

pydirectinput.PAUSE = 0.0
# pyautogui.PAUSE = 0
# pyautogui.FAILSAFE = False
# mousex = pyautogui.position().x
# mousey = pyautogui.position().y
mousex = mouse.get_position()[0]
mousey = mouse.get_position()[1]
lastx = None
lasty = None

async def handler(websocket):
    global mousex, mousey, lastx, lasty
    SENSITIVITY = 4 * 0.1
    async for message in websocket:
        event = json.loads(message)
        if event['type'] == 'touchstart':
            lastx = int(event['x'] / SENSITIVITY)
            lasty = int(event['y'] / SENSITIVITY)
        elif event['type'] == 'touchmove':
            mousex += int((event['x'] / SENSITIVITY) - lastx)
            mousey += int((event['y'] / SENSITIVITY) - lasty)
            pydirectinput.moveTo(mousex, mousey)
            # mouse.move(mousex, mousey, absolute=True, duration=0.0)
            lastx = int(event['x'] / SENSITIVITY)
            lasty = int(event['y'] / SENSITIVITY)
        elif event['type'] == 'touchend':
            lastx = None
            lasty = None

async def main():
    async with websockets.serve(handler, "0.0.0.0", 8001):
        await asyncio.Future() # run forever

if __name__ == "__main__":
    asyncio.run(main())
