// Code generated by "getter -type Grammar,GrammarRef,Custom,String,Structure,StructRef,Binary,Number,Offset,ScriptElement /Users/bramp/go/src/bramp.net/dsector/ufwb"; DO NOT EDIT

package ufwb

func (b *Binary) Description() string {
	if b.description != "" {
		return b.description
	}
	return ""
}

func (b *Binary) SetDescription(description string) {
	b.description = description
}

func (b *Binary) ElemType() string {
	if b.elemType != "" {
		return b.elemType
	}
	return ""
}

func (b *Binary) FillColour() Colour {
	if b.fillColour != nil {
		return *b.fillColour
	}
	if b.extends != nil {
		return b.extends.FillColour()
	}
	return White
}

func (b *Binary) SetFillColour(fillColour Colour) {
	b.fillColour = &fillColour
}

func (b *Binary) Id() int {
	if b.id != 0 {
		return b.id
	}
	return 0
}

func (b *Binary) SetId(id int) {
	b.id = id
}

func (b *Binary) Length() Reference {
	if b.length != Reference("") {
		return b.length
	}
	if b.extends != nil {
		return b.extends.Length()
	}
	return Reference("")
}

func (b *Binary) SetLength(length Reference) {
	b.length = length
}

func (b *Binary) LengthUnit() LengthUnit {
	if b.lengthUnit != LengthUnit(0) {
		return b.lengthUnit
	}
	if b.extends != nil {
		return b.extends.LengthUnit()
	}
	if b.parent != nil {
		return b.parent.LengthUnit()
	}
	return ByteLengthUnit
}

func (b *Binary) SetLengthUnit(lengthUnit LengthUnit) {
	b.lengthUnit = lengthUnit
}

func (b *Binary) MustMatch() Bool {
	if b.mustMatch != Bool(0) {
		return b.mustMatch
	}
	if b.extends != nil {
		return b.extends.MustMatch()
	}
	return True
}

func (b *Binary) SetMustMatch(mustMatch Bool) {
	b.mustMatch = mustMatch
}

func (b *Binary) Name() string {
	if b.name != "" {
		return b.name
	}
	return ""
}

func (b *Binary) SetName(name string) {
	b.name = name
}

func (b *Binary) RepeatMax() Reference {
	if b.repeatMax != Reference("") {
		return b.repeatMax
	}
	if b.extends != nil {
		return b.extends.RepeatMax()
	}
	return Reference("1")
}

func (b *Binary) SetRepeatMax(repeatMax Reference) {
	b.repeatMax = repeatMax
}

func (b *Binary) RepeatMin() Reference {
	if b.repeatMin != Reference("") {
		return b.repeatMin
	}
	if b.extends != nil {
		return b.extends.RepeatMin()
	}
	return Reference("1")
}

func (b *Binary) SetRepeatMin(repeatMin Reference) {
	b.repeatMin = repeatMin
}

func (b *Binary) StrokeColour() Colour {
	if b.strokeColour != nil {
		return *b.strokeColour
	}
	if b.extends != nil {
		return b.extends.StrokeColour()
	}
	return Black
}

func (b *Binary) SetStrokeColour(strokeColour Colour) {
	b.strokeColour = &strokeColour
}

func (b *Binary) Values() []*FixedBinaryValue {
	if b.values != nil {
		return b.values
	}
	if b.extends != nil {
		return b.extends.Values()
	}
	return nil
}

func (b *Binary) SetValues(values []*FixedBinaryValue) {
	b.values = values
}

func (c *Custom) Description() string {
	if c.description != "" {
		return c.description
	}
	return ""
}

func (c *Custom) SetDescription(description string) {
	c.description = description
}

func (c *Custom) ElemType() string {
	if c.elemType != "" {
		return c.elemType
	}
	return ""
}

func (c *Custom) FillColour() Colour {
	if c.fillColour != nil {
		return *c.fillColour
	}
	if c.extends != nil {
		return c.extends.FillColour()
	}
	return White
}

