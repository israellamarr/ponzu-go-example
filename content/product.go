package content

import (
	"fmt"

	"github.com/bosssauce/reference"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Product struct {
	item.Item

	ProductId   int    `json:"product_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Url         string `json:"url"`
}

// MarshalEditor writes a buffer of html to edit a Product within the CMS
// and implements editor.Editable
func (p *Product) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(p,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Product field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("ProductId", p, map[string]string{
				"label":       "ProductId",
				"type":        "text",
				"placeholder": "Enter the ProductId here",
			}),
		},
		editor.Field{
			View: editor.Input("Name", p, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Description", p, map[string]string{
				"label":       "Description",
				"placeholder": "Enter the Description here",
			}),
		},
		editor.Field{
			View: reference.Select("Category", p, map[string]string{
				"label": "Category",
			},
				"Category",
				`{{ .name }} `,
			),
		},
		editor.Field{
			View: editor.Input("Url", p, map[string]string{
				"label":       "Url",
				"type":        "text",
				"placeholder": "Enter the Url here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Product editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Product"] = func() interface{} { return new(Product) }
}

// String defines how a Product is printed. Update it using more descriptive
// fields from the Product struct type
func (p *Product) String() string {
	return fmt.Sprintf("Product: %s", p.UUID)
}
