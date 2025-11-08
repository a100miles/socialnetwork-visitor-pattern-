package Entities

import "MiniSocial/Visitors"

type Story struct {
    Views int
    Title string
}

func (s *Story) Accept(visitor Visitors.IVisitor) {
    visitor.VisitStory(s)
}
