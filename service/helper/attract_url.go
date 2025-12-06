package helper

func AttachBaseURL(url, BucketName string, path *string) {
	*path = url + "/" + BucketName + "/" + *path
}
