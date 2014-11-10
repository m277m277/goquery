package goquery

import (
	"testing"
)

const (
	wrapHtml = "<div id=\"ins\">test string<div><p><em><b></b></em></p></div></div>"
)

func TestAfter(t *testing.T) {
	doc := Doc2Clone()
	doc.Find("#main").After("#nf6")

	assertLength(t, doc.Find("#main #nf6").Nodes, 0)
	assertLength(t, doc.Find("#foot #nf6").Nodes, 0)
	assertLength(t, doc.Find("#main + #nf6").Nodes, 1)
	printSel(t, doc.Selection)
}

func TestAfterMany(t *testing.T) {
	doc := Doc2Clone()
	doc.Find(".one").After("#nf6")

	assertLength(t, doc.Find("#foot #nf6").Nodes, 1)
	assertLength(t, doc.Find("#main #nf6").Nodes, 1)
	assertLength(t, doc.Find(".one + #nf6").Nodes, 2)
	printSel(t, doc.Selection)
}

func TestAfterWithRemoved(t *testing.T) {
	doc := Doc2Clone()
	s := doc.Find("#main").Remove()
	s.After("#nf6")

	assertLength(t, s.Find("#nf6").Nodes, 0)
	assertLength(t, doc.Find("#nf6").Nodes, 0)
	printSel(t, doc.Selection)
}

func TestAfterSelection(t *testing.T) {
	doc := Doc2Clone()
	doc.Find("#main").AfterSelection(doc.Find("#nf1, #nf2"))

	assertLength(t, doc.Find("#main #nf1, #main #nf2").Nodes, 0)
	assertLength(t, doc.Find("#foot #nf1, #foot #nf2").Nodes, 0)
	assertLength(t, doc.Find("#main + #nf1, #nf1 + #nf2").Nodes, 2)
	printSel(t, doc.Selection)
}

func TestAfterHtml(t *testing.T) {
	doc := Doc2Clone()
	doc.Find("#main").AfterHtml("<strong>new node</strong>")

	assertLength(t, doc.Find("#main + strong").Nodes, 1)
	printSel(t, doc.Selection)
}

func TestAppend(t *testing.T) {
	doc := Doc2Clone()
	doc.Find("#main").Append("#nf6")

	assertLength(t, doc.Find("#foot #nf6").Nodes, 0)
	assertLength(t, doc.Find("#main #nf6").Nodes, 1)
	printSel(t, doc.Selection)
}

func TestAppendBody(t *testing.T) {
	doc := Doc2Clone()
	doc.Find("body").Append("#nf6")

	assertLength(t, doc.Find("#foot #nf6").Nodes, 0)
	assertLength(t, doc.Find("#main #nf6").Nodes, 0)
	assertLength(t, doc.Find("body > #nf6").Nodes, 1)
	printSel(t, doc.Selection)
}

func TestAppendSelection(t *testing.T) {
	doc := Doc2Clone()
	doc.Find("#main").AppendSelection(doc.Find("#nf1, #nf2"))

	assertLength(t, doc.Find("#foot #nf1").Nodes, 0)
	assertLength(t, doc.Find("#foot #nf2").Nodes, 0)
	assertLength(t, doc.Find("#main #nf1").Nodes, 1)
	assertLength(t, doc.Find("#main #nf2").Nodes, 1)
	printSel(t, doc.Selection)
}

func TestAppendSelectionExisting(t *testing.T) {
	doc := Doc2Clone()
	doc.Find("#main").AppendSelection(doc.Find("#n1, #n2"))

	assertClass(t, doc.Find("#main :nth-child(1)"), "three")
	assertClass(t, doc.Find("#main :nth-child(5)"), "one")
	assertClass(t, doc.Find("#main :nth-child(6)"), "two")
	printSel(t, doc.Selection)
}

func TestAppendClone(t *testing.T) {
	doc := Doc2Clone()
	doc.Find("#n1").AppendSelection(doc.Find("#nf1").Clone())

	assertLength(t, doc.Find("#foot #nf1").Nodes, 1)
	assertLength(t, doc.Find("#main #nf1").Nodes, 1)
	printSel(t, doc.Selection)
}

func TestAppendHtml(t *testing.T) {
	doc := Doc2Clone()
	doc.Find("div").AppendHtml("<strong>new node</strong>")

	assertLength(t, doc.Find("strong").Nodes, 14)
	printSel(t, doc.Selection)
}

func TestBefore(t *testing.T) {
	doc := Doc2Clone()
	doc.Find("#main").Before("#nf6")

	assertLength(t, doc.Find("#main #nf6").Nodes, 0)
	assertLength(t, doc.Find("#foot #nf6").Nodes, 0)
	assertLength(t, doc.Find("body > #nf6:first-child").Nodes, 1)
	printSel(t, doc.Selection)
}

func TestBeforeWithRemoved(t *testing.T) {
	doc := Doc2Clone()
	s := doc.Find("#main").Remove()
	s.Before("#nf6")

	assertLength(t, s.Find("#nf6").Nodes, 0)
	assertLength(t, doc.Find("#nf6").Nodes, 0)
	printSel(t, doc.Selection)
}

