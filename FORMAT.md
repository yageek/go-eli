# Endianness
All integer values are stored as little-endian in the ELI file.

# ELI File
The _ELI_ format represents a list or _records_.

## | Start of file (0x00) |

| Record 0
| --------
| Record 1
| Record 2
| ...
| Record N

Depending on its type, a record does not have the same size.

At the beginning of an ELI file, we have the following sequence of records:

| Start of file (0x00)
| -------------------------------
| File Descriptor
| Physical Parameters
| Rate Resolution
| Calibration Configuration
| Calibration Configuration Fixed
| A4 Bilinear
| Time Date

### Record
#### General structure

Field   | Size    | Signedness
------- | ------- | ----------
Type    | 1 byte  | unsigned
Length  | 1 byte  | unsigned
Payload | M bytes |

#### Record Length
To obtain the correct record length as 4 bytes integer, you have to proceed the value in the `Length` field as follow:

```
// Compute base and exp
var length = N
var base = length & 0x3f
var exp = length >> 6
var recordLength = base * Math.pow(6.0, exp)
// Use recordLength as record length
```

#### Record Types

Type Name                       | Type Byte
------------------------------- | ---------
File Descriptor                 | 0x01      |
Physical Parameters             | 0x11      |
Rate Resolution                 | 0x21      |
Calibration Configuration       | 0x22      |
Calibration Configuration Fixed | 0x26      |
A4 Bilinear                     | 0x23      |
Time Date                       | 0x31      |
2D Position Delay               | 0x41      |
2D Position Distance            | 0x52      |
2D Position Coordination        | 0x61      |
Data Test Point                 | 0x02      |
Thrown Points                   | 0xa1      |
Timeout                         | 0xb1      |
Status Change                   | 0xf1      |
Unknown                         | 0xff      |

### Record Definition
#### File Descriptor

Field       | Size   | Value | Signedness
----------- | ------ | ----- | ----------
Type        | 1 byte | 0x01  | unsigned
Length      | 1 byte | 0x06  | unsigned
File Type   | 1 byte |       | unsigned
Pos Data    | 1 byte |       | unsigned
Cancel Flag | 1 byte |       | unsigned
Reserved    | 1 byte |       | unsigned

#### Physical Parameters

Field      | Size    | Value | Signedness
---------- | ------- | ----- | ----------
Type       | 1 byte  | 0x11  | unsigned
Length     | 1 byte  | 0x09  | unsigned
Misc Sepr  | 1 byte  |       | unsigned
Left Bias  | 2 bytes |       | unsigned
Right Bias | 2 bytes |       | unsigned
Air Temp   | 1 byte  |       | unsigned
Reserved   | 1 byte  |       | unsigned

#### Rate Resolution

Field      | Size    | Value | Signedness
---------- | ------- | ----- | ----------
Type       | 1 byte  | 0x10  | unsigned
Length     | 1 byte  | 0x0F  | unsigned
F Samp     | 4 bytes |       | unsigned
Log 2N FFT | 1 bytes |       | unsigned
Dist Res   | 2 bytes |       | unsigned
Coord Res  | 1 byte  |       | unsigned
Timing     | 4 bytes |       | unsigned
Reserved   | 1 byte  |       | unsigned
Reserved   | 1 byte  |       | unsigned

#### Calibration Configuration

Field   | Size    | Value | Signedness
------- | ------- | ----- | ----------
Type    | 1 byte  | 0x22  | unsigned
Length  | 1 byte  | 0x24  | unsigned
AL1     | 4 bytes |       | unsigned
AR1     | 4 bytes |       | unsigned
BL1     | 4 bytes |       | unsigned
BR1     | 4 bytes |       | unsigned
AL2     | 4 bytes |       | unsigned
AR2     | 4 bytes |       | unsigned
BL2     | 4 bytes |       | unsigned
BR2     | 4 bytes |       | unsigned
Padding | 2 bytes |       | unsigned

#### Calibration Configuration Fixed

Field    | Size    | Value | Signedness
-------- | ------- | ----- | ----------
Type     | 1 byte  | 0x26  | unsigned
Length   | 1 byte  |       | unsigned
Padding  | 9 bytes |       | unsigned
Reserved | 3 bytes |       | unsigned
AL1      | 3 bytes |       | unsigned
BL1      | 3 bytes |       | unsigned
AR1      | 3 bytes |       | unsigned
BR1      | 3 bytes |       | unsigned
AL2      | 3 bytes |       | unsigned
BL2      | 3 bytes |       | unsigned
AR2      | 3 bytes |       | unsigned
BR2      | 3 bytes |       | unsigned

#### A4 Bilinear

Field          | Size     | Value | Signedness
-------------- | -------- | ----- | ----------
Type           | 1 byte   | 0x23  | unsigned
Length         | 1 byte   | 0x24  | unsigned
Algorithm Data | 32 bytes |       | unsigned
Reserved       | 2 bytes  |       | unsigned

#### Date Time

Field    | Size   | Value | Signedness
-------- | ------ | ----- | ----------
Type     | 1 byte | 0x31  | unsigned
Length   | 1 byte | 0x09  | unsigned
Time H   | 1 byte |       | unsigned
Time M   | 1 byte |       | unsigned
Time S   | 1 byte |       | unsigned
Date D   | 1 byte |       | unsigned
Date M   | 1 byte |       | unsigned
Date Y   | 1 byte |       | unsigned
Reserved | 1 byte |       | unsigned