func (c *Custom) SetFillColour(fillColour Colour) {
	c.fillColour = &fillColour
}

func (c *Custom) Id() int {
	if c.id != 0 {
		return c.id
	}
	return 0
}

func (c *Custom) SetId(id int) {
	c.id = id
}

func (c *Custom) Length() Reference {
	if c.length != Reference("") {
		return c.length
	}
	if c.extends != nil {
		return c.extends.Length()
	}
	return Reference("")
}

func (c *Custom) SetLength(length Reference) {
	c.length = length
}

func (c *Custom) LengthUnit() LengthUnit {
	if c.lengthUnit != LengthUnit(0) {
		return c.lengthUnit
	}
	if c.extends != nil {
		return c.extends.LengthUnit()
	}
	return ByteLengthUnit
}

func (c *Custom) SetLengthUnit(lengthUnit LengthUnit) {
	c.lengthUnit = lengthUnit
}

func (c *Custom) Name() string {
	if c.name != "" {
		return c.name
	}
	return ""
}

func (c *Custom) SetName(name string) {
	c.name = name
}

func (c *Custom) Script() *ScriptElement {
	if c.script != nil {
		return c.script
	}
	if c.extends != nil {
		return c.extends.Script()
	}
	return nil
}

func (c *Custom) SetScript(script *ScriptElement) {
	c.script = script
}

func (c *Custom) StrokeColour() Colour {
	if c.strokeColour != nil {
		return *c.strokeColour
	}
	if c.extends != nil {
		return c.extends.StrokeColour()
	}
	return Black
}

func (c *Custom) SetStrokeColour(strokeColour Colour) {
	c.strokeColour = &strokeColour
}

func (g *Grammar) Description() string {
	if g.description != "" {
		return g.description
	}
	return ""
}

func (g *Grammar) SetDescription(description string) {
	g.description = description
}

func (g *Grammar) ElemType() string {
	if g.elemType != "" {
		return g.elemType
	}
	return ""
}

func (g *Grammar) Id() int {
	if g.id != 0 {
		return g.id
	}
	return 0
}

func (g *Grammar) SetId(id int) {
	g.id = id
}

func (g *Grammar) Name() string {
	if g.name != "" {
		return g.name
	}
	return ""
}

func (g *Grammar) SetName(name string) {
	g.name = name
}

func (g *Grammar) RepeatMax() Reference {
	if g.repeatMax != Reference("") {
		return g.repeatMax
	}
	return Reference("1")
}

func (g *Grammar) SetRepeatMax(repeatMax Reference) {
	g.repeatMax = repeatMax
}

func (g *Grammar) RepeatMin() Reference {
	if g.repeatMin != Reference("") {
		return g.repeatMin
	}
	return Reference("1")
}

func (g *Grammar) SetRepeatMin(repeatMin Reference) {
	g.repeatMin = repeatMin
}

func (g *GrammarRef) Description() string {
	if g.description != "" {
		return g.description
	}
	return ""
}

func (g *GrammarRef) SetDescription(description string) {
	g.description = description
}

func (g *GrammarRef) Disabled() Bool {
	if g.disabled != Bool(0) {
		return g.disabled
	}
	if g.extends != nil {
		return g.extends.Disabled()
	}
	return Bool(0)
}

func (g *GrammarRef) SetDisabled(disabled Bool) {
	g.disabled = disabled
}

func (g *GrammarRef) ElemType() string {
	if g.elemType != "" {
		return g.elemType
	}
	return ""
}

func (g *GrammarRef) Filename() string {
	if g.filename != "" {
		return g.filename
	}
	if g.extends != nil {
		return g.extends.Filename()
	}
	return ""
}

func (g *GrammarRef) SetFilename(filename string) {
	g.filename = filename
}

