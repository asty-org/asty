package asty

import (
	"encoding/json"
	"go/ast"
)

type Node struct {
	NodeType string `json:"NodeType"`
}

type PositionNode struct {
	Node
	Filename string `json:"Filename,omitempty"`
	Offset   int    `json:"Offset,omitempty"`
	Line     int    `json:"Line,omitempty"`
	Column   int    `json:"Column,omitempty"`
}

type CommentNode struct {
	Node
	Slash *PositionNode `json:"Slash,omitempty"`
	Text  string        `json:"Text"`
}

type CommentGroupNode struct {
	Node
	List []*CommentNode `json:"List,omitempty"`
}

// ---------------------------------------------------------------------------

type FieldNode struct {
	Node
	Doc     *CommentGroupNode `json:"Doc,omitempty"`
	Names   []*IdentNode      `json:"Names"`
	Type    IExprNode         `json:"Type"`
	Tag     *BasicLitNode     `json:"Tag,omitempty"`
	Comment *CommentGroupNode `json:"Comment,omitempty"`
}
type FieldNodeAlias struct {
	Node
	Doc     *CommentGroupNode
	Names   []*IdentNode
	Type    json.RawMessage
	Tag     *BasicLitNode
	Comment *CommentGroupNode
}

type FieldListNode struct {
	Node
	Opening *PositionNode `json:"Opening,omitempty"`
	List    []*FieldNode  `json:"List"`
	Closing *PositionNode `json:"Closing,omitempty"`
}

// ---------------------------------------------------------------------------

type BadExprNode struct {
	Node
	From *PositionNode `json:"From,omitempty"`
	To   *PositionNode `json:"To,omitempty"`
}

type IdentNode struct {
	Node
	NamePos *PositionNode `json:"NamePos,omitempty"`
	Name    string        `json:"Name"`
	// Obj     *Object   // denoted object; or nil
}

type EllipsisNode struct {
	Node
	Ellipsis *PositionNode `json:"Ellipsis,omitempty"`
	Elt      IExprNode     `json:"Elt"`
}
type EllipsisNodeAlias struct {
	Node
	Ellipsis *PositionNode
	Elt      json.RawMessage
}

type BasicLitNode struct {
	Node
	ValuePos *PositionNode `json:"ValuePos,omitempty"`
	Kind     string        `json:"Kind"`
	Value    string        `json:"Value"`
}

type FuncLitNode struct {
	Node
	Type *FuncTypeNode  `json:"Type"`
	Body *BlockStmtNode `json:"Body"`
}

type CompositeLitNode struct {
	Node
	Type       IExprNode     `json:"Type"`
	Lbrace     *PositionNode `json:"Lbrace,omitempty"`
	Elts       []IExprNode   `json:"Elts"`
	Rbrace     *PositionNode `json:"Rbrace,omitempty"`
	Incomplete bool          `json:"Incomplete"`
}
type CompositeLitNodeAlias struct {
	Node
	Type       json.RawMessage
	Lbrace     *PositionNode
	Elts       []json.RawMessage
	Rbrace     *PositionNode
	Incomplete bool
}

type ParenExprNode struct {
	Node
	Lparen *PositionNode `json:"Lparen,omitempty"`
	X      IExprNode     `json:"X"`
	Rparen *PositionNode `json:"Rparen,omitempty"`
}
type ParenExprNodeAlias struct {
	Node
	Lparen *PositionNode
	X      json.RawMessage
	Rparen *PositionNode
}

type SelectorExprNode struct {
	Node
	X   IExprNode  `json:"X,omitempty"`
	Sel *IdentNode `json:"Sel,omitempty"`
}
type SelectorExprNodeAlias struct {
	Node
	X   json.RawMessage
	Sel *IdentNode
}

type IndexExprNode struct {
	Node
	X      IExprNode     `json:"X"`
	Lbrack *PositionNode `json:"Lbrack,omitempty"`
	Index  IExprNode     `json:"Index"`
	Rbrack *PositionNode `json:"Rbrack,omitempty"`
}
type IndexExprNodeAlias struct {
	Node
	X      json.RawMessage
	Lbrack *PositionNode
	Index  json.RawMessage
	Rbrack *PositionNode
}

type IndexListExprNode struct {
	Node
	X       IExprNode     `json:"X"`
	Lbrack  *PositionNode `json:"Lbrack,omitempty"`
	Indices []IExprNode   `json:"Indices"`
	Rbrack  *PositionNode `json:"Rbrack,omitempty"`
}
type IndexListExprNodeAlias struct {
	Node
	X       json.RawMessage
	Lbrack  *PositionNode
	Indices []json.RawMessage
	Rbrack  *PositionNode
}

type SliceExprNode struct {
	Node
	X      IExprNode     `json:"X"`
	Lbrack *PositionNode `json:"Lbrack,omitempty"`
	Low    IExprNode     `json:"Low"`
	High   IExprNode     `json:"High"`
	Max    IExprNode     `json:"Max"`
	Slice3 bool          `json:"Slice3"`
	Rbrack *PositionNode `json:"Rbrack,omitempty"`
}
type SliceExprNodeAlias struct {
	Node
	X      json.RawMessage
	Lbrack *PositionNode
	Low    json.RawMessage
	High   json.RawMessage
	Max    json.RawMessage
	Slice3 bool
	Rbrack *PositionNode
}

type TypeAssertExprNode struct {
	Node
	X      IExprNode     `json:"X"`
	Lparen *PositionNode `json:"Lparen,omitempty"`
	Type   IExprNode     `json:"Type"`
	Rparen *PositionNode `json:"Rparen,omitempty"`
}
type TypeAssertExprNodeAlias struct {
	Node
	X      json.RawMessage
	Lparen *PositionNode
	Type   json.RawMessage
	Rparen *PositionNode
}

type CallExprNode struct {
	Node
	Fun      IExprNode     `json:"Fun"`
	Lparen   *PositionNode `json:"Lparen,omitempty"`
	Args     []IExprNode   `json:"Args"`
	Ellipsis *PositionNode `json:"Ellipsis,omitempty"`
	Rparen   *PositionNode `json:"Rparen,omitempty"`
}
type CallExprNodeAlias struct {
	Node
	Fun      json.RawMessage
	Lparen   *PositionNode
	Args     []json.RawMessage
	Ellipsis *PositionNode
	Rparen   *PositionNode
}

