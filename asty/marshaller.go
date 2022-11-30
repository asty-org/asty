package asty

import (
	"go/ast"
	"go/token"
	"reflect"
)

type Marshaller struct {
	WithPositions  bool
	WithComments   bool
	WithReferences bool

	fset       *token.FileSet
	references map[any]any
	refcount   int
}

func NewMarshaller(withComments, withPositions, withReferences bool) *Marshaller {
	return &Marshaller{
		WithComments:   withComments,
		WithPositions:  withPositions,
		WithReferences: withReferences,
		fset:           token.NewFileSet(),
		references:     make(map[any]any),
		refcount:       0,
	}
}

func (m *Marshaller) FileSet() *token.FileSet {
	return m.fset
}

func wrapMarshal[T any, R any](m *Marshaller, node *T, marshal func() *R) *R {
	if node == nil {
		return nil
	}

	if !m.WithReferences {
		return marshal()
	}

	if ref, ok := m.references[node]; ok {
		return ref.(*R)
	}
	result := marshal()
	m.references[node] = result
	return result
}

// ---------------------------------------------------------------------------

func (m *Marshaller) MarshalNode(nodeType string, _ ast.Node) Node {
	ref := 0
	if m.WithReferences {
		m.refcount++
		ref = m.refcount
	}
	return Node{
		NodeType: nodeType,
		RefId:    ref,
	}
}

func (m *Marshaller) MarshalPosition(pos token.Pos) *PositionNode {
	if !m.WithPositions {
		return nil
	}
	if pos == token.NoPos {
		return nil
	}
	position := m.fset.PositionFor(pos, false)
	return &PositionNode{
		Node:     m.MarshalNode("Position", nil),
		Filename: position.Filename,
		Line:     position.Line,
		Offset:   position.Offset,
		Column:   position.Column,
	}
}

func (m *Marshaller) MarshalComment(comment *ast.Comment) *CommentNode {
	return wrapMarshal(m, comment, func() *CommentNode {
		return &CommentNode{
			Node:  m.MarshalNode("Comment", comment),
			Slash: m.MarshalPosition(comment.Slash),
			Text:  comment.Text,
		}
	})
}

func (m *Marshaller) MarshalComments(comments []*ast.Comment) []*CommentNode {
	if comments == nil {
		return nil
	}
	nodes := make([]*CommentNode, len(comments))
	for index, comment := range comments {
		nodes[index] = m.MarshalComment(comment)
	}
	return nodes
}

func (m *Marshaller) MarshalCommentGroup(group *ast.CommentGroup) *CommentGroupNode {
	if !m.WithComments {
		return nil
	}
	return wrapMarshal(m, group, func() *CommentGroupNode {
		return &CommentGroupNode{
			Node: m.MarshalNode("CommentGroup", group),
			List: m.MarshalComments(group.List),
		}
	})
}

func (m *Marshaller) MarshalCommentGroups(groups []*ast.CommentGroup) []*CommentGroupNode {
	if groups == nil {
		return nil
	}
	nodes := make([]*CommentGroupNode, len(groups))
	for index, comment := range groups {
		nodes[index] = m.MarshalCommentGroup(comment)
	}
	return nodes
}

// ---------------------------------------------------------------------------

func (m *Marshaller) MarshalField(node *ast.Field) *FieldNode {
	return wrapMarshal(m, node, func() *FieldNode {
		return &FieldNode{
			Node:    m.MarshalNode("Field", node),
			Doc:     m.MarshalCommentGroup(node.Doc),
			Names:   m.MarshalIdents(node.Names),
			Type:    m.MarshalExpr(node.Type),
			Tag:     m.MarshalBasicLit(node.Tag),
			Comment: m.MarshalCommentGroup(node.Comment),
		}
	})
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
	return wrapMarshal(m, node, func() *FieldListNode {
		return &FieldListNode{
			Node:    m.MarshalNode("FieldList", node),
			Opening: m.MarshalPosition(node.Opening),
			List:    m.MarshalFields(node.List),
			Closing: m.MarshalPosition(node.Closing),
		}
	})
}

// ---------------------------------------------------------------------------

