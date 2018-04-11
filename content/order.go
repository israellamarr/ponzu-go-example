package content

import (
	"fmt"

	"github.com/bosssauce/reference"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Order struct {
	item.Item

	OrderId  int    `json:"order_id"`
	Products string `json:"products"`
	Customer string `json:"customer"`
}

// MarshalEditor writes a buffer of html to edit a Order within the CMS
// and implements editor.Editable
func (o *Order) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(o,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Order field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("OrderId", o, map[string]string{
				"label":       "OrderId",
				"type":        "text",
				"placeholder": "Enter the OrderId here",
			}),
		},
		editor.Field{
			View: reference.Select("Products", o, map[string]string{
				"label": "Products",
			},
				"Product",
				`{{ .name }} {{ .price }} `,
			),
		},
		editor.Field{
			View: editor.Input("Customer", o, map[string]string{
				"label":       "Customer",
				"type":        "text",
				"placeholder": "Enter the Customer here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Order editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Order"] = func() interface{} { return new(Order) }
}

// String defines how a Order is printed. Update it using more descriptive
// fields from the Order struct type
func (o *Order) String() string {
	return fmt.Sprintf("Order: %s", o.UUID)
}
