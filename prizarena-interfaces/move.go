package prizarena_interfaces

type MoveDto struct {
	Bid int `json:",omitempty"`
	Target string `json:",omitempty"`
}
