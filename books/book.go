package books

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	PublishedAt string `json:"published_at"`
	Language    string `json:"language"`
}
