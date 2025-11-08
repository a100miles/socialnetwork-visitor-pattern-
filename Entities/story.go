package Entities

import "MiniSocial/Core"

type Story struct {
    Views int
    Title string
}

func (s *Story) Accept(visitor Core.IVisitor) {
    visitor.VisitStory(s)
}
