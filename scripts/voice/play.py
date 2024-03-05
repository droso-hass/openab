#!/sbin/python3
import asyncio
import json
import nats
import os
from time import time
import cfg

async def main():
    nc = await nats.connect(cfg.NATS_URL, user=cfg.NATS_USER, password=cfg.NATS_PASS)

    print("Player Connected")

    async def blank():
        for i in range(30):
            await send(b'0'*1024)

    async def send(audio):
        inbox = nc.new_inbox()
        sub = await nc.subscribe(inbox)
        await nc.publish(f"openab.{cfg.MAC}.player.data", audio, reply=inbox)
        msg = await sub.next_msg()
        if msg.data != b'':
            await asyncio.sleep(0.01)
            await send(audio)

    snd = os.open("/tmp/openabsnd", os.O_RDONLY|os.O_NONBLOCK)
    data = b''
    lastread = time()
    was_streaming = False
    while True:
        try:
            r = os.read(snd, 4)
            if r != b'':
                if not was_streaming:
                    # stop the recorder while we send audio
                    await nc.publish(f"openab.{cfg.MAC}.recorder.command", bytes(json.dumps({"state": 1}), "utf8"))
                    await blank()
                was_streaming = True
                lastread = time()
                data += r
                if len(data) >= 1024:
                    await send(data)
                    data = b''
            elif was_streaming:
                if time()-lastread > 1:
                    print("End of stream")
                    if len(data) > 0:
                        await send(data)
                    await blank()
                    was_streaming = False
                    break
        except BlockingIOError:
            pass

    await nc.drain()

if __name__ == '__main__':
    asyncio.run(main())
