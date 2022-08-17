package linkly

type FocusElement struct {
	XPath string
}

func NewFocusElement(xpath string) *FocusElement {
	return &FocusElement{
		XPath: xpath,
	}
}
