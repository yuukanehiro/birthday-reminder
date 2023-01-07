package Response

type Error struct {
  Message string `json:"message"`
  Property string `json:"property"`
}