func (g *GrammarRef) Grammar() *Grammar {
	if g.grammar != nil {
		return g.grammar
	}
	if g.extends != nil {
		return g.extends.Grammar()
	}
	return nil
}

func (g *GrammarRef) SetGrammar(grammar *Grammar) {
	g.grammar = grammar
}

func (g *GrammarRef) Id() int {
	if g.id != 0 {
		return g.id
	}
	return 0
}

func (g *GrammarRef) SetId(id int) {
	g.id = id
}

func (g *GrammarRef) Name() string {
	if g.name != "" {
		return g.name
	}
	return ""
}

func (g *GrammarRef) SetName(name string) {
	g.name = name
}

func (g *GrammarRef) RepeatMax() Reference {
	if g.repeatMax != Reference("") {
		return g.repeatMax
	}
	if g.extends != nil {
		return g.extends.RepeatMax()
	}
	return Reference("1")
}

func (g *GrammarRef) SetRepeatMax(repeatMax Reference) {
	g.repeatMax = repeatMax
}

func (g *GrammarRef) RepeatMin() Reference {
	if g.repeatMin != Reference("") {
		return g.repeatMin
	}
	if g.extends != nil {
		return g.extends.RepeatMin()
	}
	return Reference("1")
}

func (g *GrammarRef) SetRepeatMin(repeatMin Reference) {
	g.repeatMin = repeatMin
}

func (g *GrammarRef) Uti() string {
	if g.uti != "" {
		return g.uti
	}
	if g.extends != nil {
		return g.extends.Uti()
	}
	return ""
}

func (g *GrammarRef) SetUti(uti string) {
	g.uti = uti
}

func (n *Number) Description() string {
	if n.description != "" {
		return n.description
	}
	return ""
}

func (n *Number) SetDescription(description string) {
	n.description = description
}

func (n *Number) Display() Display {
	if n.display != Display(0) {
		return n.display
	}
	if n.extends != nil {
		return n.extends.Display()
	}
	if n.parent != nil {
		return n.parent.Display()
	}
	return DecDisplay
}

func (n *Number) SetDisplay(display Display) {
	n.display = display
}

func (n *Number) ElemType() string {
	if n.elemType != "" {
		return n.elemType
	}
	return ""
}

func (n *Number) Endian() Endian {
	if n.endian != Endian(0) {
		return n.endian
	}
	if n.extends != nil {
		return n.extends.Endian()
	}
	if n.parent != nil {
		return n.parent.Endian()
	}
	return LittleEndian
}

func (n *Number) SetEndian(endian Endian) {
	n.endian = endian
}

func (n *Number) FillColour() Colour {
	if n.fillColour != nil {
		return *n.fillColour
	}
	if n.extends != nil {
		return n.extends.FillColour()
	}
	return White
}

func (n *Number) SetFillColour(fillColour Colour) {
	n.fillColour = &fillColour
}

func (n *Number) Id() int {
	if n.id != 0 {
		return n.id
	}
	return 0
}

func (n *Number) SetId(id int) {
	n.id = id
}

func (n *Number) Length() Reference {
	if n.length != Reference("") {
		return n.length
	}
	if n.extends != nil {
		return n.extends.Length()
	}
	return Reference("")
}

func (n *Number) SetLength(length Reference) {
	n.length = length
}

func (n *Number) LengthUnit() LengthUnit {
	if n.lengthUnit != LengthUnit(0) {
		return n.lengthUnit
	}
	if n.extends != nil {
		return n.extends.LengthUnit()
	}
	if n.parent != nil {
		return n.parent.LengthUnit()
	}
	return ByteLengthUnit
}

func (n *Number) SetLengthUnit(lengthUnit LengthUnit) {
	n.lengthUnit = lengthUnit
}

func (n *Number) Masks() []*Mask {
	if n.masks != nil {
		return n.masks
	}
	if n.extends != nil {
		return n.extends.Masks()
	}
	return nil
}

func (n *Number) SetMasks(masks []*Mask) {
	n.masks = masks
}

