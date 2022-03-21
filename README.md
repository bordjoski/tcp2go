# tcp2go
tcp2go is tcp utility tool with following capabilites:
* tcp scanner
* tcp proxy
* Replication of netcat arbitrary bash commands execution
* client (to interact with remote listener)


## Usage
### Scanning
    tcp2go scan <target>    

This command will scan all the ports and will output only the ports that are open. 

### Proxy
    tcp2go proxy -p 3001 -t example.com:3032

This command will start listener on port 3001 and will proxy the clients to given target example.com:3032

### Arbitary code execution
netcat contains feature that allows stdin and stdout of any arbitrary programs to be redirected over tcp, turning a single command execution vulnerability into operating system shell access. This feature is usually excluded from builds and tcp2go exec command is replication of netcats feature.

    tcp2go exec -p 3002
This command will start listener on port 3002 and will execute the input

To connect and interact with server, interactive client can be started by running

    tcp2go connect -t localhost:3002    

From there, please behaive.
## Installation
tcp2go is not published module. To install it, clone this repository and run

    git clone https://github.com/bordjoski/tcp2go.git
    cd tcp2go  
    go install    


## References and Credits
* Black Hat Go by Tom Steele, Chris Patten and Dan Kottmann https://www.amazon.com/Black-Hat-Go-Programming-Pentesters-ebook-dp-B073NPY29N/dp/B073NPY29N/ref=mt_other?_encoding=UTF8&me=&qid=
* RFC793 https://datatracker.ietf.org/doc/html/rfc793
* Cobra https://github.com/spf13/cobra
* Viper https://github.com/spf13/viper
* https://music.youtube.com/watch?v=5txYYSo0jWs&feature=share