type StarExprNode struct {
	Node
	Star *PositionNode `json:"Star,omitempty"`
	X    IExprNode     `json:"X"`
}
type StarExprNodeAlias struct {
	Node
	Star *PositionNode
	X    json.RawMessage
}

type UnaryExprNode struct {
	Node
	OpPos *PositionNode `json:"OpPos,omitempty"`
	Op    string        `json:"Op"`
	X     IExprNode     `json:"X"`
}
type UnaryExprNodeAlias struct {
	Node
	OpPos *PositionNode
	Op    string
	X     json.RawMessage
}

type BinaryExprNode struct {
	Node
	X     IExprNode     `json:"X"`
	OpPos *PositionNode `json:"OpPos,omitempty"`
	Op    string        `json:"Op"`
	Y     IExprNode     `json:"Y"`
}
type BinaryExprNodeAlias struct {
	Node
	X     json.RawMessage
	OpPos *PositionNode
	Op    string
	Y     json.RawMessage
}

type KeyValueExprNode struct {
	Node
	Key   IExprNode     `json:"Key"`
	Colon *PositionNode `json:"Colon,omitempty"`
	Value IExprNode     `json:"Value"`
}
type KeyValueExprNodeAlias struct {
	Node
	Key   json.RawMessage
	Colon *PositionNode
	Value json.RawMessage
}

// ---------------------------------------------------------------------------

type ArrayTypeNode struct {
	Node
	Lbrack *PositionNode `json:"Lbrack,omitempty"`
	Len    IExprNode     `json:"Len"`
	Elt    IExprNode     `json:"Elt"`
}
type ArrayTypeNodeAlias struct {
	Node
	Lbrack *PositionNode
	Len    json.RawMessage
	Elt    json.RawMessage
}

type StructTypeNode struct {
	Node
	Struct     *PositionNode  `json:"Struct,omitempty"`
	Fields     *FieldListNode `json:"Fields"`
	Incomplete bool           `json:"Incomplete"`
}

type FuncTypeNode struct {
	Node
	Func       *PositionNode  `json:"Func,omitempty"`
	TypeParams *FieldListNode `json:"TypeParams"`
	Params     *FieldListNode `json:"Params"`
	Results    *FieldListNode `json:"Results"`
}

type InterfaceTypeNode struct {
	Node
	Interface  *PositionNode  `json:"Interface,omitempty"`
	Methods    *FieldListNode `json:"Methods"`
	Incomplete bool           `json:"Incomplete"`
}

type MapTypeNode struct {
	Node
	Map   *PositionNode `json:"Map,omitempty"`
	Key   IExprNode     `json:"Key"`
	Value IExprNode     `json:"Value"`
}
type MapTypeNodeAlias struct {
	Node
	Map   *PositionNode
	Key   json.RawMessage
	Value json.RawMessage
}

type ChanTypeNode struct {
	Node
	Begin *PositionNode `json:"Begin,omitempty"`
	Arrow *PositionNode `json:"Arrow,omitempty"`
	Dir   string        `json:"Dir"`
	Value IExprNode     `json:"Value"`
}
type ChanTypeNodeAlias struct {
	Node
	Begin *PositionNode
	Arrow *PositionNode
	Dir   string
	Value json.RawMessage
}

// ---------------------------------------------------------------------------

type BadStmtNode struct {
	Node
	From *PositionNode `json:"From,omitempty"`
	To   *PositionNode `json:"To,omitempty"`
}

type DeclStmtNode struct {
	Node
	Decl IDeclNode `json:"Decl"`
}
type DeclStmtNodeAlias struct {
	Node
	Decl json.RawMessage
}

type EmptyStmtNode struct {
	Node
	Semicolon *PositionNode `json:"Semicolon,omitempty"`
	Implicit  bool          `json:"Implicit"`
}

type LabeledStmtNode struct {
	Node
	Label *IdentNode    `json:"Label"`
	Colon *PositionNode `json:"Colon,omitempty"`
	Stmt  IStmtNode     `json:"Stmt"`
}
type LabeledStmtNodeAlias struct {
	Node
	Label *IdentNode
	Colon *PositionNode
	Stmt  json.RawMessage
}

type ExprStmtNode struct {
	Node
	X IExprNode `json:"X,omitempty"`
}
type ExprStmtNodeAlias struct {
	Node
	X json.RawMessage
}

type SendStmtNode struct {
	Node
	Chan  IExprNode     `json:"Chan"`
	Arrow *PositionNode `json:"Arrow,omitempty"`
	Value IExprNode     `json:"Value"`
}
type SendStmtNodeAlias struct {
	Node
	Chan  json.RawMessage
	Arrow *PositionNode
	Value json.RawMessage
}

type IncDecStmtNode struct {
	Node
	X      IExprNode     `json:"X"`
	TokPos *PositionNode `json:"TokPos,omitempty"`
	Tok    string        `json:"Tok"`
}
type IncDecStmtNodeAlias struct {
	Node
	X      json.RawMessage
	TokPos *PositionNode
	Tok    string
}

type AssignStmtNode struct {
	Node
	Lhs    []IExprNode   `json:"Lhs"`
	TokPos *PositionNode `json:"TokPos,omitempty"`
	Tok    string        `json:"Tok"`
	Rhs    []IExprNode   `json:"Rhs"`
}
type AssignStmtNodeAlias struct {
	Node
	Lhs    []json.RawMessage
	TokPos *PositionNode
	Tok    string
	Rhs    []json.RawMessage
}

type GoStmtNode struct {
	Node
	Go   *PositionNode `json:"Go,omitempty"`
	Call *CallExprNode `json:"Call"`
}

type DeferStmtNode struct {
	Node
	Defer *PositionNode `json:"Defer,omitempty"`
	Call  *CallExprNode `json:"Call"`
}

