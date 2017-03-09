Calling go from C.

http://stackoverflow.com/questions/32215509/using-go-code-in-an-existing-c-project

- Opening multiple sockets and closing them is bug-prone and possibly expensive.
- Unfortunately, I'm not familiar enough with non-blocking IO in C.
  So do this on a different thread.
- Really don't want to block.

This starts out as a simple prototype.
