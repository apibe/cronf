package internal

type UD interface {
	Upload() error
	UploadRetry() error
	Download() error
	DownloadRetry() error
}

type Task interface {
	Startup()
}