type ReturnStmtNode struct {
	Node
	Return  *PositionNode `json:"Return,omitempty"`
	Results []IExprNode   `json:"Results"`
}
type ReturnStmtNodeAlias struct {
	Node
	Return  *PositionNode
	Results []json.RawMessage
}

type BranchStmtNode struct {
	Node
	TokPos *PositionNode `json:"TokPos,omitempty"`
	Tok    string        `json:"Tok"`
	Label  *IdentNode    `json:"Label"`
}

type BlockStmtNode struct {
	Node
	Lbrace *PositionNode `json:"Lbrace,omitempty"`
	List   []IStmtNode   `json:"List"`
	Rbrace *PositionNode `json:"Rbrace,omitempty"`
}
type BlockStmtNodeAlias struct {
	Node
	Lbrace *PositionNode
	List   []json.RawMessage
	Rbrace *PositionNode
}

type IfStmtNode struct {
	Node
	If   *PositionNode  `json:"If,omitempty"`
	Init IStmtNode      `json:"Init"`
	Cond IExprNode      `json:"Cond"`
	Body *BlockStmtNode `json:"Body"`
	Else IStmtNode      `json:"Else"`
}
type IfStmtNodeAlias struct {
	Node
	If   *PositionNode
	Init json.RawMessage
	Cond json.RawMessage
	Body *BlockStmtNode
	Else json.RawMessage
}

type CaseClauseNode struct {
	Node
	Case  *PositionNode `json:"Case,omitempty"`
	List  []IExprNode   `json:"List"`
	Colon *PositionNode `json:"Colon,omitempty"`
	Body  []IStmtNode   `json:"Body"`
}
type CaseClauseNodeAlias struct {
	Node
	Case  *PositionNode
	List  []json.RawMessage
	Colon *PositionNode
	Body  []json.RawMessage
}

type SwitchStmtNode struct {
	Node
	Switch *PositionNode  `json:"Switch,omitempty"`
	Init   IStmtNode      `json:"Init"`
	Tag    IExprNode      `json:"Tag"`
	Body   *BlockStmtNode `json:"Body"`
}
type SwitchStmtNodeAlias struct {
	Node
	Switch *PositionNode
	Init   json.RawMessage
	Tag    json.RawMessage
	Body   *BlockStmtNode
}

type TypeSwitchStmtNode struct {
	Node
	Switch *PositionNode  `json:"Switch,omitempty"`
	Init   IStmtNode      `json:"Init"`
	Assign IStmtNode      `json:"Assign"`
	Body   *BlockStmtNode `json:"Body"`
}
type TypeSwitchStmtNodeAlias struct {
	Node
	Switch *PositionNode
	Init   json.RawMessage
	Assign json.RawMessage
	Body   *BlockStmtNode
}

type CommClauseNode struct {
	Node
	Case  *PositionNode `json:"Case,omitempty"`
	Comm  IStmtNode     `json:"Comm"`
	Colon *PositionNode `json:"Colon,omitempty"`
	Body  []IStmtNode   `json:"Body"`
}
type CommClauseNodeAlias struct {
	Node
	Case  *PositionNode
	Comm  json.RawMessage
	Colon *PositionNode
	Body  []json.RawMessage
}

type SelectStmtNode struct {
	Node
	Select *PositionNode  `json:"Select,omitempty"`
	Body   *BlockStmtNode `json:"Body"`
}

type ForStmtNode struct {
	Node
	For  *PositionNode  `json:"For,omitempty"`
	Init IStmtNode      `json:"Init"`
	Cond IExprNode      `json:"Cond"`
	Post IStmtNode      `json:"Post"`
	Body *BlockStmtNode `json:"Body"`
}
type ForStmtNodeAlias struct {
	Node
	For  *PositionNode
	Init json.RawMessage
	Cond json.RawMessage
	Post json.RawMessage
	Body *BlockStmtNode
}

type RangeStmtNode struct {
	Node
	For    *PositionNode  `json:"For,omitempty"`
	Key    IExprNode      `json:"Key"`
	Value  IExprNode      `json:"Value"`
	TokPos *PositionNode  `json:"TokPos,omitempty"`
	Tok    string         `json:"Tok"`
	X      IExprNode      `json:"X"`
	Body   *BlockStmtNode `json:"Body"`
}
type RangeStmtNodeAlias struct {
	Node
	For    *PositionNode
	Key    json.RawMessage
	Value  json.RawMessage
	TokPos *PositionNode
	Tok    string
	X      json.RawMessage
	Body   *BlockStmtNode
}

// ---------------------------------------------------------------------------

type ImportSpecNode struct {
	Node
	Doc     *CommentGroupNode `json:"Doc,omitempty"`
	Name    *IdentNode        `json:"Name"`
	Path    *BasicLitNode     `json:"Path"`
	Comment *CommentGroupNode `json:"Comment,omitempty"`
	EndPos  *PositionNode     `json:"EndPos,omitempty"`
}

type ValueSpecNode struct {
	Node
	Doc     *CommentGroupNode `json:"Doc,omitempty"`
	Names   []*IdentNode      `json:"Names"`
	Type    IExprNode         `json:"Type"`
	Values  []IExprNode       `json:"Values"`
	Comment *CommentGroupNode `json:"Comment,omitempty"`
}
type ValueSpecNodeAlias struct {
	Node
	Doc     *CommentGroupNode
	Names   []*IdentNode
	Type    json.RawMessage
	Values  []json.RawMessage
	Comment *CommentGroupNode
}

type TypeSpecNode struct {
	Node
	Doc        *CommentGroupNode `json:"Doc,omitempty"`
	Name       *IdentNode        `json:"Name"`
	TypeParams *FieldListNode    `json:"TypeParams"`
	Assign     *PositionNode     `json:"Assign,omitempty"`
	Type       IExprNode         `json:"Type"`
	Comment    *CommentGroupNode `json:"Comment,omitempty"`
}
type TypeSpecNodeAlias struct {
	Node
	Doc        *CommentGroupNode
	Name       *IdentNode
	TypeParams *FieldListNode
	Assign     *PositionNode
	Type       json.RawMessage
	Comment    *CommentGroupNode
}

// ---------------------------------------------------------------------------

