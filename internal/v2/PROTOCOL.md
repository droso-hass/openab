# Protocol used to communicate with the v2 bootcode

Communication is done over a TCP socket on port 5000 on the nabaztag.

## Packet Structure

[Type (len 2); Data] encoded as a UTF-8 string

Max Size is x

## Packet Types

|Type|Description|Data|
|--|--|--|
|00|Ping|ping|
|01|Reboot||
|02|Ear|[Ear (left: 0, right: 1); Start After (start sequence only after x ms);position (00-16); direction (0/1); Time (interval in miliseconds);]|
|03|Led|[Led (id from 0 to 4: nose, left, center, right, bottom); Start After (start sequence only after x ms); Color Hex color code (6 characters, 000000 for off); Time (0 to stay fixed, or time interval in miliseconds); ]|
|04|Button event|Type (1: click, 2: double click, 3: long click)|
|05|Rfid Read|rfid id|
|06|Recorder|Type (0: stop, 1: start)|
|07|Player|Type (0: stop, 1: start, 2: volume); Volume (only for type 2, percentage); Linkto media (only for type 1)|
|08|Record Data|audio chunk|
|09|Play Midi|bytes|
|10|Wheel status|Value (0-255)|
|11|Ear State Change|[Ear ID; position (0-16)]|


Led example:

03;1;1500;00FF00;200;00FF11;300;00FF22;400;00FF33;500

Ear example: 
02;1;0;7;1;0