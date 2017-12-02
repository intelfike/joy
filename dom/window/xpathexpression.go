package window

import "github.com/matthewmueller/golly/js"

// XPathExpression struct
// js:"XPathExpression,omit"
type XPathExpression struct {
}

// Evaluate fn
// js:"evaluate"
func (*XPathExpression) Evaluate(contextNode Node, kind uint8, result *XPathResult) (x *XPathExpression) {
	js.Rewrite("$_.evaluate($1, $2, $3)", contextNode, kind, result)
	return x
}