package domain_linkly

type FocusElement struct {
	XPath string
}

func NewFocusElement(xpath string) *FocusElement {
	return &FocusElement{
		XPath: xpath,
	}
}