func (m *Marshaller) MarshalBadExpr(node *ast.BadExpr) *BadExprNode {
	return wrapMarshal(m, node, func() *BadExprNode {
		return &BadExprNode{
			Node: m.MarshalNode("BadExpr", node),
			From: m.MarshalPosition(node.From),
			To:   m.MarshalPosition(node.To),
		}
	})
}

func (m *Marshaller) MarshalIdent(node *ast.Ident) *IdentNode {
	return wrapMarshal(m, node, func() *IdentNode {
		return &IdentNode{
			Node:    m.MarshalNode("Ident", node),
			NamePos: m.MarshalPosition(node.NamePos),
			Name:    node.Name,
		}
	})
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
	return wrapMarshal(m, node, func() *EllipsisNode {
		return &EllipsisNode{
			Node:     m.MarshalNode("Ellipsis", node),
			Ellipsis: m.MarshalPosition(node.Ellipsis),
			Elt:      m.MarshalExpr(node.Elt),
		}
	})
}

func (m *Marshaller) MarshalBasicLit(node *ast.BasicLit) *BasicLitNode {
	return wrapMarshal(m, node, func() *BasicLitNode {
		return &BasicLitNode{
			Node:     m.MarshalNode("BasicLit", node),
			ValuePos: m.MarshalPosition(node.ValuePos),
			Kind:     node.Kind.String(),
			Value:    node.Value,
		}
	})
}

func (m *Marshaller) MarshalFuncLit(node *ast.FuncLit) *FuncLitNode {
	return wrapMarshal(m, node, func() *FuncLitNode {
		return &FuncLitNode{
			Node: m.MarshalNode("FuncLit", node),
			Type: m.MarshalFuncType(node.Type),
			Body: m.MarshalBlockStmt(node.Body),
		}
	})
}

func (m *Marshaller) MarshalCompositeLit(node *ast.CompositeLit) *CompositeLitNode {
	return wrapMarshal(m, node, func() *CompositeLitNode {
		return &CompositeLitNode{
			Node:       m.MarshalNode("CompositeLit", node),
			Type:       m.MarshalExpr(node.Type),
			Lbrace:     m.MarshalPosition(node.Lbrace),
			Elts:       m.MarshalExprs(node.Elts),
			Rbrace:     m.MarshalPosition(node.Rbrace),
			Incomplete: node.Incomplete,
		}
	})
}

func (m *Marshaller) MarshalParenExpr(node *ast.ParenExpr) *ParenExprNode {
	return wrapMarshal(m, node, func() *ParenExprNode {
		return &ParenExprNode{
			Node:   m.MarshalNode("ParenExpr", node),
			Lparen: m.MarshalPosition(node.Lparen),
			X:      m.MarshalExpr(node.X),
			Rparen: m.MarshalPosition(node.Rparen),
		}
	})
}

func (m *Marshaller) MarshalSelectorExpr(node *ast.SelectorExpr) *SelectorExprNode {
	return wrapMarshal(m, node, func() *SelectorExprNode {
		return &SelectorExprNode{
			Node: m.MarshalNode("SelectorExpr", node),
			X:    m.MarshalExpr(node.X),
			Sel:  m.MarshalIdent(node.Sel),
		}
	})
}

func (m *Marshaller) MarshalIndexExpr(node *ast.IndexExpr) *IndexExprNode {
	return wrapMarshal(m, node, func() *IndexExprNode {
		return &IndexExprNode{
			Node:   m.MarshalNode("IndexExpr", node),
			X:      m.MarshalExpr(node.X),
			Lbrack: m.MarshalPosition(node.Lbrack),
			Index:  m.MarshalExpr(node.Index),
			Rbrack: m.MarshalPosition(node.Rbrack),
		}
	})
}

func (m *Marshaller) MarshalIndexListExpr(node *ast.IndexListExpr) *IndexListExprNode {
	return wrapMarshal(m, node, func() *IndexListExprNode {
		return &IndexListExprNode{
			Node:    m.MarshalNode("IndexListExpr", node),
			X:       m.MarshalExpr(node.X),
			Lbrack:  m.MarshalPosition(node.Lbrack),
			Indices: m.MarshalExprs(node.Indices),
			Rbrack:  m.MarshalPosition(node.Rbrack),
		}
	})
}

