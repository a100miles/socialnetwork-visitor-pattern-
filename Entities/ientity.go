package Entities

import "MiniSocial/Visitors"

type IEntity interface {
    Accept(visitor Visitors.IVisitor)
}
