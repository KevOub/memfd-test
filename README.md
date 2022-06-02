## memfd

> memfd_create() creates an anonymous file and returns a file
       descriptor that refers to it.  The file behaves like a regular
       file, and so can be modified, truncated, memory-mapped, and so
       on.  However, unlike a regular file, it lives in RAM and has a
       volatile backing storage.  Once all references to the file are
       dropped, it is automatically released.  Anonymous memory is used
       for all backing pages of the file.  Therefore, files created by
       memfd_create() have the same semantics as other anonymous memory
       allocations such as those allocated using mmap(2) with the
       MAP_ANONYMOUS flag.

(https://man7.org/linux/man-pages/man2/memfd_create.2.html)

> fexecve - execute program specified via file descriptor

(https://man7.org/linux/man-pages/man3/fexecve.3.html) 

### How To Run

how to compile golang binary within go
```bash
 go run main.go -mode build -gofile samplebins/hello.go  
```

how to load binary into memory and execute it
```bash
go run main.go -mode loadexec -bin samplebins/lol 
```

how to load a binary without executing it 
```bash
go run main.go -mode load -bin samplebins/lol
```


### Files

###### api.go

PoC to see if you can post a go func 

###### builder.go

Takes the given golang file as base64 and compiles it. However it has to touch disk

###### dropper.go

Implements unix-specific functions pertaining to `memfd_create` and `fexecve` to run code in memomry



### Notes

fexecve uses `proc` since fexecve is not implemented natively
If you read `/proc/self/fd/%d` where `%d` is the `fd` given by `memfd`

Also if `/dev/shm` is read only no dice. TODO workaround

### How to detect

https://www.sandflysecurity.com/blog/detecting-linux-memfd-create-fileless-malware-with-command-line-forensics/

