# AMI-cli
an asterisk manager interface client with go


```conf 
# in /etc/asterisk/cdr.conf 
[general]
enable=yes
```

```conf 
# in /etc/asterisk/cdr_manager.conf 
[general]
enable=yes
```

```conf 
# in /etc/asterisk/manager.d/go.conf 
[general]
secret = passpass
read = cdr
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

- to make AMI client to get CDRs :
 * edit cdr_manager.conf
 * add custom conf for the AMI inside manager.d folder
