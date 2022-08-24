package asty

import (
	"go/ast"
	"go/token"
	"os"
)

type Marshaller struct {
	fset          *token.FileSet
	WithPositions bool
	WithComments  bool
}

func NewMarshaller(withComments, withPositions bool) *Marshaller {
	return &Marshaller{
		WithComments:  withComments,
		WithPositions: withPositions,
		fset:          token.NewFileSet(),
	}
}

func (m *Marshaller) FileSet() *token.FileSet {
	return m.fset
}

func (m *Marshaller) AddFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	info, err := file.Stat()
	if err != nil {
		return err
	}

	size := info.Size()
	err = file.Close()
	if err != nil {
		return err
	}

	m.fset.AddFile(filename, -1, int(size))
	return nil
}

// ---------------------------------------------------------------------------

func (m *Marshaller) MarshalNode(nodeType string) Node {
	return Node{
		NodeType: nodeType,
	}
}

func (m *Marshaller) MarshalPosition(pos token.Pos) *PositionNode {
	if !m.WithPositions {
		return nil
	}
	position := m.fset.Position(pos)
	return &PositionNode{
		Node:     m.MarshalNode("Position"),
		Filename: position.Filename,
		Line:     position.Line,
		Offset:   position.Offset,
		Column:   position.Column,
	}
}

func (m *Marshaller) MarshalComment(comment *ast.Comment) *CommentNode {
	if comment == nil {
		return nil
	}
	return &CommentNode{
		Node:  m.MarshalNode("Comment"),
		Slash: m.MarshalPosition(comment.Slash),
		Text:  comment.Text,
	}
}

func (m *Marshaller) MarshalComments(node []*ast.Comment) []*CommentNode {
	if node == nil {
		return nil
	}
	nodes := make([]*CommentNode, len(node))
	for index, comment := range node {
		nodes[index] = m.MarshalComment(comment)
	}
	return nodes
}

func (m *Marshaller) MarshalCommentGroup(group *ast.CommentGroup) *CommentGroupNode {
	if !m.WithComments {
		return nil
	}
	if group == nil {
		return nil
	}
	return &CommentGroupNode{
		Node: m.MarshalNode("CommentGroup"),
		List: m.MarshalComments(group.List),
	}
}

// ---------------------------------------------------------------------------

func (m *Marshaller) MarshalField(node *ast.Field) *FieldNode {
	return &FieldNode{
		Node:    m.MarshalNode("Field"),
		Doc:     m.MarshalCommentGroup(node.Doc),
		Names:   m.MarshalIdents(node.Names),
		Type:    m.MarshalExpr(node.Type),
		Tag:     m.MarshalBasicLit(node.Tag),
		Comment: m.MarshalCommentGroup(node.Comment),
	}
}

func (m *Marshaller) MarshalFields(fields []*ast.Field) []*FieldNode {
	if fields == nil {
		return nil
	}
	nodes := make([]*FieldNode, len(fields))
	for index, field := range fields {
		nodes[index] = m.MarshalField(field)
	}
	return nodes
}

func (m *Marshaller) MarshalFieldList(node *ast.FieldList) *FieldListNode {
	if node == nil {
		return nil
	}
	return &FieldListNode{
		Node:    m.MarshalNode("FieldList"),
		Opening: m.MarshalPosition(node.Opening),
		List:    m.MarshalFields(node.List),
		Closing: m.MarshalPosition(node.Closing),
	}
}

// ---------------------------------------------------------------------------

func (m *Marshaller) MarshalBadExpr(node *ast.BadExpr) *BadExprNode {
	return &BadExprNode{
		Node: m.MarshalNode("BadExpr"),
		From: m.MarshalPosition(node.From),
		To:   m.MarshalPosition(node.To),
	}
}