func (n *Number) MaxVal() string {
	if n.maxVal != "" {
		return n.maxVal
	}
	if n.extends != nil {
		return n.extends.MaxVal()
	}
	return ""
}

func (n *Number) SetMaxVal(maxVal string) {
	n.maxVal = maxVal
}

func (n *Number) MinVal() string {
	if n.minVal != "" {
		return n.minVal
	}
	if n.extends != nil {
		return n.extends.MinVal()
	}
	return ""
}

func (n *Number) SetMinVal(minVal string) {
	n.minVal = minVal
}

func (n *Number) MustMatch() Bool {
	if n.mustMatch != Bool(0) {
		return n.mustMatch
	}
	if n.extends != nil {
		return n.extends.MustMatch()
	}
	return True
}

func (n *Number) SetMustMatch(mustMatch Bool) {
	n.mustMatch = mustMatch
}

func (n *Number) Name() string {
	if n.name != "" {
		return n.name
	}
	return ""
}

func (n *Number) SetName(name string) {
	n.name = name
}

func (n *Number) RepeatMax() Reference {
	if n.repeatMax != Reference("") {
		return n.repeatMax
	}
	if n.extends != nil {
		return n.extends.RepeatMax()
	}
	return Reference("1")
}

func (n *Number) SetRepeatMax(repeatMax Reference) {
	n.repeatMax = repeatMax
}

func (n *Number) RepeatMin() Reference {
	if n.repeatMin != Reference("") {
		return n.repeatMin
	}
	if n.extends != nil {
		return n.extends.RepeatMin()
	}
	return Reference("1")
}

func (n *Number) SetRepeatMin(repeatMin Reference) {
	n.repeatMin = repeatMin
}

func (n *Number) StrokeColour() Colour {
	if n.strokeColour != nil {
		return *n.strokeColour
	}
	if n.extends != nil {
		return n.extends.StrokeColour()
	}
	return Black
}

func (n *Number) SetStrokeColour(strokeColour Colour) {
	n.strokeColour = &strokeColour
}

func (n *Number) ValueExpression() string {
	if n.valueExpression != "" {
		return n.valueExpression
	}
	if n.extends != nil {
		return n.extends.ValueExpression()
	}
	return ""
}

func (n *Number) SetValueExpression(valueExpression string) {
	n.valueExpression = valueExpression
}

func (n *Number) Values() []*FixedValue {
	if n.values != nil {
		return n.values
	}
	if n.extends != nil {
		return n.extends.Values()
	}
	return nil
}

func (n *Number) SetValues(values []*FixedValue) {
	n.values = values
}

func (o *Offset) Additional() string {
	if o.additional != "" {
		return o.additional
	}
	if o.extends != nil {
		return o.extends.Additional()
	}
	return ""
}

func (o *Offset) SetAdditional(additional string) {
	o.additional = additional
}

func (o *Offset) Description() string {
	if o.description != "" {
		return o.description
	}
	return ""
}

func (o *Offset) SetDescription(description string) {
	o.description = description
}

func (o *Offset) Display() Display {
	if o.display != Display(0) {
		return o.display
	}
	if o.extends != nil {
		return o.extends.Display()
	}
	if o.parent != nil {
		return o.parent.Display()
	}
	return DecDisplay
}

func (o *Offset) SetDisplay(display Display) {
	o.display = display
}

func (o *Offset) ElemType() string {
	if o.elemType != "" {
		return o.elemType
	}
	return ""
}

func (o *Offset) Endian() Endian {
	if o.endian != Endian(0) {
		return o.endian
	}
	if o.extends != nil {
		return o.extends.Endian()
	}
	if o.parent != nil {
		return o.parent.Endian()
	}
	return LittleEndian
}

func (o *Offset) SetEndian(endian Endian) {
	o.endian = endian
}

func (o *Offset) FillColour() Colour {
	if o.fillColour != nil {
		return *o.fillColour
	}
	if o.extends != nil {
		return o.extends.FillColour()
	}
	return White
}