type BadDeclNode struct {
	Node
	From *PositionNode `json:"From,omitempty"`
	To   *PositionNode `json:"To,omitempty"`
}

type GenDeclNode struct {
	Node
	Doc    *CommentGroupNode `json:"Doc,omitempty"`
	TokPos *PositionNode     `json:"TokPos,omitempty"`
	Tok    string            `json:"Tok"`
	Lparen *PositionNode     `json:"Lparen,omitempty"`
	Specs  []ISpecNode       `json:"Specs"`
	Rparen *PositionNode     `json:"Rparen,omitempty"`
}
type GenDeclNodeAlias struct {
	Node
	Doc    *CommentGroupNode
	TokPos *PositionNode
	Tok    string
	Lparen *PositionNode
	Specs  []json.RawMessage
	Rparen *PositionNode
}

type FuncDeclNode struct {
	Node
	Doc  *CommentGroupNode `json:"Doc,omitempty"`
	Recv *FieldListNode    `json:"Recv"`
	Name *IdentNode        `json:"Name"`
	Type *FuncTypeNode     `json:"Type"`
	Body *BlockStmtNode    `json:"Body"`
}

// ---------------------------------------------------------------------------

type FileNode struct {
	Node
	Doc        *CommentGroupNode   `json:"Doc,omitempty"`
	Package    *PositionNode       `json:"Package,omitempty"`
	Name       *IdentNode          `json:"Name"`
	Decls      []IDeclNode         `json:"Decls"`
	Imports    []*ImportSpecNode   `json:"Imports,omitempty"`
	Unresolved []*IdentNode        `json:"Unresolved,omitempty"`
	Comments   []*CommentGroupNode `json:"Comments,omitempty"`
	//	Scope      *Scope
}
type FileNodeAlias struct {
	Node
	Doc        *CommentGroupNode
	Package    *PositionNode
	Name       *IdentNode
	Decls      []json.RawMessage
	Imports    []*ImportSpecNode
	Unresolved []*IdentNode
	Comments   []*CommentGroupNode
	//	Scope      *Scope
}

type PackageNode struct {
	Node
	Name string
	//	Scope   *Scope
	//	Imports map[string]*Object
	Files map[string]*FileNode
}

func (node *BadExprNode) UnmarshalExpr(um IExprUnmarshaller) ast.Expr {
	return um.UnmarshalBadExprNode(node)
}
func (node *IdentNode) UnmarshalExpr(um IExprUnmarshaller) ast.Expr {
	return um.UnmarshalIdentNode(node)
}
func (node *BasicLitNode) UnmarshalExpr(um IExprUnmarshaller) ast.Expr {
	return um.UnmarshalBasicLitNode(node)
}
func (node *EllipsisNode) UnmarshalExpr(um IExprUnmarshaller) ast.Expr {
	return um.UnmarshalEllipsisNode(node)
}
func (node *FuncLitNode) UnmarshalExpr(um IExprUnmarshaller) ast.Expr {
	return um.UnmarshalFuncLitNode(node)
}
func (node *CompositeLitNode) UnmarshalExpr(um IExprUnmarshaller) ast.Expr {
	return um.UnmarshalCompositeLitNode(node)
}
func (node *ParenExprNode) UnmarshalExpr(um IExprUnmarshaller) ast.Expr {
	return um.UnmarshalParenExprNode(node)
}
func (node *SelectorExprNode) UnmarshalExpr(um IExprUnmarshaller) ast.Expr {
	return um.UnmarshalSelectorExprNode(node)
}
func (node *IndexExprNode) UnmarshalExpr(um IExprUnmarshaller) ast.Expr {
	return um.UnmarshalIndexExprNode(node)
}
func (node *IndexListExprNode) UnmarshalExpr(um IExprUnmarshaller) ast.Expr {
	return um.UnmarshalIndexListExprNode(node)
}
func (node *SliceExprNode) UnmarshalExpr(um IExprUnmarshaller) ast.Expr {
	return um.UnmarshalSliceExprNode(node)
}
func (node *TypeAssertExprNode) UnmarshalExpr(um IExprUnmarshaller) ast.Expr {
	return um.UnmarshalTypeAssertExprNode(node)
}
func (node *CallExprNode) UnmarshalExpr(um IExprUnmarshaller) ast.Expr {
	return um.UnmarshalCallExprNode(node)
}
func (node *StarExprNode) UnmarshalExpr(um IExprUnmarshaller) ast.Expr {
	return um.UnmarshalStarExprNode(node)
}
func (node *UnaryExprNode) UnmarshalExpr(um IExprUnmarshaller) ast.Expr {
	return um.UnmarshalUnaryExprNode(node)
}
func (node *BinaryExprNode) UnmarshalExpr(um IExprUnmarshaller) ast.Expr {
	return um.UnmarshalBinaryExprNode(node)
}
func (node *KeyValueExprNode) UnmarshalExpr(um IExprUnmarshaller) ast.Expr {
	return um.UnmarshalKeyValueExprNode(node)
}
func (node *ArrayTypeNode) UnmarshalExpr(um IExprUnmarshaller) ast.Expr {
	return um.UnmarshalArrayTypeNode(node)
}
func (node *StructTypeNode) UnmarshalExpr(um IExprUnmarshaller) ast.Expr {
	return um.UnmarshalStructTypeNode(node)
}
func (node *FuncTypeNode) UnmarshalExpr(um IExprUnmarshaller) ast.Expr {
	return um.UnmarshalFuncTypeNode(node)
}
func (node *InterfaceTypeNode) UnmarshalExpr(um IExprUnmarshaller) ast.Expr {
	return um.UnmarshalInterfaceTypeNode(node)
}
func (node *MapTypeNode) UnmarshalExpr(um IExprUnmarshaller) ast.Expr {
	return um.UnmarshalMapTypeNode(node)
}
func (node *ChanTypeNode) UnmarshalExpr(um IExprUnmarshaller) ast.Expr {
	return um.UnmarshalChanTypeNode(node)
}