func (m *Marshaller) MarshalIdent(node *ast.Ident) *IdentNode {
	if node == nil {
		return nil
	}
	return &IdentNode{
		Node:    m.MarshalNode("Ident"),
		NamePos: m.MarshalPosition(node.NamePos),
		Name:    node.Name,
	}
}

func (m *Marshaller) MarshalIdents(idents []*ast.Ident) []*IdentNode {
	if idents == nil {
		return nil
	}
	nodes := make([]*IdentNode, len(idents))
	for index, ident := range idents {
		nodes[index] = m.MarshalIdent(ident)
	}
	return nodes
}

func (m *Marshaller) MarshalEllipsis(node *ast.Ellipsis) *EllipsisNode {
	if node == nil {
		return nil
	}
	return &EllipsisNode{
		Node:     m.MarshalNode("Ellipsis"),
		Ellipsis: m.MarshalPosition(node.Ellipsis),
		Elt:      m.MarshalExpr(node.Elt),
	}
}

func (m *Marshaller) MarshalBasicLit(node *ast.BasicLit) *BasicLitNode {
	if node == nil {
		return nil
	}
	return &BasicLitNode{
		Node:     m.MarshalNode("BasicLit"),
		ValuePos: m.MarshalPosition(node.ValuePos),
		Kind:     node.Kind.String(),
		Value:    node.Value,
	}
}

func (m *Marshaller) MarshalFuncLit(node *ast.FuncLit) *FuncLitNode {
	return &FuncLitNode{
		Node: m.MarshalNode("FuncLit"),
		Type: m.MarshalFuncType(node.Type),
		Body: m.MarshalBlockStmt(node.Body),
	}
}

func (m *Marshaller) MarshalCompositeLit(node *ast.CompositeLit) *CompositeLitNode {
	return &CompositeLitNode{
		Node:       m.MarshalNode("CompositeLit"),
		Type:       m.MarshalExpr(node.Type),
		Lbrace:     m.MarshalPosition(node.Lbrace),
		Elts:       m.MarshalExprs(node.Elts),
		Rbrace:     m.MarshalPosition(node.Rbrace),
		Incomplete: node.Incomplete,
	}
}

func (m *Marshaller) MarshalParenExpr(node *ast.ParenExpr) *ParenExprNode {
	return &ParenExprNode{
		Node:   m.MarshalNode("ParenExpr"),
		Lparen: m.MarshalPosition(node.Lparen),
		X:      m.MarshalExpr(node.X),
		Rparen: m.MarshalPosition(node.Rparen),
	}
}

func (m *Marshaller) MarshalSelectorExpr(node *ast.SelectorExpr) *SelectorExprNode {
	return &SelectorExprNode{
		Node: m.MarshalNode("SelectorExpr"),
		X:    m.MarshalExpr(node.X),
		Sel:  m.MarshalIdent(node.Sel),
	}
}

func (m *Marshaller) MarshalIndexExpr(node *ast.IndexExpr) *IndexExprNode {
	return &IndexExprNode{
		Node:   m.MarshalNode("IndexExpr"),
		X:      m.MarshalExpr(node.X),
		Lbrack: m.MarshalPosition(node.Lbrack),
		Index:  m.MarshalExpr(node.Index),
		Rbrack: m.MarshalPosition(node.Rbrack),
	}
}

func (m *Marshaller) MarshalIndexListExpr(node *ast.IndexListExpr) *IndexListExprNode {
	return &IndexListExprNode{
		Node:    m.MarshalNode("IndexListExpr"),
		X:       m.MarshalExpr(node.X),
		Lbrack:  m.MarshalPosition(node.Lbrack),
		Indices: m.MarshalExprs(node.Indices),
		Rbrack:  m.MarshalPosition(node.Rbrack),
	}
}

