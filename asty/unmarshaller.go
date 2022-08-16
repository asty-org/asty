package asty

import (
	"go/ast"
	"go/token"
)

var StringToToken = map[string]token.Token{}

func init() {
	for t := token.ILLEGAL; t <= token.TILDE; t++ {
		StringToToken[t.String()] = t
	}
}

// TODO: implement positions and comments unmarshalling

type Unmarshaller struct {
}

func NewUnmarshaller() *Unmarshaller {
	return &Unmarshaller{}
}

func (um *Unmarshaller) UnmarshalFieldNode(node *FieldNode) *ast.Field {
	return &ast.Field{
		Names: um.UnmarshalIdentNodes(node.Names),
		Type:  um.UnmarshalExpr(node.Type),
		Tag:   um.UnmarshalBasicLitNode(node.Tag),
	}
}

func (um *Unmarshaller) UnmarshalFieldNodes(nodes []*FieldNode) []*ast.Field {
	if nodes == nil {
		return nil
	}
	fields := make([]*ast.Field, len(nodes))
	for index, node := range nodes {
		fields[index] = um.UnmarshalFieldNode(node)
	}
	return fields
}

func (um *Unmarshaller) UnmarshalFieldListNode(node *FieldListNode) *ast.FieldList {
	if node == nil {
		return nil
	}
	return &ast.FieldList{
		List: um.UnmarshalFieldNodes(node.List),
	}
}

func (um *Unmarshaller) UnmarshalBadExprNode(_ *BadExprNode) *ast.BadExpr {
	return &ast.BadExpr{}
}

func (um *Unmarshaller) UnmarshalIdentNode(node *IdentNode) *ast.Ident {
	if node == nil {
		return nil
	}
	return &ast.Ident{
		Name: node.Name,
	}
}

func (um *Unmarshaller) UnmarshalIdentNodes(nodes []*IdentNode) []*ast.Ident {
	if nodes == nil {
		return nil
	}
	idents := make([]*ast.Ident, len(nodes))
	for index, node := range nodes {
		idents[index] = um.UnmarshalIdentNode(node)
	}
	return idents
}

func (um *Unmarshaller) UnmarshalEllipsisNode(node *EllipsisNode) *ast.Ellipsis {
	return &ast.Ellipsis{
		Elt: um.UnmarshalExpr(node.Elt),
	}
}
func (um *Unmarshaller) UnmarshalBasicLitNode(node *BasicLitNode) *ast.BasicLit {
	if node == nil {
		return nil
	}
	kind, ok := StringToToken[node.Kind]
	if !ok {
		panic("unsupported token kind " + node.Kind)
	}
	return &ast.BasicLit{
		Kind:  kind,
		Value: node.Value,
	}
}

func (um *Unmarshaller) UnmarshalFuncLitNode(node *FuncLitNode) *ast.FuncLit {
	return &ast.FuncLit{
		Type: um.UnmarshalFuncTypeNode(node.Type),
		Body: um.UnmarshalBlockStmtNode(node.Body),
	}
}

func (um *Unmarshaller) UnmarshalCompositeLitNode(node *CompositeLitNode) *ast.CompositeLit {
	return &ast.CompositeLit{
		Type: um.UnmarshalExpr(node.Type),
	}
}

func (um *Unmarshaller) UnmarshalParenExprNode(node *ParenExprNode) *ast.ParenExpr {
	return &ast.ParenExpr{
		X: um.UnmarshalExpr(node.X),
	}
}

func (um *Unmarshaller) UnmarshalSelectorExprNode(node *SelectorExprNode) *ast.SelectorExpr {
	return &ast.SelectorExpr{
		X:   um.UnmarshalExpr(node.X),
		Sel: um.UnmarshalIdentNode(node.Sel),
	}
}

func (um *Unmarshaller) UnmarshalIndexExprNode(node *IndexExprNode) *ast.IndexExpr {
	return &ast.IndexExpr{
		X:     um.UnmarshalExpr(node.X),
		Index: um.UnmarshalExpr(node.Index),
	}
}

func (um *Unmarshaller) UnmarshalExprNodes(nodes []IExprNode) []ast.Expr {
	if nodes == nil {
		return nil
	}
	exprs := make([]ast.Expr, len(nodes))
	for index, node := range nodes {
		exprs[index] = um.UnmarshalExpr(node)
	}
	return exprs
}

func (um *Unmarshaller) UnmarshalIndexListExprNode(node *IndexListExprNode) *ast.IndexListExpr {
	return &ast.IndexListExpr{
		X:       um.UnmarshalExpr(node.X),
		Indices: um.UnmarshalExprNodes(node.Indices),
	}
}

