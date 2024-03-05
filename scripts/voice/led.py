#!/sbin/python3
import asyncio
import nats
import json
import sys
import cfg

async def main():
    nc = await nats.connect(cfg.NATS_URL, user=cfg.NATS_USER, password=cfg.NATS_PASS)

    if sys.argv[1] == "000000":
        await asyncio.sleep(2)
    await nc.publish(f"openab.{cfg.MAC}.led", bytes(json.dumps({
        "delay": 0,
        "id": 0,
        "sequence": [
            {
                "color": sys.argv[1],
                "duration": 0
            }
        ]
    }), "utf8"))

    await nc.drain()

if __name__ == '__main__':
    asyncio.run(main())
