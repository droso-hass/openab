# OpeNab

An opensource nabaztag/nabaztag:tag/karotz server.

## Getting Started

Make sure that you have cloned this repo with the option `--recurse-submodules`.

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

Go into `scripts/voice`:
 - install the requirements: `pip3 install -r requirements.txt`
 - clone [wyoming-satellite](https://github.com/rhasspy/wyoming-satellite/)
 - run `cd wyoming-satellite && script/setup`
 - make sure to also install the webrtc utils `cd wyoming-satellite && .venv/bin/pip3 install 'webrtc-noise-gain==1.2.3'`
 - turn on your nabaztag and wait for it to connect to the server
 - run `./voice.sh`