func (o *Offset) SetFillColour(fillColour Colour) {
	o.fillColour = &fillColour
}

func (o *Offset) FollowNullReference() Bool {
	if o.followNullReference != Bool(0) {
		return o.followNullReference
	}
	if o.extends != nil {
		return o.extends.FollowNullReference()
	}
	return Bool(0)
}

func (o *Offset) SetFollowNullReference(followNullReference Bool) {
	o.followNullReference = followNullReference
}

func (o *Offset) Id() int {
	if o.id != 0 {
		return o.id
	}
	return 0
}

func (o *Offset) SetId(id int) {
	o.id = id
}

func (o *Offset) Length() Reference {
	if o.length != Reference("") {
		return o.length
	}
	if o.extends != nil {
		return o.extends.Length()
	}
	return Reference("")
}

func (o *Offset) SetLength(length Reference) {
	o.length = length
}

func (o *Offset) LengthUnit() LengthUnit {
	if o.lengthUnit != LengthUnit(0) {
		return o.lengthUnit
	}
	if o.extends != nil {
		return o.extends.LengthUnit()
	}
	if o.parent != nil {
		return o.parent.LengthUnit()
	}
	return ByteLengthUnit
}

func (o *Offset) SetLengthUnit(lengthUnit LengthUnit) {
	o.lengthUnit = lengthUnit
}

func (o *Offset) Name() string {
	if o.name != "" {
		return o.name
	}
	return ""
}

func (o *Offset) SetName(name string) {
	o.name = name
}

func (o *Offset) ReferencedSize() Element {
	if o.referencedSize != Element(nil) {
		return o.referencedSize
	}
	if o.extends != nil {
		return o.extends.ReferencedSize()
	}
	return Element(nil)
}

func (o *Offset) SetReferencedSize(referencedSize Element) {
	o.referencedSize = referencedSize
}

func (o *Offset) References() Element {
	if o.references != Element(nil) {
		return o.references
	}
	if o.extends != nil {
		return o.extends.References()
	}
	return Element(nil)
}

func (o *Offset) SetReferences(references Element) {
	o.references = references
}

func (o *Offset) RelativeTo() Element {
	if o.relativeTo != Element(nil) {
		return o.relativeTo
	}
	if o.extends != nil {
		return o.extends.RelativeTo()
	}
	return Element(nil)
}

func (o *Offset) SetRelativeTo(relativeTo Element) {
	o.relativeTo = relativeTo
}

func (o *Offset) RepeatMax() Reference {
	if o.repeatMax != Reference("") {
		return o.repeatMax
	}
	if o.extends != nil {
		return o.extends.RepeatMax()
	}
	return Reference("1")
}

func (o *Offset) SetRepeatMax(repeatMax Reference) {
	o.repeatMax = repeatMax
}

func (o *Offset) RepeatMin() Reference {
	if o.repeatMin != Reference("") {
		return o.repeatMin
	}
	if o.extends != nil {
		return o.extends.RepeatMin()
	}
	return Reference("1")
}

func (o *Offset) SetRepeatMin(repeatMin Reference) {
	o.repeatMin = repeatMin
}

func (o *Offset) StrokeColour() Colour {
	if o.strokeColour != nil {
		return *o.strokeColour
	}
	if o.extends != nil {
		return o.extends.StrokeColour()
	}
	return Black
}

func (o *Offset) SetStrokeColour(strokeColour Colour) {
	o.strokeColour = &strokeColour
}

func (s *ScriptElement) Description() string {
	if s.description != "" {
		return s.description
	}
	return ""
}

func (s *ScriptElement) SetDescription(description string) {
	s.description = description
}

func (s *ScriptElement) ElemType() string {
	if s.elemType != "" {
		return s.elemType
	}
	return ""
}

func (s *ScriptElement) Id() int {
	if s.id != 0 {
		return s.id
	}
	return 0
}

