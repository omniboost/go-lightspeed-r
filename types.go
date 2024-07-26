package lightspeed_r

// const (
// 	WebhookEventCreated         WebhookEventType = "event.created"
// 	WebhookEventUpdated         WebhookEventType = "event.updated"
// 	WebhookEventDeleted         WebhookEventType = "event.deleted"
// 	WebhookEventActivityCreated WebhookEventType = "event_activity.created"
// 	WebhookEventActivityUpdated WebhookEventType = "event_activity.updated"
// 	WebhookEventActivityDeleted WebhookEventType = "event_activity.deleted"
// )

// type WebhookEventType string

// type Page struct {
// 	Status       string    `json:"status"`
// 	StatusCode   StringInt `json:"status_code"`
// 	Locale       string    `json:"locale,omitempty"`
// 	NextUrl      string    `json:"next_url,omitempty"`
// 	PreviousUrl  string    `json:"previous_url,omitempty"`
// 	ItemsPerPage int       `json:"items_per_page,omitempty"`
// 	TotalItems   int       `json:"total_items,omitempty"`
// 	TotalPages   int       `json:"total_pages,omitempty"`
// 	Message      string    `json:"message,omitempty"`
// }

// type PaginatedResponse[T any] struct {
// 	Data []T  `json:"data"`
// 	Page Page `json:"page"`
// }

// type PageStatusResponse[T any] struct {
// 	Data T    `json:"data"`
// 	Page Page `json:"page"`
// }

type Account struct {
	Attributes struct {
		Next     string `json:"next"`
		Previous string `json:"previous"`
	} `json:"@attributes"`
	Account struct {
		AccountID string `json:"accountID"`
		Name      string `json:"name"`
		Link      struct {
			Attributes struct {
				Href string `json:"href"`
			} `json:"@attributes"`
		} `json:"link"`
	} `json:"Account"`
}
