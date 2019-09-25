package chromedominate

type ChromeDOM struct {
	nodeId int64
	target *ChromeTargetDominate
}

func NewChromeDom(nodeId int64, target *ChromeTargetDominate) *ChromeDOM {
	dom := &ChromeDOM{
		nodeId: nodeId,
		target: target,
	}

	return dom
}

func (d *ChromeDOM) Click() error {
	param := GetBoxModelParam{
		NodeId: &d.nodeId,
	}
	r, err := d.target.GetBoxModel(param)

	if err != nil {
		return err
	}

	button := "left"
	buttons := int64(1)
	clickCount := int64(1)
	deltaX := r.Content[0] + 10
	deltaY := r.Content[1] + 10
	pointerType := "mouse"

	inputParam := DispatchMouseEventParam{
		Type:        "mousePressed",
		X:           deltaX,
		Y:           deltaY,
		Button:      &button,
		Buttons:     &buttons,
		ClickCount:  &clickCount,
		DeltaX:      &deltaX,
		DeltaY:      &deltaY,
		PointerType: &pointerType,
	}

	return d.target.DispatchMouseEvent(inputParam)
}

func (d *ChromeDOM) Focus() {

}

func (d *ChromeDOM) SendKeys() {

}

func (d *ChromeDOM) Find() {

}
