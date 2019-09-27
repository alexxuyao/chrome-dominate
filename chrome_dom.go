package chromedominate

type ChromeDOM struct {
	nodeId NodeId
	target *ChromeTargetDominate
}

func NewChromeDom(nodeId NodeId, target *ChromeTargetDominate) *ChromeDOM {
	dom := &ChromeDOM{
		nodeId: nodeId,
		target: target,
	}

	return dom
}

func (d *ChromeDOM) Click() error {
	return d.ClickTimes(1)
}

func (d *ChromeDOM) ClickTimes(times int) error {
	param := GetBoxModelParam{
		NodeId: &d.nodeId,
	}

	r, err := d.target.GetBoxModel(param)

	if err != nil {
		return err
	}

	button := "left"
	buttons := int64(1)
	clickCount := int64(times)
	deltaX := r.Content[0] + 10
	deltaY := r.Content[1] + 10
	pointerType := "mouse"

	inputParam := DispatchMouseEventParam{
		X:           deltaX,
		Y:           deltaY,
		Button:      &button,
		Buttons:     &buttons,
		ClickCount:  &clickCount,
		DeltaX:      &deltaX,
		DeltaY:      &deltaY,
		PointerType: &pointerType,
	}

	for i := 0; i < int(clickCount); i++ {

		inputParam.Type = "mousePressed"
		err = d.target.DispatchMouseEvent(inputParam)
		if err != nil {
			return err
		}

		inputParam.Type = "mouseReleased"
		err = d.target.DispatchMouseEvent(inputParam)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *ChromeDOM) SetAttributeValue(name, value string) error {
	return d.target.SetAttributeValue(d.nodeId, name, value)
}

func (d *ChromeDOM) GetOuterHTML() (string, error) {
	param := GetOuterHTMLParam{
		NodeId: &d.nodeId,
	}
	return d.target.GetOuterHTML(param)
}

func (d *ChromeDOM) Focus() {

}

func (d *ChromeDOM) SendKeys() {

}

func (d *ChromeDOM) Find() {

}
