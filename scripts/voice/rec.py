#!/sbin/python3
import asyncio
import nats
import json
import os
import cfg

async def main():
    nc = await nats.connect(cfg.NATS_URL, user=cfg.NATS_USER, password=cfg.NATS_PASS)

    print("Recorder Connected")

    rec = os.open("/tmp/openabrec", os.O_RDWR|os.O_NONBLOCK)
    # start recorder
    async def message_handler(msg):
        tp = msg.subject.split(".")
        if tp[2] == "recorder" and tp[3] == "data":
            os.write(rec, msg.data)
            return
        print(tp)
        print(msg.data.decode())

    print("recording")
    sub = await nc.subscribe(f"openab.{cfg.MAC}.recorder.>", cb=message_handler)
    await nc.publish(f"openab.{cfg.MAC}.recorder.command", bytes(json.dumps({"state": 2}), "utf8"))
    print("k")

    await asyncio.sleep(1000000)
    await sub.unsubscribe()
    await nc.drain()

if __name__ == '__main__':
    asyncio.run(main())