func (node *BadStmtNode) UnmarshalStmt(um IStmtUnmarshaller) ast.Stmt {
	return um.UnmarshalBadStmtNode(node)
}
func (node *DeclStmtNode) UnmarshalStmt(um IStmtUnmarshaller) ast.Stmt {
	return um.UnmarshalDeclStmtNode(node)
}
func (node *EmptyStmtNode) UnmarshalStmt(um IStmtUnmarshaller) ast.Stmt {
	return um.UnmarshalEmptyStmtNode(node)
}
func (node *LabeledStmtNode) UnmarshalStmt(um IStmtUnmarshaller) ast.Stmt {
	return um.UnmarshalLabeledStmtNode(node)
}
func (node *ExprStmtNode) UnmarshalStmt(um IStmtUnmarshaller) ast.Stmt {
	return um.UnmarshalExprStmtNode(node)
}
func (node *SendStmtNode) UnmarshalStmt(um IStmtUnmarshaller) ast.Stmt {
	return um.UnmarshalSendStmtNode(node)
}
func (node *IncDecStmtNode) UnmarshalStmt(um IStmtUnmarshaller) ast.Stmt {
	return um.UnmarshalIncDecStmtNode(node)
}
func (node *AssignStmtNode) UnmarshalStmt(um IStmtUnmarshaller) ast.Stmt {
	return um.UnmarshalAssignStmtNode(node)
}
func (node *GoStmtNode) UnmarshalStmt(um IStmtUnmarshaller) ast.Stmt {
	return um.UnmarshalGoStmtNode(node)
}
func (node *DeferStmtNode) UnmarshalStmt(um IStmtUnmarshaller) ast.Stmt {
	return um.UnmarshalDeferStmtNode(node)
}
func (node *ReturnStmtNode) UnmarshalStmt(um IStmtUnmarshaller) ast.Stmt {
	return um.UnmarshalReturnStmtNode(node)
}
func (node *BranchStmtNode) UnmarshalStmt(um IStmtUnmarshaller) ast.Stmt {
	return um.UnmarshalBranchStmtNode(node)
}
func (node *BlockStmtNode) UnmarshalStmt(um IStmtUnmarshaller) ast.Stmt {
	return um.UnmarshalBlockStmtNode(node)
}
func (node *IfStmtNode) UnmarshalStmt(um IStmtUnmarshaller) ast.Stmt {
	return um.UnmarshalIfStmtNode(node)
}
func (node *CaseClauseNode) UnmarshalStmt(um IStmtUnmarshaller) ast.Stmt {
	return um.UnmarshalCaseClauseNode(node)
}
func (node *SwitchStmtNode) UnmarshalStmt(um IStmtUnmarshaller) ast.Stmt {
	return um.UnmarshalSwitchStmtNode(node)
}
func (node *TypeSwitchStmtNode) UnmarshalStmt(um IStmtUnmarshaller) ast.Stmt {
	return um.UnmarshalTypeSwitchStmtNode(node)
}
func (node *CommClauseNode) UnmarshalStmt(um IStmtUnmarshaller) ast.Stmt {
	return um.UnmarshalCommClauseNode(node)
}
func (node *SelectStmtNode) UnmarshalStmt(um IStmtUnmarshaller) ast.Stmt {
	return um.UnmarshalSelectStmtNode(node)
}
func (node *ForStmtNode) UnmarshalStmt(um IStmtUnmarshaller) ast.Stmt {
	return um.UnmarshalForStmtNode(node)
}
func (node *RangeStmtNode) UnmarshalStmt(um IStmtUnmarshaller) ast.Stmt {
	return um.UnmarshalRangeStmtNode(node)
}

func (node *ImportSpecNode) UnmarshalSpec(um ISpecUnmarshaller) ast.Spec {
	return um.UnmarshalImportSpecNode(node)
}
func (node *ValueSpecNode) UnmarshalSpec(um ISpecUnmarshaller) ast.Spec {
	return um.UnmarshalValueSpecNode(node)
}
func (node *TypeSpecNode) UnmarshalSpec(um ISpecUnmarshaller) ast.Spec {
	return um.UnmarshalTypeSpecNode(node)
}

func (node *BadDeclNode) UnmarshalDecl(um IDeclUnmarshaller) ast.Decl {
	return um.UnmarshalBadDeclNode(node)
}
func (node *GenDeclNode) UnmarshalDecl(um IDeclUnmarshaller) ast.Decl {
	return um.UnmarshalGenDeclNode(node)
}
func (node *FuncDeclNode) UnmarshalDecl(um IDeclUnmarshaller) ast.Decl {
	return um.UnmarshalFuncDeclNode(node)
}

func MakeExpr(nodeType string) IExprNode {
	switch nodeType {
	case "BadExpr":
		return &BadExprNode{}
	case "Ident":
		return &IdentNode{}
	case "BasicLit":
		return &BasicLitNode{}
	case "FuncLit":
		return &FuncLitNode{}
	case "CompositeLit":
		return &CompositeLitNode{}
	case "ParenExpr":
		return &ParenExprNode{}
	case "SelectorExpr":
		return &SelectorExprNode{}
	case "IndexExpr":
		return &IndexExprNode{}
	case "SliceExpr":
		return &SliceExprNode{}
	case "TypeAssertExpr":
		return &TypeAssertExprNode{}
	case "CallExpr":
		return &CallExprNode{}
	case "StarExpr":
		return &StarExprNode{}
	case "UnaryExpr":
		return &UnaryExprNode{}
	case "BinaryExpr":
		return &BinaryExprNode{}
	case "KeyValueExpr":
		return &KeyValueExprNode{}
	case "ArrayType":
		return &ArrayTypeNode{}
	case "StructType":
		return &StructTypeNode{}
	case "FuncType":
		return &FuncTypeNode{}
	case "InterfaceType":
		return &InterfaceTypeNode{}
	case "MapType":
		return &MapTypeNode{}
	case "ChanType":
		return &ChanTypeNode{}
	default:
		panic("implement me")
	}
}

