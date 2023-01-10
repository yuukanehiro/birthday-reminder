package User

type UserId struct {
  value int64
}

// construct
func NewUserId(id int64) UserId {
  return UserId{
    value: id,
  }
}

func (u UserId) GetValue() int64 {
  return u.value
}