func (s *ScriptElement) SetId(id int) {
	s.id = id
}

func (s *ScriptElement) Name() string {
	if s.name != "" {
		return s.name
	}
	return ""
}

func (s *ScriptElement) SetName(name string) {
	s.name = name
}

func (s *ScriptElement) RepeatMax() Reference {
	if s.repeatMax != Reference("") {
		return s.repeatMax
	}
	if s.extends != nil {
		return s.extends.RepeatMax()
	}
	return Reference("1")
}

func (s *ScriptElement) SetRepeatMax(repeatMax Reference) {
	s.repeatMax = repeatMax
}

func (s *ScriptElement) RepeatMin() Reference {
	if s.repeatMin != Reference("") {
		return s.repeatMin
	}
	if s.extends != nil {
		return s.extends.RepeatMin()
	}
	return Reference("1")
}

func (s *ScriptElement) SetRepeatMin(repeatMin Reference) {
	s.repeatMin = repeatMin
}

func (s *String) Delimiter() byte {
	if s.delimiter != 0 {
		return s.delimiter
	}
	if s.extends != nil {
		return s.extends.Delimiter()
	}
	return 0
}

func (s *String) SetDelimiter(delimiter byte) {
	s.delimiter = delimiter
}

func (s *String) Description() string {
	if s.description != "" {
		return s.description
	}
	return ""
}

func (s *String) SetDescription(description string) {
	s.description = description
}

func (s *String) ElemType() string {
	if s.elemType != "" {
		return s.elemType
	}
	return ""
}

func (s *String) Encoding() string {
	if s.encoding != "" {
		return s.encoding
	}
	if s.extends != nil {
		return s.extends.Encoding()
	}
	if s.parent != nil {
		return s.parent.Encoding()
	}
	return "UTF-8"
}

func (s *String) SetEncoding(encoding string) {
	s.encoding = encoding
}

func (s *String) FillColour() Colour {
	if s.fillColour != nil {
		return *s.fillColour
	}
	if s.extends != nil {
		return s.extends.FillColour()
	}
	return White
}

func (s *String) SetFillColour(fillColour Colour) {
	s.fillColour = &fillColour
}

func (s *String) Id() int {
	if s.id != 0 {
		return s.id
	}
	return 0
}

func (s *String) SetId(id int) {
	s.id = id
}

func (s *String) Length() Reference {
	if s.length != Reference("") {
		return s.length
	}
	if s.extends != nil {
		return s.extends.Length()
	}
	return Reference("")
}

func (s *String) SetLength(length Reference) {
	s.length = length
}

func (s *String) LengthUnit() LengthUnit {
	if s.lengthUnit != LengthUnit(0) {
		return s.lengthUnit
	}
	if s.extends != nil {
		return s.extends.LengthUnit()
	}
	if s.parent != nil {
		return s.parent.LengthUnit()
	}
	return ByteLengthUnit
}

func (s *String) SetLengthUnit(lengthUnit LengthUnit) {
	s.lengthUnit = lengthUnit
}

func (s *String) MustMatch() Bool {
	if s.mustMatch != Bool(0) {
		return s.mustMatch
	}
	if s.extends != nil {
		return s.extends.MustMatch()
	}
	return True
}

func (s *String) SetMustMatch(mustMatch Bool) {
	s.mustMatch = mustMatch
}

func (s *String) Name() string {
	if s.name != "" {
		return s.name
	}
	return ""
}

func (s *String) SetName(name string) {
	s.name = name
}

func (s *String) RepeatMax() Reference {
	if s.repeatMax != Reference("") {
		return s.repeatMax
	}
	if s.extends != nil {
		return s.extends.RepeatMax()
	}
	return Reference("1")
}

func (s *String) SetRepeatMax(repeatMax Reference) {
	s.repeatMax = repeatMax
}