func MakeStmt(nodeType string) IStmtNode {
	switch nodeType {
	case "BadStmt":
		return &BadStmtNode{}
	case "DeclStmt":
		return &DeclStmtNode{}
	case "EmptyStmt":
		return &EmptyStmtNode{}
	case "LabeledStmt":
		return &LabeledStmtNode{}
	case "ExprStmt":
		return &ExprStmtNode{}
	case "SendStmt":
		return &SendStmtNode{}
	case "IncDecStmt":
		return &IncDecStmtNode{}
	case "AssignStmt":
		return &AssignStmtNode{}
	case "GoStmt":
		return &GoStmtNode{}
	case "DeferStmt":
		return &DeferStmtNode{}
	case "ReturnStmt":
		return &ReturnStmtNode{}
	case "BranchStmt":
		return &BranchStmtNode{}
	case "BlockStmt":
		return &BlockStmtNode{}
	case "IfStmt":
		return &IfStmtNode{}
	case "CaseClause":
		return &CaseClauseNode{}
	case "SwitchStmt":
		return &SwitchStmtNode{}
	case "TypeSwitchStmt":
		return &TypeSwitchStmtNode{}
	case "CommClause":
		return &CommClauseNode{}
	case "SelectStmt":
		return &SelectStmtNode{}
	case "ForStmt":
		return &ForStmtNode{}
	case "RangeStmt":
		return &RangeStmtNode{}
	default:
		panic("implement me")
	}
}

func MakeSpec(nodeType string) ISpecNode {
	switch nodeType {
	case "ImportSpec":
		return &ImportSpecNode{}
	case "ValueSpec":
		return &ValueSpecNode{}
	case "TypeSpec":
		return &TypeSpecNode{}
	default:
		panic("implement me")
	}
}

func MakeDecl(nodeType string) IDeclNode {
	switch nodeType {
	case "BadDecl":
		return &BadDeclNode{}
	case "GenDecl":
		return &GenDeclNode{}
	case "FuncDecl":
		return &FuncDeclNode{}
	default:
		panic("implement me")
	}
}

