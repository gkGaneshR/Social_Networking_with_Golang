package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	//"image"
	"github.com/boltdb/bolt"
	"github.com/eahydra/gouuid"
	"github.com/nfnt/resize"
	//"html/template"
	"bytes"
	"html/template"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

const templa = `
<html>
<head>

<style>
body {
	margin:0px 200px;
    background-color: #cccccc;
}
#views
{
  position:relative;
  height:932px;
  min-heigth:932px;
  margin:16px;
 // overflow:auto;
}

#leftview
{
  position:absolute;
  top:0px;
}

#Layer1
{
  position:absolute;
  top:0px;
  left: 300px;	
  background-color:while;
  width:400px;
  height:400px;
  min-height:400px;	
}
</style>
<script language="javascript">
function showPic(sUrl)
{
 var x,y;
 x = event.clientX;
 y = event.clientY;
 document.getElementById("Layer1").style.left = x;
 document.getElementById("Layer1").style.top = y;
 document.getElementById("Layer1").innerHTML = "<img height=500 width=500 src=\"" + sUrl + "\">";
 document.getElementById("Layer1").style.display = "block";
}
function hiddenPic(){
 document.getElementById("Layer1").innerHTML = "";
 document.getElementById("Layer1").style.display = "none";
}
</script></head>
<body align="middle">
The images are : <br>

<div id="views">
<div id="leftview">
{{range .Imm}}
<img src="/static/server/scaledloc/{{.}}" border="3" width="200" height="200" onmouseout="hiddenPic();" onmousemove="showPic(this.src);"/> <br>
{{end}}
 </div>
<div id="Layer1" style="display:none;position:absolute;z-index:1;"></div>

</div>
</div>

</body>
</html>
`
const templ = `
<html>
    <head>
        <script src="http://ajax.googleapis.com/ajax/libs/jquery/1/jquery.js"></script>
        <script src="/static/galleria/galleria-1.4.2.min.js"></script>
		<style>
    .galleria{ width: 800px; height: 500px; background: #000 ; margin-left : 100px ;}
</style>
    </head>
    <body >
		<br><br>
        <div class="galleria">
		{{range .Imm}}
		<img src="/static/server/scaledloc/{{.}}" />
         {{end}}   
            
        </div>
        <script>
            Galleria.loadTheme('/static/galleria/themes/classic/galleria.classic.min.js');
            Galleria.run('.galleria');
        </script>
    </body>
</html>
`

var world = []byte("world")
var key [10][]byte
var value [10][]byte
var val []byte
var i int

type Imagee struct {
	Imagename [10][]byte
	Imm       []string
}

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("v r at helloserver")
	file, handler, err := req.FormFile("id-file-d")

	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	currenttime := time.Now().Nanosecond()
	//fmt.Println("Current time : ", currenttime.Format("2006-01-02 15:04:05 +0800"))
	// change both atime and mtime to currenttime

	tempstr := handler.Filename
	str := string(tempstr)
	fname := strconv.Itoa(int(currenttime)) + "_" + str
	fmt.Println("current time is ", currenttime)
	fmt.Println("fname is ", fname)
	//fmt.Println(string(currenttime.Format("2006-01-02 15:04:05 ")) + str)

	err = ioutil.WriteFile("C:/mygo/gowiki/client/bin/static/server/pictures/"+fname, data, 0777)
	if err != nil {
		fmt.Println(err)
	}

	file1, err := os.Open("C:/mygo/gowiki/client/bin/static/server/pictures/" + fname)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file1)
	if err != nil {
		log.Fatal(err)
	}
	file1.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(600, 400, img, resize.Lanczos3)

	out, err := os.Create("C:/mygo/gowiki/client/bin/static/server/scaledLoc/" + fname)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)

	//writing to database bolt.db
	db, err := bolt.Open("bolt.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	u := gouuid.NewUUID()
	uuidStr := fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	fmt.Println(uuidStr)

	key[i] = []byte(uuidStr)
	value[i] = []byte(fname)

	// store some data
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(world)
		if err != nil {
			return err
		}

		err = bucket.Put(key[i], value[i])
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	i = i + 1

}

func accessdb(w http.ResponseWriter, req *http.Request) {
	fmt.Println("after update")
	v := Imagee{}
	v.Imm = make([]string, 10, 10)

	t := template.New("Imagee template")
	t, err := t.Parse(templ)
	checkError(err)

	db, err := bolt.Open("bolt.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	j := 0
	defer db.Close()
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(world)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", world)
		}
		c := bucket.Cursor()

		for keyy, valuu := c.First(); keyy != nil; keyy, valuu = c.Next() {
			// retrieve the data

			//v.Imagename[j] = bucket.Get(keyy)

			v.Imagename[j] = valuu
			fmt.Println(string(valuu))
			//fmt.Fprint(w, string(key[j]))

			//v.Imagename[j] = val

			//v.Imm[j] = string("C:/mygo/gowiki/server/pictures/" + v.Imm[j])

			//fmt.Fprintln(w, "---", string(val))
			//fmt.Fprintln(w, "---", v.Imm[j])
			j = j + 1

		}
		fmt.Println("no of files is", j)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	for k := 0; k <= j; k++ {
		v.Imagename[k] = bytes.Trim(v.Imagename[k], "") // leading and trailing

		v.Imm[k] = string(v.Imagename[k])
		//fmt.Println(v.Imagename[k])

	}

	v.Imm = v.Imm[:j]

	err = t.Execute(w, v)
	checkError(err)

	//err = t.Execute(os.Stdout, person)
	//checkError(err)

}

func HandleImages(w http.ResponseWriter, req *http.Request) {
	//fmt.Fprint(w, "v r at disp")

	files, _ := filepath.Glob("C:/mygo/gowiki/client/bin/static/server/pictures/*")

	l := len(files)
	for i := 0; i < l; i++ {
		files[i] = filepath.Base(files[i])
	}
	for i := 0; i < l; i++ {
		fmt.Fprintln(w, files[i])
	}
	data, _ := json.Marshal(files)
	fmt.Fprint(w, string(data))

	return
	//l = len(files)
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < l; i++ {
		file, err := os.Open(files[i])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		buff := make([]byte, 512000) // why 512 bytes ? see http://golang.org/pkg/net/http/#DetectContentType
		_, err = file.Read(buff)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		filetype := http.DetectContentType(buff)

		switch filetype {
		case "image/jpeg", "image/jpg", "image/gif", "image/png":
			files[i] = filepath.Base(files[i])
			fmt.Fprintln(w, files[i])

		}
	}
	return
}

func HandleImage(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	var imageName string = vars["ImageName"]
	imageLocation := "C:/mygo/gowiki/client/bin/static/server/pictures/" + imageName
	reader, err := os.Open(imageLocation)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer reader.Close()

	imageData, err := jpeg.Decode(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	jpeg.Encode(w, imageData, &jpeg.Options{100})
	fmt.Fprintln(w, "image displayed")
}

func main() {
	i = 0
	r := mux.NewRouter()
	r.HandleFunc("/upload/", HelloServer)

	r.HandleFunc("/app/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, appHtml)
	})

	r.HandleFunc("/images/", HandleImages)
	//r.HandleFunc("/images/{ImageName}/", HandleImage)

	r.HandleFunc("/accessdb/", accessdb)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", r)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
"# Social_Networking_with_Golang" 