func (m *Marshaller) MarshalSliceExpr(node *ast.SliceExpr) *SliceExprNode {
	return &SliceExprNode{
		Node:   m.MarshalNode("SliceExpr"),
		X:      m.MarshalExpr(node.X),
		Lbrack: m.MarshalPosition(node.Lbrack),
		Low:    m.MarshalExpr(node.Low),
		High:   m.MarshalExpr(node.High),
		Max:    m.MarshalExpr(node.Max),
		Slice3: node.Slice3,
		Rbrack: m.MarshalPosition(node.Rbrack),
	}
}

func (m *Marshaller) MarshalTypeAssertExpr(node *ast.TypeAssertExpr) *TypeAssertExprNode {
	return &TypeAssertExprNode{
		Node:   m.MarshalNode("TypeAssertExpr"),
		X:      m.MarshalExpr(node.X),
		Lparen: m.MarshalPosition(node.Lparen),
		Type:   m.MarshalExpr(node.Type),
		Rparen: m.MarshalPosition(node.Rparen),
	}
}

func (m *Marshaller) MarshalCallExpr(node *ast.CallExpr) *CallExprNode {
	return &CallExprNode{
		Node:     m.MarshalNode("CallExpr"),
		Fun:      m.MarshalExpr(node.Fun),
		Lparen:   m.MarshalPosition(node.Lparen),
		Args:     m.MarshalExprs(node.Args),
		Ellipsis: m.MarshalPosition(node.Ellipsis),
		Rparen:   m.MarshalPosition(node.Rparen),
	}
}

func (m *Marshaller) MarshalStarExpr(node *ast.StarExpr) *StarExprNode {
	return &StarExprNode{
		Node: m.MarshalNode("StarExpr"),
		Star: m.MarshalPosition(node.Star),
		X:    m.MarshalExpr(node.X),
	}
}

func (m *Marshaller) MarshalUnaryExpr(node *ast.UnaryExpr) *UnaryExprNode {
	return &UnaryExprNode{
		Node:  m.MarshalNode("UnaryExpr"),
		OpPos: m.MarshalPosition(node.OpPos),
		Op:    node.Op.String(),
		X:     m.MarshalExpr(node.X),
	}
}

func (m *Marshaller) MarshalBinaryExpr(node *ast.BinaryExpr) *BinaryExprNode {
	return &BinaryExprNode{
		Node:  m.MarshalNode("BinaryExpr"),
		X:     m.MarshalExpr(node.X),
		OpPos: m.MarshalPosition(node.OpPos),
		Op:    node.Op.String(),
		Y:     m.MarshalExpr(node.Y),
	}
}

func (m *Marshaller) MarshalKeyValueExpr(node *ast.KeyValueExpr) *KeyValueExprNode {
	return &KeyValueExprNode{
		Node:  m.MarshalNode("KeyValueExpr"),
		Key:   m.MarshalExpr(node.Key),
		Colon: m.MarshalPosition(node.Colon),
		Value: m.MarshalExpr(node.Value),
	}
}

func (m *Marshaller) MarshalExpr(node ast.Expr) IExprNode {
	if node == nil {
		return nil
	}
	switch expr := node.(type) {
	case *ast.BadExpr:
		return m.MarshalBadExpr(expr)
	case *ast.Ident:
		return m.MarshalIdent(expr)
	case *ast.Ellipsis:
		return m.MarshalEllipsis(expr)
	case *ast.BasicLit:
		return m.MarshalBasicLit(expr)
	case *ast.FuncLit:
		return m.MarshalFuncLit(expr)
	case *ast.CompositeLit:
		return m.MarshalCompositeLit(expr)
	case *ast.ParenExpr:
		return m.MarshalParenExpr(expr)
	case *ast.SelectorExpr:
		return m.MarshalSelectorExpr(expr)
	case *ast.IndexExpr:
		return m.MarshalIndexExpr(expr)
	case *ast.IndexListExpr:
		return m.MarshalIndexListExpr(expr)
	case *ast.SliceExpr:
		return m.MarshalSliceExpr(expr)
	case *ast.TypeAssertExpr:
		return m.MarshalTypeAssertExpr(expr)
	case *ast.CallExpr:
		return m.MarshalCallExpr(expr)
	case *ast.StarExpr:
		return m.MarshalStarExpr(expr)
	case *ast.UnaryExpr:
		return m.MarshalUnaryExpr(expr)
	case *ast.BinaryExpr:
		return m.MarshalBinaryExpr(expr)
	case *ast.KeyValueExpr:
		return m.MarshalKeyValueExpr(expr)
	case *ast.ArrayType:
		return m.MarshalArrayType(expr)
	case *ast.StructType:
		return m.MarshalStructType(expr)
	case *ast.FuncType:
		return m.MarshalFuncType(expr)
	case *ast.InterfaceType:
		return m.MarshalInterfaceType(expr)
	case *ast.MapType:
		return m.MarshalMapType(expr)
	case *ast.ChanType:
		return m.MarshalChanType(expr)
	default:
		panic("implement me")
	}
}