func (um *Unmarshaller) UnmarshalSliceExprNode(node *SliceExprNode) *ast.SliceExpr {
	return &ast.SliceExpr{
		X:      um.UnmarshalExpr(node.X),
		Low:    um.UnmarshalExpr(node.Low),
		High:   um.UnmarshalExpr(node.High),
		Max:    um.UnmarshalExpr(node.Max),
		Slice3: node.Slice3,
	}
}

func (um *Unmarshaller) UnmarshalTypeAssertExprNode(node *TypeAssertExprNode) *ast.TypeAssertExpr {
	return &ast.TypeAssertExpr{
		X:    um.UnmarshalExpr(node.X),
		Type: um.UnmarshalExpr(node.Type),
	}
}

func (um *Unmarshaller) UnmarshalCallExprNode(node *CallExprNode) *ast.CallExpr {
	return &ast.CallExpr{
		Fun:  um.UnmarshalExpr(node.Fun),
		Args: um.UnmarshalExprNodes(node.Args),
	}
}

func (um *Unmarshaller) UnmarshalStarExprNode(node *StarExprNode) *ast.StarExpr {
	return &ast.StarExpr{
		X: um.UnmarshalExpr(node.X),
	}
}

func (um *Unmarshaller) UnmarshalUnaryExprNode(node *UnaryExprNode) *ast.UnaryExpr {
	return &ast.UnaryExpr{
		Op: StringToToken[node.Op],
		X:  um.UnmarshalExpr(node.X),
	}
}

func (um *Unmarshaller) UnmarshalBinaryExprNode(node *BinaryExprNode) *ast.BinaryExpr {
	return &ast.BinaryExpr{
		X:  um.UnmarshalExpr(node.X),
		Op: StringToToken[node.Op],
		Y:  um.UnmarshalExpr(node.Y),
	}
}

func (um *Unmarshaller) UnmarshalKeyValueExprNode(node *KeyValueExprNode) *ast.KeyValueExpr {
	return &ast.KeyValueExpr{
		Key:   um.UnmarshalExpr(node.Key),
		Value: um.UnmarshalExpr(node.Value),
	}
}

func (um *Unmarshaller) UnmarshalArrayTypeNode(node *ArrayTypeNode) *ast.ArrayType {
	return &ast.ArrayType{
		Len: um.UnmarshalExpr(node.Len),
		Elt: um.UnmarshalExpr(node.Elt),
	}
}

func (um *Unmarshaller) UnmarshalStructTypeNode(node *StructTypeNode) *ast.StructType {
	return &ast.StructType{
		Fields:     um.UnmarshalFieldListNode(node.Fields),
		Incomplete: node.Incomplete,
	}
}

func (um *Unmarshaller) UnmarshalFuncTypeNode(node *FuncTypeNode) *ast.FuncType {
	return &ast.FuncType{
		TypeParams: um.UnmarshalFieldListNode(node.TypeParams),
		Params:     um.UnmarshalFieldListNode(node.Params),
		Results:    um.UnmarshalFieldListNode(node.Results),
	}
}

func (um *Unmarshaller) UnmarshalInterfaceTypeNode(node *InterfaceTypeNode) *ast.InterfaceType {
	return &ast.InterfaceType{
		Methods:    um.UnmarshalFieldListNode(node.Methods),
		Incomplete: node.Incomplete,
	}
}

func (um *Unmarshaller) UnmarshalMapTypeNode(node *MapTypeNode) *ast.MapType {
	return &ast.MapType{
		Key:   um.UnmarshalExpr(node.Key),
		Value: um.UnmarshalExpr(node.Value),
	}
}

var StringToChanDir = map[string]ast.ChanDir{
	"SEND": ast.SEND,
	"RECV": ast.RECV,
}

func (um *Unmarshaller) UnmarshalChanTypeNode(node *ChanTypeNode) *ast.ChanType {
	return &ast.ChanType{
		Dir:   StringToChanDir[node.Dir],
		Value: um.UnmarshalExpr(node.Value),
	}
}

func (um *Unmarshaller) UnmarshalBadStmtNode(_ *BadStmtNode) *ast.BadStmt {
	return &ast.BadStmt{}
}

func (um *Unmarshaller) UnmarshalDeclStmtNode(node *DeclStmtNode) *ast.DeclStmt {
	return &ast.DeclStmt{
		Decl: um.UnmarshalDecl(node.Decl),
	}
}

func (um *Unmarshaller) UnmarshalEmptyStmtNode(_ *EmptyStmtNode) *ast.EmptyStmt {
	return &ast.EmptyStmt{}
}

