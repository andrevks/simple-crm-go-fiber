package models

import (
)

type Model interface {
	GetID() uint
	SetID(uint)
}