func (m *Marshaller) MarshalExprs(exprs []ast.Expr) []IExprNode {
	if exprs == nil {
		return nil
	}
	nodes := make([]IExprNode, len(exprs))
	for index, expr := range exprs {
		nodes[index] = m.MarshalExpr(expr)
	}
	return nodes
}

// ---------------------------------------------------------------------------

func (m *Marshaller) MarshalArrayType(node *ast.ArrayType) *ArrayTypeNode {
	return &ArrayTypeNode{
		Node:   m.MarshalNode("ArrayType"),
		Lbrack: m.MarshalPosition(node.Lbrack),
		Len:    m.MarshalExpr(node.Len),
		Elt:    m.MarshalExpr(node.Elt),
	}
}

func (m *Marshaller) MarshalStructType(node *ast.StructType) *StructTypeNode {
	return &StructTypeNode{
		Node:       m.MarshalNode("StructType"),
		Struct:     m.MarshalPosition(node.Struct),
		Fields:     m.MarshalFieldList(node.Fields),
		Incomplete: node.Incomplete,
	}
}

func (m *Marshaller) MarshalFuncType(node *ast.FuncType) *FuncTypeNode {
	return &FuncTypeNode{
		Node:       m.MarshalNode("FuncType"),
		Func:       m.MarshalPosition(node.Func),
		TypeParams: m.MarshalFieldList(node.TypeParams),
		Params:     m.MarshalFieldList(node.Params),
		Results:    m.MarshalFieldList(node.Results),
	}
}

func (m *Marshaller) MarshalInterfaceType(node *ast.InterfaceType) *InterfaceTypeNode {
	return &InterfaceTypeNode{
		Node:       m.MarshalNode("InterfaceType"),
		Interface:  m.MarshalPosition(node.Interface),
		Methods:    m.MarshalFieldList(node.Methods),
		Incomplete: node.Incomplete,
	}
}

func (m *Marshaller) MarshalMapType(node *ast.MapType) *MapTypeNode {
	return &MapTypeNode{
		Node:  m.MarshalNode("MapType"),
		Map:   m.MarshalPosition(node.Map),
		Key:   m.MarshalExpr(node.Key),
		Value: m.MarshalExpr(node.Value),
	}
}

var ChanDirToString = map[ast.ChanDir]string{
	ast.SEND: "SEND",
	ast.RECV: "RECV",
}

func (m *Marshaller) MarshalChanType(node *ast.ChanType) *ChanTypeNode {
	return &ChanTypeNode{
		Node:  m.MarshalNode("ChanType"),
		Begin: m.MarshalPosition(node.Begin),
		Arrow: m.MarshalPosition(node.Arrow),
		Dir:   ChanDirToString[node.Dir],
		Value: m.MarshalExpr(node.Value),
	}
}

// ---------------------------------------------------------------------------

