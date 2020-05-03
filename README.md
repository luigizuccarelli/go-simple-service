# Go simple service

Simple utility for testing web services on [Microlib](https://github.com/microlib).
Uses a generic script `script.sh` to simulate start & stop for Linux & MacOS based systems. 

## Usage 

```bash
# cd to project directory and build executable
$ go build .
$ chmod u+x script.sh


## Note
The http server by @luigizuccarelli uses signals to allow for graceful shutdown. 
Use this as a standard pattern when creating all web services. 