func (um *Unmarshaller) UnmarshalLabeledStmtNode(node *LabeledStmtNode) *ast.LabeledStmt {
	return &ast.LabeledStmt{
		Label: um.UnmarshalIdentNode(node.Label),
		Stmt:  um.UnmarshalStmt(node.Stmt),
	}
}

func (um *Unmarshaller) UnmarshalExprStmtNode(node *ExprStmtNode) *ast.ExprStmt {
	return &ast.ExprStmt{
		X: um.UnmarshalExpr(node.X),
	}
}

func (um *Unmarshaller) UnmarshalSendStmtNode(node *SendStmtNode) *ast.SendStmt {
	return &ast.SendStmt{
		Chan:  um.UnmarshalExpr(node.Chan),
		Value: um.UnmarshalExpr(node.Value),
	}
}

func (um *Unmarshaller) UnmarshalIncDecStmtNode(node *IncDecStmtNode) *ast.IncDecStmt {
	return &ast.IncDecStmt{
		X:   um.UnmarshalExpr(node.X),
		Tok: StringToToken[node.Tok],
	}
}

func (um *Unmarshaller) UnmarshalAssignStmtNode(node *AssignStmtNode) *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: um.UnmarshalExprNodes(node.Lhs),
		Tok: StringToToken[node.Tok],
		Rhs: um.UnmarshalExprNodes(node.Rhs),
	}
}

func (um *Unmarshaller) UnmarshalGoStmtNode(node *GoStmtNode) *ast.GoStmt {
	return &ast.GoStmt{
		Call: um.UnmarshalCallExprNode(node.Call),
	}
}

func (um *Unmarshaller) UnmarshalDeferStmtNode(node *DeferStmtNode) *ast.DeferStmt {
	return &ast.DeferStmt{
		Call: um.UnmarshalCallExprNode(node.Call),
	}
}

func (um *Unmarshaller) UnmarshalReturnStmtNode(node *ReturnStmtNode) *ast.ReturnStmt {
	return &ast.ReturnStmt{
		Results: um.UnmarshalExprNodes(node.Results),
	}
}

func (um *Unmarshaller) UnmarshalBranchStmtNode(node *BranchStmtNode) *ast.BranchStmt {
	return &ast.BranchStmt{
		Tok:   StringToToken[node.Tok],
		Label: um.UnmarshalIdentNode(node.Label),
	}
}

func (um *Unmarshaller) UnmarshalStmtNodes(nodes []IStmtNode) []ast.Stmt {
	if nodes == nil {
		return nil
	}
	stmts := make([]ast.Stmt, len(nodes))
	for i, node := range nodes {
		stmts[i] = um.UnmarshalStmt(node)
	}
	return stmts
}

func (um *Unmarshaller) UnmarshalBlockStmtNode(node *BlockStmtNode) *ast.BlockStmt {
	return &ast.BlockStmt{
		List: um.UnmarshalStmtNodes(node.List),
	}
}

func (um *Unmarshaller) UnmarshalIfStmtNode(node *IfStmtNode) *ast.IfStmt {
	return &ast.IfStmt{
		Init: um.UnmarshalStmt(node.Init),
		Cond: um.UnmarshalExpr(node.Cond),
		Body: um.UnmarshalBlockStmtNode(node.Body),
		Else: um.UnmarshalStmt(node.Else),
	}
}

func (um *Unmarshaller) UnmarshalCaseClauseNode(node *CaseClauseNode) *ast.CaseClause {
	return &ast.CaseClause{
		List: um.UnmarshalExprNodes(node.List),
		Body: um.UnmarshalStmtNodes(node.Body),
	}
}

func (um *Unmarshaller) UnmarshalSwitchStmtNode(node *SwitchStmtNode) *ast.SwitchStmt {
	return &ast.SwitchStmt{
		Init: um.UnmarshalStmt(node.Init),
		Tag:  um.UnmarshalExpr(node.Tag),
		Body: um.UnmarshalBlockStmtNode(node.Body),
	}
}

func (um *Unmarshaller) UnmarshalTypeSwitchStmtNode(node *TypeSwitchStmtNode) *ast.TypeSwitchStmt {
	return &ast.TypeSwitchStmt{
		Init:   um.UnmarshalStmt(node.Init),
		Assign: um.UnmarshalStmt(node.Assign),
		Body:   um.UnmarshalBlockStmtNode(node.Body),
	}
}

func (um *Unmarshaller) UnmarshalCommClauseNode(node *CommClauseNode) *ast.CommClause {
	return &ast.CommClause{
		Comm: um.UnmarshalStmt(node.Comm),
		Body: um.UnmarshalStmtNodes(node.Body),
	}
}