func (m *Marshaller) MarshalSliceExpr(node *ast.SliceExpr) *SliceExprNode {
	return wrapMarshal(m, node, func() *SliceExprNode {
		return &SliceExprNode{
			Node:   m.MarshalNode("SliceExpr", node),
			X:      m.MarshalExpr(node.X),
			Lbrack: m.MarshalPosition(node.Lbrack),
			Low:    m.MarshalExpr(node.Low),
			High:   m.MarshalExpr(node.High),
			Max:    m.MarshalExpr(node.Max),
			Slice3: node.Slice3,
			Rbrack: m.MarshalPosition(node.Rbrack),
		}
	})
}

func (m *Marshaller) MarshalTypeAssertExpr(node *ast.TypeAssertExpr) *TypeAssertExprNode {
	return wrapMarshal(m, node, func() *TypeAssertExprNode {
		return &TypeAssertExprNode{
			Node:   m.MarshalNode("TypeAssertExpr", node),
			X:      m.MarshalExpr(node.X),
			Lparen: m.MarshalPosition(node.Lparen),
			Type:   m.MarshalExpr(node.Type),
			Rparen: m.MarshalPosition(node.Rparen),
		}
	})
}

func (m *Marshaller) MarshalCallExpr(node *ast.CallExpr) *CallExprNode {
	return wrapMarshal(m, node, func() *CallExprNode {
		return &CallExprNode{
			Node:     m.MarshalNode("CallExpr", node),
			Fun:      m.MarshalExpr(node.Fun),
			Lparen:   m.MarshalPosition(node.Lparen),
			Args:     m.MarshalExprs(node.Args),
			Ellipsis: m.MarshalPosition(node.Ellipsis),
			Rparen:   m.MarshalPosition(node.Rparen),
		}
	})
}

func (m *Marshaller) MarshalStarExpr(node *ast.StarExpr) *StarExprNode {
	return wrapMarshal(m, node, func() *StarExprNode {
		return &StarExprNode{
			Node: m.MarshalNode("StarExpr", node),
			Star: m.MarshalPosition(node.Star),
			X:    m.MarshalExpr(node.X),
		}
	})
}

func (m *Marshaller) MarshalUnaryExpr(node *ast.UnaryExpr) *UnaryExprNode {
	return wrapMarshal(m, node, func() *UnaryExprNode {
		return &UnaryExprNode{
			Node:  m.MarshalNode("UnaryExpr", node),
			OpPos: m.MarshalPosition(node.OpPos),
			Op:    node.Op.String(),
			X:     m.MarshalExpr(node.X),
		}
	})
}

func (m *Marshaller) MarshalBinaryExpr(node *ast.BinaryExpr) *BinaryExprNode {
	return wrapMarshal(m, node, func() *BinaryExprNode {
		return &BinaryExprNode{
			Node:  m.MarshalNode("BinaryExpr", node),
			X:     m.MarshalExpr(node.X),
			OpPos: m.MarshalPosition(node.OpPos),
			Op:    node.Op.String(),
			Y:     m.MarshalExpr(node.Y),
		}
	})
}

func (m *Marshaller) MarshalKeyValueExpr(node *ast.KeyValueExpr) *KeyValueExprNode {
	return wrapMarshal(m, node, func() *KeyValueExprNode {
		return &KeyValueExprNode{
			Node:  m.MarshalNode("KeyValueExpr", node),
			Key:   m.MarshalExpr(node.Key),
			Colon: m.MarshalPosition(node.Colon),
			Value: m.MarshalExpr(node.Value),
		}
	})
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
	return wrapMarshal(m, node, func() *ArrayTypeNode {
		return &ArrayTypeNode{
			Node:   m.MarshalNode("ArrayType", node),
			Lbrack: m.MarshalPosition(node.Lbrack),
			Len:    m.MarshalExpr(node.Len),
			Elt:    m.MarshalExpr(node.Elt),
		}
	})
}

func (m *Marshaller) MarshalStructType(node *ast.StructType) *StructTypeNode {
	return wrapMarshal(m, node, func() *StructTypeNode {
		return &StructTypeNode{
			Node:       m.MarshalNode("StructType", node),
			Struct:     m.MarshalPosition(node.Struct),
			Fields:     m.MarshalFieldList(node.Fields),
			Incomplete: node.Incomplete,
		}
	})
}

