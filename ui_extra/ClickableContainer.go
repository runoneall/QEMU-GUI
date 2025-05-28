package ui_extra

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ClickableContainer struct {
	widget.BaseWidget
	content  fyne.CanvasObject
	OnTapped func()
}

func NewClickableContainer(content fyne.CanvasObject) *ClickableContainer {
	c := &ClickableContainer{
		content: content,
	}
	c.ExtendBaseWidget(c)
	return c
}

func (c *ClickableContainer) CreateRenderer() fyne.WidgetRenderer {
	return &clickableContainerRenderer{
		container: c,
		objects:   []fyne.CanvasObject{c.content},
	}
}

func (c *ClickableContainer) Tapped(*fyne.PointEvent) {
	if c.OnTapped != nil {
		c.OnTapped()
	}
}

func (c *ClickableContainer) MinSize() fyne.Size {
	return c.content.MinSize()
}

type clickableContainerRenderer struct {
	container *ClickableContainer
	objects   []fyne.CanvasObject
}

func (r *clickableContainerRenderer) Layout(size fyne.Size) {
	r.objects[0].Resize(size)
}

func (r *clickableContainerRenderer) MinSize() fyne.Size {
	return r.container.MinSize()
}

func (r *clickableContainerRenderer) Refresh() {
	r.objects[0].Refresh()
}

func (r *clickableContainerRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *clickableContainerRenderer) Destroy() {}
