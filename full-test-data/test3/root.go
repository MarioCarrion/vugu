package main

// DO NOT EDIT: This file was generated by vugu. Please regenerate instead of editing or add additional code in a separate file.

import "fmt"
import "reflect"
import "encoding/json"
import "github.com/vugu/vugu"
import js "github.com/vugu/vugu/js"

type Root struct {
	ItemCount int `vugu:"data"`
}

func (c *Root) BeforeBuild() {
	if c.ItemCount == 0 {
		c.ItemCount = 5
	}
}

func (c *Root) OnAdd() {
	c.ItemCount++
}

func (c *Root) Build(vgin *vugu.BuildIn) (vgout *vugu.BuildOut) {

	vgout = &vugu.BuildOut{}

	var vgn *vugu.VGNode
	vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "html", Attr: []vugu.VGAttribute(nil)}
	vgout.Out = append(vgout.Out, vgn)	// root for output
	{
		vgparent := vgn
		_ = vgparent
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "head", Attr: []vugu.VGAttribute(nil)}
		vgparent.AppendChild(vgn)
		{
			vgparent := vgn
			_ = vgparent
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "title", Attr: []vugu.VGAttribute(nil)}
			vgparent.AppendChild(vgn)
			{
				vgparent := vgn
				_ = vgparent
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "Test page"}
				vgparent.AppendChild(vgn)
			}
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "link", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "rel", Val: "stylesheet"}, vugu.VGAttribute{Namespace: "", Key: "href", Val: "https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"}, vugu.VGAttribute{Namespace: "", Key: "integrity", Val: "sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T"}, vugu.VGAttribute{Namespace: "", Key: "crossorigin", Val: "anonymous"}}}
			vgout.CSS = append(vgout.CSS, vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
			vgparent.AppendChild(vgn)
		}
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "body", Attr: []vugu.VGAttribute(nil)}
		vgparent.AppendChild(vgn)
		{
			vgparent := vgn
			_ = vgparent
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "test-div"}, vugu.VGAttribute{Namespace: "", Key: "id", Val: "test_div_id"}}}
			vgparent.AppendChild(vgn)
			{
				vgparent := vgn
				_ = vgparent
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\n            "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute(nil)}
				vgparent.AppendChild(vgn)
				{
					vghtml := "Let\x26#39;s see how this goes:"
					vgn.InnerHTML = &vghtml
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\n            "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "ul", Attr: []vugu.VGAttribute(nil)}
				vgparent.AppendChild(vgn)
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					for i := 0; i < c.ItemCount; i++ {
						{
							compKey := vugu.CompKey{ID: 0x5D75B8DFCC2A928B, IterKey: vgiterkey}
							// ask BuildEnv for prior instance of this specific component
							comp, _ := vgin.BuildEnv.CachedComponent(compKey).(*demoline)
							if comp == nil {
								// create new one if needed
								comp = new(demoline)
							}
							vgin.BuildEnv.UseComponent(compKey, comp)	// ensure we can use this in the cache next time around
							comp.AttrMap = make(map[string]interface{}, 8)
							comp.AttrMap["num"] = i
							vgout.Components = append(vgout.Components, comp)
							vgn = &vugu.VGNode{Component: comp}
							vgparent.AppendChild(vgn)
						}
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
					vgparent.AppendChild(vgn)
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\n            "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "button", Attr: []vugu.VGAttribute(nil)}
				vgparent.AppendChild(vgn)
				vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
					EventType:	"click",
					Func:		func(event *vugu.DOMEvent) { c.OnAdd() },
					// TODO: implement capture, etc. mostly need to decide syntax
				})
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "Add"}
					vgparent.AppendChild(vgn)
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n\n        "}
				vgparent.AppendChild(vgn)
			}
		}
	}
	return vgout
}

// 'fix' unused imports
var _ fmt.Stringer
var _ reflect.Type
var _ json.RawMessage
var _ js.Value