func (m *Marshaller) MarshalFuncType(node *ast.FuncType) *FuncTypeNode {
	return wrapMarshal(m, node, func() *FuncTypeNode {
		return &FuncTypeNode{
			Node:       m.MarshalNode("FuncType", node),
			Func:       m.MarshalPosition(node.Func),
			TypeParams: m.MarshalFieldList(node.TypeParams),
			Params:     m.MarshalFieldList(node.Params),
			Results:    m.MarshalFieldList(node.Results),
		}
	})
}

func (m *Marshaller) MarshalInterfaceType(node *ast.InterfaceType) *InterfaceTypeNode {
	return wrapMarshal(m, node, func() *InterfaceTypeNode {
		return &InterfaceTypeNode{
			Node:       m.MarshalNode("InterfaceType", node),
			Interface:  m.MarshalPosition(node.Interface),
			Methods:    m.MarshalFieldList(node.Methods),
			Incomplete: node.Incomplete,
		}
	})
}

func (m *Marshaller) MarshalMapType(node *ast.MapType) *MapTypeNode {
	return wrapMarshal(m, node, func() *MapTypeNode {
		return &MapTypeNode{
			Node:  m.MarshalNode("MapType", node),
			Map:   m.MarshalPosition(node.Map),
			Key:   m.MarshalExpr(node.Key),
			Value: m.MarshalExpr(node.Value),
		}
	})
}

var ChanDirToString = map[ast.ChanDir]string{
	ast.SEND:            "SEND",
	ast.RECV:            "RECV",
	ast.SEND | ast.RECV: "BOTH",
}

func (m *Marshaller) MarshalChanType(node *ast.ChanType) *ChanTypeNode {
	return wrapMarshal(m, node, func() *ChanTypeNode {
		return &ChanTypeNode{
			Node:  m.MarshalNode("ChanType", node),
			Begin: m.MarshalPosition(node.Begin),
			Arrow: m.MarshalPosition(node.Arrow),
			Dir:   ChanDirToString[node.Dir],
			Value: m.MarshalExpr(node.Value),
		}
	})
}

// ---------------------------------------------------------------------------

func (m *Marshaller) MarshalBadStmt(stmt *ast.BadStmt) *BadStmtNode {
	return wrapMarshal(m, stmt, func() *BadStmtNode {
		return &BadStmtNode{
			Node: m.MarshalNode("BadStmt", stmt),
			From: m.MarshalPosition(stmt.From),
			To:   m.MarshalPosition(stmt.To),
		}
	})
}

func (m *Marshaller) MarshalDeclStmt(stmt *ast.DeclStmt) *DeclStmtNode {
	return wrapMarshal(m, stmt, func() *DeclStmtNode {
		return &DeclStmtNode{
			Node: m.MarshalNode("DeclStmt", stmt),
			Decl: m.MarshalDecl(stmt.Decl),
		}
	})
}

func (m *Marshaller) MarshalEmptyStmt(stmt *ast.EmptyStmt) *EmptyStmtNode {
	return wrapMarshal(m, stmt, func() *EmptyStmtNode {
		return &EmptyStmtNode{
			Node:      m.MarshalNode("EmptyStmt", stmt),
			Semicolon: m.MarshalPosition(stmt.Semicolon),
			Implicit:  stmt.Implicit,
		}
	})
}

func (m *Marshaller) MarshalLabeledStmt(stmt *ast.LabeledStmt) *LabeledStmtNode {
	return wrapMarshal(m, stmt, func() *LabeledStmtNode {
		return &LabeledStmtNode{
			Node:  m.MarshalNode("LabeledStmt", stmt),
			Label: m.MarshalIdent(stmt.Label),
			Colon: m.MarshalPosition(stmt.Colon),
			Stmt:  m.MarshalStmt(stmt.Stmt),
		}
	})
}

func (m *Marshaller) MarshalExprStmt(stmt *ast.ExprStmt) *ExprStmtNode {
	return wrapMarshal(m, stmt, func() *ExprStmtNode {
		return &ExprStmtNode{
			Node: m.MarshalNode("ExprStmt", stmt),
			X:    m.MarshalExpr(stmt.X),
		}
	})
}

