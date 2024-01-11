# Protocol used to communicate with the v2 bootcode

Communication is done over a TCP socket on port 5000 on the nabaztag.

## Packet Structure

[Type (len 2); Data] encoded as a UTF-8 string

Max Size is x

## Packet Types

|Type|Description|Data|
|--|--|--|
|01|Reboot||
|02|Ear|Get/Set ear position|[Ear (left: 0, right: 1, both: 2); position (00-16)]|
|03|Led|Set led|[Led (5 digits: 1 to use, 0 to ignore, order: bottom, left, center, right, nose); Color Hex color code (6 characters, 000000 for off); Blinking Speed (0 to stay fixed, or time interval in miliseconds) ]|
|04|Button event|Type (1: click, 2: double click, 3: long click)|
|05|Rfid Read|rfid id|
|06|Recorder|Type (0: stop, 1: start, 2: volume); Volume (only for type 2, percentage on 3 digits)|
|07|Player|Type (0: stop, 1: start, 2: volume); Volume (only for type 2, percentage on 3 digits)|
|08|Record Data|audio chunk|
|09|Player Data|audio chunk|
|10|Wheel status|Value (0-255)|
|12|Play Midi|bytes|