func (m *Marshaller) MarshalBadStmt(stmt *ast.BadStmt) *BadStmtNode {
	return &BadStmtNode{
		Node: m.MarshalNode("BadStmt"),
		From: m.MarshalPosition(stmt.From),
		To:   m.MarshalPosition(stmt.To),
	}
}

func (m *Marshaller) MarshalDeclStmt(stmt *ast.DeclStmt) *DeclStmtNode {
	return &DeclStmtNode{
		Node: m.MarshalNode("DeclStmt"),
		Decl: m.MarshalDecl(stmt.Decl),
	}
}

func (m *Marshaller) MarshalEmptyStmt(stmt *ast.EmptyStmt) *EmptyStmtNode {
	return &EmptyStmtNode{
		Node:      m.MarshalNode("EmptyStmt"),
		Semicolon: m.MarshalPosition(stmt.Semicolon),
		Implicit:  stmt.Implicit,
	}
}

func (m *Marshaller) MarshalLabeledStmt(stmt *ast.LabeledStmt) *LabeledStmtNode {
	return &LabeledStmtNode{
		Node:  m.MarshalNode("LabeledStmt"),
		Label: m.MarshalIdent(stmt.Label),
		Colon: m.MarshalPosition(stmt.Colon),
		Stmt:  m.MarshalStmt(stmt.Stmt),
	}
}

func (m *Marshaller) MarshalExprStmt(stmt *ast.ExprStmt) *ExprStmtNode {
	return &ExprStmtNode{
		Node: m.MarshalNode("ExprStmt"),
		X:    m.MarshalExpr(stmt.X),
	}
}

func (m *Marshaller) MarshalSendStmt(stmt *ast.SendStmt) *SendStmtNode {
	return &SendStmtNode{
		Node:  m.MarshalNode("SendStmt"),
		Chan:  m.MarshalExpr(stmt.Chan),
		Arrow: m.MarshalPosition(stmt.Arrow),
		Value: m.MarshalExpr(stmt.Value),
	}
}

func (m *Marshaller) MarshalIncDecStmt(stmt *ast.IncDecStmt) *IncDecStmtNode {
	return &IncDecStmtNode{
		Node:   m.MarshalNode("IncDecStmt"),
		X:      m.MarshalExpr(stmt.X),
		TokPos: m.MarshalPosition(stmt.TokPos),
		Tok:    stmt.Tok.String(),
	}
}

func (m *Marshaller) MarshalAssignStmt(stmt *ast.AssignStmt) *AssignStmtNode {
	return &AssignStmtNode{
		Node:   m.MarshalNode("AssignStmt"),
		Lhs:    m.MarshalExprs(stmt.Lhs),
		TokPos: m.MarshalPosition(stmt.TokPos),
		Tok:    stmt.Tok.String(),
		Rhs:    m.MarshalExprs(stmt.Rhs),
	}
}

func (m *Marshaller) MarshalGoStmt(stmt *ast.GoStmt) *GoStmtNode {
	return &GoStmtNode{
		Node: m.MarshalNode("GoStmt"),
		Go:   m.MarshalPosition(stmt.Go),
		Call: m.MarshalCallExpr(stmt.Call),
	}
}

func (m *Marshaller) MarshalDeferStmt(stmt *ast.DeferStmt) *DeferStmtNode {
	return &DeferStmtNode{
		Node:  m.MarshalNode("DeferStmt"),
		Defer: m.MarshalPosition(stmt.Defer),
		Call:  m.MarshalCallExpr(stmt.Call),
	}
}

func (m *Marshaller) MarshalReturnStmt(stmt *ast.ReturnStmt) *ReturnStmtNode {
	return &ReturnStmtNode{
		Node:    m.MarshalNode("ReturnStmt"),
		Return:  m.MarshalPosition(stmt.Return),
		Results: m.MarshalExprs(stmt.Results),
	}
}

