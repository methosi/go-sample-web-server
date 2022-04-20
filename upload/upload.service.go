package upload

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/kienit/be_brankas_test/auth"
	"github.com/kienit/be_brankas_test/cors"
)

var uploadDirectory string = filepath.Join("uploaded-image")

const uploadBasePath = "upload"

// SetupRoutes routes for upload
func SetupRoutes() {
	uploadHandler := http.HandlerFunc(uploadHandler)
	http.Handle(fmt.Sprintf("/%s", uploadBasePath), auth.AuthMiddleware(cors.CORSMiddleware(uploadHandler)))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		r.ParseMultipartForm(9 << 20) // Maximum file size is 9Mb

		// Reading file from field "data"
		file, handler, err := r.FormFile("data")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Check if file size exceed limit (8Mb)
		if handler.Size > (8 << 20) {
			fmt.Println("Exceed file size limit (max 8Mb)")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Exceed file size limit (max 8Mb)")
			return
		}

		if !checkIfValidImg(file) {
			fmt.Println("Invalid file type")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Invalid file type")
			return
		}

		// Since we read the first 512 bytes from the file in order to determine the content type
		// We need to reset pointer to start of the file before we copy
		file.Seek(0, io.SeekStart)

		err = saveImgFile(file)
		if err != nil {
			fmt.Println("Failed to save image " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		imgMeta := Image{handler.Filename, handler.Size, handler.Header.Get("Content-Type")}
		err = insertImageData(&imgMeta)
		if err != nil {
			fmt.Println("Failed to insert database " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return

	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func checkIfValidImg(file multipart.File) bool {
	buff := make([]byte, 512)
	_, err := file.Read(buff)
	if err != nil {
		fmt.Println(err)
		return false
	}

	fileType := http.DetectContentType(buff)

	switch fileType {
	case "image/jpeg", "image/jpg", "image/gif", "image/png":
		fmt.Println(fileType)
		return true
	case "application/pdf":
		fmt.Println(fileType)
		return false
	default:
		fmt.Println("unknown file type uploaded")
		return false
	}
}

func saveImgFile(file multipart.File) error {
	// Create a temporary file within our uploadDirectory directory that follows pattern
	tempFile, err := ioutil.TempFile(uploadDirectory, "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer tempFile.Close()

	// Read all of the contents of our uploaded file into a byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Write this byte array to our temporary file
	_, err = tempFile.Write(fileBytes)
	if err != nil {
		return err
	}

	return nil
}
