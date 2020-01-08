package main

import "net/http"

import "fmt"

import "sync/atomic"

const page = `<!DOCTYPE html>
<html>
    <head>
        <style>
            .title {
				text-align: center;
				display: block;
            };
		</style>
		<title>modokiアクセスカウンター</title>
    </head>
	<body>
		<div class="title">
			<h1 class="title">modokiアクセスカウンター</h1>
			<p>あなたは%d番目の来訪者です！</p>
			<a href="https://twitter.com/share" class="twitter-share-button" data-url="https://counter.apps.tsuzu.xyz/" data-text="あなたは%d番目の来訪者です！"data-size="default" data-count="none">Tweet</a> 
			<script>!function(d,s,id){var js,fjs=d.getElementsByTagName(s)[0],p=/^http:/.test(d.location)?'http':'https';if(!d.getElementById(id)){js=d.createElement(s);js.id=id;js.src=p+'://platform.twitter.com/widgets.js';fjs.parentNode.insertBefore(js,fjs);}}(document, 'script', 'twitter-wjs');</script>
		</div>
    </body>
</html>
`

func main() {
	var counter int32

	http.ListenAndServe(":80", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" || req.Method != "GET" {
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		cnt := atomic.AddInt32(&counter, 1)
		fmt.Fprintf(rw, page, cnt, cnt)
	}))
}
