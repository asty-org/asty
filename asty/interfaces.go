package asty

import "go/ast"

type IExprUnmarshaller interface {
	UnmarshalBadExprNode(node *BadExprNode) *ast.BadExpr
	UnmarshalIdentNode(node *IdentNode) *ast.Ident
	UnmarshalEllipsisNode(node *EllipsisNode) *ast.Ellipsis
	UnmarshalBasicLitNode(node *BasicLitNode) *ast.BasicLit
	UnmarshalFuncLitNode(node *FuncLitNode) *ast.FuncLit
	UnmarshalCompositeLitNode(node *CompositeLitNode) *ast.CompositeLit
	UnmarshalParenExprNode(node *ParenExprNode) *ast.ParenExpr
	UnmarshalSelectorExprNode(node *SelectorExprNode) *ast.SelectorExpr
	UnmarshalIndexExprNode(node *IndexExprNode) *ast.IndexExpr
	UnmarshalIndexListExprNode(node *IndexListExprNode) *ast.IndexListExpr
	UnmarshalSliceExprNode(node *SliceExprNode) *ast.SliceExpr
	UnmarshalTypeAssertExprNode(node *TypeAssertExprNode) *ast.TypeAssertExpr
	UnmarshalCallExprNode(node *CallExprNode) *ast.CallExpr
	UnmarshalStarExprNode(node *StarExprNode) *ast.StarExpr
	UnmarshalUnaryExprNode(node *UnaryExprNode) *ast.UnaryExpr
	UnmarshalBinaryExprNode(node *BinaryExprNode) *ast.BinaryExpr
	UnmarshalKeyValueExprNode(node *KeyValueExprNode) *ast.KeyValueExpr
	UnmarshalArrayTypeNode(node *ArrayTypeNode) *ast.ArrayType
	UnmarshalStructTypeNode(node *StructTypeNode) *ast.StructType
	UnmarshalFuncTypeNode(node *FuncTypeNode) *ast.FuncType
	UnmarshalInterfaceTypeNode(node *InterfaceTypeNode) *ast.InterfaceType
	UnmarshalMapTypeNode(node *MapTypeNode) *ast.MapType
	UnmarshalChanTypeNode(node *ChanTypeNode) *ast.ChanType
}

type IStmtUnmarshaller interface {
	UnmarshalBadStmtNode(node *BadStmtNode) *ast.BadStmt
	UnmarshalDeclStmtNode(node *DeclStmtNode) *ast.DeclStmt
	UnmarshalEmptyStmtNode(node *EmptyStmtNode) *ast.EmptyStmt
	UnmarshalLabeledStmtNode(node *LabeledStmtNode) *ast.LabeledStmt
	UnmarshalExprStmtNode(node *ExprStmtNode) *ast.ExprStmt
	UnmarshalSendStmtNode(node *SendStmtNode) *ast.SendStmt
	UnmarshalIncDecStmtNode(node *IncDecStmtNode) *ast.IncDecStmt
	UnmarshalAssignStmtNode(node *AssignStmtNode) *ast.AssignStmt
	UnmarshalGoStmtNode(node *GoStmtNode) *ast.GoStmt
	UnmarshalDeferStmtNode(node *DeferStmtNode) *ast.DeferStmt
	UnmarshalReturnStmtNode(node *ReturnStmtNode) *ast.ReturnStmt
	UnmarshalBranchStmtNode(node *BranchStmtNode) *ast.BranchStmt
	UnmarshalBlockStmtNode(node *BlockStmtNode) *ast.BlockStmt
	UnmarshalIfStmtNode(node *IfStmtNode) *ast.IfStmt
	UnmarshalCaseClauseNode(node *CaseClauseNode) *ast.CaseClause
	UnmarshalSwitchStmtNode(node *SwitchStmtNode) *ast.SwitchStmt
	UnmarshalTypeSwitchStmtNode(node *TypeSwitchStmtNode) *ast.TypeSwitchStmt
	UnmarshalCommClauseNode(node *CommClauseNode) *ast.CommClause
	UnmarshalSelectStmtNode(node *SelectStmtNode) *ast.SelectStmt
	UnmarshalForStmtNode(node *ForStmtNode) *ast.ForStmt
	UnmarshalRangeStmtNode(node *RangeStmtNode) *ast.RangeStmt
}

type ISpecUnmarshaller interface {
	UnmarshalImportSpecNode(node *ImportSpecNode) *ast.ImportSpec
	UnmarshalValueSpecNode(node *ValueSpecNode) *ast.ValueSpec
	UnmarshalTypeSpecNode(node *TypeSpecNode) *ast.TypeSpec
}

type IDeclUnmarshaller interface {
	UnmarshalBadDeclNode(node *BadDeclNode) *ast.BadDecl
	UnmarshalGenDeclNode(node *GenDeclNode) *ast.GenDecl
	UnmarshalFuncDeclNode(node *FuncDeclNode) *ast.FuncDecl
}

type INode interface {
	GetRefId() int
}

type IDeclNode interface {
	INode
	UnmarshalDecl(IDeclUnmarshaller) ast.Decl
}

type IExprNode interface {
	INode
	UnmarshalExpr(IExprUnmarshaller) ast.Expr
}

type ISpecNode interface {
	INode
	UnmarshalSpec(ISpecUnmarshaller) ast.Spec
}

type IStmtNode interface {
	INode
	UnmarshalStmt(IStmtUnmarshaller) ast.Stmt
}