func (m *Marshaller) MarshalBranchStmt(stmt *ast.BranchStmt) *BranchStmtNode {
	return &BranchStmtNode{
		Node:   m.MarshalNode("BranchStmt"),
		TokPos: m.MarshalPosition(stmt.TokPos),
		Tok:    stmt.Tok.String(),
		Label:  m.MarshalIdent(stmt.Label),
	}
}

func (m *Marshaller) MarshalBlockStmt(stmt *ast.BlockStmt) *BlockStmtNode {
	return &BlockStmtNode{
		Node:   m.MarshalNode("BlockStmt"),
		Lbrace: m.MarshalPosition(stmt.Lbrace),
		List:   m.MarshalStmts(stmt.List),
		Rbrace: m.MarshalPosition(stmt.Rbrace),
	}
}

func (m *Marshaller) MarshalIfStmt(stmt *ast.IfStmt) *IfStmtNode {
	return &IfStmtNode{
		Node: m.MarshalNode("IfStmt"),
		If:   m.MarshalPosition(stmt.If),
		Init: m.MarshalStmt(stmt.Init),
		Cond: m.MarshalExpr(stmt.Cond),
		Body: m.MarshalBlockStmt(stmt.Body),
		Else: m.MarshalStmt(stmt.Else),
	}
}

func (m *Marshaller) MarshalCaseClause(stmt *ast.CaseClause) *CaseClauseNode {
	return &CaseClauseNode{
		Node:  m.MarshalNode("CaseClause"),
		Case:  m.MarshalPosition(stmt.Case),
		List:  m.MarshalExprs(stmt.List),
		Colon: m.MarshalPosition(stmt.Colon),
		Body:  m.MarshalStmts(stmt.Body),
	}
}

func (m *Marshaller) MarshalSwitchStmt(stmt *ast.SwitchStmt) *SwitchStmtNode {
	return &SwitchStmtNode{
		Node:   m.MarshalNode("SwitchStmt"),
		Switch: m.MarshalPosition(stmt.Switch),
		Init:   m.MarshalStmt(stmt.Init),
		Tag:    m.MarshalExpr(stmt.Tag),
		Body:   m.MarshalBlockStmt(stmt.Body),
	}
}

func (m *Marshaller) MarshalTypeSwitchStmt(stmt *ast.TypeSwitchStmt) *TypeSwitchStmtNode {
	return &TypeSwitchStmtNode{
		Node:   m.MarshalNode("TypeSwitchStmt"),
		Switch: m.MarshalPosition(stmt.Switch),
		Init:   m.MarshalStmt(stmt.Init),
		Assign: m.MarshalStmt(stmt.Assign),
		Body:   m.MarshalBlockStmt(stmt.Body),
	}
}

func (m *Marshaller) MarshalCommClause(stmt *ast.CommClause) *CommClauseNode {
	return &CommClauseNode{
		Node: m.MarshalNode("CommClause"),
		Case: m.MarshalPosition(stmt.Case),
		Comm: m.MarshalStmt(stmt.Comm),
		Body: m.MarshalStmts(stmt.Body),
	}
}

func (m *Marshaller) MarshalSelectStmt(stmt *ast.SelectStmt) *SelectStmtNode {
	return &SelectStmtNode{
		Node:   m.MarshalNode("SelectStmt"),
		Select: m.MarshalPosition(stmt.Select),
		Body:   m.MarshalBlockStmt(stmt.Body),
	}
}

func (m *Marshaller) MarshalForStmt(stmt *ast.ForStmt) *ForStmtNode {
	return &ForStmtNode{
		Node: m.MarshalNode("ForStmt"),
		For:  m.MarshalPosition(stmt.For),
		Init: m.MarshalStmt(stmt.Init),
		Cond: m.MarshalExpr(stmt.Cond),
		Post: m.MarshalStmt(stmt.Post),
		Body: m.MarshalBlockStmt(stmt.Body),
	}
}

