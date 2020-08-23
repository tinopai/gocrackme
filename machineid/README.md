# Machine ID CrackMe
Simple Messagebox protected by a password (SHA256 hashed: [machineid[:6]](https://github.com/tinopai/gocrackme/blob/master/machineid/crackme.go#L31))

### Rules
1. You must not bruteforce it.
2. It should run on any machine, not just yours.
3. Maximum runtime is 30 seconds.
4. You must not hex edit the "Wrong Password" string to a "Congratulations! You solved it" string.

### Building
Building on Windows is easy!
1. `git clone https://github.com/tinopai/gocrackme.git`
2. `cd gocrackme/machineid`
3. `go get github.com/denisbrodbeck/machineid`
4. `go build -o MID_CrackMe.exe crackme.go`
5. `MID_CrackMe.exe` **(OPTIONAL)**
