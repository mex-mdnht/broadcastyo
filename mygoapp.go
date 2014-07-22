package hello

import (
	"net/http"
	"net/url"
	"appengine"
	"appengine/urlfetch"
)

const API_TOKEN string = "b0514855-2c5e-f5c9-ec87-54ae99e1ccdd"
const YO_ALL string = "http://api.justyo.co/yoall/"

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query()["username"]
	/*path := r.URL.Path
	apikey := path[1:len(path)-1]*/
	
	c := appengine.NewContext(r)
	client := http.Client{Transport: &urlfetch.Transport{Context:c}}
	//&& apikey != ""
	if username != nil {
		res, err := client.PostForm(YO_ALL, url.Values{"api_token":{API_TOKEN}})
		if err != nil {
			c.Errorf("Error %s",err)
		}
		if res.StatusCode != 200{
			c.Warningf("yo failed with: %s", res)
		}
	}else
	{
		c.Errorf("username and key required")
	}
}


