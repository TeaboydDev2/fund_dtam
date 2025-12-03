package helper

func AttachBaseURL(url, BucketName, path string) string {

	baseUrl := url + "/" + BucketName + "/" + path

	return baseUrl
}
