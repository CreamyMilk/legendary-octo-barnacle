package engine

type Provider interface {
	GetChapter(id string) (Chapter, error)
	ListChapters(id string) ([]Chapter, error)
	Search(name string) ([]Title, error)
	ToChapterModel() Chapter
	ToTitleModel() Title
}

type Title struct {
	// path component
	ID           string `json:"id"`
	Name         string `json:"name"`
	URL          string `json:"url"`
	ChaptersUrls string `json:"chapters_urls"`
}

type Chapter struct {
	// path component
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Pages []string `json:"pages"`
}

type StatusType string

const (
	StatusOk  StatusType = "ok"
	StatusErr StatusType = "err"
)

type Status struct {
	StatusType `json:"status"`
	Error      string `json:"error"`
}

func GetRightProvider(providerName string) Provider {
	switch providerName {
	case "manganato":
		return &ManganatoProvider{}
	case "mangadex":
		return &MangadexProvider{}
	default:
		return nil
	}
}
