package config

type GCS struct {
	GCSKey      string `json:"gcs_key"`
	GCSSecret   string `json:"gcs_secret"`
	GCSRegion   string `json:"gcs_region"`
	GCSBucket   string `json:"gcs_bucket"`
	GCSUrl      string `json:"gcs_url"`
	GCSEndpoint string `json:"gcs_endpoint"`
	GCSFolder   string `json:"gcs_folder"`
}
