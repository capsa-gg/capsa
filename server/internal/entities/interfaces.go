package entities

// Note: these interfaces must be placed in entities, not in interactors, as that would cause a circular import.

// EmailSender is used to make email implementations hidden in calling code.
type EmailSender interface {
	// SendEmail sends an email with the provided arguments.
	SendEmail(toEmail, toName, subject, htmlContents string) error
}

// BlobStorage is a generic way to download and upload files to blob storage.
type BlobStorage interface {
	// DownloadFile retrieves the file contents for a given path.
	DownloadFile(path string) ([]byte, error)
	// UploadFile stores the contents on a given path
	UploadFile(path string, contents []byte) error
}