func (m *Marshaller) MarshalSendStmt(stmt *ast.SendStmt) *SendStmtNode {
	return wrapMarshal(m, stmt, func() *SendStmtNode {
		return &SendStmtNode{
			Node:  m.MarshalNode("SendStmt", stmt),
			Chan:  m.MarshalExpr(stmt.Chan),
			Arrow: m.MarshalPosition(stmt.Arrow),
			Value: m.MarshalExpr(stmt.Value),
		}
	})
}

func (m *Marshaller) MarshalIncDecStmt(stmt *ast.IncDecStmt) *IncDecStmtNode {
	return wrapMarshal(m, stmt, func() *IncDecStmtNode {
		return &IncDecStmtNode{
			Node:   m.MarshalNode("IncDecStmt", stmt),
			X:      m.MarshalExpr(stmt.X),
			TokPos: m.MarshalPosition(stmt.TokPos),
			Tok:    stmt.Tok.String(),
		}
	})
}

func (m *Marshaller) MarshalAssignStmt(stmt *ast.AssignStmt) *AssignStmtNode {
	return wrapMarshal(m, stmt, func() *AssignStmtNode {
		return &AssignStmtNode{
			Node:   m.MarshalNode("AssignStmt", stmt),
			Lhs:    m.MarshalExprs(stmt.Lhs),
			TokPos: m.MarshalPosition(stmt.TokPos),
			Tok:    stmt.Tok.String(),
			Rhs:    m.MarshalExprs(stmt.Rhs),
		}
	})
}

func (m *Marshaller) MarshalGoStmt(stmt *ast.GoStmt) *GoStmtNode {
	return wrapMarshal(m, stmt, func() *GoStmtNode {
		return &GoStmtNode{
			Node: m.MarshalNode("GoStmt", stmt),
			Go:   m.MarshalPosition(stmt.Go),
			Call: m.MarshalCallExpr(stmt.Call),
		}
	})
}

func (m *Marshaller) MarshalDeferStmt(stmt *ast.DeferStmt) *DeferStmtNode {
	return wrapMarshal(m, stmt, func() *DeferStmtNode {
		return &DeferStmtNode{
			Node:  m.MarshalNode("DeferStmt", stmt),
			Defer: m.MarshalPosition(stmt.Defer),
			Call:  m.MarshalCallExpr(stmt.Call),
		}
	})
}

func (m *Marshaller) MarshalReturnStmt(stmt *ast.ReturnStmt) *ReturnStmtNode {
	return wrapMarshal(m, stmt, func() *ReturnStmtNode {
		return &ReturnStmtNode{
			Node:    m.MarshalNode("ReturnStmt", stmt),
			Return:  m.MarshalPosition(stmt.Return),
			Results: m.MarshalExprs(stmt.Results),
		}
	})
}

func (m *Marshaller) MarshalBranchStmt(stmt *ast.BranchStmt) *BranchStmtNode {
	return wrapMarshal(m, stmt, func() *BranchStmtNode {
		return &BranchStmtNode{
			Node:   m.MarshalNode("BranchStmt", stmt),
			TokPos: m.MarshalPosition(stmt.TokPos),
			Tok:    stmt.Tok.String(),
			Label:  m.MarshalIdent(stmt.Label),
		}
	})
}

func (m *Marshaller) MarshalBlockStmt(stmt *ast.BlockStmt) *BlockStmtNode {
	return wrapMarshal(m, stmt, func() *BlockStmtNode {
		return &BlockStmtNode{
			Node:   m.MarshalNode("BlockStmt", stmt),
			Lbrace: m.MarshalPosition(stmt.Lbrace),
			List:   m.MarshalStmts(stmt.List),
			Rbrace: m.MarshalPosition(stmt.Rbrace),
		}
	})
}

func (m *Marshaller) MarshalIfStmt(stmt *ast.IfStmt) *IfStmtNode {
	return wrapMarshal(m, stmt, func() *IfStmtNode {
		return &IfStmtNode{
			Node: m.MarshalNode("IfStmt", stmt),
			If:   m.MarshalPosition(stmt.If),
			Init: m.MarshalStmt(stmt.Init),
			Cond: m.MarshalExpr(stmt.Cond),
			Body: m.MarshalBlockStmt(stmt.Body),
			Else: m.MarshalStmt(stmt.Else),
		}
	})
}