func (m *Marshaller) MarshalRangeStmt(stmt *ast.RangeStmt) *RangeStmtNode {
	return &RangeStmtNode{
		Node:   m.MarshalNode("RangeStmt"),
		For:    m.MarshalPosition(stmt.For),
		Key:    m.MarshalExpr(stmt.Key),
		Value:  m.MarshalExpr(stmt.Value),
		TokPos: m.MarshalPosition(stmt.TokPos),
		Tok:    stmt.Tok.String(),
		X:      m.MarshalExpr(stmt.X),
		Body:   m.MarshalBlockStmt(stmt.Body),
	}
}

func (m *Marshaller) MarshalStmt(node ast.Stmt) IStmtNode {
	if node == nil {
		return nil
	}
	switch stmt := node.(type) {
	case *ast.BadStmt:
		return m.MarshalBadStmt(stmt)
	case *ast.DeclStmt:
		return m.MarshalDeclStmt(stmt)
	case *ast.EmptyStmt:
		return m.MarshalEmptyStmt(stmt)
	case *ast.LabeledStmt:
		return m.MarshalLabeledStmt(stmt)
	case *ast.ExprStmt:
		return m.MarshalExprStmt(stmt)
	case *ast.SendStmt:
		return m.MarshalSendStmt(stmt)
	case *ast.IncDecStmt:
		return m.MarshalIncDecStmt(stmt)
	case *ast.AssignStmt:
		return m.MarshalAssignStmt(stmt)
	case *ast.GoStmt:
		return m.MarshalGoStmt(stmt)
	case *ast.DeferStmt:
		return m.MarshalDeferStmt(stmt)
	case *ast.ReturnStmt:
		return m.MarshalReturnStmt(stmt)
	case *ast.BranchStmt:
		return m.MarshalBranchStmt(stmt)
	case *ast.BlockStmt:
		return m.MarshalBlockStmt(stmt)
	case *ast.IfStmt:
		return m.MarshalIfStmt(stmt)
	case *ast.CaseClause:
		return m.MarshalCaseClause(stmt)
	case *ast.SwitchStmt:
		return m.MarshalSwitchStmt(stmt)
	case *ast.TypeSwitchStmt:
		return m.MarshalTypeSwitchStmt(stmt)
	case *ast.CommClause:
		return m.MarshalCommClause(stmt)
	case *ast.SelectStmt:
		return m.MarshalSelectStmt(stmt)
	case *ast.ForStmt:
		return m.MarshalForStmt(stmt)
	case *ast.RangeStmt:
		return m.MarshalRangeStmt(stmt)
	default:
		panic("implement me")
	}
}

func (m *Marshaller) MarshalStmts(stmts []ast.Stmt) []IStmtNode {
	if stmts == nil {
		return nil
	}
	nodes := make([]IStmtNode, len(stmts))
	for index, stmt := range stmts {
		nodes[index] = m.MarshalStmt(stmt)
	}
	return nodes
}

// ---------------------------------------------------------------------------

func (m *Marshaller) MarshalImportSpec(spec *ast.ImportSpec) *ImportSpecNode {
	return &ImportSpecNode{
		Node:    m.MarshalNode("ImportSpec"),
		Doc:     m.MarshalCommentGroup(spec.Doc),
		Name:    m.MarshalIdent(spec.Name),
		Path:    m.MarshalBasicLit(spec.Path),
		Comment: m.MarshalCommentGroup(spec.Comment),
		EndPos:  m.MarshalPosition(spec.EndPos),
	}
}

func (m *Marshaller) MarshalValueSpec(spec *ast.ValueSpec) *ValueSpecNode {
	return &ValueSpecNode{
		Node:    m.MarshalNode("ValueSpec"),
		Doc:     m.MarshalCommentGroup(spec.Doc),
		Names:   m.MarshalIdents(spec.Names),
		Type:    m.MarshalExpr(spec.Type),
		Values:  m.MarshalExprs(spec.Values),
		Comment: m.MarshalCommentGroup(spec.Comment),
	}
}

