package Entities

import "MiniSocial/Core"

type User struct {
    Name  string
    Posts []Post
}

func (u *User) Accept(visitor Core.IVisitor) {
    visitor.VisitUser(u)
}
