package config

type AirpaxCloud struct {
	Upload        string `json:"upload"`
	Download      string `json:"download"`
	Provider      string `json:"provider"`
	BucketName    string `json:"bucket_name"`
	Folder        string `json:"folder"`
	Authorization string `json:"authorization"`
}
