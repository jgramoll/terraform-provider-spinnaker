package client

import (
  // "fmt"
)

type Pipeline struct {
  Id    string
  Name  string
}

// func (pipeline *Pipeline) String() string {
//   return fmt.Sprintf("pipeline %s", pipeline.id)
// }

func (client *Client) GetPipeline(path string) (*Pipeline, error) {
  // client.NewRequest("get", path)

  return &Pipeline{}, nil
}
