rm /tmp/openabsnd /tmp/openabrec
mkfifo /tmp/openabsnd
mkfifo /tmp/openabrec
dir=$(pwd)

./rec.py &
REC_PID=$!
./play.py &
PLAY_PID=$!

trap "kill $REC_PID; kill $PLAY_PID; exit" SIGINT

cd wyoming-satellite && script/run \
 --name "testsat" \
 --uri "tcp://0.0.0.0:10700" \
 --mic-command 'cat /tmp/openabrec' \
 --mic-command-rate 16000 \
 --mic-noise-suppression 4 \
 --debug \
 --snd-command-rate 44100 \
 --snd-command 'dd of=/tmp/openabsnd' \
 --detection-command "/sbin/python3 $dir/led.py ff0000" \
 --stt-stop-command "/sbin/python3 $dir/led.py 0000ff" \
 --tts-start-command "/sbin/python3 $dir/led.py 00ff00" \
 --tts-played-command "/sbin/python3 $dir/led.py 000000"

#--debug-recording-dir '/home/thomas/Documents/dev/openab/doc/rec' \
#--wake-uri 'tcp://127.0.0.1:10400' \
#--wake-word-name 'perroquet' \
#--snd-command 'aplay -r 22050 -c 1 -f S16_LE -t raw' \

# docker run -it -p 10400:10400 rhasspy/wyoming-porcupine1:latest --debug &
