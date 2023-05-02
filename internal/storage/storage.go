package storage

type FileTable struct {
	Files   map[string]string `json:"files"`
	IdChain map[string]string `json:"chain"`
}

func (f *FileTable) AddFile(names []string) {

}
