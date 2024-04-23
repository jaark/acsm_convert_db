package main

import (
	"log"
	"fmt"
	"github.com/gobeam/stringy"
	"github.com/boltdb/bolt"
	"os"

)

func main() {
	db, err := bolt.Open("server_manager.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	os.Mkdir("shared_store.json", 0755)

	db.View(func(tx *bolt.Tx) error {
		tx.ForEach(func(name []byte, b *bolt.Bucket) error{
			bucket_name := string(name[:])
			dirname := stringy.New(bucket_name).SnakeCase().ToLower()
			os.Mkdir("shared_store.json/"+dirname, 0755)
			b.ForEach(func(key, val []byte) error {
				key_name := stringy.New(string(key)).SnakeCase("-", "-").ToLower()
				file_path_name := "shared_store.json/" + dirname +"/" + key_name + ".json"
				switch bucket_name {
				case "audit":
					file_path_name = "shared_store.json/audit.json"
				case "customChecksums":
					file_path_name = "shared_store.json/custom_checksums.json"
				case "customRainPresets":
					file_path_name = "shared_store.json/custom_rain_presets.json"
				case "serverOptions":
					file_path_name = "shared_store.json/" + key_name +".json"
				} 
				
				os.WriteFile(file_path_name, val, 0644)
				return nil
			})
			fmt.Println(string(dirname))
			return nil
		})
		return nil
	})

}

