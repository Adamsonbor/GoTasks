package db

type Product struct {
	Id			int		`json:"id"`
	Name		string	`json:"name"`
	Size		string	`json:"size"`
	Quantity	string	`json:"quantity"`
}

type Warehouse struct {
	Id				int		`json:"id"`
	Title			string	`json:"name"`
	Availability	bool	`json:"availability"`
}
