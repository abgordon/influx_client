// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	// "net/url"
	"flag"
	"html/template"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	// "fmt"
    "github.com/influxdb/influxdb/client/v2"
	"time"
)

const (
    username = "root"
    password = "root"
    INFLUX_URL_LOCAL = "localhost:8086"
    MyDB = "temp_db"
    db = "temp_db"
)



var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)

        /*
            write to influx here
        */ 



		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}

func main() {
	/*
		influx init
	*/
	// u, err := url.Parse(INFLUX_URL_LOCAL)
	// if err != nil {
	// 	fmt.Println("err on URL parse:", err)
	// }
	// // make client
	// c := client.NewClient(client.Config{
	// 	URL: u,
	// 	Username: username,
	// 	Password: password,
	// })


    
    

    // Create a new point batch
 //    bp,_ := client.NewBatchPoints(client.BatchPointsConfig{
 //        Database:  MyDB,
 //        Precision: "s",
 //    })

	// tags := map[string]string{"cpu": "cpu-total"}
 //    fields := map[string]interface{}{
 //        "idle":   10.1,
 //        "system": 53.3,
 //        "user":   46.6,
 //    }
 //    pt, _ := client.NewPoint("cpu_usage", tags, fields, time.Now())
 //    bp.AddPoint(pt)


	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a new password to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))