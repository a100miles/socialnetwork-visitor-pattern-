package Entities

import "MiniSocial/Core"

type Post struct {
    Title   string
    Likes   int
    Tags    []string
    Content string
}

func (p *Post) Accept(visitor Core.IVisitor) {
    visitor.VisitPost(p)
}
