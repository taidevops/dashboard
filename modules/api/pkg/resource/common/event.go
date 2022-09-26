package common

type EventList struct {
	ListMeta api.ListMeta `json:"listMeta"`

	Events []Event `json:"events"`

	Errors []error `json:"errors"`
}

type Event struct {
	ObjectMeta api.ObjectMeta `json:"objectMeta"`
	TypeMeta api.TypeMeta `json:"typeMeta"`

	Message string `json:"message"`
}