func (s *String) RepeatMin() Reference {
	if s.repeatMin != Reference("") {
		return s.repeatMin
	}
	if s.extends != nil {
		return s.extends.RepeatMin()
	}
	return Reference("1")
}

func (s *String) SetRepeatMin(repeatMin Reference) {
	s.repeatMin = repeatMin
}

func (s *String) StrokeColour() Colour {
	if s.strokeColour != nil {
		return *s.strokeColour
	}
	if s.extends != nil {
		return s.extends.StrokeColour()
	}
	return Black
}

func (s *String) SetStrokeColour(strokeColour Colour) {
	s.strokeColour = &strokeColour
}

func (s *String) Typ() string {
	if s.typ != "" {
		return s.typ
	}
	if s.extends != nil {
		return s.extends.Typ()
	}
	return ""
}

func (s *String) SetTyp(typ string) {
	s.typ = typ
}

func (s *String) Values() []*FixedStringValue {
	if s.values != nil {
		return s.values
	}
	if s.extends != nil {
		return s.extends.Values()
	}
	return nil
}

func (s *String) SetValues(values []*FixedStringValue) {
	s.values = values
}

func (s *StructRef) Description() string {
	if s.description != "" {
		return s.description
	}
	return ""
}

func (s *StructRef) SetDescription(description string) {
	s.description = description
}

func (s *StructRef) ElemType() string {
	if s.elemType != "" {
		return s.elemType
	}
	return ""
}

func (s *StructRef) FillColour() Colour {
	if s.fillColour != nil {
		return *s.fillColour
	}
	if s.extends != nil {
		return s.extends.FillColour()
	}
	return White
}

func (s *StructRef) SetFillColour(fillColour Colour) {
	s.fillColour = &fillColour
}

func (s *StructRef) Id() int {
	if s.id != 0 {
		return s.id
	}
	return 0
}

func (s *StructRef) SetId(id int) {
	s.id = id
}

func (s *StructRef) Name() string {
	if s.name != "" {
		return s.name
	}
	return ""
}

func (s *StructRef) SetName(name string) {
	s.name = name
}

func (s *StructRef) RepeatMax() Reference {
	if s.repeatMax != Reference("") {
		return s.repeatMax
	}
	if s.extends != nil {
		return s.extends.RepeatMax()
	}
	return Reference("1")
}

func (s *StructRef) SetRepeatMax(repeatMax Reference) {
	s.repeatMax = repeatMax
}

func (s *StructRef) RepeatMin() Reference {
	if s.repeatMin != Reference("") {
		return s.repeatMin
	}
	if s.extends != nil {
		return s.extends.RepeatMin()
	}
	return Reference("1")
}

func (s *StructRef) SetRepeatMin(repeatMin Reference) {
	s.repeatMin = repeatMin
}

func (s *StructRef) StrokeColour() Colour {
	if s.strokeColour != nil {
		return *s.strokeColour
	}
	if s.extends != nil {
		return s.extends.StrokeColour()
	}
	return Black
}

func (s *StructRef) SetStrokeColour(strokeColour Colour) {
	s.strokeColour = &strokeColour
}

func (s *StructRef) Structure() *Structure {
	if s.structure != nil {
		return s.structure
	}
	if s.extends != nil {
		return s.extends.Structure()
	}
	return nil
}

func (s *StructRef) SetStructure(structure *Structure) {
	s.structure = structure
}

func (s *Structure) Description() string {
	if s.description != "" {
		return s.description
	}
	return ""
}

func (s *Structure) SetDescription(description string) {
	s.description = description
}

func (s *Structure) Display() Display {
	if s.display != Display(0) {
		return s.display
	}
	if s.extends != nil {
		return s.extends.Display()
	}
	if s.parent != nil {
		return s.parent.Display()
	}
	return DecDisplay
}

func (s *Structure) SetDisplay(display Display) {
	s.display = display
}

func (s *Structure) ElemType() string {
	if s.elemType != "" {
		return s.elemType
	}
	return ""
}

