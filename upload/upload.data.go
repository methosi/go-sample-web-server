package upload

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/kienit/be_brankas_test/database"
)

func insertImageData(image *Image) error {
	// Add timeout if database execution over 3s
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := fmt.Sprintf("INSERT INTO %v.%v (file_name, file_size, content_type) VALUES ($1, $2, $3)", os.Getenv("schema"), "images")
	_, err := database.DBConn.ExecContext(ctx, query,
		image.Filename,
		image.Size,
		image.ContentType)
	if err != nil {
		return err
	}

	return nil
}
