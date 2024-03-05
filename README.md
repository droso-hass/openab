# OpeNab

An opensource nabaztag/nabaztag:tag/karotz server.

## Getting Started

Make sure that you have cloned this repo with the option `--recurse-submodules`.

Warning: This project is in a very early stage, it has not been fully tested and is not (yet) recommended to use for an extended period of time.

Ports:
- tcp/80: For V1/V2 (to download bootcode)
- udp/4000: For V2 (to send/receive audio)

v2 TODO:
- Implement button long click
- Fix ear positionning
- Fix player

Global TODO:
- V3 Implementation
- V1 Implementation
- API Stabilization
- Home Assistant Integration

### Testing the voice satellite (highly experimental and hacky)

Make sure that you have python and docker installed. We assume that you are running a linux system.

Run `scripts/nats.sh` to run the nats server.

Run `make run` to build the openab server and run it.

Press the head of the nabaztag:tag while powering it (all leds should be blue), connect to the new wifi (NabaztagXX) and go to `192.168.0.1` and at the bottom of the page, change the `Violet Platform` field to the address of the pc running the server (ex: `http://192.168.1.102/vl`), then click `update and start`.

Go into `scripts/voice`:
 - install the requirements: `pip3 install -r requirements.txt`
 - clone [wyoming-satellite](https://github.com/rhasspy/wyoming-satellite/)
 - run `cd wyoming-satellite && script/setup`
 - make sure to also install the webrtc utils `cd wyoming-satellite && .venv/bin/pip3 install 'webrtc-noise-gain==1.2.3'`
 - turn on your nabaztag and wait for it to connect to the server
 - run `./voice.sh` (currently you need to restart this script, the server and the nab after each voice command)