func (m *Marshaller) MarshalTypeSpec(spec *ast.TypeSpec) *TypeSpecNode {
	return &TypeSpecNode{
		Node:       m.MarshalNode("TypeSpec"),
		Doc:        m.MarshalCommentGroup(spec.Doc),
		Name:       m.MarshalIdent(spec.Name),
		TypeParams: m.MarshalFieldList(spec.TypeParams),
		Assign:     m.MarshalPosition(spec.Assign),
		Type:       m.MarshalExpr(spec.Type),
		Comment:    m.MarshalCommentGroup(spec.Comment),
	}
}

func (m *Marshaller) MarshalSpec(node ast.Spec) ISpecNode {
	switch spec := node.(type) {
	case *ast.ImportSpec:
		return m.MarshalImportSpec(spec)
	case *ast.ValueSpec:
		return m.MarshalValueSpec(spec)
	case *ast.TypeSpec:
		return m.MarshalTypeSpec(spec)
	default:
		panic("implement me")
	}
}

func (m *Marshaller) MarshalSpecs(specs []ast.Spec) []ISpecNode {
	if specs == nil {
		return nil
	}
	nodes := make([]ISpecNode, len(specs))
	for index, spec := range specs {
		nodes[index] = m.MarshalSpec(spec)
	}
	return nodes
}

// ---------------------------------------------------------------------------

func (m *Marshaller) MarshalBadDecl(decl *ast.BadDecl) *BadDeclNode {
	return &BadDeclNode{
		Node: m.MarshalNode("BadDecl"),
		From: m.MarshalPosition(decl.From),
		To:   m.MarshalPosition(decl.To),
	}
}

func (m *Marshaller) MarshalGenDecl(node *ast.GenDecl) *GenDeclNode {
	return &GenDeclNode{
		Node:   m.MarshalNode("GenDecl"),
		Doc:    m.MarshalCommentGroup(node.Doc),
		TokPos: m.MarshalPosition(node.TokPos),
		Tok:    node.Tok.String(),
		Lparen: m.MarshalPosition(node.Lparen),
		Specs:  m.MarshalSpecs(node.Specs),
		Rparen: m.MarshalPosition(node.Rparen),
	}
}

func (m *Marshaller) MarshalFuncDecl(node *ast.FuncDecl) *FuncDeclNode {
	return &FuncDeclNode{
		Node: m.MarshalNode("FuncDecl"),
		Doc:  m.MarshalCommentGroup(node.Doc),
		Recv: m.MarshalFieldList(node.Recv),
		Name: m.MarshalIdent(node.Name),
		Type: m.MarshalFuncType(node.Type),
		Body: m.MarshalBlockStmt(node.Body),
	}
}

func (m *Marshaller) MarshalDecl(node ast.Decl) IDeclNode {
	switch decl := node.(type) {
	case *ast.BadDecl:
		return m.MarshalBadDecl(decl)
	case *ast.GenDecl:
		return m.MarshalGenDecl(decl)
	case *ast.FuncDecl:
		return m.MarshalFuncDecl(decl)
	default:
		panic("implement me")
	}
}

func (m *Marshaller) MarshalDecls(decls []ast.Decl) []IDeclNode {
	if decls == nil {
		return nil
	}
	nodes := make([]IDeclNode, len(decls))
	for index, decl := range decls {
		nodes[index] = m.MarshalDecl(decl)
	}
	return nodes
}

// ---------------------------------------------------------------------------

func (m *Marshaller) MarshalFile(node *ast.File) *FileNode {
	return &FileNode{
		Node:    m.MarshalNode("File"),
		Doc:     m.MarshalCommentGroup(node.Doc),
		Package: m.MarshalPosition(node.Package),
		Name:    m.MarshalIdent(node.Name),
		Decls:   m.MarshalDecls(node.Decls),
	}
}