func UnmarshalJSONExpr(data json.RawMessage) (IExprNode, error) {
	var node *Node
	err := json.Unmarshal(data, &node)
	if err != nil {
		return nil, err
	}
	if node == nil {
		return nil, nil
	}

	result := MakeExpr(node.NodeType)
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UnmarshalJSONExprs(data []json.RawMessage) ([]IExprNode, error) {
	if data == nil {
		return nil, nil
	}
	result := make([]IExprNode, len(data))
	for i, d := range data {
		expr, err := UnmarshalJSONExpr(d)
		if err != nil {
			return nil, err
		}
		result[i] = expr
	}
	return result, nil
}

func UnmarshalJSONStmt(data json.RawMessage) (IStmtNode, error) {
	var node *Node
	err := json.Unmarshal(data, &node)
	if err != nil {
		return nil, err
	}
	if node == nil {
		return nil, nil
	}

	result := MakeStmt(node.NodeType)
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UnmarshalJSONStmts(data []json.RawMessage) ([]IStmtNode, error) {
	if data == nil {
		return nil, nil
	}
	result := make([]IStmtNode, len(data))
	for i, d := range data {
		stmt, err := UnmarshalJSONStmt(d)
		if err != nil {
			return nil, err
		}
		result[i] = stmt
	}
	return result, nil
}

func UnmarshalJSONSpec(data json.RawMessage) (ISpecNode, error) {
	var node *Node
	err := json.Unmarshal(data, &node)
	if err != nil {
		return nil, err
	}
	if node == nil {
		return nil, nil
	}

	result := MakeSpec(node.NodeType)
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UnmarshalJSONSpecs(data []json.RawMessage) ([]ISpecNode, error) {
	if data == nil {
		return nil, nil
	}
	result := make([]ISpecNode, len(data))
	for i, d := range data {
		spec, err := UnmarshalJSONSpec(d)
		if err != nil {
			return nil, err
		}
		result[i] = spec
	}
	return result, nil
}

func UnmarshalJSONDecl(data json.RawMessage) (IDeclNode, error) {
	var node *Node
	err := json.Unmarshal(data, &node)
	if err != nil {
		return nil, err
	}
	if node == nil {
		return nil, nil
	}

	result := MakeDecl(node.NodeType)
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UnmarshalJSONDecls(data []json.RawMessage) ([]IDeclNode, error) {
	if data == nil {
		return nil, nil
	}
	result := make([]IDeclNode, len(data))
	for i, d := range data {
		decl, err := UnmarshalJSONDecl(d)
		if err != nil {
			return nil, err
		}
		result[i] = decl
	}
	return result, nil
}

func (node *FieldNode) UnmarshalJSON(data []byte) error {
	var alias FieldNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Doc = alias.Doc
	node.Names = alias.Names
	node.Type, err = UnmarshalJSONExpr(alias.Type)
	if err != nil {
		return err
	}
	node.Tag = alias.Tag
	node.Comment = alias.Comment
	return nil
}

func (node *EllipsisNode) UnmarshalJSON(data []byte) error {
	var alias EllipsisNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Ellipsis = alias.Ellipsis
	node.Elt, err = UnmarshalJSONExpr(alias.Elt)
	if err != nil {
		return err
	}
	return nil
}

func (node *CompositeLitNode) UnmarshalJSON(data []byte) error {
	var alias CompositeLitNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Type, err = UnmarshalJSONExpr(alias.Type)
	if err != nil {
		return err
	}
	node.Lbrace = alias.Lbrace
	node.Elts, err = UnmarshalJSONExprs(alias.Elts)
	if err != nil {
		return err
	}
	node.Rbrace = alias.Rbrace
	node.Incomplete = alias.Incomplete
	return nil
}

func (node *ParenExprNode) UnmarshalJSON(data []byte) error {
	var alias ParenExprNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Lparen = alias.Lparen
	node.X, err = UnmarshalJSONExpr(alias.X)
	if err != nil {
		return err
	}
	node.Rparen = alias.Rparen
	return nil
}

func (node *SelectorExprNode) UnmarshalJSON(data []byte) error {
	var alias SelectorExprNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.X, err = UnmarshalJSONExpr(alias.X)
	if err != nil {
		return err
	}
	node.Sel = alias.Sel
	return nil
}

func (node *IndexExprNode) UnmarshalJSON(data []byte) error {
	var alias IndexExprNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.X, err = UnmarshalJSONExpr(alias.X)
	if err != nil {
		return err
	}
	node.Lbrack = alias.Lbrack
	node.Index, err = UnmarshalJSONExpr(alias.Index)
	if err != nil {
		return err
	}
	node.Rbrack = alias.Rbrack
	return nil
}

func (node *IndexListExprNode) UnmarshalJSON(data []byte) error {
	var alias IndexListExprNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.X, err = UnmarshalJSONExpr(alias.X)
	if err != nil {
		return err
	}
	node.Lbrack = alias.Lbrack
	node.Indices, err = UnmarshalJSONExprs(alias.Indices)
	if err != nil {
		return err
	}
	node.Rbrack = alias.Rbrack
	return nil
}

func (node *SliceExprNode) UnmarshalJSON(data []byte) error {
	var alias SliceExprNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.X, err = UnmarshalJSONExpr(alias.X)
	if err != nil {
		return err
	}
	node.Lbrack = alias.Lbrack
	node.Low, err = UnmarshalJSONExpr(alias.Low)
	if err != nil {
		return err
	}
	node.High, err = UnmarshalJSONExpr(alias.High)
	if err != nil {
		return err
	}
	node.Max, err = UnmarshalJSONExpr(alias.Max)
	if err != nil {
		return err
	}
	node.Rbrack = alias.Rbrack
	node.Slice3 = alias.Slice3
	return nil
}

func (node *TypeAssertExprNode) UnmarshalJSON(data []byte) error {
	var alias TypeAssertExprNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.X, err = UnmarshalJSONExpr(alias.X)
	if err != nil {
		return err
	}
	node.Lparen = alias.Lparen
	node.Type, err = UnmarshalJSONExpr(alias.Type)
	if err != nil {
		return err
	}
	node.Rparen = alias.Rparen
	return nil
}

func (node *CallExprNode) UnmarshalJSON(data []byte) error {
	var alias CallExprNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Fun, err = UnmarshalJSONExpr(alias.Fun)
	if err != nil {
		return err
	}
	node.Lparen = alias.Lparen
	node.Args, err = UnmarshalJSONExprs(alias.Args)
	if err != nil {
		return err
	}
	node.Ellipsis = alias.Ellipsis
	node.Rparen = alias.Rparen
	return nil
}

func (node *StarExprNode) UnmarshalJSON(data []byte) error {
	var alias StarExprNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Star = alias.Star
	node.X, err = UnmarshalJSONExpr(alias.X)
	if err != nil {
		return err
	}
	return nil
}

func (node *UnaryExprNode) UnmarshalJSON(data []byte) error {
	var alias UnaryExprNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.OpPos = alias.OpPos
	node.Op = alias.Op
	node.X, err = UnmarshalJSONExpr(alias.X)
	if err != nil {
		return err
	}
	return nil
}

func (node *BinaryExprNode) UnmarshalJSON(data []byte) error {
	var alias BinaryExprNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.X, err = UnmarshalJSONExpr(alias.X)
	if err != nil {
		return err
	}
	node.OpPos = alias.OpPos
	node.Op = alias.Op
	node.Y, err = UnmarshalJSONExpr(alias.Y)
	if err != nil {
		return err
	}
	return nil
}

func (node *KeyValueExprNode) UnmarshalJSON(data []byte) error {
	var alias KeyValueExprNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Key, err = UnmarshalJSONExpr(alias.Key)
	if err != nil {
		return err
	}
	node.Colon = alias.Colon
	node.Value, err = UnmarshalJSONExpr(alias.Value)
	if err != nil {
		return err
	}
	return nil
}

func (node *ArrayTypeNode) UnmarshalJSON(data []byte) error {
	var alias ArrayTypeNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Lbrack = alias.Lbrack
	node.Len, err = UnmarshalJSONExpr(alias.Len)
	if err != nil {
		return err
	}
	node.Elt, err = UnmarshalJSONExpr(alias.Elt)
	if err != nil {
		return err
	}
	return nil
}

func (node *MapTypeNode) UnmarshalJSON(data []byte) error {
	var alias MapTypeNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Map = alias.Map
	node.Key, err = UnmarshalJSONExpr(alias.Key)
	if err != nil {
		return err
	}
	node.Value, err = UnmarshalJSONExpr(alias.Value)
	if err != nil {
		return err
	}
	return nil
}

func (node *ChanTypeNode) UnmarshalJSON(data []byte) error {
	var alias ChanTypeNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Begin = alias.Begin
	node.Arrow = alias.Arrow
	node.Dir = alias.Dir
	node.Value, err = UnmarshalJSONExpr(alias.Value)
	if err != nil {
		return err
	}
	return nil
}

func (node *DeclStmtNode) UnmarshalJSON(data []byte) error {
	var alias DeclStmtNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Decl, err = UnmarshalJSONDecl(alias.Decl)
	if err != nil {
		return err
	}
	return nil
}

func (node *LabeledStmtNode) UnmarshalJSON(data []byte) error {
	var alias LabeledStmtNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Label = alias.Label
	node.Colon = alias.Colon
	node.Stmt, err = UnmarshalJSONStmt(alias.Stmt)
	if err != nil {
		return err
	}
	return nil
}

func (node *ExprStmtNode) UnmarshalJSON(data []byte) error {
	var alias ExprStmtNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.X, err = UnmarshalJSONExpr(alias.X)
	if err != nil {
		return err
	}
	return nil
}

func (node *SendStmtNode) UnmarshalJSON(data []byte) error {
	var alias SendStmtNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Chan, err = UnmarshalJSONExpr(alias.Chan)
	if err != nil {
		return err
	}
	node.Arrow = alias.Arrow
	node.Value, err = UnmarshalJSONExpr(alias.Value)
	if err != nil {
		return err
	}
	return nil
}

func (node *IncDecStmtNode) UnmarshalJSON(data []byte) error {
	var alias IncDecStmtNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.X, err = UnmarshalJSONExpr(alias.X)
	if err != nil {
		return err
	}
	node.TokPos = alias.TokPos
	node.Tok = alias.Tok
	return nil
}

func (node *AssignStmtNode) UnmarshalJSON(data []byte) error {
	var alias AssignStmtNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Lhs, err = UnmarshalJSONExprs(alias.Lhs)
	if err != nil {
		return err
	}
	node.TokPos = alias.TokPos
	node.Tok = alias.Tok
	node.Rhs, err = UnmarshalJSONExprs(alias.Rhs)
	if err != nil {
		return err
	}
	return nil
}

func (node *ReturnStmtNode) UnmarshalJSON(data []byte) error {
	var alias ReturnStmtNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Return = alias.Return
	node.Results, err = UnmarshalJSONExprs(alias.Results)
	if err != nil {
		return err
	}
	return nil
}

func (node *BlockStmtNode) UnmarshalJSON(data []byte) error {
	var alias BlockStmtNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Lbrace = alias.Lbrace
	node.List, err = UnmarshalJSONStmts(alias.List)
	if err != nil {
		return err
	}
	node.Rbrace = alias.Rbrace
	return nil
}

func (node *IfStmtNode) UnmarshalJSON(data []byte) error {
	var alias IfStmtNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.If = alias.If
	node.Init, err = UnmarshalJSONStmt(alias.Init)
	if err != nil {
		return err
	}
	node.Cond, err = UnmarshalJSONExpr(alias.Cond)
	if err != nil {
		return err
	}
	node.Body = alias.Body
	if err != nil {
		return err
	}
	node.Else, err = UnmarshalJSONStmt(alias.Else)
	if err != nil {
		return err
	}
	return nil
}

func (node *CaseClauseNode) UnmarshalJSON(data []byte) error {
	var alias CaseClauseNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Case = alias.Case
	node.List, err = UnmarshalJSONExprs(alias.List)
	if err != nil {
		return err
	}
	node.Colon = alias.Colon
	node.Body, err = UnmarshalJSONStmts(alias.Body)
	if err != nil {
		return err
	}
	return nil
}

func (node *SwitchStmtNode) UnmarshalJSON(data []byte) error {
	var alias SwitchStmtNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Switch = alias.Switch
	node.Init, err = UnmarshalJSONStmt(alias.Init)
	if err != nil {
		return err
	}
	node.Tag, err = UnmarshalJSONExpr(alias.Tag)
	if err != nil {
		return err
	}
	node.Body = alias.Body
	return nil
}

func (node *TypeSwitchStmtNode) UnmarshalJSON(data []byte) error {
	var alias TypeSwitchStmtNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Switch = alias.Switch
	node.Init, err = UnmarshalJSONStmt(alias.Init)
	if err != nil {
		return err
	}
	node.Assign, err = UnmarshalJSONStmt(alias.Assign)
	if err != nil {
		return err
	}
	node.Body = alias.Body
	return nil
}

func (node *CommClauseNode) UnmarshalJSON(data []byte) error {
	var alias CommClauseNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Case = alias.Case
	node.Comm, err = UnmarshalJSONStmt(alias.Comm)
	if err != nil {
		return err
	}
	node.Colon = alias.Colon
	node.Body, err = UnmarshalJSONStmts(alias.Body)
	if err != nil {
		return err
	}
	return nil
}

func (node *ForStmtNode) UnmarshalJSON(data []byte) error {
	var alias ForStmtNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.For = alias.For
	node.Init, err = UnmarshalJSONStmt(alias.Init)
	if err != nil {
		return err
	}
	node.Cond, err = UnmarshalJSONExpr(alias.Cond)
	if err != nil {
		return err
	}
	node.Post, err = UnmarshalJSONStmt(alias.Post)
	if err != nil {
		return err
	}
	node.Body = alias.Body
	return nil
}

func (node *RangeStmtNode) UnmarshalJSON(data []byte) error {
	var alias RangeStmtNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.For = alias.For
	node.Key, err = UnmarshalJSONExpr(alias.Key)
	if err != nil {
		return err
	}
	node.Value, err = UnmarshalJSONExpr(alias.Value)
	if err != nil {
		return err
	}
	node.TokPos = alias.TokPos
	node.Tok = alias.Tok
	node.X, err = UnmarshalJSONExpr(alias.X)
	if err != nil {
		return err
	}
	node.Body = alias.Body
	return nil
}

func (node *ValueSpecNode) UnmarshalJSON(data []byte) error {
	var alias ValueSpecNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Doc = alias.Doc
	node.Names = alias.Names
	node.Type, err = UnmarshalJSONExpr(alias.Type)
	if err != nil {
		return err
	}
	node.Values, err = UnmarshalJSONExprs(alias.Values)
	if err != nil {
		return err
	}
	node.Comment = alias.Comment
	return nil
}

func (node *TypeSpecNode) UnmarshalJSON(data []byte) error {
	var alias TypeSpecNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Doc = alias.Doc
	node.Name = alias.Name
	node.TypeParams = alias.TypeParams
	node.Assign = alias.Assign
	node.Type, err = UnmarshalJSONExpr(alias.Type)
	if err != nil {
		return err
	}
	node.Comment = alias.Comment
	return nil
}

func (node *GenDeclNode) UnmarshalJSON(data []byte) error {
	var alias GenDeclNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Doc = alias.Doc
	node.TokPos = alias.TokPos
	node.Tok = alias.Tok
	node.Lparen = alias.Lparen
	node.Specs, err = UnmarshalJSONSpecs(alias.Specs)
	if err != nil {
		return err
	}
	node.Rparen = alias.Rparen
	return nil
}

func (node *FileNode) UnmarshalJSON(data []byte) error {
	var alias FileNodeAlias
	err := json.Unmarshal(data, &alias)
	if err != nil {
		return err
	}

	node.NodeType = alias.NodeType
	node.Doc = alias.Doc
	node.Package = alias.Package
	node.Name = alias.Name
	node.Decls, err = UnmarshalJSONDecls(alias.Decls)
	if err != nil {
		return err
	}
	node.Imports = alias.Imports
	node.Unresolved = alias.Unresolved
	node.Comments = alias.Comments
	return nil
}