func (um *Unmarshaller) UnmarshalSelectStmtNode(node *SelectStmtNode) *ast.SelectStmt {
	return &ast.SelectStmt{
		Body: um.UnmarshalBlockStmtNode(node.Body),
	}
}

func (um *Unmarshaller) UnmarshalForStmtNode(node *ForStmtNode) *ast.ForStmt {
	return &ast.ForStmt{
		Init: um.UnmarshalStmt(node.Init),
		Cond: um.UnmarshalExpr(node.Cond),
		Post: um.UnmarshalStmt(node.Post),
		Body: um.UnmarshalBlockStmtNode(node.Body),
	}
}

func (um *Unmarshaller) UnmarshalRangeStmtNode(node *RangeStmtNode) *ast.RangeStmt {
	return &ast.RangeStmt{
		Key:   um.UnmarshalExpr(node.Key),
		Value: um.UnmarshalExpr(node.Value),
		Tok:   StringToToken[node.Tok],
		X:     um.UnmarshalExpr(node.X),
		Body:  um.UnmarshalBlockStmtNode(node.Body),
	}
}

func (um *Unmarshaller) UnmarshalImportSpecNode(node *ImportSpecNode) *ast.ImportSpec {
	return &ast.ImportSpec{
		Name: um.UnmarshalIdentNode(node.Name),
		Path: um.UnmarshalBasicLitNode(node.Path),
	}
}

func (um *Unmarshaller) UnmarshalValueSpecNode(node *ValueSpecNode) *ast.ValueSpec {
	return &ast.ValueSpec{
		Names:  um.UnmarshalIdentNodes(node.Names),
		Type:   um.UnmarshalExpr(node.Type),
		Values: um.UnmarshalExprNodes(node.Values),
	}
}

func (um *Unmarshaller) UnmarshalTypeSpecNode(node *TypeSpecNode) *ast.TypeSpec {
	return &ast.TypeSpec{
		Name:       um.UnmarshalIdentNode(node.Name),
		TypeParams: um.UnmarshalFieldListNode(node.TypeParams),
		Type:       um.UnmarshalExpr(node.Type),
	}
}

func (um *Unmarshaller) UnmarshalSpecNodes(nodes []ISpecNode) []ast.Spec {
	if nodes == nil {
		return nil
	}
	specs := make([]ast.Spec, len(nodes))
	for i, node := range nodes {
		specs[i] = um.UnmarshalSpec(node)
	}
	return specs
}

func (um *Unmarshaller) UnmarshalBadDeclNode(_ *BadDeclNode) *ast.BadDecl {
	return &ast.BadDecl{}
}

func (um *Unmarshaller) UnmarshalGenDeclNode(node *GenDeclNode) *ast.GenDecl {

	return &ast.GenDecl{
		Tok:   StringToToken[node.Tok],
		Specs: um.UnmarshalSpecNodes(node.Specs),
	}
}

func (um *Unmarshaller) UnmarshalFuncDeclNode(node *FuncDeclNode) *ast.FuncDecl {
	return &ast.FuncDecl{
		Recv: um.UnmarshalFieldListNode(node.Recv),
		Name: um.UnmarshalIdentNode(node.Name),
		Type: um.UnmarshalFuncTypeNode(node.Type),
		Body: um.UnmarshalBlockStmtNode(node.Body),
	}
}

func (um *Unmarshaller) UnmarshalDeclNodes(nodes []IDeclNode) []ast.Decl {
	if nodes == nil {
		return nil
	}
	decls := make([]ast.Decl, len(nodes))
	for i, node := range nodes {
		decls[i] = um.UnmarshalDecl(node)
	}
	return decls
}

func (um *Unmarshaller) UnmarshalFileNode(node *FileNode) *ast.File {
	return &ast.File{
		Name:  um.UnmarshalIdentNode(node.Name),
		Decls: um.UnmarshalDeclNodes(node.Decls),
	}
}

func (um *Unmarshaller) UnmarshalExpr(expr IExprNode) ast.Expr {
	if expr == nil {
		return nil
	}
	return expr.UnmarshalExpr(um)
}

func (um *Unmarshaller) UnmarshalStmt(stmt IStmtNode) ast.Stmt {
	if stmt == nil {
		return nil
	}
	return stmt.UnmarshalStmt(um)
}

func (um *Unmarshaller) UnmarshalSpec(spec ISpecNode) ast.Spec {
	if spec == nil {
		return nil
	}
	return spec.UnmarshalSpec(um)
}

func (um *Unmarshaller) UnmarshalDecl(decl IDeclNode) ast.Decl {
	if decl == nil {
		return nil
	}
	return decl.UnmarshalDecl(um)
}
