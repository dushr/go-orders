// Package order provides shared routines, structs and JSON mapping for order
// processing.
package orders

const (
	// Order status
	StatusSent       = "sent"
	StatusInProgress = "progress"
	StatusSuccess    = "success"
	StatusSuccessCC  = "successcc"
	StatusRejected   = "rejected"
	StatusMissed     = "missed"
	StatusFailed     = "failed"
	StatusFailedCC   = "failedcc"
	StatusUnknown    = "unknown"
)

type Order struct {
	ID            int64        `datastore:"" json:"id"`
	CreatedTS     string       `datastore:"" json:"created_at"`
	RestaurantID  string       `datastore:"" json:"restaurant_id"`
	Status        string       `datastore:"" json:"status"`
	PaymentMethod string       `datastore:",noindex" json:"payment_method"`
	TotalPrice    int          `datastore:",noindex" json:"-"`
	MinOrderValue int          `datastore:",noindex" json:"-"`
	DeliveryFee   int          `datastore:",noindex" json:"-"`
	DeliveryInfo  DeliveryInfo `datastore:",noindex" json:"delivery_info"`
	Items         []OrderItem  `datastore:",noindex" json:"items"`
	DeliveryTime  string       `datastore:",noindex" json:"delivery_time,omitempty"`
	Restaurant    Restaurant   `datastore:",noindex" json:"restaurant"`
}

type DeliveryInfo struct {
	Firstname string `datastore:",noindex" json:"name"`
	Lastname  string `datastore:",noindex" json:"lastname"`
	Company   string `datastore:",noindex" json:"company,omitempty"`
	Street    string `datastore:",noindex" json:"street_name"`
	StreetNo  string `datastore:",noindex" json:"street_number"`
	Zipcode   string `datastore:",noindex" json:"zipcode"`
	City      string `datastore:",noindex" json:"city"`
	Door      string `datastore:",noindex" json:"door,omitempty"`
	Floor     string `datastore:",noindex" json:"etage,omitempty"`
	Phone     string `datastore:",noindex" json:"phone"`
	Email     string `datastore:",noindex" json:"email"`
	Comment   string `datastore:",noindex" json:"comments,omitempty"`
}

type OrderItem struct {
	ID          string `datastore:",noindex" json:"id"`
	Name        string `datastore:",noindex" json:"name"`
	Size        string `datastore:",noindex" json:"size"`
	Quantity    int    `datastore:",noindex" json:"quantity"`
	Price       int    `datastore:",noindex" json:"-"`
	Description string `datastore:",noindex" json:"description"`
}

type Restaurant struct {
	Name             string `datastore:",noindex" json:"name"`
	Zipcode          string `datastore:",noindex" json:"zipcode"`
	City             string `datastore:",noindex" json:"city"`
	Street           string `datastore:",noindex" json:"street"`
	PhoneAreaCode    string `datastore:",noindex" json:"phone_area_code"`
	PhoneLocalNumber string `datastore:",noindex" json:"phone_local_number"`
	StreetNo         string `datastore:",noindex" json:"street_no"`
}

type Orders []Order
