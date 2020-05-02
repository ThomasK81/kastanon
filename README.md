# kastanon
Very Basic Console for Browsing CEX Alignment Libraries as used in Brucheion and Ducat

## Using it

- currently the file has to be `output.cex`

### Starting Up
```
YOURCOMPUTER:kastanon user$ ./kastanon
2020/05/02 16:44:34 LoadCEX: loading output.cex
2020/05/02 16:44:34 Read !#relations succesfully!
2020/05/02 16:44:34 Read !#ctsdata succesfully!
2020/05/02 16:44:34 Starting Kastanon Shell
---------------------
Kastanon Shell
---------------------

->
```

### Retrieving a Specific Passage
```
-> passage urn:cts:greekLit:tlg0016.tlg001.eng.token:1.74  
```

### Retrieving All Alignments Given an Alignment URN
```
-> alignment urn:cite2:ducat:alignments.temp:2020416_40_49_56_0
```

### Retrieving All Alignments Given a Token URN
```
-> token urn:cts:greekLit:tlg0016.tlg001.grc.token:1.74.7
```

### Exit
```
-> q
```
