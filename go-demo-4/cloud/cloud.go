package cloud

type CloudDb struct {
	url string
}

func NewCloudDb(url string) *CloudDb {
	return &CloudDb{
		url: url,
	}
}

func (db *CloudDb) Read() ([]byte, error) {
	// Реализуйте чтение данных из облачного хранилища
	return nil, nil
}

func (db *CloudDb) Write(content []byte) error {
	// Реализуйте запись данных в облачное хранилище
	return nil
}
