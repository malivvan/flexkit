package flex

import (
	"testing"
	"fmt"
)


func TestFlex(t *testing.T) {
	c := NewContainer().Wrap(Wrap).JustifyContent(SpaceBetween, ScreenSizeLarge).Append(NewItem(nil).Grow(1))
	fmt.Println(c.Style())
}
