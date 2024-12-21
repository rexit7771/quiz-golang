package structs

type Buku struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Image_url    string `json:"image_url"`
	Release_year int    `json:"release_year"`
	Price        int    `json:"price"`
	Total_page   int    `json:"total_page"`
	Thickness    string `json:"thickness"`
	Category_id  int    `json:"category_id"`
	Created_at   string `json:"created_at"`
	Created_by   string `json:"created_by"`
	Modified_at  string `json:"modified_at"`
	Modified_by  string `json:"modified_by"`
}