func TestBeforeSelection(t *testing.T) {
	doc := Doc2Clone()
	doc.Find("#main").BeforeSelection(doc.Find("#nf1, #nf2"))

	assertLength(t, doc.Find("#main #nf1, #main #nf2").Nodes, 0)
	assertLength(t, doc.Find("#foot #nf1, #foot #nf2").Nodes, 0)
	assertLength(t, doc.Find("body > #nf1:first-child, #nf1 + #nf2").Nodes, 2)
	printSel(t, doc.Selection)
}

func TestBeforeHtml(t *testing.T) {
	doc := Doc2Clone()
	doc.Find("#main").BeforeHtml("<strong>new node</strong>")

	assertLength(t, doc.Find("body > strong:first-child").Nodes, 1)
	printSel(t, doc.Selection)
}

func TestEmpty(t *testing.T) {
	doc := Doc2Clone()
	s := doc.Find("#main").Empty()

	assertLength(t, doc.Find("#main").Children().Nodes, 0)
	assertLength(t, s.Filter("div").Nodes, 6)
	printSel(t, doc.Selection)
}

func TestRemove(t *testing.T) {
	doc := Doc2Clone()
	doc.Find("#nf1").Remove()

	assertLength(t, doc.Find("#foot #nf1").Nodes, 0)
	printSel(t, doc.Selection)
}

func TestRemoveAll(t *testing.T) {
	doc := Doc2Clone()
	doc.Find("*").Remove()

	assertLength(t, doc.Find("*").Nodes, 0)
	printSel(t, doc.Selection)
}

func TestRemoveRoot(t *testing.T) {
	doc := Doc2Clone()
	doc.Find("html").Remove()

	assertLength(t, doc.Find("html").Nodes, 0)
	printSel(t, doc.Selection)
}

func TestRemoveFiltered(t *testing.T) {
	doc := Doc2Clone()
	nf6 := doc.Find("#nf6")
	s := doc.Find("div").RemoveFiltered("#nf6")

	assertLength(t, doc.Find("#nf6").Nodes, 0)
	assertLength(t, s.Nodes, 1)
	if nf6.Nodes[0] != s.Nodes[0] {
		t.Error("Removed node does not match original")
	}
	printSel(t, doc.Selection)
}

func TestWrap(t *testing.T) {
	doc := Doc2Clone()
	doc.Find("#nf1").Wrap("#nf2")
	nf1 := doc.Find("#foot #nf2 #nf1")
	assertLength(t, nf1.Nodes, 1)

	nf2 := doc.Find("#nf2")
	assertLength(t, nf2.Nodes, 2)

	printSel(t, doc.Selection)
}

func TestWrapEmpty(t *testing.T) {
	doc := Doc2Clone()
	doc.Find("#nf1").Wrap("#doesnt-exist")

	origHtml, _ := Doc2().Html()
	newHtml, _ := doc.Html()

	if origHtml != newHtml {
		t.Error("Expected the two documents to be identical.")
	}

	printSel(t, doc.Selection)
}

func TestWrapHtml(t *testing.T) {
	doc := Doc2Clone()
	doc.Find("#nf1").WrapHtml(wrapHtml)
	nf1 := doc.Find("#foot div#ins:first-child div p em b #nf1")
	assertLength(t, nf1.Nodes, 1)
	printSel(t, doc.Selection)
}

func TestWrapSelection(t *testing.T) {
	doc := Doc2Clone()
	doc.Find("#nf1").WrapSelection(doc.Find("#nf2"))
	nf1 := doc.Find("#foot #nf2 #nf1")
	assertLength(t, nf1.Nodes, 1)

	nf2 := doc.Find("#nf2")
	assertLength(t, nf2.Nodes, 2)

	printSel(t, doc.Selection)
}

func TestWrapAll(t *testing.T) {
	doc := Doc2Clone()
	doc.Find(".odd").WrapAll("#nf1")
	nf1 := doc.Find("#main #nf1")
	assertLength(t, nf1.Nodes, 1)

	sel := nf1.Find("#n2 ~ #n4 ~ #n6 ~ #nf2 ~ #nf4 ~ #nf6")
	assertLength(t, sel.Nodes, 1)

	printSel(t, doc.Selection)
}

func TestWrapAllHtml(t *testing.T) {
	doc := Doc2Clone()
	doc.Find(".odd").WrapAllHtml(wrapHtml)
	nf1 := doc.Find("#main div#ins div p em b #n2 ~ #n4 ~ #n6 ~ #nf2 ~ #nf4 ~ #nf6")
	assertLength(t, nf1.Nodes, 1)
	printSel(t, doc.Selection)
}

func TestWrapInner(t *testing.T) {
	doc := Doc2Clone()
	doc.Find(".one").WrapInner(".two")

	nf1 := doc.Find("#foot #n2 ~ #nf2")
	assertLength(t, nf1.Nodes, 1)

	printSel(t, doc.Selection)
}

func TestWrapInnerHtml(t *testing.T) {
	doc := Doc2Clone()
	doc.Find("#foot").WrapInnerHtml(wrapHtml)

	foot := doc.Find("#foot div#ins div p em b #nf1 ~ #nf2 ~ #nf3")
	assertLength(t, foot.Nodes, 1)

	printSel(t, doc.Selection)
}