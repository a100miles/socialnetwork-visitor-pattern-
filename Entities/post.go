package Entities

import "MiniSocial/Visitors"

type Post struct {
    Title   string
    Likes   int
    Tags    []string
    Content string
}

func (p *Post) Accept(visitor Visitors.IVisitor) {
    visitor.VisitPost(p)
}
