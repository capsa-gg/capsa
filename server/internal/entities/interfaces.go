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
	// UploadFile stores the contents on a given path.
	UploadFile(path string, contents []byte) error
	// DeleteFile deletes a file from blob storage.
	DeleteFile(path string) error
}

// FilteredLineLoader is a way to abstract getting log lines, used for merging logs.
type FilteredLineLoader interface {
	// HasNextLine returns whether the log has a next line available.
	HasNextLine() (bool, error)

	// ReadNextLineMetadata returns the metadata for the next log.
	// In case error is nil and lineMetadata is also nil, there are no more log lines.
	// NOTE: there must be a check for `lineMetadata != nil`, as nil is a valid return value, even without an error.
	ReadNextLineMetadata() (lineMetadata *LogChunkLineMetadata, err error)

	// GetNextLine returns the next line of a log.
	// In case error is nil and logLine is also nil, there are no more log lines.
	// By calling GetNextLine, the line gets consumed and the return value for ReadNextLineMetadata gets populated.
	// The line content should not include the final line break.
	// NOTE: there must be a check for `logLine != nil`, as nil is a valid return value, even without an error.
	GetNextLine() (logLine *string, err error)
}
