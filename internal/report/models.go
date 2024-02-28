package report

type RestaurantRequest struct {
	Userid uint   `json:"userid"`
	Name   string `json:"name"`
	Format string `json:"format"`
}