func (m *Marshaller) MarshalCaseClause(stmt *ast.CaseClause) *CaseClauseNode {
	return wrapMarshal(m, stmt, func() *CaseClauseNode {
		return &CaseClauseNode{
			Node:  m.MarshalNode("CaseClause", stmt),
			Case:  m.MarshalPosition(stmt.Case),
			List:  m.MarshalExprs(stmt.List),
			Colon: m.MarshalPosition(stmt.Colon),
			Body:  m.MarshalStmts(stmt.Body),
		}
	})
}

func (m *Marshaller) MarshalSwitchStmt(stmt *ast.SwitchStmt) *SwitchStmtNode {
	return wrapMarshal(m, stmt, func() *SwitchStmtNode {
		return &SwitchStmtNode{
			Node:   m.MarshalNode("SwitchStmt", stmt),
			Switch: m.MarshalPosition(stmt.Switch),
			Init:   m.MarshalStmt(stmt.Init),
			Tag:    m.MarshalExpr(stmt.Tag),
			Body:   m.MarshalBlockStmt(stmt.Body),
		}
	})
}

func (m *Marshaller) MarshalTypeSwitchStmt(stmt *ast.TypeSwitchStmt) *TypeSwitchStmtNode {
	return wrapMarshal(m, stmt, func() *TypeSwitchStmtNode {
		return &TypeSwitchStmtNode{
			Node:   m.MarshalNode("TypeSwitchStmt", stmt),
			Switch: m.MarshalPosition(stmt.Switch),
			Init:   m.MarshalStmt(stmt.Init),
			Assign: m.MarshalStmt(stmt.Assign),
			Body:   m.MarshalBlockStmt(stmt.Body),
		}
	})
}

func (m *Marshaller) MarshalCommClause(stmt *ast.CommClause) *CommClauseNode {
	return wrapMarshal(m, stmt, func() *CommClauseNode {
		return &CommClauseNode{
			Node: m.MarshalNode("CommClause", stmt),
			Case: m.MarshalPosition(stmt.Case),
			Comm: m.MarshalStmt(stmt.Comm),
			Body: m.MarshalStmts(stmt.Body),
		}
	})
}

func (m *Marshaller) MarshalSelectStmt(stmt *ast.SelectStmt) *SelectStmtNode {
	return wrapMarshal(m, stmt, func() *SelectStmtNode {
		return &SelectStmtNode{
			Node:   m.MarshalNode("SelectStmt", stmt),
			Select: m.MarshalPosition(stmt.Select),
			Body:   m.MarshalBlockStmt(stmt.Body),
		}
	})
}

func (m *Marshaller) MarshalForStmt(stmt *ast.ForStmt) *ForStmtNode {
	return wrapMarshal(m, stmt, func() *ForStmtNode {
		return &ForStmtNode{
			Node: m.MarshalNode("ForStmt", stmt),
			For:  m.MarshalPosition(stmt.For),
			Init: m.MarshalStmt(stmt.Init),
			Cond: m.MarshalExpr(stmt.Cond),
			Post: m.MarshalStmt(stmt.Post),
			Body: m.MarshalBlockStmt(stmt.Body),
		}
	})
}

func (m *Marshaller) MarshalRangeStmt(stmt *ast.RangeStmt) *RangeStmtNode {
	return wrapMarshal(m, stmt, func() *RangeStmtNode {
		return &RangeStmtNode{
			Node:   m.MarshalNode("RangeStmt", stmt),
			For:    m.MarshalPosition(stmt.For),
			Key:    m.MarshalExpr(stmt.Key),
			Value:  m.MarshalExpr(stmt.Value),
			TokPos: m.MarshalPosition(stmt.TokPos),
			Tok:    stmt.Tok.String(),
			X:      m.MarshalExpr(stmt.X),
			Body:   m.MarshalBlockStmt(stmt.Body),
		}
	})
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
		panic("implement me " + reflect.TypeOf(stmt).String())
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
	return wrapMarshal(m, spec, func() *ImportSpecNode {
		return &ImportSpecNode{
			Node:    m.MarshalNode("ImportSpec", spec),
			Doc:     m.MarshalCommentGroup(spec.Doc),
			Name:    m.MarshalIdent(spec.Name),
			Path:    m.MarshalBasicLit(spec.Path),
			Comment: m.MarshalCommentGroup(spec.Comment),
			EndPos:  m.MarshalPosition(spec.EndPos),
		}
	})
}