func (s *Structure) SetElements(elements []Element) {
	s.elements = elements
}

func (s *Structure) Encoding() string {
	if s.encoding != "" {
		return s.encoding
	}
	if s.extends != nil {
		return s.extends.Encoding()
	}
	if s.parent != nil {
		return s.parent.Encoding()
	}
	return "UTF-8"
}

func (s *Structure) SetEncoding(encoding string) {
	s.encoding = encoding
}

func (s *Structure) Endian() Endian {
	if s.endian != Endian(0) {
		return s.endian
	}
	if s.extends != nil {
		return s.extends.Endian()
	}
	if s.parent != nil {
		return s.parent.Endian()
	}
	return LittleEndian
}

func (s *Structure) SetEndian(endian Endian) {
	s.endian = endian
}

func (s *Structure) FillColour() Colour {
	if s.fillColour != nil {
		return *s.fillColour
	}
	if s.extends != nil {
		return s.extends.FillColour()
	}
	return White
}

func (s *Structure) SetFillColour(fillColour Colour) {
	s.fillColour = &fillColour
}

func (s *Structure) Id() int {
	if s.id != 0 {
		return s.id
	}
	return 0
}

func (s *Structure) SetId(id int) {
	s.id = id
}

func (s *Structure) Length() Reference {
	if s.length != Reference("") {
		return s.length
	}
	if s.extends != nil {
		return s.extends.Length()
	}
	return Reference("")
}

func (s *Structure) SetLength(length Reference) {
	s.length = length
}

func (s *Structure) LengthOffset() Reference {
	if s.lengthOffset != Reference("") {
		return s.lengthOffset
	}
	if s.extends != nil {
		return s.extends.LengthOffset()
	}
	if s.parent != nil {
		return s.parent.LengthOffset()
	}
	return Reference("")
}

func (s *Structure) SetLengthOffset(lengthOffset Reference) {
	s.lengthOffset = lengthOffset
}

func (s *Structure) LengthUnit() LengthUnit {
	if s.lengthUnit != LengthUnit(0) {
		return s.lengthUnit
	}
	if s.extends != nil {
		return s.extends.LengthUnit()
	}
	if s.parent != nil {
		return s.parent.LengthUnit()
	}
	return ByteLengthUnit
}

func (s *Structure) SetLengthUnit(lengthUnit LengthUnit) {
	s.lengthUnit = lengthUnit
}

func (s *Structure) Name() string {
	if s.name != "" {
		return s.name
	}
	return ""
}

func (s *Structure) SetName(name string) {
	s.name = name
}

func (s *Structure) Order() Order {
	if s.order != Order(0) {
		return s.order
	}
	if s.extends != nil {
		return s.extends.Order()
	}
	if s.parent != nil {
		return s.parent.Order()
	}
	return FixedOrder
}

func (s *Structure) SetOrder(order Order) {
	s.order = order
}

func (s *Structure) RepeatMax() Reference {
	if s.repeatMax != Reference("") {
		return s.repeatMax
	}
	if s.extends != nil {
		return s.extends.RepeatMax()
	}
	return Reference("1")
}

func (s *Structure) SetRepeatMax(repeatMax Reference) {
	s.repeatMax = repeatMax
}

func (s *Structure) RepeatMin() Reference {
	if s.repeatMin != Reference("") {
		return s.repeatMin
	}
	if s.extends != nil {
		return s.extends.RepeatMin()
	}
	return Reference("1")
}

func (s *Structure) SetRepeatMin(repeatMin Reference) {
	s.repeatMin = repeatMin
}

func (s *Structure) StrokeColour() Colour {
	if s.strokeColour != nil {
		return *s.strokeColour
	}
	if s.extends != nil {
		return s.extends.StrokeColour()
	}
	return Black
}

func (s *Structure) SetStrokeColour(strokeColour Colour) {
	s.strokeColour = &strokeColour
}
