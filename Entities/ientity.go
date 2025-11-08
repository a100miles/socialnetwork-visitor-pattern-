package Entities

import "MiniSocial/Core"

type IEntity interface {
    Accept(visitor Core.IVisitor)
}
