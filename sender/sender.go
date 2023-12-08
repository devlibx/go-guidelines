package sender

import (
	"context"
	"fmt"
)

type Service struct {
}

func (s *Service) Send(ctx context.Context, data map[string]interface{}) error {
	fmt.Println("Sending data to outside world", data)
	return nil
}
