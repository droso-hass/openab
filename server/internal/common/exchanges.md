
Client      Server

Play "url" -> play ""
Play Stop -> play_stop
Player Volume -> set_volume

Record -> start_rec
Record End -> stop_rec

Move Ear -> ear_move(num pos dir)
Set Led -> set_led(color/time sequence)

Get Image -> snapshot
Image <- nab

Reboot -> reboot

Button (1,2,3) <- nab
Rfid <- nab
Record Data <- nab
Ear Move <- nab
Connected/Disconnected <- nab