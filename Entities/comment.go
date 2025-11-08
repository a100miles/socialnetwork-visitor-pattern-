package Entities

import "MiniSocial/Core"

type Comment struct {
    Text  string
    Likes int
}

func (c *Comment) Accept(visitor Core.IVisitor) {
    visitor.VisitComment(c)
}
