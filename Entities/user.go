package Entities

import "MiniSocial/Visitors"

type User struct {
    Name  string
    Posts []Post
}

func (u *User) Accept(visitor Visitors.IVisitor) {
    visitor.VisitUser(u)
}
