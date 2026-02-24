# AMI-cli
an asterisk manager interface client with go

```conf 
# in /etc/asterisk/cdr_custom.conf 
[mappings]
Simple.csv => ${CSV_QUOTE(${CDR(src)})},${CSV_QUOTE(${CDR(dst)})},${CSV_QUOTE(${CDR(channel)})},${CSV_QUOTE(${CDR(duration)})},${CSV_QUOTE(${CDR(billsec)})}
```
```conf 
# in /etc/asterisk/cdr.conf 
[general]
enable=yes
```

## Note

- Each CDR includes the following times:
 * Start time - the time at which the CDR was created for Party A
 * Answer time - the time at which Party A and Party B could begin communicating
 * End time - the time at which Party A and Party B could no longer communicate

- From these times, two durations are computed:
 * Duration - the End time minus the Start time.
 * Billsec - the End time minus the Answer time. 
    (Whether or not you actually bill for this period of time is up to you)
