package Entities

import "MiniSocial/Visitors"

type Comment struct {
    Text  string
    Likes int
}

func (c *Comment) Accept(visitor Visitors.IVisitor) {
    visitor.VisitComment(c)
}
