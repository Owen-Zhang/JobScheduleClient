package main


import "os"
import "io"
import "fmt"
import (
	"net/http"
)

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.Handle("/", http.StripPrefix("/",http.FileServer(http.Dir("./staticfile"))))
	fmt.Println("文件下载")
	
	http.ListenAndServe(":8988", nil)
}

func uploadHandler (w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "POST":
			err := r.ParseMultipartForm(100000)
	        if err != nil {
	            http.Error(w, err.Error(), http.StatusInternalServerError)
	            return
	        }
			m := r.MultipartForm
			files := m.File["uploadfile"]

	        for i, _ := range files {
	            //for each fileheader, get a handle to the actual file
	            file, err := files[i].Open()
	            defer file.Close()
	            if err != nil {
	                http.Error(w, err.Error(), http.StatusInternalServerError)
	                return
	            }
				dst, err := os.Create("./staticfile/" + files[i].Filename)
	            defer dst.Close()
	            if err != nil {
	                http.Error(w, err.Error(), http.StatusInternalServerError)
	                return
	            }
	            //copy the uploaded file to the destination file
	            if _, err := io.Copy(dst, file); err != nil {
	                http.Error(w, err.Error(), http.StatusInternalServerError)
	                return
	            }
			}
		default:
        	w.WriteHeader(http.StatusMethodNotAllowed)
	}
}