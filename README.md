Installing MonkieToes for Google App Engine
--------------------------------------------
 
1) Install Go or Google App Engine Go Runtime
	Heroku Instructions will come later

	...

2) Install Gorilla web toolkit 
	- http://www.gorillatoolkit.org/pkg/mux

	$ go get github.com/gorilla/mux

3) Install MonkieToes

	Rename app.settings.yaml --> app.yaml

	Replace "<the application id>" with the appid you're using on App Engine
	Replace "<the version id>" with a version id of your choosing.


4)  ...

5) Profit!



Go App Command
Run on the current directory (if you're in the directory your app lives in) 
on port 8090 and admin port 8009
$ goapp serve -port=8090 -admin_port=8009 ./

Regular Command
$ go run main.go
