package digine

type Label struct {
	*string
}

var NilLabel = &Label{nil}

func NewLabel(label string) *Label {
	labelPtr := (*string)(nil)
	if label != "" {
		labelPtr = &label
	}
	return &Label{labelPtr}
}

func (l *Label) Get() *string {
	return l.string
}
