package app

import (
	"context"
	"io"
	"os"

	"github.com/jinzhu/copier"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// FS returns the instance of fsUtil (filesystem utility).
// If the fsClient is nil, it initializes it by calling the configure method of fsUtil.
// It then returns the fsClient.
func FS() *fsUtil {
	if fsClient == nil {
		fsClient = fsUtil{}.configure()
	}
	return fsClient
}

// fsClient holds the instance of fsUtil (filesystem utility).
// It is used to access the filesystem utility throughout the package.
var fsClient *fsUtil

// fsUtil represents a filesystem utility.
// It contains various fields to store configuration related to the filesystem.
// It also has a field mClient of type minio.Client to interact with the any S3 compatible object storage server.
// The type implements several methods to configure the filesystem and perform operations like uploading and deleting files.
type fsUtil struct {
	Driver        string
	LocalDirPath  string
	PublicDirPath string
	EndPoint      string
	Port          int
	Region        string
	BucketName    string
	AccessKey     string
	SecretKey     string
	mClient       *minio.Client
	ctx           context.Context
	err           error
}

// configure configures the filesystem utility based on the provided environment variables.
// It checks the FS_DRIVER environment variable to determine whether to use a local filesystem or a cloud storage like AWS S3.
// If it's a local filesystem, it sets the necessary fields and returns the updated fsUtil.
// If it's a cloud storage, it sets the fields required for the MinIO client, creates a bucket if it doesn't exist, and returns the updated fsUtil.
// If any error occurs during configuration, it falls back to using the local filesystem.
func (f fsUtil) configure() *fsUtil {
	if FS_DRIVER == "local" {
		return f.setLocalFS()
	}

	// Cloud Storage Like AWS S3, Google Cloud Storage, etc
	f.Driver = FS_DRIVER
	f.EndPoint = FS_END_POINT
	f.Port = FS_PORT
	f.Region = FS_REGION
	f.BucketName = FS_BUCKET_NAME
	f.AccessKey = FS_ACCESS_KEY
	f.SecretKey = FS_SECRET_KEY
	f.mClient, f.err = minio.New(f.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(f.AccessKey, f.SecretKey, ""),
		Secure: true,
	})
	if f.err != nil {
		return f.setLocalFS()
	}
	f.ctx = context.Background()
	isBucketExists := false
	isBucketExists, f.err = f.mClient.BucketExists(f.ctx, f.BucketName)
	if !isBucketExists {
		f.err = f.mClient.MakeBucket(f.ctx, f.BucketName, minio.MakeBucketOptions{Region: f.Region})
	}
	if f.err != nil {
		return f.setLocalFS()
	}
	return &f
}

// setLocalFS sets the fields of fsUtil to use a local filesystem.
// It handles any error that occurred during configuration and logs an error message.
// It sets the Driver to "local" and sets the local directory paths for storing files.
// It also creates the local directory path if it doesn't exist.
func (f fsUtil) setLocalFS() *fsUtil {
	if f.err != nil {
		Logger().Error().Msg(f.err.Error() + ", local filesystem will be used.")
	}
	f.Driver = "local"
	f.LocalDirPath = FS_LOCAL_DIR_PATH
	if f.LocalDirPath == "" {
		f.LocalDirPath = "storages"
	}
	f.PublicDirPath = FS_PUBLIC_DIR_PATH
	if f.PublicDirPath == "" {
		f.PublicDirPath = "storages"
	}
	f.createLocalDirPath()
	return &f
}

// createLocalDirPath creates the local directory path if it doesn't exist.
// It checks if the directory path exists and creates it with the appropriate permissions if it doesn't.
func (f fsUtil) createLocalDirPath() {
	_, err := os.Stat(f.LocalDirPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(f.LocalDirPath, 0755)
		if err != nil {
			Logger().Error().
				Err(err).
				Str("FS_DRIVER", FS_DRIVER).
				Str("LocalDirPath", f.LocalDirPath).
				Msg("Failed to create local dir path.")
		}
	}
}

// GetFileUrl constructs and returns the URL for accessing a file.
// It considers the configuration of the filesystem utility and the provided filename and path.
// The resulting URL depends on the storage driver and the endpoint being used.
func (f *fsUtil) GetFileUrl(fileName string, path ...string) string {
	res := APP_URL + "/" + f.PublicDirPath
	if f.EndPoint == "s3.amazonaws.com" {
		res = "https://" + f.BucketName + ".s3." + f.Region + ".amazonaws.com/"
	} else if f.Driver != "local" {
		res = "https://" + f.BucketName + "." + f.EndPoint + "/"
	}

	for _, p := range path {
		res += p + "/"
	}
	res += fileName

	return res
}

// Upload uploads a file to the configured storage.
// It takes the filename, source reader, file size, and optional upload options.
// If the storage driver is local, it copies the file from the source reader to the local directory.
// If it's a cloud storage, it uses the MinIO client to upload the file to the configured bucket.
func (f *fsUtil) Upload(fileName string, src io.Reader, fileSize int64, opts ...FileUploadOption) (FileUploadInfo, error) {
	if f.Driver == "local" {
		dst, err := os.Create(f.LocalDirPath + "/" + fileName)
		if err != nil {
			return FileUploadInfo{}, nil
		}
		defer dst.Close()
		_, err = io.Copy(dst, src)
		return FileUploadInfo{}, err
	}

	opt := minio.PutObjectOptions{}
	opt.UserMetadata = map[string]string{"x-amz-acl": "public-read"}
	if len(opts) > 0 {
		copier.Copy(opt, opts[0])
	}
	info, err := f.mClient.PutObject(f.ctx, f.BucketName, fileName, src, fileSize, opt)
	return FileUploadInfo{info}, err
}

// Delete deletes a file from the configured storage.
// It takes the filename and optional delete options.
// If the storage driver is local, it deletes the file from the local directory.
// If it's a cloud storage, it uses the MinIO client to remove the file from the configured bucket.
func (f *fsUtil) Delete(fileName string, opts ...FileDeleteOption) error {
	if f.Driver == "local" {
		return os.Remove(f.LocalDirPath + "/" + fileName)
	}

	opt := minio.RemoveObjectOptions{}
	opt.GovernanceBypass = true
	if len(opts) > 0 {
		copier.Copy(opt, opts[0])
	}
	return f.mClient.RemoveObject(f.ctx, f.BucketName, fileName, opt)
}

// This type represents the options for file upload.
// It embeds minio.PutObjectOptions to store upload-specific details.
type FileUploadOption struct {
	minio.PutObjectOptions
}

// FileUploadOption represents the information related to a file upload.
// It embeds minio.UploadInfo to store upload-specific details.
type FileUploadInfo struct {
	minio.UploadInfo
}

// FileDeleteOption represents the options for file deletion.
// It embeds minio.RemoveObjectOptions to provide additional delete options.
type FileDeleteOption struct {
	minio.RemoveObjectOptions
}
