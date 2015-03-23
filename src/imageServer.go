package main

import (
	"net/http"
	"github.com/codegangsta/martini"
	"log"
	"github.com/disintegration/imaging"
	"strconv"
	"image"

)

func main() {
	// lets start martini and the real code
	m := martini.Classic()
	m.Get("/", func(res http.ResponseWriter, req *http.Request) {
			req.ParseForm();
			source := req.FormValue("s")
			if source != "" {
				resp, err := http.Get(source)
				defer resp.Body.Close()
				if err != nil {
					res.Write([]byte("error"))
					return
				}

				srcimage, err := imaging.Decode(resp.Body);
				if (err != nil) {
					res.Write([]byte("error"))
					return
				}

				w, _ := strconv.Atoi(req.FormValue("w"));
				h, _ := strconv.Atoi(req.FormValue("h"));
				if w < 0 {
					w = 0;
				}
				if (h < 0) {
					h = 0;
				}

				ds := &image.NRGBA{}
				if req.FormValue("c") == "" {
					ds = imaging.Resize(srcimage, w, h, imaging.Lanczos)
				}else {
					ds = imaging.Thumbnail(srcimage, w, h, imaging.Lanczos)
				}
				imaging.Encode(res, ds, imaging.JPEG);
			}

		});

	m.RunOnAddr(":8012")
	m.Run()

}