func (m *Marshaller) MarshalImportSpecs(imports []*ast.ImportSpec) []*ImportSpecNode {
	if imports == nil {
		return nil
	}
	nodes := make([]*ImportSpecNode, len(imports))
	for index, importSpec := range imports {
		nodes[index] = m.MarshalImportSpec(importSpec)
	}
	return nodes
}

func (m *Marshaller) MarshalValueSpec(spec *ast.ValueSpec) *ValueSpecNode {
	return wrapMarshal(m, spec, func() *ValueSpecNode {
		return &ValueSpecNode{
			Node:    m.MarshalNode("ValueSpec", spec),
			Doc:     m.MarshalCommentGroup(spec.Doc),
			Names:   m.MarshalIdents(spec.Names),
			Type:    m.MarshalExpr(spec.Type),
			Values:  m.MarshalExprs(spec.Values),
			Comment: m.MarshalCommentGroup(spec.Comment),
		}
	})
}

func (m *Marshaller) MarshalTypeSpec(spec *ast.TypeSpec) *TypeSpecNode {
	return wrapMarshal(m, spec, func() *TypeSpecNode {
		return &TypeSpecNode{
			Node:       m.MarshalNode("TypeSpec", spec),
			Doc:        m.MarshalCommentGroup(spec.Doc),
			Name:       m.MarshalIdent(spec.Name),
			TypeParams: m.MarshalFieldList(spec.TypeParams),
			Assign:     m.MarshalPosition(spec.Assign),
			Type:       m.MarshalExpr(spec.Type),
			Comment:    m.MarshalCommentGroup(spec.Comment),
		}
	})
}

func (m *Marshaller) MarshalSpec(node ast.Spec) ISpecNode {
	if node == nil {
		return nil
	}
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
	return wrapMarshal(m, decl, func() *BadDeclNode {
		return &BadDeclNode{
			Node: m.MarshalNode("BadDecl", decl),
			From: m.MarshalPosition(decl.From),
			To:   m.MarshalPosition(decl.To),
		}
	})
}

func (m *Marshaller) MarshalGenDecl(decl *ast.GenDecl) *GenDeclNode {
	return wrapMarshal(m, decl, func() *GenDeclNode {
		return &GenDeclNode{
			Node:   m.MarshalNode("GenDecl", decl),
			Doc:    m.MarshalCommentGroup(decl.Doc),
			TokPos: m.MarshalPosition(decl.TokPos),
			Tok:    decl.Tok.String(),
			Lparen: m.MarshalPosition(decl.Lparen),
			Specs:  m.MarshalSpecs(decl.Specs),
			Rparen: m.MarshalPosition(decl.Rparen),
		}
	})
}

func (m *Marshaller) MarshalFuncDecl(decl *ast.FuncDecl) *FuncDeclNode {
	return wrapMarshal(m, decl, func() *FuncDeclNode {
		return &FuncDeclNode{
			Node: m.MarshalNode("FuncDecl", decl),
			Doc:  m.MarshalCommentGroup(decl.Doc),
			Recv: m.MarshalFieldList(decl.Recv),
			Name: m.MarshalIdent(decl.Name),
			Type: m.MarshalFuncType(decl.Type),
			Body: m.MarshalBlockStmt(decl.Body),
		}
	})
}

func (m *Marshaller) MarshalDecl(node ast.Decl) IDeclNode {
	if node == nil {
		return nil
	}
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
	nodes := make([]IDeclNode, len(decls))
	for index, decl := range decls {
		nodes[index] = m.MarshalDecl(decl)
	}
	return nodes
}

// ---------------------------------------------------------------------------

func (m *Marshaller) MarshalFile(node *ast.File) *FileNode {
	return wrapMarshal(m, node, func() *FileNode {
		return &FileNode{
			Node:       m.MarshalNode("File", node),
			Doc:        m.MarshalCommentGroup(node.Doc),
			Package:    m.MarshalPosition(node.Package),
			Name:       m.MarshalIdent(node.Name),
			Decls:      m.MarshalDecls(node.Decls),
			Imports:    m.MarshalImportSpecs(node.Imports),
			Unresolved: m.MarshalIdents(node.Unresolved),
			Comments:   m.MarshalCommentGroups(node.Comments),
			FileSet:    m.fset,
		}
	})
}
